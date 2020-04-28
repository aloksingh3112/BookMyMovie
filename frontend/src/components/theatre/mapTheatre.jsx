import React from "react";
import { useState, useEffect } from "react";
import Axios from "axios";
import { LIST_MOVIE, options, LIST_THEATRE } from "../../config/url";
import MultiSelect from "react-multi-select-component";

const MapTheatre = (props) => {
  const [movies, setMovies] = useState([]);
  const [theatres, setTheatres] = useState([]);
  const [cities, setCity] = useState([]);

  const [selectedTheatres, setSelectedTheatres] = useState([]);
  const [selectedDate, setSelectedDate] = useState([]);
  const [selectedTime, setSelectedTime] = useState([]);

  const dateOptions = [
    { label: "28 April", value: "28 April" },
    { label: "29 April", value: "29 April" },
    { label: "30 April", value: "30 April" },
  ];
  const timeOptions = [
    { label: "10:30 am", value: "10:30 am" },
    { label: "12:30 pm", value: "12:30 pm" },
    { label: "03:30 pm", value: "03:30 pm" },
  ];

  useEffect(() => {
    Axios.all([
      Axios.get(LIST_MOVIE, options),
      Axios.get(LIST_THEATRE, options),
    ]).then((responseArr) => {
      console.log(responseArr[0]);
      console.log(responseArr[1]);
    });
  }, [props]);
  return (
    <React.Fragment>
      <form>
        <div className="form-group">
          <label htmlFor="">Movie</label>
          <select name="movieID" id="" className="form-control"></select>
        </div>
        <div className="form-group">
          <label htmlFor="">City</label>
          <select name="" id="" className="form-control"></select>
        </div>
        <div className="form-group">
          <label htmlFor="">Theatres</label>
          <select name="" id="" className="form-control"></select>
        </div>
        <div className="form-group">
          <label htmlFor="">Select Date</label>
          <MultiSelect
            options={dateOptions}
            value={selectedDate}
            onChange={setSelectedDate}
            labelledBy={"Select"}
          />
        </div>

        <div className="form-group">
          <label htmlFor="">Select Time</label>
          <MultiSelect
            options={timeOptions}
            value={selectedTime}
            onChange={selectedTime}
            labelledBy={"Select"}
          />
        </div>
      </form>
    </React.Fragment>
  );
};

export default MapTheatre;
