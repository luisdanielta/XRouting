import uuid
import random
from datetime import datetime
from ports.componentServices import getCurrentTimestamp
from adapters.http.apiDtos import LaunchDTO
from core.entities.maintenance import Maintenance, FailurePrediction, FailureRiskLevel
from ports.commentServices import getAdminUserId
from typing import List

def generateMaintenanceFromLaunch(launch: LaunchDTO) -> List[Maintenance]:
    adminUserId = getAdminUserId()
    dateNowStr = getCurrentTimestamp()  
    dateNow = datetime.fromisoformat(dateNowStr)
    
    maintenances = []
    
    if launch.success:
        for capsule in launch.capsules:
            maintenances.append(Maintenance(
                id=str(uuid.uuid4()),
                maintenanceId=str(uuid.uuid4()),
                componentId=capsule,
                subcomponentsAffected=[],
                date=dateNowStr,
                performedBy=adminUserId,
                cost=int(random.uniform(200000, 500000)),
                description=launch.details or "Routine maintenance",
                nextMaintenance=(dateNow.replace(year=dateNow.year + 1)).isoformat(),
                failurePrediction=FailurePrediction(
                    riskLevel=FailureRiskLevel.low,
                    reason="Routine maintenance after successful launch"
                )
            ))
        
        subcomponents = [core["id"] for core in launch.cores if "id" in core]
        maintenances.append(Maintenance(
            id=str(uuid.uuid4()),
            maintenanceId=str(uuid.uuid4()),
            componentId=launch.rocket,
            subcomponentsAffected=subcomponents,
            date=dateNowStr,
            performedBy=adminUserId,
            cost=int(random.uniform(50000, 200000)),
            description=launch.details or "Routine maintenance",
            nextMaintenance=(dateNow.replace(year=dateNow.year + 1)).isoformat(),
            failurePrediction=FailurePrediction(
                riskLevel=FailureRiskLevel.low,
                reason="Routine maintenance after successful launch"
            )
        ))
    else:
        subcomponents = [core["id"] for core in launch.cores if "id" in core]
        maintenances.append(Maintenance(
            id=str(uuid.uuid4()),
            maintenanceId=str(uuid.uuid4()),
            componentId=launch.rocket,
            subcomponentsAffected=subcomponents,
            date=dateNowStr,
            performedBy=adminUserId,
            cost=int(random.uniform(500000, 1000000)),
            description=launch.details or "Emergency repair due to failed launch",
            nextMaintenance=(dateNow.replace(year=dateNow.year + 1)).isoformat(),
            failurePrediction=FailurePrediction(
                riskLevel=FailureRiskLevel.high,
                reason="Emergency repair after failed launch"
            )
        ))
    
    return maintenances
