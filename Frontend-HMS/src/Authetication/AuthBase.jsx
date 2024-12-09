// most of the pages look alike so this is the base component for all the pages
import Navbar from "../components/Navbar";
import sideImg from "../assets/signup-side-image.svg";
import PropTypes from "prop-types";

const AuthBase = ({ children }) => {
  const currentYear = new Date().getFullYear(); // Dynamically get the current year for copyright section
  return (
    <>
      <>
        <Navbar />
        <div
          className="flex justify-center items-center h-[auto]  flex-col gap-[52px]"
          style={{ backgroundColor: "var(--auth-bg-color)" }}
        >
          {/* Center Item container*/}
          <div className="bg-white w-[61.87rem] h-[42.125rem] rounded-[1.5rem] flex justify-between mt-40 mb-10">
            <div className="hidden sm:hidden lg:block">
              <img
                src={sideImg}
                alt="Palmtrees and a pool"
                className="h-[42.125rem] w-[22.75rem] rounded-tl-3xl rounded-bl-3xl"
              />
            </div>
            <div className="rounded-[1.5rem] lg:w-[40.25rem] bg-white p-12 w-80">
              {children}
            </div>
          </div>

          <div>
            <p className="font-normal text-base leading-[22.4px] text-center text-[#393b3a]">
              &copy; {currentYear} FindPeace Ltd
            </p>
          </div>
        </div>
      </>
    </>
  );
};

AuthBase.propTypes = {
  children: PropTypes.node.isRequired,
};

export default AuthBase;
