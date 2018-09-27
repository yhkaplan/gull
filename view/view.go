package view

import (
	"log"

	"github.com/jroimartin/gocui"
)

type DashboardView struct {
	g            *gocui.Gui
	categoryView *gocui.View
	listView     *gocui.View
}

// Initializes DashboardView
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
	err = v.setListView(g, horizOffset, maxX, maxY)
	if err != nil {
		return err
	}

	return nil
}

// Setup left side category window
func (v *DashboardView) setCategoryView(g *gocui.Gui, horizOffset int, maxY int) error {
	if categoryView, err := g.SetView("Category", 0, 0, horizOffset, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		categoryView.Frame = true
		categoryView.BgColor = gocui.ColorMagenta //TODO: color here for testing
		categoryView.FgColor = gocui.ColorCyan

		v.categoryView = categoryView

		//TODO: Add go func call to get info here
	}

	return nil
}

// Setup right side list window
func (v *DashboardView) setListView(g *gocui.Gui, horizOffset int, maxX int, maxY int) error {
	//TODO: Set view name to change dynamically to PR, ALL, etc
	if listView, err := g.SetView("Events", horizOffset, 0, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		listView.Frame = true
		listView.BgColor = gocui.ColorGreen
		listView.FgColor = gocui.ColorYellow

		v.listView = listView

		//TODO: Add go func call to get info here
	}

	return nil
}

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
