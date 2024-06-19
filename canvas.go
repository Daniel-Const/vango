package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
    colora = lipgloss.Color("#424242")
    colorb = lipgloss.Color("#545353")
)

type Canvas struct {
    cells    [][]lipgloss.Color
    color    lipgloss.Color
    table     *table.Table
    width     int
    height    int
}

func NewCanvas(width int, height int) Canvas {
    c := Canvas{}
    c.cells = make([][]lipgloss.Color, Height)
    c.width = width
    c.height = height

    cells := make([][]lipgloss.Color, c.height)
    grid := make([][]string, c.height)
    for i := 0; i < c.height; i++ {
        cells[i] = make([]lipgloss.Color, c.width)
        grid[i] = make([]string, c.width)
    }

    c.cells = cells
    c.Clear()

	c.table = table.New().
		BorderRow(false).
		BorderColumn(false).
		Rows(grid...).
		StyleFunc(func(row, col int) lipgloss.Style {
            return lipgloss.NewStyle().Padding(0, 1).Background(c.cells[row-1][col]).Width(1).Height(1)
		})

    return c
}

func (c *Canvas) ClearCell(x int, y int) {
    col := x % 2
    row := y % 2
    if (col+row) % 2 == 0 {
        c.cells[y][x] = EmptyPalette[0] 
    } else {
        c.cells[y][x] = EmptyPalette[1] 
    }
}

func (c *Canvas) Clear() {
    for y := range c.cells {
        for x := range c.cells[y] {
            c.ClearCell(x, y)
        }
    }
}

func (c *Canvas) ColorCell(x int, y int) {
    c.cells[y][x] = c.color
}


func (c *Canvas) SetColor(color lipgloss.Color) {
    c.color = color
}

func (c Canvas) String() string {
    return c.table.Render()
}

