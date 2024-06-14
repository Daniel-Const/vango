package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Point struct {
    X     int
    Y     int
    Color lipgloss.Color 
}

func (p Point) Key() string {
    return fmt.Sprintf("%d:%d", p.X, p.Y)
}

func PointKey(x int, y int) string {
    return fmt.Sprintf("%d:%d", x, y)
}

type Canvas struct {
    points   map[string] Point
    pic      string
    offsetx  int
    offsety  int
}

func (c *Canvas) AddPoint(x int, y int, color lipgloss.Color) {
    p := Point{x, y, color}
    c.points[p.Key()] = p
}

func (c *Canvas) Draw() {
    style := lipgloss.NewStyle().Background(lipgloss.Color(lipgloss.Color("0")))
    view := strings.Builder{}
    for y := OffsetY; y <= Height; y++ {
        for x := OffsetX; x <= Width; x++ {
            if p, ok := c.points[PointKey(x, y)]; ok {
                style = style.Background(p.Color)
                view.WriteString(style.Render((" ")))
            } else {
                view.WriteString(" ")
            }
        }
        view.WriteString("\n")
    }
    c.pic = view.String()
}
