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
from core.repositories.maintenanceRepository import MaintenanceRepository
from core.repositories.componentRepository import ComponentRepository
from core.repositories.subComponentRepository import SubComponentRepository
from core.repositories.commentsRepository import CommentRepository
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

def transformData(
        spaceXData: Dict[str, List[Union[RocketDTO, LaunchDTO, LaunchpadDTO]]]
    ) -> tuple[
        List[Component], 
        List[Comment],
        List[SubComponent], 
        List[Analytic[Union[Decimal, str]]]
    ]:
    components, comments, subcomponets= extractComponents(spaceXData)
    maintenanceRepo = MaintenanceRepository()
    maintenances: List[Maintenance] = maintenanceRepo.getAll()
    analytics = generateMetrics(spaceXData, components, maintenances)
    return components, comments, subcomponets, analytics

def extractComponents(
        spaceXData: Dict[str, List[Union[RocketDTO, LaunchDTO, LaunchpadDTO]]]
    ) -> tuple[List[Component], List[Comment], List[SubComponent]]:
    components = (
        list(map(lambda rocket: rocketToComponent(rocket, spaceXData["launches"]), spaceXData["rockets"])) +
        list(map(capsuleToComponent, spaceXData["capsules"])) +
        list(map(coreToComponent, spaceXData["cores"])) +
        list(map(launchpadToComponent, spaceXData["launchpads"])) +
        list(map(shipToComponent, spaceXData["ships"]))
    )
    subcomponets = [engineToSubComponent(rocket) for rocket in spaceXData["rockets"]]
    comments = [
        comment
        for launch in spaceXData["launches"]
        if launch.details
        for comment in detailToComments(launch)
    ]

    return components, comments, subcomponets

def generateMetrics(
        spaceXData: Dict[str, List[Union[RocketDTO, LaunchDTO, LaunchpadDTO]]], 
        components: List[Component],
        maintenances: List[Maintenance]
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
        generateRegionSuccessMetric(launches, launchpads),
        generateMaintenanceFrequencyMetric(maintenances, components)
    ]
    
    convertMetricsToDecimal(metrics)
    return metrics

def convertMetricsToDecimal(metrics: List[Analytic[Union[Decimal, str]]]) -> None:
    for metric in metrics:
        metric.convertToDecimal()

def saveToDynamoDB(
        components: List[Component],
        comments: List[Comment], 
        subComponents: List[SubComponent],
        analytics: List[Analytic[float | str]]
    ) -> Dict[str, str | int]:
    try:
        analyticRepo = AnalyticRepository()
        componentRepo = ComponentRepository()
        subComponentRepo = SubComponentRepository()
        commentsRepo = CommentRepository()
        
        componentRepo.saveBatch(components)
        subComponentRepo.saveBatch(subComponents)
        analyticRepo.saveBatch(analytics)
        commentsRepo.saveBatch(comments)

        logger.debug(f"Saved {len(components)} components, {len(subComponents)} sub-components, {len(comments)} comments, and updated {len(analytics)} analytics.")

        return {
            "statusCode": 200,
            "body": json.dumps({
                "message": "Data successfully saved to DynamoDB.",
                "componentsSaved": len(components),
                "subComponentsSaved": len(subComponents),
                "commentsSaved": len(comments),
                "analyticsUpdated": len(analytics)
            }),
        }
    except Exception as e:
        logger.debug(f"Failed to save data: {e}")
        return {"statusCode": 500, "body": json.dumps({"error": str(e)})}

def lambda_handler(event: Dict[str, str], context: object) -> Dict[str, str | int]:
    logger.debug("Fetching SpaceX data...")
    spaceXData = fetchSpaceXData()

    logger.debug("Transforming data...")
    components, comments, subcomponets, analytics = transformData(spaceXData)

    logger.debug("Saving data to DynamoDB...")
    return saveToDynamoDB(components, comments, subcomponets, analytics)

if __name__ == "__main__":
    print("Running locally...")
    response = lambda_handler({}, {})
    print(json.dumps(response, indent=4))
