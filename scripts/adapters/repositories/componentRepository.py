from core.entities.component import Component
from adapters.repositories.baseRepository import BaseRepository

class ComponentRepository(BaseRepository[Component]):
    def __init__(self, regionName="us-east-1"):
        super().__init__(tableName="components", regionName=regionName)

    def createTableIfNotExists(self):
        super().createTableIfNotExists(
            keySchema=[{"AttributeName": "componentId", "KeyType": "HASH"}],
            attributeDefinitions=[{"AttributeName": "componentId", "AttributeType": "S"}]
        )
