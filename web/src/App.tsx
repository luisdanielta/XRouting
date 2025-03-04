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

import Plot from 'react-plotly.js';

const componentService = new ComponentService();
const maintenanceService = new MaintenanceService();

export default function App() {
  const [components, setComponents] = useState<Component[]>([]);
  const [maintenances, setMaintenances] = useState<Maintenance[]>([]);

  useEffect(() => {
    const fetchComponents = async () => {
      const data = await componentService.findAll();
      setComponents(data);

      const maintenanceData = await maintenanceService.findAll();
      setMaintenances(maintenanceData);
    };

    fetchComponents();
  }, []);

  const tableComponent = useComponentTable(components);
  const tableMaintenance = useMaintenanceTable(maintenances);

  return (
    <>
      <Layout>
        <NavBar />
        <main className="container mx-auto p-4">
          <section className="grid grid-cols-4 gap-4">
            <article className="col-span-2">
              <h1 className="text-4xl font-bold text-center mb-6 bg-gradient-to-r from-blue-500 via-blue-200 to-green-500 bg-clip-text text-transparent">
                Components
              </h1>
              <Plot
                data={[
                  {
                    x: components.map((component) => component.status),
                    y: components.map((component) => component.type),
                    type: 'bar',
                  },
                ]}
                layout={{ width: 450, title: 'Component Maintenance Count' }}
              />
            </article>
            <article className="col-span-2">
              <h1 className="text-4xl font-bold text-center mb-6 bg-gradient-to-r from-blue-500 via-blue-200 to-green-500 bg-clip-text text-transparent">
                Maintenances
              </h1>
              <Plot
                data={[
                  {
                    x: maintenances.map((maintenance) => maintenance.date),
                    y: maintenances.map((maintenance) => maintenance.cost),
                    mode: 'markers',
                    type: 'scatter',
                  },
                ]}
                layout={{ width: 450, title: 'Maintenance Count' }}
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
