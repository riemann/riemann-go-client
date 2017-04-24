// +build integration

package riemanngo

import (
	"testing"
)

func TestSendEventTcp(t *testing.T) {
	c := NewTcpClient("127.0.0.1:5555")
	err := c.Connect(5)
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
	c := NewTcpClient("127.0.0.1:5555")
	err := c.Connect(5)
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
	c := NewTcpClient("127.0.0.1:5555")
	err := c.Connect(5)
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
	c := NewTcpClient("does.not.exists:8888")
	// should produce an error
	err := c.Connect(2)
	if err == nil {
		t.Error("Error, should fail")
	}
}
