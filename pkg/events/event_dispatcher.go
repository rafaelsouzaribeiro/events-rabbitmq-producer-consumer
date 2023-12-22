package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("Handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispartcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}

}

func (e *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if handlers, ok := e.handlers[eventName]; ok {
		for i, v := range handlers {
			if v == handler {
				/// Vai adicionando mas +1 e o restante ... de acordo com :1 primeiro item
				e.handlers[eventName] = append(handlers[:i], handlers[i+1:]...)
				return nil
			}
		}
	}
	return nil
}

func (e *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := e.handlers[eventName]; ok {
		for _, v := range e.handlers[eventName] {
			if v == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	e.handlers[eventName] = append(e.handlers[eventName], handler)
	return nil
}

func (e *EventDispatcher) Clear() {

	e.handlers = make(map[string][]EventHandlerInterface)
}

func (e *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := e.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}

		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}

	return nil
}

func (e *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {

	if _, ok := e.handlers[eventName]; ok {
		for _, v := range e.handlers[eventName] {
			if v == handler {
				return true
			}
		}
	}

	return false
}
