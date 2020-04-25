import React from "react";
import { Route, Redirect } from "react-router-dom";
import getToken from "./jwt";

const AuthorisedRoute = ({ path, component: Component, render, ...rest }) => {
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
          const data = getToken(token);
          console.log(data.typeofuser);
          if (data.typeofuser === "ADMIN") {
            return Component ? (
              <Component {...props}></Component>
            ) : (
              render(props)
            );
          } else {
            return <Redirect to="/login" />;
          }
        }
      }}
    />
  );
};

export default AuthorisedRoute;
