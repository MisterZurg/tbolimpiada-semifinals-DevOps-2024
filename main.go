package main

import (
	"fmt"

	"github.com/MisterZurg/tbolimpiada_semifinals_DevOps_2024/hexagon"
)

func main() {
	grid := hexagon.RectangleGrid(47, 21)
	for _, cell := range grid {
		fmt.Println(cell.GetColor())
	}

	// TODO: add winning logic
}
