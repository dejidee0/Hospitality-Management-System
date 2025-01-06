import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';

// Set API base URL
const API_URL = 'https://api.example.com/auth'; // Replace with your API endpoint

// Async Thunks
export const login = createAsyncThunk(
  'auth/login',
  async (credentials, { rejectWithValue }) => {
    try {
      const response = await axios.post(`${API_URL}/login`, credentials);
      const { token, user } = response.data;
      return { token, user };
    } catch (error) {
      return rejectWithValue(error.response.data || 'Failed to log in');
    }
  }
);

export const logout = createAsyncThunk('auth/logout', async (_, { dispatch }) => {
  localStorage.removeItem('auth');
  dispatch(authSlice.actions.logoutSuccess());
});

export const checkAuth = createAsyncThunk('auth/checkAuth', async () => {
  const auth = JSON.parse(localStorage.getItem('auth'));
  if (auth && auth.token) {
    return auth; // Returns the token and user
  }
  throw new Error('Not authenticated');
});

// Slice
const authSlice = createSlice({
  name: 'auth',
  initialState: {
    user: null,
    token: null,
    isAuthenticated: false,
    isLoading: false,
    error: null,
  },
  reducers: {
    logoutSuccess(state) {
      state.user = null;
      state.token = null;
      state.isAuthenticated = false;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(login.pending, (state) => {
        state.isLoading = true;
        state.error = null;
      })
      .addCase(login.fulfilled, (state, action) => {
        state.isLoading = false;
        state.isAuthenticated = true;
        state.user = action.payload.user;
        state.token = action.payload.token;
        localStorage.setItem('auth', JSON.stringify(action.payload));
      })
      .addCase(login.rejected, (state, action) => {
        state.isLoading = false;
        state.error = action.payload;
      })
      .addCase(checkAuth.fulfilled, (state, action) => {
        state.isAuthenticated = true;
        state.user = action.payload.user;
        state.token = action.payload.token;
      })
      .addCase(checkAuth.rejected, (state) => {
        state.isAuthenticated = false;
      });
  },
});

export const { logoutSuccess } = authSlice.actions;
export default authSlice.reducer;
