import React from "react";
import { Switch, Route, Link, NavLink } from "react-router-dom";
import ListTheatres from "./listTheatres";
import AddTheatres from "./addTheatres";

const TheatreContainer = () => {
  return (
    <React.Fragment>
      <div className="sidenav">
        <NavLink
          activeStyle={{
            color: "white",
          }}
          to="/addTheatre/add"
        >
          Add Theatre
        </NavLink>
        <NavLink
          activeStyle={{
            color: "white",
          }}
          to="/addTheatre/list"
        >
          List Theatre
        </NavLink>

        <a href="#clients">Clients</a>
        <a href="#contact">Contact</a>
      </div>

      <div className="main">
        <Switch>
          <Route path="/addTheatre/add" component={AddTheatres}></Route>
          <Route path="/addTheatre/list" component={ListTheatres}></Route>
        </Switch>
      </div>
    </React.Fragment>
  );
};

export default TheatreContainer;
