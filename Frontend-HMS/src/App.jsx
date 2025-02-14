import { Routes, Route } from 'react-router-dom';
import './App.css';
import UsersRoutes from '@/routes/users/userRoutes';


function App() {
  return (
    <>
      <Routes>
        <Route path="/*" element={<UsersRoutes />} />
      </Routes>    
    </>
  );
}

export default App;
