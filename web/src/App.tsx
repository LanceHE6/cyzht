import { Route, Routes } from "react-router-dom";

import DocsPage from "@/pages/docs";
import PricingPage from "@/pages/pricing";
import BlogPage from "@/pages/blog";
import AboutPage from "@/pages/about";
import LoginPage from "@/pages/login";
import Home from "@/pages/home.tsx";

function App() {
  return (
    // 定义路由
    <Routes>
      <Route element={<Home />} path="/" />
      <Route element={<LoginPage />} path="/login" />
      <Route element={<DocsPage />} path="/docs" />
      <Route element={<PricingPage />} path="/pricing" />
      <Route element={<BlogPage />} path="/blog" />
      <Route element={<AboutPage />} path="/about" />
    </Routes>
  );
}

export default App;
