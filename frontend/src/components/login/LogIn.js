
import { useState } from "react";
const initialValues = {
  email: "",
  password: "",
};
const Login = () => {

  const [inputValues, setValues] = useState(initialValues);

  const handleInputChange = (e) => {
    //const name = e.target.name 
    //const value = e.target.value 
    const { name, value } = e.target;

    setValues({
      ...inputValues,
      [name]: value,
    });
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
  
    const response = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(inputValues)
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
      <h2>Please LogIn!</h2>
      <form onSubmit={handleSubmit}>
      <label>Email:
      <input 
        type="email" 
        name="email" 
        value={inputValues.email} 
        onChange={handleInputChange}
      />
      </label>
      <label>Password:
      <input 
        type="password" 
        name="password" 
        value={inputValues.password} 
        onChange={handleInputChange}
      />
      </label>
      <input type="submit" />
      </form>
    </div>
  );
};

export default Login;
