import React from "react";
import { useState, useEffect } from "react";
import Axios from "axios";
import { LIST_MOVIE, options, DELETE_MOVIE } from "../../config/url";

const ListMovies = (props) => {
  const [movies, setMovies] = useState([]);
  const [deleted, setDeleted] = useState(false);
  useEffect(() => {
    Axios.get(LIST_MOVIE, options).then((responseData) => {
      console.log(responseData);
      if (responseData.data.statusCode < 300) {
        setMovies(responseData.data.data);
      } else {
        setMovies([]);
      }
    });
    //     return () => {
    //       cleanup;
    //     };
  }, [props, deleted]);

  const deleteMovie = (id) => {
    Axios.delete(`${DELETE_MOVIE}/${id}`, options)
      .then((responseData) => {
        console.log(responseData.data.message);
        if (responseData.data.statusCode < 300) {
          alert(responseData.data.message);
          setDeleted(!deleted);
        } else {
          alert(responseData.data.message);
        }
      })
      .catch((err) => {
        console.log(err);
      });
    console.log(id);
  };

  return (
    <React.Fragment>
      {movies && movies.length > 0 ? (
        <table className="table mt-4">
          <thead className="thead-light">
            <tr>
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
            {movies.map((movie) => {
              return (
                <tr key={movie.ID}>
                  <td>{movie.title}</td>

                  <td>{movie.year}</td>

                  <td>{movie.genre}</td>

                  <td>{movie.starcast}</td>

                  <td>{movie.director}</td>

                  <td>{movie.language}</td>

                  <td>{movie.duration}</td>

                  <td>
                    <button
                      className="btn btn-primary"
                      onClick={() => {
                        deleteMovie(movie.ID);
                      }}
                    >
                      <img
                        src="https://cdn2.iconfinder.com/data/icons/apple-inspire-white/100/Apple-64-512.png"
                        width="15"
                      />
                    </button>
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      ) : (
        <h4>No movies Found</h4>
      )}
    </React.Fragment>
  );
};

export default ListMovies;
