import React, { useState } from 'react';
import Login from './components/LogIn';
import SignUp from './components/SignUp';
import './App.css';
function App() {
  const [user, setUser] = useState(null);

  const handleSignIn = (username) => {
    setUser(username);
  };

  const handleLogout = () => {
    setUser(null);
  };

  return (
    <div className="App">
      {!user ? (
        <SignUp handleSignIn={handleSignIn} />
      ) : (
        <Login username={user} handleLogout={handleLogout} />
      )}
    </div>
  );
}

export default App;
