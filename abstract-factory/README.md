# Abstract Factory Pattern

## Purpose

Provide an interface for creating families of related or dependent objects without specifying their concrete classes.

## Example
### Scenario
Consider making an application with a GUI that needs to support multiple platforms, with each platform having its own look and feel to GUI widgets. In order to keep the core of the application simple, it must be isolated from and ignorant of the GUI differences between the platforms. This isolation can be maintained by having the application interact only with abstract, platform-independent widgets. Building an abstract factory structure hides the platform specific implementation details and gives the application core a single place to get abstract widgets from.

### Structure
Each type of widget, a button to close a window for instance, requires an abstract interface, so the application can interact with the interface, not caring about the platform type. Each platform will then have a concrete implementation of the interface, hiding the details of the widget.

A structure to swap out all of the different widgets of one platform for all of the different widgets of another platform requires an abstract interface for creating all of the different widgets in one place - this is the abstract interface for a factory, which is what the abstract factory will interact with, but is not the actual abstract factory the pattern name refers to.

Each platform will have a concrete implementation of the factory, creating all of the concrete widgets of that platform. Now the abstract factory can be made. The abstract factory is a class that implements the factory interface, but does so by holding a reference to a concrete factory and forwarding requests for widgets to the concrete factory, but returning the widgets as the abstract widget class so the application doesn't know the specific platform of the widget it asked for.
