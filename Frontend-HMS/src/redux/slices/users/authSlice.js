import { createSlice } from '@reduxjs/toolkit';
import { saveToken, removeToken, getToken } from '../../../utils/authUtils';


// Initial state
const initialState = {
  user: null,
  token: getToken(), // Retrieve token from localStorage
  isAuthenticated: !!getToken(), // Check if token exists
};

// Create the slice
const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    // Set user and token after successful login/register
    setCredentials: (state, action) => {
      const { user, token } = action.payload;
      state.user = user;
      state.token = token;
      state.isAuthenticated = true;
      saveToken(token); // Save token to localStorage
    },
    // Clear user and token on logout
    logout: (state) => {
      state.user = null;
      state.token = null;
      state.isAuthenticated = false;
      removeToken(); // Remove token from localStorage
    },
  },
});

// Export actions
export const { setCredentials, logout } = authSlice.actions;

// Export reducer
export default authSlice.reducer;
