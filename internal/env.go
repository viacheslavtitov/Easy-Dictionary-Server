package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	firebase "firebase.google.com/go"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

type Env struct {
	AppEnv         string
	ServerAddress  string
	ContextTimeout int
	DBHost         string
	DBPort         string
	DBUser         string
	DBPass         string
	DBName         string
}

func LoadEnv(environment string) *Env {
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
	defer client.Close()
	if err != nil {
		log.Default().Fatal("Couldn't init", err)
		return nil
	}
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
	}
	env := Env{AppEnv: viper.GetString("APP_ENV")}
	switch env.AppEnv {
	case "local":
		{
			log.Default().Println("The App is running in local env")
		}
	case "development":
		{
			log.Default().Println("The App is running in development env")
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
	return &env
}
