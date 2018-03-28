// +build integration

package riemanngo

import (
	"testing"
	"time"
)

func TestSendEventTcp(t *testing.T) {
	c := NewTCPClient("127.0.0.1:5555", 5*time.Second)
	err := c.Connect()
	defer c.Close()
	if err != nil {
		t.Error("Error Tcp client Connect")
	}
	result, err := SendEvent(c, &Event{
		Service: "LOOOl",
		Metric:  100,
		Tags:    []string{"nonblocking"},
	})
	if !*result.Ok {
		t.Error("Error Tcp client SendEvent")
	}
}

func TestSendEventsTcp(t *testing.T) {
	c := NewTCPClient("127.0.0.1:5555", 5*time.Second)
	err := c.Connect()
	defer c.Close()
	if err != nil {
		t.Error("Error Tcp client Connect")
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
	if !*result.Ok {
		t.Error("Error Tcp client SendEvent")
	}
}

func TestQueryIndex(t *testing.T) {
	c := NewTCPClient("127.0.0.1:5555", 5*time.Second)
	err := c.Connect()
	defer c.Close()
	if err != nil {
		t.Error("Error Tcp client Connect")
	}
	events := []Event{
		{
			Host:    "foobaz",
			Service: "golang",
			Metric:  100,
			Tags:    []string{"hello"},
		},
		{
			Host:    "foobar",
			Service: "golang",
			Metric:  200,
			Tags:    []string{"goodbye"},
		},
	}
	result, err := SendEvents(c, &events)
	if !*result.Ok {
		t.Error("Error Tcp client SendEvent")
	}
	queryResult, err := c.QueryIndex("(service = \"golang\")")
	if len(queryResult) != 2 {
		t.Error("Error Tcp client QueryIndex")
	}
}

func TestTcpConnec(t *testing.T) {
	c := NewTCPClient("does.not.exists:8888", 5*time.Second)
	// should produce an error
	err := c.Connect()
	if err == nil {
		t.Error("Error, should fail")
	}
}
