//import { IGetProfessional } from "../interfaces/IGetProfissional";
import { off } from "process";
import { IProfissional } from "../interfaces/IProfissional";
import { IProfissionalAdd } from "../interfaces/IProfissionalAdd";
import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";
import axios from "axios";

const getProfessionalbyName = (data: { name: string }) =>
  Api.post("/professionals/name", data);

const getProfessionalbyID = (id: Number) => Api.get("/professional/" + id);
const postProfessionalXLSX2 = () =>
  Api.post("/export-professionals-XLSX", { responseType: "blob" });
const postProfessionalXLSX = () => {
  const baseURL = import.meta.env.VITE_BASE_URL;
  const token = localStorage.getItem("token"); //"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9";]
  axios.defaults.headers.common = { Authorization: `Bearer ${token}` };
  axios({
    url: baseURL + "/export-professionals-XLSX", //your url
    method: "POST",
    responseType: "blob", // important
  }).then((response) => {
    // create file link in browser's memory
    console.log(response.headers);
    const href = URL.createObjectURL(response.data);

    // create "a" HTML element with href to file & click
    const link = document.createElement("a");
    link.href = href;
    link.setAttribute("download", "profissionais.xlsx"); //or any other extension
    document.body.appendChild(link);
    link.click();

    // clean up "a" element & remove ObjectURL
    document.body.removeChild(link);
    URL.revokeObjectURL(href);
    axios.defaults.headers.common = {};
  });
};

const getProfessionals = (
  limit: number,
  offset: number,
  filter: string,
  uf: string,
  professionId: number,
  order:string |null
) =>
  Api.get(
    "/professionals?limit=" +
      limit +
      "&offset=" +
      offset +
      "&filter=" +
      filter +
      "&uf=" +
      uf +
      "&profession_id=" +
      professionId+"&order="+order
  );

const postProfessional = (data: IProfissional) =>
  Api.post("/professional", data);
const saveProfessional = (data: IProfissionalAdd) =>
  Api.post("/saveProfessional", data);
const changePassword = (data: {
  oid: number;
  name: string;
  password: string;
}) => Api.post("/professional", data);
const getProfessionalByCityAndProfession = (data: any) =>
  ApiPublica.post("/search-all-professionals-and-city-and-profession", data);

const countProfessionalByProfession = () =>
  Api.post("/count/professional/profession");

const countProfessionalByState = (data: { limit: number; offset: number }) =>
  Api.post("/count/professionals/state", data);
const ProfessionalByState = (data: { state:string, limit: number; offset: number }) =>
  Api.post("professionals/state", data);

const CountProfessionalsByProfessionByCitie = (data: { cityID: string }) =>
  Api.post("count/professionals/city", data);



const lastProfessionals = (data: { quantity: number }) =>
  Api.post("/last/professionals", data);

const postProfessionalPublic = (data: IProfissional) =>
  ApiPublica.post("/save/professional", data);

const deleteProfessional = (id: Number) => Api.delete("/professional/" + id);

export const ProfessionalService = {
  getProfessionalbyID,
  getProfessionals,
  postProfessional,
  postProfessionalPublic,
  changePassword,
  saveProfessional,
  getProfessionalByCityAndProfession,
  countProfessionalByProfession,
  countProfessionalByState,
  lastProfessionals,
  deleteProfessional,
  getProfessionalbyName,
  postProfessionalXLSX,
  postProfessionalXLSX2,
  ProfessionalByState,
  CountProfessionalsByProfessionByCitie,
};
