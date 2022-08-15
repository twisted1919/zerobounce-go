package zerobounce

import "strconv"

type CreditsResponseSuccess struct {
	Credits string `json:"Credits"` //nolint:tagliatelle
}

func (c *CreditsResponseSuccess) AsInt() int {
	credits, err := strconv.Atoi(c.Credits)
	if err != nil {
		return 0
	}

	return credits
}
