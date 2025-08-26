import { ICity } from "./ICity";

export interface IProfissional {
  oid: number;
  nome: string;
  email: string;
  telephone: string;
  //LgpdAceito: string;
  cep: string;
  street: string;
  neighborhood: string;
  password: string | null;
  cityId: number;
  cidade: ICity;
  professionIds: string[];
  image: string | null;
  verified: boolean|null;
}
