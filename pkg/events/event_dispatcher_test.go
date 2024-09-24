package events

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface) {
	m.Called(event)
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := s.EventDispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.EventDispatcher.handlers[s.event.GetName()]))

	err = s.EventDispatcher.Register(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.EventDispatcher.handlers[s.event.GetName()]))

	assert.Equal(s.T(), &s.handler, s.EventDispatcher.handlers[s.event.GetName()][0])
	assert.Equal(s.T(), &s.handler2, s.EventDispatcher.handlers[s.event.GetName()][1])
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := s.EventDispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.EventDispatcher.handlers[s.event.GetName()]))

	err = s.EventDispatcher.Register(s.event.GetName(), &s.handler)
	s.Equal(EventNameAlreadyRegistered, err)
	s.Equal(1, len(s.EventDispatcher.handlers[s.event.GetName()]))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	err := s.EventDispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.EventDispatcher.handlers[s.event.GetName()]))

	err = s.EventDispatcher.Register(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.EventDispatcher.handlers[s.event.GetName()]))

	err = s.EventDispatcher.Register(s.event2.GetName(), &s.handler3)
	s.Nil(err)
	s.Equal(1, len(s.EventDispatcher.handlers[s.event2.GetName()]))

	s.EventDispatcher.Clear()
	s.Equal(0, len(s.EventDispatcher.handlers))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := s.EventDispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.EventDispatcher.handlers[s.event.GetName()]))

	err = s.EventDispatcher.Register(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.EventDispatcher.handlers[s.event.GetName()]))

	assert.True(s.T(), s.EventDispatcher.Has(s.event.GetName(), &s.handler))
	assert.True(s.T(), s.EventDispatcher.Has(s.event.GetName(), &s.handler2))
	assert.False(s.T(), s.EventDispatcher.Has(s.event.GetName(), &s.handler3))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	evtHdlr := &MockHandler{}
	evtHdlr.On("Handle", &s.event)

	err := s.EventDispatcher.Register(s.event.GetName(), evtHdlr)
	s.Nil(err)
	err = s.EventDispatcher.Dispatch(&s.event)
	s.Nil(err)

	evtHdlr.AssertExpectations(s.T())
	evtHdlr.AssertNumberOfCalls(s.T(), "Handle", 1)
}
