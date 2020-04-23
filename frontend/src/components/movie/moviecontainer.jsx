import React from "react";
import AddMovie from "./addmovie";
import { Route, Switch } from "react-router-dom";

const MovieContainer = () => {
  return (
    <div>
      <Switch>
        <Route path="/addMovie/movie" component={AddMovie}></Route>
      </Switch>
      <h1>Movie Container</h1>;
    </div>
  );
};

export default MovieContainer;
