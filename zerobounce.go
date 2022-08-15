package zerobounce

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/netip"
	"os"
	"strings"
	"time"
)

const (
	APIURL                   = "https://api.zerobounce.net/v2"
	httpClientSecondsTimeout = 30
)

var (
	errUnexpectedHTTPStatusCode = errors.New("expected http status code")
)

type API struct {
	apiURL string
	apiKey string
	client *http.Client
}

func NewAPI(apiKey string) *API {
	return &API{
		apiURL: APIURL,
		apiKey: apiKey,
		client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       time.Second * httpClientSecondsTimeout,
		},
	}
}

func (a *API) SetAPIKey(key string) *API {
	a.apiKey = key

	return a
}

func (a *API) GetAPIKey() string {
	return a.apiKey
}

func (a *API) SetAPIURL(url string) *API {
	a.apiURL = strings.TrimRight(url, "/")

	return a
}

func (a *API) GetAPIURL() string {
	return a.apiURL
}

func (a *API) GetCredits() (*ResponseType[CreditsResponseSuccess], error) {
	url := fmt.Sprintf("%s/getcredits?api_key=%s", a.GetAPIURL(), a.GetAPIKey())

	response, err := a.client.Get(url)

	if err != nil {
		if os.IsTimeout(err) {
			return nil, fmt.Errorf("http client timeout reached: %w", err)
		}

		return nil, fmt.Errorf("http client error: %w", err)
	}

	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, errUnexpectedHTTPStatusCode
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %w", err)
	}

	var successResponse CreditsResponseSuccess
	err = json.Unmarshal(b, &successResponse)

	if err == nil {
		return newResponseSuccess(&successResponse), nil
	}

	var errorResponse ResponseError
	err = json.Unmarshal(b, &errorResponse)

	if err == nil {
		return newResponseError[CreditsResponseSuccess](&errorResponse), nil
	}

	return nil, fmt.Errorf("invalid response from the server: %w", err)
}

func (a *API) Validate(email string, ipAddress *netip.Addr) (*ResponseType[ValidateResponseSuccess], error) {
	url := fmt.Sprintf("%s/validate?api_key=%s&email=%s", a.GetAPIURL(), a.GetAPIKey(), email)
	if ipAddress != nil {
		url = fmt.Sprintf("%s&ip_address=%s", url, ipAddress.String())
	}

	response, err := a.client.Get(url)

	if err != nil {
		if os.IsTimeout(err) {
			return nil, fmt.Errorf("http client timeout reached: %w", err)
		}

		return nil, fmt.Errorf("http client error: %w", err)
	}

	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, errUnexpectedHTTPStatusCode
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %w", err)
	}

	var successResponse ValidateResponseSuccess
	err = json.Unmarshal(b, &successResponse)

	if err == nil {
		return newResponseSuccess(&successResponse), nil
	}

	var errorResponse ResponseError
	err = json.Unmarshal(b, &errorResponse)

	if err == nil {
		return newResponseError[ValidateResponseSuccess](&errorResponse), nil
	}

	return nil, fmt.Errorf("invalid response from the server: %w", err)
}
