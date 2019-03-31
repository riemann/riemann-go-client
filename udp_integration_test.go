// +build integration

package riemanngo

import (
	"testing"
	"time"
)

func TestSendEventUdp(t *testing.T) {
	c := NewUDPClient("127.0.0.1:5555", 5*time.Second)
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

func TestSendEventsUDP(t *testing.T) {
	c := NewUDPClient("127.0.0.1:5555", 5*time.Second)
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

func TestUDPConnec(t *testing.T) {
	c := NewUDPClient("does.not.exists:8888", 5*time.Second)
	// should produce an error
	err := c.Connect()
	if err == nil {
		t.Error("Error, should fail")
	}
}
