package riemanngo

import (
	"testing"
)

func TestNewTlsClientWithInsecure(t *testing.T) {
	client, err := NewTlsClient("127.0.0.1:5555", "tls/client.crt", "tls/client.key", true)
	if err != nil {
		t.Error("Error creating a new tls client")
	}
	if client.addr != "127.0.0.1:5555" {
		t.Error("Error creating a new tls client")
	}
}

func TestNewTlsClientWithoutInsecure(t *testing.T) {
	client, err := NewTlsClient("127.0.0.1:5555", "tls/client.crt", "tls/client.key", false)
	if err != nil {
		t.Error("Error creating a new tls client")
	}
	if client.addr != "127.0.0.1:5555" {
		t.Error("Error creating a new tls client")
	}
}
