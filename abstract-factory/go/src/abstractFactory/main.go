package abstractFactory

import (
	//	bp "abstractFactory/bluePlatform"
	gp "abstractFactory/greenPlatform"
	"fmt"
)

func Main() {
	// Initialization
	// When your program is just starting up, or just getting ready to
	// make some widgets, you specify the concrete type of widgets you
	// want to make - perhaps decided by checking that you're on
	// "Blue OS" and need to make blue buttons.

	// the config file created by installation says we're on "Green OS", so we'll
	// need green widgets. We'll instantiate our GreenWidgetFactory in this one
	// place, then we'll hide it in an AbstractWidgetFactory, and the rest of
	// our application will ask the AbstractWidgetFactory for widgets, not
	// caring what concrete type they're getting.
	greenWidgetFactory := new(gp.GreenWidgetFactory)

	// Hide the concrete factory in an AbstractWidgetFactory.
	widgetFactory := &AbstractWidgetFactory{greenWidgetFactory}

	// Now we're done with initialization, so it's time to run the
	// heart of the program, passing it the AbstractWidgetFactory.
	heartOfProgram(widgetFactory)
}

func heartOfProgram(wf *AbstractWidgetFactory) {
	// Now the heart of our program has an Abstract widget factory, and has
	// no idea what concrete type of widgets it will make. Perfect! Now we can
	// program against the widget interfaces, and if we change from "Green OS"
	// to "Blue OS" the heart of our program won't have to change at all.

	button := wf.MakeButton()
	fmt.Println(button.Click())
	windowCloser := wf.MakeWindowCloser()
	window := wf.MakeWindow()
	fmt.Println(windowCloser.CloseWindow(window))
}
