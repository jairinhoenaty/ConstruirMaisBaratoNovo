import { useLocation, useNavigate } from "react-router-dom";
import { useState, useEffect, useRef } from "react";
import { BudgetService } from "../services/Budget";
import { IUnlockBudgetPaymentOutput } from "../interfaces/IUnlockedBudget";
import { Copy, Check, Loader, Lock, CheckCircle2, AlertCircle } from "lucide-react";
import Swal from "sweetalert2";
import { usePaymentStatus } from "../hooks/usePaymentStatus";

interface CheckoutUnlockState {
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

function CheckoutUnlock() {
  const location = useLocation();
  const navigate = useNavigate();
  const [checkoutData, setCheckoutData] = useState<IUnlockBudgetPaymentOutput | null>(null);
  const [loading, setLoading] = useState(true);
  const [copied, setCopied] = useState(false);

  const hasCalledApi = useRef(false);

  const state = location.state as CheckoutUnlockState | null;

  const paymentStatus = usePaymentStatus(checkoutData?.statusToken, () => {
    setTimeout(() => navigate("/professional-panel"), 3000);
  });

  useEffect(() => {
    if (!state || !state.budgetId) {
      Swal.fire({
        icon: "error",
        title: "Erro",
        text: "Dados de pagamento não encontrados",
      });
      navigate("/quotes");
      return;
    }

    if (hasCalledApi.current) {
      return;
    }

    hasCalledApi.current = true;

    const createUnlockPayment = async () => {
      try {
        const response = await BudgetService.unlockBudget(state.budgetId, {
          budgetId: state.budgetId,
          payer: state.payer,
        });

        if (response.data) {
          setCheckoutData(response.data);
        }
      } catch (error: any) {
        console.error("Erro ao criar pagamento:", error);
        Swal.fire({
          icon: "error",
          title: "Erro ao processar pagamento",
          text: error.response?.data?.error || "Tente novamente mais tarde",
        });
        navigate("/quotes");
      } finally {
        setLoading(false);
      }
    };

    createUnlockPayment();
  }, []); // eslint-disable-line react-hooks/exhaustive-deps

  const copyToClipboard = () => {
    if (checkoutData?.qrCode) {
      navigator.clipboard.writeText(checkoutData.qrCode);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <Loader className="w-12 h-12 text-blue-600 animate-spin mx-auto mb-4" />
          <p className="text-gray-600">Gerando QR Code...</p>
        </div>
      </div>
    );
  }

  if (!checkoutData) {
    return null;
  }

  return (
    <div className="min-h-screen bg-gray-50 p-4 py-6">
      <div className="max-w-2xl mx-auto bg-white rounded-lg shadow-lg overflow-hidden">
        {/* Header */}
        <div className="bg-gradient-to-r from-blue-600 to-indigo-600 p-4 text-white">
          <div className="flex items-center gap-3 mb-1">
            <Lock className="w-6 h-6" />
            <h1 className="text-xl font-bold">Desbloqueio de Orçamento</h1>
          </div>
          <p className="text-sm text-blue-100">
            Pague para acessar os dados completos deste orçamento
          </p>
        </div>

        {/* Content */}
        <div className="p-6">
          {/* Payment Info */}
          <div className="mb-6">
            <div className="bg-blue-50 border border-blue-200 rounded-lg p-4">
              <div className="flex items-center justify-between">
                <div>
                  <h2 className="text-lg font-semibold text-gray-900">
                    Desbloqueio Avulso de Orçamento
                  </h2>
                  <p className="text-sm text-gray-600">
                    Pagamento único - Orçamento #{state?.budgetId}
                  </p>
                </div>
                <div className="text-right">
                  <p className="text-3xl font-bold text-blue-600">
                    R$ {checkoutData.amount.toFixed(2)}
                  </p>
                  <p className="text-sm text-gray-500">pagamento único</p>
                </div>
              </div>
            </div>
          </div>

          {/* Payment Status */}
          <div className="mb-6">
            {paymentStatus === "approved" ? (
              <div className="flex items-center justify-between bg-green-50 border border-green-200 rounded-lg p-4">
                <div className="flex items-center">
                  <CheckCircle2 className="w-5 h-5 text-green-600 mr-3" />
                  <div>
                    <p className="text-sm font-medium text-green-800">
                      Pagamento confirmado!
                    </p>
                    <p className="text-xs text-green-700">
                      Orçamento desbloqueado. Redirecionando...
                    </p>
                  </div>
                </div>
              </div>
            ) : paymentStatus === "timeout" ? (
              <div className="flex items-center justify-between bg-gray-50 border border-gray-200 rounded-lg p-4">
                <div className="flex items-center">
                  <AlertCircle className="w-5 h-5 text-gray-500 mr-3" />
                  <div>
                    <p className="text-sm font-medium text-gray-800">
                      Não identificamos seu pagamento
                    </p>
                    <p className="text-xs text-gray-600">
                      Se já pagou, atualize a página. O código PIX continua válido.
                    </p>
                  </div>
                </div>
              </div>
            ) : (
              <div className="flex items-center justify-between bg-yellow-50 border border-yellow-200 rounded-lg p-4">
                <div className="flex items-center">
                  <div className="w-3 h-3 bg-yellow-500 rounded-full animate-pulse mr-3"></div>
                  <p className="text-sm font-medium text-yellow-800">
                    Aguardando pagamento
                  </p>
                </div>
              </div>
            )}
          </div>

          {/* Benefits */}
          <div className="mb-6 bg-green-50 border border-green-200 rounded-lg p-4">
            <h3 className="font-semibold text-gray-900 mb-2 flex items-center gap-2">
              <Check className="w-5 h-5 text-green-600" />
              Após o pagamento você terá acesso a:
            </h3>
            <ul className="space-y-1 text-sm text-gray-700">
              <li className="flex items-center gap-2">
                <span className="text-green-600">✓</span>
                <span>Nome completo do cliente</span>
              </li>
              <li className="flex items-center gap-2">
                <span className="text-green-600">✓</span>
                <span>Telefone/WhatsApp para contato</span>
              </li>
              <li className="flex items-center gap-2">
                <span className="text-green-600">✓</span>
                <span>Email do cliente</span>
              </li>
              <li className="flex items-center gap-2">
                <span className="text-green-600">✓</span>
                <span>Endereço completo do serviço</span>
              </li>
            </ul>
          </div>

          {/* QR Code Section */}
          <div className="border-t pt-6">
            <h3 className="font-semibold text-gray-900 mb-4 text-center">
              Escaneie o QR Code com o app do seu banco
            </h3>

            {checkoutData.qrCodeBase64 && (
              <div className="flex flex-col items-center">
                {/* QR Code Image */}
                <div className="bg-white p-3 rounded-lg shadow-md mb-4">
                  <img
                    src={`data:image/png;base64,${checkoutData.qrCodeBase64}`}
                    alt="QR Code PIX"
                    className="w-48 h-48"
                  />
                </div>

                {/* Divider */}
                <div className="flex items-center w-full mb-3">
                  <div className="flex-1 border-t border-gray-300"></div>
                  <span className="px-4 text-sm text-gray-500">ou</span>
                  <div className="flex-1 border-t border-gray-300"></div>
                </div>

                {/* Copy Code */}
                <div className="w-full">
                  <p className="text-sm text-gray-600 mb-2 text-center">
                    Copie o código PIX Copia e Cola:
                  </p>
                  <div className="bg-gray-50 border border-gray-200 rounded-lg p-4">
                    <div className="flex items-center justify-between">
                      <code className="text-xs break-all text-gray-800 flex-1 pr-4">
                        {checkoutData.qrCode}
                      </code>
                      <button
                        onClick={copyToClipboard}
                        className={`flex-shrink-0 p-2 rounded-md transition-colors ${
                          copied
                            ? "bg-green-100 text-green-600"
                            : "bg-blue-100 text-blue-600 hover:bg-blue-200"
                        }`}
                        title={copied ? "Copiado!" : "Copiar código"}
                      >
                        {copied ? (
                          <Check className="w-5 h-5" />
                        ) : (
                          <Copy className="w-5 h-5" />
                        )}
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            )}
          </div>

          {/* Instructions */}
          <div className="mt-6 bg-gray-50 rounded-lg p-4">
            <h4 className="font-semibold text-gray-900 mb-2">
              Como pagar com PIX:
            </h4>
            <ol className="space-y-1.5 text-sm text-gray-600">
              <li className="flex">
                <span className="font-semibold mr-2">1.</span>
                <span>
                  Abra o app do seu banco e escolha a opção de pagar com PIX
                </span>
              </li>
              <li className="flex">
                <span className="font-semibold mr-2">2.</span>
                <span>
                  Escaneie o QR Code acima ou copie o código PIX Copia e Cola
                </span>
              </li>
              <li className="flex">
                <span className="font-semibold mr-2">3.</span>
                <span>Confirme o pagamento no app do seu banco</span>
              </li>
              <li className="flex">
                <span className="font-semibold mr-2">4.</span>
                <span>
                  Pronto! O orçamento será desbloqueado automaticamente após a
                  confirmação
                </span>
              </li>
            </ol>
          </div>

          {/* Footer Info */}
          <div className="mt-4 text-center">
            <p className="text-xs text-gray-500">
              O pagamento PIX é processado instantaneamente.
              <br />
              Retorne à lista de orçamentos para visualizar os dados completos.
            </p>
          </div>

          {/* Back Button */}
          <div className="mt-6 text-center">
            <button
              onClick={() => navigate("/quotes")}
              className="text-blue-600 hover:text-blue-700 font-medium text-sm"
            >
              ← Voltar para orçamentos
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default CheckoutUnlock;
