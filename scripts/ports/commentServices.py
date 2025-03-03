import random
import uuid
from ports.componentServices import getCurrentTimestamp
from typing import List
from core.entities.comment import Comment
from adapters.http.apiDtos import LaunchDTO

def detailToComments(launch: LaunchDTO, userId="admin") -> List[Comment]:
    details = launch.details.lower() if launch.details else ""

    keywords = {
        "rocket": [launch.rocket] if launch.rocket else [],
        "launchpad": [launch.launchpad] if launch.launchpad else [],
        "core": [core["id"] for core in launch.cores if "id" in core],
        "capsule": launch.capsules or [],
        "ship": launch.ships or [],
    }

    comments = [
        Comment(
            id=str(uuid.uuid4()),
            componentId=random.choice(values),
            userId=userId,
            message=launch.details,
            updatedDate=getCurrentTimestamp()
        )
        for key, values in keywords.items() if key in details and values
    ]

    return comments
