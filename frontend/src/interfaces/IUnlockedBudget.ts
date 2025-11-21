export interface IUnlockedBudget {
  id: number;
  professionalId: number;
  budgetId: number;
  status: 'pending' | 'paid' | 'failed' | 'cancelled';
  paymentId: string;
  amount: number;
  qrCode?: string;
  qrCodeBase64?: string;
  paidAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface ICheckBudgetAccess {
  hasAccess: boolean;
  isPremium: boolean;
  isPaid: boolean;
  reason?: string;
}

export interface IUnlockBudgetPaymentInput {
  budgetId: number;
  payer: {
    first_name: string;
    last_name: string;
    email: string;
    identification: {
      type: string;
      number: string;
    };
    address: {
      zip_code: string;
      street_name: string;
      street_number: string;
      neighborhood: string;
      city: string;
      federal_unit: string;
    };
  };
}

export interface IUnlockBudgetPaymentOutput {
  unlockedBudgetId: number;
  paymentId: string;
  amount: number;
  qrCode: string;
  qrCodeBase64: string;
  status: string;
}
