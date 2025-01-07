import { Routes, Route } from "react-router-dom";
import { useSelector, useDispatch } from 'react-redux';
import { useEffect } from "react";
import "./App.css";
import Signup from "./users/auth/Signup";
import Signin from "./users/auth/Signin";
import ForgetPassword from "./users/auth/ForgetPassword";
import ResetPassword from "./users/auth/ResetPassword";
import LandingPage from "./components/LandingPage";
import Search from "./users/hotel/Home";
import Home from "./users/hotel/Home";
import { checkAuth } from "./redux/slices/users/authSlice";


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
          element={!isAuthenticated ? <Signup /> : <Navigate to="/LandingPage" />}
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
