import { useLocation, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { ProfessionalService } from "../services/ProfessionalService";
import { CheckoutState, CheckoutPremiumOutput } from "../interfaces";
import { Copy, Check, Loader } from "lucide-react";
import Swal from "sweetalert2";

function Checkout() {
  const location = useLocation();
  const navigate = useNavigate();
  const [checkoutData, setCheckoutData] = useState<CheckoutPremiumOutput | null>(null);
  const [loading, setLoading] = useState(true);
  const [copied, setCopied] = useState(false);

  // Dados vindos do registro com tipagem
  const state = location.state as CheckoutState | null;

  useEffect(() => {
    if (!state || !state.userId) {
      Swal.fire({
        icon: "error",
        title: "Erro",
        text: "Dados de checkout não encontrados",
      });
      navigate("/register");
      return;
    }

    // Chamar API de checkout
    const createCheckout = async () => {
      try {
        const response = await ProfessionalService.checkoutUserPremium({
          userId: state.userId,
          payer: state.payer,
        });

        if (response.data) {
          setCheckoutData(response.data);
        }
      } catch (error: any) {
        console.error("Erro ao criar checkout:", error);
        Swal.fire({
          icon: "error",
          title: "Erro ao processar pagamento",
          text: error.response?.data?.error || "Tente novamente mais tarde",
        });
      } finally {
        setLoading(false);
      }
    };

    createCheckout();
  }, [state, navigate]);

  const copyToClipboard = () => {
    if (checkoutData?.qr_code) {
      navigator.clipboard.writeText(checkoutData.qr_code);
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
    <div className="min-h-screen bg-gray-50 p-4 py-12">
      <div className="max-w-2xl mx-auto bg-white rounded-lg shadow-lg overflow-hidden">
        {/* Header */}
        <div className="bg-gradient-to-r from-blue-600 to-indigo-600 p-6 text-white">
          <h1 className="text-2xl font-bold mb-2">🚀 Pagamento Premium</h1>
          <p className="text-blue-100">
            Último passo para ativar seus benefícios exclusivos!
          </p>
        </div>

        {/* Content */}
        <div className="p-8">
          {/* Plano Info */}
          <div className="mb-8">
            <div className="bg-blue-50 border border-blue-200 rounded-lg p-4">
              <div className="flex items-center justify-between">
                <div>
                  <h2 className="text-lg font-semibold text-gray-900">
                    Plano Premium Mensal
                  </h2>
                  <p className="text-sm text-gray-600">
                    Assinatura recorrente
                  </p>
                </div>
                <div className="text-right">
                  <p className="text-3xl font-bold text-blue-600">
                    R$ {state?.planPrice?.toFixed(2)}
                  </p>
                  <p className="text-sm text-gray-500">/mês</p>
                </div>
              </div>
            </div>
          </div>

          {/* Payment Status */}
          <div className="mb-6">
            <div className="flex items-center justify-between bg-yellow-50 border border-yellow-200 rounded-lg p-4">
              <div className="flex items-center">
                <div className="w-3 h-3 bg-yellow-500 rounded-full animate-pulse mr-3"></div>
                <p className="text-sm font-medium text-yellow-800">
                  Aguardando pagamento
                </p>
              </div>
              <p className="text-xs text-gray-600">ID: {checkoutData.paymentId}</p>
            </div>
          </div>

          {/* QR Code Section */}
          <div className="border-t pt-6">
            <h3 className="font-semibold text-gray-900 mb-4 text-center">
              Escaneie o QR Code com o app do seu banco
            </h3>

            {checkoutData.qr_code_base64 && (
              <div className="flex flex-col items-center">
                {/* QR Code Image */}
                <div className="bg-white p-4 rounded-lg shadow-md mb-6">
                  <img
                    src={`data:image/png;base64,${checkoutData.qr_code_base64}`}
                    alt="QR Code PIX"
                    className="w-64 h-64"
                  />
                </div>

                {/* Divider */}
                <div className="flex items-center w-full mb-4">
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
                        {checkoutData.qr_code}
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
          <div className="mt-8 bg-gray-50 rounded-lg p-4">
            <h4 className="font-semibold text-gray-900 mb-3">
              Como pagar com PIX:
            </h4>
            <ol className="space-y-2 text-sm text-gray-600">
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
                  Pronto! Seu plano será ativado automaticamente após a
                  confirmação
                </span>
              </li>
            </ol>
          </div>

          {/* Footer Info */}
          <div className="mt-6 text-center">
            <p className="text-xs text-gray-500">
              O pagamento PIX é processado instantaneamente.
              <br />
              Após a confirmação, você receberá um e-mail de boas-vindas.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Checkout;
