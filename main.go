package main

import (
	"fmt"

	"github.com/MisterZurg/tbolimpiada-semifinals-DevOps-2024/hexagon"
)

func main() {
	grid := hexagon.RectangleGrid(47, 21)
	for _, cell := range grid {
		fmt.Println(cell.GetColor())
	}

	// TODO: add winning logic
}
