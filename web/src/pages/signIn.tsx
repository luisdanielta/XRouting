export default function SignIn() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-r from-blue-50 via-purple-50 to-pink-50">
      <div className="bg-white shadow-xl rounded-xl p-8 sm:p-12 w-full max-w-md container cq-container">
        <h1 className="text-3xl font-bold text-center mb-6 text-gray-800 animate-fadeInDown">
          Sign In
        </h1>
        <form className="space-y-5">
          <div>
            <label className="block text-gray-700">Email</label>
            <input
              type="email"
              required
              className="mt-1 w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring focus:border-blue-200 transition-all duration-200 animate-fadeInUp"
            />
          </div>
          <div>
            <label className="block text-gray-700">Password</label>
            <input
              type="password"
              required
              className="mt-1 w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring focus:border-blue-200 transition-all duration-200 animate-fadeInUp"
            />
          </div>
          <button
            type="submit"
            className="w-full py-3 rounded-lg bg-blue-500 text-white font-bold hover:bg-blue-600 transition-all duration-200 animate-fadeInUp"
          >
            Sign In
          </button>
        </form>
        <p className="mt-4 text-center text-sm text-gray-600 animate-fadeInUp">
          Don't have an account?{" "}
          <a href="/signup" className="text-blue-500 hover:underline">
            Sign up
          </a>
        </p>
      </div>
    </div>
  );
}
