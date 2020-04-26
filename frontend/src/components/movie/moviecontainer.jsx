import React from "react";
import AddMovies from "./addmovie";
import ListMovies from "./listMovies";
import { useState } from "react";

const MovieContainer = () => {
  const [movie, setMovie] = useState();
  const addMovie = (m) => {
    console.log("gg", m);
    setMovie(m);
  };
  return (
    <div>
      <h4>Add Movie</h4>
      <hr />
      <AddMovies addMovie={addMovie} />
      <h4 className="mt-3">List Movies</h4>
      <hr />
      <ListMovies movie={movie} />
    </div>
  );
};

export default MovieContainer;
