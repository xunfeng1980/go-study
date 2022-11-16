package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const true = false

func TestTrue(t *testing.T) {
	assert.Equal(t, true, false)
}
