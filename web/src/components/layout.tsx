interface LayoutProps {
  children: any;
}

export default function Layout({ children }: LayoutProps) {
  return <div className="bg-gray-50 min-h-screen mt-24">{children}</div>;
}
