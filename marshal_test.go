package riemanngo

import (
	"math"
	"testing"
	"time"

	pb "github.com/golang/protobuf/proto"
	"github.com/riemann/riemann-go-client/proto"
)

func TestEventToProtocolBuffer(t *testing.T) {
	testCases := []struct {
		desc     string
		event    Event
		expected proto.Event
	}{
		{
			desc: "simple event, metric int32",
			event: Event{
				Host:    "baz",
				Service: "foobar",
				Metric:  int32(100),
				Tags:    []string{"hello"},
				Time:    time.Unix(100, 0),
			},
			expected: proto.Event{
				Host:         pb.String("baz"),
				Time:         pb.Int64(100),
				TimeMicros:   pb.Int64(100000000),
				MetricSint64: pb.Int64(100),
				Service:      pb.String("foobar"),
				Tags:         []string{"hello"},
			},
		},
		{
			desc: "simple event, metric int",
			event: Event{
				Host:    "baz",
				Service: "foobar",
				Metric:  100,
				Tags:    []string{"hello"},
				Time:    time.Unix(100, 0),
			},
			expected: proto.Event{
				Host:         pb.String("baz"),
				Time:         pb.Int64(100),
				TimeMicros:   pb.Int64(100000000),
				MetricSint64: pb.Int64(100),
				Service:      pb.String("foobar"),
				Tags:         []string{"hello"},
			},
		},
		{
			desc: "event with attributes, metric float",
			event: Event{
				Host:    "baz",
				Service: "foobar",
				Metric:  100.1,
				Tags:    []string{"hello"},
				Time:    time.Unix(100, 0),
				Ttl:     10,
				Attributes: map[string]string{
					"foo": "bar",
					"bar": "baz",
				},
			},
			expected: proto.Event{
				Host:       pb.String("baz"),
				Time:       pb.Int64(100),
				TimeMicros: pb.Int64(100000000),
				MetricD:    pb.Float64(100.1),
				Service:    pb.String("foobar"),
				Tags:       []string{"hello"},
				Ttl:        pb.Float32(10),
				Attributes: []*proto.Attribute{
					{
						Key:   pb.String("bar"),
						Value: pb.String("baz"),
					},
					{
						Key:   pb.String("foo"),
						Value: pb.String("bar"),
					},
				},
			},
		},
		{
			desc: "full event",
			event: Event{
				Host:        "baz",
				Service:     "foobar",
				Ttl:         20,
				Description: "desc",
				State:       "critical",
				Metric:      100,
				Tags:        []string{"hello"},
				Time:        time.Unix(100, 0),
			},
			expected: proto.Event{
				Host:         pb.String("baz"),
				Time:         pb.Int64(100),
				TimeMicros:   pb.Int64(100000000),
				Ttl:          pb.Float32(20),
				Description:  pb.String("desc"),
				State:        pb.String("critical"),
				MetricSint64: pb.Int64(100),
				Service:      pb.String("foobar"),
				Tags:         []string{"hello"},
			},
		},
		{
			desc: "test int64",
			event: Event{
				Host:        "baz",
				Service:     "foobar",
				Ttl:         20,
				Description: "desc",
				State:       "critical",
				Metric:      int64(100),
				Tags:        []string{"hello"},
				Time:        time.Unix(100, 0),
			},
			expected: proto.Event{
				Host:         pb.String("baz"),
				Time:         pb.Int64(100),
				TimeMicros:   pb.Int64(100000000),
				Ttl:          pb.Float32(20),
				Description:  pb.String("desc"),
				State:        pb.String("critical"),
				MetricSint64: pb.Int64(100),
				Service:      pb.String("foobar"),
				Tags:         []string{"hello"},
			},
		},
		{
			desc: "test float32",
			event: Event{
				Host:        "baz",
				Service:     "foobar",
				Ttl:         20,
				Description: "desc",
				State:       "critical",
				Metric:      float32(100.0),
				Tags:        []string{"hello"},
				Time:        time.Unix(100, 0),
			},
			expected: proto.Event{
				Host:        pb.String("baz"),
				Time:        pb.Int64(100),
				TimeMicros:  pb.Int64(100000000),
				Ttl:         pb.Float32(20),
				Description: pb.String("desc"),
				State:       pb.String("critical"),
				MetricD:     pb.Float64(100.0),
				Service:     pb.String("foobar"),
				Tags:        []string{"hello"},
			},
		},
		{
			desc: "test float64",
			event: Event{
				Host:        "baz",
				Service:     "foobar",
				Ttl:         20,
				Description: "desc",
				State:       "critical",
				Metric:      float64(100.12),
				Tags:        []string{"hello"},
				Time:        time.Unix(100, 0),
			},
			expected: proto.Event{
				Host:        pb.String("baz"),
				Time:        pb.Int64(100),
				TimeMicros:  pb.Int64(100000000),
				Ttl:         pb.Float32(20),
				Description: pb.String("desc"),
				State:       pb.String("critical"),
				MetricD:     pb.Float64(100.12),
				Service:     pb.String("foobar"),
				Tags:        []string{"hello"},
			},
		},
		{
			desc: "simple event with time in nanosecond",
			event: Event{
				Host:    "baz",
				Service: "foobar",
				Metric:  uint32(100),
				Tags:    []string{"hello"},
				Time:    time.Unix(100, 123456789),
			},
			expected: proto.Event{
				Host:         pb.String("baz"),
				Time:         pb.Int64(100),
				TimeMicros:   pb.Int64(100123456),
				MetricSint64: pb.Int64(100),
				Service:      pb.String("foobar"),
				Tags:         []string{"hello"},
			},
		},
		{
			desc: "Event without metrics",
			event: Event{
				Host:    "baz",
				Service: "foobar",
				Time:    time.Unix(100, 123456789),
			},
			expected: proto.Event{
				Host:       pb.String("baz"),
				Service:    pb.String("foobar"),
				Time:       pb.Int64(100),
				TimeMicros: pb.Int64(100123456),
			},
		},
		{
			desc: "Event with uint type",
			event: Event{
				Host:    "baz",
				Metric:  uint64(5),
				Service: "foobar",
				Time:    time.Unix(100, 123456789),
			},
			expected: proto.Event{
				Host:         pb.String("baz"),
				Service:      pb.String("foobar"),
				Time:         pb.Int64(100),
				TimeMicros:   pb.Int64(100123456),
				MetricSint64: pb.Int64(5),
			},
		},
		{
			desc: "Event with uint type, overflow",
			event: Event{
				Host:    "baz",
				Metric:  uint64(math.MaxUint64),
				Service: "foobar",
				Time:    time.Unix(100, 123456789),
			},
			expected: proto.Event{
				Host:         pb.String("baz"),
				Service:      pb.String("foobar"),
				Time:         pb.Int64(100),
				TimeMicros:   pb.Int64(100123456),
				MetricSint64: pb.Int64(-1),
			},
		},
	}

	for _, tc := range testCases {
		obtained, err := EventToProtocolBuffer(&tc.event)

		if err != nil {
			t.Errorf(
				"Marshal error (%s)", tc.desc,
			)
		}

		if !pb.Equal(obtained, &tc.expected) {
			t.Errorf(
				"Error during event to protobuf conversion (%s)",
				tc.desc,
			)
		}
	}
}

