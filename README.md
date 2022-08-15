# ZeroBounce Email Verification Library for Go

## Getting Started
You will need a [zerobounce account](https://www.zerobounce.net) to get started.  
Once you get an account, you will need to [get an api key](https://www.zerobounce.net/members/apikey/)
to use it in the API calls.

## Installation

Require the package

```bash
$ go get https://github.com/twisted1919/zerobounce-go
```

#### Go Version

Requires Go >= 1.8

## Usage

```go
package main

import (
	"fmt"
	"github.com/twisted1919/zerobounce-go"
	"log"
	"net/netip"
)

// Example function to show available credits
func credits(api *zerobounce.API) {
	response, err := api.GetCredits()
	if err != nil {
		log.Fatal(fmt.Errorf("unable to fetch the credits: %w", err))
	}

	if response.IsSuccess() {
		fmt.Println("you have", response.Success.Credits, "credits left")
	} else if response.IsError() {
		fmt.Println("the api returned following error", response.Error.Error)
	}
}

// Example function to validate an email address
func validate(api *zerobounce.API, email string, ipAddress *netip.Addr) {
	response, err := api.Validate(email, ipAddress)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to validate email address: %w", err))
	}

	if response.IsSuccess() {
		fmt.Println("the email address", email, "is", response.Success.Status)
	} else if response.IsError() {
		fmt.Println("the api returned following error", response.Error.Error)
	}
}

func main() {

	// instantiate the api
	api := zerobounce.NewAPI("your-api-key")

	// output the result of validation call for a valid email address
	validate(api, "valid@example.com", nil)

	// output the result of validation call for an invalid email address
	validate(api, "invalid@example.com", nil)

	// output the result of validation call for an valid email address but for a different IP Address
	ipAddress := netip.MustParseAddr("127.0.0.1")
	validate(api, "valid@example.com", &ipAddress)

	// output the result of the get credits call
	credits(api)

}

```

## License
MIT

## Test
Set your api key in the `ZEROBOUNCE_API_KEY` environment variable, then run:  
```bash
$ go test ./...
``` 

## Bug Reports
Report [here](https://github.com/twisted1919/zerobounce-go/issues).