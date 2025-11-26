import React from "react";
import { Briefcase } from "lucide-react";
import Pagination from "../components/Pagination";
import { JobService } from "../services/JobService";
import { IJob } from "../interfaces/IJob";

function JobsPanel() {
  const [jobs, setJobs] = React.useState<IJob[]>([]);
  const [page, setPage] = React.useState(1);
  const [totalPage, setTotalPage] = React.useState(1);
  const [isUpdate, setIsUpdate] = React.useState(false);

  const limit = 5;

  React.useEffect(() => {
    const fetchData = async () => {
      const offset = (page - 1) * limit;
      const response = await JobService.getJobs(limit, offset);

      if (response.status === 200) {
        const { jobs, total } = response.data;
        setJobs(jobs);
        setTotalPage(Math.ceil(total / limit));
      }
    };
    fetchData();
  }, [page, isUpdate]);

  return (
    <div className="bg-white rounded-lg shadow-md p-8">
      <div className="flex items-center gap-3 mb-8">
        <Briefcase className="w-8 h-8 text-blue-600" />
        <h2 className="text-2xl font-bold text-gray-900">
          Balcão de vagas
        </h2>
      </div>

      {totalPage === 0 && <h1>Nenhuma vaga cadastrada</h1>}

      {totalPage > 0 && (
        <Pagination
          currentPage={page}
          totalPages={totalPage}
          handleNextPage={() => setPage(page + 1)}
          handlePrevPage={() => setPage(page - 1)}
        />
      )}

      <div className="space-y-6 mt-6">
        {jobs.map((job) => (
          <div
            key={job.id}
            className="border border-gray-200 rounded-lg p-6 hover:border-blue-500 transition-colors"
          >
            {/* Cabeçalho – igual Mensagens */}
            <div className="flex justify-between items-center text-sm text-gray-500 mb-4">
              <span>
                📅{" "}
                {job.created_at
                  ? new Date(job.created_at).toLocaleString("pt-BR")
                  : ""}
              </span>
            </div>

            <h3 className="text-lg font-semibold text-gray-900 mb-2">
              Cargo: <span className="font-normal">{job.cargo}</span>
            </h3>

            <p className="text-sm text-gray-600 mb-4">
              Empresa: <span className="font-medium">{job.empresa}</span> •{" "}
              Local: <span className="font-medium">{job.local}</span> •
              Contratação:{" "}
              <span className="font-medium">{job.contratacao}</span>
            </p>

            {/* Bloco com 2 colunas, igual a estrutura de Mensagens */}
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-4">
              <div className="space-y-2 text-sm text-gray-700">
                <p>
                  <span className="font-semibold">Horário: </span>
                  {job.horario}
                </p>
                <p>
                  <span className="font-semibold">Salário: </span>
                  {job.salario}
                </p>
                <p>
                  <span className="font-semibold">Contato: </span>
                  {job.contato}
                </p>
              </div>

              <div className="space-y-2 text-sm text-gray-700">
                <p>
                  <span className="font-semibold">
                    Quantidade de vagas:{" "}
                  </span>
                  {job.quantidade_vagas}
                </p>
                <p>
                  <span className="font-semibold">Benefícios: </span>
                  {job.beneficios}
                </p>
              </div>
            </div>

            {/* Descrição / Requisitos */}
            <div className="border-t border-gray-100 pt-4 space-y-3 text-sm text-gray-700">
              <div>
                <span className="font-semibold block mb-1">
                  Descrição:
                </span>
                <p>{job.descricao}</p>
              </div>

              <div>
                <span className="font-semibold block mb-1">
                  Requisitos:
                </span>
                <p>{job.requisitos}</p>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default JobsPanel;
