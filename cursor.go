package main

type Cursor struct {
    Pos int
    Max int
    Min int
}

func (c *Cursor) Next() {
    if c.Pos < c.Max {
        c.Pos++
    } else {
        c.Pos = 0
    }
}

func (c *Cursor) Prev() {
    if c.Pos > 0 {
        c.Pos--
    } else {
        c.Pos = c.Max
    }
}
