package main

import (
	"fmt"
	"log"
)

type Brush interface {
    Paint(c *Canvas, x int, y int)
}

type Erase struct {}
func (e Erase) Paint(c *Canvas, x int, y int) {
    col := x % 2
    row := y % 2
    if (col+row) % 2 == 0 {
        c.cells[y][x].setColor(EmptyPalette[0])
    } else {
        c.cells[y][x].setColor(EmptyPalette[1])
    }
    c.cells[y][x].empty = true
}

type Pen struct {}
func (p Pen) Paint(c *Canvas, x int, y int) {
    c.cells[y][x].setColor(c.color)
    c.cells[y][x].empty = false
}

type Bucket struct {}
func (b Bucket) Paint(c *Canvas, x int, y int) {
    visited := map[string]bool{}
    stack := []Pos{}
    color := c.cells[y][x].rawcolor
    isempty := c.cells[y][x].empty
    
    stack = append(stack, Pos{x, y})

    var p Pos
    
    for len(stack) > 0 {
        log.Println(len(stack))
        // Pop stack
        l := len(stack) - 1
        stack, p = stack[:l], stack[l]

        if !c.isValidPos(p.x, p.y) {
            continue
        }

        if isempty && !c.cells[p.y][p.x].empty {
            continue
        }

        if !isempty && c.cells[p.y][p.x].rawcolor != color {
            continue
        }

        if key := fmt.Sprintf("%d-%d", p.x, p.y); visited[key] {
            continue
        } else {
            log.Println("Bucket coloring!")
            c.cells[p.y][p.x].setColor(c.color)
            c.cells[p.y][p.x].empty = false
            visited[key] = true
        }

        log.Println("Pushing to stack")

        // Push to stack
        stack = append(stack, Pos{p.x-1, p.y})
        stack = append(stack, Pos{p.x, p.y-1})
        stack = append(stack, Pos{p.x+1, p.y})
        stack = append(stack, Pos{p.x, p.y+1})
    }
}

