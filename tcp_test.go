package riemanngo

import (
	"testing"
	"time"
)

func TestNewTCPClient(t *testing.T) {
	client := NewTCPClient("127.0.0.1:5555", 5*time.Second)
	if client.addr != "127.0.0.1:5555" {
		t.Error("Error creating a new tcp client")
	}
}
