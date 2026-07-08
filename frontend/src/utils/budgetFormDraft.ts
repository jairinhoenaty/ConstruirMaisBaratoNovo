export const BUDGET_FORM_DRAFT_KEYS = {
  searchResults: "budget-form-search-results",
  quoteProducts: "budget-form-quote-products",
} as const

export type BudgetFormDraftKey =
  (typeof BUDGET_FORM_DRAFT_KEYS)[keyof typeof BUDGET_FORM_DRAFT_KEYS]

export const saveBudgetFormDraft = <T>(key: BudgetFormDraftKey, data: T): void => {
  sessionStorage.setItem(key, JSON.stringify(data))
}

export const loadBudgetFormDraft = <T>(key: BudgetFormDraftKey): T | null => {
  const raw = sessionStorage.getItem(key)
  if (!raw) return null
  try {
    return JSON.parse(raw) as T
  } catch {
    return null
  }
}

export const clearBudgetFormDraft = (key: BudgetFormDraftKey): void => {
  sessionStorage.removeItem(key)
}
