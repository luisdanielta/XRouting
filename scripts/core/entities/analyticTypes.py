from typing import TypedDict
from decimal import Decimal

class CostSuccessRateData(TypedDict):
    rocketId: str
    rocketName: str
    costPerLaunch: int
    successRate: Decimal
    engineCount: int
    active: bool
    totalLaunches: int

class LaunchCostDistributionData(TypedDict):
    region: str
    totalLaunchCost: Decimal

class HistoricalSuccessFailureData(TypedDict):
    year: int
    totalLaunches: int
    successfulLaunches: int
    failedLaunches: int
    successRate: Decimal

class MissionsPerComponentTypeData(TypedDict):
    componentType: str
    totalMissions: int

class ComponentStatusDistributionData(TypedDict):
    status: str
    count: int

class RegionSuccessRateData(TypedDict):
    launchpad: str
    region: str
    latitude: Decimal
    longitude: Decimal
    totalLaunches: int
    successfulLaunches: int
    successRate: Decimal

class MaintenanceFrequencyData(TypedDict):
    componentType: str
    maintenanceCount: int
