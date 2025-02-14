import AuthBase from '@components/organisms/AuthBase';
import SigninForm from '@features/users/SigninForm';
import Navbar from '@components/templates/Navbar';

const Signin = () => {
  return (
    <>
      <Navbar />
      <AuthBase>
        <SigninForm />
      </AuthBase>
    </>
  );
};

export default Signin;
