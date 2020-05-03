import React from "react";
import { useState, useEffect } from "react";
import Axios from "axios";
import { MAP_DATE_THEATRE } from "../../config/url";

const TheatreDate = (props) => {
  const [movieDates, setMovieSates] = useState([]);
  const [theatresDate, setTheatreDate] = useState([]);
  const [theatreId, setTheatreId] = useState();
  const [dateId, setDateId] = useState();
  const [timeId, setTimeId] = useState();

  useEffect(() => {
    const { id } = props.match.params;

    Axios.get(`${MAP_DATE_THEATRE}/${id}`)
      .then((responseData) => {
        console.log(responseData);
        setMovieSates(responseData.data.data);
      })
      .catch((err) => console.log(err));
  }, [props]);

  const handleChange = (e) => {
    if (e.target.value == "noData") {
      setTheatreDate([]);
      setDateId(null);
      return;
    }

    console.log(e.target.value);
    const index = movieDates.findIndex((m) => m.Date.ID == e.target.value);
    //console.log(movieDates[index]["Theatre"]);
    setTheatreDate(movieDates[index]["Theatre"]);
    setDateId(e.target.value);
  };

  const handleSeat = (theatreId, timeId) => {
    const { id } = props.match.params;

    console.log(theatreId, timeId, Number(dateId), Number(id));
    props.history.push(`/seat/${id}/${dateId}/${theatreId}/${timeId}`);
  };
  return (
    <React.Fragment>
      <div className="form-group">
        <label htmlFor="">
          <strong>Select Date</strong>
        </label>
        <select name="" id="" className="form-control" onChange={handleChange}>
          <option value="noData">Select date</option>
          {movieDates &&
            movieDates.length > 0 &&
            movieDates.map((m, i) => {
              return (
                <option value={m.Date.ID} key={i}>
                  {m.Date.date}
                </option>
              );
            })}
        </select>

        <div className="d-flex flex-wrap mt-4">
          {theatresDate &&
            theatresDate.length > 0 &&
            theatresDate.map((td, i) => {
              return (
                <div className="card mt-4 mr-4" key={i}>
                  <div className="card-header">
                    <h3>{td.Theatre.theatrename}</h3>
                  </div>
                  <div className="card-body">
                    <h5 className="card-title">
                      Location:{td.Theatre.location}
                    </h5>
                    <div className="mt-2">
                      {td.Time.map((t, i) => {
                        return (
                          <button
                            key={i}
                            type="button"
                            className="btn btn-outline-primary mr-1"
                            onClick={() => handleSeat(td.Theatre.ID, t.ID)}
                          >
                            {t.time}
                          </button>
                        );
                      })}
                    </div>
                  </div>
                </div>
              );
            })}
        </div>
      </div>
    </React.Fragment>
  );
};

export default TheatreDate;
