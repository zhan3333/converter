package converter

import "testing"
import "github.com/stretchr/testify/assert"

func TestMac(t *testing.T) {
	assert.NoError(t, Mac())
}
