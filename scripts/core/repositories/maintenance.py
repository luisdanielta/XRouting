from core.entities.maintenance import Maintenance
from typing import List, Optional


class MaintenanceRepository:
    def __init__(self):
        self.maintenances: List[Maintenance] = []

    def addMaintenance(self, maintenance: Maintenance):
        self.maintenances.append(maintenance)

    def getMaintenance(self, id: str) -> Optional[Maintenance]:
        return next((m for m in self.maintenances if m.id == id), None)

    def getMaintenances(self) -> List[Maintenance]:
        return self.maintenances

    def getMaintenancesByComponent(self, componentId: str) -> List[Maintenance]:
        return [m for m in self.maintenances if m.componentId == componentId]

    def updateMaintenance(self, id: str, updated_maintenance: Maintenance):
        for i, m in enumerate(self.maintenances):
            if m.id == id:
                self.maintenances[i] = updated_maintenance
                return

    def deleteMaintenance(self, id: str):
        self.maintenances = [m for m in self.maintenances if m.id != id]
