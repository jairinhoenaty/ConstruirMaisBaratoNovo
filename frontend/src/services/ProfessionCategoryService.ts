import { IProfessionCategory } from "../interfaces/IProfessionCategory";
import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

const getActiveCategories = () => ApiPublica.get("/profession-categories");

const getProfessionsByCategoryPublic = (categoryId: number) =>
  ApiPublica.get(`/profession-categories/${categoryId}/professions`);

const getCategories = (limit: number, offset: number) =>
  Api.get(`/profession-categories?limit=${limit}&offset=${offset}`);

const getCategoryById = (id: number) => Api.get(`/profession-category/${id}`);

const getProfessionsByCategory = (categoryId: number) =>
  Api.get(`/profession-category/${categoryId}/professions`);

const postCategory = (data: Partial<IProfessionCategory>) =>
  Api.post("/profession-category", data);

const deleteCategory = (id: number) => Api.delete(`/profession-category/${id}`);

export const ProfessionCategoryService = {
  getActiveCategories,
  getProfessionsByCategoryPublic,
  getCategories,
  getCategoryById,
  getProfessionsByCategory,
  postCategory,
  deleteCategory,
};
