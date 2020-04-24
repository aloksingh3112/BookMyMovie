import React, { useState } from "react";
import { useForm } from "react-hook-form";
import axios from "axios";
import { SIGN_UP } from "../../config/url";
const Signup = () => {
  const [state, setstate] = useState({});
  const { register, handleSubmit, errors, reset } = useForm();

  const onSubmit = (data, e) => {
    axios
      .post(SIGN_UP, data)
      .then((responseData) => {
        setstate(responseData.data);
        if (responseData.data.statusCode < 300) {
          e.target.reset();
        }
        setTimeout(() => {
          setstate({});
        }, 1000);
      })
      .catch((error) => {
        console.log(error);
      });
  };

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
          ref={register({ required: true, minLength: 6 })}
        />
        {errors.password && errors.password.type === "required" && (
          <p>Password is required</p>
        )}
        {errors.password && errors.password.type === "minLength" && (
          <p>Password is of minimum 6 digit </p>
        )}
      </div>

      <div className="form-group">
        <label>Email</label>
        <input
          type="email"
          className="form-control"
          name="email"
          ref={register({ required: true })}
        />
        {errors.email && <p>Email is required</p>}
      </div>

      <div className="form-group">
        <label>Name</label>
        <input
          type="text"
          className="form-control"
          name="name"
          ref={register({ required: true })}
        />
        {errors.name && <p>Name is required</p>}
      </div>
      <button className="btn btn-primary">Submit</button>
      <button className="btn btn-secondary ml-2" type="button">
        Cancel
      </button>

      {state && state.statusCode > 300 && (
        <div className="alert alert-error mt-2">{state.message}</div>
      )}
      {state && state.statusCode < 300 && (
        <div className="alert alert-success mt-2">{state.message}</div>
      )}
    </form>
  );
};

export default Signup;
