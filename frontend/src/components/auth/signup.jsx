import React from "react";
import { useForm } from "react-hook-form";

const Signup = () => {
  const onSubmit = (data) => {
    console.log(data);
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
    </form>
  );
};

export default Signup;
