// Save token to localStorage
export const saveToken = (token) => localStorage.setItem('authToken', token);

// Remove token from localStorage
export const removeToken = () => localStorage.removeItem('authToken');

// Get token from localStorage
export const getToken = () => localStorage.getItem('authToken');
