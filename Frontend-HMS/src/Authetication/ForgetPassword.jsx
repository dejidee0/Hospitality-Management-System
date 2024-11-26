// import { Link } from "react-router-dom";
// import { CgDanger } from "react-icons/cg";
import Navbar from "../components/Navbar";
import sideImg from "../assets/signup-side-image.svg";

const ForgetPassword = () => {
  const currentYear = new Date().getFullYear(); // Dynamically get the current year for copyright section
  // State to hold the form data
  return (
    <>
      <Navbar />
      <div
        className="flex justify-center items-center h-[140vh]  flex-col gap-[52px]"
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
            <form className="flex flex-col gap-8">
              {/* FORM HEADING */}
              <div className="flex flex-col gap-2">
                <h4 className="font-semibold text-[22px] leading-[26.4px] text-[black]">
                  Forget Password
                </h4>
                <span className="font-normal text-sm leading-[19.6px] text-[black] flex gap-[5px]">
                  <p>
                    You can easily reset your password here. We&#39;ll send the
                    reset link to your email.
                  </p>
                </span>
              </div>
              {/* Form content */}
              <div className="flex flex-col gap-[32px]">
                {/* Default signup */}

                <div className="flex flex-col gap-5">
                  <div className="input-group">
                    <input
                      className="input  w-[548px] h-12 border [borderRadius:8px] pl-2.5 pr-3 py-3.5 border-solid border-[#dcdcdc]"
                      type="text"
                      id="email"
                      required
                      name="email"
                      autoComplete="email"
                      placeholder=""
                    />

                    <label className="placeholder" htmlFor="email">
                      {" "}
                      Email{" "}
                    </label>
                  </div>

                  <button className="w-[548px] h-12 border border-[color:var(--primary-purple)] bg-[color:var(--primary-purple)] flex justify-center items-center font-medium [font-style:14px] leading-[19.6px] text-white rounded-lg border-solid">
                    Reset password
                  </button>
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
export default ForgetPassword;