func compareEvents(e1 *Event, e2 *Event, t *testing.T) {
	if e1.Tags[0] != e2.Tags[0] {
		t.Error("Error during event to events conversion to protobuf (Tags)")
	}
	if e1.Host != e2.Host {
		t.Error("Error during event to events conversion to protobuf (Host)")
	}
	if e1.Time != e2.Time {
		t.Error("Error during event to events conversion to protobuf (Time)")
	}
	if e1.Ttl != e2.Ttl {
		t.Error("Error during event to events conversion to protobuf (Ttl)")
	}
	if e1.Description != e2.Description {
		t.Error("Error during event to events conversion to protobuf (Description)")
	}
	if e1.Metric != e2.Metric {
		t.Error("Error during event to events conversion to protobuf (Metric)")
	}
	if e1.State != e2.State {
		t.Error("Error during event to events conversion to protobuf (State)")
	}
	if e1.Service != e2.Service {
		t.Error("Error during event to events conversion to protobuf (Service)")
	}
	if len(e1.Tags) != len(e2.Tags) {
		t.Error("Error during event to events conversion to protobuf (Tags)")
	}
	for i, v := range e1.Tags {
		if v != e2.Tags[i] {
			t.Error("Error during event to events conversion to protobuf (Tags)")
		}
	}
	if len(e1.Attributes) != len(e2.Attributes) {
		t.Error("Error during event to events conversion to protobuf (Attributes)")
	}
	for i, v := range e1.Attributes {
		if v != e2.Attributes[i] {
			t.Error("Error during event to events conversion to protobuf (Attributes)")
		}
	}
}

