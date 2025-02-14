import { useState} from 'react';
import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';
import { useForgotPasswordMutation } from '@/api/authApi';
import AuthBase from '@/components/organisms/AuthBase'; 


// Define validation schema using Yup
const schema = yup.object().shape({
  email: yup
    .string()
    .email('Email address is invalid')
    .required('Email is required'),
});

const SigninForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: yupResolver(schema),
  });

  const [forgotPasswordMutation, { isLoading }] = useForgotPasswordMutation();
  const [message, setMessage] = useState('');
  const navigate = useNavigate();

  const onSubmit = async (data) => {
    try {
      await forgotPasswordMutation(data).unwrap();
      setMessage('Password reset link sent to your email.');
    } catch (err) {
      setMessage('Error sending reset link. Please try again.');
      navigate('/reset-password');
      console.error('Forget password error:', err);
    }
  };

  return (
      <form className="flex flex-col gap-8" onSubmit={handleSubmit(onSubmit)}>
        {/* FORM HEADING */}
        <div className="flex flex-col gap-2">
          <h4 className="font-semibold text-[22px] leading-[26.4px] text-[black]">
            Forget Password
          </h4>
          <span className="font-normal text-sm leading-[19.6px] text-[black]">
            <p>
              You can easily reset your password here. We&#39;ll send the reset
              link to your email.
            </p>
          </span>
        </div>

        {/* Form content */}
        <div className="flex flex-col gap-[32px]">
          <div className="flex flex-col gap-5">
            <div className="input-group">
              <input
                className="input w-[548px] h-12 border border-[#dcdcdc] rounded-lg pl-2.5 pr-3 py-3.5"
                type="email"
                {...register('email')}
                autoComplete="email"
                placeholder=" "
              />
              <label className="placeholder" htmlFor="email">
                {' '}
                Email{' '}
              </label>
              {errors.email && (
                <p className="text-red-500 text-sm">{errors.email.message}</p>
              )}
            </div>

            <button
              type="submit"
              className="w-[548px] h-12 bg-primary-purple text-white rounded-lg flex justify-center items-center font-medium"
            >
              {isLoading ? 'Loading...' : 'Reset password'}
            </button>

            {message && (
              <p className="text-sm text-center text-gray-600">{message}</p>
            )}
          </div>
        </div>
      </form>
  );
};

export default SigninForm;
