import axios from "axios";

export const axiosInstance = axios.create({
  baseURL: "/api",
  headers: {
    "Content-type": "application/json"
  },
  withCredentials: true
});

