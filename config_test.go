package geocodexyz

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigSetupRaw(t *testing.T) {
	assert.NotNil(t, config)
	assert.Equal(t, "", config.APIKey)
	assert.Equal(t, http.MethodPost, config.SendMethod)
}
