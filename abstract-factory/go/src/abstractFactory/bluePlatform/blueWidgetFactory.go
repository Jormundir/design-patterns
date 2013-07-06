package bluePlatform

import (
	wi "abstractFactory/widgets"
)

type BlueWidgetFactory struct {
}

// What's important here is the BlueWidgetFactory implements the generic
// WidgetFactory functions, but does so by instantiating the concrete
// Blue* widgets.

func (bwf *BlueWidgetFactory) MakeButton() wi.Button {
	return new(BlueButton)
}

func (bwf *BlueWidgetFactory) MakeWindow() wi.Window {
	return new(BlueWindow)
}

func (bwf *BlueWidgetFactory) MakeWindowCloser() wi.WindowCloser {
	return new(BlueWindowCloser)
}
