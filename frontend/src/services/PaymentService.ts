import { PaymentStatusOutput } from "../interfaces/ICheckout";
import ApiPublica from "../providers/ApiPublica";

// Consulta se o PIX já foi confirmado. O backend reconsulta o MercadoPago
// quando o pagamento ainda está pendente, então esse endpoint destrava o fluxo
// mesmo se o webhook não chegar.
//
// A chave é o token opaco devolvido no checkout. O id do pagamento no
// MercadoPago nunca chega ao cliente: a rota é pública, e um identificador
// adivinhável permitiria ler pagamentos alheios.
const getPaymentStatus = (statusToken: string) =>
  ApiPublica.get<PaymentStatusOutput>(`/payment/status/${statusToken}`);

export const PaymentService = {
  getPaymentStatus,
};
