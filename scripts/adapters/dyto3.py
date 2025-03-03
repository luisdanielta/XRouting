from boto3.session import Session
from botocore.exceptions import BotoCoreError, ClientError

class Dyto3(Session):
    def __init__(self, *args, **kwargs):
        super(Dyto3, self).__init__(*args, **kwargs)
        self.resource = self.resource('dynamodb')