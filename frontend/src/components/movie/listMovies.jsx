import React from "react";
import { useState } from "react";
import Axios from "axios";

const ListMovies = (props) => {
  const [movies, setMovies] = useState([]);

  useEffect(() => {
    // Axios.get()
    //     return () => {
    //       cleanup;
    //     };
  });

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
            <tr>
              <td>{movie.Title}</td>

              <td>{movie.Year}</td>

              <td>{movie.Genre}</td>

              <td>{movie.Actors}</td>

              <td>{movie.Director}</td>

              <td>{movie.Language}</td>

              <td>{movie.Runtime}</td>

              <td>
                <button className="btn btn-primary">+</button>
              </td>
            </tr>
          </tbody>
        </table>
      ) : (
        <h4>No movies Found</h4>
      )}
    </React.Fragment>
  );
};
