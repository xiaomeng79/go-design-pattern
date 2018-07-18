package single

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.Exactly(t, New(), New(), "no equal")
}
