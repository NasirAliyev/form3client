package inputs

// DeleteAccountInput represents the request to delete an account.
type DeleteAccountInput struct {
	AccountId string `json:"account_id"`
	Version   *int64 `json:"version"`
}
