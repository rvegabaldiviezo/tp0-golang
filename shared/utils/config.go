package utils

import (
	"encoding/json"
	"log"
	"os"
)

// type Config struct {
// 	Ip      string `json:"ip"`
// 	Puerto  int    `json:"puerto"`
// 	Mensaje string `json:"mensaje"`
// }

// var ClientConfig *Config

// func IniciarConfiguracion(filePath string) *Config {
// 	var config *Config
// 	configFile, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	defer configFile.Close()
// 	jsonParser := json.NewDecoder(configFile)
// 	jsonParser.Decode(&config)
// 	return config
// }

func IniciarConfiguracion[T any]() *T {
	filePath := "configs/config.json"
	var config T
	configFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("error abriendo archivo de config: %s", err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		log.Fatalf("error decodificando config: %s", err.Error())
	}

	return &config
}
