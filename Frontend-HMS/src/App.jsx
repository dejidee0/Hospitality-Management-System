import { Routes } from 'react-router-dom';
import './App.css';
import UsersRoutes from './routes/users/userRoutes';

function App() {
  return (
    <>
      <Routes>
        <UsersRoutes />
      </Routes>
    </>
  );
}

export default App;
