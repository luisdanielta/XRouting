import json
from utils.logger import Logger
from decimal import Decimal
from typing import Dict, List, Union
from ports.analyticServices import *
from ports.componentServices import *
from ports.subComponentServices import engineToSubComponent
from ports.commentServices import detailToComments
from ports.maintainceServices import generateMaintenanceFromLaunch
from core.entities.analytic import Analytic
from core.entities.component import Component
from core.entities.comment import Comment
from core.entities.maintenance import Maintenance
from core.entities.subComponent import SubComponent
from adapters.http.apiClient import ApiClient
from core.repositories.analyticRepository import AnalyticRepository
from adapters.http.apiDtos import RocketDTO, LaunchDTO, LaunchpadDTO, CoreDTO, CapsuleDTO, ShipDTO

logger = Logger()

def fetchSpaceXData() -> Dict[str, List[RocketDTO | LaunchDTO | LaunchpadDTO | CoreDTO | CapsuleDTO | ShipDTO]]:
    client = ApiClient.getInstance()
    return {
        "rockets": client.getRockets(),
        "launches": client.getLaunches(),
        "launchpads": client.getLaunchpads(),
        "cores": client.getCores(),
        "capsules": client.getCapsules(),
        "ships": client.getShips()
    }

def transformData(spaceXData: Dict[str, List[Union[RocketDTO, LaunchDTO, LaunchpadDTO]]]) -> List[Analytic[Union[Decimal, str]]]:
    components, comments, maintenances, subcomponets= extractComponents(spaceXData)
    analytics = generateMetrics(spaceXData, components)
    return analytics

def extractComponents(
        spaceXData: Dict[str, List[Union[RocketDTO, LaunchDTO, LaunchpadDTO]]]
    ) -> tuple[List[Component], List[Comment], List[Maintenance], List[SubComponent]]:
    components = (
        list(map(lambda rocket: rocketToComponent(rocket, spaceXData["launches"]), spaceXData["rockets"])) +
        list(map(capsuleToComponent, spaceXData["capsules"])) +
        list(map(coreToComponent, spaceXData["cores"])) +
        list(map(launchpadToComponent, spaceXData["launchpads"])) +
        list(map(shipToComponent, spaceXData["ships"]))
    )
    subcomponets = [engineToSubComponent(rocket) for rocket in spaceXData["rockets"]]
    comments =[detailToComments(launch) for launch in spaceXData["launches"]]

    return components, comments, subcomponets

def generateMetrics(
        spaceXData: Dict[str, List[Union[RocketDTO, LaunchDTO, LaunchpadDTO]]], 
        components: List[Component]
    ) -> List[Analytic[Union[Decimal, str]]]:
    rockets: List[RocketDTO] = spaceXData["rockets"]
    launches: List[LaunchDTO] = spaceXData["launches"]
    launchpads: List[LaunchpadDTO] = spaceXData["launchpads"]

    metrics = [
        generateCostSuccessMetric(rockets, launches),
        generateLaunchCostDistributionMetric(rockets, launches, launchpads),
        generateHistoricalSuccessFailureMetric(launches),
        generateMissionsPerComponentTypeMetric(components),
        generateComponentStatusDistributionMetric(components),
        generateRegionSuccessMetric(launches, launchpads)
    ]
    
    convertMetricsToDecimal(metrics)
    return metrics

def convertMetricsToDecimal(metrics: List[Analytic[Union[Decimal, str]]]) -> None:
    for metric in metrics:
        metric.convertToDecimal()

def saveToDynamoDB(analytics: List[Analytic[float | str]]) -> Dict[str, str | int]:
    try:
        analyticRepo = AnalyticRepository()
        [analyticRepo.save(analytic) for analytic in analytics]
        return {"statusCode": 200, "body": json.dumps({"message": "Data saved successfully!"})}
    except Exception as e:
        logger.debug(f"Failed to save data: {e}")
        return {"statusCode": 500, "body": json.dumps({"error": str(e)})}

def lambdaHandler(event: Dict[str, str], context: object) -> Dict[str, str | int]:
    logger.debug("Fetching SpaceX data...")
    spaceXData = fetchSpaceXData()

    logger.debug("Transforming data...")
    analytics = transformData(spaceXData)

    logger.debug("Saving data to DynamoDB...")
    return saveToDynamoDB(analytics)

if __name__ == "__main__":
    print("Running locally...")
    response = lambdaHandler({}, {})
    print(json.dumps(response, indent=4))
