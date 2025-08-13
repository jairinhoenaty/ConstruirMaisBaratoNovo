//import { IGetProfessional } from "../interfaces/IGetProfissional";
import axios from "axios";
import { IClient } from "../interfaces/IClient";

import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

const getClientbyID = (id: Number) => Api.get("/client/" + id);
const getClients = (limit: Number, offset: Number) =>
  Api.get("/clients?limit=" + limit + "&offset=" + offset);
const postClient = (data: IClient) => Api.post("/client", data);
const postClientPublic = (data: IClient) =>
  ApiPublica.post("save/client", data);
const lastClients = (data: { quantity: number }) =>
  Api.post("/last/clients", data);

const deleteClient = (id: Number) => Api.delete("/client/" + id);

const postClientXLSX = () => {
  const baseURL = import.meta.env.VITE_BASE_URL;
  const token = localStorage.getItem("token"); //"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9";]
  axios.defaults.headers.common = { Authorization: `Bearer ${token}` };
  axios({
    url: baseURL + "/export-clients-XLSX", //your url
    method: "POST",
    responseType: "blob", // important
  }).then((response) => {
    // create file link in browser's memory
    console.log(response.headers);
    const href = URL.createObjectURL(response.data);

    // create "a" HTML element with href to file & click
    const link = document.createElement("a");
    link.href = href;
    link.setAttribute("download", "clientes.xlsx"); //or any other extension
    document.body.appendChild(link);
    link.click();

    // clean up "a" element & remove ObjectURL
    document.body.removeChild(link);
    URL.revokeObjectURL(href);
    axios.defaults.headers.common = {};
  });
};

export const ClientService = {
  postClient,
  getClients,
  postClientPublic,
  getClientbyID,
  lastClients,
  deleteClient,
  postClientXLSX,
};
