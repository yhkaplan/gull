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
	categoryView     *gocui.View   // Leftmost view
	categoryList     *categoryList // Left viewmodel
	activityView     *gocui.View   // Rightmost view
	activityList     *ActivityList // Right viewmodel
	selectedView     *gocui.View   // The view currently selected by user
	selectedRowIndex int
}

type categoryList struct {
	title         string
	items         []string
	isHighlighted bool
}

// ActivityList is a generic activity type
type ActivityList struct {
	title         string
	items         []github.GitHubActivity
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

func (l *ActivityList) displayItem(a github.GitHubActivity, v *gocui.View) string {
	item := fmt.Sprintf("%s: %s %s", a.EventType, a.Title, a.Link)
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

func (v *DashboardView) drawListView() error {
	activities := v.activityList.items

	for i := 0; i < len(activities); i++ {
		fmt.Printf("%d", i)
		//l.Clear //TODO: to implement
		a := activities[i]
		_, err := fmt.Fprintln(v.activityView, v.activityList.displayItem(a, v.activityView))
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
func New(a []github.GitHubActivity) *DashboardView {
	activityList := ActivityList{items: a}
	return &DashboardView{
		activityList: &activityList, //TODO: does this work?
	}
}

// Returns window's width and height
func (v *DashboardView) size() (int, int) {
	return v.g.Size()
}

func (v *DashboardView) layout(g *gocui.Gui) error {
	maxX, maxY := v.size()
	horizOffset := maxX / 4

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

		eventTypes := append([]string{"All"}, github.EventTypes...)
		v.categoryList = &categoryList{
			title: name,
			items: eventTypes,
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
		listView.BgColor = gocui.ColorBlack
		listView.FgColor = gocui.ColorGreen

		v.activityView = listView

		if err := v.drawListView(); err != nil {
			return err
		}
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
