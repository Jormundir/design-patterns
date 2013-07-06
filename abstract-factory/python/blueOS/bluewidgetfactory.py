from bluewidgets import *

class BlueWidgetFactory:

	def makeButton(self):
		return BlueButton()

	def makeWindow(self):
		return BlueWindow()

	def makeWindowCloser(self):
		return BlueWindowCloser()
