package main

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
    ColorUp key.Binding
    ColorDown key.Binding
    BrushUp key.Binding
    BrushDown key.Binding
    Clear key.Binding
    Quit key.Binding
}

var keys = keyMap {
    ColorDown: key.NewBinding(
        key.WithKeys("c"),
        key.WithHelp("c↓/C↑", "color select"),
    ),
    ColorUp: key.NewBinding(
        key.WithKeys("C"),
    ),
    BrushDown: key.NewBinding(
        key.WithKeys("b"),
        key.WithHelp("b↓/B↑", "brush select"),
    ),
    BrushUp: key.NewBinding(
        key.WithKeys("B"),
    ),
    Clear: key.NewBinding(
        key.WithKeys("w"),
        key.WithHelp("w", "wipe canvas"),
    ),
    Quit: key.NewBinding(
        key.WithKeys("ctrl+c", "q"),
        key.WithHelp("ctrl+c/q", "quit"),
    ),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.ColorDown, k.BrushDown, k.Clear, k.Quit},
	}
}

