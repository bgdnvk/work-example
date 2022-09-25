package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/go-chi/chi/v5"
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
	user := cfg.User

	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World the user is " + user))
		if err != nil {
			log.Fatal(err)
		}
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
	if err != nil {
		log.Fatal(err)
	}

}

//load the env file
//env arg is the prefix inside the env file ie dev_user
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
