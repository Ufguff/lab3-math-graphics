package types

import (
	"image/color"
	"log"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var scale = []Point{NewPoint(-1, -1), NewPoint(-1, 1), NewPoint(1, 1), NewPoint(1, -1)}

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

type Square struct {
	points []Point
	center Point
	move   []int
	len    float64
	color  color.RGBA
}

func NewSquare(point Point, len float64, color color.RGBA) *Square {
	arr := make([]Point, 0)
	c := NewPoint(point.x+(len/2), point.y+(len/2))

	arr = append(arr, NewPoint(point.x, point.y))
	arr = append(arr, NewPoint(point.x, point.y+len))
	arr = append(arr, NewPoint(point.x+len, point.y+len))
	arr = append(arr, NewPoint(point.x+len, point.y))

	return &Square{points: arr, center: c, move: []int{-1, 1}, len: len, color: color}

}

func (s *Square) DrawSquare() {

	rl.DrawLine(int32(s.points[0].x), int32(s.points[0].y), int32(s.points[1].x), int32(s.points[1].y), s.color)
	rl.DrawLine(int32(s.points[1].x), int32(s.points[1].y), int32(s.points[2].x), int32(s.points[2].y), s.color)
	rl.DrawLine(int32(s.points[2].x), int32(s.points[2].y), int32(s.points[3].x), int32(s.points[3].y), s.color)
	rl.DrawLine(int32(s.points[3].x), int32(s.points[3].y), int32(s.points[0].x), int32(s.points[0].y), s.color)
}

func (s *Square) Changes(angle float64) {

	log.Println("====================")
	for i := 0; i < 4; i++ {
		tempX := s.points[i].x
		tempY := s.points[i].y

		//s.points[i].x = (tempX+scale[i].x)*(math.Cos(angle)) - (tempY+scale[i].y)*(math.Sin(angle)) + s.center.x
		//s.points[i].y = (tempX+scale[i].x)*(math.Sin(angle)) + (tempY+scale[i].y)*(math.Cos(angle)) + s.center.y

		if s.points[i].x >= 800 {
			s.move[0] = 1
		} else if s.points[i].x <= 0 {
			s.move[0] = -1
		}

		if s.points[i].y >= 600 {
			s.move[1] = 1
		} else if s.points[i].y <= 0 {
			s.move[1] = -1
		}

		s.center.x += float64(s.move[0])
		s.center.y += float64(s.move[1])

		s.points[i].x = (tempX+scale[i].x+-s.center.x)*(math.Cos(angle)) - (tempY+scale[i].y-s.center.y)*(math.Sin(angle)) + s.center.x
		s.points[i].y = (tempX+scale[i].x-s.center.x)*(math.Sin(angle)) + (tempY+scale[i].y-s.center.y)*(math.Cos(angle)) + s.center.y

		log.Println(s.points[i].x, s.points[i].x)

	}
	log.Println("====================")
}
