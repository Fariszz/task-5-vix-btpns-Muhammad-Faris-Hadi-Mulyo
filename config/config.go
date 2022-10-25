package config

import "GOlangRakamin/helpers"

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Configuration {

	configuration := Configuration{
		DB_USERNAME: helpers.GoDotEnv("DB_USERNAME"),
		DB_PASSWORD: helpers.GoDotEnv("DB_PASSWORD"),
		DB_PORT:     helpers.GoDotEnv("DB_PORT"),
		DB_HOST:     helpers.GoDotEnv("DB_HOST"),
		DB_NAME:     helpers.GoDotEnv("DB_NAME"),
	}

	return configuration
}