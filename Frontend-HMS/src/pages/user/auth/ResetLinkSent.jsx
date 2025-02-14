import { useState, useEffect } from 'react';
import AuthBase from '@components/organisms/AuthBase';

const ResetLinkSent = () => {
  const [email, setEmail] = useState('');

  useEffect(() => {
    // Retrieve email from localStorage
    const storedEmail = localStorage.getItem('resetEmail');
    if (storedEmail) {
      setEmail(storedEmail);
    }
  }, []);

  return (
    <AuthBase>
      <div className="flex gap-3 flex-col text-center">
        <h4 className="text-[22px] font-semibold leading-[26.4px]">
          Reset Link Sent
        </h4>
        <p className="text-black leading-6 font-normal text-base">
          The link to reset your password has been sent to {' '}
          <a
            href="mailto:"
            className="text-primary-purple font-medium underline"
          >
            {email ? <p className="font-normal">{email}</p> :
            'Unknown email'}
          </a>
        </p>

        <p className="text-black leading-6 font-normal text-base">
          If you don't receive the email, please check your spam folder.
        </p>
      </div>
    </AuthBase>
  );
};

export default ResetLinkSent;
