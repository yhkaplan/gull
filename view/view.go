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
	rightOffset := 0

	err := v.setCategoryView(g)
	if err != nil {
		return err
	}
	err = v.setListView(g)
	if err != nil {
		return err
	}

	return nil
}

// Setup left side category window
func (mainView *View) setCategoryView(g *gocui.Gui) error {
	if categoryView, err := g.SetView("Category", 0, 0, maxX/2, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		mainView.categoryView
	}
}

// Setup right side list window
func (v *View) setListView(g *gocui.Gui) error {

}

func (v *View) Run() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Fatalf("NewGui: %v", err)
	}
	defer g.Close()
	v.g = g

	defaultSettings(g)
	// Set manager and bindings
	g.SetManagerFunc(v.layout)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalf("MainLoop: %v", err)
	}

	g.SetManagerFunc(v.layout)
}

func defaultSettings(g *gocui.Gui) {
	g.InputEsc = true
	g.FgColor = gocui.ColorWhite
	g.BgColor = gocui.ColorBlack
	g.Mouse = true
	g.Highlight = true
}
