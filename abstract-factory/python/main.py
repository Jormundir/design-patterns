from abstractwidgetfactory import AbstractWidgetFactory
from blueOS.bluewidgetfactory import BlueWidgetFactory
from greenOS.greenwidgetfactory import GreenWidgetFactory


# installation script made config file saying the OS is blue
platform = "blueOS"


# later when starting the main program:
if platform == "blueOS":
	widgetFactory = AbstractWidgetFactory(BlueWidgetFactory())
else:
	widgetFactory = AbstractWidgetFactory(GreenWidgetFactory())

# In the heart of our program, we want to make and interact with some widgets:

button = widgetFactory.makeButton()
window = widgetFactory.makeWindow()
windowCloser = widgetFactory.makeWindowCloser()

print windowCloser.closeWindow(window)
print button.click()