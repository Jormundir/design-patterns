package greenPlatform

import (
	af "abstractFactory/widgets"
)

type GreenWindowCloser struct {
}

func (gwc *GreenWindowCloser) CloseWindow(w af.Window) string {
	w.Close()
	return "You closed a window with a green window closer"
}
