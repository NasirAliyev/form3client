package form3

import (
	"form3app/pkg/api/inputs"
	"form3app/pkg/api/outputs"
	"form3app/pkg/api/requests"
)

type Config struct {
	BaseUrl string
}

type client struct {
	config Config
}

func NewClient(cnf Config) *client {
	return &client{
		config: cnf,
	}
}

// CreateAccount Api responsible for creation an account resource
//
// Returned Error Types:
//   * InvalidInputException
//   This exception is thrown when the request input created not properly
//
//   * RequestFailedException
//   This exception is thrown when the request failed on server side or
//   during the request.
//
//   * AccountAlreadyExistsException
//   This exception is thrown when the account already exists.
//
//   * InvalidResponseException
//   This exception is thrown when the response from the server is not valid
//
// See also, https://api-docs.form3.tech/api.html#organisation-accounts-create
func (c client) CreateAccount(input *inputs.CreateAccountInput) (*outputs.CreateAccountOutput, error) {
	req := requests.NewCreateAccountRequest(c.config.BaseUrl)

	return req.Execute(input)
}

// FetchAccount Api responsible for fetching the account resource
//
// Returned Error Types:
//   * RequestFailedException
//   This exception is thrown when the request failed on server side or
//   during the request.
//
//   * NotFoundException
//   This exception is thrown when the account cannot be found
//   by given input
//
//   * InvalidResponseException
//   This exception is thrown when the response from the server is not valid
//
// See also, https://api-docs.form3.tech/api.html#organisation-accounts-fetch
func (c client) FetchAccount(input *inputs.FetchAccountInput) (*outputs.FetchAccountOutput, error) {
	req := requests.NewFetchAccountRequest(c.config.BaseUrl)

	return req.Execute(input)
}

// DeleteAccount Api responsible for deletion the account resource
//
// Returned Error Types:
//   * RequestFailedException
//   This exception is thrown when the request failed on server side or
//   during the request.
//
//   * VersionIncorrectException
//   This exception is thrown when the account deletion request
//   processed with incorrect version
//
//   * NotFoundException
//   This exception is thrown when the account cannot be found
//   by given input
//
//   * InvalidResponseException
//   This exception is thrown when the response from the server is not valid
//
// See also, https://api-docs.form3.tech/api.html#organisation-accounts-delete
func (c client) DeleteAccount(input *inputs.DeleteAccountInput) (*outputs.DeleteAccountOutput, error) {
	req := requests.NewDeleteAccountRequest(c.config.BaseUrl)

	return req.Execute(input)
}
