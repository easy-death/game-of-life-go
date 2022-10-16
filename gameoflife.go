package gameoflife

import (
	"context"
	"time"
)

type Config struct {
	TicksPerSecond uint8
}

type Engine struct {
	Config Config
	State  [][]bool
}

func NewEngine() *Engine {
	return &Engine{
		Config: Config{
			TicksPerSecond: 10,
		},
		State: [][]bool{},
	}
}

func (en *Engine) SetInitialState(state [][]bool) error {
	en.State = state

	return nil
}

func (en *Engine) GetStateAfterSeconds(seconds int) {

}
func (en *Engine) GetStateAfterTicks(ticks int) {

}

func (en *Engine) ListenState(ctx *context.Context) chan [][]bool {
	out := make(chan [][]bool)

	go func() {
		ticker := time.NewTicker(time.Duration(en.Config.TicksPerSecond) * time.Second)
		for range ticker.C {
			en.doTick()
			out <- en.State
		}
	}()

	return out
}

func (en *Engine) doTick() {

}
