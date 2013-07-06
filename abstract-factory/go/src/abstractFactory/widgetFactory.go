package abstractFactory

import (
	wi "abstractFactory/widgets"
)

type WidgetFactory interface {
	MakeButton() wi.Button
	MakeWindow() wi.Window
	MakeWindowCloser() wi.WindowCloser
}
