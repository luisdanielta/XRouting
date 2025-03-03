import uuid
from datetime import datetime, timezone
from core.entities.component import Component, ComponentType, ComponentStatus, ComponentCategory
from adapters.http.apiDtos import *
from typing import List

def getCurrentTimestamp():
    return datetime.now(timezone.utc).isoformat()

def rocketToComponent(rocket: RocketDTO, launches: List[LaunchDTO]) -> Component:
    rocketLaunched = any(launch.rocket == rocket.id for launch in launches)    
    missions = [launch.id for launch in launches if launch.rocket == rocket.id]
    return Component(
        id=str(uuid.uuid4()),
        componentId=rocket.id,
        type=ComponentType.rocket,
        status=ComponentStatus.active if rocket.active else ComponentStatus.inactive,
        category=ComponentCategory.launched if rocketLaunched else ComponentCategory.fixed,
        missions=missions,
        hasSubcomponents=rocket.engines is not None,
        created_at=getCurrentTimestamp(),
        last_updated=getCurrentTimestamp()
    )

def mapCapsuleStatus(status: CapsuleStatusDTO) -> ComponentStatus:
    statusMapping = {
        CapsuleStatusDTO.active: ComponentStatus.active,
        CapsuleStatusDTO.retired: ComponentStatus.inactive,
        CapsuleStatusDTO.unknown: ComponentStatus.damaged,
        CapsuleStatusDTO.destroyed: ComponentStatus.destroyed
    }
    return statusMapping.get(status, ComponentStatus.inactive)

def capsuleToComponent(capsule: CapsuleDTO) -> Component:
    return Component(
        id=str(uuid.uuid4()),
        componentId=capsule.id,
        type=ComponentType.capsule,
        status=mapCapsuleStatus(capsule.status),
        category=ComponentCategory.launched if bool(capsule.launches) else ComponentCategory.fixed,
        missions=capsule.launches,
        hasSubcomponents=False,
        createdAt=getCurrentTimestamp(),
        lastUpdated=getCurrentTimestamp()
    )

def mapCoreStatus(status: CoreStatusDTO) -> ComponentStatus:
    statusMapping = {
        CoreStatusDTO.active: ComponentStatus.active,
        CoreStatusDTO.lost: ComponentStatus.destroyed,
        CoreStatusDTO.expended: ComponentStatus.damaged,
        CoreStatusDTO.inactive: ComponentStatus.inactive
    }
    return statusMapping.get(status, ComponentStatus.inactive)

def coreToComponent(core: CoreDTO) -> Component:
    return Component(
        id=str(uuid.uuid4()),
        componentId=core.id,
        type=ComponentType.core,
        status=mapCoreStatus(core.status),
        category=ComponentCategory.launched if bool(core.launches) else ComponentCategory.fixed,
        missions=core.launches,
        hasSubcomponents=False,
        createdAt=getCurrentTimestamp(),
        lastUpdated=getCurrentTimestamp()
    )

def shipToComponent(ship: ShipDTO) -> Component:
    return Component(
        id=str(uuid.uuid4()),
        componentId=ship.id,
        type=ComponentType.ship,
        status=ComponentStatus.active if ship.active else ComponentStatus.inactive,
        category=ComponentCategory.launched if bool(ship.launches) else ComponentCategory.fixed,
        missions=ship.launches,
        hasSubcomponents=False,
        createdAt=getCurrentTimestamp(),
        lastUpdated=getCurrentTimestamp()
    )

def launchpadToComponent(launchpad: LaunchpadDTO) -> Component:
    """Convierte un LaunchpadDTO en un Component con camelCase."""
    return Component(
        id=str(uuid.uuid4()),
        componentId=launchpad.id,
        type=ComponentType.launchpad,
        status=ComponentStatus.active if launchpad.status == "active" else ComponentStatus.inactive,
        category=ComponentCategory.launched if bool(launchpad.launches) else ComponentCategory.fixed,
        missions=launchpad.launches,
        hasSubcomponents=False,
        createdAt=getCurrentTimestamp(),
        lastUpdated=getCurrentTimestamp()
    )

