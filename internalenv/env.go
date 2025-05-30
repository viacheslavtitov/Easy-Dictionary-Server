package internalenv

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
	AppEnv            string `json:"app_env"`
	JwtExpTimeMinutes int    `json:"JWT_EXP_TIME_MINUTES"`
	JwtSecret         string `json:"JWT_SECRET"`
	ServerAddress     string `json:"address"`
	ServerPort        string `json:"port"`
	TimeOut           int    `json:"timeout"`
	DBName            string `json:"db_name"`
	DBHost            string `json:"db_host"`
	DBPort            int    `json:"db_port"`
	DBUser            string `json:"db_user"`
	DBPassword        string `json:"db_password"`
}

type EnvInteface interface {
	CombineServerAddress() string
}

func (env *Env) CombineServerAddress() string {
	return env.ServerAddress + ":" + env.ServerPort
}

const (
	envName           = "app_env"
	jwtExpTimeMinutes = "JWT_EXP_TIME_MINUTES"
	jwtSecret         = "JWT_SECRET"
	serverAddress     = "address"
	serverPort        = "port"
	timeOut           = "timeout"
	dbname            = "db_name"
	dbhost            = "db_host"
	dbport            = "db_port"
	dbuser            = "db_user"
	dbpassword        = "db_password"
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

	env := parseEnv()
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
		newEnv := parseEnv()
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

func parseEnv() Env {
	return Env{AppEnv: viper.GetString(envName),
		JwtExpTimeMinutes: viper.GetInt(jwtExpTimeMinutes),
		JwtSecret:         viper.GetString(jwtSecret),
		ServerAddress:     viper.GetString(serverAddress),
		ServerPort:        viper.GetString(serverPort),
		TimeOut:           viper.GetInt(timeOut),
		DBName:            viper.GetString(dbname),
		DBHost:            viper.GetString(dbhost),
		DBPort:            viper.GetInt(dbport),
		DBUser:            viper.GetString(dbuser),
		DBPassword:        viper.GetString(dbpassword)}
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
