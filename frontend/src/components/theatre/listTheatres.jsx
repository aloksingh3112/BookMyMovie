import React from "react";
import { useState, useEffect } from "react";
import Axios from "axios";
import { options, LIST_THEATRE, DELETE_THEATRE } from "../../config/url";

const ListTheatres = (props) => {
  const [theatres, setTheatres] = useState([]);
  const [deleted, setDeleted] = useState(false);
  useEffect(() => {
    Axios.get(LIST_THEATRE, options).then((responseData) => {
      console.log(responseData);
      if (responseData.data.statusCode < 300) {
        setTheatres(responseData.data.data);
      } else {
        setTheatres([]);
      }
    });
    //     return () => {
    //       cleanup;
    //     };
  }, [props, deleted]);

  const deleteTheatre = (id) => {
    Axios.delete(`${DELETE_THEATRE}/${id}`, options)
      .then((responseData) => {
        console.log(responseData.data.message);
        if (responseData.data.statusCode < 300) {
          alert(responseData.data.message);
          setDeleted(!deleted);
        } else {
          alert(responseData.data.message);
        }
      })
      .catch((err) => {
        console.log(err);
      });
    console.log(id);
  };

  return (
    <React.Fragment>
      {theatres && theatres.length > 0 ? (
        <table className="table mt-4">
          <thead className="thead-light listTheatreTable">
            <tr>
              <th>Theatre Name</th>
              <th>Location</th>
              <th>City</th>
              <th>Action</th>
            </tr>
          </thead>
          <tbody>
            {theatres.map((theatre) => {
              return (
                <tr key={theatre.ID}>
                  <td>{theatre.theatrename}</td>

                  <td>{theatre.location}</td>

                  <td>{theatre.city}</td>

                  <td>
                    <button
                      className="btn btn-danger"
                      onClick={() => {
                        deleteTheatre(theatre.ID);
                      }}
                    >
                      <img
                        src="https://cdn2.iconfinder.com/data/icons/apple-inspire-white/100/Apple-64-512.png"
                        width="15"
                      />
                    </button>
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      ) : (
        <h4>No Theatre Found</h4>
      )}
    </React.Fragment>
  );
};

export default ListTheatres;
