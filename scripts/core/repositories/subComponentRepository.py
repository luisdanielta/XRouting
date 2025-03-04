from core.entities.subComponent import SubComponent
from core.repositories.baseRepository import BaseRepository

class SubComponentRepository(BaseRepository[SubComponent]):
    def __init__(self, regionName="us-east-1"):
        super().__init__(tableName="subcomponents", regionName=regionName)

    def createTableIfNotExists(self):
        super().createTableIfNotExists(
            keySchema=[{"AttributeName": "subComponentId", "KeyType": "HASH"}],
            attributeDefinitions=[{"AttributeName": "subComponentId", "AttributeType": "S"}]
        )
