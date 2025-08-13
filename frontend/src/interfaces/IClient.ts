export interface IClient {
  oid: number;
  Name: string;
  Email: string;
  Telephone: string;
  LgpdAceito: string;
  cep: string;
  street: string;
  Password: string | null;
  neighborhood: string;
  cityId: number;
  image: string | null;
}
