export interface IBudget {
  id?: number;
  name: string;
  email: string;
  telephone: string;
  description: string;
  professionalsId: number[];
  cityId: number;
  termResponsabilityAccepted: boolean;
  // clientId: number;
}
