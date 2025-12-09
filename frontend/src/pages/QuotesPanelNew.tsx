import React, { useState } from "react";
import {
  Calendar,
  User,
  Phone,
  MessageSquare,
  HardHat,
  MapPin,
  Trash2,
  Mail,
  Lock,
  Eye,
} from "lucide-react";
import { useNavigate } from "react-router-dom";
import { states } from "../data";
import { BudgetService } from "../services/Budget";
import Pagination from "../components/Pagination";
import { IBudget } from "../interfaces";
import { ICheckBudgetAccess } from "../interfaces/IUnlockedBudget";
import BudgetLockModal from "../components/BudgetLockModal";
import Swal from "sweetalert2";

interface Quote extends IBudget {
  budgetAccess?: ICheckBudgetAccess;
}

function QuotesPanel() {
  const navigate = useNavigate();
  const [quotes, setQuotes] = useState<Quote[]>([]);
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(0);
  const [isUpdate, setIsUpdate] = useState(false);
  const [selectedBudgetId, setSelectedBudgetId] = useState<number | null>(null);
  const [showLockModal, setShowLockModal] = useState(false);
  const profile = localStorage.getItem("profile");
  const [budgetFilter, setBudgetFilter] = useState<
    "all" | "professional" | "store"
  >("all");

  const [professionals, setProfessionals] = useState([]);
  const citiesByState = [];

  React.useEffect(() => {
    console.log("useEffect");
    const fetchData = async () => {
      const limit = 10;
      let json;
      let total = 0;
      const id = parseInt(localStorage.getItem("id") || "");
      const post_id = parseInt(localStorage.getItem("post_id") || "");
      let response;

      if (profile == "admin") {
        response = await BudgetService.getBudgetsAll(limit, (page - 1) * limit);
        console.log(response);
        json = response.data.budgets;
        total = response.data.total;
      } else if (profile == "store") {
        response = await BudgetService.getBudgetsbyMonth({
          month: "April",
          storeID: profile == "store" ? id : 0,
          clientID: profile == "client" ? post_id : 0,
          page: 1,
          pagesize: 10,
        });
        json = response.data;
      } else {
        response = await BudgetService.getBudgetsbyMonth({
          month: "April",
          professionalID: profile == "profissional" ? id : 0,
          clientID: profile == "client" ? post_id : 0,
          page: 1,
          pagesize: 10,
        });
        json = response.data;
      }

      console.log(json);
      setProfessionals([]);

      if (response.status == 200) {
        // Log para debug - mostrar estrutura dos dados
        if (profile == "admin" && json?.length > 0) {
          console.log("=== DEBUG: Estrutura dos orçamentos ===");
          json.slice(0, 3).forEach((quote: any, index: number) => {
            console.log(`Orçamento ${index + 1}:`, {
              id: quote.id,
              professionalsId: quote.professionalsId,
              storesId: quote.storesId,
              professionals: quote.professionals,
              stores: quote.stores,
              fullData: quote,
            });
          });
          console.log("======================================");
        }

        // Se for profissional, verificar acesso para cada orçamento
        if (
          (profile == "profissional" || profile == "store") &&
          json?.length > 0
        ) {
          const quotesWithAccess = await Promise.all(
            json.map(async (quote: IBudget) => {
              try {
                const accessResponse = await BudgetService.checkBudgetAccess(
                  quote.id
                );
                return {
                  ...quote,
                  budgetAccess: accessResponse.data,
                };
              } catch (error) {
                console.error(
                  `Erro ao verificar acesso do orçamento ${quote.id}:`,
                  error
                );
                return {
                  ...quote,
                  budgetAccess: {
                    hasAccess: true, // Em caso de erro, permitir acesso para não bloquear
                    isPremium: false,
                    isPaid: false,
                    reason: "error",
                  },
                };
              }
            })
          );
          setQuotes(quotesWithAccess);
        } else {
          setQuotes(json);
        }

        setTotalPage(Math.ceil(total / limit));
      }
    };

    fetchData();
  }, [page, isUpdate, profile]);

  const handleDeleteQuote = async (id: string) => {
    if (window.confirm("Tem certeza que deseja excluir este orçamento?")) {
      const response = await BudgetService.deleteBudget(id);
      if (response.status == 200) {
        Swal.fire({
          icon: "success",
          text: "Mensagem excluída",
          showConfirmButton: false,
          timer: 1500,
        });
        setIsUpdate(!isUpdate);
      }
    }
  };

  const handleStatusChange = async (
    budget_approved: IBudget,
    approved: boolean
  ) => {
    budget_approved.approved = !approved;
    console.log(budget_approved);

    const response = await BudgetService.saveBudget(budget_approved);
    if (response.status == 200) {
      Swal.fire({
        icon: "success",
        text: "Orçamento alterado",
        showConfirmButton: false,
        timer: 2000,
      });
      setIsUpdate(!isUpdate);
    }
  };

  const handleViewBudgetDetails = (quote: Quote) => {
    // Se não tem acesso, mostra modal de bloqueio
    if (!quote.budgetAccess?.hasAccess) {
      setSelectedBudgetId(quote.id);
      setShowLockModal(true);
    }
    // Se tem acesso, não faz nada (dados já estão visíveis)
  };

  const handlePayUnlock = () => {
    if (!selectedBudgetId) return;

    // Redireciona para checkout de desbloqueio com dados do profissional
    const professionalEmail = localStorage.getItem("email") || "";
    const professionalName = localStorage.getItem("name") || "Profissional";

    navigate("/checkout-unlock", {
      state: {
        budgetId: selectedBudgetId,
        payer: {
          first_name: professionalName.split(" ")[0],
          last_name: professionalName.split(" ").slice(1).join(" ") || "Silva",
          email: professionalEmail,
          identification: {
            type: "CPF",
            number: "00000000000",
          },
          address: {
            zip_code: "00000-000",
            street_name: "Rua Exemplo",
            street_number: "123",
            neighborhood: "Centro",
            city: "São Paulo",
            federal_unit: "SP",
          },
        },
      },
    });

    setShowLockModal(false);
  };

  const renderClientInfoPreview = (message: Quote) => {
    const hasAccess = message.budgetAccess?.hasAccess ?? true;

    return (
      <>
        <div className="border-t border-gray-100 pt-4"></div>
        <div className="space-y-4">
          <div className="flex items-center justify-between">
            <h2 className="text-lg font-semibold text-gray-700">
              Dados do Cliente
            </h2>
            {!hasAccess && (
              <span className="flex items-center gap-1 text-sm text-orange-600 font-medium">
                <Lock className="w-4 h-4" />
                Bloqueado
              </span>
            )}
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {/* Coluna 1 */}
            <div className="space-y-4">
              {budgetFilter !== "store" && (
                <div className="flex items-start gap-3">
                  <User className="w-5 h-5 mt-1 text-gray-400" />
                  <div>
                    <span className="block text-sm text-gray-500">Nome</span>
                    {hasAccess ? (
                      <p className="text-gray-900 font-medium">
                        {message.name}
                      </p>
                    ) : (
                      <p className="text-gray-400 italic">****** *******</p>
                    )}
                  </div>
                </div>
              )}

              <div className="flex items-start gap-3">
                <Mail className="w-5 h-5 mt-1 text-gray-400" />
                <div>
                  <span className="block text-sm text-gray-500">Email:</span>
                  {hasAccess ? (
                    <p className="text-gray-900 font-medium">{message.email}</p>
                  ) : (
                    <p className="text-gray-400 italic">******@*****.com</p>
                  )}
                </div>
              </div>
            </div>

            {/* Coluna 2 */}
            <div className="space-y-4">
              <div className="flex items-start gap-3">
                <Phone className="w-5 h-5 mt-1 text-gray-400" />
                <div>
                  <span className="block text-sm text-gray-500">WhatsApp:</span>
                  {hasAccess ? (
                    <p className="text-gray-900 font-medium">
                      {message.telephone}
                    </p>
                  ) : (
                    <p className="text-gray-400 italic">(XX) XXXXX-XXXX</p>
                  )}
                </div>
              </div>

              <div className="flex items-start gap-3">
                <MapPin className="w-5 h-5 mt-1 text-gray-400" />
                <div>
                  <span className="block text-sm text-gray-500">Região:</span>
                  <p className="text-gray-900 font-medium">
                    {message.city?.name || "Não informado"}
                  </p>
                </div>
              </div>
            </div>
          </div>

          {/* Unlock Button */}
          {!hasAccess && (
            <div className="mt-4">
              <button
                onClick={() => handleViewBudgetDetails(message)}
                className="w-full flex items-center justify-center gap-2 px-4 py-3 bg-gradient-to-r from-blue-600 to-blue-700 text-white rounded-lg hover:from-blue-700 hover:to-blue-800 transition-all font-medium shadow-md"
              >
                <Eye className="w-5 h-5" />
                Visualizar Orçamento Completo
              </button>
            </div>
          )}
        </div>
      </>
    );
  };

  const renderProfessionalInfo = (message: Quote) => {
    console.log("message", message);
    console.log("Professions", message.professionals?.[0]?.professions);

    message.professionals?.[0]?.professions?.map((profession) => {
      console.log(profession.name);
    });

    return (
      <>
        <div className="border-t border-gray-100 pt-4"></div>
        <div className="space-y-4">
          <h2 className="text-lg font-semibold text-gray-700">
            Dados do Profissional
          </h2>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {/* Coluna 1 */}
            <div className="space-y-4">
              <div className="flex items-start gap-3">
                <User className="w-5 h-5 mt-1 text-gray-400" />
                <div>
                  <span className="block text-sm text-gray-500">Nome</span>
                  <p className="text-gray-900 font-medium">
                    {message.professionals !== null
                      ? message.professionals?.[0]?.name
                      : message.stores?.[0]?.name}
                  </p>
                </div>
              </div>

              <div className="flex items-start gap-3">
                <Mail className="w-5 h-5 mt-1 text-gray-400" />
                <div>
                  <span className="block text-sm text-gray-500">Email:</span>
                  <p className="text-gray-900 font-medium">
                    {message.professionals !== null
                      ? message.professionals?.[0]?.email
                      : message.stores?.[0]?.email}
                  </p>
                </div>
              </div>
            </div>

            {/* Coluna 2 */}
            <div className="space-y-4">
              <div className="flex items-start gap-3">
                <Phone className="w-5 h-5 mt-1 text-gray-400" />
                <div>
                  <span className="block text-sm text-gray-500">WhatsApp:</span>
                  <p className="text-gray-900 font-medium">
                    {message.professionals !== null
                      ? message.professionals?.[0]?.telephone
                      : message.stores?.[0]?.telephone}
                  </p>
                </div>
              </div>

              <div className="flex items-start gap-3">
                <MapPin className="w-5 h-5 mt-1 text-gray-400" />
                <div>
                  <span className="block text-sm text-gray-500">
                    UF/Cidade:
                  </span>
                  <p className="text-gray-900 font-medium">
                  {message.professionals !== null
                      ? `${message.professionals?.[0]?.city?.name}/${message.professionals?.[0]?.city?.uf}`
                      : message.stores !== null
                      ? `${message.stores?.[0]?.city?.name}/${message.stores?.[0]?.city?.uf}`
                      :"Não informado"
                      }

                    {/* {message.professionals && message.professionals.length > 0
                      ? `${message.professionals[0]?.city?.name}/${message.professionals[0]?.city?.uf}`
                      : message.stores && message.stores.length > 0
                      ? `${message.stores[0]?.city?.name}/${message.stores[0]?.city?.uf}`
                      : "Não informado"} */}
                  </p>
                </div>
              </div>
            </div>

            {/* Profissões */}
            {message.professionals !== null && (
              <div className="md:col-span-2 flex flex-wrap items-center gap-2">
                <HardHat className="w-5 h-5 text-gray-400" />
                <div className="flex flex-wrap gap-2">
                  <span className="text-sm text-gray-500">Profissão:</span>
                  {message.professionals?.[0]?.professions?.map(
                    (element: any, index: number) => (
                      <span
                        key={index}
                        className="text-gray-900 bg-blue-100 text-blue-800 text-xs font-medium px-2.5 py-0.5 rounded dark:bg-blue-900 dark:text-blue-300"
                      >
                        {element.name}
                      </span>
                    )
                  )}
                </div>
              </div>
            )}
          </div>
        </div>
      </>
    );
  };

  // Filtrar orçamentos baseado no filtro selecionado (apenas para admin)
  const filteredQuotes =
    profile === "admin"
      ? quotes.filter((quote) => {
          if (budgetFilter === "all") return true;

          // Para profissionais: verificar se tem professionals ou professionalsId preenchido
          if (budgetFilter === "professional") {
            const hasProfessionalsArray =
              (quote as any).professionals &&
              Array.isArray((quote as any).professionals) &&
              (quote as any).professionals.length > 0;
            const hasProfessionalsId =
              quote.professionalsId &&
              Array.isArray(quote.professionalsId) &&
              quote.professionalsId.length > 0;
            console.log(
              `[FILTRO PROFISSIONAL] Quote ${quote.id}: hasProfessionalsArray=${hasProfessionalsArray}, hasProfessionalsId=${hasProfessionalsId}`
            );
            return hasProfessionalsArray || hasProfessionalsId;
          }

          // Para lojistas: verificar múltiplas possibilidades de campo
          if (budgetFilter === "store") {
            // Verificar arrays (plural)
            const hasStoresArray =
              (quote as any).stores &&
              Array.isArray((quote as any).stores) &&
              (quote as any).stores.length > 0;
            const hasStoresId =
              quote.storesId &&
              Array.isArray(quote.storesId) &&
              quote.storesId.length > 0;
            const hasStoresIds =
              (quote as any).storesIds &&
              Array.isArray((quote as any).storesIds) &&
              (quote as any).storesIds.length > 0;

            return hasStoresArray || hasStoresId || hasStoresIds;
          }

          return true;
        })
      : quotes;

  // Log adicional para debug do filtro de lojistas
  if (profile === "admin" && budgetFilter === "store") {
    console.log(
      `[RESULTADO FILTRO LOJISTA] Total encontrado: ${filteredQuotes.length} de ${quotes.length} orçamentos`
    );
  }

  return (
    <div className="bg-white rounded-lg shadow-md p-8">
      <div className="flex items-center justify-between mb-8">
        <div className="flex items-center gap-3">
          <Calendar className="w-8 h-8 text-blue-600" />
          <h2 className="text-2xl font-bold text-gray-900">Orçamentos</h2>
        </div>

        {/* Botões de filtro - apenas para admin */}
        {profile === "admin" && (
          <div className="flex gap-2">
            <button
              onClick={() => setBudgetFilter("all")}
              className={`px-4 py-2 rounded-lg font-medium transition-colors ${
                budgetFilter === "all"
                  ? "bg-blue-600 text-white"
                  : "bg-gray-200 text-gray-700 hover:bg-gray-300"
              }`}
            >
              Todos
            </button>
            <button
              onClick={() => setBudgetFilter("professional")}
              className={`px-4 py-2 rounded-lg font-medium transition-colors ${
                budgetFilter === "professional"
                  ? "bg-blue-600 text-white"
                  : "bg-gray-200 text-gray-700 hover:bg-gray-300"
              }`}
            >
              Profissionais
            </button>
            <button
              onClick={() => setBudgetFilter("store")}
              className={`px-4 py-2 rounded-lg font-medium transition-colors ${
                budgetFilter === "store"
                  ? "bg-blue-600 text-white"
                  : "bg-gray-200 text-gray-700 hover:bg-gray-300"
              }`}
            >
              Lojistas
            </button>
          </div>
        )}
      </div>

      {filteredQuotes.length == 0 && <div>Nenhum orçamento</div>}
      {quotes.length != 0 && totalPage !== 0 && (
        <Pagination
          currentPage={page}
          totalPages={totalPage}
          handleNextPage={() => {
            setPage(page + 1);
          }}
          handlePrevPage={() => {
            setPage(page - 1);
          }}
        />
      )}

      <div className="space-y-6">
        {filteredQuotes.map((quote) => {
          const profession = professionals.find(
            (p) => p.id === quote.profession
          );
          const state = states.find((s) => s.id === quote.state);
          const city = citiesByState[quote.state]?.find(
            (c) => c.id === quote.city
          );

          return (
            <div
              key={quote.id}
              className="border border-gray-200 rounded-lg p-6 hover:border-blue-500 transition-colors"
            >
              <div className="flex justify-between items-start mb-4">
                <div className="flex items-center gap-2">
                  <Calendar className="w-5 h-5 text-gray-400" />
                  <span className="text-gray-900">
                    {new Date(quote.created_at).toLocaleDateString("pt-BR")}
                  </span>
                </div>
                <div className="flex justify-between items-center mb-4">
                  {profile == "admin" && (
                    <label className="inline-flex items-center cursor-pointer">
                      <input
                        type="checkbox"
                        className="form-checkbox h-5 w-5 text-blue-600"
                        checked={quote.approved}
                        onChange={() =>
                          handleStatusChange(quote, quote.approved)
                        }
                      />
                      <span className="ml-2 text-sm text-gray-700">
                        Aprovar
                      </span>
                    </label>
                  )}
                  <button
                    onClick={() => handleDeleteQuote(quote.id)}
                    className="text-red-600 hover:bg-red-50 p-2 rounded-lg transition-colors"
                  >
                    <Trash2 className="w-5 h-5" />
                  </button>
                </div>
              </div>

              {(profile == "admin" || profile == "client") &&
                renderProfessionalInfo(quote)}
              {(profile == "admin" ||
                profile == "profissional" ||
                profile == "store") &&
                renderClientInfoPreview(quote)}
              {budgetFilter !== "store" && (
                <div className="border-t border-gray-100 pt-4">
                  <div className="flex items-start gap-2">
                    <MessageSquare className="w-5 h-5 text-gray-400 mt-1" />
                    <div>
                      <span className="text-sm text-gray-500">Descrição:</span>
                      <p className="mt-2 text-gray-900">{quote.description}</p>
                    </div>
                  </div>
                </div>
              )}

              {profile == "admin" && (
                <div className="flex justify-end gap-3 mt-4">
                  <a
                    href={`https://wa.me/${quote.telephone.replace(/\D/g, "")}`}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="flex items-center gap-2 px-4 py-2 text-green-600 hover:bg-green-50 rounded-lg transition-colors"
                  >
                    <Phone className="w-5 h-5" />
                    Responder via WhatsApp
                  </a>
                </div>
              )}
            </div>
          );
        })}
      </div>

      {/* Lock Modal */}
      <BudgetLockModal
        isOpen={showLockModal}
        onClose={() => setShowLockModal(false)}
        budgetId={selectedBudgetId || 0}
        onPayUnlock={handlePayUnlock}
      />
    </div>
  );
}

export default QuotesPanel;
