package requests

import (
	"encoding/json"
	"fmt"
	"form3app/pkg/api/exceptions"
	"form3app/pkg/api/inputs"
	"form3app/pkg/api/outputs"
	"net/http"
)

const (
	deleteAccountUri = "/v1/organisation/accounts/"
)

// deleteAccountRequest the child struct of Request and represents delete account operation
type deleteAccountRequest struct {
	Request
}

// NewDeleteAccountRequest the constructor method of deleteAccountRequest
func NewDeleteAccountRequest(baseUrl string) *deleteAccountRequest {
	return &deleteAccountRequest{
		Request{
			url: baseUrl + deleteAccountUri,
		},
	}
}

// Execute will make the request, returning error if errors are encountered,
// otherwise expected success response will be retrieved
func (c deleteAccountRequest) Execute(input *inputs.DeleteAccountInput) (*outputs.DeleteAccountOutput, error) {
	resp, err := http.Get(c.url + c.buildQueryString(input))

	if err != nil {
		return nil, exceptions.RequestFailedException{}
	}

	if resp.StatusCode == http.StatusConflict {
		return nil, exceptions.VersionIncorrectException{}
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, exceptions.NotFoundException{}
	}

	var respData outputs.DeleteAccountOutput
	err = json.NewDecoder(resp.Body).Decode(&respData)

	if err != nil {
		return nil, exceptions.InvalidResponseException{}
	}

	return &respData, nil
}

// buildQueryString is a private method for generating proper query string from input.
func (c deleteAccountRequest) buildQueryString(input *inputs.DeleteAccountInput) string {
	return fmt.Sprintf("/%s?version=%d", input.AccountId, input.Version)
}
