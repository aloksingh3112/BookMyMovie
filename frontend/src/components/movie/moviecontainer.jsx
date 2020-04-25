import React from "react";
import AddMovies from "./addmovie";

const MovieContainer = () => {
  const addMovie = (movie) => {
    console.log("gg", movie);
  };
  return (
    <div>
      <h4>Add Movie</h4>
      <hr />
      <AddMovies addMovie={addMovie} />
    </div>
  );
};

export default MovieContainer;
