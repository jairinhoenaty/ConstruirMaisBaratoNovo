import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

const productCategoriesByProfession = (profession_id: number) =>
  Api.get("/product_category/" + profession_id);

const findTopLevelCategory = () =>
  ApiPublica.get("/categories/top-level");

const findSubcategoriesByParentID = (categoryId: number) =>
  ApiPublica.get(`/categories/${categoryId}/subcategories`);

// Admin CRUD operations
const create = (data: { name: string; parent_id?: number | null; profession_id?: number }) =>
  Api.post("/categories", data);

const update = (id: number, data: { name: string; profession_id?: number }) =>
  Api.put(`/categories/${id}`, data);

const deleteCategory = (id: number) =>
  Api.delete(`/categories/${id}`);

export const ProductCategoryService = {
  productCategoriesByProfession,
  findTopLevelCategory,
  findSubcategoriesByParentID,
  create,
  update,
  deleteCategory,
};
