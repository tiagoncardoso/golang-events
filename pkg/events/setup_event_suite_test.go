package events

import (
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
	"time"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

type TestEventHandler struct {
	ID int
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	EventDispatcher *EventDispatcher
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

func (s *EventDispatcherTestSuite) SetupTest() {
	s.EventDispatcher = NewEventDispatcher()

	s.handler = TestEventHandler{
		ID: 1,
	}
	s.handler2 = TestEventHandler{
		ID: 2,
	}
	s.handler3 = TestEventHandler{
		ID: 3,
	}

	s.event = TestEvent{Name: "test event", Payload: "test"}
	s.event = TestEvent{Name: "test event 2", Payload: "test2"}
}
