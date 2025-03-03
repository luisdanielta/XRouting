from aws_lambda_typing.events import APIGatewayProxyEventV2
from aws_lambda_typing.context import Context
from aws_lambda_typing.responses import APIGatewayProxyResponseV2

# utils
from utils.logger import Logger

from adapters.dyto3 import Dyto3

__log = Logger()

def lambda_handler(event: APIGatewayProxyEventV2, context: Context) -> APIGatewayProxyResponseV2:
    db = Dyto3()

if __name__ == '__main__':
    lambda_handler(None, None)
