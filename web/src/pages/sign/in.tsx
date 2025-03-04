import Layout from "@/components/layout";
import PrimaryButton from "@/components/ui/button/primaryButton";
import Input from "@/components/ui/input";

export default function SignIn() {
  return (
    <Layout>
      <main className="min-h-screen flex items-center justify-center bg-gradient-to-r from-blue-100 via-blue-50 to-green-50">
        <section className="shadow rounded-md p-8 sm:p-12 w-full max-w-md bg-gray-50 mx-4">
          <h1 className="text-4xl font-bold text-center mb-6 bg-gradient-to-r from-blue-500 via-blue-200 to-green-500 bg-clip-text text-transparent">
            Sign In
          </h1>

          <form className="space-y-4">
            <div>
              <label htmlFor="email" className="">
                Email
              </label>
              <Input type="email" />
            </div>
            <div>
              <label htmlFor="password" className="">
                Password
              </label>
              <Input type="password" />
            </div>
            <PrimaryButton text="Sign In" />
          </form>
          <p className="mt-4 text-center text-sm text-gray-600">
            Don't have an account?{" "}
            <a href="/sign/up" className="text-blue-600 hover:underline">
              Sign Up
            </a>
          </p>
        </section>
      </main>
    </Layout>
  );
}
