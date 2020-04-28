import React from "react";
import { useState, useEffect } from "react";
import Axios from "axios";
import {
  LIST_MOVIE,
  options,
  LIST_THEATRE,
  MAP_MOVIE_THEATRE,
} from "../../config/url";
import MultiSelect from "react-multi-select-component";
import { useForm } from "react-hook-form";

const MapTheatre = (props) => {
  const [movies, setMovies] = useState([]);
  const [theatres, setTheatres] = useState([]);
  const [cities, setCity] = useState([]);
  const { register, handleSubmit, errors } = useForm();

  const [selectedTheatres, setSelectedTheatres] = useState([]);
  const [selectedDate, setSelectedDate] = useState([]);
  const [selectedTime, setSelectedTime] = useState([]);
  const [theatreOptions, setTheatreOptions] = useState([]);

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
      console.log(responseArr[0].data.data);
      console.log(responseArr[1].data.data);

      const obj = {};
      for (let data of responseArr[1].data.data) {
        obj[data.city] = data.city;
      }
      var cities = Object.keys(obj);
      setCity(cities);
      setMovies(responseArr[0].data.data);
      setTheatres(responseArr[1].data.data);
    });
  }, [props]);

  const selectCity = (e) => {
    setSelectedTheatres([]);
    console.log(e.target.value);
    var data = [];
    var theatreData = [];
    if (e.target.value != "noData") {
      data = theatres.filter((t) => t.city === e.target.value);
    }

    for (let d of data) {
      const obj = { label: d.theatrename, value: d.ID };
      theatreData.push(obj);
    }
    setTheatreOptions(theatreData);
  };

  const mapMovieTheatre = (data) => {
    if (data.city == "noData") {
      alert("Please select city");
      return;
    }
    if (selectedTheatres.length <= 0) {
      alert("Please select theatre");
      return;
    }
    if (selectedDate.length <= 0) {
      alert("Please select date");
      return;
    }
    if (selectedTime.length <= 0) {
      alert("Please select time");
      return;
    }
    console.log(data, selectedTheatres, selectedDate, selectedTime);
    var dates = [];
    var theatreIDS = [];
    var times = [];
    for (let d of selectedDate) {
      dates.push(d.value);
    }
    for (let d of selectedTheatres) {
      theatreIDS.push(d.value);
    }
    for (let d of selectedTime) {
      times.push(d.value);
    }
    const obj = {
      movieID: data.movieID,
      dates: dates,
      theatreIDS: theatreIDS,
      times: times,
    };
    console.log(obj);
    Axios.post(MAP_MOVIE_THEATRE, obj, options)
      .then((responseData) => {
        console.log(responseData.data);
        alert(responseData.data.message);
      })
      .catch((err) => console.log(err));
  };

  return (
    <React.Fragment>
      <form onSubmit={handleSubmit(mapMovieTheatre)}>
        <div className="form-group">
          <label htmlFor="">Movie</label>
          <select
            name="movieID"
            id=""
            className="form-control"
            ref={register({ required: true })}
          >
            {movies &&
              movies.length > 0 &&
              movies.map((m) => {
                return (
                  <option value={m.ID} key={m.ID}>
                    {m.title}
                  </option>
                );
              })}
          </select>
          {errors.movieID && <p>Movie is required</p>}
        </div>

        <div className="form-group">
          <label htmlFor="">City</label>
          <select
            name="city"
            id=""
            className="form-control"
            onChange={selectCity}
            ref={register}
          >
            <option value="noData" selected>
              select city
            </option>
            {cities &&
              cities.length > 0 &&
              cities.map((c, index) => {
                return (
                  <option value={c} key={index}>
                    {c}
                  </option>
                );
              })}
          </select>
          {errors.city && <p>City is required</p>}
        </div>

        <div className="form-group">
          <label htmlFor="">Theatres</label>
          <MultiSelect
            options={theatreOptions}
            value={selectedTheatres}
            onChange={setSelectedTheatres}
            labelledBy={"Select"}
          />
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
            onChange={setSelectedTime}
            labelledBy={"Select"}
          />
        </div>
        <button className="btn btn-primary" type="submit">
          Map Movie
        </button>
      </form>
    </React.Fragment>
  );
};

export default MapTheatre;
