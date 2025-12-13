export interface IStore {
  oid: number;
  nome?: string; // Nome no backend
  Name: string;
  email?: string; // Email no backend
  Email: string;
  telefone?: string; // Telefone no backend
  Telephone: string;
  lgpdaceito?: string; // LgpdAceito no backend
  LgpdAceito: string;
  cep: string;
  endereco?: string; // Street no backend
  street: string;
  bairro?: string; // Neighborhood no backend
  neighborhood: string;
  Password: string | null;
  cityId: number;
  image: string | null;
  categoryId?: number; // CategoryId no backend
  categoryProductID: number | null;
  subCategoriesId?: number[]; // SubCategoriesId no backend
	subCategories : number[] | null;
  isPremiumStore?: boolean;
}
