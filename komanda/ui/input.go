package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/logger"
)

// InputView creates a new view for input for the cui
func InputView(g *gocui.Gui, x, y, maxX, maxY int) error {

	if v, err := g.SetView("input", x, y, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		_, err := g.SetCurrentView("input")

		if err != nil {
			return err
		}

		logger.Logger.Println(" CHANGE:", "input", x, y, maxX, maxY)

		v.Editor = Editor

		v.FgColor = gocui.Attribute(15 + 1)
		v.BgColor = gocui.Attribute(0)

		v.Autoscroll = false
		v.Editable = true
		v.Wrap = false
		v.Frame = false

	}

	return nil
}
