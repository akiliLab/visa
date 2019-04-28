package visa

import (
	"bytes"
	"encoding/json"
	"fmt"
	pb "github.com/ubunifupay/visa/pb"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetForexChangeRate : get foreignExchange
func (c *Client) GetForexChangeRate(transactionID string, request pb.VisaForexRequest) (*pb.VisaForexReply, error) {

	urlFull := &url.URL{
		Scheme: "https",
		Host:   c.BaseURL,
		Path:   "/forexrates/v1/foreignexchangerates",
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", urlFull.String(), bytes.NewBuffer(body))
	req.Header.Set("X-Client-Transaction-ID", transactionID)
	req.Header.Set("Authorization", c.Authorization)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Could not load HTTPS client: %v", err)
	}
	defer resp.Body.Close()

	// HTTP 200 and 404 are not errors
	serializedBody, _ := ioutil.ReadAll(resp.Body)

	response := &pb.VisaForexReply{}

	jsonErr := json.Unmarshal(serializedBody, response)

	if jsonErr != nil {
		fmt.Println(err)
	}

	switch resp.StatusCode {
	case 200, 202, 404:
		return response, nil
	default:
		return nil, fmt.Errorf("HTTP error: %s, %v", resp.Status, string(serializedBody))
	}
}
