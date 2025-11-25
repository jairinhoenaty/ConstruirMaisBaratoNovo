import { Plan, UserType } from "../interfaces/IPlan";
import ApiPublica from "../providers/ApiPublica";

const getAllActivePlans = () =>
  ApiPublica.get<Plan[]>("/plans");

const getPlanByUserType = (userType: UserType) =>
  ApiPublica.get<Plan>(`/plans/${userType}`);

export const PlanService = {
  getAllActivePlans,
  getPlanByUserType,
};
