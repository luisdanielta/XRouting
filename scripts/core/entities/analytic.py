from pydantic import BaseModel, Field
from datetime import datetime
from typing import List, Dict, TypeVar, Generic
from decimal import Decimal

T = TypeVar("T")

class Analytic(BaseModel, Generic[T]):
    id: str
    metricName: str = Field(..., description="The name of the metric (e.g., 'Core Reuse Rate')")
    description: str = Field(..., description="Detailed explanation of the metric")
    data: List[Dict[str, T]] = Field([], description="Aggregated data related to the metric")
    updatedAt: str = Field(..., description="Last updated timestamp for this metric")

    def convertToDecimal(self) -> None:
        """
        Converts all float values in 'data' to Decimal for DynamoDB compatibility.
        """
        for item in self.data:
            for key, value in item.items():
                if isinstance(value, float):
                    item[key] = Decimal(str(value))
