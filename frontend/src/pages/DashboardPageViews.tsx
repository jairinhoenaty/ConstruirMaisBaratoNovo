import { useEffect, useMemo, useState } from "react"
import { BarChart3, Calendar } from "lucide-react"
import { IPageView, PageViewService } from "../services/PageViewService"

const formatViewDate = (viewDate: string): string => {
  if (!viewDate) return "-"
  const [year, month, day] = viewDate.split("-")
  if (!year || !month || !day) return viewDate
  return `${day}/${month}/${year}`
}

function DashboardPageViews() {
  const [pageViews, setPageViews] = useState<IPageView[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [dateFrom, setDateFrom] = useState("")
  const [dateTo, setDateTo] = useState("")

  useEffect(() => {
    const fetchPageViews = async () => {
      try {
        const response = await PageViewService.getPageViews()
        if (response.status === 200) {
          const sorted = [...response.data].sort((a, b) => {
            if (a.viewDate !== b.viewDate) {
              return b.viewDate.localeCompare(a.viewDate)
            }
            return b.count - a.count
          })
          setPageViews(sorted)
        }
      } catch (error) {
        console.error("Erro ao buscar acessos às páginas:", error)
      } finally {
        setIsLoading(false)
      }
    }

    fetchPageViews()
  }, [])

  const filteredPageViews = useMemo(() => {
    return pageViews.filter((pageView) => {
      if (!pageView.viewDate) return false
      if (dateFrom && pageView.viewDate < dateFrom) return false
      if (dateTo && pageView.viewDate > dateTo) return false
      return true
    })
  }, [pageViews, dateFrom, dateTo])

  const hasActiveFilter = dateFrom !== "" || dateTo !== ""

  const handleClearFilter = () => {
    setDateFrom("")
    setDateTo("")
  }

  if (isLoading) {
    return <p className="text-gray-600">Carregando acessos...</p>
  }

  return (
    <div className="bg-white rounded-lg shadow-md p-6">
      <div className="flex items-center gap-3 mb-6">
        <div className="p-3 bg-indigo-100 rounded-lg">
          <BarChart3 className="w-6 h-6 text-indigo-600" />
        </div>
        <div>
          <h3 className="text-lg font-semibold text-gray-900">
            Acessos por Página
          </h3>
          <p className="text-sm text-gray-600">
            Contagem diária por rota. Acessos de admin logado não são
            contabilizados.
          </p>
        </div>
      </div>

      <div className="mb-6 flex flex-col gap-3 sm:flex-row sm:items-end sm:flex-wrap">
        <div>
          <label
            htmlFor="page-view-date-from"
            className="mb-1 block text-sm font-medium text-gray-700"
          >
            Data inicial
          </label>
          <div className="relative">
            <Calendar className="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-gray-400" />
            <input
              id="page-view-date-from"
              type="date"
              value={dateFrom}
              onChange={(event) => setDateFrom(event.target.value)}
              max={dateTo || undefined}
              className="block w-full rounded-md border border-gray-300 py-2 pl-9 pr-3 text-sm focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500"
              aria-label="Filtrar a partir da data"
            />
          </div>
        </div>

        <div>
          <label
            htmlFor="page-view-date-to"
            className="mb-1 block text-sm font-medium text-gray-700"
          >
            Data final
          </label>
          <div className="relative">
            <Calendar className="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-gray-400" />
            <input
              id="page-view-date-to"
              type="date"
              value={dateTo}
              onChange={(event) => setDateTo(event.target.value)}
              min={dateFrom || undefined}
              className="block w-full rounded-md border border-gray-300 py-2 pl-9 pr-3 text-sm focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500"
              aria-label="Filtrar até a data"
            />
          </div>
        </div>

        {hasActiveFilter && (
          <button
            type="button"
            onClick={handleClearFilter}
            className="rounded-md border border-gray-300 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50"
          >
            Limpar filtro
          </button>
        )}
      </div>

      {pageViews.length === 0 ? (
        <p className="text-gray-500">Nenhum acesso registrado ainda.</p>
      ) : filteredPageViews.length === 0 ? (
        <p className="text-gray-500">
          Nenhum acesso encontrado para o período selecionado.
        </p>
      ) : (
        <div className="overflow-x-auto">
          <table className="w-full table-auto text-sm">
            <thead className="bg-gray-50">
              <tr className="border-b border-gray-200">
                <th className="px-4 py-3 text-left font-medium text-gray-500">
                  Data
                </th>
                <th className="px-4 py-3 text-left font-medium text-gray-500">
                  Página
                </th>
                <th className="px-4 py-3 text-right font-medium text-gray-500">
                  Acessos
                </th>
              </tr>
            </thead>
            <tbody>
              {filteredPageViews.map((pageView) => (
                <tr key={pageView.id} className="border-b border-gray-100">
                  <td className="px-4 py-3 text-gray-700">
                    {formatViewDate(pageView.viewDate)}
                  </td>
                  <td className="px-4 py-3 text-gray-900 font-mono">
                    {pageView.path}
                  </td>
                  <td className="px-4 py-3 text-right text-gray-900 font-semibold">
                    {pageView.count.toLocaleString("pt-BR")}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  )
}

export default DashboardPageViews
