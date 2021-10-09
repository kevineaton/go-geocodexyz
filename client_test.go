package geocodexyz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	params := ReverseLookupRequest{
		Locate: "42.8428226,-71.2094162",
		Region: "us",
	}

	// if there is no auth key, this could either work or be throttled
	// so we need to check for either; setting up an auth key will
	// greatly increase the likelihood of this succeeding
	responseBytes, apiError, err := makeCall(params)
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
