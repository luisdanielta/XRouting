from pydantic import BaseModel, Field
from enum import Enum
from typing import List

class ComponentType(str, Enum):
    core = 'core'
    capsule = 'capsule'
    engine = 'engine'
    ship = 'ship'
    launchpad = 'launchpad'

class ComponentStatus(str, Enum):
    active = 'active'
    inactive = 'inactive'
    damaged = 'damaged'
    destroyed = 'destroyed'

class ComponentCategory(str, Enum):
    fixed = 'fixed'
    launched = 'launched'

class Component(BaseModel):
    id: str
    componentId: str = Field(..., description='The ID of the component')
    type: ComponentType = Field(..., description='The type of the component')
    status: ComponentStatus = Field(..., description='The status of the component')
    category: ComponentCategory = Field(..., description='The category of the component')
    missions: List[str] = Field([], description='The missions this component has been a part of')
    hasSubcomponents: bool = Field(False, description='Whether this component has subcomponents')