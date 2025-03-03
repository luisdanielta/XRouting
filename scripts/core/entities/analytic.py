from pydantic import BaseModel, Field
from datetime import datetime
from typing import List, Dict, Any

class Analytic(BaseModel):
    id: str
    metricId: str = Field(..., description="The unique ID of the metric")
    metricName: str = Field(..., description="The name of the metric (e.g., 'Core Reuse Rate')")
    data: List[Dict[str, float]] = Field([], description="Aggregated data related to the metric")
    updatedAt: datetime = Field(..., description="Last updated timestamp for this metric")
