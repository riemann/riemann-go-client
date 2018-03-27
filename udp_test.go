package riemanngo

import (
	"testing"
	"time"
)

func TestNewUdpClient(t *testing.T) {
	client := NewUdpClient("127.0.0.1:5555", 5*time.Second)
	if client.addr != "127.0.0.1:5555" {
		t.Error("Error creating a new tcp client")
	}
}
