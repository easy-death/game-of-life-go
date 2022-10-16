package gameoflife

import (
	"context"
	"testing"
	"time"
)

func TestStateChange(t *testing.T) {
	en := NewEngine()
	state := State{
		[]bool{true, true},
		[]bool{true, true},
	}
	er1 := en.SetInitialState(state)
	er2 := en.SetInitialState(state)

	if er1 != nil {
		t.Error("failed to set state")
	}
	if er2 == nil {
		t.Error("possible to change state")
	}
}

func TestTicksOnEmptyState(t *testing.T) {
	en := NewEngine()
	_, err := en.GetStateAfterTicks(1)
	if err == nil {
		t.Error("empty engine do ticks")
	}
}

func TestTicksPerSecondRespected(t *testing.T) {
	en := NewEngine()
	en.SetInitialState(State{
		[]bool{true, true},
		[]bool{true, true},
	})
	en.Config.TicksPerSecond = 2
	todoCtx := context.TODO()
	ctx, cancel := context.WithTimeout(todoCtx, time.Second*2)
	defer cancel()
	channel, _ := en.ListenState(ctx)
	counter := 0
	expectCounter := 4
	for range channel {
		counter++
	}

	if counter != expectCounter {
		t.Error("bad tickrate")
	}
}

func TestDoTick(t *testing.T) {
	en := NewEngine()
	state := State{
		StateRow{true, false, false},
		StateRow{true, false, false},
		StateRow{true, false, false},
	}
	expectedState := State{
		StateRow{true, true, true},
		StateRow{true, true, true},
		StateRow{true, true, true},
	}
	en.SetInitialState(state)
	newState, _ := en.GetStateAfterTicks(1)
	if !compareStates(newState, expectedState) {
		t.Error("bad life calculation")
	}
}

func TestListenEmptyState(t *testing.T) {
	en := NewEngine()

	_, err := en.ListenState(context.Background())
	if err == nil {
		t.Error("listen empty state make no sence")
	}
}

func TestGetStateAfterSeconds(t *testing.T) {
	en1 := NewEngine()
	en2 := NewEngine()
	state := State{
		StateRow{false, true, false},
		StateRow{false, true, true},
		StateRow{false, true, false},
	}
	en1.SetInitialState(state)
	en2.SetInitialState(state)

	en1.Config.TicksPerSecond = 3
	r1, _ := en1.GetStateAfterSeconds(1)
	r2, _ := en2.GetStateAfterTicks(3)

	if !compareStates(r1, r2) {
		t.Error("GetStateAfterSeconds fails tickrate")
	}
}

func TestChangeWithListeners(t *testing.T) {
	en := NewEngine()
	en.SetInitialState(State{
		StateRow{false, true, false},
		StateRow{false, true, true},
		StateRow{false, true, false},
	})

	en.ListenState(context.Background())

	_, err := en.GetStateAfterSeconds(1)
	if err == nil {
		t.Error("can't interupt listeners by direct state change")
	}
}

func compareStates(s1 State, s2 State) bool {
	if len(s1) == len(s2) {
		for i := 0; i < len(s1); i++ {
			if len(s1[i]) == len(s2[i]) {
				for j := 0; j < len(s1[i]); j++ {
					if s1[i][j] != s2[i][j] {
						return false
					}
				}
			} else {
				return false
			}
		}
	} else {
		return false
	}
	return true
}
