package view

import (
	"log"

	"github.com/jroimartin/gocui"
)

var g gocui.Gui
var m gocui.Manager

type manager struct {
}

func (manager) Layout(*gocui.Gui) error {
	panic("not implemented")
}

// TODO: move this elsewhere
type View struct {
	g *gocui.Gui
}

func (v *View) size() (int, int) {
	return v.g.size
}

func (v *View) layout(g *gocui.Gui) error {
	maxX, maxY := v.size()

	v.setCategoryPane(g)
	v.setListPane(g)
}

// Setup left side category window
func (v *View) setCategoryPane(g *gocui.Gui) error {
	if
}

// Setup right side list window
func (v *View) setListPane(g *gocui.Gui) error {

}

func (v *View) Run() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Fatalf("NewGui: %v", err)
	}
	defer g.Close()
	v.g = g

	g.defaultSettings()
	// Set manager and bindings
	g.SetManagerFunc(v.layout)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}

	g.SetManager(gocui.Manager)

	return nil
}

func (g *gocui.Gui) defaultSettings() {
	g.InputEsc = true
	g.FgColor = gocui.ColorWhite
	g.BgColor = gocui.ColorBlack
	g.Mouse = true
	g.Highlight = true
}
