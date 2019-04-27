package visa

import (
	"crypto/tls"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestForexConvesion(t *testing.T) {

	err := godotenv.Load()

	if err != nil {
		t.Errorf("Error loading variables: %v", err)
	}

	testData := &ForexRequest{
		DestinationCurrencyCode:  "840",
		MarkUpRate:               "1",
		RetrievalReferenceNumber: "201010101031",
		SourceAmount:             "100",
		SourceCurrencyCode:       "643",
		SystemsTraceAuditNumber:  "350421",
		CardAcceptor: CardAcceptor{
			Address: Address{
				City:    "Foster City",
				Country: "RU",
				County:  "San Mateo",
				State:   "CA",
				ZipCode: "94404",
			},
			IDCode:     "ABCD1234ABCD123",
			Name:       "ABCD",
			TerminalID: "ABCD1234",
		},
	}

	userID := os.Getenv("USER_ID")
	password := os.Getenv("PASSWORD")
	certPath := os.Getenv("CERT")
	keyPath := os.Getenv("KEY")
	caPath := os.Getenv("CA")
	cert, _ := tls.LoadX509KeyPair(certPath, keyPath)
	ca, _ := ioutil.ReadFile(caPath)

	client, _ := NewClient(userID, password, cert, ca, SANDBOX)

	response, err := client.GetForexChangeRate("EYaCiuJW6LS", testData)

	if err != nil {
		t.Errorf("Error when getting response: %v", err)
	}

	if reflect.TypeOf(response).String() != "visa.ForexRequestResponse" {
		t.Errorf("Return should be of type ForexRequestResponse. Looking for %s, got %s", "visa.ForexRequestResponse", reflect.TypeOf(response).String())
	}

}
