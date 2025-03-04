import SecondaryButton from "./button/secondaryButton";
import Input from "./input";
import IconBox from "./iconBox";
import ISearch from "@/icons/iSearch";

export default function NavBar() {
  return (
    <nav className="bg-gradient-to-r from-blue-100 via-blue-50 to-green-50 shadow-md shadow-gray-100 fixed top-0 w-full z-10">
      <div className="container mx-auto flex flex-wrap justify-between items-center px-4 py-4">
        <h1 className="text-2xl font-bold bg-gradient-to-r from-blue-500 via-blue-300 to-green-500 bg-clip-text text-transparent">
          XRounting
        </h1>

        <section className="hidden md:flex items-center gap-2 w-full md:w-96">
          <Input placeholder="Buscar" className="w-full max-w-xs" />
          <span>
            <IconBox>
              <ISearch />
            </IconBox>
          </span>
        </section>

        <div className="flex items-center gap-2">
          <div className="md:hidden">
            <IconBox>
              <ISearch />
            </IconBox>
          </div>
          <SecondaryButton text="luisdanielta" />
        </div>
      </div>
    </nav>
  );
}
