import ApiPublica from "../providers/ApiPublica";
import { IAccountDeletionRequest } from "../interfaces/IAccountDeletion";

const create = (data: IAccountDeletionRequest) => {
  return ApiPublica.post("/account-deletion-request", data);
};

export const AccountDeletionService = {
  create,
};