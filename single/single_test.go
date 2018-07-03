package single

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.Exactly(t,New(),New(),"no equal")
}
