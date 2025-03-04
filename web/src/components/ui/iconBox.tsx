export default function IconBox({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex flex-col items-center text-gray-800 justify-center w-full h-full p-4 bg-gradient-to-r from-blue-200 via-green-200 w-full to-green-300 text-gray-800 hover:text-gray-900 hover:bg-gradient-to-r hover:from-blue-300 hover:to-green-300 focus:ring-0 hover:ease-in shadow-gray-100 hover:shadow-gray-200 transition-colors duration-200 ease-out cursor-pointer shadow-md rounded">
      {children}
    </div>
  );
}
