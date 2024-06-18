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
    clicked      bool
    keys         keyMap
    help         help.Model
}

var (
    OffsetY = 1
    OffsetX = 0
)

// TODO Get terminal dimensions + handle resize
const (
    Width = 16 
    Height = 16
)

var (
   colors = [6]lipgloss.Color{
        lipgloss.Color(lipgloss.Color("100")),
        lipgloss.Color(lipgloss.Color("120")),
        lipgloss.Color(lipgloss.Color("140")),
        lipgloss.Color(lipgloss.Color("160")),
        lipgloss.Color(lipgloss.Color("180")),
        lipgloss.Color(lipgloss.Color("190")),
    }
)

func NewModel() Model {
    h := help.New()
    h.ShowAll = true
    return Model{
        canvas:      NewCanvas(OffsetX, OffsetY),
        keys:        keys,
        help:        h,
        colorcursor: 0,
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
            if m.colorcursor < len(colors)-1 {
                m.colorcursor++
                m.canvas.SetColor(colors[m.colorcursor])
            }

        case key.Matches(msg, m.keys.ColorUp):
            if m.colorcursor > 0 {
                m.colorcursor--
                m.canvas.SetColor(colors[m.colorcursor])
            }
        case key.Matches(msg, m.keys.Clear):
            m.canvas.Clear()
        }
    case tea.MouseMsg:
        switch msg.Action {
        case tea.MouseActionPress:
            m.DrawPoint(msg.X, msg.Y)
            m.clicked = true
        case tea.MouseActionRelease:
            m.clicked = false
        case tea.MouseActionMotion:
            if m.clicked {
                m.DrawPoint(msg.X, msg.Y)
            }
        }
    }
    return m, nil
}

func (m *Model) DrawPoint(x int, y int) {
    mapx := (x - 1) / 2
    mapy := y - (1 + m.canvas.offsety)
    if mapx < Width && mapy < Height {
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

    for i, c := range colors {
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
