import { ILogin } from "../interfaces/ILogin";
import ApiPublica from "../providers/ApiPublica";

const login = (data: ILogin) => ApiPublica.post('/login', data);

export const LoginService= {
   login
}