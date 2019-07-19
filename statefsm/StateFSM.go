package statefsm

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"sync"
)

type FSMHandler func() (string, error)

//Finite State Machine
type FSM struct {
	mu           sync.Mutex
	currentState string
	flowDiagram  map[string][]string //State Machine Map  {key:state,value:event}
}

func (f *FSM) setState(newState string) {
	f.currentState = newState
}

func NewFSM(initState string) *FSM {
	return &FSM{
		currentState: initState,
		flowDiagram:  make(map[string][]string),
	}
}

func (f *FSM) AddHandler(state string, events []string) *FSM {
	if _, ok := f.flowDiagram[state]; !ok {
		f.flowDiagram[state] = make([]string, len(events))
	}
	if _, ok := f.flowDiagram[state]; ok {
		log.Warnf("The state (%s) event (%s) has been defined.", state, events)
	}
	f.flowDiagram[state] = events
	return f
}

func (f *FSM) Call(event string, fsmHandler FSMHandler) error {
	if f == nil || f.currentState == "" {
		return errors.New("FSM data error.")
	}
	if event == "" {
		return errors.New("EventState is nil.")
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	events := f.flowDiagram[f.currentState]
	if events == nil || len(events) == 0 {
		return errors.New("Event undefined. State : " + f.currentState)
	}
	for _, e := range events {
		if e == event {
			log.Infof("State changed from %s to %s", f.currentState, event)
			state, err := fsmHandler()
			if err != nil {
				return err
			}
			f.setState(state)
			return nil
		}
	}
	return errors.New("State transition error." + f.currentState + " -> " + event)
}
