interface LayoutProps {
  children: any;
}

export default function Layout({ children }: LayoutProps) {
  return <div className="bg-gray-100 min-h-screen">{children}</div>;
}
