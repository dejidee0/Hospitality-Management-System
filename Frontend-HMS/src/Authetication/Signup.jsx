import { Link } from "react-router-dom";
import { useState } from "react";
import { LuEye, LuEyeOff } from "react-icons/lu";
import { CgDanger } from "react-icons/cg";
import Navbar from "../components/Navbar";
import sideImg from "../assets/signup-side-image.svg";
import googleIcon from "../assets/google_icon.svg";
import facebookIcon from "../assets/facebook_icon.svg";

const Signup = () => {
  const [showPassword, setShowPassword] = useState(false); //for password toggle
  const currentYear = new Date().getFullYear(); // Dynamically get the current year for copyright section
  // State to hold the form data
  const [formData, setFormData] = useState({
    email: "",
    password: "",
    confirmPassword: "",
  });

  // State to hold error messages for each field
  const [errors, setErrors] = useState({
    email: "",
    password: "",
    confirmPassword: "",
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
      [field]: "", // Clear the error message for the focused field
    });
  };

  // Validation logic for the form fields
  const validate = () => {
    const newErrors = {
      email: "",
      password: "",
      confirmPassword: "",
    };

    // Validate email format using regex
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(formData.email)) {
      newErrors.email = "Email address is invalid";
    }

    // Validate password length (must be at least 8 characters)
    if (formData.password.length < 8) {
      newErrors.password = "Your password must contain 8 or more characters.";
    }

    // Validate that password and confirmPassword match
    if (formData.password !== formData.confirmPassword) {
      newErrors.confirmPassword = "Password doesn't match.";
    }

    // Update the errors state with validation results
    setErrors(newErrors);

    // Return true if there are no errors
    return Object.values(newErrors).every((error) => error === "");
  };

  // Handle form submission
  const handleSubmit = (e) => {
    e.preventDefault(); // Prevent the default form submission behavior

    if (validate()) {
      console.log("Form submitted successfully:", formData);

      // Simulate dummy data submission
      alert("Form submitted with dummy data. API not ready.");

      // Clear the form
      setFormData({
        email: "",
        password: "",
        confirmPassword: "",
      });
      setErrors({
        email: "",
        password: "",
        confirmPassword: "",
      });
    }
  };
  return (
    <>
      <Navbar />
      <div
        className="flex justify-center items-center h-[160vh]  flex-col gap-[52px]"
        style={{ backgroundColor: "var(--auth-bg-color)" }}
      >
        {/* Center Item container*/}
        <div className="bg-white w-[61.87rem] h-[42.125rem] rounded-[1.5rem] flex justify-between">
          <div>
            <img
              src={sideImg}
              alt="Palmtrees and a pool"
              className="h-[42.125rem] w-[22.75rem] rounded-tl-3xl rounded-bl-3xl"
            />
          </div>
          <div className="rounded-[1.5rem] w-[644px] bg-white p-12">
            <form className="flex flex-col gap-8" onSubmit={handleSubmit}>
              {/* FORM HEADING */}
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
              {/* Form content */}
              <div className="flex flex-col gap-[32px]">
                {/* Other means of authetication */}
                <div className="flex gap-[15px]">
                  {/* Sign up with Google */}
                  <div className="bg-[color:var(--default-grey)] h-12 w-[266.5px] flex justify-center items-center gap-2.5 rounded-lg">
                    <img
                      src={googleIcon}
                      alt="Google Icon"
                      className="w-6 h-6"
                    />
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
                      name="email"
                      autoComplete="email"
                      value={formData.email}
                      onChange={handleChange}
                      onFocus={() => handleFocus("email")} // Clear the email error on focus
                      placeholder=""
                      required
                    />

                    <label className="placeholder" htmlFor="email">
                      {" "}
                      Email{" "}
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
                      type={showPassword ? "text" : "password"}
                      placeholder=""
                      required
                      id="password"
                      name="password"
                      value={formData.password}
                      onChange={handleChange}
                      onFocus={() => handleFocus("password")} // Clear the password error on focus
                    />

                    <label className="placeholder" htmlFor="password">
                      {" "}
                      Password{" "}
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
                  </div>

                  <div className="input-group">
                    <input
                      className="input w-[548px] h-12 border [borderRadius:8px] pl-2.5 pr-3 py-3.5 border-solid border-[#dcdcdc]"
                      type={showPassword ? "text" : "password"}
                      placeholder=""
                      required
                      id="confirmPassword"
                      name="confirmPassword"
                      value={formData.confirmPassword}
                      onChange={handleChange}
                      onFocus={() => handleFocus("confirmPassword")} // Clear the confirm password error on focus
                    />
                    <label className="placeholder" htmlFor="confirmPassword">
                      Confirm Password{" "}
                    </label>
                    <span
                      className="password__toggle hidden absolute -translate-y-2/4 cursor-pointer text-[#666] right-5 top-[25px]"
                      onClick={() => setShowPassword(!showPassword)}
                    >
                      {showPassword ? <LuEye /> : <LuEyeOff />}
                    </span>

                    {errors.confirmPassword && (
                      <div className="flex items-center gap-1 h-5 mt-[5px]">
                        <CgDanger className="w-[16.7px] h-[16.7px] bg-[#EF1212] text-white rounded-[999px]" />
                        <p className="text-[#EF1212] text-xs font-normal leading-[16.8px]">
                          {errors.confirmPassword}
                        </p>
                      </div>
                    )}
                  </div>
                  <button className="w-[548px] h-12 border border-[color:var(--primary-purple)] bg-[color:var(--primary-purple)] flex justify-center items-center font-medium [font-style:14px] leading-[19.6px] text-white rounded-lg border-solid">
                    Sign Up
                  </button>
                  <span className="flex justify-center gap-[5px] font-normal text-sm leading-[19.6px] text-center">
                    <p>By signing up you accept our</p>
                    <Link className="text-[color:var(--primary-purple)]">
                      Terms{" "}
                    </Link>{" "}
                    and{" "}
                    <Link className=" text-[color:var(--primary-purple)]">
                      Privacy Policy
                    </Link>
                  </span>
                </div>
              </div>
            </form>
          </div>
        </div>

        <div>
          <p className="font-normal text-base leading-[22.4px] text-center text-[#393b3a]">
            &copy; {currentYear} FindPeace Ltd
          </p>
        </div>
      </div>
    </>
  );
};
export default Signup;
