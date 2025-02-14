import { Routes, Route } from 'react-router-dom';
import Signup from '@/pages/user/auth/Signup';
import Signin from '@/pages/user/auth/Signin';
// import { ConfirmChange } from '../../pages/user/auth/ComfirmChange';
import ForgetPassword from '@/pages/user/auth/ForgetPassword';
import ResetLinkSent from '@/pages/user/auth/ResetLinkSent';
// import { Home } from '../../pages/user/Home';
// import { Search } from '../../pages/user/Search';

const UsersRoutes = () => {
  return (
    <Routes>
      <Route path="/signup" element={<Signup />} />
      <Route path="/signin" element={<Signin />} />
      {/* <Route path="confirm-change/" element={<ConfirmChange />} /> */}
      <Route path="reset-password/" element={<ForgetPassword />} />
      <Route path="reset-sent/" element={<ResetLinkSent />} />
      {/* <Route path="home/" element={<Home />} /> */}
      {/* <Route path="search/" element={<Search />} /> */}
    </Routes>
  );
};

export default UsersRoutes;
