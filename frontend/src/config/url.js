const BASE_URL = `http://localhost:3000`;

export const options = {
  headers: {
    Authorization: localStorage.getItem("token")
      ? localStorage.getItem("token")
      : "",
  },
};
export const SIGN_UP = `${BASE_URL}/auth/signup`;
export const LOGIN = `${BASE_URL}/auth/login`;
