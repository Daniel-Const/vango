package main

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
    canvas       Canvas
    colorcursor  int
    clicked      bool
}

var (
    OffsetY = 1
    OffsetX = 2
)

// TODO Get terminal dimensions + handle resize
const (
    Width = 100
    Height = 40
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
    return Model{
        canvas: Canvas{
            points: make(map[string] Point),
            pic: "",
            offsetx: OffsetX,
            offsety: OffsetY,
        },
        colorcursor: 0,
        clicked: false,
    }
}

func (m Model) Init() tea.Cmd {
    return nil
}


func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg :
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
        case "c":
            if m.colorcursor < len(colors) {
                m.colorcursor++
            }
        case "C":
            if m.colorcursor >= 0 {
                m.colorcursor--
            }
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
    m.canvas.AddPoint(x-OffsetX+1, y-OffsetY, colors[m.colorcursor])
    m.canvas.Draw()
}

func (m Model) View() string {
    border := lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

    title := lipgloss.NewStyle().
                Foreground(lipgloss.Color("40")).
                MarginLeft(OffsetX).
                SetString("Vango - Terminal Paint")

    highlight := lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	canvas := lipgloss.NewStyle().
		Border(border, true).
		BorderForeground(highlight).
        Width(Width).
        Height(Height).
        MarginLeft(OffsetX)

    info := lipgloss.NewStyle().
                Width(Width).
                Height(Height).
                Align(lipgloss.Left).
                MarginLeft(2)

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

    layout := lipgloss.JoinHorizontal(lipgloss.Top, canvas.Render(m.canvas.pic), info.Render(colorinfo.String()))
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
