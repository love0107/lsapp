import React, { useState } from 'react';

function SignUp() {
  const [formData, setFormData] = useState({
    userName: '',
    fName: '',
    sName: '',
    mobile: '',
    email: '',
    gender: '',
    password: '',
  });
  const [isRegistered, setIsRegistered] = useState(false); 
  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:8080/signup', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response) {
        console.log('User registered successfully!');
        // Optionally, you can redirect the user to another page or show a success message
      } else {
        console.error('Failed to register user');
        // Handle error scenarios
      }
    } catch (error) {
      console.error('Failed to connect to server:', error);
      // Handle network errors
    }
    if (isRegistered) {
      return (
        <div>
          <h2>Registration Successful! + {setIsRegistered}</h2>
          {/* Display additional content for registered users */}
        </div>
      );
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" name="userName" placeholder="Username" value={formData.userName} onChange={handleChange} />
      <input type="text" name="fName" placeholder="First Name" value={formData.fName} onChange={handleChange} />
      <input type="text" name="sName" placeholder="Last Name" value={formData.sName} onChange={handleChange} />
      <input type="text" name="mobile" placeholder="Mobile" value={formData.mobile} onChange={handleChange} />
      <input type="email" name="email" placeholder="Email" value={formData.email} onChange={handleChange} />
      <input type="text" name="gender" placeholder="Gender" value={formData.gender} onChange={handleChange} />
      <input type="password" name="password" placeholder="Password" value={formData.password} onChange={handleChange} />
      <button type="submit">Register</button>
    </form>
  );
}

export default SignUp;

