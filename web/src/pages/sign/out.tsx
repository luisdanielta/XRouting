import Layout from "@/components/layout";
import PrimaryButton from "@/components/ui/button/primaryButton";
import SecondaryButton from "@/components/ui/button/secondaryButton";

export default function SignOut() {
  return (
    <Layout>
      <main className="min-h-screen flex items-center justify-center bg-gradient-to-r from-blue-100 via-blue-50 to-green-50">
        <section className="shadow rounded-sm p-8 sm:p-12 w-full max-w-md bg-gray-50 mx-4">
          <h1 className="text-4xl font-bold text-center mb-6 bg-gradient-to-r from-blue-500 via-blue-200 to-green-500 bg-clip-text text-transparent">
            Sign Out
          </h1>

          <article>
            <p className="mt-4 text-center text-sm text-gray-600">
              You are signed out successfully.
            </p>
            <div className="flex justify-between mt-4 gap-4">
              <SecondaryButton text="Sing Up" />
              <PrimaryButton text="Sign In" />
            </div>
          </article>
        </section>
      </main>
    </Layout>
  );
}
