import { useDispatch, useSelector } from 'react-redux';
import { login } from '../../redux/slices/users/authSlice';
import { Link, useNavigate } from 'react-router-dom';
import { useState } from 'react';
import AuthBase from './AuthBase';
import Loading from './components/ButtonLoader';
import { LuEye, LuEyeOff } from 'react-icons/lu';
import { CgDanger } from 'react-icons/cg';
import googleIcon from '../../assets/google_icon.svg';
import facebookIcon from '../../assets/facebook_icon.svg';

const Signin = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [showPassword, setShowPassword] = useState(false);
  const { isLoading, error } = useSelector((state) => state.auth);
  // State to hold the form data
  const [formData, setFormData] = useState({
    email: '',
    password: '',
  });

  // State to hold error messages for each field
  const [errors, setErrors] = useState({
    email: '',
    password: '',
  });

  // Handle input changes and update formData state
  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value, // Update the value of the field dynamically
    });
  };

  // Handle focus event to clear the specific error message
  const handleFocus = (field) => {
    setErrors({
      ...errors,
      [field]: '', // Clear the error message for the focused field
    });
  };

  // Validation logic for the form fields
  const validate = () => {
    const newErrors = {
      email: '',
      password: '',
    };

    // Validate email format using regex
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(formData.email)) {
      newErrors.email = 'Email address is invalid';
    }

    // Validate password length (must be at least 8 characters)
    if (formData.password.length < 8) {
      newErrors.password = 'Your password must contain 8 or more characters.';
    }

    // Update the errors state with validation results
    setErrors(newErrors);

    // Return true if there are no errors
    return Object.values(newErrors).every((error) => error === '');
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Handle form submission
    if (validate()) {
      dispatch(login(formData))
        .unwrap()
        .then(() => {
          navigate('/home');
        })
        .catch((err) => {
          console.error('Signup error:', err);
        });
    }
  };

  return (
    <>
      <AuthBase>
        <form className="flex flex-col gap-8" onSubmit={handleSubmit}>
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
              Or Sign up with
              <span className=" w-[200.5px] grow h-px bg-[#E6E7E6] mr-2.5"></span>
            </div>

            <div className="flex flex-col gap-5">
              <div className="input-group">
                <input
                  className="input  w-[548px] h-12 border [borderRadius:8px] pl-2.5 pr-3 py-3.5 border-solid border-[#dcdcdc]"
                  type="text"
                  id="email"
                  required
                  name="email"
                  autoComplete="email"
                  value={formData.email}
                  onChange={handleChange}
                  onFocus={() => handleFocus('email')} // Clear the email error on focus
                  placeholder=""
                />

                <label className="placeholder" htmlFor="email">
                  {' '}
                  Email{' '}
                </label>

                {/* Display email error message */}
                {errors.email && (
                  <div className="flex items-center gap-1 h-5 mt-[5px]">
                    <CgDanger className="w-[16.7px] h-[16.7px] bg-[#EF1212] text-white rounded-[999px]" />
                    <p className="text-[#EF1212] text-xs font-normal leading-[16.8px]">
                      {errors.email}
                    </p>
                  </div>
                )}
              </div>

              <div className="input-group">
                <input
                  className="input w-[548px] h-12 border [borderRadius:8px] pl-2.5 pr-3 py-3.5 border-solid border-[#dcdcdc]"
                  type={showPassword ? 'text' : 'password'}
                  placeholder=""
                  required
                  id="password"
                  name="password"
                  value={formData.password}
                  onChange={handleChange}
                  onFocus={() => handleFocus('password')} // Clear the password error on focus
                />

                <label className="placeholder" htmlFor="password">
                  {' '}
                  Password{' '}
                </label>

                <span
                  className="password__toggle hidden absolute -translate-y-2/4 cursor-pointer text-[#666] right-5 top-[25px]"
                  onClick={() => setShowPassword(!showPassword)}
                  onMouseDown={(e) => e.preventDefault()}
                >
                  {showPassword ? <LuEye /> : <LuEyeOff />}
                </span>

                {errors.password && (
                  <div className="flex items-center gap-1 h-5 mt-[5px]">
                    <CgDanger className="w-[16.7px] h-[16.7px] bg-[#EF1212] text-white rounded-[999px]" />
                    <p className="text-[#EF1212] text-xs font-normal leading-[16.8px]">
                      {errors.password}
                    </p>
                  </div>
                )}
                {error && (
                  <div className="text-[#EF1212] pt-2 text-xs text-center">
                    {error}
                  </div>
                )}
              </div>

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
                {isLoading ? <Loading /> : 'Sign In'}
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
      </AuthBase>
    </>
  );
};
export default Signin;
