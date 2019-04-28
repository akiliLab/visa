package main

import (
	"context"
	"crypto/tls"
	"github.com/joho/godotenv"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	pb "github.com/ubunifupay/visa/pb"
	"github.com/ubunifupay/visa/server/visa"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var client *visa.Client

type server struct{}

func (s *server) GetForexConversion(ctx context.Context, request *pb.VisaForexRequest, response *pb.VisaForexReply) error {
	// TODO: use the exact value supplied
	transactionID := ""
	resp, err := client.GetForexChangeRate(transactionID, request)

	//TODO: Add logics to handle after this
	response = resp
	return err
}

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

	service := grpc.NewService(
		micro.Name("ubunifupay.visa"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	pb.RegisterVisaServiceHandler(service.Server(), new(server))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
