import { ISignOut } from "@/icons/signOut";

export default function NavBar() {
  return (
    <nav className="bg-gradient-to-r from-blue-100 via-blue-50 to-green-50 shadow-md shadow-gray-100">
      <div className="container mx-auto flex justify-between items-center h-full p-4">
        <h1 className="font-bold text-xl text-green-700">XRounting</h1>

        <div className="flex items-center gap-4">
          <section className="flex items-center gap-2 border-2 border-green-300 py-1 px-4 rounded-full hover:bg-green-100 cursor-pointer transition-colors duration-200 ease-in hover:ease-out shadow-md shadow-green-100">
            <p className="text-sm font-bold">John Doe</p>
          </section>

          <div className="text-red-700 cursor-pointer">
            <ISignOut />
          </div>
        </div>
      </div>
    </nav>
  );
}
