package view

import "github.com/jroimartin/gocui"

var g gocui.Gui
var m gocui.Manager

type manager struct {
}

func (manager) Layout(*gocui.Gui) error {
	panic("not implemented")
}

func SetupView() error {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		return err
	}
	defer g.Close()

	// Set manager and bindings

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}

	g.SetManager(gocui.Manager)

	return nil
}

func guiSettings(g *gocui.Gui) {
	g.FgColor = gocui.ColorWhite
	g.BgColor = gocui.ColorBlack
	g.Mouse = true
	g.Highlight = true
}
