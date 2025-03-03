import json
import urllib.request
from typing import Type, TypeVar, Dict, Optional, Union, Generic, List
from pydantic import BaseModel, Field

T = TypeVar("T", bound=BaseModel)

class HttpRequestConfig(BaseModel):
    method: str = "GET"
    headers: Dict[str, str] = Field(default_factory=dict)
    body: Optional[Union[BaseModel, Dict[str, str]]] = None


class HttpResponse(BaseModel, Generic[T]):
    data: Union[List[T], T]
    status: int


class HttpClient:
    def __init__(self, base_url: str):
        self.base_url = base_url

    def request(self, endpoint: str, response_model: Type[T], config: HttpRequestConfig = HttpRequestConfig()) -> HttpResponse[T]:
        url = f"{self.base_url}{endpoint}"
        headers = {"Content-Type": "application/json", **config.headers}
        json_data = json.dumps(config.body.model_dump() if isinstance(config.body, BaseModel) else config.body).encode("utf-8") if config.body else None
        req = urllib.request.Request(url, data=json_data, headers=headers, method=config.method)

        with urllib.request.urlopen(req) as response:
            response_data = json.loads(response.read().decode("utf-8"))
            if isinstance(response_data, list):
                parsed_data = [response_model(**item) for item in response_data]
            else:
                parsed_data = response_model(**response_data)
            return HttpResponse[T](data=parsed_data, status=response.getcode())
