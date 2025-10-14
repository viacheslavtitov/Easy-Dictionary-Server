package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	domain "easy-dictionary-server/domain"
	domainAuth "easy-dictionary-server/domain"
	dictionaryDomain "easy-dictionary-server/domain/dictionary"
	languageDomain "easy-dictionary-server/domain/language"
	userDomain "easy-dictionary-server/domain/user"
)

type Session struct {
	BaseURL     string
	BearerToken string
	HTTP        *http.Client
	Timeout     time.Duration
}

func NewSession() *Session {
	return &Session{
		BaseURL: "http://127.0.0.1:8080/api",
		HTTP:    &http.Client{Timeout: 10 * time.Second},
		Timeout: 10 * time.Second,
	}
}

var in = bufio.NewReader(os.Stdin)

func ask(prompt string, def string) string {
	fmt.Printf("%s", prompt)
	text, _ := in.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		return def
	}
	return text
}

func askInt(prompt string, def int) int {
	for {
		s := ask(fmt.Sprintf("%s", prompt), fmt.Sprintf("%d", def))
		v, err := strconv.Atoi(strings.TrimSpace(s))
		if err == nil {
			return v
		}
		fmt.Println("  ✖ Input number, please.")
	}
}

func askDuration(prompt string, def time.Duration) time.Duration {
	for {
		s := ask(fmt.Sprintf("%s", prompt), def.String())
		d, err := time.ParseDuration(strings.TrimSpace(s))
		if err == nil {
			return d
		}
		fmt.Println("  ✖ Duration time, example.: 200ms, 1s, 2s.")
	}
}

