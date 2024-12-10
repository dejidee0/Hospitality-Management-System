import Navbar from "../components/Navbar";
import checkMark from "../../assets/confirmPasswordCheck.svg";

const ConfirmChange = () => {
  const currentYear = new Date().getFullYear(); // Dynamically get the current year for copyright section
  return (
    <>
      <>
        <Navbar />
        <div className="p-5 pt-0 md:bg-auth-bg-color">
          <div className="bg-white md:bg-auth-bg-color flex items-center justify-center min-h-screen">
            {/* Center Item container*/}
            <div>
              <div className="flex flex-col items-center text-center justify-center rounded-3xl md:bg-white gap-6 md:w-[61.88rem] md:h-[440px]">
                {/* Image holder */}
                <div>
                  <img
                    src={checkMark}
                    alt="confirm password check"
                    className="w-20 h-20 md:w-32 md:h-32"
                  />
                </div>

                {/* Text holder */}
                <div className="flex flex-col gap-3">
                  <h3 className="leading-6 font-medium text-xl md:font-semibold md:leading-7 md:text-2xl">
                    New password set successfully
                  </h3>
                  <p className="font-normal text-sm leading-5">
                    Your new password has been changed successfully
                  </p>
                  <button className="bg-primary-purple font-medium text-sm leading-5 text-white px-6 py-3 rounded-lg">
                    Continue
                  </button>
                </div>
              </div>
            </div>
          </div>
          <footer>
            <p className="font-normal text-base leading-[22.4px] text-center text-[#393b3a]">
              &copy; {currentYear} FindPeace Ltd
            </p>
          </footer>
        </div>
      </>
    </>
  );
};

export default ConfirmChange;
