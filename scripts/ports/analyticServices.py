from collections import defaultdict
from decimal import Decimal
from typing import List, Dict, Union
from adapters.http.apiDtos import RocketDTO, LaunchDTO, LaunchpadDTO
from core.entities.maintenance import Maintenance
from core.entities.component import Component, ComponentStatus, ComponentType
from core.entities.analytic import Analytic
from ports.componentServices import getCurrentTimestamp
from datetime import datetime

def generateCostSuccessMetric(rockets: List[RocketDTO], launches: List[LaunchDTO]) -> Analytic:
    """
    Generates the cost vs. success rate metric for rockets.

    Args:
        rockets (List[RocketDTO]): List of rockets from the API.
        launches (List[LaunchDTO]): List of launches from the API.

    Returns:
        Analytic: Metric data ready to be stored in the Analytic table.
    """
    rocketLaunchStats = defaultdict(lambda: {"total": 0, "success": 0})

    for launch in launches:
        rocketId = launch.rocket
        if rocketId:
            rocketLaunchStats[rocketId]["total"] += 1
            if launch.success:
                rocketLaunchStats[rocketId]["success"] += 1

    data = []

    for rocket in rockets:
        rocketId = rocket.id
        launchStats = rocketLaunchStats.get(rocketId, {"total": 0, "success": 0})

        totalLaunches = launchStats["total"]
        successCount = launchStats["success"]
        successRate = (successCount / totalLaunches * 100) if totalLaunches > 0 else 0

        data.append({
            "rocketId": rocket.id,
            "rocketName": rocket.name,
            "costPerLaunch": rocket.cost_per_launch,
            "successRate": successRate,
            "engineCount": rocket.engines.number,
            "active": rocket.active,
            "totalLaunches": totalLaunches
        })

    analytic = Analytic(
        id="costSuccessRate",
        metricName="Cost vs. Success Rate of Rockets",
        description="This metric compares the cost per launch of different rockets "
                    "against their actual success rates, helping to evaluate the "
                    "economic efficiency of each rocket model.",
        data=data,
        updatedAt=getCurrentTimestamp()
    )
    return analytic


def generateLaunchCostDistributionMetric(
        rockets: List[RocketDTO], 
        launches: List[LaunchDTO], 
        launchpads: List[LaunchpadDTO]
    ) -> Analytic[Decimal | str]:
        """
        Generates a metric for launch cost distribution by region.

        Args:
            rockets (List[RocketDTO]): List of rockets.
            launches (List[LaunchDTO]): List of launches.
            launchpads (List[LaunchpadDTO]): List of launchpads.

        Returns:
            Analytic[Decimal | str]: Launch cost distribution metric.
        """

        rocketCostMap: Dict[str, Decimal] = {
            rocket.id: Decimal(str(rocket.cost_per_launch)) for rocket in rockets
        }
        launchpadRegionMap: Dict[str, str] = {
            launchpad.id: launchpad.region for launchpad in launchpads
        }

        regionCostMap: Dict[str, Decimal] = defaultdict(Decimal)

        for launch in launches:
            rocketId = launch.rocket
            launchpadId = launch.launchpad

            if rocketId in rocketCostMap and launchpadId in launchpadRegionMap:
                region = launchpadRegionMap[launchpadId]
                regionCostMap[region] += rocketCostMap[rocketId]

        data = [{"region": region, "totalLaunchCost": cost} for region, cost in regionCostMap.items()]
        analytic = Analytic[Decimal | str](
            id="launchCostDistribution",
            metricName="Launch Cost Distribution by Region",
            description="This metric analyzes the total launch costs in different geographic regions, "
                        "helping to understand the financial distribution of launch activities across "
                        "various spaceports worldwide.",
            data=data,
            updatedAt=getCurrentTimestamp()
        )

        analytic.convertToDecimal()

        return analytic


def generateHistoricalSuccessFailureMetric(launches: List[LaunchDTO]) -> Analytic:
    """
    Generates a metric showing historical success and failure rates of launches per year.

    Args:
        launches (List[LaunchDTO]): List of launches.

    Returns:
        Analytic: Metric containing success and failure rates per year.
    """

    yearlyLaunchStats: Dict[int, Dict[str, int]] = defaultdict(lambda: {"total": 0, "success": 0})

    for launch in launches:
        launchYear = datetime.fromisoformat(launch.date_utc).year
        yearlyLaunchStats[launchYear]["total"] += 1
        if launch.success:
            yearlyLaunchStats[launchYear]["success"] += 1

    data = []
    for year, stats in sorted(yearlyLaunchStats.items()):
        totalLaunches = stats["total"]
        successCount = stats["success"]
        failureCount = totalLaunches - successCount
        successRate = (successCount / totalLaunches * 100) if totalLaunches > 0 else 0

        data.append({
            "year": year,
            "totalLaunches": totalLaunches,
            "successfulLaunches": successCount,
            "failedLaunches": failureCount,
            "successRate": successRate
        })

    analytic = Analytic(
        id="historicalSuccessFailure",
        metricName="Historical Success & Failure Rates per Year",
        description="This metric tracks the yearly success and failure rates of space launches, "
                    "providing insights into trends in reliability and launch performance over time.",
        data=data,
        updatedAt=getCurrentTimestamp()
    )
    return analytic


