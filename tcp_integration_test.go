// +build integration

package riemanngo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSendEventTcp(t *testing.T) {
	c := NewTCPClient("127.0.0.1:5555", 5*time.Second)
	err := c.Connect()
	defer c.Close()
	assert.NoError(t, err)
	result, err := SendEvent(c, &Event{
		Service: "LOOOl",
		Metric:  100,
		Tags:    []string{"nonblocking"},
	})
	assert.True(t, *result.Ok)
}

func TestSendEventsTcp(t *testing.T) {
	c := NewTCPClient("127.0.0.1:5555", 5*time.Second)
	err := c.Connect()
	defer c.Close()
	assert.NoError(t, err)
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
	assert.True(t, *result.Ok)
}

func TestQueryIndex(t *testing.T) {
	c := NewTCPClient("127.0.0.1:5555", 5*time.Second)
	err := c.Connect()
	defer c.Close()
	assert.NoError(t, err)
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
	assert.NoError(t, err)
	assert.True(t, *result.Ok)
	queryResult, err := c.QueryIndex("(service = \"golang\")")
	assert.True(t, *result.Ok)
	assert.Equal(t, 2, len(queryResult))
}

func TestTcpConnect(t *testing.T) {
	c := NewTCPClient("does.not.exists:8888", 5*time.Second)
	// should produce an error
	err := c.Connect()
	assert.Error(t, err)
}
