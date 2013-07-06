package greenPlatform

import (
	wi "abstractFactory/widgets"
)

type GreenWidgetFactory struct {
}

func (gwf *GreenWidgetFactory) MakeButton() wi.Button {
	return new(GreenButton)
}

func (gwf *GreenWidgetFactory) MakeWindow() wi.Window {
	return new(GreenWindow)
}

func (gwf *GreenWidgetFactory) MakeWindowCloser() wi.WindowCloser {
	return new(GreenWindowCloser)
}
