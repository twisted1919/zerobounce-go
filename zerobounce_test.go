package zerobounce_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/twisted1919/zerobounce-go"
	"net/netip"
	"os"
	"testing"
)

func TestAPI_SetAPIKey(t *testing.T) {
	t.Parallel()

	api := zerobounce.NewAPI("test-1")
	assert.Equal(t, api.GetAPIKey(), "test-1")

	api.SetAPIKey("test-2")
	assert.Equal(t, api.GetAPIKey(), "test-2")
}

func TestAPI_SetAPIURL(t *testing.T) {
	t.Parallel()

	api := zerobounce.NewAPI("test-1")
	assert.Equal(t, api.GetAPIURL(), zerobounce.APIURL)

	api.SetAPIURL("https://example.com")
	assert.Equal(t, api.GetAPIURL(), "https://example.com")
}

func TestAPI_GetCredits(t *testing.T) {
	t.Parallel()

	apiKey, _ := os.LookupEnv("ZEROBOUNCE_API_KEY")
	api := zerobounce.NewAPI(apiKey)

	response, err := api.GetCredits()
	assert.Nil(t, err)
	assert.Nil(t, response.Error)
	assert.NotEmpty(t, response.Success)
	assert.NotEmpty(t, response.Success.Credits)
	assert.GreaterOrEqual(t, response.Success.AsInt(), 0)
}

func TestAPI_Validate(t *testing.T) {
	t.Parallel()

	apiKey, _ := os.LookupEnv("ZEROBOUNCE_API_KEY")
	api := zerobounce.NewAPI(apiKey)

	type EmailQuery struct {
		email          string
		expectedStatus string
		statusCb       func(success *zerobounce.ValidateResponseSuccess) bool
	}

	emails := []EmailQuery{
		{
			email:          "valid@example.com",
			expectedStatus: "valid",
			statusCb:       func(success *zerobounce.ValidateResponseSuccess) bool { return success.IsValid() },
		},
		{
			email:          "invalid@example.com",
			expectedStatus: "invalid",
			statusCb:       func(success *zerobounce.ValidateResponseSuccess) bool { return success.IsInvalid() },
		},
		{
			email:          "catch_all@example.com",
			expectedStatus: "catch-all",
			statusCb:       func(success *zerobounce.ValidateResponseSuccess) bool { return success.IsCatchAll() },
		},
		{
			email:          "unknown@example.com",
			expectedStatus: "unknown",
			statusCb:       func(success *zerobounce.ValidateResponseSuccess) bool { return success.IsUnknown() },
		},
		{
			email:          "spamtrap@example.com",
			expectedStatus: "spamtrap",
			statusCb:       func(success *zerobounce.ValidateResponseSuccess) bool { return success.IsSpamtrap() },
		},
		{
			email:          "abuse@example.com",
			expectedStatus: "abuse",
			statusCb:       func(success *zerobounce.ValidateResponseSuccess) bool { return success.IsAbuse() },
		},
		{
			email:          "do_not_mail@example.com",
			expectedStatus: "do_not_mail",
			statusCb:       func(success *zerobounce.ValidateResponseSuccess) bool { return success.IsDoNotMail() },
		},
	}

	for _, email := range emails {
		response, err := api.Validate(email.email, nil)
		
		assert.Nil(t, err)
		assert.Nil(t, response.Error)
		assert.Equal(t, response.Success.Status, email.expectedStatus)
		assert.True(t, email.statusCb(response.Success))
	}
}

func TestAPI_ValidateWithIP(t *testing.T) {
	t.Parallel()

	apiKey, _ := os.LookupEnv("ZEROBOUNCE_API_KEY")
	api := zerobounce.NewAPI(apiKey)

	ipAddress := netip.MustParseAddr("127.0.0.1")
	response, err := api.Validate("valid@example.com", &ipAddress)

	assert.Nil(t, err)
	assert.Nil(t, response.Error)
	assert.Equal(t, response.Success.Status, "valid")
	assert.True(t, response.Success.IsValid())
}
