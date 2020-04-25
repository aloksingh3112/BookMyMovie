import React from "react";
import { Route, Redirect } from "react-router-dom";
import getToken from "./jwt";

const ProtectedRoute = ({ path, component: Component, render, ...rest }) => {
  const token = localStorage.getItem("token");
  //console.log(decodeUser);
  return (
    <Route
      path={path}
      {...rest}
      render={(props) => {
        if (!getToken(token)) {
          localStorage.clear();
          return <Redirect to="/login" />;
        } else {
          return Component ? <Component {...props}></Component> : render(props);
        }
      }}
    />
  );
};

export default ProtectedRoute;
