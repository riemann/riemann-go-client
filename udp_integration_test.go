// +build integration

package riemanngo

import (
	"testing"
	"time"
)

func TestSendEventUdp(t *testing.T) {
	c := NewUdpClient("127.0.0.1:5555", 5*time.Second)
	err := c.Connect()
	defer c.Close()
	if err != nil {
		t.Error("Error Udp client Connect")
	}
	result, err := SendEvent(c, &Event{
		Service: "LOOOl",
		Metric:  100,
		Tags:    []string{"nonblocking"},
	})
	if result != nil || err != nil {
		t.Error("Error Udp client SendEvent")
	}
}

func TestSendEventsUdp(t *testing.T) {
	c := NewUdpClient("127.0.0.1:5555", 5*time.Second)
	err := c.Connect()
	defer c.Close()
	if err != nil {
		t.Error("Error Udp client Connect")
	}
	events := []Event{
		{
			Service: "hello",
			Metric:  100,
			Tags:    []string{"hello"},
		},
		{
			Service: "goodbye",
			Metric:  200,
			Tags:    []string{"goodbye"},
		},
	}
	result, err := SendEvents(c, &events)
	if result != nil || err != nil {
		t.Error("Error Udp client SendEvent")
	}
}

func TestUdpConnec(t *testing.T) {
	c := NewUdpClient("does.not.exists:8888", 5*time.Second)
	// should produce an error
	err := c.Connect()
	if err == nil {
		t.Error("Error, should fail")
	}
}
