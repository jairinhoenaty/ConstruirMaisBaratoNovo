import { useEffect, useState } from "react"
import { BarChart3 } from "lucide-react"
import { IPageView, PageViewService } from "../services/PageViewService"

function DashboardPageViews() {
  const [pageViews, setPageViews] = useState<IPageView[]>([])
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    const fetchPageViews = async () => {
      try {
        const response = await PageViewService.getPageViews()
        if (response.status === 200) {
          const sorted = [...response.data].sort((a, b) => b.count - a.count)
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
            Páginas com mais visitas podem indicar gargalos ou oportunidades.
            Acessos de admin logado não são contabilizados.
          </p>
        </div>
      </div>

      {pageViews.length === 0 ? (
        <p className="text-gray-500">Nenhum acesso registrado ainda.</p>
      ) : (
        <div className="overflow-x-auto">
          <table className="w-full table-auto text-sm">
            <thead className="bg-gray-50">
              <tr className="border-b border-gray-200">
                <th className="px-4 py-3 text-left font-medium text-gray-500">
                  Página
                </th>
                <th className="px-4 py-3 text-right font-medium text-gray-500">
                  Acessos
                </th>
              </tr>
            </thead>
            <tbody>
              {pageViews.map((pageView) => (
                <tr key={pageView.id} className="border-b border-gray-100">
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
