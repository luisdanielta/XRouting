from pydantic import BaseModel, Field
from enum import Enum
from typing import Optional, List, Dict


class CapsuleStatusDTO(str, Enum):
    active = 'active'
    retired = 'retired'
    unknown = 'unknown'
    destroyed = 'destroyed'

class CoreStatusDTO(str, Enum):
    active = 'active'
    lost = 'lost'
    expended = 'expended'
    inactive = 'inactive'

class EngineDTO(BaseModel):
    number: int
    type: str
    version: str
    thrust_to_weight: float
    
class RocketDTO(BaseModel):
    id: str
    name: str
    active: bool
    description: str
    cost_per_launch: float
    success_rate_pct: float
    engines: EngineDTO

class CoreDTO(BaseModel):
    id: str
    core_serial: str = Field(None)
    status: Optional[CoreStatusDTO]
    reuse_count: int
    launches: List[str]

class CapsuleDTO(BaseModel):
    id: str
    serial: str
    status: Optional[CapsuleStatusDTO]
    reuse_count: int
    last_update: Optional[str]
    launches: List[str]
    water_landings: int

class ShipDTO(BaseModel):
    id: str
    name: str
    launches: List[str]
    active: bool
    home_port: Optional[str]

class LaunchpadDTO(BaseModel):
    id: str
    name: str
    status: str
    region: str
    latitude: float
    longitude: float
    launches: List[str]

class LaunchDTO(BaseModel):
    id: str
    launchpad: str
    rocket: str
    cores: List[Dict]
    capsules: List[str]
    ships: List[str]
    success: Optional[bool]
    details: Optional[str]
    date_utc: Optional[str]

