package main

import (
	"context"
	"crypto/tls"
	"github.com/joho/godotenv"
	pb "github.com/ubunifupay/visa/pb"
	"github.com/ubunifupay/visa/pkg/visa"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"os"
)

var client *visa.Client

type server struct{}

func (s *server) GetForexConversion(ctx context.Context, request *pb.VisaForexRequest) (*pb.VisaForexReply, error) {
	// TODO: use the exact value supplied
	transactionID := ""
	response, err := client.GetForexChangeRate(transactionID, request)

	//TODO: Add logics to handle after this

	return response, err
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

	lis, err := net.Listen("tcp", ":5088")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterVisaServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
