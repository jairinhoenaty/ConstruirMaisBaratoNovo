import Api from "../providers/Api";

const productCategoriesByProfession = (profession_id: number) =>
  Api.get("/product_category/" + profession_id);


export const ProductCategoryService = {
  productCategoriesByProfession,

};
