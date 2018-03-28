package riemanngo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewTCPClient(t *testing.T) {
	client := NewTCPClient("127.0.0.1:5555", 5*time.Second)
	assert.Equal(t, "127.0.0.1:5555", client.addr)
}

func TestNewTLSClientWithInsecure(t *testing.T) {
	config, err := GetTLSConfig("127.0.0.1", "tls/client.crt", "tls/client.key", true)
	assert.NoError(t, err)
	client, err := NewTLSClient("127.0.0.1:5555", config, 5*time.Second)
	assert.NoError(t, err)

	assert.Equal(t, "127.0.0.1:5555", client.addr)
}

func TestNewTLSClientWithoutInsecure(t *testing.T) {
	config, err := GetTLSConfig("127.0.0.1", "tls/client.crt", "tls/client.key", false)
	assert.NoError(t, err)
	client, err := NewTLSClient("127.0.0.1:5555", config, 5*time.Second)
	assert.NoError(t, err)

	assert.Equal(t, "127.0.0.1:5555", client.addr)
}