func doJSON(ctx context.Context, s *Session, method, url string, reqBody any, out any) error {
	var body io.Reader
	if reqBody != nil {
		b, _ := json.Marshal(reqBody)
		body = bytes.NewReader(b)
		br, err := json.MarshalIndent(reqBody, "", "  ")
		if err != nil {
			log.Println("marshal:", err)
		} else {
			fmt.Printf("Request:\n%s\n", br)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if s.BearerToken != "" {
		bt := s.BearerToken
		if !strings.HasPrefix(strings.ToLower(bt), "bearer ") {
			bt = "Bearer " + bt
		}
		req.Header.Set("Authorization", bt)
	}

	resp, err := s.HTTP.Do(req)
	if resp != nil {
		fmt.Printf("\nstatus=%d\n", resp.StatusCode)
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var apiErr domain.ErrorResponse
		if json.Unmarshal(b, &apiErr) == nil && apiErr.Message != "" {
			return fmt.Errorf("status=%d api_error=%s", resp.StatusCode, apiErr.Message)
		}
		return fmt.Errorf("status=%d body=%s", resp.StatusCode, string(b))
	}

	if out != nil && len(b) > 0 {
		if err := json.Unmarshal(b, out); err != nil {
			return fmt.Errorf("decode: %w, body=%s", err, string(b))
		}
	}
	return nil
}

type menuItem struct {
	Key   string
	Title string
	Do    func(*Session)
}

func actionSetConfig(s *Session) {
	base := ask(fmt.Sprintf("Base URL [%s]: ", s.BaseURL), s.BaseURL)
	token := ask("Bearer token (enter to skip): ", s.BearerToken)
	timeout := askDuration("HTTP timeout (e.g. 10s) [10s]: ", s.Timeout)

	s.BaseURL = strings.TrimRight(base, "/")
	s.BearerToken = strings.TrimSpace(token)
	s.Timeout = timeout
	s.HTTP.Timeout = timeout

	fmt.Println("✓ Set up finished.")
}

func actionCreateDictionary(s *Session) {
	langFromId := askInt("Language from id: ", 0)
	langToId := askInt("Language to id: ", 0)
	dialect := ask("Dialect  (enter to skip): ", "")
	if err := doJSON(context.Background(), s, http.MethodPost, s.BaseURL+"/dictionary/create",
		dictionaryDomain.DictionaryRequest{LangFromId: langFromId, LangToId: langToId, Dialect: &dialect}, nil); err != nil {
		fmt.Println("✖ Failed:", err)
		return
	}
	fmt.Printf("✓ Dictionary created")
}

func actionGetAllShortDictionariesForUser(s *Session) {
	var resp *[]dictionaryDomain.DetailShortDictionary
	if err := doJSON(context.Background(), s, http.MethodGet, s.BaseURL+"/dictionary/all/short", nil, resp); err != nil {
		fmt.Println("✖ Failed:", err)
		return
	}
	if resp != nil && len(*resp) > 0 {
		b, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			log.Println("marshal:", err)
		} else {
			fmt.Printf("✓ Dictionaries:\n%s\n", b)
		}
	} else {
		fmt.Printf("✓ User doesn't have any dictionaries")
	}
}

func actionCreateUser(s *Session) {
	email := ask("Email: ", "")
	firstName := ask("First Name: ", "")
	lastName := ask("Last Name: ", "")
	password := ask("Password: ", "")
	provider := ask("provider: ", "")
	var resp userDomain.User
	if err := doJSON(context.Background(), s, http.MethodPost, s.BaseURL+"/signup", userDomain.RegisterUserRequest{
		Email:         email,
		Provider:      provider,
		Password:      password,
		FirstName:     firstName,
		ProviderToken: "",
		LastName:      lastName}, &resp); err != nil {
		fmt.Println("✖ Failed:", err)
		return
	}
	fmt.Printf("✓ User created")
}

func actionGetToken(s *Session) {
	email := ask("Email: ", "")
	password := ask("Password: ", "")
	provider := ask("Provider  (enter to skip):", "email")
	var resp string
	if err := doJSON(context.Background(), s, http.MethodPost, s.BaseURL+"/signin", domainAuth.AuthRequest{
		Email:         email,
		Provider:      provider,
		Password:      password,
		ProviderToken: ""}, &resp); err != nil {
		fmt.Println("✖ Failed:", err)
		return
	}
	fmt.Printf("Baearer Token: %s", resp)
}

func actionCreateLanguage(s *Session) {
	name := ask("Name: ", "")
	code := ask("Code: ", "")
	var resp languageDomain.Language
	if err := doJSON(context.Background(), s, http.MethodPost, s.BaseURL+"/languages/create", languageDomain.LanguageRequest{
		Name: name,
		Code: &code}, &resp); err != nil {
		fmt.Println("✖ Failed:", err)
		return
	}
	fmt.Printf("✓ Language created: \n%s", resp)
}

func actionGetAllLanguagesForUser(s *Session) {
	var resp *[]languageDomain.Language
	if err := doJSON(context.Background(), s, http.MethodGet, s.BaseURL+"/languages/all", nil, &resp); err != nil {
		fmt.Println("✖ Failed:", err)
		return
	}
	if resp != nil && len(*resp) > 0 {
		b, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			log.Println("marshal:", err)
		} else {
			fmt.Printf("✓ Languages:\n%s\n", b)
		}
	} else {
		fmt.Printf("✓ User doesn't have any languages")
	}
}

func main() {
	s := NewSession()

	items := []menuItem{
		{"1", "Set base URL / token / timeout", actionSetConfig},
		{"2", "Create user", actionCreateUser},
		{"3", "Get access token", actionGetToken},
		{"4", "Get user languages", actionGetAllLanguagesForUser},
		{"5", "Create language", actionCreateLanguage},
		{"6", "Get user dictionaries", actionGetAllShortDictionariesForUser},
		{"7", "Create dictionary", actionCreateDictionary},
		{"0", "Exit", nil},
	}
	for {
		fmt.Println("\n=== EasyDictionary Seeder ===")
		fmt.Printf("Base: %s\n", s.BaseURL)
		if s.BearerToken != "" {
			fmt.Println("Auth: (set)")
		} else {
			fmt.Println("Auth: (none)")
		}
		fmt.Println("-----------------------------")
		for _, it := range items {
			fmt.Printf("%s) %s\n", it.Key, it.Title)
		}
		choice := ask("Select: ", "")
		found := false
		for _, it := range items {
			if it.Key == choice {
				found = true
				if it.Do == nil {
					fmt.Println("Bye!")
					return
				}
				it.Do(s)
				break
			}
		}
		if !found {
			fmt.Println("✖ Unreachable choice.")
		}
	}
}
