import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';
import { useLoginMutation } from '@/api/authApi';
import { useDispatch } from 'react-redux';
import { setCredentials } from '@/redux/slices/users/authSlice';
import { CgDanger } from 'react-icons/cg';
import googleIcon from '@/assets/google_icon.svg';
import facebookIcon from '@/assets/facebook_icon.svg';
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
});

const SigninForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: yupResolver(schema),
  });

  const [showPassword, setShowPassword] = useState(false);
  const [loginMutation, { isLoading }] = useLoginMutation();
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const onSubmit = async (data) => {
    console.log(data)
    try {
      const response = await loginMutation(data).unwrap();
      dispatch(setCredentials(response)); // Update auth state
      navigate('/signin'); // Redirect to sign-in page after successful signup
    } catch (err) {
      console.error('Signin error:', err);
    }
  };

  return (
    <>
      <form className="flex flex-col gap-8" onSubmit={handleSubmit(onSubmit)}>
        {/* FORM HEADING */}
        <div className="flex flex-col gap-2">
          {/* to change */}
          <h4 className="font-semibold text-[22px] leading-[26.4px] text-[black]">
            Sign in to account
          </h4>
          <span className="font-normal text-sm leading-[19.6px] text-[black] flex gap-[5px]">
            {/* to change */}
            <p>Don&#39;t have an account up</p>
            <Link
              to="/signup"
              className="text-[color:var(--primary-purple)] underline"
            >
              Sign Up
            </Link>
          </span>
        </div>
        {/* Form content */}
        <div className="flex flex-col gap-[32px]">
          {/* Other means of authetication */}
          <div className="flex gap-[15px]">
            {/* Sign up with Google */}
            <div className="bg-[color:var(--default-grey)] h-12 w-[266.5px] flex justify-center items-center gap-2.5 rounded-lg">
              <img src={googleIcon} alt="Google Icon" className="w-6 h-6" />
              <Link className="text-sm font-normal leading-[19.6px]">
                Google
              </Link>
            </div>
            <div className="bg-[color:var(--default-grey)] h-12 w-[266.5px] flex justify-center items-center gap-2.5 rounded-lg">
              <img
                src={facebookIcon}
                alt="facebook login"
                className="w-6 h-6"
              />
              <Link className="text-sm font-normal leading-[19.6px]">
                Facebook
              </Link>
            </div>
          </div>

          {/* Default signup */}
          <div className="flex items-center justify-center gap-3 relative text-center text-sm font-normal text-black">
            <span className=" w-[200.5px] grow h-px bg-[#E6E7E6] mr-2.5"></span>
            Or Sign in with
            <span className=" w-[200.5px] grow h-px bg-[#E6E7E6] mr-2.5"></span>
          </div>

          <div className="flex flex-col gap-5">
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
                  <p className="text-[#EF1212] text-xs">
                    {errors.email.message}
                  </p>
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
          </div>
          <div>
            <div>
              <Link
                to="/forget-password"
                style={{
                  color: 'var(--primary-purple)',
                  float: 'right',
                }}
              >
                Forget Password?
              </Link>
            </div>

            <button className="w-[548px] h-12 border border-[color:var(--primary-purple)] bg-[color:var(--primary-purple)] flex justify-center items-center font-medium [font-style:14px] leading-[19.6px] text-white rounded-lg border-solid">
              {isLoading ? 'Signing in...' : 'Sign In'}
            </button>
            <span className="flex justify-center gap-[5px] font-normal text-sm leading-[19.6px] text-center">
              <p>By signing up you accept our</p>
              <Link className="text-[color:var(--primary-purple)]">
                Terms{' '}
              </Link>{' '}
              and{' '}
              <Link className=" text-[color:var(--primary-purple)]">
                Privacy Policy
              </Link>
            </span>
          </div>
        </div>
      </form>
    </>
  );
};

export default SigninForm;
