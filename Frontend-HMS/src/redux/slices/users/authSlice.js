import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';

// Set API base URL
const API_URL = `${import.meta.env.VITE_API_URL}/v1/auth`; // Replace with your API endpoint

// Async Thunks
export const signup = createAsyncThunk(
    'auth/signup',
    async (userDetails, { rejectWithValue }) => {
      try {
        const response = await axios.post(`${API_URL}/signup`, userDetails);
        const { message } = response.data;
        return { message };
      } catch (error) {
        return rejectWithValue(
          error.response?.data?.message || 'Signup failed. Please try again.'
        );
      }
    }
  );

// Async Thunks
export const login = createAsyncThunk(
  'auth/login',
  async (credentials, { rejectWithValue }) => {
    try {
      const response = await axios.post(`${API_URL}/login`, credentials);
      const { token } = response.data;
      return { token };
    } catch (error) {
        return rejectWithValue(
            error.response?.data?.message || 'Login failed. Password or Email is incorrect'
        );
    }
  }
);

export const logout = createAsyncThunk('auth/logout', async (_, { dispatch }) => {
  localStorage.removeItem('auth');
  dispatch(authSlice.actions.logoutSuccess());
});

export const checkAuth = createAsyncThunk('auth/checkAuth', async () => {
  const token = JSON.parse(localStorage.getItem('auth'));
  if (token) {
    return token; // Returns the token and user
  }
  else 
    return null
});

// Slice
const authSlice = createSlice({
  name: 'auth',
  initialState: {
    token: null,
    isAuthenticated: false,
    isLoading: false,
    error: null,
  },
  reducers: {
    logoutSuccess(state) {
      state.token = null;
      state.isAuthenticated = false;
    },
  },
  extraReducers: (builder) => {
    builder
    .addCase(signup.pending, (state) => {
        state.isLoading = true;
        state.error = null;
      })
      .addCase(signup.fulfilled, (state, action) => {
        state.isLoading = false;
        state.isAuthenticated = true;
        state.token = action.payload.token;
        localStorage.setItem('auth', JSON.stringify(action.payload));
      })
      .addCase(signup.rejected, (state, action) => {
        state.isLoading = false;
        state.error = action.payload;
      })
      .addCase(login.pending, (state) => {
        state.isLoading = true;
        state.error = null;
      })
      .addCase(login.fulfilled, (state, action) => {
        state.isLoading = false;
        state.isAuthenticated = true;
        state.token = action.payload.token;
        localStorage.setItem('auth', JSON.stringify(action.payload.token));
      })
      .addCase(login.rejected, (state, action) => {
        state.isLoading = false;
        state.error = action.payload;
      })
      .addCase(checkAuth.fulfilled, (state, action) => {
        state.isAuthenticated = true;
        state.token = action.payload;
      })
      .addCase(checkAuth.rejected, (state) => {
        state.isAuthenticated = false;
      });
  },
});

export const { logoutSuccess } = authSlice.actions;
export default authSlice.reducer;
