import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router";

import "./index.css";
import App from "./App.tsx";
import SignLayout from "./pages/sign/index.tsx";
import SignIn from "./pages/sign/in.tsx";
import SignUp from "./pages/sign/up.tsx";
import SignOut from "./pages/sign/out.tsx";
import { UserProvider } from "./context/userContext.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <UserProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<App />} />
          <Route element={<></>}></Route>
          <Route path="/sign" element={<SignLayout />}>
            <Route path="in" element={<SignIn />} />
            <Route path="up" element={<SignUp />} />
            <Route path="out" element={<SignOut />} />
          </Route>
          <Route path="*" element={<div>404</div>} />
        </Routes>
      </BrowserRouter>
    </UserProvider>
  </StrictMode>,
);
