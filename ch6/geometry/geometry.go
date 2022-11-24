package geometry

import (
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type IntList struct {
	Value int
	Tail  *IntList
}

type ColorPoint struct {
	Point
	Color color.RGBA
}

type Path []Point

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point

	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}

	}
	return sum
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}

	return list.Value + list.Tail.Sum()
}

type Logger struct {
	flags  int
	prefix string
}

func (l *Logger) Flags() int {
	return l.flags
}
func (l *Logger) SetFlag(flag int) {
	l.flags = flag
}
func (l *Logger) Prefix() string {
	return l.prefix
}
func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
}
