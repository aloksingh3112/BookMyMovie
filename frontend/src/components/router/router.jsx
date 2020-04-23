import React from "react";
import { Switch, Route } from "react-router-dom";
import Home from "../home/home";
import TheatreDate from "../theatreDate/theatredate";
import SeatContainer from "../seatContainer/seatcontainer";
import Confirmation from "../confirmation/confirmation";
import MovieContainer from "../movie/moviecontainer";
import ThereContainer from "../theatre/theatreContainer";
import Login from "../auth/login";
import Signup from "../auth/signup";

const Router = () => {
  return (
    <Switch>
      <Route exact path="/" component={Home}></Route>
      <Route path="/theatreDate/:id" component={TheatreDate}></Route>
      <Route path="/seat" component={SeatContainer}></Route>
      <Route path="/booked" component={Confirmation}></Route>
      <Route path="/login" component={Login}></Route>
      <Route path="/signup" component={Signup}></Route>

      <Route path="/addMovie" component={MovieContainer}></Route>
      <Route path="/addTheatre" component={ThereContainer}></Route>
    </Switch>
  );
};

export default Router;
