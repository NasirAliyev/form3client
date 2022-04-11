package form3

import (
	"form3app/pkg/api/exceptions"
	"form3app/pkg/api/inputs"
	"form3app/pkg/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// Testing CreateAccount api for positive flow
func TestClient_CreateAccount(t *testing.T) {
	id := uuid.New().String()
	account := getTestAccount(id)

	input := inputs.CreateAccountInput{Data: account}

	config := Config{
		BaseUrl: os.Getenv("API_URL"),
	}

	resp, _ := NewClient(config).CreateAccount(&input)
	actual := resp.Data.ID

	assert.Equal(t, id, actual, "ID is not the same")
}

// Testing CreateAccount api for negative flow.
// When client tries to create the same account with the same ID.
// It should throw an AccountAlreadyExistsException.
func TestClient_CreateAccountWithTheSameId(t *testing.T) {
	id := uuid.New().String()
	account := getTestAccount(id)

	input := inputs.CreateAccountInput{Data: account}

	config := Config{
		BaseUrl: os.Getenv("API_URL"),
	}

	client := NewClient(config)
	_, _ = client.CreateAccount(&input)
	_, err := client.CreateAccount(&input)

	expected := exceptions.AccountAlreadyExistsException{}

	assert.IsType(t, expected, err, "Proper exception is not thrown")
}

// Testing FetchAccount api for positive flow.
// When client tries to get existing account.
// In response ID should be the same as in request.
func TestClient_FetchAccount(t *testing.T) {
	id := uuid.New().String()
	account := getTestAccount(id)

	createInput := inputs.CreateAccountInput{Data: account}

	config := Config{
		BaseUrl: os.Getenv("API_URL"),
	}

	client := NewClient(config)
	_, _ = client.CreateAccount(&createInput)

	input := inputs.FetchAccountInput{AccountId: id}
	resp, _ := client.FetchAccount(&input)

	assert.Equal(t, id, resp.Data.ID, "Cannot get existing account")
}

// Testing FetchAccount api for negative flow.
// When client tries to get the account which should not exist.
// It should throw an NotFoundException.
func TestClient_FetchAccountShouldThrowNotFound(t *testing.T) {
	id := uuid.New().String()

	config := Config{
		BaseUrl: os.Getenv("API_URL"),
	}

	client := NewClient(config)

	getInput := inputs.FetchAccountInput{AccountId: id}
	_, err := client.FetchAccount(&getInput)
	expected := exceptions.NotFoundException{}

	assert.IsType(t, expected, err, "NotFoundException is not thrown")
}

// Testing DeleteAccount api for positive flow.
// When client tries to delete existing account.
// Response should be empty content without error
func TestClient_DeleteAccount(t *testing.T) {
	id := uuid.New().String()
	account := getTestAccount(id)

	createInput := inputs.CreateAccountInput{Data: account}

	config := Config{
		BaseUrl: os.Getenv("API_URL"),
	}

	client := NewClient(config)
	createdAccount, _ := client.CreateAccount(&createInput)

	input := inputs.DeleteAccountInput{
		AccountId: createdAccount.Data.ID,
		Version:   createdAccount.Data.Version,
	}

	resp, _ := client.DeleteAccount(&input)

	fetchInput := inputs.FetchAccountInput{AccountId: id}
	fetchResp, _ := client.FetchAccount(&fetchInput)

	assert.Empty(t, resp, "Request failed")
	assert.Empty(t, fetchResp, "The account still exists")
}

// Testing DeleteAccount api for negative flow.
// When client tries to delete the account which not exist.
// It should throw an NotFoundException.
func TestClient_DeleteAccountShouldThrowNotFound(t *testing.T) {
	config := Config{
		BaseUrl: os.Getenv("API_URL"),
	}

	client := NewClient(config)

	id := uuid.New().String()
	version := int64(1)

	input := inputs.DeleteAccountInput{
		AccountId: id,
		Version:   &version,
	}

	_, err := client.DeleteAccount(&input)
	expected := exceptions.NotFoundException{}

	assert.IsType(t, expected, err, "NotFoundException is not thrown")
}

func getTestAccount(id string) model.AccountData {
	cls := "Personal"
	AccountMatchingOptOut := false
	country := "GB"
	JointAccount := false
	Version := int64(1)

	attr := model.AccountAttributes{
		AccountClassification: &cls,
		AccountMatchingOptOut: &AccountMatchingOptOut,
		AlternativeNames: []string{
			"Abu Amir",
		},
		BankID:       "400300",
		BankIDCode:   "GBDSC",
		BaseCurrency: "GBP",
		Bic:          "NWBKGB22",
		Country:      &country,
		JointAccount: &JointAccount,
		Name: []string{
			"Nasir Aliyev",
		},
		SecondaryIdentification: "A1B2C3D4",
	}

	account := model.AccountData{
		Attributes:     &attr,
		ID:             id,
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
		Version:        &Version,
	}

	return account
}
