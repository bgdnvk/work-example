package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Cfg struct {
	User string
	Port int
}

func main() {
	cfg := *loadCFG("dev")

	fmt.Printf("my cfg: %+v", cfg)

}

func loadCFG(env string) *Cfg {
	var cfg Cfg
	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		log.Fatal("Error loading .env file")
	}

	err2 := envconfig.Process(env, &cfg)
	if err2 != nil {
		log.Fatal(dotEnvErr.Error())
	}

	return &cfg
}
