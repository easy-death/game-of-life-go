# game-of-life-go

[![Go Reference](https://pkg.go.dev/badge/github.com/easy-death/game-of-life-go.svg)](https://pkg.go.dev/github.com/easy-death/game-of-life-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/easy-death/game-of-life-go)](https://goreportcard.com/report/github.com/easy-death/game-of-life-go)
[![Release](https://img.shields.io/github/v/release/easy-death/game-of-life-go.svg)](https://github.com/easy-death/game-of-life-go/releases/latest)

Go implementation of Conway's Game of Life. Possible to use as library or cli.
## Installation
Get latest version
```cmd
go get github.com/easy-death/game-of-life-go
```
Import
```go
import (
  gameoflife "github.com/easy-death/game-of-life-go"
)
```
## Usage
Library works in two modes: "time machine" and observer.
### Time machine
Returns state after given time delay or ticks count:
```go
engine := gameoflife.NewEngine()
state := gameoflife.State{
  gameoflife.StateRow{false, true, false},
  gameoflife.StateRow{false, true, false},
  gameoflife.StateRow{false, true, false},
}
engine.SetInitialState(state)
newState, err := engine.GetStateAfterTicks(10)

engine.Config.TicksPerSecond = 15
engine.GetStateAfterSeconds(10)
```
### Observer
State can be listened by multiple listeners with desired frequency:
```go
channel, err := engine.ListenState(context.Background())
if err == nil {
  for state := range channel {
    // ...
  }
}
```

## Cli usage
Navigate into `cmd/gameoflife`, exec `go build`, and you can call it with `./gameoflife`. Be aware that table is filled randomly,
restart to check how other state will look.
```text
Console app which will print 2 dimensional array of bools which will follow Conway's Game of life

Usage:
  gameoflife [flags]

Flags:
      --columns uint   How many columns in matrix (default 10)
  -h, --help           help for gameoflife
      --rows uint      How many rows in matrix (default 10)

```
