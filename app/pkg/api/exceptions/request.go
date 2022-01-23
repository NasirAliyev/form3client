package exceptions

import "fmt"

// InvalidInputException this exception describes that
// the input performed by user is not valid.
type InvalidInputException struct{}

func (ex InvalidInputException) Error() string {
	return fmt.Sprintf("The input is not valid")
}

// RequestFailedException this exception describes that
// the request failed during the process or failed on server side
type RequestFailedException struct{}

func (ex RequestFailedException) Error() string {
	return fmt.Sprintf("Request failed")
}

// InvalidResponseException this exception describes that
// the response from server is not valid
type InvalidResponseException struct{}

func (ex InvalidResponseException) Error() string {
	return fmt.Sprintf("The response is not valid")
}
