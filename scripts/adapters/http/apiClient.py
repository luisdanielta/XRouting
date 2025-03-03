
from adapters.http.httpClient import HttpClient, HttpRequestConfig
from adapters.http.apiDtos import RocketDTO, CoreDTO, CapsuleDTO, ShipDTO, LaunchpadDTO, LaunchDTO
from typing import List, Type
from pydantic import BaseModel

class ApiClient(HttpClient):
    _instance = None  # Singleton

    def __init__(self):
        super().__init__("https://api.spacexdata.com/v4")  # Base URL de SpaceX

    @classmethod
    def getInstance(cls):
        if cls._instance is None:
            cls._instance = cls()
        return cls._instance

    def fetchData(self, endpoint: str, response_model: Type[BaseModel]) -> List[BaseModel]:
        response = super().request(endpoint, response_model, HttpRequestConfig(method="GET"))
        if isinstance(response.data, list):
            return response.data  
        return [response_model(**item.dict()) for item in response.data]


    def getRockets(self) -> List[RocketDTO]:
        return self.fetchData("/rockets", RocketDTO)

    def getCores(self) -> List[CoreDTO]:
        return self.fetchData("/cores", CoreDTO)

    def getCapsules(self) -> List[CapsuleDTO]:
        return self.fetchData("/capsules", CapsuleDTO)

    def getShips(self) -> List[ShipDTO]:
        return self.fetchData("/ships", ShipDTO)

    def getLaunchpads(self) -> List[LaunchpadDTO]:
        return self.fetchData("/launchpads", LaunchpadDTO)
    
    def getLaunches(self) -> List[LaunchDTO]:
        return self.fetchData("/launches", LaunchDTO)
