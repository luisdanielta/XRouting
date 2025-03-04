import { z } from "zod";

export const AnalyticSchema = z.object({
  id: z.string(),
  metricName: z
    .string()
    .describe("The name of the metric (e.g., 'Core Reuse Rate')"),
  description: z.string().describe("Detailed explanation of the metric"),
  data: z
    .array(z.record(z.unknown()))
    .default([])
    .describe("Aggregated data related to the metric"),
  updatedAt: z.string().describe("Last updated timestamp for this metric"),
});

export type Analytic = z.infer<typeof AnalyticSchema>;

export function convertToDecimal<T extends Record<string, unknown>>(
  data: T[],
): T[] {
  return data.map((item) => {
    const newItem: Record<string, unknown> = { ...item };
    for (const key in newItem) {
      if (typeof newItem[key] === "number") {
        newItem[key] = parseFloat((newItem[key] as number).toString()).toFixed(
          2,
        );
      }
    }
    return newItem as T;
  });
}

type RegionSuccessRateData = {
  latitude: number;
  longitude: number;
  launchpad: string;
  region: string;
  successRate: number;
  successfulLaunches: number;
  totalLaunches: number;
};

type ComponentStatusData = {
  count: number;
  status: string;
};

type CostSuccessRateData = {
  active: boolean;
  costPerLaunch: number;
  engineCount: number;
  rocketId: string;
  rocketName: string;
  successRate: number;
  totalLaunches: number;
};

type HistoricalSuccessFailureData = {
  year: number;
  failedLaunches: number;
  successfulLaunches: number;
  totalLaunches: number;
  successRate: number;
};

type MissionsPerComponentData = {
  componentType: string;
  totalMissions: number;
};

type LaunchCostDistributionData = {
  region: string;
  totalLaunchCost: number;
};

type MaintenanceFrequencyData = {
  componentType: string;
  maintenanceCount: number;
};

type Metric<T> = {
  id: string;
  metricId: string;
  metricName: string;
  data: T[];
  updatedAt: string;
};

export type {
  RegionSuccessRateData,
  ComponentStatusData,
  CostSuccessRateData,
  HistoricalSuccessFailureData,
  MissionsPerComponentData,
  LaunchCostDistributionData,
  MaintenanceFrequencyData,
  Metric,
};
