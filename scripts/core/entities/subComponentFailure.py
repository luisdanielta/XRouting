from core.entities.maintenance import FailureRiskLevel
from pydantic import BaseModel, Field
from enum import Enum
from datetime import datetime
from typing import Optional

class FailureType(str, Enum):
    wear = "wear"
    overheating = "overheating"
    vibrations = "vibrations"

class FailureStatus(str, Enum):
    pending = "pending"
    underRepair = "underRepair"
    resolved = "resolved"

class SubComponentFailure(BaseModel):
    id: str
    subComponentId: str = Field(..., description="The ID of the affected subComponent")
    detectedAt: datetime = Field(..., description="Timestamp when the failure was detected")
    failureType: FailureType = Field(..., description="The type of failure (wear, overheating, vibrations)")
    severity: FailureRiskLevel = Field(..., description="The severity of the failure (low, medium, high)")
    reportedBy: str = Field(..., description="The ID of the user who reported the failure")
    status: FailureStatus = Field(..., description="Current status of the failure (pending, under repair, resolved)")
    maintenanceId: Optional[str] = Field(None, description="Optional: ID of the maintenance that fixed the failure")
