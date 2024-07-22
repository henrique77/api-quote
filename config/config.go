package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port                string
	DBHost              string
	DBUserName          string
	DBPassword          string
	DBPort              string
	DBName              string
	UrlAPIFreteRapido   string
	RegisteredNumber    string
	TokenAPIFreteRapido string
	PlatformCode        string
}

func ReadEnvs() *Env {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}
	return &Env{
		Port:                os.Getenv("PORT"),
		DBHost:              os.Getenv("DB_HOST"),
		DBUserName:          os.Getenv("DB_USERNAME"),
		DBPassword:          os.Getenv("DB_PASSWORD"),
		DBPort:              os.Getenv("DB_PORT"),
		DBName:              os.Getenv("DB_DATABASE"),
		UrlAPIFreteRapido:   os.Getenv("URL_API_FRETE_RAPIDO"),
		RegisteredNumber:    os.Getenv("REGISTERED_NUMBER"),
		TokenAPIFreteRapido: os.Getenv("TOKEN_API_FRETE_RAPIDO"),
		PlatformCode:        os.Getenv("PLATFORM_CODE"),
	}
}
