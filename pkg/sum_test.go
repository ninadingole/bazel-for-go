package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sum(t *testing.T) {
	assert.Equal(t, 3, Sum(1, 2))
}
