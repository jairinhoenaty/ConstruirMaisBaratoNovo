import { useEffect, useRef, useState } from "react"
import { PaymentService } from "../services/PaymentService"

// Intervalo entre consultas enquanto o QR Code está na tela.
const POLL_INTERVAL_MS = 5000
// Depois disso paramos de consultar: o usuário provavelmente fechou o app do
// banco ou desistiu. O PIX continua válido — basta recarregar a página.
const POLL_TIMEOUT_MS = 15 * 60 * 1000

export type PaymentPollStatus = "idle" | "waiting" | "approved" | "timeout"

/**
 * Consulta periodicamente se o pagamento foi confirmado.
 *
 * O backend reconsulta o MercadoPago quando o registro ainda está pendente,
 * então o fluxo destrava mesmo que o webhook não chegue.
 *
 * @param statusToken token opaco do checkout; vazio enquanto o QR Code não foi gerado
 * @param onApproved chamado uma única vez, assim que o pagamento é aprovado
 */
export const usePaymentStatus = (
  statusToken: string | null | undefined,
  onApproved?: () => void
) => {
  const [status, setStatus] = useState<PaymentPollStatus>("idle")

  // Mantém a referência mais recente sem reiniciar o polling a cada render.
  const onApprovedRef = useRef(onApproved)
  onApprovedRef.current = onApproved

  useEffect(() => {
    if (!statusToken) return

    let cancelled = false
    const startedAt = Date.now()
    setStatus("waiting")

    const check = async () => {
      try {
        const response = await PaymentService.getPaymentStatus(statusToken)
        if (cancelled) return

        if (response.data?.approved) {
          setStatus("approved")
          clearInterval(timer)
          onApprovedRef.current?.()
          return
        }

        if (Date.now() - startedAt > POLL_TIMEOUT_MS) {
          setStatus("timeout")
          clearInterval(timer)
        }
      } catch {
        // 404 (pagamento ainda não gravado) e falhas de rede são transitórios:
        // seguimos tentando até o timeout.
        if (!cancelled && Date.now() - startedAt > POLL_TIMEOUT_MS) {
          setStatus("timeout")
          clearInterval(timer)
        }
      }
    }

    const timer = setInterval(check, POLL_INTERVAL_MS)
    check()

    return () => {
      cancelled = true
      clearInterval(timer)
    }
  }, [statusToken])

  return status
}
