import axios from "axios";

const baseURL = import.meta.env.VITE_BASE_URL_PUBLICA;

const ApiPublica = axios.create({
  baseURL: baseURL,
});

export default ApiPublica;