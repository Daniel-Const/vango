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
    offsetx  int
    offsety  int
    color    lipgloss.Color
    table     *table.Table 
}

func NewCanvas(offsetx int, offsety int) Canvas {
    c := Canvas{}
    c.cells = make([][]lipgloss.Color, Height)
    c.offsetx = offsetx
    c.offsety = offsety
    
    c.SetColor(colors[0])

    cells := make([][]lipgloss.Color, Height)
    grid := make([][]string, Height)
    for i := 0; i < Height; i++ {
        cells[i] = make([]lipgloss.Color, Width)
        grid[i] = make([]string, Width)
    }

    for h := range cells{
        for w := range cells[h] {
            grid[h][w] = ""
            c := w % 2
            r := h % 2
            if (c+r) % 2 == 0 {
                cells[h][w] = colora
            } else {
                cells[h][w] = colorb
            }
        }
    }

    c.cells = cells
	c.table = table.New().
		BorderRow(false).
		BorderColumn(false).
		Rows(grid...).
		StyleFunc(func(row, col int) lipgloss.Style {
            return lipgloss.NewStyle().Padding(0, 1).Background(c.cells[row-1][col]).Width(1).Height(1)
		})

    return c
}

func (c *Canvas) Clear() {
    for h := range c.cells{
        for w := range c.cells[h] {
            col := w % 2
            row := h % 2
            if (col+row) % 2 == 0 {
                c.cells[h][w] = colora
            } else {
                c.cells[h][w] = colorb
            }
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

