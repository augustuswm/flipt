import axios from "axios";

let host =
  process.env.NODE_ENV === "production"
    ? window.location.host
    : "localhost:8080";

let cookies = document.cookie.split(';');
cookies = cookies.map(cookie => cookie.split('=')).filter(([name, value]) => name === 'token');

let auth = cookies.length > 0 ? cookies[0][1] : '';
let headers = auth ? { 'Authorization': `Bearer ${auth}` } : {};

export const Api = axios.create({
  baseURL: "//" + host + "/api/v1/",
  headers
});