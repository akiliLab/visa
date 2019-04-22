package visa

// Address contains the full address of a merchant
type Address struct {
	City    string
	Country string
	County  string
	State   string
	ZipCode string
}

// CardAcceptor for card Acceptor
type CardAcceptor struct {
	IDCode     string
	Name       string
	TerminalID string
	Address
}

// ForexRequest for  ForexRequest
type ForexRequest struct {
	DestinationCurrencyCode  string
	MarkUpRate               string
	RetrievalReferenceNumber string
	SourceAmount             string
	SourceCurrencyCode       string
	SystemsTraceAuditNumber  string
	CardAcceptor
}

// ForexRequestResponse for ForexRequestResponse
type ForexRequestResponse struct {
	ConversionRate               float64
	DestinationAmount            string
	MarkUpRateApplied            string
	OriginalDestnAmtBeforeMarkUp string
}

// GetForexChangeRate : get foreignExchange
func (c *Client) GetForexChangeRate(request *ForexRequest) (*ForexRequestResponse, error) {

	return nil, nil
}
