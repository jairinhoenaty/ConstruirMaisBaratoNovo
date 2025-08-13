export interface IProduct {
  id: number;
  name: string;
  description: string;
  image: string;
  price: number;
  originalprice: number;
  discount: number;
  approved: boolean;
  dayoffer: boolean;
  professionID: number | null;
  categoryID: number;
  professionalID: number | null;
  storeID: number | null;
}
