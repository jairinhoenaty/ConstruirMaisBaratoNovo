import { ArrowLeft } from "lucide-react"
import { useLocation, useNavigate } from "react-router-dom"
import PrivacyPolicyContent from "../components/PrivacyPolicyContent"

interface PrivacyPolicyLocationState {
  returnTo?: string
  returnLabel?: string
  formDraftKey?: string
  returnState?: unknown
  restoreFormDraft?: string
}

interface PrivacyPolicyProps {
  onBack?: () => void
}

function PrivacyPolicy({ onBack }: PrivacyPolicyProps) {
  const navigate = useNavigate()
  const location = useLocation()
  const navigationState = location.state as PrivacyPolicyLocationState | null
  const returnTo = navigationState?.returnTo
  const returnLabel =
    navigationState?.returnLabel ?? "Voltar para página inicial"

  const handleBack = () => {
    if (onBack) {
      onBack()
      return
    }
    if (returnTo) {
      navigate(returnTo, {
        state: {
          ...((navigationState?.returnState as object) ?? {}),
          restoreFormDraft: navigationState?.formDraftKey,
        },
      })
      return
    }
    navigate("/")
  }

  return (
    <div className="min-h-screen bg-gray-50 py-12">
      <div className="max-w-4xl mx-auto px-4">
        <button
          type="button"
          onClick={handleBack}
          className="flex items-center gap-2 text-gray-600 hover:text-blue-600 mb-8 transition-colors"
          aria-label={returnLabel}
        >
          <ArrowLeft className="w-5 h-5" />
          {returnLabel}
        </button>

        <div className="bg-white rounded-xl shadow-xl p-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-2">
            Política de Privacidade e Termos de Uso
          </h1>

          <PrivacyPolicyContent />
        </div>
      </div>
    </div>
  );
}

export default PrivacyPolicy;
