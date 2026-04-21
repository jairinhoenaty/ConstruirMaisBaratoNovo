import { ILogin } from "../interfaces/ILogin";
import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

const login = (data: ILogin) => ApiPublica.post('/login', data);

const redeemCode = (code:string) => ApiPublica.post('/redeem-code',{code});

export const LoginService= {
   login,
   redeemCode
}
