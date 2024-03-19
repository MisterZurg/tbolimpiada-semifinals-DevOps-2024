package hexagon

import (
	"math"
	"math/rand"
)

type direction int

const (
	directionSE = iota
	directionNE
	directionN
	directionNW
	directionSW
	directionS
)

var directions = []Hex{
	NewHex(1, 0),
	NewHex(1, -1),
	NewHex(0, -1),
	NewHex(-1, 0),
	NewHex(-1, +1),
	NewHex(0, +1),
}

// Hex - структура представляющая Гекс.
type Hex struct {
	q     int
	r     int
	s     int
	Color int
}

// NewHex - конструктор, возвращающий Гекс заполненный случайным цветом.
func NewHex(q, r int) Hex {
	color := rand.Intn(Orange-White) + White
	return Hex{q: q, r: r, s: -q - r, Color: color}
}

// Список доступных цветов для Гекса.
const (
	White  = iota + 1 // 1. Белый
	Red               // 2. Ярко-красный
	Green             // 3. Зеленый
	Lime              // 4. Ярко-зеленый
	Blue              // 5. Синий
	Cayan             // 6. Светло-синий
	Yellow            // 7. Желтый
	Pink              // 8. Розовый
	Orange            // 9. Оранжевый
	// Black             // 10. Черный (стены, цвет не доступен для выбора)
)

// GetColor - возвращает текстовое представление цвета.
func (h Hex) GetColor() string {
	switch h.Color {
	case White:
		return "White"
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Lime:
		return "Lime"
	case Blue:
		return "Blue"
	case Cayan:
		return "Cayan"
	case Yellow:
		return "Yellow"
	case Pink:
		return "Pink"
	case Orange:
		return "Orange"
	default:
		return ""
	}
}

// Neighbor - возвращает соседей Hex относительно заданного направления.
func (h Hex) Neighbor(direction direction) Hex {
	directionOffset := directions[direction]
	return NewHex(h.q+directionOffset.q, h.r+directionOffset.r)
}

// RectangleGrid - возвращает слайс Hex образующих прямоугольник заданной ширины и высоты.
func RectangleGrid(width, height int) []Hex {
	grid := make([]Hex, 0)

	for q := 0; q < width; q++ {
		qOffset := int(math.Floor(float64(q) / 2.))

		for r := -qOffset; r < height-qOffset; r++ {
			grid = append(grid, NewHex(q, r))
		}
	}

	return grid
}
