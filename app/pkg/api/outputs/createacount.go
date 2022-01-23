package outputs

import "form3app/pkg/model"

// CreateAccountOutput the response from the server for a CreateAccount request.
type CreateAccountOutput struct {
	Data model.AccountData `json:"data"`
}
