package main

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
    ColorUp key.Binding
    ColorDown key.Binding
    Clear key.Binding
    Quit key.Binding
}

var keys = keyMap {
    ColorDown: key.NewBinding(
        key.WithKeys("c"),
        key.WithHelp("c", "Change color down"),
    ),
    ColorUp: key.NewBinding(
        key.WithKeys("C"),
        key.WithHelp("C", "Change color up"),
    ),
    Clear: key.NewBinding(
        key.WithKeys("w"),
        key.WithHelp("w", "Wipe canvas"),
    ),
    Quit: key.NewBinding(
        key.WithKeys("ctrl+c", "q"),
        key.WithHelp("ctrl+c/q", "Quit"),
    ),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.ColorDown, k.ColorUp, k.Clear, k.Quit},
	}
}

