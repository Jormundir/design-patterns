package bluePlatform

import (
	af "abstractFactory/widgets"
)

type BlueWindowCloser struct {
}

// This type represents a blue window closer. It takes, however, a generic window type,
// so it is capable of closing any type of window, not just a blue window. This is because
// the program will pass it a generic window; the program doesn't know there is a more
// specific type.
func (bwc *BlueWindowCloser) CloseWindow(w af.Window) string {
	w.Close()
	return "You closed a window with a blue window closer"
}
