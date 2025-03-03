from typing import List, TypeVar, Generic
from adapters.db.dynamoDbManager import DynamoDBManager
from adapters.db.dbConnection import getDynamodb
from pydantic import BaseModel

T = TypeVar("T", bound=BaseModel)

class BaseRepository(Generic[T]):
    def __init__(self, tableName: str, regionName="us-east-1"):
        self.dynamodb = getDynamodb()
        self.tableName = tableName
        self.table = self.dynamodb.Table(tableName)
        self.dbManager = DynamoDBManager(regionName)

    def createTableIfNotExists(self, keySchema, attributeDefinitions):
        self.dbManager.createTable(
            tableName=self.tableName,
            keySchema=keySchema,
            attributeDefinitions=attributeDefinitions
        )

    def save(self, item: T):
        self.table.put_item(Item=item.model_dump())

    def saveBatch(self, items: List[T]):
        with self.table.batch_writer() as batch:
            for item in items:
                batch.put_item(Item=item.model_dump())

    def getById(self, itemId: str, keyName: str):
        response = self.table.get_item(Key={keyName: itemId})
        return response.get("Item")

    def getAll(self):
        response = self.table.scan()
        return response.get("Items", [])
