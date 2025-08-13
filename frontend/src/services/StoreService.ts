import axios from "axios";
import { IStore } from "../interfaces/IStore";
import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

const getStorebyID = (id: Number) => Api.get("/store/" + id);
const getStores = (limit: Number, offset: Number) =>
  Api.get("/stores?limit=" + limit + "&offset=" + offset);
const postStore = (data: IStore) => Api.post("/store", data);
const postStorePublic = (data: IStore) => ApiPublica.post("save/store", data);
const lastStores = (data: { quantity: number }) =>
  Api.post("/last/stores", data);

const deleteStore = (id: Number) => Api.delete("/store/" + id);

const postStoreXLSX = () => {
  const baseURL = import.meta.env.VITE_BASE_URL;
  const token = localStorage.getItem("token"); //"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9";]
  axios.defaults.headers.common = { Authorization: `Bearer ${token}` };
  axios({
    url: baseURL + "/export-stores-XLSX", //your url
    method: "POST",
    responseType: "blob", // important
  }).then((response) => {
    // create file link in browser's memory
    console.log(response.headers);
    const href = URL.createObjectURL(response.data);

    // create "a" HTML element with href to file & click
    const link = document.createElement("a");
    link.href = href;
    link.setAttribute("download", "lojistas.xlsx"); //or any other extension
    document.body.appendChild(link);
    link.click();

    // clean up "a" element & remove ObjectURL
    document.body.removeChild(link);
    URL.revokeObjectURL(href);
    axios.defaults.headers.common = {};
  });
};

export const StoreService = {
  postStorePublic,
  getStores,
  getStorebyID,
  postStore,
  lastStores,
  deleteStore,
  postStoreXLSX,
};
