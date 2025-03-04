import Layout from "@/components/layout";
import NavBar from "./components/ui/navBar";

import { ComponentService } from "./ports/components";
import { ComponentTable } from "./components/table/ComponentTable";
import { Component } from "./domain/entities/component";

import { useEffect, useState } from "react";

const componentService = new ComponentService();

export default function App() {
  const [components, setComponents] = useState<Component[]>([]);

  useEffect(() => {
    const fetchComponents = async () => {
      const data = await componentService.findAll();
      setComponents(data);
    };

    fetchComponents();
  }, []);

  return (
    <>
      <Layout>
        <NavBar />
        <main className="container mx-auto p-4">
          <div className="flex w-1/2">
          <ComponentTable data={components} />
          </div>
        </main>
      </Layout>
    </>
  );
}
