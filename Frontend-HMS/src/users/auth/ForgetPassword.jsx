import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import AuthBase from './AuthBase';
import Loading from './components/ButtonLoader';
const ForgetPassword = () => {
  const [email, setEmail] = useState('');
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    try {
      const response = await fetch(
        `http://localhost:8080/v1/auth/reset-password?email=${email}`,
        {
          method: 'GET',
        }
      );

      if (response.ok) {
        setMessage('Check your email for the reset link.');
        // Redirect to reset link page with email passed as state
        // Save email to localStorage
        localStorage.setItem('resetEmail', email);
        navigate('/reset-link');
      } else {
        const errorData = await response.json();
        console.log('Error data:', errorData);
        setMessage('Error: Unable to process your request.');
      }
    } catch (error) {
      console.error('Error:', error.message);
      setMessage('Something went wrong. Please try again later.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <AuthBase>
        <form className="flex flex-col gap-8" onSubmit={handleSubmit}>
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
                  onChange={(e) => setEmail(e.target.value)}
                />

                <label className="placeholder" htmlFor="email">
                  {' '}
                  Email{' '}
                </label>
              </div>

              <button
                type="submit"
                className="w-[548px] h-12 border border-primary-purple bg-primary-purple flex justify-center items-center font-medium [font-style:14px] leading-[19.6px] text-white rounded-lg border-solid"
              >
                {loading ? <Loading /> : ' Reset password'}
              </button>
              {message && <p>{message}</p>}
            </div>
          </div>
        </form>
      </AuthBase>
    </>
  );
};
export default ForgetPassword;
