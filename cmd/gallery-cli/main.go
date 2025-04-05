package main

import (
	"image/png"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	ascii "github.com/qeesung/image2ascii/convert"
)

func newApp() app {
	f, err := os.Open("images/puppy1.png")
	if err != nil {
		panic(err)
	}
	image, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	converter := ascii.NewImageConverter()
	asciiString := converter.Image2ASCIIString(image, &ascii.DefaultOptions)
	return app{asciiString: asciiString}
}

type app struct {
	asciiString string
}

func (a app) Init() tea.Cmd {
	return nil
}

func (a app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		if keyMsg.Type == tea.KeyCtrlC {
			return a, tea.Quit
		}
	}
	return a, nil
}

func (a app) View() string {
	return a.asciiString
}

func main() {
	app := newApp()
	prog := tea.NewProgram(app, tea.WithAltScreen())
	_, err := prog.Run()
	if err != nil {
		panic(err)
	}
}
