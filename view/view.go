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
type DashboardView struct {
	g         *gocui.Gui
	categoryV *gocui.View
}

// Initializes GullView
func New() *DashboardView {
	v := &DashboardView{}
	return v
}

// Returns window's width and height
func (v *DashboardView) size() (int, int) {
	return v.g.Size()
}

func (v *DashboardView) layout(g *gocui.Gui) error {
	maxX, maxY := v.size()
	horizOffset := maxX / 2

	err := v.setCategoryView(g, horizOffset, maxY)
	if err != nil {
		return err
	}
	// err = v.setListView(g)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// Setup left side category window
func (v *DashboardView) setCategoryView(g *gocui.Gui, horizOffset int, maxY int) error {
	if categoryV, err := g.SetView("Category", 0, 0, horizOffset, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		categoryV.Frame = true
		categoryV.BgColor = gocui.ColorMagenta //TODO: color here for testing
		categoryV.FgColor = gocui.ColorCyan
		v.categoryV = categoryV

		//TODO: Add go func call to get info here
	}

	return nil
}

// // Setup right side list window
// func (v *GullView) setListView(g *gocui.Gui) error {

// }

func (v *DashboardView) Run() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Fatalf("NewGui: %v", err)
	}
	defer g.Close()
	v.g = g

	defaultSettings(g)

	g.SetManagerFunc(v.layout)
	//TODO: set keybindings w/ v.keybindings(g)
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalf("MainLoop: %v", err)
	}
}

func defaultSettings(g *gocui.Gui) {
	g.InputEsc = true
	g.FgColor = gocui.ColorWhite
	g.BgColor = gocui.ColorBlack
	g.Mouse = true
	g.Highlight = true
}
