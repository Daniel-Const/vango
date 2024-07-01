package main

import (
	"fmt"
	"log"
)

type Brush interface {
    Paint(c *Canvas, x int, y int, a *ActionStack)
}

type Erase struct {}
func (e Erase) Paint(c *Canvas, x int, y int, a *ActionStack) {
    if c.cells[y][x].isEmpty {
        return
    }

    a.push([]Action{{Pos{x, y}, c.cells[y][x]}})
    col := x % 2
    row := y % 2
    if (col+row) % 2 == 0 {
        c.cells[y][x].setColor(EmptyPalette[0])
    } else {
        c.cells[y][x].setColor(EmptyPalette[1])
    }
    c.cells[y][x].isEmpty = true
}

type Pen struct {}
func (p Pen) Paint(c *Canvas, x int, y int, a *ActionStack) {
    if c.cells[y][x].rawcolor == c.color {
        return
    }

    a.push([]Action{{Pos{x, y}, c.cells[y][x]}})
    c.cells[y][x].setColor(c.color)
    c.cells[y][x].isEmpty = false
}

type Bucket struct {}
func (b Bucket) Paint(c *Canvas, x int, y int, a *ActionStack) {
    actions := []Action{}
    visited := map[string]bool{}
    stack := []Pos{{x, y}}
    color := c.cells[y][x].rawcolor
    isempty := c.cells[y][x].isEmpty

    var p Pos
    for len(stack) > 0 {
        log.Println(len(stack))
        // Pop stack
        l := len(stack) - 1
        stack, p = stack[:l], stack[l]

        if !c.isValidPos(p.x, p.y) {
            continue
        }

        if isempty && !c.cells[p.y][p.x].isEmpty {
            continue
        }

        if !isempty && c.cells[p.y][p.x].rawcolor != color {
            continue
        }

        if key := fmt.Sprintf("%d-%d", p.x, p.y); visited[key] {
            continue
        } else {
            actions = append(actions, Action{Pos{p.x, p.y}, c.cells[p.y][p.x]})
            c.cells[p.y][p.x].setColor(c.color)
            c.cells[p.y][p.x].isEmpty = false
            visited[key] = true
        }

        // Push to stack
        stack = append(stack, Pos{p.x-1, p.y})
        stack = append(stack, Pos{p.x, p.y-1})
        stack = append(stack, Pos{p.x+1, p.y})
        stack = append(stack, Pos{p.x, p.y+1})
    }
    a.push(actions)
}

