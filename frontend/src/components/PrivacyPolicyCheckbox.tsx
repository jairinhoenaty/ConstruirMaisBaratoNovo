import { Link } from "react-router-dom"
import type { ChangeEvent } from "react"

interface PrivacyPolicyCheckboxProps {
  checked: boolean
  onChange: (checked: boolean) => void
  id?: string
}

function PrivacyPolicyCheckbox({
  checked,
  onChange,
  id = "privacy-policy-accept",
}: PrivacyPolicyCheckboxProps) {
  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    onChange(event.target.checked)
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
        <Link
          to="/privacy"
          target="_blank"
          rel="noopener noreferrer"
          className="text-blue-600 hover:text-blue-800 underline"
        >
          Política de Privacidade
        </Link>
      </span>
    </label>
  )
}

export default PrivacyPolicyCheckbox
