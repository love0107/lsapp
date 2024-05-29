import React from 'react';
import Signup from './components/signup/SignUp';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './components/login/LogIn';
import Otp from './components/otp/otp';
import Password from './components/password/password';




function App() {
  return (
        <div className="App">
    <Router>
  <Routes>
    <Route path="/login" element={<Login />} />
    <Route path="/signup" element={<Signup />} />
    <Route path="/otp" element={<Otp />} />
    <Route path="/password" element={<Password />} />
  </Routes>
</Router>
   </div>
  );
}

export default App;
