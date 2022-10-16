package main

import (
	"context"
	"fmt"
	"gameoflife"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var (
	rows uint
	cols uint
)

func main() {
	cmd := cobra.Command{
		Run:   haveFun,
		Use:   "gameoflife",
		Short: "Golang implementation of Conway's Game of life",
		Long:  "Console app which will print 2 dimensional array of bools which will follow Conway's Game of life",
	}
	cmd.PersistentFlags().UintVar(&rows, "rows", 10, "How many rows in matrix")
	cmd.PersistentFlags().UintVar(&cols, "columns", 10, "How many columns in matrix")
	cmd.Execute()
}

func haveFun(_ *cobra.Command, _ []string) {
	var format = func(in bool) string {
		if in {
			return "*"
		} else {
			return " "
		}
	}
	state := make([][]bool, rows)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < int(rows); i++ {
		state[i] = make([]bool, cols)
		for j := 0; j < int(cols); j++ {
			state[i][j] = r.Intn(10) == 0
		}
	}

	engine := gameoflife.NewEngine()
	engine.SetInitialState(state)
	engine.Config.TicksPerSecond = 15
	ctx := context.Background()
	channel, _ := engine.ListenState(ctx)
	for s := range channel {
		for _, x := range s {
			for _, y := range x {
				fmt.Print("|", format(y))
			}
			fmt.Print("|\n")
		}
		fmt.Print("\n\n\n")
	}
}
