import React from "react";

const SeatContainer = () => {
  return (
    <h1>
      <div className="screen">SCREEN</div>
      <div className="row mt-4">
        <div className="col-md-10">
          <center>
            <table className="table table-borderless seat-table">
              <thead>
                <tr>
                  <th>1</th>
                  <th>2</th>
                  <th>3</th>
                  <th>4</th>
                  <th>5</th>
                  <th>6</th>
                  <th></th>
                  <th></th>
                  <th>7</th>
                  <th>8</th>
                  <th>9</th>
                  <th>10</th>
                  <th>11</th>
                  <th>12</th>
                </tr>
              </thead>
              <tbody></tbody>
            </table>
          </center>
        </div>
        <div className="col-md-2 seat-selection ">
          <div className="smallBox greenBox">Selected Seat</div>
          <div className="smallBox emptyBox">Empty Seat</div>
        </div>
      </div>
    </h1>
  );
};

export default SeatContainer;

// {/* <React.Fragment>
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
//               </td> */}

//               <br />
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
//       </div>
//     </React.Fragment> */}
