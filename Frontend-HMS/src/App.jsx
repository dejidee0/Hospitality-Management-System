import { Routes, Route, Navigate } from 'react-router-dom';
import './App.css';
import Signup from './users/auth/Signup';
import Signin from './users/auth/Signin';
import ForgetPassword from './users/auth/ForgetPassword';
import ResetPassword from './users/auth/ResetPassword';
import LandingPage from './components/LandingPage';
import Search from './users/hotel/Search';
import Home from './users/hotel/Home';

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
        <Route path="/home" element={<Home />} />
      </Routes>
    </>
  );
}

export default App;
