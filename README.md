# Form3 API client. Task description [here](Form3README.md).

### About

This is a client library in Go to access form3 fake account API, which is provided as a Docker
container in the file `docker-compose.yaml` of this repository. Please refer to the
[Form3 documentation](http://api-docs.form3.tech/api.html#organisation-accounts) for information on how to interact with the API. Please note that the fake account API does not require any authorisation or authentication.

Main README of task. [here](Form3README.md).

### Instruction
*In order to run tests, it is enough just to run this command :*
```
docker compose up
```


### Structure

```
├── go.mod
├── go.sum
├── main.go
└── pkg
    ├── api
    │ ├── exceptions
    │ │ ├── account.go
    │ │ └── request.go
    │ ├── inputs
    │ │ ├── createacount.go
    │ │ ├── deleteacount.go
    │ │ └── fetchacount.go
    │ ├── outputs
    │ │ ├── createacount.go
    │ │ ├── deleteaccount.go
    │ │ └── fetchaccount.go
    │ └── requests
    │     ├── createaccount.go
    │     ├── deleteaccount.go
    │     ├── fetchaccount.go
    │     └── request.go
    ├── client.go
    ├── client_test.go
    └── model
        └── account.go
```

## Implementation
`client.go` the main api for using this package.
Package can be extracted with minimal changes and can be used as independent package.
I did not do that due to run it from one place.
Every API divided into separate files in order to make it clean. E.g: `requests/createacount.go` 


## Example

```shell
config := form3.Config{
    BaseUrl: "https://form3-api.example.com",
}

client := form3.NewClient(config)

input := inputs.FetchAccountInput{
    AccountId: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
}

account, err := client.FetchAccount(&input)
```