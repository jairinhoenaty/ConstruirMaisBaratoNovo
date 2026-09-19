export interface IAccountDeletionRequest {
  userId?: number | null;
  name: string;
  email: string;
  telephone?: string;
  reason?: string;
}