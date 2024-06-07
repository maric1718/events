package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// environment variables
type (
	ENV struct {
		App   *App
		DB    *DB
		HTTP  *HTTP
		Kafka *Kafka
	}

	App struct {
		Name string
		Env  string
	}

	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}

	HTTP struct {
		URL            string
		Port           string
		AllowedOrigins []string
		AllowedHeaders []string
	}

	Kafka struct {
		Brokers []string
		Timeout string
		Topics  []string
	}
)

func New() (*ENV, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Name:       os.Getenv("DB_NAME"),
	}

	http := &HTTP{
		URL:            os.Getenv("HTTP_URL"),
		Port:           os.Getenv("HTTP_PORT"),
		AllowedOrigins: strings.Split(os.Getenv("HTTP_ALLOWED_ORIGINS"), ","),
		AllowedHeaders: strings.Split(os.Getenv("HTTP_ALLOWED_HEADERS"), ","),
	}

	kafka := &Kafka{
		Brokers: strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
		Timeout: os.Getenv("KAFKA_TIMEOUT"),
		Topics:  strings.Split(os.Getenv("KAFKA_TOPICS"), ","),
	}

	return &ENV{
		app,
		db,
		http,
		kafka,
	}, nil
}
