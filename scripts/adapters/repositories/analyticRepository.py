from core.entities.analytic import Analytic
from adapters.repositories.baseRepository import BaseRepository

class AnalyticRepository(BaseRepository[Analytic]):
    def __init__(self, regionName="us-east-1"):
        super().__init__(tableName="analytics", regionName=regionName)

    def createTableIfNotExists(self):
        super().createTableIfNotExists(
            keySchema=[{"AttributeName": "analyticId", "KeyType": "HASH"}],
            attributeDefinitions=[{"AttributeName": "analyticId", "AttributeType": "S"}]
        )
