package geocodexyz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseRequestToMap(t *testing.T) {
	input := ReverseLookupRequest{}
	mapped := input.toMap()
	assert.Equal(t, "json", mapped["geoit"])
	assert.Equal(t, "en", mapped["lang"])
	assert.Equal(t, "0,0", mapped["locate"])
}

func TestReverseLookup(t *testing.T) {
	input := &ReverseLookupRequest{
		Lat:    42.8428226,
		Lng:    -71.2094162,
		Region: "us",
	}
	result, apiError, err := ReverseLookup(input)
	if apiError != nil {
		assert.NotNil(t, err)
		// it could be either 006 (rate limit) or 003 (auth)
		assert.True(t, apiError.Code == "006" || apiError.Code == "003")
		if apiError.Code == "006" {
			assert.NotZero(t, apiError.Requests)
		}
		assert.NotEmpty(t, apiError.Description)
	} else {
		// check just a couple fields
		assert.Equal(t, "NH", result.State)
		assert.Equal(t, "03079", result.Postal)
		assert.Equal(t, "103 Haverhill RD", result.Street)
		assert.Equal(t, "United States of America", result.Country)
		assert.NotZero(t, result.Confidence)
		assert.NotNil(t, result.RawRequest)
	}
}

func TestReverseLookupNotPrecise(t *testing.T) {
	input := &ReverseLookupRequest{
		Lat:    40.8428229,
		Lng:    -69.2094162,
		Region: "us",
	}
	result, apiError, err := ReverseLookup(input)
	if apiError != nil {
		assert.NotNil(t, err)
		// it could be either 006 (rate limit) or 003 (auth) or 008 (not suitable)
		assert.True(t, apiError.Code == "006" || apiError.Code == "003" || apiError.Code == "008")
		if apiError.Code == "006" {
			assert.NotZero(t, apiError.Requests)
		}
		assert.NotEmpty(t, apiError.Description)
	} else {
		// check just a couple fields
		assert.Nil(t, err)
		assert.Equal(t, "NH", result.State)
		assert.Equal(t, "03079", result.Postal)
		assert.Equal(t, "103 Haverhill RD", result.Street)
		assert.Equal(t, "United States of America", result.Country)
		assert.NotZero(t, result.Confidence)
		assert.NotNil(t, result.RawRequest)
	}
}
