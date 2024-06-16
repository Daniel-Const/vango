package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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
    cells    []string
    offsetx  int
    offsety  int
    cellstyle lipgloss.Style
    table     *table.Table 
}

func NewCanvas(offsetx int, offsety int) Canvas {
    c := Canvas{}
    c.cells = make([]string, Height*Width + Height)
    c.offsetx = offsetx
    c.offsety = offsety
    c.cellstyle = lipgloss.NewStyle().SetString(" ")
    
    // Init cells
    for i := range c.cells {
        if i % (Width+1) == 0 {
            c.cells[i] = "\n"
            continue
        }
        c.cells[i] = c.cellstyle.Render() 
    }
    c.cells[0] = c.cellstyle.Render()

    c.SetColor(colors[0])
    return c
}

func (c *Canvas) Clear() {
    clearstyle := lipgloss.NewStyle().SetString(" ")
    for i := range c.cells {
        if c.cells[i] == "\n" {
            continue
        }
        c.cells[i] = clearstyle.Render()
    }
}

func (c *Canvas) AddPoint(x int, y int) {
    if pos, valid := c.mapPos(x, y); valid {
        c.cells[pos] = c.cellstyle.Render()
    }
}

func (c *Canvas) mapPos(x int, y int) (int, bool) {
    newx := x-c.offsetx
    newy := y-c.offsety
    if newx < 0 || newx > Width || newy < 0 || newy > Height {
        return 0, false
    }

    point := newy*(Width+1)+newx
    if point >= 0 && point < Height*Width+Height {
        if c.cells[point] == "\n" {
            log.Println("Ah newline here")
            return 0, false
        }

        log.Printf("Adding (%d, %d) => (%d, %d)", x, y, newx, newy)
        return point, true
    }
    
    return 0, false
}  
func (c *Canvas) SetColor(color lipgloss.Color) {
    c.cellstyle = c.cellstyle.Background(color)
}

func (c Canvas) String() string {
    return strings.Join(c.cells, "")
}

