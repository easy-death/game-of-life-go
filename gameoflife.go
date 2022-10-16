package gameoflife

import (
	"context"
	"errors"
	"time"
)

type Config struct {
	TicksPerSecond uint8
}

type State [][]bool

type Engine struct {
	Config       Config
	State        [][]bool
	isChanelOpen bool
	lifeIsDead   bool
}

func NewEngine() *Engine {
	return &Engine{
		Config: Config{
			TicksPerSecond: 10,
		},
		State: [][]bool{},
	}
}

func (en *Engine) SetInitialState(state State) error {
	if len(en.State) > 0 {
		return errors.New("state changing not allowed")
	}
	en.State = state

	return nil
}

func (en *Engine) GetStateAfterSeconds(seconds uint) (State, error) {
	return en.GetStateAfterTicks(seconds * uint(en.Config.TicksPerSecond))
}

func (en *Engine) GetStateAfterTicks(ticks uint) (State, error) {
	if en.isChanelOpen {
		return nil, errors.New("can't change state when channel opened")
	}
	if len(en.State) == 0 {
		return nil, errors.New("GetStateAfterTicks called on empty state")
	}
	var i uint
	for i = 0; i < ticks; i++ {
		en.doTick()
	}

	return en.State, nil
}

func (en *Engine) ListenState(ctx context.Context) (chan State, error) {
	if len(en.State) == 0 {
		return nil, errors.New("ListenState called on empty state")
	}

	out := make(chan State)
	en.isChanelOpen = true

	go func() {
		ticker := time.NewTicker(time.Duration(1 / float64(en.Config.TicksPerSecond) * float64(time.Second)))
		for range ticker.C {
			select {
			case <-ctx.Done():
				return
			default:
				en.doTick()
				out <- en.State
			}
		}
	}()

	return out, nil
}

func (en *Engine) doTick() {
	if en.lifeIsDead {
		return
	}

	var countLiveNeighbours = func(state State, row int, col int) int {
		alive := 0
		for i := row - 1; i <= row+1; i++ {
			for j := col - 1; j <= col+1; j++ {
				var x int
				var y int
				if i == row && j == col {
					continue
				}
				if i < 0 {
					x = len(state) - 1
				} else if i >= len(state) {
					x = 0
				} else {
					x = i
				}

				if j < 0 {
					y = len(state[x]) - 1
				} else if j >= len(state[x]) {
					y = 0
				} else {
					y = j
				}

				if state[x][y] {
					alive++
				}
			}
		}

		return alive
	}
	totalAliveCells := 0

	for i := 0; i < len(en.State); i++ {
		for j := 0; j < len(en.State[i]); j++ {
			aliveNeighbours := countLiveNeighbours(en.State, i, j)
			if en.State[i][j] {
				totalAliveCells++

				if aliveNeighbours < 2 || aliveNeighbours > 3 {
					en.State[i][j] = false
				}
			} else {
				if aliveNeighbours == 3 {
					en.State[i][j] = true
				}
			}
		}
	}

	if totalAliveCells < 3 {
		en.lifeIsDead = true
	}
}
