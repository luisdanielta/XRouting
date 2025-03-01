from aws_lambda_typing.events import APIGatewayProxyEventV2
from aws_lambda_typing.context import Context
from aws_lambda_typing.responses import APIGatewayProxyResponseV2

# utils
from utils.logger import Logger

__log = Logger()

def lambda_handler(event: APIGatewayProxyEventV2, context: Context) -> APIGatewayProxyResponseV2:
    pass