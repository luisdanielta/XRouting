from core.entities.comment import Comment
from core.repositories.baseRepository import BaseRepository

class CommentRepository(BaseRepository[Comment]):
    def __init__(self, regionName="us-east-1"):
        super().__init__(tableName="comments", regionName=regionName)

    def createTableIfNotExists(self):
        super().createTableIfNotExists(
            keySchema=[{"AttributeName": "id", "KeyType": "HASH"}],
            attributeDefinitions=[{"AttributeName": "id", "AttributeType": "S"}]
        )
