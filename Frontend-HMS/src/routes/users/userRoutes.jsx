import { Routes, Route } from 'react-router-dom';
import Signup from '@/pages/user/auth/Signup';
import Signin from '@/pages/user/auth/Signin';
// import { ConfirmChange } from '../../pages/user/auth/ComfirmChange';
// import { ResetPassword } from '../../pages/user/auth/ResetPassword';
// import { Home } from '../../pages/user/Home';
// import { Search } from '../../pages/user/Search';

const UsersRoutes = () => {
  return (
    <Routes>
      <Route path="/signup" element={<Signup />} />
      <Route path="/signin" element={<Signin />} />
      {/* <Route path="confirm-change/" element={<ConfirmChange />} />
      <Route path="reset-password/" element={<ResetPassword />} />
      <Route path="home/" element={<Home />} />
      <Route path="search/" element={<Search />} /> */}
    </Routes>
  );
};

export default UsersRoutes;
