import { useState } from "react";

const initialValues = {
  fName: "",
  sName: "",
  mobile: "",
  email: "",
  gender: "male",
  password: "",
};

const SignUp = () => {
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
  
    const response = await fetch('http://localhost:8080/signup', {
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
    <form onSubmit={handleSubmit}>
      <label>First Name:
      <input 
        type="text" 
        name="fName" 
        value={inputValues.fName} 
        onChange={handleInputChange}
      />
      </label>

      <label>Second Name:
      <input 
        type="text" 
        name="sName" 
        value={inputValues.sName} 
        onChange={handleInputChange}
      />
      </label>

      <label>Mobile:
      <input 
        type="text" 
        name="mobile" 
        value={inputValues.mobile} 
        onChange={handleInputChange}
      />
      </label>

      <label>Email:
      <input 
        type="email" 
        name="email" 
        value={inputValues.email} 
        onChange={handleInputChange}
      />
      </label>

      <label>Gender:
      <select name="gender" value={inputValues.gender} onChange={handleInputChange}>
    <option value="male">Male</option>
    <option value="female">Female</option>
    <option value="other">Other</option>
  </select>
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

  )
};

export default SignUp;
