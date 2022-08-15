package zerobounce_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/twisted1919/zerobounce-go"
	"testing"
)

func TestCreditsResponseSuccess_AsInt(t *testing.T) {
	t.Parallel()

	var credits zerobounce.CreditsResponseSuccess
	err := json.Unmarshal([]byte(`{"Credits": "100"}`), &credits)

	assert.Nil(t, err)
	assert.Equal(t, credits.AsInt(), 100)
	assert.Equal(t, credits.Credits, "100")
}
