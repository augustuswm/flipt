import axios from "axios";

let host =
  process.env.NODE_ENV === "production"
    ? window.location.host
    : "localhost:8080";

// Under the assumption that the token is the only cookie
let auth = document.cookie ? document.cookie.split('=')[1] : '';
let headers = auth ? { 'Authorization': `Bearer ${auth}` } : {};

export const Api = axios.create({
  baseURL: "//" + host + "/api/v1/",
  headers
});