func TestProtocolBuffersToEvents(t *testing.T) {
	pbEvents := []*proto.Event{
		{
			Host:         pb.String("baz"),
			Time:         pb.Int64(100),
			Ttl:          pb.Float32(20),
			Description:  pb.String("desc"),
			State:        pb.String("critical"),
			MetricSint64: pb.Int64(100),
			Service:      pb.String("foobar"),
			Tags:         []string{"hello"},
		},
	}
	event := Event{
		Host:        "baz",
		Time:        time.Unix(100, 0),
		Ttl:         20,
		Description: "desc",
		State:       "critical",
		Metric:      int64(100),
		Service:     "foobar",
		Tags:        []string{"hello"},
	}
	events := ProtocolBuffersToEvents(pbEvents)
	compareEvents(&events[0], &event, t)
	pbEvents = []*proto.Event{
		{
			Host:         pb.String("baz"),
			Time:         pb.Int64(100),
			Ttl:          pb.Float32(20),
			Description:  pb.String("desc"),
			State:        pb.String("critical"),
			MetricSint64: pb.Int64(100),
			Service:      pb.String("foobar"),
			Tags:         []string{"hello"},
		},
		{
			Host:        pb.String("baz"),
			Time:        pb.Int64(100),
			Ttl:         pb.Float32(20),
			Description: pb.String("desc"),
			State:       pb.String("critical"),
			MetricD:     pb.Float64(100.1),
			Service:     pb.String("foobar"),
			Tags:        []string{"hello"},
			Attributes: []*proto.Attribute{
				{
					Key:   pb.String("foo"),
					Value: pb.String("bar"),
				},
				{
					Key:   pb.String("bar"),
					Value: pb.String("baz"),
				},
			},
		},
	}
	event1 := Event{
		Host:        "baz",
		Time:        time.Unix(100, 0),
		Ttl:         20,
		Description: "desc",
		State:       "critical",
		Metric:      int64(100),
		Service:     "foobar",
		Tags:        []string{"hello"},
	}
	event2 := Event{
		Host:        "baz",
		Service:     "foobar",
		Description: "desc",
		State:       "critical",
		Metric:      100.1,
		Tags:        []string{"hello"},
		Time:        time.Unix(100, 0),
		Ttl:         20,
		Attributes: map[string]string{
			"foo": "bar",
			"bar": "baz",
		},
	}
	events = ProtocolBuffersToEvents(pbEvents)
	compareEvents(&events[0], &event1, t)
	compareEvents(&events[1], &event2, t)

	pbEvents = []*proto.Event{
		{
			Host:         pb.String("baz"),
			Time:         pb.Int64(100),
			TimeMicros:   pb.Int64(100123456),
			Ttl:          pb.Float32(20),
			Description:  pb.String("desc"),
			State:        pb.String("critical"),
			MetricSint64: pb.Int64(100),
			Service:      pb.String("foobar"),
			Tags:         []string{"hello"},
		},
	}
	event = Event{
		Host:        "baz",
		Time:        time.Unix(100, 123456000),
		Ttl:         20,
		Description: "desc",
		State:       "critical",
		Metric:      int64(100),
		Service:     "foobar",
		Tags:        []string{"hello"},
	}
	events = ProtocolBuffersToEvents(pbEvents)
	compareEvents(&events[0], &event, t)

	pbEvents = []*proto.Event{
		{
			Host:         pb.String("baz"),
			TimeMicros:   pb.Int64(100123456),
			Ttl:          pb.Float32(20),
			Description:  pb.String("desc"),
			State:        pb.String("critical"),
			MetricSint64: pb.Int64(100),
			Service:      pb.String("foobar"),
			Tags:         []string{"hello"},
		},
	}
	event = Event{
		Host:        "baz",
		Time:        time.Unix(100, 123456000),
		Ttl:         20,
		Description: "desc",
		State:       "critical",
		Metric:      int64(100),
		Service:     "foobar",
		Tags:        []string{"hello"},
	}
	events = ProtocolBuffersToEvents(pbEvents)
	compareEvents(&events[0], &event, t)
}

func BenchmarkEventToProtocolBuffer(b *testing.B) {
	e := &Event{
		Host:    "baz",
		Service: "foobar",
		Metric:  123,
		Tags:    []string{"hello"},
		Time:    time.Unix(100, 0),
		Attributes: map[string]string{
			"d": "4",
			"c": "3",
			"b": "2",
			"a": "1",
		},
	}

	for n := 0; n < b.N; n++ {
		_, err := EventToProtocolBuffer(e)

		if err != nil {
			b.Fatal("Marshaling error")
		}
	}
}
