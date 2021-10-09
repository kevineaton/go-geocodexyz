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
	assert.Equal(t, "", mapped["locate"])
}
