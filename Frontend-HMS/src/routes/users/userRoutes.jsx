import { Routes, Route } from 'react-router-dom';
import { Signup } from '../../pages/user/auth';
import { Signin } from '../../pages/user/Signin';
import { ConfirmChange } from '../../pages/user/ConfirmChange';
import { ResetPassword } from '../../pages/user/ResetPassword';
import { Home } from '../../pages/user/Home'
import { Search } from '../../pages/user/Search';

const UsersRoutes = () => {
  return (
    <Routes>
      <Route path="signup/" element={<Signup />} />
      <Route path="signin/" element={<Signin />} />
      <Route path="confirm-change/" element={<ConfirmChange />} />
      <Route path="reset-password/" element={<ResetPassword />} />
      <Route path="home/" element={<Home />} />
      <Route path="search/" element={<Search />} />
    </Routes>
  );
};

export default UsersRoutes;
