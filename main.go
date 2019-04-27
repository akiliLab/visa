package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ubunifupay/visa/pkg/visa"
)

var client *visa.Client

func main() {

	log.Println("Service has started")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// apiKey := os.Getenv("API_KEY")
	// sharedSecret := os.Getenv("SHARED_SECRET")
	userID := os.Getenv("USER_ID")
	password := os.Getenv("PASSWORD")
	certPath := os.Getenv("CERT")
	keyPath := os.Getenv("KEY")
	caPath := os.Getenv("CA")
	cert, _ := tls.LoadX509KeyPair(certPath, keyPath)
	ca, _ := ioutil.ReadFile(caPath)
	client, _ = visa.NewClient(userID, password, cert, ca, visa.SANDBOX)

}
