import { IUser } from "../interfaces/IUser";
import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

const getUserById = (id: number) => Api.get("/user/" + id);
const saveUser = (data: IUser) => Api.post("/user", data);
const findbyemailPublic = (data: any) =>
  ApiPublica.post("/user/find-by-email", data);
const findbyemail = (data: any) => Api.post("/find-by-email", data);
const sendMail = (data: {email:string}) => ApiPublica.post("/user/send-mail", data);
const resetPassword = (data: {email:string,password:string}) => ApiPublica.post("/reset/password", data);

export const UserService = {
  getUserById,
  saveUser,
  findbyemailPublic,
  findbyemail,
  sendMail,
  resetPassword
};
