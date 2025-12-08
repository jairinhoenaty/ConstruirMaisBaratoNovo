import React, { useEffect, useState } from "react";
import { JobService } from "../services/JobService";
import { ProfessionService } from "../services/ProfessionService";
import { CityService } from "../services/CityService";
import ErrorAlert from "../components/ErrorAlert";
import { useNavigate } from "react-router-dom";
import { ArrowLeft } from "lucide-react";


import {
  Briefcase,
  DollarSign,
  MapPin,
  Mail,
  Phone,
  FileText,
  Clock,
  ClipboardList,
  CheckCircle,
  Layers,
  Hash,
} from "lucide-react";

// Estados locais (igual Register.tsx)
import { states } from "../data";

const RegisterJob: React.FC = () => {
const navigate = useNavigate();

  // 🔵 Formulário do job
  const [form, setForm] = useState({
    title: "",
    hiring_type: "",
    salary: "",
    salary_type: "",
    location: "",
    description: "",
    schedule: "",
    requirements: "",
    benefits: "",
    contact_email: "",
    contact_phone: "",
    openings_quantity: "",
    city_id: "",
    profession_id: "",
    state: "",
    status: "aberto",
  });

  // 🟣 Listas carregadas do banco
  const [professions, setProfessions] = useState<any[]>([]);
  const [citiesByState, setCitiesByState] = useState<any[]>([]);

  // Feedback
  const [success, setSuccess] = useState("");
  const [error, setError] = useState("");

  // 🔵 Atualiza formulário
  function handleChange(e: any) {
    setForm({ ...form, [e.target.name]: e.target.value });
  }

  // 🔥 Busca profissões ao iniciar
  useEffect(() => {
    async function loadProfessions() {
      try {
        const res = await ProfessionService.getProfessionsPublic();
        setProfessions(res.data.professions || res.data || []);
      } catch (err) {
        console.error("Erro ao carregar profissões", err);
      }
    }
    loadProfessions();
  }, []);

  // 🔥 Busca cidades quando o estado mudar
  useEffect(() => {
    async function loadCities() {
      if (!form.state) return;

      try {
        const res = await CityService.citiesByStatePublic({
          uf: form.state,
        });

        setCitiesByState(res.data.cities || res.data || []);
      } catch (err) {
        console.error("Erro ao carregar cidades", err);
      }
    }

    loadCities();
  }, [form.state]);

  

  // Enviar vaga
  async function submit() {
    if (!form.title ||
  !form.hiring_type ||
  !form.salary ||
  !form.salary_type ||
  !form.location ||
  !form.description ||
  !form.schedule ||
  !form.requirements ||
  !form.benefits ||
  !form.contact_email ||
  !form.contact_phone ||
  !form.openings_quantity ||
  !form.profession_id ||
  !form.city_id ||
  !form.state) {
      setError("Preencha todos os campos obrigatórios!");
      return;
    }

    try {
      const payload = {
        ...form,
        salary: Number(form.salary),
        openings_quantity: Number(form.openings_quantity),
        profession_id: Number(form.profession_id),
        city_id: Number(form.city_id),
      };

      await JobService.create(payload);
      


      setSuccess("Vaga cadastrada com sucesso!");
      setError("");
      setTimeout(() => {
      navigate("/"); // home
      }, 1500);
      
    } catch (err) {
      console.error(err);
      setError("Erro ao cadastrar vaga.");
    }
  }

  // 🔹 Input padrão igual ao Register.tsx
  const InputWrapper = ({ label, name, type = "text", Icon, placeholder }: any) => (
    <div>
      <label className="block text-sm font-medium text-gray-700">{label}</label>
      <div className="mt-1 relative rounded-md shadow-sm">
        <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
          <Icon className="h-5 w-5 text-gray-400" />
        </div>
        <input
          id={name}
          name={name}
          type={type}
          value={(form as any)[name]}
          onChange={handleChange}
          className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          placeholder={placeholder}
        />
      </div>
    </div>
  );

  const TextAreaWrapper = ({ label, name, Icon }: any) => (
    <div>
      <label className="block text-sm font-medium text-gray-700">{label}</label>
      <div className="mt-1 relative rounded-md shadow-sm">
        <div className="absolute inset-y-0 left-0 pl-3 flex items-start pt-3 pointer-events-none">
          <Icon className="h-5 w-5 text-gray-400" />
        </div>
        <textarea
          id={name}
          name={name}
          value={(form as any)[name]}
          onChange={handleChange}
          rows={4}
          className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
        ></textarea>
      </div>
    </div>
  );

  return (
    <div className="min-h-screen bg-gray-50 flex flex-col justify-center py-12 px-4 sm:px-6 lg:px-8">

      {/* Botão de voltar */}
      <button
        onClick={() => navigate(-1)}
        className="flex items-center text-blue-600 hover:text-blue-800 mb-4 w-fit"
      >
        <ArrowLeft className="w-6 h-6 mr-1" />
        Voltar
      </button>
      
      <div className="sm:mx-auto sm:w-full sm:max-w-2xl">
        <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900 flex items-center justify-center gap-2">
          <Briefcase className="w-8 h-8 text-blue-600" />
          Cadastrar Vaga
        </h2>
      </div>

      <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-2xl">
        <div className="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">

          {error && <ErrorAlert message={error} onClose={() => setError("")} />}
          {success && (
            <p className="bg-green-100 text-green-700 p-3 rounded mb-4 text-center font-medium">
              {success}
            </p>
          )}

          <div className="space-y-6">

            {/* PROFISSÃO */}
            <div>
              <label className="block text-sm font-medium text-gray-700">Cargo / Profissão</label>
              <div className="mt-1 relative rounded-md shadow-sm">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <Briefcase className="h-5 w-5 text-gray-400" />
                </div>

                <select
                  name="profession_id"
                  value={form.profession_id}
                  onChange={handleChange}
                  className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm 
                             focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="">Selecione a profissão</option>

                  {professions.map((p) => (
                    <option key={p.id} value={p.id}>{p.name}</option>
                  ))}
                </select>
              </div>
            </div>

            {/* Inputs */}
            <InputWrapper 
              label="Cargo / Título da vaga" 
              name="title" 
              Icon={Briefcase} 
              placeholder="Ex: Pedreiro, Eletricista, Auxiliar..." 
            />
            <InputWrapper label="Tipo de Contratação" name="hiring_type" Icon={Layers} />
            <InputWrapper label="Salário" name="salary" type="number" Icon={DollarSign} />
            <InputWrapper label="Tipo de Salário (Mensal, Diário...)" name="salary_type" Icon={Hash} />
            <InputWrapper label="Local" name="location" Icon={MapPin} />

            <TextAreaWrapper label="Descrição" name="description" Icon={FileText} />
            <TextAreaWrapper label="Horário" name="schedule" Icon={Clock} />
            <TextAreaWrapper label="Requisitos" name="requirements" Icon={ClipboardList} />
            <TextAreaWrapper label="Benefícios" name="benefits" Icon={CheckCircle} />

            <InputWrapper label="Email para contato" name="contact_email" Icon={Mail} />
            <InputWrapper label="Telefone para contato" name="contact_phone" Icon={Phone} />

            <InputWrapper label="Quantidade de vagas" name="openings_quantity" Icon={Hash} type="number" />

            {/* ESTADO */}
            <div>
              <label className="block text-sm font-medium text-gray-700">Estado</label>
              <div className="mt-1 relative rounded-md shadow-sm">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <MapPin className="h-5 w-5 text-gray-400" />
                </div>

                <select
                  name="state"
                  value={(form as any).state}
                  onChange={handleChange}
                  className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm 
                             focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="">Selecione o estado</option>
                  {states.map((state) => (
                    <option key={state.id} value={state.id}>
                      {state.name}
                    </option>
                  ))}
                </select>
              </div>
            </div>

            {/* CIDADE */}
            <div>
              <label className="block text-sm font-medium text-gray-700">Cidade</label>
              <div className="mt-1 relative rounded-md shadow-sm">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <MapPin className="h-5 w-5 text-gray-400" />
                </div>

                <select
                  name="city_id"
                  value={form.city_id}
                  onChange={handleChange}
                  disabled={!form.state}
                  className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm 
                             focus:outline-none focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100"
                >
                  <option value="">Selecione a cidade</option>

                  {citiesByState.map((city: any) => (
                    <option key={city.id} value={city.id}>{city.name}</option>
                  ))}
                </select>
              </div>
            </div>

            {/* Botão */}
            <button
              onClick={submit}
              className="w-full flex justify-center items-center py-3 border border-transparent rounded-md 
                         shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700"
            >
              Cadastrar Vaga
            </button>

          </div>

        </div>
      </div>

    </div>
  );
};

export default RegisterJob;
