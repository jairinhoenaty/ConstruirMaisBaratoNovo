//import { IBanner } from "../interfaces/IBanner";
import { IBudget } from "../interfaces/IBudget";
import { ICheckBudgetAccess, IUnlockBudgetPaymentInput, IUnlockBudgetPaymentOutput } from "../interfaces/IUnlockedBudget";
import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";


//const getBannerById = (id: number) => Api.get("/banner/" + id);
//const getBannerByPage = (data: { page: string }) => Api.post("/banners/page", data);
const saveBudget = (data: IBudget) => {
  return ApiPublica.post("/save/budget", data);
};
const deleteBudget = (id: number) => Api.delete("/budget/"+ id);
// Recusa: some só para o destinatário; não exclui o orçamento.
const refuseBudget = (
  id: number,
  recipientId: number,
  recipientType: "professional" | "store"
) => Api.patch("/budget/" + id + "/refuse", { recipientId, recipientType });
const getBudgetsbyMonth = (data: {month:number,year:number}) => Api.post("/budgets/month", data);
const getBudgetsAll = (limit:number,offset:number) => Api.get("/budgets?limit=" +
  limit +
  "&offset=" +
  offset);

const checkBudgetAccess = (budgetId: number) =>
  Api.get<ICheckBudgetAccess>(`/budget/${budgetId}/check-access`);

const unlockBudget = (budgetId: number, data: IUnlockBudgetPaymentInput) =>
  Api.post<IUnlockBudgetPaymentOutput>(`/budget/${budgetId}/unlock`, data);

export const BudgetService = {
  saveBudget,
  getBudgetsbyMonth,
  getBudgetsAll,
  deleteBudget,
  refuseBudget,
  checkBudgetAccess,
  unlockBudget,
};
