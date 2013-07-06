class AbstractWidgetFactory:

	def __init__(self, factory):
		self.factory = factory

	def makeButton(self):
		return self.factory.makeButton()

	def makeWindow(self):
		return self.factory.makeWindow()

	def makeWindowCloser(self):
		return self.factory.makeWindowCloser()
