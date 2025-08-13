//import { IBanner } from "../interfaces/IBanner";
import { IContact } from "../interfaces/IContact";
import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

//const getBannerById = (id: number) => Api.get("/banner/" + id);
const getContacts = (limit: number, offset: number) =>
  Api.get("/contacts?limit=" + limit + "&offset=" + offset);
const saveContact = (data: IContact) => ApiPublica.post("/contact", data);
const deleteContact = (id: any) => Api.delete("/contact/"+ id);
const getContactsUser = (data: {
  limit: number;
  offset: number;
  professional_id: number;
  client_id: number;
  store_id: number;
}) => Api.post("/contacts", data);

export const ContactService = {
  saveContact,
  getContacts,
  getContactsUser,
  deleteContact,
};
