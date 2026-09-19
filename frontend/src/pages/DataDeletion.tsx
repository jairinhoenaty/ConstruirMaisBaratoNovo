import { FormEvent, useState } from "react";
import { CheckCircle, Trash2 } from "lucide-react";
import { AccountDeletionService } from "../services/AccountDeletionService";

const DataDeletion = () => {
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    telephone: "",
    reason: "",
  });

  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState("");

  const handleChange = (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = event.target;

    setFormData((previous) => ({
      ...previous,
      [name]: value,
    }));
  };

  const handleSubmit = async (event: FormEvent) => {
    event.preventDefault();

    setError("");

    if (!formData.name.trim() || !formData.email.trim()) {
      setError("Preencha o nome e o e-mail cadastrado.");
      return;
    }

    try {
      setLoading(true);

      await AccountDeletionService.create({
        name: formData.name.trim(),
        email: formData.email.trim(),
        telephone: formData.telephone.trim(),
        reason: formData.reason.trim(),
      });

      setSuccess(true);
    } catch (err: any) {
      console.error("Erro ao solicitar exclusão:", err);

      const message =
        err?.response?.data?.error ||
        "Não foi possível enviar sua solicitação. Tente novamente.";

      setError(message);
    } finally {
      setLoading(false);
    }
  };

  if (success) {
    return (
      <main className="min-h-screen bg-gray-50 flex items-center justify-center px-4 py-12">
        <div className="bg-white shadow-md rounded-xl p-8 max-w-xl w-full text-center">
          <CheckCircle className="w-16 h-16 text-green-600 mx-auto mb-5" />

          <h1 className="text-2xl font-bold text-gray-900 mb-4">
            Solicitação recebida
          </h1>

          <p className="text-gray-600">
            Recebemos sua solicitação de exclusão. Nossa equipe irá
            processá-la.
          </p>

          <p className="text-sm text-gray-500 mt-4">
            A exclusão será realizada de acordo com nossa Política de
            Privacidade e com as obrigações legais aplicáveis.
          </p>
        </div>
      </main>
    );
  }

  return (
    <main className="min-h-screen bg-gray-50 px-4 py-12">
      <div className="max-w-2xl mx-auto">
        <div className="bg-white rounded-xl shadow-md p-6 sm:p-10">

          <div className="flex items-start gap-3 mb-5">
            <Trash2 className="w-8 h-8 text-red-600 flex-shrink-0 mt-1" />

            <h1 className="text-2xl sm:text-3xl font-bold text-gray-900">
              Exclusão de Conta e Dados — Hassis Conecta
            </h1>
          </div>

          <p className="text-gray-600 mb-4">
            Você pode solicitar a exclusão da sua conta e dos dados pessoais
            associados ao Hassis Conecta.
          </p>

          <p className="text-gray-600 mb-8">
            Preencha o formulário abaixo utilizando o e-mail cadastrado em sua
            conta. Após o recebimento, nossa equipe analisará e processará sua
            solicitação.
          </p>

          <form onSubmit={handleSubmit} className="space-y-5">

            <div>
              <label
                htmlFor="name"
                className="block text-sm font-medium text-gray-700 mb-1"
              >
                Nome *
              </label>

              <input
                id="name"
                name="name"
                type="text"
                value={formData.name}
                onChange={handleChange}
                required
                autoComplete="name"
                className="w-full border border-gray-300 rounded-lg px-4 py-3 focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label
                htmlFor="email"
                className="block text-sm font-medium text-gray-700 mb-1"
              >
                E-mail cadastrado *
              </label>

              <input
                id="email"
                name="email"
                type="email"
                value={formData.email}
                onChange={handleChange}
                required
                autoComplete="email"
                className="w-full border border-gray-300 rounded-lg px-4 py-3 focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label
                htmlFor="telephone"
                className="block text-sm font-medium text-gray-700 mb-1"
              >
                Telefone (opcional)
              </label>

              <input
                id="telephone"
                name="telephone"
                type="tel"
                value={formData.telephone}
                onChange={handleChange}
                autoComplete="tel"
                className="w-full border border-gray-300 rounded-lg px-4 py-3 focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label
                htmlFor="reason"
                className="block text-sm font-medium text-gray-700 mb-1"
              >
                Motivo da solicitação (opcional)
              </label>

              <textarea
                id="reason"
                name="reason"
                value={formData.reason}
                onChange={handleChange}
                rows={4}
                className="w-full border border-gray-300 rounded-lg px-4 py-3 resize-none focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            {error && (
              <div
                role="alert"
                className="bg-red-50 border border-red-200 text-red-700 rounded-lg p-3"
              >
                {error}
              </div>
            )}

            <button
              type="submit"
              disabled={loading}
              className="w-full bg-red-600 hover:bg-red-700 disabled:bg-gray-400 text-white font-medium py-3 px-4 rounded-lg transition-colors"
            >
              {loading
                ? "Enviando solicitação..."
                : "Solicitar exclusão da conta"}
            </button>

          </form>

          <p className="text-xs text-gray-500 mt-6">
            Determinadas informações poderão ser mantidas quando sua
            conservação for necessária para o cumprimento de obrigações legais
            ou regulatórias aplicáveis.
          </p>

        </div>
      </div>
    </main>
  );
};

export default DataDeletion;