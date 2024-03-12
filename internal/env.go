package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Env struct {
	AppEnv string `json:"app_env"`
	Test   string `json:"test"`
	// ServerAddress  string
	// ContextTimeout int
	// DBHost         string
	// DBPort         string
	// DBUser         string
	// DBPass         string
	// DBName         string
}

const (
	envName = "app_env"
)

func LoadEnv(environment string, changeEnvChan chan Env) *Env {
	// token := loadToekn()
	// if token == nil {
	// 	log.Default().Println("Token wasn't loaded")
	// 	return nil
	// }
	// log.Default().Printf("AccessToken = %s", token.AccessToken)

	log.Default().Printf("Load environment %s", environment)
	viper.AddRemoteProvider("firestore", "google-cloud-project-id", "collection/document")
	viper.SetConfigType("json")
	opt := option.WithCredentialsFile("../service-account.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Default().Fatal("Couldn't load env credentials file", err)
		return nil
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Default().Fatal("Couldn't init", err)
		return nil
	}
	// defer client.Close()
	doc, err := client.Collection("Environment").Doc(environment).Get(context.Background())
	if err != nil {
		log.Default().Fatal("Couldn't load config collection from Cloud Firestore", err)
		return nil
	}
	data := doc.Data()
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Default().Fatal("Couldn't encode JSON data", err)
		return nil
	}
	err = viper.ReadConfig(bytes.NewBuffer(jsonData))
	if err != nil {
		log.Default().Fatal("Couldn't read config file", err)
		return nil
	} else {
		log.Default().Println("Config was loaded")
		log.Default().Println(viper.AllKeys())
	}

	env := Env{AppEnv: viper.GetString(envName), Test: viper.GetString("test")}
	switch env.AppEnv {
	case "local":
		{
			log.Default().Println("The App is running in local env")
			log.Default().Printf("%s", jsonData)
		}
	case "development":
		{
			log.Default().Println("The App is running in development env")
			log.Default().Printf("%s", jsonData)
		}
	case "production":
		{
			log.Default().Println("The App is running in production env")
		}
	default:
		{
			log.Default().Fatal("Unrecognized environment")
			return nil
		}
	}
	streamChanges := client.Collection("Environment").Doc(environment).Snapshots(context.Background())
	go readNewConfig(streamChanges, client, changeEnvChan, &env)
	return &env
}

func readNewConfig(streamChanges *firestore.DocumentSnapshotIterator, client *firestore.Client, changeChan chan Env, currentEnv *Env) {
	for {
		snap, err := streamChanges.Next()
		if err != nil {
			if err == iterator.Done {
				continue
			}
			log.Default().Println("Failed to read next stream changes")
			log.Default().Println(err)
			continue
		}
		jsonData, err := json.Marshal(snap.Data())
		if err != nil {
			log.Default().Println(err)
			continue
		}
		log.Default().Println("Config was changed")
		log.Default().Printf("%s", jsonData)
		err = viper.ReadConfig(bytes.NewBuffer(jsonData))
		if err != nil {
			log.Default().Fatal("Couldn't read remote config", err)
			return
		} else {
			log.Default().Printf("Config was loaded keys = %d", len(viper.AllKeys()))
			log.Default().Println(viper.AllKeys())
			if len(viper.AllKeys()) == 0 {
				continue
			}
		}
		newEnv := Env{AppEnv: viper.GetString(envName), Test: viper.GetString("test")}
		if newEnv == *currentEnv {
			log.Default().Println("New config equals previous")
		} else {
			log.Default().Println("New config is not equal previous")
			changeChan <- newEnv
			break
		}
	}
	defer streamChanges.Stop()
	defer client.Close()
}

// func loadToekn() *oauth2.Token {
// 	var token *oauth2.Token
// 	ctx := context.Background()
// 	scopes := []string{
// 		"https://www.googleapis.com/auth/cloud-platform",
// 	}
// 	credentials, err := auth.FindDefaultCredentials(ctx, scopes...)
// 	if err == nil {
// 		log.Default().Printf("found default credentials. %v", credentials)
// 		token, err = credentials.TokenSource.Token()
// 		log.Default().Printf("token: %v, err: %v", token, err)
// 		if err != nil {
// 			log.Default().Println(err)
// 		}
// 	} else {
// 		log.Default().Println(err)
// 	}
// 	return token
// }
