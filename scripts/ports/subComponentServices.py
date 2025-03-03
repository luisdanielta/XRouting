import uuid
from core.entities.subComponent import SubComponent, SubComponentStatus
from adapters.http.apiDtos import RocketDTO
from ports.componentServices import getCurrentTimestamp

def engineToSubComponent(rocket: RocketDTO) -> SubComponent:
    return SubComponent(
        id=str(uuid.uuid4()),
        subComponentId=str(rocket.engines.number),
        parentComponentId=rocket.id,
        name=f"{rocket.engines.type} {rocket.engines.version}",
        status=SubComponentStatus.good,
        created_at=getCurrentTimestamp(),
        last_updated=getCurrentTimestamp()
    )