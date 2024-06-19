package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const (
    NormalBrush = iota
    BucketBrush
    EraseBrush
)

type Color lipgloss.Color

type Cell struct {
    color   lipgloss.Color
    empty   bool
}

type Canvas struct {
    cells    [][]Cell
    color    lipgloss.Color
    table     *table.Table
    brush     int
    Width     int
    Height    int
}

func NewCanvas(width int, height int) Canvas {
    c := Canvas{}
    c.Width = width
    c.Height = height
    cells := make([][]Cell, c.Height)
    grid := make([][]string, c.Height)
    
    for i := 0; i < c.Height; i++ {
        cells[i] = make([]Cell, c.Width)
        grid[i] = make([]string, c.Width)
    }

    c.cells = cells
	c.table = table.New().
		BorderRow(false).
		BorderColumn(false).
		Rows(grid...).
		StyleFunc(func(row, col int) lipgloss.Style {
            return lipgloss.NewStyle().
                Padding(0, 1).
                Background(c.cells[row-1][col].color).
                Width(1).
                Height(1)
		})

    c.Clear()
    c.brush = NormalBrush
    return c
}

// Erase all cells on the canvas
func (c *Canvas) Clear() {
    for y := range c.cells {
        for x := range c.cells[y] {
            c.colorErase(x, y)
        }
    }
}

// Color a cell according to the selected brush
func (c *Canvas) ColorCell(x int, y int) {
    if c.isValidPos(x, y) {
        switch c.brush {
        case NormalBrush:
            c.colorNormal(x, y)
        case BucketBrush:
            var color *lipgloss.Color
            if cell := c.cells[y][x]; !cell.empty {
                color = &cell.color
            }
            c.colorBucket(x, y, color, map[string]bool{})
        case EraseBrush:
            c.colorErase(x, y)
        }
    }
}

func (c *Canvas) colorNormal(x int, y int) {
    c.cells[y][x].color = c.color
    c.cells[y][x].empty = false

}

// Floodfill paint
func (c *Canvas) colorBucket(x int, y int, color *lipgloss.Color, visited map[string]bool) {
    if !c.isValidPos(x, y) {
        return
    }
    
    // Cell should be empty
    if color == nil && !c.cells[y][x].empty {
       return 
    }
    
    // Color is wrong 
    if color != nil && c.cells[y][x].color != *color {
        return
    }

    if pointkey := fmt.Sprintf("%d-%d", x, y); visited[pointkey] {
        return
    } else {
        c.colorNormal(x, y)
        visited[pointkey] = true
    }

    c.colorBucket(x-1, y, color, visited)
    c.colorBucket(x, y-1, color, visited)
    c.colorBucket(x+1, y, color, visited)
    c.colorBucket(x, y+1, color, visited)
}

func (c *Canvas) colorErase(x int, y int) {
    col := x % 2
    row := y % 2
    if (col+row) % 2 == 0 {
        c.cells[y][x].color = EmptyPalette[0]
        c.cells[y][x].empty = true
    } else {
        c.cells[y][x].color = EmptyPalette[1]
        c.cells[y][x].empty = true
    }
}

func (c *Canvas) SetColor(color lipgloss.Color) {
    c.color = color
}

func (c *Canvas) SetBrush(b int) {
    c.brush = b
}

func (c Canvas) isValidPos(x int, y int) bool {
    return x >= 0 && y >= 0 && x < c.Width && y < c.Height
}

func (c Canvas) String() string {
    return c.table.Render()
}

