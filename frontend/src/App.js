import React from 'react';
import SignUp from './components/signup/SignUp';
import Login from './components/login/LogIn';
import Otp from './components/otp/otp';
import Password from './components/password/password';

function App() {
  return (
    <div className="App">
      <h1>User Registration</h1>
      <SignUp />
      <Login />
      <Otp />
      <Password />
    </div>
  );
}

export default App;
