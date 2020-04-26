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
export const ADD_MOVIE = `${BASE_URL}/movie/addMovie`;
export const LIST_MOVIE = `${BASE_URL}/movie/getMovies`;
export const DELETE_MOVIE = `${BASE_URL}/movie/deleteMovie`;
