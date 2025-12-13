import { IStore } from "./IStore";

export interface IBudget {
  id?: number;
  name: string;
  email: string;
  telephone: string;
  description: string;
  professionalsId: number[];
  storesId: number[];
  cityId: number;
  termResponsabilityAccepted: boolean;
  approved: boolean;
  // clientId: number;
}
