import { useLocation, useNavigate } from "react-router-dom";
import { useState, useEffect, useRef } from "react";
import { ProfessionalService } from "../services/ProfessionalService";
import { StoreService } from "../services/StoreService";
import { PlanService } from "../services/PlanService";
import { CheckoutState, CheckoutPremiumOutput } from "../interfaces";
import { Plan } from "../interfaces/IPlan";
import { Copy, Check, Loader, CheckCircle } from "lucide-react";
import Swal from "sweetalert2";

 // coloquei aqui pois fica mais facil de regular 
const SHOW_MANUAL_BUTTON_AFTER_MS = 30_000;
const POLLING_INTERVAL_MS = 5_000;

function Checkout() {
  const location = useLocation();
  const navigate = useNavigate();

  const [checkoutData, setCheckoutData] = useState<CheckoutPremiumOutput | null>(null);
  const [plan, setPlan] = useState<Plan | null>(null);
  const [loading, setLoading] = useState(true);
  const [copied, setCopied] = useState(false);
  const [paymentStatus, setPaymentStatus] = useState<"pending" | "approved" | "failed">("pending");
  const [showManualButton, setShowManualButton] = useState(false);

  const hasCalledApi = useRef(false);
  const pollingRef = useRef<ReturnType<typeof setInterval> | null>(null);
  const manualButtonTimerRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  const state = location.state as CheckoutState | null;
  const isUpgrade = (state as any)?.isUpgrade || false;

  const stopPolling = () => {
    if (pollingRef.current) {
      clearInterval(pollingRef.current);
      pollingRef.current = null;
    }
  };

  const startPolling = (paymentId: number) => {
    pollingRef.current = setInterval(async () => {
      try {
        const res = await ProfessionalService.getPaymentStatus(paymentId);
        const status = res.data?.status;

        if (status === "approved") {
          stopPolling();
          setPaymentStatus("approved");
          Swal.fire({
            position: "center",
            icon: "success",
            title: "Pagamento confirmado! 🎉",
            text: "Seu plano premium foi ativado. Redirecionando para o login...",
            showConfirmButton: false,
            timer: 3000,
          });
          setTimeout(() => navigate("/login"), 3000);
        } else if (
          status === "rejected" ||
          status === "cancelled" ||
          status === "canceled"
        ) {
          stopPolling();
          setPaymentStatus("failed");
        }
      } catch {
        // Silencioso — continua polling mesmo com erros pontuais
      }
    }, POLLING_INTERVAL_MS);
  };

  useEffect(() => {
    if (!state || !state.userId) {
      Swal.fire({ icon: "error", title: "Erro", text: "Dados de checkout não encontrados" });
      navigate("/register");
      return;
    }

    if (hasCalledApi.current) return;
    hasCalledApi.current = true;

    const createCheckout = async () => {
      try {
        if (state.userType) {
          const planResponse = await PlanService.getPlanByUserType(state.userType);
          if (planResponse.data) setPlan(planResponse.data);
        }

        const checkoutService =
          state.userType === "store"
            ? StoreService.checkoutStorePremium
            : ProfessionalService.checkoutUserPremium;

        const response = await checkoutService({ userId: state.userId, payer: state.payer });

        if (response.data) {
          setCheckoutData(response.data);
          startPolling(response.data.paymentId);
          manualButtonTimerRef.current = setTimeout(
            () => setShowManualButton(true),
            SHOW_MANUAL_BUTTON_AFTER_MS
          );
        }
      } catch (error: any) {
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

    return () => {
      stopPolling();
      if (manualButtonTimerRef.current) clearTimeout(manualButtonTimerRef.current);
    };
  }, []);

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

  if (!checkoutData) return null;

  if (paymentStatus === "approved") {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center p-4">
        <div className="max-w-sm w-full bg-white rounded-2xl shadow-lg p-8 text-center">
          <CheckCircle className="w-16 h-16 text-green-500 mx-auto mb-4" />
          <h2 className="text-xl font-bold text-gray-900 mb-2">Pagamento confirmado!</h2>
          <p className="text-gray-600 text-sm mb-6">
            Seu plano premium foi ativado. Redirecionando para o login...
          </p>
          <Loader className="w-6 h-6 text-blue-600 animate-spin mx-auto" />
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50 p-4 py-6">
      <div className="max-w-2xl mx-auto bg-white rounded-lg shadow-lg overflow-hidden">
        <div className="bg-gradient-to-r from-blue-600 to-indigo-600 p-4 text-white">
          <h1 className="text-xl font-bold mb-1">
            {isUpgrade ? "Upgrade para Premium" : "Pagamento"}
          </h1>
          <p className="text-sm text-blue-100">
            {isUpgrade
              ? "Complete seu upgrade e desbloqueie todos os benefícios!"
              : "Último passo para ativar seus benefícios exclusivos!"}
          </p>
        </div>

        <div className="p-6">

          <div className="mb-6">
            <div className="bg-blue-50 border border-blue-200 rounded-lg p-4">
              <div className="flex items-center justify-between">
                <div>
                  <h2 className="text-lg font-semibold text-[#FF6B35]">
                    {plan?.name || "Carregando..."}
                  </h2>
                  <p className="text-sm text-gray-600">
                    {plan?.description || "Assinatura recorrente"}
                  </p>
                </div>
                <div className="text-right">
                  <p className="text-3xl font-bold text-blue-600">
                    {plan ? `R$ ${plan.price.toFixed(2).replace(".", ",")}` : "..."}
                  </p>
                  <p className="text-sm text-gray-500">/mês</p>
                </div>
              </div>
            </div>
          </div>

          <div className="mb-6">
            {paymentStatus === "pending" && (
              <div className="flex items-center justify-between bg-yellow-50 border border-yellow-200 rounded-lg p-4">
                <div className="flex items-center">
                  <div className="w-3 h-3 bg-yellow-500 rounded-full animate-pulse mr-3" />
                  <p className="text-sm font-medium text-yellow-800">
                    Aguardando confirmação do pagamento...
                  </p>
                </div>
                <p className="text-xs text-gray-500 ml-2">ID: {checkoutData.paymentId}</p>
              </div>
            )}
            {paymentStatus === "failed" && (
              <div className="flex items-center bg-red-50 border border-red-200 rounded-lg p-4">
                <div className="w-3 h-3 bg-red-500 rounded-full mr-3 flex-shrink-0" />
                <p className="text-sm font-medium text-red-700">
                  Pagamento não confirmado. Tente novamente ou contate o suporte.
                </p>
              </div>
            )}
          </div>

          <div className="border-t pt-6">
            <h3 className="font-semibold text-gray-900 mb-4 text-center">
              Escaneie o QR Code com o app do seu banco
            </h3>

            {checkoutData.qr_code_base64 && (
              <div className="flex flex-col items-center">
                <div className="bg-white p-3 rounded-lg shadow-md mb-4">
                  <img
                    src={`data:image/png;base64,${checkoutData.qr_code_base64}`}
                    alt="QR Code PIX"
                    className="w-48 h-48"
                  />
                </div>

                <div className="flex items-center w-full mb-3">
                  <div className="flex-1 border-t border-gray-300" />
                  <span className="px-4 text-sm text-gray-500">ou</span>
                  <div className="flex-1 border-t border-gray-300" />
                </div>

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
                        {copied ? <Check className="w-5 h-5" /> : <Copy className="w-5 h-5" />}
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            )}
          </div>

          <div className="mt-6 bg-gray-50 rounded-lg p-4">
            <h4 className="font-semibold text-gray-900 mb-2">Como pagar com PIX:</h4>
            <ol className="space-y-1.5 text-sm text-gray-600">
              <li className="flex">
                <span className="font-semibold mr-2">1.</span>
                <span>Abra o app do seu banco e escolha a opção de pagar com PIX</span>
              </li>
              <li className="flex">
                <span className="font-semibold mr-2">2.</span>
                <span>Escaneie o QR Code acima ou copie o código PIX Copia e Cola</span>
              </li>
              <li className="flex">
                <span className="font-semibold mr-2">3.</span>
                <span>Confirme o pagamento no app do seu banco</span>
              </li>
              <li className="flex">
                <span className="font-semibold mr-2">4.</span>
                <span>
                  Pronto! Você será redirecionado automaticamente assim que o pagamento for confirmado
                </span>
              </li>
            </ol>
          </div>

          <div className="mt-4 text-center">
            <p className="text-xs text-gray-500">
              O pagamento PIX é processado instantaneamente.
              <br />
              Após a confirmação, você receberá um e-mail de boas-vindas.
            </p>
          </div>

          {/* Botão manual — aparece como fallback após 30s */}
          {showManualButton && paymentStatus === "pending" && (
            <div className="mt-6 border-t border-gray-100 pt-5 text-center">
              <p className="text-xs text-gray-400 mb-3">
                Já realizou o pagamento e não foi redirecionado automaticamente?
              </p>
              <button
                onClick={() => navigate("/login")}
                className="inline-flex items-center px-4 py-2 border border-blue-400 rounded-md text-xs font-medium text-blue-600 hover:bg-blue-50 transition-colors"
              >
                Ir para o login
              </button>
            </div>
          )}

          {/* Link pequeno para voltar ao cadastro básico */}
          <div className="mt-5 text-center">
            <button
              onClick={() => navigate("/register")}
              className="text-xs text-gray-400 hover:text-gray-500 underline underline-offset-2"
            >
              Prefiro continuar como profissional básico
            </button>
          </div>

        </div>
      </div>
    </div>
  );
}

export default Checkout;