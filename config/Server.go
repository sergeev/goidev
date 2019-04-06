package config

import (
	"log"
	"net/http"
	"time"
)

// server configs
func ServerConfig(response http.Response, request *http.Request){
	s := http.Server{
		Addr:           ":8000",
		//Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// Test log server witch fatal error
	log.Fatal(s.ListenAndServe())
}
