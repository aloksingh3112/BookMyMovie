import React from "react";
import { NavLink } from "react-router-dom";

const Header = () => {
  return (
    <React.Fragment>
      <nav className="navbar navbar-expand-lg navbar-light bg-light fixed-top">
        {/* <a className="navbar-brand" href="#">
          Navbar
        </a> */}
        <NavLink
          exact
          to="/"
          className="navbar-brand"
          style={{
            fontWeight: "bold",
            fontSize: 24,
          }}
        >
          Book My Movie
        </NavLink>
        <button
          className="navbar-toggler"
          type="button"
          data-toggle="collapse"
          data-target="#navbarSupportedContent"
          aria-cont
          rols="navbarSupportedContent"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span className="navbar-toggler-icon"></span>
        </button>
      </nav>
    </React.Fragment>
  );
};

export default Header;
