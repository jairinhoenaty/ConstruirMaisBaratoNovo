export interface IProfissional {
  oid: number;
  Name: string;
  Email: string;
  Telephone: string;
  //LgpdAceito: string;
  cep: string;
  street: string;
  neighborhood: string;
  Password: string | null;
  cityId: number;
  professionIds: string[];
  image: string | null;
  verified: boolean|null;
}
