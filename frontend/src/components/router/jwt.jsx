import jwt from "jwt-decode";

function getToken(token) {
  try {
    return jwt(token);
  } catch (Error) {
    return null;
  }
}

export default getToken;
