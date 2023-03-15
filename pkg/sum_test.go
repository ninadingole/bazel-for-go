package pkg_test

import (
	"github.com/ninadingole/bazel-for-go/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sum(t *testing.T) {
	assert.Equal(t, 3, pkg.Sum(1, 2))
}
