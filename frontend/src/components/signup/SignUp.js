import { useState } from "react";

const initialValues = {
  fName: "",
  sName: "",
  mobile: "",
  email: "",
  gender: "male",
  password: "",
};

function InputField({ label, type, name, value, onChange }) {
  return (
    <label>
      {label}:
      <input type={type} name={name} value={value} onChange={onChange} />
    </label>
  );
}

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

    const response = await fetch("http://localhost:8080/signup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(inputValues),
    });

    if (response.ok) {
      console.log("Response succeeded!");
      const data = await response.json();
      console.log(data);
    } else {
      console.log("Response failed!", response.status);
    }
  };

  return (
    <div className="signup-container">
      <h1 className="signup-heading">Sign Up</h1>
      <form onSubmit={handleSubmit} className="signup-form">
        <InputField
          label="First Name"
          type="text"
          name="fName"
          value={inputValues.fName}
          onChange={handleInputChange}
        />
        <InputField
          label="Second Name"
          type="text"
          name="sName"
          value={inputValues.sName}
          onChange={handleInputChange}
        />
        <InputField
          label="Mobile"
          type="text"
          name="mobile"
          value={inputValues.mobile}
          onChange={handleInputChange}
        />
        <InputField
          label="Email"
          type="email"
          name="email"
          value={inputValues.email}
          onChange={handleInputChange}
        />
        <InputField
          label="Password"
          type="password"
          name="password"
          value={inputValues.password}
          onChange={handleInputChange}
        />

        <label>
          Gender:
          <select
            name="gender"
            value={inputValues.gender}
            onChange={handleInputChange}
          >
            <option value="male">Male</option>
            <option value="female">Female</option>
            <option value="other">Other</option>
          </select>
        </label>
        <input type="submit" className="signup-button" />
      </form>
    </div>
  );
};

export default SignUp;
