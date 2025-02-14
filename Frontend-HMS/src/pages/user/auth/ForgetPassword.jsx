import AuthBase from '@components/organisms/AuthBase';
import ForgetPasswordForm from '@features/users/ForgetPasswordForm';
import Navbar from '@components/templates/Navbar';

const ForgetPassword = () => {
  return (
    <>
      <Navbar />
      <AuthBase>
        <ForgetPasswordForm />
      </AuthBase>
    </>
  );
};

export default ForgetPassword;
