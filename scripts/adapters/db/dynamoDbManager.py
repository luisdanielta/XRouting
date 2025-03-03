import boto3

class DynamoDBManager:
    def __init__(self, region_name="us-east-1"):
        self.client = boto3.client("dynamodb", region_name=region_name)

    def tableExists(self, tableName):
        exitingTables = self.client.list_tables()["TableNames"]
        return tableName in exitingTables

    def createTable(self, tableName, keySchema, attributeDefinitions, throughput=None):
        if self.tableExists(tableName):
            print(f"Table '{tableName}' already exists.")
            return
        print(f"Creating Table'{tableName}' on DynamoDB...")
        throughput = throughput or {"ReadCapacityUnits": 5, "WriteCapacityUnits": 5}

        self.client.create_table(
            TableName=tableName,
            KeySchema=keySchema,
            AttributeDefinitions=attributeDefinitions,
            ProvisionedThroughput=throughput,
        )
        self.client.get_waiter("table_exists").wait(TableName=tableName)
        print(f"Sucessfully created'{tableName}'")

    def deleteTable(self, tableName):
        if not self.tableExists(tableName):
            print(f"Table '{tableName}' not in DB")
            return
        self.client.deleteTable(TableName=tableName)
        self.client.get_waiter("table_not_exists").wait(TableName=tableName)
        print(f"Table '{tableName}' sucessfully deleted")
