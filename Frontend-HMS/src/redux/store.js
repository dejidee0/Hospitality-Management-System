import { configureStore } from '@reduxjs/toolkit';
import authReducer from './slices/users/authSlice';
import { authApi } from '../api/authApi';


const store = configureStore({
  reducer: {
    [authApi.reducerPath]: authApi.reducer,
    auth: authReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(authApi.middleware),
});

export default store;
