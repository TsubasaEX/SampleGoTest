package tzoo

import (
	"fmt"
	"net/http"
)

type ZooClient interface {
	ListAnimalFacts(q AnimalFactsQuery) (*AnimalFactsResponse, error)
}

type ZooHTTPClient struct {
	baseURL string
	client  *http.Client
}

// Serialization of error response from zoo API service
type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (res *ErrorResponse) Error() string {
	return fmt.Sprintf("got %d error: %s", res.StatusCode, res.Message)
}

type AnimalFactsQuery struct {
	AnimalName string
	PageToken  string
}

func (c *ZooHTTPClient) ListAnimalFacts(q AnimalFactsQuery) (*AnimalFactsResponse, error) {
	// HTTP implementation here; returns an
	// AnimalFactsResponse if the HTTP request succeeds,
	// or an error, of type ErrorResponse, if the request
	// gets a non-2xx HTTP status code.

	// returning nil, nil just so the code compiles
	return nil, nil
}

type AnimalFactsResponse struct {
	Facts         []string `json:"facts"`
	AreThereMore  bool     `json:"are_there_more"`
	NextPageToken string   `json:"next_page_token"`
}
