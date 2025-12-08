import React, { useEffect, useState } from "react";
import { JobService } from "../services/JobService";
import { Briefcase, Mail, Phone, MapPin, Check, Clock, PlusCircle } from "lucide-react";

const JobsPanelPremium: React.FC = () => {
  const [activeTab, setActiveTab] = useState("profile");
  const [jobs, setJobs] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
   const profile = localStorage.getItem("profile");

async function loadJobs() {
  try {
    setLoading(true);
    const res = await JobService.getApproved();
    console.log("Vagas aprovadas", res.data); // debug
    setJobs(res.data);  // <<< AQUI
  } catch (err) {
    console.error("Erro ao carregar vagas", err);
  } finally {
    setLoading(false);
  }
}

  useEffect(() => {
    loadJobs();
  }, []);

  if (loading) return <div>Carregando...</div>;

  if (jobs.length === 0)
    return <div className="p-6">Nenhuma vaga aprovada.</div>;

  <pre>{JSON.stringify(jobs, null, 2)}</pre>

  return (
    <div className="bg-white rounded-lg shadow-md p-8">

      <div className="flex items-center gap-3 mb-6">
        <Briefcase className="w-7 h-7 text-blue-600" />
        <h2 className="text-2xl font-bold text-gray-900">Balcão de Vagas</h2>
      </div>

      <div className="space-y-6">
          {jobs.map((job) => (
            <div
              key={job.id}
              className="border rounded-lg p-5 shadow-sm hover:shadow-md transition bg-white"
            >
              <div className="flex justify-between items-start">
                <div>
                  <h3 className="text-xl font-bold text-gray-900">{job.title}</h3>
                  <p className="text-gray-700">{job.description}</p>
                </div>
                {job.approved && (
                  <Check className="w-5 h-5 text-green-600" />
                )}
              </div>

              <div className="mt-3 space-y-2 text-gray-700 text-sm">

                <p><strong>Tipo de contratação:</strong> {job.hiring_type || "-"}</p>

                <p>
                  <strong>Salário:</strong>{" "}
                  {job.salary ? `R$ ${job.salary}` : "-"}{" "}
                  {job.salary_type && `(${job.salary_type})`}
                </p>

                <p><strong>Horário:</strong> {job.schedule || "-"}</p>

                <p><strong>Requisitos:</strong> {job.requirements || "-"}</p>

                <p><strong>Benefícios:</strong> {job.benefits || "-"}</p>

                <p className="flex items-center gap-1">
                    <MapPin className="w-4 h-4" />
                    {job.city?.name ? `${job.city.name} - ${job.city.uf || job.city.state?.uf || ""}` : "-"}
                </p>


                <p className="flex items-center gap-1">
                  <Mail className="w-4 h-4" />
                  {job.contact_email || "-"}
                </p>

                <p className="flex items-center gap-1">
                  <Mail className="w-4 h-4" />
                  {job.profession?.name ? `${job.profession.name}`: "-"} 
                </p>

                <p className="flex items-center gap-1">
                  <Phone className="w-4 h-4" />
                  {job.contact_phone || "-"}
                </p>

                <p><strong>Quantidade de vagas:</strong> {job.openings_quantity}</p>

                <p className="flex items-center gap-1">
                  <Clock className="w-4 h-4" />
                  Criada em: {job.created_at ? new Date(job.published_at).toLocaleDateString() : "-"}
                </p>
              </div>
            </div>

        ))}
      </div>
    </div>
  );
};

export default JobsPanelPremium;
