// import { Link } from "@heroui/link";

import { ToastContainer } from "react-toastify";

import { Navbar } from "@/components/navbar";

export default function DefaultLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="relative flex flex-col h-screen">
      <Navbar />
      <main className="container mx-auto max-w-7xl px-6 flex-grow">
        {children}
      </main>
      {/*<footer className="w-full flex items-center justify-center py-3">*/}
      {/*  <Link*/}
      {/*    isExternal*/}
      {/*    className="flex items-center gap-1 text-current"*/}
      {/*    href="https://heroui-docs-v2.vercel.app?utm_source=next-pages-template"*/}
      {/*    title="heroui.org homepage"*/}
      {/*  >*/}
      {/*    <span className="text-default-600">Powered by</span>*/}
      {/*    <p className="text-primary">NextUI</p>*/}
      {/*  </Link>*/}
      {/*</footer>*/}
      <ToastContainer />
    </div>
  );
}
