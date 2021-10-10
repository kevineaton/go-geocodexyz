package geocodexyz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	params := ReverseLookupRequest{
		Lat:    42.8428226,
		Lng:    -71.2094162,
		Region: "us",
	}

	// if there is no auth key, this could either work or be throttled
	// so we need to check for either; setting up an auth key will
	// greatly increase the likelihood of this succeeding
	responseBytes, apiError, err := makeCall(params.toMap())
	if apiError != nil {
		assert.NotNil(t, err)
		// it could be either 006 (rate limit) or 003 (auth)
		assert.True(t, apiError.Code == "006" || apiError.Code == "003")
		if apiError.Code == "006" {
			assert.NotZero(t, apiError.Requests)
		}
		assert.NotEmpty(t, apiError.Description)
	} else {
		assert.NotNil(t, responseBytes)
		assert.NotZero(t, len(responseBytes))
	}
}

func TestClientRateLimitOver(t *testing.T) {
	// we want to specifically try to trigger the rate limiting
	// so we will set the auth to an empty string and make a bunch of calls in a loop
	// at least 1 should be a trigger
	key := config.APIKey
	config.APIKey = ""

	params := &ReverseLookupRequest{
		Lat:    42.8428226,
		Lng:    -71.2094162,
		Region: "us",
	}
	foundRateLimitError := false
	for i := 0; i < 10; i++ {
		params.Lat += float64(i)
		params.Lng += float64(i)
		_, apiErr, _ := makeCall(params.toMap())
		if apiErr != nil {
			if apiErr.Code == "006" {
				foundRateLimitError = true
			}
		}
	}
	assert.True(t, foundRateLimitError)

	config.APIKey = key
}
