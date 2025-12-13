// Interfaces para integração com Mercado Pago

export interface PayerIdentification {
  type: string;
  number: string;
}

export interface PayerAddress {
  zip_code: string;
  street_name: string;
  street_number: string;
  neighborhood: string;
  city: string;
  federal_unit: string;
}

export interface Payer {
  first_name: string;
  last_name: string;
  email: string;
  identification: PayerIdentification;
  address: PayerAddress;
}

export interface CheckoutPremiumInput {
  userId: number;
  payer: Payer;
}

export interface CheckoutPremiumOutput {
  paymentId: number;
  amount: number;
  qr_code: string;
  qr_code_base64: string;
  status: string;
}

export interface CheckoutState {
  userId: number;
  userName: string;
  userEmail: string;
  planId: number;
  userType: 'professional' | 'store';
  payer: Payer;
  isUpgrade?: boolean;
  returnUrl?: string;
}
