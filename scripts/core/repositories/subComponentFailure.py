from core.entities.subComponentFailure import SubComponentFailure
from typing import List, Optional


class SubComponentFailureRepository:
    def __init__(self):
        self.failures: List[SubComponentFailure] = []

    def addFailure(self, failure: SubComponentFailure):
        self.failures.append(failure)

    def getFailure(self, id: str) -> Optional[SubComponentFailure]:
        return next((f for f in self.failures if f.id == id), None)

    def getFailures(self) -> List[SubComponentFailure]:
        return self.failures

    def getFailuresBySubComponent(self, subcomponentId: str) -> List[SubComponentFailure]:
        return [f for f in self.failures if f.subcomponentId == subcomponentId]

    def updateFailure(self, id: str, updated_failure: SubComponentFailure):
        for i, f in enumerate(self.failures):
            if f.id == id:
                self.failures[i] = updated_failure
                return

    def deleteFailure(self, id: str):
        self.failures = [f for f in self.failures if f.id != id]
