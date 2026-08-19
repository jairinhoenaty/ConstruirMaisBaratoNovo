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

// O id do pagamento no MercadoPago não é devolvido ao cliente: o
// acompanhamento é feito pelo statusToken, que é opaco e não permite deduzir
// nem enumerar pagamentos de outras pessoas.
export interface CheckoutPremiumOutput {
  statusToken: string;
  amount: number;
  qr_code: string;
  qr_code_base64: string;
  status: string;
}

// Resposta de GET /publica/payment/status/:statusToken
export interface PaymentStatusOutput {
  status: string;
  approved: boolean;
  amount: number;
  paidAt?: string;
  expiresAt?: string;
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
