import { IProduct } from "../interfaces/IProduct";
import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

const productsDayOfferPublic = () => ApiPublica.post("/products/dayoffer");
const productsAll = (
  limit: number,
  offset: number,
  professional_id: number,
  store_id: number,
  dayoffer: String | null,
  approved: String | null
) =>
  Api.get(
    "/products?limit=" +
      limit +
      "&offset=" +
      offset +
      "&professional_id=" +
      professional_id +
      "&store_id=" +
      store_id +
      (dayoffer === null ? "" : "&dayoffer=" + dayoffer) +
      (approved === null ? "" : "&approved=" + approved)
  );

const productsAllPublic = (
  limit: number,
  offset: number,
  professional_id: number,
  store_id: number,
  dayoffer: string | null,
  approved: string | null
) =>
  ApiPublica.get(
    "/products?limit=" +
      limit +
      "&offset=" +
      offset +
      "&professional_id=" +
      professional_id +
      "&store_id=" +
      store_id +
      "&dayoffer=" +
      (dayoffer === null ? "" : dayoffer) +
      "&approved=" +
      (approved === null ? "" : approved)
  );

const productsByCity = (data: any) =>
  ApiPublica.post("/products/findbycity", data);
//const citiesByStatePublic = (data:any) => ApiPublica.post('/cities-by-state',data);
const saveProduct = (data: IProduct) => Api.post("/product", data);
const productById = (id: number) => Api.get("/product/" + id);
const deleteProduct = (id: number) => Api.delete("/product/" + id);

export const ProductService = {
  productsDayOfferPublic,
  productsAll,
  productsByCity,
  saveProduct,
  productById,
  productsAllPublic,
  deleteProduct,
  //citiesByStatePublic
};
