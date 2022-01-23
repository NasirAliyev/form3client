package requests

import (
	"encoding/json"
	"form3app/pkg/api/exceptions"
	"form3app/pkg/api/inputs"
	"form3app/pkg/api/outputs"
	"net/http"
)

const (
	fetchAccountUri = "/v1/organisation/accounts/"
)

// fetchAccountRequest the child struct of Request and represents fetch account operation
type fetchAccountRequest struct {
	Request
}

// NewFetchAccountRequest the constructor method of fetchAccountRequest
func NewFetchAccountRequest(baseUrl string) *fetchAccountRequest {
	return &fetchAccountRequest{
		Request{
			url: baseUrl + fetchAccountUri,
		},
	}
}

// Execute will make the request, returning error if errors are encountered,
// otherwise expected success response will be retrieved
func (c fetchAccountRequest) Execute(input *inputs.FetchAccountInput) (*outputs.FetchAccountOutput, error) {
	resp, err := http.Get(c.url + input.AccountId)

	if err != nil {
		return nil, exceptions.RequestFailedException{}
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, exceptions.NotFoundException{}
	}

	var respData outputs.FetchAccountOutput
	err = json.NewDecoder(resp.Body).Decode(&respData)

	if err != nil {
		return nil, exceptions.InvalidResponseException{}
	}

	return &respData, nil
}
