package main

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
    canvas       Canvas
    colorcursor  int
    mousedown      bool
    palette      Palette
    keys         keyMap
    help         help.Model
}

var (
    OffsetY = 2
    OffsetX = 1
)

// TODO Get terminal dimensions + handle resize
const (
    Width = 24 
    Height = 16
)

var (
    palettes = []Palette{SimplePalette}
)

func NewModel() Model {
    h := help.New()
    h.ShowAll = true
    p := SimplePalette
    c := NewCanvas(Width, Height)
    c.SetColor(p[0])
    return Model{
        canvas:      c,
        keys:        keys,
        help:        h,
        colorcursor: 0,
        palette:     p,
    }
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg :
        switch {
        case key.Matches(msg, m.keys.Quit):
            return m, tea.Quit

        case key.Matches(msg, m.keys.ColorDown):
            if m.colorcursor < len(m.palette)-1 {
                m.colorcursor++
                m.canvas.SetColor(m.palette[m.colorcursor])
            } else {
                m.colorcursor = 0
                m.canvas.SetColor(m.palette[m.colorcursor])
            }

        case key.Matches(msg, m.keys.ColorUp):
            if m.colorcursor > 0 {
                m.colorcursor--
                m.canvas.SetColor(m.palette[m.colorcursor])
            } else {
                m.colorcursor = len(m.palette)-1
                m.canvas.SetColor(m.palette[m.colorcursor])
            }

        case key.Matches(msg, m.keys.Clear):
            m.canvas.Clear()
        }

    case tea.MouseMsg:
        switch msg.Action {
        case tea.MouseActionPress:
            m.Draw(msg.X, msg.Y)
            m.mousedown = true

        case tea.MouseActionMotion:
            if m.mousedown {
                m.Draw(msg.X, msg.Y)
            }

        case tea.MouseActionRelease:
            m.mousedown = false
        }
    }
    return m, nil
}

func (m *Model) Draw(x int, y int) {
    mapx := (x - OffsetX) / 2
    mapy := y - OffsetY
    if mapx >= 0 && mapy >= 0 && mapx < Width && mapy < Height {
        log.Println(fmt.Sprintf("coloring: (%d, %d)", x, y))
        m.canvas.ColorCell(mapx, mapy)
    }
}

func (m Model) View() string {
   title := lipgloss.NewStyle().
                Foreground(lipgloss.Color("40")).
                MarginLeft(OffsetX).
                SetString("Vango - Terminal Paint")

    info := lipgloss.NewStyle().MarginLeft(4)

	canvas := lipgloss.NewStyle()
    colorinfo := strings.Builder{}
    colorinfo.WriteString("Colors\n\n")
    colorstyle := lipgloss.NewStyle().Background(lipgloss.Color("0"))

    for i, c := range m.palette {
        cursor := " "
        if i == m.colorcursor {
            cursor = "<"
        }
        colorstyle = colorstyle.Background(c)
        colorinfo.WriteString(fmt.Sprintf("%s %s", colorstyle.Render(" "), cursor))
        colorinfo.WriteRune('\n')
    }
    colorinfo.WriteString("\n"+m.help.View(m.keys))

    layout := lipgloss.JoinHorizontal(lipgloss.Top, canvas.Render(m.canvas.String()), info.Render(colorinfo.String()))
    return title.Render() + "\n" + layout
}

func main() {
    f, err := tea.LogToFile("debug.log", "debug")
    if err != nil {
        log.Fatal("Failed to create log file")
    }
    defer f.Close()

    p := tea.NewProgram(NewModel(), tea.WithAltScreen(), tea.WithMouseAllMotion())
    if _, err := p.Run(); err != nil {
        log.Fatal("Oops! Failed to start")
    }
}
