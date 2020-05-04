import React from "react";
import StripeCheckout from "react-stripe-checkout";
import Axios from "axios";
import { STRIPE_KEY } from "../../config/apikey";

const StripeCheckOut = (props) => {
  const onToken = (token) => {
    console.log(token);
    const body = {
      amount: 999,
      token: token,
    };

    Axios.post("http://localhost:8000/payment", body)
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
      amount={props.amount} //Amount in cents $9.99
      token={onToken}
      stripeKey={STRIPE_KEY}
      currency="INR"
      billingAddress={false}
    />
  );
};

export default StripeCheckOut;
