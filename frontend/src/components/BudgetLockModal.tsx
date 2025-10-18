import { X, Lock, Crown, Zap } from "lucide-react";
import { useNavigate } from "react-router-dom";

interface BudgetLockModalProps {
  isOpen: boolean;
  onClose: () => void;
  budgetId: number;
  onPayUnlock: () => void;
}

function BudgetLockModal({ isOpen, onClose, budgetId, onPayUnlock }: BudgetLockModalProps) {
  const navigate = useNavigate();

  if (!isOpen) return null;

  const handleUpgradePremium = () => {
    // Redireciona para página de registro premium ou upgrade
    navigate("/register", { state: { upgradeToPremium: true } });
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div className="bg-white rounded-lg shadow-xl max-w-md w-full">
        {/* Header */}
        <div className="flex items-center justify-between p-6 border-b">
          <div className="flex items-center gap-2">
            <Lock className="w-6 h-6 text-orange-600" />
            <h2 className="text-xl font-bold text-gray-900">
              Orçamento Bloqueado
            </h2>
          </div>
          <button
            onClick={onClose}
            className="text-gray-400 hover:text-gray-600 transition-colors"
          >
            <X className="w-6 h-6" />
          </button>
        </div>

        {/* Content */}
        <div className="p-6 space-y-6">
          <p className="text-gray-600">
            Este orçamento contém informações de contato do cliente. Para visualizar
            todos os detalhes, você precisa:
          </p>

          {/* Option 1: Pay for this budget */}
          <div className="border-2 border-blue-200 rounded-lg p-4 hover:border-blue-400 transition-colors">
            <div className="flex items-start gap-3">
              <Zap className="w-6 h-6 text-blue-600 flex-shrink-0 mt-1" />
              <div className="flex-1">
                <h3 className="font-semibold text-gray-900 mb-1">
                  Desbloquear este orçamento
                </h3>
                <p className="text-sm text-gray-600 mb-3">
                  Pague uma única vez e tenha acesso completo a este orçamento
                </p>
                <div className="flex items-center justify-between">
                  <span className="text-2xl font-bold text-blue-600">R$ 30,00</span>
                  <button
                    onClick={onPayUnlock}
                    className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors font-medium"
                  >
                    Pagar Agora
                  </button>
                </div>
              </div>
            </div>
          </div>

          {/* Divider */}
          <div className="flex items-center gap-3">
            <div className="flex-1 border-t border-gray-300"></div>
            <span className="text-sm text-gray-500 font-medium">OU</span>
            <div className="flex-1 border-t border-gray-300"></div>
          </div>

          {/* Option 2: Become Premium (Highlighted) */}
          <div className="bg-gradient-to-br from-orange-50 to-orange-100 border-2 border-orange-300 rounded-lg p-4 relative overflow-hidden">
            {/* Badge */}
            <div className="absolute top-0 right-0 bg-orange-500 text-white text-xs font-bold px-3 py-1 rounded-bl-lg">
              RECOMENDADO
            </div>

            <div className="flex items-start gap-3 mt-2">
              <Crown className="w-6 h-6 text-orange-600 flex-shrink-0 mt-1" />
              <div className="flex-1">
                <h3 className="font-bold text-gray-900 mb-1 text-lg">
                  Torne-se Premium
                </h3>
                <p className="text-sm text-gray-700 mb-3">
                  Acesso <strong>ilimitado</strong> a todos os orçamentos sem pagar por cada um
                </p>

                {/* Benefits */}
                <ul className="space-y-1 mb-4 text-sm text-gray-700">
                  <li className="flex items-center gap-2">
                    <span className="text-green-600">✓</span>
                    <span>Ver todos os orçamentos</span>
                  </li>
                  <li className="flex items-center gap-2">
                    <span className="text-green-600">✓</span>
                    <span>Destaque no perfil</span>
                  </li>
                  <li className="flex items-center gap-2">
                    <span className="text-green-600">✓</span>
                    <span>Suporte prioritário</span>
                  </li>
                </ul>

                <div className="flex items-center justify-between">
                  <div>
                    <span className="text-2xl font-bold text-orange-600">R$ 19,90</span>
                    <span className="text-sm text-gray-600">/mês</span>
                  </div>
                  <button
                    onClick={handleUpgradePremium}
                    className="px-6 py-2 bg-gradient-to-r from-orange-500 to-orange-600 text-white rounded-lg hover:from-orange-600 hover:to-orange-700 transition-all font-bold shadow-md hover:shadow-lg"
                  >
                    Ser Premium
                  </button>
                </div>
              </div>
            </div>
          </div>

          {/* Footer note */}
          <p className="text-xs text-gray-500 text-center">
            Ao desbloquear o orçamento ou assinar o plano Premium, você concorda com nossos termos de serviço
          </p>
        </div>
      </div>
    </div>
  );
}

export default BudgetLockModal;
