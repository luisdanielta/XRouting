from core.entities.maintenance import Maintenance
from adapters.repositories.baseRepository import BaseRepository

class MaintenanceRepository(BaseRepository[Maintenance]):
    def __init__(self, regionName="us-east-1"):
        super().__init__(tableName="maintenances", regionName=regionName)

    def createTableIfNotExists(self):
        super().createTableIfNotExists(
            keySchema=[{"AttributeName": "maintenanceId", "KeyType": "HASH"}],
            attributeDefinitions=[{"AttributeName": "maintenanceId", "AttributeType": "S"}]
        )
