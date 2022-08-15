package zerobounce_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/twisted1919/zerobounce-go"
	"testing"
)

func TestValidateResponseSuccess_IsValid(t *testing.T) {
	t.Parallel()

	assert.True(t, (&zerobounce.ValidateResponseSuccess{Status: "valid"}).IsValid())
}

func TestValidateResponseSuccess_IsInvalid(t *testing.T) {
	t.Parallel()

	assert.True(t, (&zerobounce.ValidateResponseSuccess{Status: "invalid"}).IsInvalid())
}

func TestValidateResponseSuccess_IsCatchAll(t *testing.T) {
	t.Parallel()

	assert.True(t, (&zerobounce.ValidateResponseSuccess{Status: "catch-all"}).IsCatchAll())
}

func TestValidateResponseSuccess_IsUnknown(t *testing.T) {
	t.Parallel()

	assert.True(t, (&zerobounce.ValidateResponseSuccess{Status: "unknown"}).IsUnknown())
}

func TestValidateResponseSuccess_IsSpamtrap(t *testing.T) {
	t.Parallel()

	assert.True(t, (&zerobounce.ValidateResponseSuccess{Status: "spamtrap"}).IsSpamtrap())
}

func TestValidateResponseSuccess_IsAbuse(t *testing.T) {
	t.Parallel()

	assert.True(t, (&zerobounce.ValidateResponseSuccess{Status: "abuse"}).IsAbuse())
}

func TestValidateResponseSuccess_IsDoNotMail(t *testing.T) {
	t.Parallel()

	assert.True(t, (&zerobounce.ValidateResponseSuccess{Status: "do_not_mail"}).IsDoNotMail())
}
