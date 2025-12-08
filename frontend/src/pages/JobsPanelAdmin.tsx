import React, { useEffect, useState } from "react";
import { JobService } from "../services/JobService";
import {
  Briefcase,
  Trash2,
  Clock,
  MapPin,
  Mail,
  Phone,
} from "lucide-react";
import Pagination from "../components/Pagination";
import Swal from "sweetalert2";

const JobsPanelAdmin: React.FC = () => {
  const [jobs, setJobs] = useState<any[]>([]);
  const [filteredJobs, setFilteredJobs] = useState<any[]>([]);
  const [tab, setTab] = useState<"pending" | "approved">("pending");
  const [page, setPage] = useState(1);
  const [total, setTotal] = useState(1);
  const [loading, setLoading] = useState(true);

  const limit = 10; // por página

  async function loadJobs() {
    try {
      setLoading(true);

      const offset = (page - 1) * limit;

      const res = await JobService.getAll(limit, offset);

      setJobs(res.data.jobs || []);
      setTotal(Math.ceil(res.data.total / limit));

    } catch (error) {
      console.error("Erro ao carregar vagas", error);
      Swal.fire("Erro", "Não foi possível carregar as vagas.", "error");
    } finally {
      setLoading(false);
    }
  }

  function filterByTab(all: any[], tabSelected: "pending" | "approved") {
    if (tabSelected === "pending") {
      setFilteredJobs(all.filter((j) => !j.approved));
    } else {
      setFilteredJobs(all.filter((j) => j.approved));
    }
  }

  useEffect(() => {
    loadJobs();
  }, [page]);

  useEffect(() => {
    filterByTab(jobs, tab);
  }, [tab, jobs]);

  async function handleApprove(job: any, newValue: boolean) {
    try {
      await JobService.approve(job.id, newValue);

      Swal.fire({
        icon: "success",
        text: newValue ? "Vaga aprovada!" : "Vaga reprovada!",
        timer: 1500,
        showConfirmButton: false,
      });

      loadJobs();
    } catch (err) {
      Swal.fire("Erro", "Não foi possível atualizar a vaga.", "error");
    }
  }

  async function handleDelete(id: number) {
    if (!window.confirm("Tem certeza que deseja excluir esta vaga?")) return;

    try {
      await JobService.delete(id);
      Swal.fire({
        icon: "success",
        text: "Vaga excluída!",
        timer: 1500,
        showConfirmButton: false,
      });
      loadJobs();
    } catch (err) {
      Swal.fire("Erro", "Falha ao excluir vaga.", "error");
    }
  }

  return (
    <div className="bg-white rounded-lg shadow-md p-8">

      <div className="flex items-center gap-3 mb-6">
        <Briefcase className="w-7 h-7 text-blue-600" />
        <h2 className="text-2xl font-bold text-gray-900">Balcão de Vagas (Admin)</h2>
      </div>

      {/* Tabs */}
      <div className="flex gap-4 mb-6">
        <button
          onClick={() => setTab("pending")}
          className={`px-4 py-2 rounded-lg ${
            tab === "pending"
              ? "bg-blue-600 text-white"
              : "bg-gray-100 text-gray-700"
          }`}
        >
          Vagas Pendentes
        </button>

        <button
          onClick={() => setTab("approved")}
          className={`px-4 py-2 rounded-lg ${
            tab === "approved"
              ? "bg-blue-600 text-white"
              : "bg-gray-100 text-gray-700"
          }`}
        >
          Vagas Aprovadas
        </button>
      </div>

      {/* Paginação */}
      <Pagination
        currentPage={page}
        totalPages={total}
        handleNextPage={() => setPage((p) => p + 1)}
        handlePrevPage={() => setPage((p) => p - 1)}
      />

      {/* Lista */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mt-6">
        {filteredJobs.map((job) => (
          <div
            key={job.id}
            className="border rounded-lg p-5 shadow-sm bg-white"
          >
            <h3 className="text-lg font-bold text-gray-900">{job.title}</h3>

            <p className="text-gray-700 mb-2">{job.description}</p>

            <p className="text-sm flex items-center gap-1 text-gray-700">
              <MapPin className="w-4 h-4" />
              {job.city?.name} - {job.city?.uf}
            </p>

            <p className="text-sm flex items-center gap-1 text-gray-700">
              <Mail className="w-4 h-4" />
              {job.contact_email}
            </p>

            <p className="text-sm flex items-center gap-1 text-gray-700">
              <Phone className="w-4 h-4" />
              {job.contact_phone}
            </p>

            {/* Aprovar */}
            <label className="flex items-center mt-4 gap-2 cursor-pointer">
              <input
                type="checkbox"
                checked={job.approved}
                onChange={(e) => handleApprove(job, e.target.checked)}
              />
              <span>Aprovar</span>
            </label>

            <button
              onClick={() => handleDelete(job.id)}
              className="text-red-600 mt-4 flex items-center gap-2 hover:bg-red-50 p-2 rounded-lg"
            >
              <Trash2 className="w-5 h-5" />
              Excluir
            </button>
          </div>
        ))}
      </div>
    </div>
  );
};

export default JobsPanelAdmin;
