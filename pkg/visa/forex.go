package visa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ForexRequest request definition
type ForexRequest struct {
	CardAcceptor             CardAcceptor `json:"cardAcceptor"`
	DestinationCurrencyCode  string       `json:"destinationCurrencyCode"`
	MarkUpRate               string       `json:"markUpRate"`
	RetrievalReferenceNumber string       `json:"retrievalReferenceNumber"`
	SourceAmount             string       `json:"sourceAmount"`
	SourceCurrencyCode       string       `json:"sourceCurrencyCode"`
	SystemsTraceAuditNumber  string       `json:"systemsTraceAuditNumber"`
}

// Address card address
type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
	County  string `json:"county"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}

// CardAcceptor id
type CardAcceptor struct {
	Address    Address `json:"address"`
	IDCode     string  `json:"idCode"`
	Name       string  `json:"name"`
	TerminalID string  `json:"terminalId"`
}

// ForexRequestResponse for ForexRequestResponse
type ForexRequestResponse struct {
	ConversionRate               float64
	DestinationAmount            string
	MarkUpRateApplied            string
	OriginalDestnAmtBeforeMarkUp string
}

// GetForexChangeRate : get foreignExchange
func (c *Client) GetForexChangeRate(transactionID string, request *ForexRequest) (response []byte, err error) {

	urlFull := &url.URL{
		Scheme: "https",
		Host:   c.BaseURL,
		Path:   "/forexrates/v1/foreignexchangerates",
	}

	body, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	req, err := http.NewRequest("POST", urlFull.String(), bytes.NewBuffer(body))
	req.Header.Set("X-Client-Transaction-ID", transactionID)
	req.Header.Set("Authorization", c.Authorization)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HttpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Could not load HTTPS client: %v", err)
	}
	defer resp.Body.Close()

	// HTTP 200 and 404 are not errors
	response, _ = ioutil.ReadAll(resp.Body)

	switch resp.StatusCode {
	case 200, 202, 404:
		return response, nil
	default:
		return nil, fmt.Errorf("HTTP error: %s, %v", resp.Status, string(response))
	}
}
