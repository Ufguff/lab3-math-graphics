package types

import (
	"image/color"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ufguff/config"
)

var scale = []Point{NewPoint(-1, -1), NewPoint(-1, 1), NewPoint(1, 1), NewPoint(1, -1)}
var scaleLen = 1.0

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

func NewSquare(c Point, len float64, color color.RGBA) *Square {
	arr := make([]Point, 0)

	arr = append(arr, NewPoint(c.x-(len/2), c.y-(len/2)))
	arr = append(arr, NewPoint(c.x+(len/2), c.y-(len/2)))
	arr = append(arr, NewPoint(c.x+(len/2), c.y+(len/2)))
	arr = append(arr, NewPoint(c.x-(len/2), c.y+(len/2)))

	return &Square{points: arr, center: c, move: []int{1, 1}, len: len, color: color}

}

func (s *Square) DrawSquare() {
	rl.DrawCircle(int32(s.center.x), int32(s.center.y), 3, s.color)

	rl.DrawLine(int32(s.points[0].x), int32(s.points[0].y), int32(s.points[1].x), int32(s.points[1].y), s.color)
	rl.DrawLine(int32(s.points[1].x), int32(s.points[1].y), int32(s.points[2].x), int32(s.points[2].y), s.color)
	rl.DrawLine(int32(s.points[2].x), int32(s.points[2].y), int32(s.points[3].x), int32(s.points[3].y), s.color)
	rl.DrawLine(int32(s.points[3].x), int32(s.points[3].y), int32(s.points[0].x), int32(s.points[0].y), s.color)
}

func (s *Square) Changes(angle float64) {
	s.points[0] = NewPoint(s.center.x-(s.len/2), s.center.y-(s.len/2))
	s.points[1] = NewPoint(s.center.x+(s.len/2), s.center.y-(s.len/2))
	s.points[2] = NewPoint(s.center.x+(s.len/2), s.center.y+(s.len/2))
	s.points[3] = NewPoint(s.center.x-(s.len/2), s.center.y+(s.len/2))

	log.Println("===================={begin it}")
	for i := 0; i < 4; i++ {
		//tCenterX := s.center.x
		//tCenterY := s.center.y

		//s.points[i].x = (tempX-s.center.x)*(math.Cos(angle)) - (tempY+s.center.y)*(math.Sin(angle)) + s.center.x
		//s.points[i].y = (tempX+s.center.x)*(math.Sin(angle)) + (tempY+s.center.y)*(math.Cos(angle)) + s.center.y

		//s.points[i].x = (tempX+float64(s.move[0])-s.center.x)*(math.Cos(angle)) - (tempY+float64(s.move[1])-s.center.y)*(math.Sin(angle)) + s.center.x
		//s.points[i].y = (tempX+float64(s.move[0])-s.center.x)*(math.Sin(angle)) + (tempY+float64(s.move[1])-s.center.y)*(math.Cos(angle)) + s.center.y

		//s.points[i].x = (s.points[i].x-s.center.x)*(math.Cos(angle)) - (s.points[i].y-s.center.y)*(math.Sin(angle)) + s.center.x
		//s.points[i].y = (s.points[i].x-s.center.x)*(math.Sin(angle)) + (s.points[i].y-s.center.y)*(math.Cos(angle)) + s.center.y

		s.points[i].x += float64(s.move[0])
		s.points[i].y += float64(s.move[1])

		s.center.x += float64(s.move[0])
		s.center.y += float64(s.move[1])

		s.len += scaleLen

		if s.points[i].x >= config.W {
			s.move[0] = -1
		} else if s.points[i].x <= 0 {
			s.move[0] = 1
		}

		if s.points[i].y >= config.H {
			s.move[1] = -1
		} else if s.points[i].y <= 0 {
			s.move[1] = 1
		}

		if s.len == 200 {
			scaleLen = -1.0
		} else if s.len == 50 {
			scaleLen = 1.0
		}

		log.Println(i, "-- ", s.points[i].x, s.points[i].x)

	}
	log.Println("===================={end it}")
}
