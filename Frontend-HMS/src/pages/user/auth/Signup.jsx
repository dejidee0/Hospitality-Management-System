import AuthBase from './AuthBase';
import SignupForm from '../../../features/users/SignupForm';
import Navbar from '../../../components/templates/Navbar';

const Signup = () => {
  return (
    <>
      <Navbar />
      <AuthBase>
        <SignupForm />
      </AuthBase>
    </>
  );
};

export default Signup;
