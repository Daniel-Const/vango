package main

import "github.com/charmbracelet/lipgloss"

type Palette []lipgloss.Color

var (
    EmptyPalette = Palette{
        lipgloss.Color("#424242"),
        lipgloss.Color("#545353"),
    }

    SimplePalette = Palette{
        lipgloss.Color(lipgloss.Color("#000000")),
        lipgloss.Color(lipgloss.Color("#db0fb3")),
        lipgloss.Color(lipgloss.Color("#9912c9")),
        lipgloss.Color(lipgloss.Color("#1218c9")),
        lipgloss.Color(lipgloss.Color("#12b7c9")),
        lipgloss.Color(lipgloss.Color("#c9c312")),
        lipgloss.Color(lipgloss.Color("#c93a12")),
    }
)
