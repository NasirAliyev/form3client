package requests

import (
	"bytes"
	"encoding/json"
	"form3app/pkg/api/exceptions"
	"form3app/pkg/api/inputs"
	"form3app/pkg/api/outputs"
	"net/http"
)

const (
	createAccountRequestContentType = "application/vnd.api+json"
	createAccountUri                = "/v1/organisation/accounts"
)

// createAccountRequest the child struct of Request and represents create account operation
type createAccountRequest struct {
	Request
}

// NewCreateAccountRequest the constructor method of createAccountRequest
func NewCreateAccountRequest(baseUrl string) *createAccountRequest {
	return &createAccountRequest{
		Request{
			url:         baseUrl + createAccountUri,
			contentType: createAccountRequestContentType,
		},
	}
}

// Execute will make the request, returning error if errors are encountered,
// otherwise expected success response will be retrieved
func (c createAccountRequest) Execute(input *inputs.CreateAccountInput) (*outputs.CreateAccountOutput, error) {
	jsonInput, err := json.Marshal(input)

	if err != nil {
		return nil, exceptions.InvalidInputException{}
	}

	resp, err := http.Post(
		c.url,
		c.contentType,
		bytes.NewBuffer(jsonInput),
	)

	if err != nil {
		return nil, exceptions.RequestFailedException{}
	}

	if resp.StatusCode == http.StatusConflict {
		return nil, exceptions.AccountAlreadyExistsException{}
	}

	var respData outputs.CreateAccountOutput
	err = json.NewDecoder(resp.Body).Decode(&respData)

	if err != nil {
		return nil, exceptions.InvalidResponseException{}
	}

	return &respData, nil
}
