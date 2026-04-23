import { ICity } from "./ICity";

export interface IProfissional {
  oid: number;
  nome: string;
  email: string;
  telefone: string;
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
  isPremium: boolean;
  dateOfBirth: string | null;
  experience: string | null;
  meiCnpj: string | null;
  telefoneVerificado: string | null;
  negativeCertificateNumber: string | null;
  zona: string | null;
}