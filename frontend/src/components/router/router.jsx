import React from "react";
import { Switch, Route } from "react-router-dom";
import Home from "../home/home";
import TheatreDate from "../theatreDate/theatredate";
import SeatContainer from "../seatContainer/seatcontainer";
import Confirmation from "../confirmation/confirmation";
import MovieContainer from "../movie/moviecontainer";
import TheatreContainer from "../theatre/theatreContainer";
import Login from "../auth/login";
import Signup from "../auth/signup";
import ProtectedRoute from "./protected";
import AuthorisedRoute from "./authorisedRoute";

const Router = () => {
  return (
    <Switch>
      <Route exact path="/" component={Home}></Route>
      <ProtectedRoute
        path="/theatreDate/:id"
        component={TheatreDate}
      ></ProtectedRoute>
      <ProtectedRoute
        path="/seat/:movieId/:dateId/:theatreId/:timeId"
        component={SeatContainer}
      ></ProtectedRoute>
      <ProtectedRoute path="/booked" component={Confirmation}></ProtectedRoute>
      <Route path="/login" component={Login}></Route>
      {/* <Route path="/signup" component={Signup}></Route> */}
      <Route path="/signup" component={Signup}></Route>

      <AuthorisedRoute
        path="/addMovie"
        component={MovieContainer}
      ></AuthorisedRoute>
      <AuthorisedRoute
        path="/addTheatre"
        component={TheatreContainer}
      ></AuthorisedRoute>
    </Switch>
  );
};

export default Router;
