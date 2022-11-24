package Demo

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

type PointItem struct {
	Z    int
	Desc string
}
