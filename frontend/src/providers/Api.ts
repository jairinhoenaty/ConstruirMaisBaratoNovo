import axios from "axios";


const baseURL = import.meta.env.VITE_BASE_URL;
console.log(import.meta.env);
console.log(baseURL);

const Api = axios.create({
  baseURL: baseURL,
});

Api.interceptors.request.use(async config => {
  const token = localStorage.getItem('token'); //"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9";]
  if (token) {
    config.headers['Authorization'] = `Bearer ${token}`;
    
  }  
  return config;
});

export default Api; 