package hello

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	rsp := Hello()
	assert2 := assert.New(t)
	assert2.NotEmpty(rsp)
	fmt.Println(rsp)
	assert2.Greater(2, 1)
}
