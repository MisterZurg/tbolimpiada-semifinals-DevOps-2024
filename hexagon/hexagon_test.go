package hexagon_test

import (
	"testing"

	"github.com/MisterZurg/tbolimpiada_semifinals_DevOps_2024/hexagon"
)

func TestHex_GetColor(t *testing.T) {
	tests := map[string]struct {
		color     int
		expResult string
	}{
		"White":         {hexagon.White, "White"},
		"Red":           {hexagon.Red, "Red"},
		"Green":         {hexagon.Green, "Green"},
		"Lime":          {hexagon.Lime, "Lime"},
		"Blue":          {hexagon.Blue, "Blue"},
		"Cayan":         {hexagon.Cayan, "Cayan"},
		"Yellow":        {hexagon.Yellow, "Yellow"},
		"Pink":          {hexagon.Pink, "Pink"},
		"Orange":        {hexagon.Orange, "Orange"},
		"Unimplemented": {1337, ""},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := hexagon.Hex{Color: test.color}.GetColor()
			if got != test.expResult {
				t.Errorf("GetColor(%d) = %s; want %s",
					test.color,
					got,
					test.expResult,
				)
			}
		})
	}
}
