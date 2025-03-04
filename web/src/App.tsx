import Layout from "@/components/layout";
import NavBar from "./components/ui/navBar";

import { ComponentService } from "./ports/components";

export default function App() {

  const componentService = new ComponentService();
  componentService.findAll().then((res) => console.log(res));

  return (
    <>
      <Layout>
        <NavBar />
        <main className="container mx-auto p-4">
          <h1 className="font-bold font-lg">Hi</h1>
        </main>
      </Layout>
    </>
  );
}
