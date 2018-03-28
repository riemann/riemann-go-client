package riemanngo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewUDPClient(t *testing.T) {
	client := NewUDPClient("127.0.0.1:5555", 5*time.Second)
	assert.Equal(t, "127.0.0.1:5555", client.addr)
}
