import Layout from "@/components/layout";
import NavBar from "./components/ui/navBar";

export default function App() {
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
