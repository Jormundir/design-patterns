from greenwidgets import *

class GreenWidgetFactory:

	def makeButton(self):
		return GreenButton()

	def makeWindow(self):
		return GreenWindow()

	def makeWindowCloser(self):
		return GreenWindowCloser()
