package view

import (
	"bytes"
	"fmt"

	"github.com/jroimartin/gocui" //TODO: alias as c instead of gocuit

	"github.com/yhkaplan/gull/github"
)

// DashboardView represents entire dashboard view
type DashboardView struct {
	g                *gocui.Gui
	categoryView     *gocui.View
	categoryList     *categoryList
	activityView     *gocui.View
	selectedView     *gocui.View
	selectedRowIndex int
}

type categoryList struct {
	title         string
	items         []string
	isHighlighted bool
}

// Name always equals title
func (l *categoryList) name() string {
	return l.title
}

func (l *categoryList) Focus(g *gocui.Gui) error {
	l.isHighlighted = true
	_, err := g.SetCurrentView(l.name())

	return err
}

func (l *categoryList) displayItem(i int, v *gocui.View) string {
	item := fmt.Sprint(l.items[i])
	sp := spaces(maxWidth(v) - len(item) - 3)
	return fmt.Sprintf(" %v%v", item, sp)
}

func maxWidth(v *gocui.View) int {
	_, y := v.Size()
	return y
}

func spaces(n int) string {
	var s bytes.Buffer
	for i := 0; i < n; i++ {
		s.WriteString(" ")
	}
	return s.String()
}

func (v *DashboardView) drawCategories() error {
	categories := v.categoryList.items

	for i := 0; i < len(categories); i++ {
		fmt.Printf("%d", i)
		//l.Clear //TODO: to implement
		_, err := fmt.Fprintln(v.categoryView, v.categoryList.displayItem(i, v.categoryView))
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO: will need this for navigation
// func (v *DashboardView) currentIndex() int {
// 	return v.selectedRowIndex
// }

// New initializes the DashboardView
func New() *DashboardView {
	return &DashboardView{}
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
	name := "Category"

	if categoryView, err := g.SetView(name, 0, 0, horizOffset, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		categoryView.Frame = true
		categoryView.BgColor = gocui.ColorBlack
		categoryView.FgColor = gocui.ColorWhite

		v.categoryView = categoryView
		v.selectedView = categoryView // Category view is always initially selected
		v.selectedRowIndex = 0        // Top row always the default

		v.categoryList = &categoryList{
			title: name,
			items: github.EventTypes,
		}
		if err := v.categoryList.Focus(g); err != nil {
			return err
		}
		if err := v.drawCategories(); err != nil {
			return err
		}
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

		v.activityView = listView

		//TODO: Add go func call to get info here
	}

	return nil
}

// Run starts up the cui
func (v *DashboardView) Run() error {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		return fmt.Errorf("NewGui: %v", err)
	}
	defer g.Close()
	v.g = g

	defaultSettings(g)

	g.SetManagerFunc(v.layout)
	//TODO: move keybindings to separate file
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return fmt.Errorf("MainLoop: %v", err)
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func defaultSettings(g *gocui.Gui) {
	g.InputEsc = true
	g.FgColor = gocui.ColorWhite
	g.BgColor = gocui.ColorBlack
	g.Mouse = true
	g.Highlight = true
}
