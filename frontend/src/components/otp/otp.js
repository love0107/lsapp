
import React, { useState } from 'react';

const Otp = () => {


const [otp, setOtp] = useState(Array(6).fill('')); // Initialize state


const handleChange = (element, index) => {
  if (isNaN(element.value)) return false;
  setOtp([...otp.map((d, idx) => (idx === index ? element.value : d))]);
  // focus on next input
  if (element.nextSibling) {
    element.nextSibling.focus();
  }
};

const handleSubmit = async (event) => {
    event.preventDefault();
    const otpValue = otp.join('');
    const response = await fetch('http://localhost:8080/password/otp/validate', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ otp: otpValue }),
      credentials: 'include'
    });
  
    if (response.ok) {
      console.log('Response succeeded!');
      const data = await response.json();
      console.log(data);
    } else {
      console.log('Response failed!', response.status);
    }
  };

return (
  <div>
    <h2>Please Enter OTP!</h2>
    <form onSubmit={handleSubmit}>
      <label>Otp:
        {[...Array(6)].map((_, i) => (
          <input
            type="text"
            name="otp"
            value={otp[i]}
            onChange={e => handleChange(e.target, i)}
            maxLength="1"
            key={i}
          />
        ))}
      </label>
      <input type="submit" />
    </form>
  </div>
);
}
export default Otp;