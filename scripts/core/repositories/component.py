from core.entities.component import Component
from typing import List


class ComponentRepository:
    def __init__(self):
        self.components = []

    def addComponent(self, component: Component):
        self.components.append(component)

    def getComponent(self, id: str) -> Component:
        for component in self.components:
            if component.id == id:
                return component

    def getComponents(self) -> List[Component]:
        return self.components

    def updateComponent(self, id: str, component: Component):
        for i, c in enumerate(self.components):
            if c.id == id:
                self.components[i] = component
                return

    def deleteComponent(self, id: str):
        for i, c in enumerate(self.components):
            if c.id == id:
                del self.components[i]
                return