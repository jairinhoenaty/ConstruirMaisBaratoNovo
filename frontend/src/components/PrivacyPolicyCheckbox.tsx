import type { ChangeEvent, MouseEvent } from "react"
import { useLocation, useNavigate } from "react-router-dom"
import {
  BudgetFormDraftKey,
  saveBudgetFormDraft,
} from "../utils/budgetFormDraft"

interface PrivacyPolicyCheckboxProps {
  checked: boolean
  onChange: (checked: boolean) => void
  id?: string
  formDraftKey?: BudgetFormDraftKey
  getFormDraft?: () => Record<string, unknown>
}

function PrivacyPolicyCheckbox({
  checked,
  onChange,
  id = "privacy-policy-accept",
  formDraftKey,
  getFormDraft,
}: PrivacyPolicyCheckboxProps) {
  const navigate = useNavigate()
  const location = useLocation()

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    onChange(event.target.checked)
  }

  const handlePrivacyClick = (event: MouseEvent<HTMLButtonElement>) => {
    event.preventDefault()
    event.stopPropagation()

    if (formDraftKey && getFormDraft) {
      saveBudgetFormDraft(formDraftKey, getFormDraft())
    }

    navigate("/privacy", {
      state: {
        returnTo: `${location.pathname}${location.search}`,
        returnLabel: "Voltar ao formulário",
        formDraftKey,
        returnState: location.state,
      },
    })
  }

  return (
    <label
      htmlFor={id}
      className="flex items-start gap-3 cursor-pointer text-sm text-gray-600"
    >
      <input
        type="checkbox"
        id={id}
        checked={checked}
        onChange={handleChange}
        required
        className="mt-1 w-4 h-4 rounded border-gray-300 text-blue-600 focus:ring-blue-500"
        aria-label="Aceitar política de privacidade"
      />
      <span>
        Li e concordo com a{" "}
        <button
          type="button"
          onClick={handlePrivacyClick}
          className="text-blue-600 hover:text-blue-800 underline"
        >
          Política de Privacidade
        </button>
      </span>
    </label>
  )
}

export default PrivacyPolicyCheckbox
