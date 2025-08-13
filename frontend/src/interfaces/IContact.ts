export interface IContact {
  id: number;
  name: string;
  telefone: string;
  email: string;
  mensagem: string;
  status: string;
  city_id: number | null;
  professional_id: number | null;
  client_id: number | null;
  store_id: number | null;
  product_id: number | null;
  approved: boolean;
  product: { Name: string };
  created_at: Date;
}
