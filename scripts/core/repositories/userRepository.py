from core.entities.user import User
from core.repositories.baseRepository import BaseRepository

class UserRepository(BaseRepository[User]):
    def __init__(self, regionName="us-east-1"):
        super().__init__(tableName="users", regionName=regionName)

    def createTableIfNotExists(self):
        super().createTableIfNotExists(
            keySchema=[{"AttributeName": "id", "KeyType": "HASH"}],
            attributeDefinitions=[{"AttributeName": "id", "AttributeType": "S"}]
        )