package riemanngo

import (
	"testing"
	"time"
)

func TestNewTcpClient(t *testing.T) {
	client := NewTcpClient("127.0.0.1:5555", 5*time.Second)
	if client.addr != "127.0.0.1:5555" {
		t.Error("Error creating a new tcp client")
	}
}
