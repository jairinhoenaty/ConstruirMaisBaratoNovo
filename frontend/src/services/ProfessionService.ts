//import { IGetProfessional } from "../interfaces/IGetProfissional";
import { IProfession } from "../interfaces/IProfession";
import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

const getProfessions = (limit: number, offset: number) =>
  Api.get("/professions?limit=" + limit + "&offset=" + offset);
const getProfessionsPublic = () => ApiPublica.get("/professions/all");
const postProfession = (data: IProfession) => Api.post("/profession", data);
const getProfessionbyID = (id: number) => Api.get("/profession/" + id);
const deleteProfession = (id: number) => Api.delete("/profession/" + id);

export const ProfessionService = {
  getProfessions,
  getProfessionsPublic,
  postProfession,
  getProfessionbyID,
  deleteProfession,
};
