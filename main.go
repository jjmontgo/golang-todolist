package main

import (
	"fmt"
	"net/http"
	"os"
	// get config from .env when run locally
	localEnvFile "github.com/joho/godotenv"
	// map ApiGateway to http.ResponseWriter/Request
	lambdaGoServerAdapter "github.com/akrylysov/algnhsa"
)

func init() {
	router := InitRouter()
	// prod will not have a .env, so it can ignore errors
	_ = localEnvFile.Load()
	mode := os.Getenv("MODE")
	if mode == "prod" {
		// running lambda from api gateway
		lambdaGoServerAdapter.ListenAndServe(router, nil)
	} else if mode == "dev" {
		// running from local machine
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		fmt.Println("Listening on port " + port + "...");
		http.ListenAndServe(":"+port, router)
	} else {
		panic("Failed to determine application mode: 'prod' or 'dev'")
	}
}

func main() {
}




