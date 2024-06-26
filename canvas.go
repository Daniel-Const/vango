package main

import (
	"image"
	"image/color"
	"log"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Pos struct { x int; y int }
type Cell struct {
    color    lipgloss.Color
    empty    bool
    rawcolor string
}

func (c *Cell) setColor(color string) {
    c.color = lipgloss.Color(color)
    c.rawcolor = color
}

type Canvas struct {
    cells    [][]Cell
    color    string 
    table     *table.Table
    brush     Brush 
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
    c.brush = Pen{}
    return c
}

// Erase all cells on the canvas
func (c *Canvas) Clear() {
    erase := Erase{}
    for y := range c.Height {
        for x := range c.Width {
            erase.Paint(c, x, y)
            c.cells[y][x].empty = true
        }
    }
}

// Color a cell according to the selected brush
func (c *Canvas) ColorCell(x int, y int) {
    if c.isValidPos(x, y) {
        c.brush.Paint(c, x, y)
    }
}

func (c *Canvas) SetColor(color string) {
    c.color = color
}

func (c *Canvas) SetBrush(b Brush) {
    c.brush = b
}

func (c Canvas) isValidPos(x int, y int) bool {
    return x >= 0 && y >= 0 && x < c.Width && y < c.Height
}

func (c Canvas) String() string {
    return c.table.Render()
}

func (c Canvas) Image() image.Image {
    start := image.Point{0, 0}
    end := image.Point{c.Width, c.Height}
    img := image.NewRGBA(image.Rectangle{start, end})

    for y := range c.Height {
        for x := range c.Width {
            if !c.cells[y][x].empty {
                img.Set(x, y, hexToRGBA(c.cells[y][x].rawcolor))
            }
        }
    }

    return img
}

func hexToRGBA(hex string) color.Color {
    if hex[0] != '#' && len(hex) != 7 {
        log.Fatalf("Failed to convert color %s to RGBA (Invalid format)", hex)
    }

    val, err := strconv.ParseUint(hex[1:], 16, 32)
    if err != nil {
        log.Fatal(err)
    }
    
    // #RRGGBB
    // RRRR  RRRR  GGGG  GGGG  BBBB  BBBB
    
    // Red 
    // 0000  0000  0000  0000  RRRR  RRRR (>>16)

    // Green
    // 0000  0000  RRRR  RRRR  GGGG  GGGG (>> 8)
    // 0000  0000  0000  0000  GGGG  GGGG (& 0xFF)

    // Blue (& 0xFF)
    // 0000  0000  0000  0000  BBBB  BBBB (& 0xFF) 

    return color.RGBA{
        uint8(val >> 16),
        uint8((val >> 8) & 0xFF),
        uint8(val & 0xFF),
        0xff,
    }
}
