import React, { useState, useEffect } from "react";
import Axios from "axios";
import { GET_SEAT_DATA, options } from "../../config/url";
import StripeCheckOut from "./stripe";
const seatsData = [
  { seatnumber: "A1", price: 500, disabled: false },
  { seatnumber: "A2", price: 500, disabled: false },
  { seatnumber: "A3", price: 500, disabled: false },
  { seatnumber: "A4", price: 500, disabled: false },
  { seatnumber: "A5", price: 500, disabled: false },
  { seatnumber: "A6", price: 500, disabled: false },
  { seatnumber: "A7", price: 500, disabled: false },
  { seatnumber: "A8", price: 500, disabled: false },
  { seatnumber: "A9", price: 500, disabled: false },
  { seatnumber: "A10", price: 500, disabled: false },
  { seatnumber: "A11", price: 500, disabled: false },
  { seatnumber: "B1", price: 500, disabled: false },
  { seatnumber: "B2", price: 500, disabled: false },
  { seatnumber: "B3", price: 500, disabled: false },
  { seatnumber: "B4", price: 500, disabled: false },
  { seatnumber: "B5", price: 500, disabled: false },
  { seatnumber: "B6", price: 500, disabled: false },
  { seatnumber: "B7", price: 500, disabled: false },
  { seatnumber: "B8", price: 500, disabled: false },
  { seatnumber: "B9", price: 500, disabled: false },
  { seatnumber: "B10", price: 500, disabled: false },
  { seatnumber: "B11", price: 500, disabled: false },
];

const SeatContainer = (props) => {
  const [seatsResponse, setSeatsResponse] = useState([]);
  const [seats, setSeats] = useState(seatsData);
  const [noofSeats, setNumberOfSeats] = useState(0);
  const [seatSelected, setSeatSelected] = useState([]);

  const [seatsNumber, seatNumber] = useState();
  const [amount, setAmount] = useState(0);

  useEffect(() => {
    const { movieId, dateId, theatreId, timeId } = props.match.params;
    console.log(movieId, dateId, theatreId, timeId);
    const data = {
      movieId,
      dateId,
      theatreId,
      timeId,
    };
    console.log(data);
    Axios.post(GET_SEAT_DATA, data, options).then((responseData) => {
      console.log(responseData);
    });
  }, [props]);

  const handleChange = (e) => {
    console.log(e.target.checked);
    const index = seats.findIndex((s) => s.seatnumber == e.target.value);
    const data = seats[index];
    if (e.target.checked) {
      setSeatSelected([...seatSelected, data]);
      var count = noofSeats;

      count = count + 1;
      setNumberOfSeats(count);
      var price = amount + data.price;
      setAmount(price);
    } else {
      var s = [...seatSelected];
      const index = s.findIndex((s) => s.seatnumber == e.target.value);
      s.splice(index, 1);
      setSeatSelected(s);
      var count = noofSeats;
      count = count - 1;
      setNumberOfSeats(count);
      var price = amount - data.price;
      setAmount(price);
    }
  };
  return (
    <React.Fragment>
      <div className="screen">SCREEN</div>
      <div className="row mt-4">
        <div className="col-md-8">
          <div className="row">
            <div className="col-md-9 ">
              <div className="d-flex flex-wrap ml-4">
                {seats &&
                  seats.map((s) => {
                    return (
                      <input
                        type="checkbox"
                        className="seats"
                        value={s.seatnumber}
                        disabled={s.disabled}
                        onChange={handleChange}
                      />
                    );
                  })}
              </div>
            </div>
            <div className="col-md-3 seat-selection ">
              <div className="smallBox greenBox">Selected Seat</div>
              <div className="smallBox emptyBox">Empty Seat</div>
            </div>
          </div>
        </div>

        <div className="col-md-4">
          <div className="card" style={{ width: 350 }}>
            <div className="card-body">
              <h3 className="card-title">Booking Summary</h3>
              <hr />
              <div className="row">
                <div className="col-md-8">No. of seat selected</div>
                <div className="col-md-2 text-nowrap">{noofSeats}</div>
              </div>
              <hr />

              <div className="row">
                <div className="col-md-8">Seat No</div>
                <div className="col-md-4 ">
                  {seatSelected &&
                    seatSelected.length > 0 &&
                    seatSelected.map((s, i) => {
                      return (
                        <span>
                          {s.seatnumber}
                          {i != seatSelected.length - 1 && ","}
                        </span>
                      );
                    })}
                </div>
              </div>
              <hr />

              <div className="row">
                <div className="col-md-8">Sub Total</div>
                <div className="col-md-4 text-nowrap">{amount}</div>
              </div>
              <hr />
              <div className="row">
                <div className="col-md-8">Internet handling fees</div>
                <div className="col-md-4">20</div>
              </div>
              <hr />

              <div className="row">
                <div className="col-md-8">Amount payable</div>
                <div className="col-md-4 text-nowrap">
                  {amount == 0 ? 0 : amount + 20}
                </div>
              </div>
              <hr />

              <StripeCheckOut amount={amount + 20} />
            </div>
          </div>
        </div>
      </div>
    </React.Fragment>
  );
};

export default SeatContainer;

{
  /* // <React.Fragment>
//       <div class="seatStructure">
//         <center>
//           <table id="seatsBlock">
//             <tr>
//               <td colspan="14">
//                 <div class="screen">SCREEN</div>
//               </td>
//               {/* <td rowspan="20">
//                 <div class="smallBox greenBox">Selected Seat</div> <br />
//                 <div class="smallBox emptyBox">Empty Seat</div>
//                 <br />
//               </td> */
}

{
  /* //               <br />
//             </tr>

//             <tr>
//               <td></td>
//               <td>1</td>
//               <td>2</td>
//               <td>3</td>
//               <td>4</td>
//               <td>5</td>
//               <td></td>
//               <td>6</td>
//               <td>7</td>
//               <td>8</td>
//               <td>9</td>
//               <td>10</td>
//               <td>11</td>
//               <td>12</td>
//             </tr>

//             <tr>
//               <td>A</td>
//               <td>
//                 <input type="checkbox" class="seats" value="A1" />
//               </td>
//               <td>
//                 <input type="checkbox" class="seats" value="A2" />
//               </td>
//               <td>
//                 <input type="checkbox" class="seats" value="A3" />
//               </td>
//               <td>
//                 <input type="checkbox" class="seats" value="A4" />
//               </td>
//               <td>
//                 <input type="checkbox" class="seats" value="A5" />
//               </td>
//               <td class="seatGap"></td>
//               <td>
//                 <input type="checkbox" class="seats" value="A6" />
//               </td>
//               <td>
//                 <input type="checkbox" class="seats" value="A7" />
//               </td>
//               <td>
//                 <input type="checkbox" class="seats" value="A8" />
//               </td>
//               <td>
//                 <input type="checkbox" class="seats" value="A9" />
//               </td>
//               <td>
//                 <input type="checkbox" class="seats" value="A10" />
//               </td>
//               <td>
//                 <input type="checkbox" class="seats" value="A11" />
//               </td>
//               <td>
//                 <input type="checkbox" class="seats" value="A12" />
//               </td>
//             </tr>
//           </table>
//         </center>
//       </div> */
}
//     </React.Fragment> */}
