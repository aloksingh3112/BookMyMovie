import React, { useState, useEffect } from "react";
import Axios from "axios";
import { LIST_MOVIE, options } from "../../config/url";

const Home = (props) => {
  const [movies, setMovies] = useState([]);
  useEffect(() => {
    Axios.get(LIST_MOVIE, options)
      .then((responseData) => {
        console.log(responseData.data.data);
        setMovies(responseData.data.data);
      })
      .catch((err) => console.log(err));
  }, [props]);
  return (
    <React.Fragment>
      <div className="d-flex flex-wrap">
        {movies &&
          movies.length > 0 &&
          movies.map((m) => {
            return (
              <div className="card mr-3 mt-4" style={{ width: 250 }}>
                <img
                  className="card-img-top"
                  src={m.poster}
                  alt="Card image cap"
                  height="200"
                />
                <div className="card-body">
                  <h4 className="card-title">{m.title}</h4>
                  <div className="card-text">
                    <h6>
                      <strong>Language</strong>: {m.language}
                    </h6>
                    <h6>
                      <strong>Genre</strong>: {m.genre}
                    </h6>
                    <h6>
                      <strong>Actors</strong>: {m.starcast}
                    </h6>
                  </div>
                </div>
                <div className="card-footer">
                  <a href="#" className="btn btn-danger">
                    Book Now
                  </a>
                </div>
              </div>
            );
          })}
      </div>
    </React.Fragment>
  );
};

export default Home;
