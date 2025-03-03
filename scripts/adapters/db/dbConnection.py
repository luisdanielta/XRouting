import boto3

def getDynamodb():
    return boto3.resource("dynamodb", region_name="us-east-1")
