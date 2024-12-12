import { Routes, Route } from "react-router-dom";
import "./App.css";
import Signup from "./Authetication/Signup";
import Signin from "./Authetication/Signin";
import ForgetPassword from "./Authetication/ForgetPassword";
import ResetPassword from "./Authetication/ResetPassword";
import LandingPage from "./components/LandingPage";
import Search from "./pages/Search";

function App() {
  return (
    <>
      <Routes>
        <Route path="/signup" element={<Signup />} />
        <Route path="/signin" element={<Signin />} />
        <Route path="/forget-password" element={<ForgetPassword />} />
        <Route path="/reset-link" element={<ResetPassword />} />
        <Route path="/landing" element={<LandingPage />} />
        <Route path="/search" element={<Search />} />


      </Routes>
    </>
  );
}

export default App;
