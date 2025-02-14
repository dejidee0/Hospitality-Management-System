import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';
import { useRegisterMutation } from '../../api/authApi';
import { useDispatch } from 'react-redux';
import { setCredentials } from '../../redux/slices/users/authSlice';
import { CgDanger } from 'react-icons/cg';
import googleIcon from '../../assets/google_icon.svg';
import facebookIcon from '../../assets/facebook_icon.svg';
import { LuEye, LuEyeOff } from 'react-icons/lu';

// Define validation schema using Yup
const schema = yup.object().shape({
  email: yup
    .string()
    .email('Email address is invalid')
    .required('Email is required'),
  password: yup
    .string()
    .min(8, 'Password must be at least 8 characters')
    .required('Password is required'),
  confirmPassword: yup
    .string()
    .oneOf([yup.ref('password'), null], "Passwords don't match")
    .required('Confirm Password is required'),
});

const SignupForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: yupResolver(schema),
  });

  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [registerMutation, { isLoading }] = useRegisterMutation();
  // const dispatch = useDispatch();
  const navigate = useNavigate();

  const onSubmit = async (data) => {
    try {
      const response = await registerMutation(data).unwrap();
      // dispatch(setCredentials(response)); // Update auth state
      navigate('/signin'); // Redirect to sign-in page after successful signup
    } catch (err) {
      console.error('Signup error:', err);
    }
  };

  return (
    <form className="flex flex-col gap-8" onSubmit={handleSubmit(onSubmit)}>
      <div className="flex flex-col gap-2">
        <h4 className="font-semibold text-[22px] leading-[26.4px] text-[black]">
          Create an account
        </h4>
        <span className="font-normal text-sm leading-[19.6px] text-[black] flex gap-[5px]">
          <p>Have an account?</p>
          <Link
            to="/signin"
            className="text-[color:var(--primary-purple)] underline"
          >
            Sign In
          </Link>
        </span>
      </div>

      <div className="flex flex-col gap-[32px]">
        <div className="flex gap-[15px]">
          <div className="bg-[color:var(--default-grey)] h-12 w-[266.5px] flex justify-center items-center gap-2.5 rounded-lg">
            <img src={googleIcon} alt="Google Icon" className="w-6 h-6" />
            <Link className="text-sm font-normal leading-[19.6px]">Google</Link>
          </div>
          <div className="bg-[color:var(--default-grey)] h-12 w-[266.5px] flex justify-center items-center gap-2.5 rounded-lg">
            <img src={facebookIcon} alt="facebook login" className="w-6 h-6" />
            <Link className="text-sm font-normal leading-[19.6px]">
              Facebook
            </Link>
          </div>
        </div>

        <div className="flex items-center justify-center gap-3 relative text-center text-sm font-normal text-black">
          <span className="w-[200.5px] grow h-px bg-[#E6E7E6] mr-2.5"></span>
          Or Sign up with
          <span className="w-[200.5px] grow h-px bg-[#E6E7E6] mr-2.5"></span>
        </div>

        <div className="flex flex-col gap-5">
          {/* Email Input */}
          <div className="input-group">
            <input
              className="input w-[548px] h-12 border rounded-lg pl-2.5 pr-3 py-3.5 border-solid border-[#dcdcdc]"
              type="text"
              id="email"
              {...register('email')}
              placeholder=""
            />
            <label className="placeholder" htmlFor="email">
              Email
            </label>
            {errors.email && (
              <div className="flex items-center gap-1 h-5 mt-[5px]">
                <CgDanger className="w-[16.7px] h-[16.7px] bg-[#EF1212] text-white rounded-full" />
                <p className="text-[#EF1212] text-xs">{errors.email.message}</p>
              </div>
            )}
          </div>

          {/* Password Input */}
          <div className="input-group">
            <input
              className="input w-[548px] h-12 border rounded-lg pl-2.5 pr-3 py-3.5 border-solid border-[#dcdcdc]"
              type={showPassword ? 'text' : 'password'}
              id="password"
              {...register('password')}
              placeholder=""
            />
            <label className="placeholder" htmlFor="password">
              Password
            </label>
            <span
              className="password__toggle absolute -translate-y-2/4 cursor-pointer text-[#666] right-5 top-[25px]"
              onClick={() => setShowPassword(!showPassword)}
              onMouseDown={(e) => e.preventDefault()}
            >
              {showPassword ? <LuEye /> : <LuEyeOff />}
            </span>
            {errors.password && (
              <div className="flex items-center gap-1 h-5 mt-[5px]">
                <CgDanger className="w-[16.7px] h-[16.7px] bg-[#EF1212] text-white rounded-full" />
                <p className="text-[#EF1212] text-xs">
                  {errors.password.message}
                </p>
              </div>
            )}
          </div>

          {/* Confirm Password Input */}
          <div className="input-group">
            <input
              className="input w-[548px] h-12 border rounded-lg pl-2.5 pr-3 py-3.5 border-solid border-[#dcdcdc]"
              type={showConfirmPassword ? 'text' : 'password'}
              id="confirmPassword"
              {...register('confirmPassword')}
              placeholder=""
            />
            <label className="placeholder" htmlFor="confirmPassword">
              Confirm Password
            </label>
            <span
              className="password__toggle absolute -translate-y-2/4 cursor-pointer text-[#666] right-5 top-[25px]"
              onClick={() => setShowConfirmPassword(!showConfirmPassword)}
              onMouseDown={(e) => e.preventDefault()}
            >
              {showConfirmPassword ? <LuEye /> : <LuEyeOff />}
            </span>
            {errors.confirmPassword && (
              <div className="flex items-center gap-1 h-5 mt-[5px]">
                <CgDanger className="w-[16.7px] h-[16.7px] bg-[#EF1212] text-white rounded-full" />
                <p className="text-[#EF1212] text-xs">
                  {errors.confirmPassword.message}
                </p>
              </div>
            )}
          </div>
        </div>
      </div>

      <button type="submit" className="btn-primary" disabled={isLoading}>
        {isLoading ? 'Signing up...' : 'Sign Up'}
      </button>
    </form>
  );
};

export default SignupForm;
