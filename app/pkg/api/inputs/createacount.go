package inputs

import "form3app/pkg/model"

// CreateAccountInput represents the request to create an account.
type CreateAccountInput struct {
	Data model.AccountData `json:"data"`
}