def generateMissionsPerComponentTypeMetric(components: List[Component]) -> Analytic[int]:
    """
    Generates a metric showing the number of missions per component type.

    Args:
        components (List[Component]): List of processed components.

    Returns:
        Analytic[int]: Number of missions per component type.
    """
    missionCountMap: Dict[str, int] = defaultdict(int)

    for component in components:
        missionCountMap[component.type] += len(component.missions)

    data = [{"componentType": componentType, "totalMissions": totalMissions}
            for componentType, totalMissions in missionCountMap.items()]

    analytic = Analytic[Union[Decimal, int, str]](
        id="missionsPerComponentType",
        metricName="Number of Missions per Component Type",
        description="This metric tracks how many missions each type of component has participated in, "
                    "helping to analyze component utilization trends across space missions.",
        data=data,
        updatedAt=getCurrentTimestamp()
    )

    return analytic

def generateComponentStatusDistributionMetric(components: List[Component]) -> Analytic[int]:
    """
    Generates a metric showing the distribution of component statuses.

    Args:
        components (List[Component]): List of processed components.

    Returns:
        Analytic[int]: Count of components per status.
    """
    statusCountMap: Dict[str, int] = defaultdict(int)

    for component in components:
        statusCountMap[component.status] += 1
    data = [{"status": status, "count": count} for status, count in statusCountMap.items()]

    analytic = Analytic[Union[int, ComponentStatus]](
        id="componentStatusDistribution",
        metricName="Overall Component Status Distribution",
        description="This metric tracks the overall distribution of component statuses, "
                    "providing insights into the health and operational status of space components.",
        data=data,
        updatedAt=getCurrentTimestamp()
    )

    return analytic



def generateRegionSuccessMetric(launches: List[LaunchDTO], launchpads: List[LaunchpadDTO]) -> Analytic[Union[int, str, float]]:
    """
    Generates a metric tracking how launch success rates vary across different launch sites (with coordinates).

    Args:
        launches (List[LaunchDTO]): List of launches.
        launchpads (List[LaunchpadDTO]): List of launchpads.

    Returns:
        Analytic[Union[int, str, float]]: Success rate per launch site with coordinates.
    """

    launchpadMap: Dict[str, Dict[str, Union[str, float]]] = {
        pad.id: {"region": pad.region, "latitude": pad.latitude, "longitude": pad.longitude} for pad in launchpads
    }

    launchpadStats: Dict[str, Dict[str, int]] = defaultdict(lambda: {"total": 0, "success": 0})

    for launch in launches:
        launchpadId = launch.launchpad
        if launchpadId in launchpadMap:
            launchpadStats[launchpadId]["total"] += 1
            if launch.success:
                launchpadStats[launchpadId]["success"] += 1

    data = [
        {
            "launchpad": launchpadId,
            "region": launchpadMap[launchpadId]["region"],
            "latitude": launchpadMap[launchpadId]["latitude"],
            "longitude": launchpadMap[launchpadId]["longitude"],
            "totalLaunches": stats["total"],
            "successfulLaunches": stats["success"],
            "successRate": (stats["success"] / stats["total"] * 100) if stats["total"] > 0 else 0
        }
        for launchpadId, stats in launchpadStats.items()
    ]

    analytic = Analytic[Union[str, float]](
        id="regionSuccessRate",
        metricName="Impact of Region on Launch Success",
        description="This metric analyzes how launch success rates vary across different launch sites, "
                    "including geographic coordinates for spatial analysis.",
        data=data,
        updatedAt=getCurrentTimestamp()
    )

    return analytic


def generateMaintenanceFrequencyMetric(maintenances: List[dict], components: List[Component]) -> Analytic[int]:
    """
    Generates a metric showing the frequency of maintenance events per component type.

    Args:
        maintenances (List[dict]): List of maintenance records retrieved as dictionaries.
        components (List[Component]): List of components.

    Returns:
        Analytic[int]: Maintenance frequency per component type.
    """

    maintenanceObjects: List[Maintenance] = [Maintenance(**m) if isinstance(m, dict) else m for m in maintenances]
    componentTypeMap: Dict[str, str] = {comp.componentId: comp.type for comp in components}
    maintenanceCountMap: Dict[str, int] = defaultdict(int)

    for maintenance in maintenanceObjects:
        print(f"Processed Maintenance: {maintenance}")
        componentType = componentTypeMap.get(maintenance.componentId, "Unknown")
        maintenanceCountMap[componentType] += 1

    data = [{"componentType": compType, "maintenanceCount": count} for compType, count in maintenanceCountMap.items()]

    analytic = Analytic[Union[ComponentType, int]](
        id="maintenanceFrequency",
        metricName="Maintenance Frequency per Component Type",
        description="This metric tracks how frequently each type of component undergoes maintenance, "
                    "helping to identify which components require the most servicing.",
        data=data,
        updatedAt=getCurrentTimestamp()
    )

    return analytic

