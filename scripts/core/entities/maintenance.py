from pydantic import BaseModel, Field
from enum import Enum
from typing import List, Optional
from datetime import datetime

class FailureRiskLevel(str, Enum):
    low = "low"
    medium = "medium"
    high = "high"

class FailurePrediction(BaseModel):
    riskLevel: FailureRiskLevel = Field(..., description="Risk level of failure (low, medium, high)")
    estimatedFailureDate: Optional[datetime] = Field(None, description="Predicted failure date if available")
    reason: str = Field(..., description="Reason for failure prediction")

class Maintenance(BaseModel):
    id: str
    maintenanceId: str = Field(..., description="The unique ID of the maintenance")
    componentId: str = Field(..., description="The ID of the component affected")
    subcomponentsAffected: List[str] = Field([], description="List of affected subcomponent IDs")
    date: datetime = Field(..., description="Date of maintenance")
    performedBy: str = Field(..., description="User ID who performed the maintenance")
    cost: float = Field(..., description="Cost of the maintenance")
    description: Optional[str] = Field(..., description="Description of the maintenance performed")
    nextMaintenance: Optional[datetime] = Field(None, description="Scheduled next maintenance date")
    failurePrediction: Optional[FailurePrediction] = Field(None, description="Prediction of potential future failures")
