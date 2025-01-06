import { Routes, Route } from "react-router-dom";
import { useSelector, useDispatch } from 'react-redux';
import { useEffect } from "react";
import "./App.css";
import Signup from "./Authetication/Signup";
import Signin from "./Authetication/Signin";
import ForgetPassword from "./Authetication/ForgetPassword";
import ResetPassword from "./Authetication/ResetPassword";
import LandingPage from "./components/LandingPage";
import Search from "./pages/Search";
import Home from "./users/hotel/Home";

function App() {
  const dispatch = useDispatch();
  const { isAuthenticated } = useSelector((state) => state.auth);

  useEffect(() => {
    dispatch(checkAuth());
  }, [dispatch]);

  return (
    <>
      <Routes>
        <Route 
          path="/signup" 
          element={!isAuthenticated ? <Signup /> : <Navigate to="/Home" />}
        />
        {/* <Route
          path="/dashboard"
          element={isAuthenticated ? <Dashboard /> : <Navigate to="/login" />}
        /> */}
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
