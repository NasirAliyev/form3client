# Form3 API client. Task description [here](Form3README.md).

### About

- **Author**: `Nasir Aliyev`
- **Experience in web development**: `10 years (PHP, Python, JS)`
- **Experience in GO**: `0 years`

### Instruction
*In order to run tests, it is enough just to run this command :*
```
docker compose up
```
As it is required, tests are integration tests. I covered positive flow and some negative flows. In order to 
cover all negative flows mocking responses is necessary. I could not understand from requirement,
should I spend additional time for that. If I should please let me know.

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