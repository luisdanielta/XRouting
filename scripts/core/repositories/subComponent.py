from core.entities.subComponent import SubComponent
from typing import List, Optional


class SubComponentRepository:
    def __init__(self):
        self.subComponents: List[SubComponent] = []

    def addSubComponent(self, subcomponent: SubComponent):
        self.subComponents.append(subcomponent)

    def getSubComponent(self, id: str) -> Optional[SubComponent]:
        return next((sc for sc in self.subcomponents if sc.id == id), None)

    def getSubComponents(self) -> List[SubComponent]:
        return self.subComponents

    def getSubComponentsByParent(self, parentComponentId: str) -> List[SubComponent]:
        return [sc for sc in self.subcomponents if sc.parentComponentId == parentComponentId]

    def updateSubComponent(self, id: str, updated_subcomponent: SubComponent):
        for i, sc in enumerate(self.subcomponents):
            if sc.id == id:
                self.subcomponents[i] = updated_subcomponent
                return

    def deleteSubComponent(self, id: str):
        self.subcomponents = [sc for sc in self.subcomponents if sc.id != id]