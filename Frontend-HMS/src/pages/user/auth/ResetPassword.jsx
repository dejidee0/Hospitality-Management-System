// import { useState, useEffect } from 'react';
// import AuthBase from './AuthBase';

// const ResetPassword = () => {
//   const [email, setEmail] = useState('');
//   const [emailProviderUrl, setEmailProviderUrl] = useState('');

//   useEffect(() => {
//     // Retrieve the email from localStorage
//     const storedEmail = localStorage.getItem('resetEmail');
//     if (storedEmail) {
//       setEmail(storedEmail);

//       // Extract the domain and map it to the email provider URL
//       const domain = storedEmail.split('@')[1];
//       const url = getEmailProviderUrl(domain);
//       setEmailProviderUrl(url);
//     } else {
//       setEmail('No email found.'); // Fallback if email isn't available
//     }
//   }, []);

//   // Function to map email domains to provider inbox URLs
//   const getEmailProviderUrl = (domain) => {
//     const emailProviders = {
//       'gmail.com': 'https://mail.google.com',
//       'yahoo.com': 'https://mail.yahoo.com',
//       'outlook.com': 'https://outlook.live.com',
//       'hotmail.com': 'https://outlook.live.com',
//       'icloud.com': 'https://www.icloud.com/mail',
//     };

//     // Return URL for known providers, otherwise fallback to domain-based URL
//     return emailProviders[domain] || `https://${domain}`;
//   };

//   return (
//     <AuthBase>
//       <div className="flex gap-3 flex-col">
//         <h4 className="text-[22px] font-semibold leading-[26.4px]">
//           Reset Link Sent
//         </h4>
//         <span className="text-black leading-6 font-normal text-base">
//           The link to reset your password has been sent to{' '}
//           {emailProviderUrl && (
//             <a
//               href={emailProviderUrl}
//               target="_blank"
//               rel="noopener noreferrer"
//               className="text-primary-purple font-medium"
//             >
//               {email}
//             </a>
//           )}
//         </span>
//         <p className="text-black leading-6 font-normal text-base">
//           Just in case you still don&#39;t get our email, please check your spam
//           folder.
//         </p>
//       </div>
//     </AuthBase>
//   );
// };

// export default ResetPassword;
