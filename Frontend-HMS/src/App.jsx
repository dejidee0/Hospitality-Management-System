import { Routes, Route } from "react-router-dom";
import "./App.css";
import Signup from "./Authetication/Signup";
import Signin from "./Authetication/Signin";
import ForgetPassword from "./Authetication/ForgetPassword";
import ResetPassword from "./Authetication/ResetPassword";

function App() {
  return (
    <>
      <Routes>
        <Route path="/signup" element={<Signup />} />
        <Route path="/signin" element={<Signin />} />
        <Route path="/forget-password" element={<ForgetPassword />} />
        <Route path="/reset-link" element={<ResetPassword />} />
      </Routes>
    </>
  );
}

export default App;
