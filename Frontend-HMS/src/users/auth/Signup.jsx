import { Link, useNavigate } from "react-router-dom";
import { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { signup } from "../../redux/authSlice"; // Adjust the path as needed
import AuthBase from "./AuthBase";
import { LuEye, LuEyeOff } from "react-icons/lu";
import { CgDanger } from "react-icons/cg";
import googleIcon from "../../assets/google_icon.svg";
import facebookIcon from "../../assets/facebook_icon.svg";

const Signup = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const { isLoading, error } = useSelector((state) => state.auth);

  const [showPassword, setShowPassword] = useState(false);
  const [formData, setFormData] = useState({
    email: "",
    password: "",
    confirmPassword: "",
  });

  const [errors, setErrors] = useState({
    email: "",
    password: "",
    confirmPassword: "",
  });

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleFocus = (field) => {
    setErrors({
      ...errors,
      [field]: "",
    });
  };

  const validate = () => {
    const newErrors = {
      email: "",
      password: "",
      confirmPassword: "",
    };

    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(formData.email)) {
      newErrors.email = "Email address is invalid";
    }

    if (formData.password.length < 8) {
      newErrors.password = "Your password must contain 8 or more characters.";
    }

    if (formData.password !== formData.confirmPassword) {
      newErrors.confirmPassword = "Passwords don't match.";
    }

    setErrors(newErrors);
    return Object.values(newErrors).every((error) => error === "");
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (validate()) {
      dispatch(signup(formData))
        .unwrap()
        .then(() => {
          navigate("/signin");
        })
        .catch((err) => {
          console.error("Signup error:", err);
        });
    }
  };

  return (
    <AuthBase>
      <form className="flex flex-col gap-8" onSubmit={handleSubmit}>
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

          <div className="flex items-center justify-center gap-3 relative text-center text-sm font-normal text-black">
            <span className="w-[200.5px] grow h-px bg-[#E6E7E6] mr-2.5"></span>
            Or Sign up with
            <span className="w-[200.5px] grow h-px bg-[#E6E7E6] mr-2.5"></span>
          </div>

          <div className="flex flex-col gap-5">
            <div className="input-group">
              <input
                className="input w-[548px] h-12 border rounded pl-2.5 pr-3 py-3.5 border-solid border-[#dcdcdc]"
                type="text"
                id="email"
                name="email"
                value={formData.email}
                onChange={handleChange}
                onFocus={() => handleFocus("email")}
                placeholder="Email"
                required
              />
              {errors.email && (
                <div className="flex items-center gap-1 h-5 mt-[5px]">
                  <CgDanger className="w-[16.7px] h-[16.7px] bg-[#EF1212] text-white rounded-full" />
                  <p className="text-[#EF1212] text-xs">{errors.email}</p>
                </div>
              )}
            </div>

            <div className="input-group">
              <input
                className="input w-[548px] h-12 border rounded pl-2.5 pr-3 py-3.5 border-solid border-[#dcdcdc]"
                type={showPassword ? "text" : "password"}
                id="password"
                name="password"
                value={formData.password}
                onChange={handleChange}
                onFocus={() => handleFocus("password")}
                placeholder="Password"
                required
              />
              <span onClick={() => setShowPassword(!showPassword)}>
                {showPassword ? <LuEyeOff /> : <LuEye />}
              </span>
              {errors.password && (
                <div className="flex items-center gap-1 h-5 mt-[5px]">
                  <CgDanger className="w-[16.7px] h-[16.7px] bg-[#EF1212] text-white rounded-full" />
                  <p className="text-[#EF1212] text-xs">{errors.password}</p>
                </div>
              )}
            </div>

            <div className="input-group">
              <input
                className="input w-[548px] h-12 border rounded pl-2.5 pr-3 py-3.5 border-solid border-[#dcdcdc]"
                type="password"
                id="confirmPassword"
                name="confirmPassword"
                value={formData.confirmPassword}
                onChange={handleChange}
                onFocus={() => handleFocus("confirmPassword")}
                placeholder="Confirm Password"
                required
              />
              {errors.confirmPassword && (
                <div className="flex items-center gap-1 h-5 mt-[5px]">
                  <CgDanger className="w-[16.7px] h-[16.7px] bg-[#EF1212] text-white rounded-full" />
                  <p className="text-[#EF1212] text-xs">{errors.confirmPassword}</p>
                </div>
              )}
            </div>
          </div>
        </div>

        {error && (
          <div className="text-[#EF1212] text-xs text-center">{error}</div>
        )}

        <button
          type="submit"
          className="btn-primary"
          disabled={isLoading}
        >
          {isLoading ? "Signing up..." : "Sign Up"}
        </button>
      </form>
    </AuthBase>
  );
};

export default Signup;
