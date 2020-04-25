import React, { useState } from "react";
import { useForm } from "react-hook-form";
import Axios from "axios";
import { API_KEY } from "../../config/apikey";
import { ADD_MOVIE, options } from "../../config/url";

const AddMovies = (props) => {
  const [movie, setMovie] = useState({});
  const { register, handleSubmit, errors } = useForm();
  const addMovie = (data) => {
    const movie = {
      Title: data.Title,
      Language: data.Language,
      Genre: data.Genre,
      Director: data.Director,
      StarCast: data.Actors,
      Year: data.Year,
      Duration: data.Runtime,
      Poster: data.Poster,
      ImdbID: data.imdbID,
    };
    Axios.post(ADD_MOVIE, movie, options)
      .then((responseData) => {
        console.log(responseData);
      })
      .catch((err) => console.log(err));
    props.addMovie(movie);
  };
  const searchMovie = (data, e) => {
    Axios.get(
      `http://www.omdbapi.com/?t=${data.title}&y=${data.year}&apikey=${API_KEY}`
    )
      .then((responseData) => {
        console.log(responseData);

        setMovie(responseData.data);
      })
      .catch((err) => {
        console.log(err);
      });
  };
  return (
    <React.Fragment>
      <form onSubmit={handleSubmit(searchMovie)}>
        <div className="form-group">
          <label htmlFor="">Movie Title</label>
          <input
            type="text"
            className="form-control"
            name="title"
            ref={register({ required: true })}
          />
          {errors.title && <p>Movie title is required</p>}
        </div>
        <div className="form-group">
          <label htmlFor="">Year</label>
          <input
            type="text"
            className="form-control"
            name="year"
            ref={register}
          />
        </div>
        <button className="btn btn-primary" type="submit">
          Search Movie
        </button>
      </form>
      {movie && movie.Response == "False" && (
        <h4 className="mt-3">{movie.Error}</h4>
      )}
      {movie && movie.Response == "True" && (
        <table className="table mt-4">
          <thead className="thead-light">
            <tr>
              <th>Poster</th>
              <th>Title</th>
              <th>Year</th>
              <th>Genre</th>
              <th>Star Cast</th>
              <th>Director</th>
              <th>Language</th>
              <th>Duration</th>
              <th>Action</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>
                <img src={movie.Poster} width="100"></img>
              </td>
              <td>{movie.Title}</td>

              <td>{movie.Year}</td>

              <td>{movie.Genre}</td>

              <td>{movie.Actors}</td>

              <td>{movie.Director}</td>

              <td>{movie.Language}</td>

              <td>{movie.Runtime}</td>

              <td>
                <button
                  className="btn btn-primary"
                  onClick={() => addMovie(movie)}
                >
                  +
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      )}
    </React.Fragment>
  );
};

export default AddMovies;
