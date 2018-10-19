package view

import "testing"
import "github.com/jroimartin/gocui"

var l = categoryList{}
var g = gocui.Gui{}

func TestName(t *testing.T) {
	resetValues()
	title := "test phrase"
	l.title = title

	result := l.name()

	if result != title {
		t.Errorf("l.name() is supposed to be title, but was instead: %s", result)
	}
}

func TestFocus(t *testing.T) {
	resetValues()

	err := l.Focus(&g)

	if !l.isHighlighted {
		t.Error("categoryList.isHighlighted should be set to false, but it wasn't")
	}

	if err == nil {
		t.Error("Err is not supposed to be nil, but it was")
	}
}

//TODO: use mock for test focus above
type mockGoCui struct{}

// Helpers

func resetValues() {
	l = categoryList{}

	g = gocui.Gui{}
}
