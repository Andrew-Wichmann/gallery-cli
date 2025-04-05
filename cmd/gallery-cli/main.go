package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	ascii "github.com/qeesung/image2ascii/convert"
)

func newApp() app {
	var gallery gallery
	files, err := os.ReadDir("images")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			f, err := os.Open(fmt.Sprintf("%s/%s", "images", file.Name()))
			if err != nil {
				panic(err)
			}
			image, err := png.Decode(f)
			if err != nil {
				panic(err)
			}
			gallery.Insert(image)
		}
	}
	return app{gallery: gallery}
}

type galleryItem struct {
	prev        *galleryItem
	next        *galleryItem
	asciiString string
}

type gallery struct {
	tail *galleryItem
	head *galleryItem
}

func (g *gallery) Insert(image image.Image) {
	converter := ascii.NewImageConverter()
	asciiString := converter.Image2ASCIIString(image, &ascii.DefaultOptions)
	gi := galleryItem{asciiString: asciiString}
	if g.head == nil {
		if g.tail != nil {
			panic("Head of the gallery is nil but the tail isn't")
		}
		g.head = &gi
		g.tail = &gi
		return
	}
	prevHead := g.head
	g.head = &gi
	g.head.next = prevHead
	prevHead.prev = g.head
	g.head.prev = g.tail
	g.tail.next = g.head
}

func (g gallery) View() string {
	return g.head.asciiString
}

func (g *gallery) Next() {
	g.head = g.head.next
}

func (g *gallery) Prev() {
	g.head = g.head.prev
}

type app struct {
	gallery gallery
}

func (a app) Init() tea.Cmd {
	return nil
}

func (a app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		if keyMsg.Type == tea.KeyCtrlC {
			return a, tea.Quit
		}
		if keyMsg.String() == "n" {
			a.gallery.Next()
		}
		if keyMsg.String() == "p" {
			a.gallery.Prev()
		}
	}
	return a, nil
}

func (a app) View() string {
	return a.gallery.View()
}

func main() {
	app := newApp()
	prog := tea.NewProgram(app, tea.WithAltScreen())
	_, err := prog.Run()
	if err != nil {
		panic(err)
	}
}
