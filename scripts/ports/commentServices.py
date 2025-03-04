import random
import uuid
from ports.componentServices import getCurrentTimestamp
from typing import List, Optional
from core.entities.comment import Comment
from core.entities.user import UserRole
from adapters.http.apiDtos import LaunchDTO
from core.repositories.userRepository import UserRepository

def getAdminUserId() -> str:
    """
    Fetches a user with the 'admin' role from the UserRepository.

    Returns:
        str: The ID of a randomly selected user with the 'admin' role, or "admin" if no such user exists.
    """
    """Fetches a user with 'admin' role from the UserRepository."""
    userRepo = UserRepository()
    users = userRepo.getAll() 
    adminUsers = [user for user in users if user["role"] == UserRole.admin.value]
    return random.choice(adminUsers)["id"] if adminUsers else "admin"

def detailToComments(launch: LaunchDTO) -> List[Comment]:
    """Generates comments related to a launch, assigning a dynamic admin user."""
    details = launch.details.lower() if launch.details else ""
    userId = getAdminUserId()
    keywords = {
        "rocket": [launch.rocket] if launch.rocket else [],
        "launchpad": [launch.launchpad] if launch.launchpad else [],
        "core": [core["type"] for core in launch.cores if "type" in core],
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
