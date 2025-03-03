from enum import Enum
from pydantic import BaseModel, Field

class SubComponentStatus(str, Enum):
    good = 'good'
    needReview = 'needReview'
    damaged = 'damaged'
    destroyed = 'destroyed'

class SubComponent(BaseModel):
    id: str
    subComponentId: str = Field(..., description='The ID of the subcomponent')
    parentComponentId:  str = Field(..., description='The parent component ID of this subcomponent')
    name: str = Field(..., description='The name of the subcomponent')
    status: SubComponentStatus = Field(..., description='The status of the subcomponent')