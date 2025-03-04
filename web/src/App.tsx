import Layout from "@/components/layout";
import NavBar from "./components/ui/navBar";

import { ComponentService } from "./ports/components";
import { Component } from "./domain/entities/component";

import { useEffect, useState } from "react";
import { MaintenanceService } from "./ports/maintenances";
import { Maintenance } from "./domain/entities/maintenance";
import { useComponentTable } from "./hooks/useComponentTable";
import { Table } from "./components/table";
import { useMaintenanceTable } from "./hooks/useMaintenanceTable";

import { AnalyticService } from "./ports/analytics";
import {
  Analytic,
  ComponentStatusData,
  CostSuccessRateData,
  HistoricalSuccessFailureData,
  LaunchCostDistributionData,
  MaintenanceFrequencyData,
  Metric,
  MissionsPerComponentData,
  RegionSuccessRateData,
} from "./domain/entities/analytic";

import Plot from "react-plotly.js";

const componentService = new ComponentService();
const maintenanceService = new MaintenanceService();
const analyticsService = new AnalyticService();

export default function App() {
  const [components, setComponents] = useState<Component[]>([]);
  const [maintenances, setMaintenances] = useState<Maintenance[]>([]);
  const [analytics, setAnalytics] = useState<Analytic[]>([]);

  useEffect(() => {
    const fetchComponents = async () => {
      const data = await componentService.findAll();
      setComponents(data);

      const maintenanceData = await maintenanceService.findAll();
      setMaintenances(maintenanceData);

      const analyticsData = await analyticsService.findAll();
      setAnalytics(analyticsData);
    };

    fetchComponents();
  }, []);

  const tableComponent = useComponentTable(components);
  const tableMaintenance = useMaintenanceTable(maintenances);

  const rawData: any[] = analytics;

  const regionSuccessRate = rawData.find(
    (m) => m.id === "regionSuccessRate",
  ) as Metric<RegionSuccessRateData>;
  const componentStatusDistribution = rawData.find(
    (m) => m.id === "componentStatusDistribution",
  ) as Metric<ComponentStatusData>;
  const costSuccessRate = rawData.find(
    (m) => m.id === "costSuccessRate",
  ) as Metric<CostSuccessRateData>;
  //const historicalSuccessFailure = rawData.find(m => m.id === "historicalSuccessFailure") as Metric<HistoricalSuccessFailureData>;
  const missionsPerComponentType = rawData.find(
    (m) => m.id === "missionsPerComponentType",
  ) as Metric<MissionsPerComponentData>;
  // const launchCostDistribution = rawData.find(m => m.id === "launchCostDistribution") as Metric<LaunchCostDistributionData>;
  const maintenanceFrequency = rawData.find(
    (m) => m.id === "maintenanceFrequency",
  ) as Metric<MaintenanceFrequencyData>;

  return (
    <>
      <Layout>
        <NavBar />
        <main className="container mx-auto p-4 flex flex-col gap-4">
          <section className="grid grid-cols-2 gap-4">
            <article className="flex flex-col w-full justify-center items-center">
              <Plot
                data={
                  regionSuccessRate && componentStatusDistribution
                    ? [
                        {
                          x: regionSuccessRate.data.map((data) => data.region),
                          y: componentStatusDistribution.data.map(
                            (data) => data.count,
                          ),
                          type: "bar",
                        },
                      ]
                    : []
                }
                layout={{ title: { text: "Region Success Rate" }, width: 500 }}
              />
            </article>

            <article className="flex flex-col w-full justify-center items-center">
              <Plot
                data={
                  maintenanceFrequency
                    ? [
                        {
                          x: maintenanceFrequency.data.map(
                            (data) => data.componentType,
                          ),
                          y: maintenanceFrequency.data.map(
                            (data) => data.maintenanceCount,
                          ),
                          type: "bar",
                        },
                      ]
                    : []
                }
                layout={{
                  title: { text: "Maintenance Frequency" },
                  width: 500,
                }}
              />
            </article>

            <article className="flex flex-col w-full justify-center items-center">
              <Plot
                data={
                  costSuccessRate
                    ? [
                        {
                          x: costSuccessRate.data.map(
                            (data) => data.rocketName,
                          ),
                          y: costSuccessRate.data.map(
                            (data) => data.costPerLaunch,
                          ),
                          type: "bar",
                        },
                      ]
                    : []
                }
                layout={{ title: { text: "Cost per Launch" }, width: 500 }}
              />
            </article>

            <article className="flex flex-col w-full justify-center items-center">
              <Plot
                data={
                  missionsPerComponentType
                    ? [
                        {
                          x: missionsPerComponentType.data.map(
                            (data) => data.componentType,
                          ),
                          y: missionsPerComponentType.data.map(
                            (data) => data.totalMissions,
                          ),
                          type: "scatter",
                          mode: "lines+markers",
                        },
                      ]
                    : []
                }
                layout={{
                  title: { text: "Missions per Component Type" },
                  width: 500,
                }}
              />
            </article>
          </section>

          <section className="flex flex-row gap-4">
            <article>
              <h1 className="text-4xl font-bold text-center mb-6 bg-gradient-to-r from-blue-500 via-blue-200 to-green-500 bg-clip-text text-transparent">
                Components
              </h1>
              <Table table={tableComponent} />
            </article>
            <article>
              <h1 className="text-4xl font-bold text-center mb-6 bg-gradient-to-r from-blue-500 via-blue-200 to-green-500 bg-clip-text text-transparent">
                Maintenances
              </h1>
              <Table table={tableMaintenance} />
            </article>
          </section>
        </main>
      </Layout>
    </>
  );
}
