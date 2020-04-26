import React, { useState } from "react";
import { useForm } from "react-hook-form";
import Axios from "axios";
import { ADD_THEATRE, options } from "../../config/url";

const AddTheatres = () => {
  const { register, handleSubmit, errors } = useForm();
  const [state, setState] = useState({});
  const addMovie = (data, e) => {
    console.log(data);
    Axios.post(ADD_THEATRE, data, options)
      .then((responseData) => {
        console.log(responseData.data);
        setState(responseData.data);
        if (responseData.data.statusCode < 300) {
          e.target.reset();
        }
      })
      .catch((err) => {
        console.log(err);
      });
  };
  return (
    <div>
      <h3>ADD THEATRES</h3>
      <form onSubmit={handleSubmit(addMovie)}>
        <div className="form-group">
          <label>Theatre Name</label>
          <input
            type="text"
            className="form-control"
            name="theatrename"
            ref={register({ required: true })}
          />
          {errors.theatrename && <p>Theatre name is required</p>}
        </div>
        <div className="form-group">
          <label>Location</label>
          <input
            type="text"
            className="form-control"
            name="location"
            ref={register({ required: true })}
          />
          {errors.location && <p>Location is required</p>}
        </div>
        <div className="form-group">
          <label>City</label>
          <input
            type="text"
            className="form-control"
            name="city"
            ref={register({ required: true })}
          />
          {errors.city && <p>City is required</p>}
        </div>
        <button className="btn btn-primary" type="submit">
          Add
        </button>
      </form>
      {state && state.statusCode > 300 && (
        <div className="alert alert-warning mt-2">{state.message}</div>
      )}
      {state && state.statusCode < 300 && (
        <div className="alert alert-success mt-2">{state.message}</div>
      )}
    </div>
  );
};

export default AddTheatres;
