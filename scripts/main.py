import json
from adapters.http.apiClient import ApiClient
from ports.componentServices import (
    rocketToComponent, capsuleToComponent, coreToComponent, 
    shipToComponent, launchpadToComponent
)
from ports.subComponentServices import engineToSubComponent
from ports.commentServices import detailToComments
from ports.maintainceServices import generateMaintenanceFromLaunch
from adapters.repositories.componentRepository import ComponentRepository
from adapters.repositories.subComponentRepository import SubComponentRepository
from adapters.repositories.commentsRepository import CommentRepository
from adapters.repositories.maintenanceRepository import MaintenanceRepository
from utils.logger import Logger

logger = Logger()

def fetchSpaceXData():
    client = ApiClient.getInstance()
    return {
        "rockets": client.getRockets(),
        "capsules": client.getCapsules(),
        "cores": client.getCores(),
        "launches": client.getLaunches(),
        "launchpads": client.getLaunchpads(),
        "ships": client.getShips(),
    }

def transformData(spaceXData):
    components, subComponents, comments, maintenances = [], [], [], []

    for rocket in spaceXData["rockets"]:
        components.append(rocketToComponent(rocket, spaceXData["launches"]))
        subComponents.append(engineToSubComponent(rocket))

    for capsule in spaceXData["capsules"]:
        components.append(capsuleToComponent(capsule))

    for core in spaceXData["cores"]:
        components.append(coreToComponent(core))

    for launchpad in spaceXData["launchpads"]:
        components.append(launchpadToComponent(launchpad))

    for ship in spaceXData["ships"]:
        components.append(shipToComponent(ship))

    for launch in spaceXData["launches"]:
        comments.extend(detailToComments(launch))
        maintenances.extend(generateMaintenanceFromLaunch(launch))

    return components, subComponents, comments, maintenances

def saveToDynamoDB(components, subComponents, comments, maintenances):
    try:
        componentRepo = ComponentRepository()
        subComponentRepo = SubComponentRepository()
        commentsRepo = CommentRepository()
        maintenanceRepo = MaintenanceRepository()
        
        maintenanceRepo.saveBatch(maintenances)
        componentRepo.saveBatch(components)
        subComponentRepo.saveBatch(subComponents)
        commentsRepo.saveBatch(comments)

        logger.debug(f"Saved {len(components)} components, {len(subComponents)} sub-components, {len(comments)} comments, and {len(maintenances)} maintenances.")

        return {
            "statusCode": 200,
            "body": json.dumps({
                "message": "Data successfully saved to DynamoDB.",
                "componentsSaved": len(components),
                "subComponentsSaved": len(subComponents),
                "commentsSaved": len(comments),
                "maintenancesSaved": len(maintenances)
            }),
        }
    except Exception as e:
        logger.debug(f"Failed to save data: {e}")
        return {"statusCode": 500, "body": json.dumps({"error": str(e)})}

def lambdaHandler(event, context):
    logger.debug("Fetching SpaceX data...")
    spaceXData = fetchSpaceXData()

    logger.debug("Transforming data...")
    components, subComponents, comments, maintenances = transformData(spaceXData)

    logger.debug("Saving data to DynamoDB...")
    return saveToDynamoDB(components, subComponents, comments, maintenances)

if __name__ == "__main__":
    print("Running locally...")
    response = lambdaHandler({}, {})
    print(json.dumps(response, indent=4))
