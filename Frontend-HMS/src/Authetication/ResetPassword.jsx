import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import AuthBase from "./AuthBase";

const ResetPassword = () => {
  const [email, setEmail] = useState("");

  useEffect(() => {
    const storedEmail = localStorage.getItem("resetEmail");
    if (storedEmail) {
      setEmail(storedEmail);
    } else {
      setEmail("No email found."); // Fallback if email isn't available
    }
  }, []);
  return (
    <AuthBase>
      <div className="flex gap-3 flex-col">
        <h4 className="text-[22px] font-semibold leading-[26.4px]">
          Reset Link Set
        </h4>
        <span>
          The link to reset your password has been sent to{" "}
          <Link href={`mail:${email}`}>{email} </Link>
        </span>
        <p>
          Just in case you still don&#39;t get our email, please check your spam
          folder.
        </p>
      </div>
    </AuthBase>
  );
};
export default ResetPassword;
