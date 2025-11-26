import React, { useState } from "react";
import { Briefcase, ArrowRight } from "lucide-react";
import { JobService } from "../services/JobService";
import Swal from "sweetalert2";

function RegisterJob() {
  const [form, setForm] = useState({
    cargo: "",
    contratacao: "",
    salario: "",
    local: "",
    descricao: "",
    horario: "",
    requisitos: "",
    beneficios: "",
    contato: "",
    empresa: "",
    quantidade_vagas: 1,
  });

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;
    setForm((prev) => ({
      ...prev,
      [name]:
        name === "quantidade_vagas" ? Number(value || 0) : value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const response = await JobService.saveJob(form);
      if (response.status === 200 || response.status === 201) {
        Swal.fire({
          icon: "success",
          title: "Vaga cadastrada!",
          text: "Sua vaga foi enviada para o Balcão de vagas.",
        });
        // limpa form
        setForm({
          cargo: "",
          contratacao: "",
          salario: "",
          local: "",
          descricao: "",
          horario: "",
          requisitos: "",
          beneficios: "",
          contato: "",
          empresa: "",
          quantidade_vagas: 1,
        });
      }
    } catch (error) {
      Swal.fire({
        icon: "error",
        title: "Erro ao salvar vaga",
        text: "Tente novamente mais tarde.",
      });
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center px-4 py-12">
      <div className="max-w-2xl w-full bg-white rounded-lg shadow-md p-8">
        <div className="flex items-center gap-3 mb-6">
          <Briefcase className="w-8 h-8 text-blue-600" />
          <div>
            <h1 className="text-2xl font-bold text-gray-900">
              Cadastrar vaga
            </h1>
            <p className="text-gray-500 text-sm">
              Qualquer pessoa pode anunciar vagas. Os profissionais
              Premium visualizam todas no Balcão de vagas.
            </p>
          </div>
        </div>

        <form onSubmit={handleSubmit} className="space-y-4">
          {/* Linha 1 */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">
                Cargo
              </label>
              <input
                name="cargo"
                value={form.cargo}
                onChange={handleChange}
                className="mt-1 w-full border rounded-lg px-3 py-2 text-sm"
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700">
                Contratação
              </label>
              <input
                name="contratacao"
                value={form.contratacao}
                onChange={handleChange}
                placeholder="CLT, PJ, Estágio..."
                className="mt-1 w-full border rounded-lg px-3 py-2 text-sm"
              />
            </div>
          </div>

          {/* Linha 2 */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">
                Salário
              </label>
              <input
                name="salario"
                value={form.salario}
                onChange={handleChange}
                placeholder="R$ 2.000,00"
                className="mt-1 w-full border rounded-lg px-3 py-2 text-sm"
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700">
                Quantidade de vagas
              </label>
              <input
                name="quantidade_vagas"
                type="number"
                min={1}
                value={form.quantidade_vagas}
                onChange={handleChange}
                className="mt-1 w-full border rounded-lg px-3 py-2 text-sm"
              />
            </div>
          </div>

          {/* Linha 3 */}
          <div>
            <label className="block text-sm font-medium text-gray-700">
              Local
            </label>
            <input
              name="local"
              value={form.local}
              onChange={handleChange}
              placeholder="Cidade/Estado ou Remoto"
              className="mt-1 w-full border rounded-lg px-3 py-2 text-sm"
            />
          </div>

          {/* Linha 4 */}
          <div>
            <label className="block text-sm font-medium text-gray-700">
              Horário
            </label>
            <input
              name="horario"
              value={form.horario}
              onChange={handleChange}
              placeholder="Ex: Segunda a sexta, 8h às 17h"
              className="mt-1 w-full border rounded-lg px-3 py-2 text-sm"
            />
          </div>

          {/* Descrição */}
          <div>
            <label className="block text-sm font-medium text-gray-700">
              Descrição
            </label>
            <textarea
              name="descricao"
              value={form.descricao}
              onChange={handleChange}
              rows={3}
              className="mt-1 w-full border rounded-lg px-3 py-2 text-sm"
            />
          </div>

          {/* Requisitos */}
          <div>
            <label className="block text-sm font-medium text-gray-700">
              Requisitos
            </label>
            <textarea
              name="requisitos"
              value={form.requisitos}
              onChange={handleChange}
              rows={3}
              className="mt-1 w-full border rounded-lg px-3 py-2 text-sm"
            />
          </div>

          {/* Benefícios */}
          <div>
            <label className="block text-sm font-medium text-gray-700">
              Benefícios
            </label>
            <textarea
              name="beneficios"
              value={form.beneficios}
              onChange={handleChange}
              rows={3}
              className="mt-1 w-full border rounded-lg px-3 py-2 text-sm"
            />
          </div>

          {/* Contato + Empresa */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">
                Contato
              </label>
              <input
                name="contato"
                value={form.contato}
                onChange={handleChange}
                placeholder="WhatsApp, e-mail, etc."
                className="mt-1 w-full border rounded-lg px-3 py-2 text-sm"
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700">
                Empresa
              </label>
              <input
                name="empresa"
                value={form.empresa}
                onChange={handleChange}
                className="mt-1 w-full border rounded-lg px-3 py-2 text-sm"
              />
            </div>
          </div>

          <button
            type="submit"
            className="mt-6 w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg flex items-center justify-center gap-2"
          >
            Cadastrar vaga
            <ArrowRight className="w-4 h-4" />
          </button>
        </form>
      </div>
    </div>
  );
}

export default RegisterJob;
