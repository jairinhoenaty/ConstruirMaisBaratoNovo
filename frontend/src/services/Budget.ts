//import { IBanner } from "../interfaces/IBanner";
import { IBudget } from "../interfaces/IBudget";
import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

//const getBannerById = (id: number) => Api.get("/banner/" + id);
//const getBannerByPage = (data: { page: string }) => Api.post("/banners/page", data);
const saveBudget = (data: IBudget) => ApiPublica.post("/save/budget", data);
const deleteBudget = (id: number) => Api.delete("/budget/"+ id);
const getBudgetsbyMonth = (data: {month:number,year:number}) => Api.post("/budgets/month", data);
const getBudgetsAll = (limit:number,offset:number) => Api.get("/budgets?limit=" +
  limit +
  "&offset=" +
  offset);

export const BudgetService = {
  saveBudget,
  getBudgetsbyMonth,
  getBudgetsAll,
  deleteBudget,
};
