package main

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
    "golang.org/x/term"
)

var (
    palettes = []Palette{SimplePalette}
    OffsetY = 2
    OffsetX = 1
)

type Model struct {
    canvas       Canvas
    colorcursor  Cursor 
    brushcursor  Cursor 
    mousedown    bool
    palette      Palette
    keys         keyMap
    help         help.Model
}

func NewModel(width int, height int) Model {
    h := help.New()
    h.ShowAll = true
    p := SimplePalette
    c := NewCanvas(width, height)
    c.SetColor(p[0])
    return Model{
        canvas:      c,
        keys:        keys,
        help:        h,
        palette:     p,
        colorcursor: Cursor{Pos: 0, Min: 0, Max: len(p)-1},
        brushcursor: Cursor{Pos: 0, Min: 0, Max: 1},
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
            m.colorcursor.Next()
            m.canvas.SetColor(m.palette[m.colorcursor.Pos])
        case key.Matches(msg, m.keys.ColorUp):
            m.colorcursor.Prev()
            m.canvas.SetColor(m.palette[m.colorcursor.Pos])
        case key.Matches(msg, m.keys.BrushDown):
            m.brushcursor.Next()
            m.canvas.SetBrush(m.brushcursor.Pos)
        case key.Matches(msg, m.keys.BrushUp):
            m.brushcursor.Prev()
            m.canvas.SetBrush(m.brushcursor.Pos)

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



func (m Model) View() string {
   title := lipgloss.NewStyle().
                Foreground(lipgloss.Color("40")).
                MarginLeft(OffsetX).
                SetString("Vango - Terminal Paint")

    info := lipgloss.NewStyle().MarginLeft(4)
    colorinfo := strings.Builder{}
    subtitle := lipgloss.NewStyle().
                    BorderStyle(lipgloss.NormalBorder()).
                    BorderBottom(true)

    colorinfo.WriteString(subtitle.Render("Colors"))
    colorinfo.WriteString("\n")
    colorstyle := lipgloss.NewStyle().Background(lipgloss.Color("0"))

    for i, c := range m.palette {
        cursor := " "
        if i == m.colorcursor.Pos {
            cursor = "<"
        }
        colorstyle = colorstyle.Background(c)
        colorinfo.WriteString(fmt.Sprintf("%s %s", colorstyle.Render(" "), cursor))
        colorinfo.WriteRune('\n')
    }

    brushinfo := strings.Builder{}
    brushinfo.WriteString(subtitle.Render("Brushes"))
    brushinfo.WriteRune('\n')
    for i, c := range []string{"normal", "bucket"} {
        cursor := " "
        if i == m.brushcursor.Pos {
            cursor = "<"
        }
        brushinfo.WriteString(fmt.Sprintf("%s %s", c, cursor))
        brushinfo.WriteRune('\n')
    }

    helpstyle := lipgloss.NewStyle().MarginLeft(2)
    
    infolayout := lipgloss.JoinVertical(lipgloss.Top, colorinfo.String(), brushinfo.String())
    layout := lipgloss.JoinHorizontal(lipgloss.Top, lipgloss.NewStyle().Render(m.canvas.String()), info.Render(infolayout))
    return title.Render() + "\n" + layout + "\n" + helpstyle.Render(m.help.View(m.keys))
}

func (m *Model) Draw(x int, y int) {
    mapx := (x - OffsetX) / 2
    mapy := y - OffsetY
    m.canvas.ColorCell(mapx, mapy)
}

func (m Model) Resize(width int, height int) {
    m.canvas.Width = width
    m.canvas.Height = height
}

func main() {
    f, err := tea.LogToFile("debug.log", "debug")
    if err != nil {
        log.Fatal("Failed to create log file")
    }
    defer f.Close()
    
    // Terminal dimensions
    width, height, err := term.GetSize(0)
    if err != nil {
        log.Println("Error: Failed to get terminal dimensions")
        // Defaults
        width = 16
        height = 16
    }

    m := NewModel(width / 4, height / 2)
    p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseAllMotion())
    if _, err := p.Run(); err != nil {
        log.Fatal("Oops! Failed to start")
    }
}
