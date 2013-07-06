package abstractFactory

// Note this struct has "Abstract" in its name, but this is actually
// a concrete class that will be instantiated and used. The "Abstract"
// is referring to the fact that this class will only return abstract, generic
// widgets without caring about the concrete widget types. I.e. it returns Button,
// not BlueButton
type AbstractWidgetFactory struct {
	WidgetFactory
}
