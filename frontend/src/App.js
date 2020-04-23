import React from "react";
import Header from "./components/header/header";
import Router from "./components/router/router";

const App = () => {
  return (
    <React.Fragment>
      <Header></Header>
      <div className="container" style={{ marginTop: "100px" }}>
        <Router></Router>
      </div>
    </React.Fragment>
  );
};

export default App;
