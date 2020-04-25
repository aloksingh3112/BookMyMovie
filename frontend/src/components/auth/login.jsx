import React from "react";
import { useForm } from "react-hook-form";
import Axios from "axios";
import { LOGIN } from "../../config/url";
import { useState } from "react";
import { options } from "../../config/url";

const Login = (props) => {
  const [state, setState] = useState({});
  const onSubmit = (data, e) => {
    console.log(data);
    Axios.post(LOGIN, data, options)
      .then((responseData) => {
        setState(responseData.data);
        console.log(responseData.data);
        if (responseData.data.statusCode < 300) {
          localStorage.setItem("token", responseData.data.data.token);
          props.history.push("/");
        } else {
          setTimeout(() => {
            setState({});
          }, 1000);
        }
      })
      .catch((err) => console.log(err));
  };

  const { register, handleSubmit, errors } = useForm();
  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div className="form-group">
        <label>Username</label>
        <input
          type="text"
          className="form-control"
          name="username"
          ref={register({ required: true })}
        />
        {errors.username && <p>Username is required</p>}
      </div>

      <div className="form-group">
        <label>Password</label>
        <input
          type="password"
          className="form-control"
          name="password"
          ref={register({ required: true })}
        />
        {errors.password && errors.password.type === "required" && (
          <p>Password is required</p>
        )}
      </div>

      <button className="btn btn-primary">Login</button>
      <button className="btn btn-secondary ml-2" type="button">
        Cancel
      </button>
      {state && state.statusCode > 300 && (
        <div className="alert alert-warning mt-2">{state.message}</div>
      )}
    </form>
  );
};

export default Login;
