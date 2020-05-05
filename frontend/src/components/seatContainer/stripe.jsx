import React from "react";
import StripeCheckout from "react-stripe-checkout";
import Axios from "axios";
import { STRIPE_KEY } from "../../config/apikey";
import { options } from "../../config/url";

const StripeCheckOut = (props) => {
  const onToken = (token) => {
    console.log(token);
    const body = {
      token: token,
      movieId: props.movieId,
      dateId: props.dateId,
      theatreId: props.theatreId,
      timeId: props.timeId,
      seats: props.seats,
      amount: props.amount,
    };
    console.log(body);

    Axios.post("http://localhost:8000/payment", body, options)
      .then((data) => {
        console.log(data);
      })
      .catch((err) => {
        console.log(err);
      });
  };
  return (
    <StripeCheckout
      label="Proceed To Payment" //Component button text
      name="Book My Movie" //Modal Header
      panelLabel="Book" //Submit button in modal
      amount={props.amount * 100} //Amount in cents $9.99
      token={onToken}
      stripeKey={STRIPE_KEY}
      currency="INR"
      billingAddress={false}
    />
  );
};

export default StripeCheckOut;
