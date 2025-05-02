import { Route, Routes } from "react-router-dom";

import LoginPage from "@/pages/login";
import HomePage from "@/pages/home.tsx";

function App() {
  return (
    // 定义路由
    <Routes>
      <Route element={<HomePage />} path="/" />
      <Route element={<LoginPage />} path="/login" />
    </Routes>
  );
}

export default App;
