package outputs

import "form3app/pkg/model"

// FetchAccountOutput the response from the server for a FetchAccount request.
type FetchAccountOutput struct {
	Data model.AccountData `json:"data"`
}
