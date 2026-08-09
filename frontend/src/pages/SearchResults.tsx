import React, { useEffect, useState } from "react";
import {
  ArrowLeft,
  Phone,
  User,
  MessageSquare,
  X,
  Users,
  Info,
  ShieldCheck,
  Building2,
  MapPin,
  Send,
  BadgeCheck,
  Clock,
  CheckCircle,
} from "lucide-react";
import InputMask from "react-input-mask";
import { BudgetService } from "../services/Budget";
import Swal from "sweetalert2";
import Navigation from "../components/Navigation";
import { ProfessionalService } from "../services/ProfessionalService";
import { useLocation, useNavigate } from "react-router-dom";
import PrivacyPolicyCheckbox from "../components/PrivacyPolicyCheckbox";
import {
  BUDGET_FORM_DRAFT_KEYS,
  clearBudgetFormDraft,
  loadBudgetFormDraft,
} from "../utils/budgetFormDraft";
import {
  IBudget,
  IProfissional,
} from "../interfaces";

interface Professional {
  id: string;
  nome: string;
  rating: number;
  reviews: number;
  city: string;
  state: string;
  phone: string;
  photo: string;
}

interface FormData {
  name: string;
  phone: string;
  message: string;
}

interface SearchResultsFormDraft {
  formData: FormData;
  privacyAccepted: boolean;
  isBulkRequest: boolean;
  selectedProfessionalOid: number | null;
  professionalLabel: string | null;
}

const getImageSrc = (imageData: string | null | undefined): string | null => {
  if (!imageData || imageData.trim() === "") return null;

  if (imageData.startsWith("data:image")) {
    return imageData;
  }

  if (imageData.startsWith("http://") || imageData.startsWith("https://")) {
    return imageData;
  }

  if (imageData.length > 100) {
    return `data:image/jpeg;base64,${imageData}`;
  }

  return null;
};

function SearchResults() {
  const [showContactForm, setShowContactForm] = useState(false);
  const [showProfessionalSearch, setShowProfessionalSearch] = useState(true);
  const [privacyAccepted, setPrivacyAccepted] = useState(false);
  const [restoredProfessionalLabel, setRestoredProfessionalLabel] = useState<
    string | null
  >(null);
  const [searchCityId, setSearchCityId] = useState<number>(0);
  const [selectedProfessional, setSelectedProfessional] =
    useState<IProfissional | null>(null);
  const [isBulkRequest, setIsBulkRequest] = useState(false);
  const [formData, setFormData] = useState<FormData>({
    name: "",
    phone: "",
    message: "",
  });

  const [showPhoneNumbers, setShowPhoneNumbers] = useState<boolean>(false);
  const [modalInfoProfissional, setModalInfoProfissional] = useState(false);
  const [showExpandedImage, setShowExpandedImage] = useState(false);
  const [currentPage, setCurrentPage] = useState<string>("search-results");
  const [professionals, setProfessionals] = useState<IProfissional[]>([]);
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    const fetchData = async () => {
      const selectedCity = location.state?.selectedCity;
      const selectedProfessional = location.state?.selectedProfessional;

      if (selectedCity) {
        setSearchCityId(parseInt(selectedCity));
      }

      const res = await ProfessionalService.getProfessionalsRandomPublic({
        professionId: parseInt(selectedProfessional),
        limit: 1000,
        offset: 0,
      });
      const all = res.data as IProfissional[];
      const filtered =
        selectedCity && String(selectedCity).length > 0
          ? all.filter((p) => p.cidade?.oid === parseInt(selectedCity))
          : all;
      setProfessionals(filtered);
    };

    fetchData();
  }, []);

  const clearRestoreNavigationState = () => {
    const { restoreFormDraft: _, ...preservedState } =
      (location.state as Record<string, unknown>) ?? {};
    navigate(`${location.pathname}${location.search}`, {
      replace: true,
      state: preservedState,
    });
  };

  const getProfessionalLabel = (
    professional: IProfissional | null
  ): string | null => {
    if (!professional) return null;
    if (professional.nome) return professional.nome;
    const cityName = professional.cidade?.nome;
    const cityUf = professional.cidade?.uf;
    if (cityName && cityUf) return `${cityName}, ${cityUf}`;
    return null;
  };

  const getContactFormSubtitle = (): string => {
    if (isBulkRequest) {
      return `Enviar para ${professionals.length} profissionais`;
    }
    const label = getProfessionalLabel(selectedProfessional);
    if (label) return label;
    return restoredProfessionalLabel ?? "";
  };

  useEffect(() => {
    if (
      location.state?.restoreFormDraft !==
      BUDGET_FORM_DRAFT_KEYS.searchResults
    ) {
      return;
    }

    const draft = loadBudgetFormDraft<SearchResultsFormDraft>(
      BUDGET_FORM_DRAFT_KEYS.searchResults
    );
    if (!draft) return;

    setFormData(draft.formData);
    setPrivacyAccepted(draft.privacyAccepted);
    setIsBulkRequest(draft.isBulkRequest);
    setRestoredProfessionalLabel(draft.professionalLabel);
    setShowContactForm(true);
    setShowProfessionalSearch(false);

    if (draft.isBulkRequest) {
      setSelectedProfessional(null);
      clearRestoreNavigationState();
      return;
    }

    if (draft.selectedProfessionalOid && professionals.length === 0) {
      return;
    }

    if (draft.selectedProfessionalOid && professionals.length > 0) {
      const professional = professionals.find(
        (item) => item.oid === draft.selectedProfessionalOid
      );
      setSelectedProfessional(professional ?? null);
      if (professional) {
        setRestoredProfessionalLabel(getProfessionalLabel(professional));
      }
    } else {
      setSelectedProfessional(null);
    }

    clearRestoreNavigationState();
  }, [location.state?.restoreFormDraft, professionals]);

  const openQuoteForm = (
    professional: IProfissional | null = null,
    bulk = false
  ) => {
    setSelectedProfessional(professional);
    setIsBulkRequest(bulk);
    setPrivacyAccepted(false);
    setRestoredProfessionalLabel(null);
    setShowProfessionalSearch(false);
    setShowContactForm(true);
  };

  const getBudgetCityId = (): number => {
    if (searchCityId > 0) return searchCityId;
    if (selectedProfessional?.cidade?.oid) return selectedProfessional.cidade.oid;
    if (professionals.length > 0 && professionals[0].cidade?.oid) {
      return professionals[0].cidade.oid;
    }
    return 0;
  };

  const handleFormSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!privacyAccepted) {
      Swal.fire({
        position: "center",
        icon: "warning",
        title: "Aceite a política de privacidade",
        text: "É necessário aceitar a política de privacidade para continuar.",
        showConfirmButton: true,
      });
      return;
    }

    let profs: string[];
    if (selectedProfessional == null) {
      profs = professionals.map((prof) => prof.oid.toString());
    } else {
      profs = [selectedProfessional.oid.toString()];
    }

    const budget: IBudget = {
      name: formData.name,
      email: "",
      telephone: formData.phone,
      description: formData.message,
      termResponsabilityAccepted: privacyAccepted,
      cityId: getBudgetCityId(),
      professionalsId: [],
    };
    budget.professionalsId = profs.map(Number);

    const postReturn = await BudgetService.saveBudget(budget);

    if (postReturn.status == 200) {
      clearBudgetFormDraft(BUDGET_FORM_DRAFT_KEYS.searchResults);
      setShowContactForm(false);
      setShowProfessionalSearch(true);

      setFormData({
        name: "",
        phone: "",
        message: "",
      });
      setPrivacyAccepted(false);
      setRestoredProfessionalLabel(null);

      Swal.fire({
        position: "center",
        icon: "success",
        title: "Solicitação enviada com sucesso!",
        text: "Os profissionais entrarão em contato em breve.",
        showConfirmButton: false,
        timer: 3000,
      });
    } else {
      Swal.fire({
        position: "center",
        icon: "error",
        title: "Erro ao enviar orçamento",
        showConfirmButton: false,
        timer: 1500,
      });
    }
  };

  const handleInputChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  return (
    <div className="min-h-screen bg-gray-50 py-12">
      <Navigation
        currentPage={currentPage}
        setCurrentPage={setCurrentPage}
        carouselPage={""}
        setCarouselPage={function (page: string): void {
          throw new Error("Function not implemented.");
        }}
      />
      
      {/* Contact Form Modal Atualizado */}
      {showContactForm && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
          <div className="bg-white rounded-lg md:rounded-none shadow-xl w-full max-w-md">
            <div className="p-6">
              <div className="flex justify-between items-start mb-5">
                <div className="flex items-center gap-3">
                  {isBulkRequest ? (
                    <div className="w-12 h-12 rounded-full bg-blue-50 flex items-center justify-center border border-blue-100">
                      <Users className="w-6 h-6 text-blue-600" />
                    </div>
                  ) : (
                    <div className="w-12 h-12 rounded-full bg-blue-50 flex items-center justify-center border border-blue-100">
                      <User className="w-6 h-6 text-blue-600" />
                    </div>
                  )}
                  <div>
                    <h3 className="text-xl font-bold text-gray-900">
                      Solicitar Orçamento
                    </h3>
                    <p className="text-sm text-gray-600 mt-0.5">
                      {getContactFormSubtitle()}
                    </p>
                  </div>
                </div>
                <button
                  onClick={() => {
                    setShowContactForm(false);
                    setSelectedProfessional(null);
                    setShowProfessionalSearch(true);
                  }}
                  className="text-gray-400 hover:text-gray-600 transition-colors"
                >
                  <X className="w-6 h-6" />
                </button>
              </div>

              {/* Box Informativo Azul */}
              <div className="bg-[#f0f4ff] border border-blue-100 rounded-lg md:rounded-none p-4 mb-6 flex gap-3 items-start">
                <div className="flex-shrink-0 mt-0.5">
                  <div className="w-6 h-6 rounded-full bg-blue-600 flex items-center justify-center">
                    <Info className="w-4 h-4 text-white" />
                  </div>
                </div>
                <div>
                  <h4 className="text-sm font-bold text-blue-800 mb-1">
                    Gratuito e sem compromisso!
                  </h4>
                  <p className="text-sm text-blue-800/90 leading-snug">
                    Seu pedido será enviado para os profissionais da sua região. Eles entrarão em contato com você pelo WhatsApp.
                  </p>
                </div>
              </div>

              <form onSubmit={handleFormSubmit} className="space-y-4">
                <div>
                  <label
                    htmlFor="name"
                    className="block text-sm font-bold text-gray-700 mb-1"
                  >
                    Nome Completo
                  </label>
                  <div className="relative">
                    <User className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                    <input
                      type="text"
                      id="name"
                      name="name"
                      value={formData.name}
                      onChange={handleInputChange}
                      required
                      placeholder="Digite seu nome"
                      className="block w-full pl-10 pr-3 py-2.5 border border-gray-300 rounded-lg md:rounded-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-50 text-sm"
                    />
                  </div>
                </div>

                <div>
                  <label
                    htmlFor="phone"
                    className="block text-sm font-bold text-gray-700 mb-1"
                  >
                    Telefone (WhatsApp)
                  </label>
                  <div className="relative">
                    <Phone className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                    <InputMask
                      mask="(99) 99999-9999"
                      type="tel"
                      id="phone"
                      name="phone"
                      value={formData.phone}
                      onChange={handleInputChange}
                      required
                      placeholder="(xx) xxxxx-xxxx"
                      className="block w-full pl-10 pr-3 py-2.5 border border-gray-300 rounded-lg md:rounded-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-50 text-sm"
                    />
                  </div>
                </div>

                <div>
                  <label
                    htmlFor="message"
                    className="block text-sm font-bold text-gray-700 mb-1"
                  >
                    Descrição do orçamento
                  </label>
                  <div className="relative">
                    <MessageSquare className="absolute left-3 top-3 text-gray-400 w-5 h-5" />
                    <textarea
                      id="message"
                      name="message"
                      value={formData.message}
                      onChange={handleInputChange}
                      required
                      rows={3}
                      className="block w-full pl-10 pr-3 py-2.5 border border-gray-300 rounded-lg md:rounded-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-sm"
                      placeholder="Descreva o serviço que você precisa..."
                    />
                  </div>
                </div>

                {/* Seção de Vantagens (Ícones Verdes) */}
                <div className="flex items-center justify-between py-2 mb-2 mt-4">
                  <div className="flex items-center gap-1.5">
                    <CheckCircle className="w-4 h-4 sm:w-5 sm:h-5 text-green-600" />
                    <span className="text-[11px] sm:text-xs text-gray-700 font-medium">Gratuito</span>
                  </div>
                  <div className="w-px h-5 bg-gray-200"></div>
                  <div className="flex items-center gap-1.5">
                    <ShieldCheck className="w-4 h-4 sm:w-5 sm:h-5 text-green-600" />
                    <span className="text-[11px] sm:text-xs text-gray-700 font-medium">Sem compromisso</span>
                  </div>
                  <div className="w-px h-5 bg-gray-200"></div>
                  <div className="flex items-center gap-1.5">
                    <Clock className="w-4 h-4 sm:w-5 sm:h-5 text-green-600" />
                    <span className="text-[11px] sm:text-xs text-gray-700 font-medium">Resposta rápida</span>
                  </div>
                </div>

                <PrivacyPolicyCheckbox
                  checked={privacyAccepted}
                  onChange={setPrivacyAccepted}
                  formDraftKey={BUDGET_FORM_DRAFT_KEYS.searchResults}
                  getFormDraft={() => ({
                    formData,
                    privacyAccepted,
                    isBulkRequest,
                    selectedProfessionalOid: selectedProfessional?.oid ?? null,
                    professionalLabel: isBulkRequest
                      ? null
                      : getProfessionalLabel(selectedProfessional),
                  })}
                />

                <button
                  type="submit"
                  disabled={!privacyAccepted}
                  className="w-full bg-[#FF6B35] text-white py-3 px-4 rounded-lg md:rounded-none md:border md:border-[#FF6B35] hover:bg-[#E55A2B] hover:md:border-[#E55A2B] transition-colors disabled:bg-gray-300 disabled:md:border-gray-300 disabled:cursor-not-allowed flex items-center justify-center gap-2 font-bold mt-2"
                >
                  <Send className="w-5 h-5" />
                  Enviar Solicitação
                </button>
              </form>
            </div>
          </div>
        </div>
      )}
      
      {showProfessionalSearch && (
        <div className="max-w-4xl mx-auto px-4">
          <div className="flex items-center justify-between mb-8">
            <div>
              <p className="text-gray-600 mt-2">
                {professionals.length} profissionais encontrados
              </p>
            </div>
            <div className="flex gap-4">
              <button
                onClick={() => navigate("/search")}
                className="flex items-center gap-2 px-4 py-2 text-blue-600 hover:bg-blue-50 rounded-lg transition-colors"
              >
                <ArrowLeft className="w-5 h-5" />
                Nova Pesquisa
              </button>
            </div>
          </div>

          {professionals.length > 0 && (
            <div className="mb-8 rounded-2xl p-4 sm:p-5 flex flex-col sm:flex-row items-center gap-3 sm:gap-4">
              <button
                onClick={() => openQuoteForm(null, true)}
                className="w-full sm:w-auto flex items-center justify-center gap-2 bg-blue-600 hover:bg-blue-700 text-white font-semibold py-3 px-6 rounded-xl transition-colors whitespace-nowrap shadow-sm"
              >
                <Send className="w-5 h-5" />
                Solicitar para TODOS DE UMA VEZ
              </button>
            </div>
          )}

          <div className="space-y-4">
            {professionals.map((professional: IProfissional) => (
              <div
                key={professional.oid}
                className="p-4 bg-white border border-gray-200 rounded-lg hover:border-blue-500 transition-colors"
              >
                <div className="flex items-start gap-4">
                  <div className="flex-1 min-w-0">
                    {professional.isPremium && (
                      <div className="mb-3 sm:hidden">
                        <span className="whitespace-nowrap text-xs font-semibold text-white bg-green-500 py-1 px-3 rounded-full">
                          Profissional Premium
                        </span>
                      </div>
                    )}

                    <div className="hidden sm:flex items-center gap-3">
                      {professional.isPremium && (
                        <span className="text-xs font-semibold text-white bg-green-500 py-1 px-3 rounded-full self-center flex-shrink-0">
                          Profissional Premium
                        </span>
                      )}

                      <div className="flex-1">
                        <h3 className="text-lg font-semibold text-gray-900">
                          {professional.nome}
                        </h3>
                        <div className="flex items-center text-sm text-gray-600">
                          <span>
                            {professional.cidade.nome}, {professional.cidade.uf}
                          </span>
                        </div>
                      </div>
                    </div>

                    <div className="sm:hidden">
                      <h3 className="text-lg font-semibold text-gray-900">
                        {professional.nome}
                      </h3>
                      <div className="flex items-center text-sm text-gray-600">
                        <span>
                          {professional.cidade.nome}, {professional.cidade.uf}
                        </span>
                      </div>
                    </div>

                    {showPhoneNumbers &&
                      (isBulkRequest ||
                        selectedProfessional?.oid === professional.oid) && (
                        <div className="mt-2 flex items-center gap-2 text-green-600">
                          <Phone className="w-4 h-4" />
                          {professional.telefone}
                        </div>
                      )}
                  </div>

                  <div className="flex items-start gap-6">
                    {professional.isPremium && (
                      <button
                        onClick={() => {
                          setSelectedProfessional(professional);
                          setModalInfoProfissional(true);
                        }}
                        className="relative flex-shrink-0 rounded-lg hover:bg-[#fd7c4c1a] transition-colors sm:h-16 sm:w-20"
                      >
                        <img
                          className="h-28 w-28 -mt-3 sm:absolute sm:left-0 sm:top-1/2 sm:mt-0 sm:h-20 sm:w-32 sm:-translate-y-1/2"
                          src="images/saibaMais.png"
                          alt="Saiba Mais"
                        />
                      </button>
                    )}

                    <button
                      onClick={() => openQuoteForm(professional, false)}
                      className={`hidden sm:flex items-center self-start ${
                        professional.isPremium
                          ? "bg-[#FF6B35] hover:bg-[#E55A2B]"
                          : "bg-blue-600 hover:bg-blue-700"
                      } py-1 px-4 text-white rounded-lg transition-colors whitespace-nowrap h-16`}
                    >
                      Solicitar Orçamento
                    </button>
                  </div>
                </div>

                <div className="sm:hidden mt-1">
                  <button
                    onClick={() => openQuoteForm(professional, false)}
                    className={`w-44 ${
                      professional.isPremium
                        ? "bg-[#FF6B35] hover:bg-[#E55A2B]"
                        : "bg-blue-600 hover:bg-blue-700"
                    } py-2 px-4 text-white rounded-lg transition-colors`}
                  >
                    Solicitar Orçamento
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      )}
      
      {modalInfoProfissional && selectedProfessional && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
          <div className="bg-white rounded-lg shadow-xl w-full max-w-2xl max-h-[90vh] overflow-y-auto">
            <div className="relative bg-white">
              <div
                className="absolute rounded-tl-sm top-0 right-0 z-50 bg-white h-8 md:w-96 w-44"
                style={{
                  clipPath: "polygon(0 0, 100% 0, 100% 100%, 8% 100%)",
                }}
              ></div>{" "}
              <div className="flex pt-12 justify-between items-start p-8 bg-[#FF6B35] rounded-t-lg">
                <div
                  className="absolute -bottom-1 left-0 z-50 bg-white h-7 md:w-56 w-40 rounded-tl-2xl"
                  style={{
                    clipPath: "polygon(0 0, 85% 0, 100% 100%, 0 100%)",
                  }}
                ></div>
                <div className="flex items-center gap-2">
                  {getImageSrc(selectedProfessional.image) ? (
                    <img
                      src={getImageSrc(selectedProfessional.image)!}
                      alt={selectedProfessional.nome}
                      className="w-24 h-24 rounded-full object-cover border-4 border-white cursor-pointer hover:opacity-90 transition-opacity"
                      onClick={() => setShowExpandedImage(true)}
                      onError={(e) => {
                        e.currentTarget.style.display = "none";
                        const fallback = e.currentTarget
                          .nextElementSibling as HTMLElement;
                        if (fallback) fallback.classList.remove("hidden");
                      }}
                    />
                  ) : null}
                  <div
                    className={`w-24 h-24 rounded-full bg-white bg-opacity-20 flex items-center justify-center ${
                      getImageSrc(selectedProfessional.image) ? "hidden" : ""
                    }`}
                  >
                    <User className="w-12 h-12 text-white" />
                  </div>
                  <div>
                    <h3 className="text-3xl font-bold text-white">
                      {selectedProfessional.nome}
                    </h3>
                    {selectedProfessional.isPremium && (
                      <div className="flex items-center gap-1.5 mt-1 bg-green-500 rounded-full py-1 px-3">
                        <BadgeCheck className="w-4 h-4 text-white" />
                        <span className="text-xs font-bold text-white whitespace-nowrap">
                          Profissional Premium
                        </span>
                      </div>
                    )}
                  </div>
                </div>
                <button
                  onClick={() => {
                    setModalInfoProfissional(false);
                    setShowExpandedImage(false);
                  }}
                  className="text-white hover:text-white/80 transition-colors"
                >
                  <X className="w-6 h-6" />
                </button>
              </div>
            </div>

            <div className="space-y-6 p-6">
              <div className="border-b border-gray-200 pb-4">
                <h4 className="text-lg font-semibold text-gray-900 mb-3">
                  Localização
                </h4>
                <div className="space-y-3">
                  <div className="flex items-center gap-3">
                    <MapPin className="w-5 h-5 text-blue-600" />
                    <div>
                      <p className="text-sm text-gray-600">Cidade</p>
                      <p className="text-gray-900">
                        {selectedProfessional.cidade.nome},{" "}
                        {selectedProfessional.cidade.uf}
                      </p>
                    </div>
                  </div>
                  {selectedProfessional.neighborhood && (
                    <div className="flex items-center gap-3">
                      <Building2 className="w-5 h-5 text-blue-600" />
                      <div>
                        <p className="text-sm text-gray-600">Bairro</p>
                        <p className="text-gray-900">
                          {selectedProfessional.neighborhood}
                        </p>
                      </div>
                    </div>
                  )}
                  {selectedProfessional.street && (
                    <div className="flex items-center gap-3">
                      <Building2 className="w-5 h-5 text-blue-600" />
                      <div>
                        <p className="text-sm text-gray-600">Endereço</p>
                        <p className="text-gray-900">
                          {selectedProfessional.street}
                        </p>
                      </div>
                    </div>
                  )}
                  {selectedProfessional.cep && (
                    <div className="flex items-center gap-3">
                      <MapPin className="w-5 h-5 text-blue-600" />
                      <div>
                        <p className="text-sm text-gray-600">CEP</p>
                        <p className="text-gray-900">
                          {selectedProfessional.cep}
                        </p>
                      </div>
                    </div>
                  )}
                </div>
              </div>

              <div className="border-b border-gray-200 pb-4">
                <h4 className="text-lg font-semibold text-gray-900 mb-3">
                  Informações Profissionais
                </h4>
                <div className="space-y-3">
                  {selectedProfessional.experience && (
                    <div className="flex items-start gap-3">
                      <Info className="w-5 h-5 text-blue-600 mt-0.5" />
                      <div>
                        <p className="text-sm text-gray-600">Experiência</p>
                        <p className="text-gray-900">
                          {selectedProfessional.experience}
                        </p>
                      </div>
                    </div>
                  )}
                  {selectedProfessional.meiCnpj && (
                    <div className="flex items-start gap-3">
                      <Building2 className="w-5 h-5 text-blue-600 mt-0.5" />
                      <div>
                        <p className="text-sm text-gray-600">MEI/CNPJ</p>
                        <p className="text-gray-900">
                          {selectedProfessional.meiCnpj}
                        </p>
                      </div>
                    </div>
                  )}
                  {selectedProfessional.verified && (
                    <div className="flex items-start gap-3">
                      <ShieldCheck className="w-5 h-5 text-green-600 mt-0.5" />
                      <div>
                        <p className="text-sm text-gray-600">Status</p>
                        <p className="text-green-600 font-semibold">
                          Profissional Verificado
                        </p>
                      </div>
                    </div>
                  )}
                </div>
              </div>

              {selectedProfessional.isPremium && selectedProfessional.youtubeUrl && (
                <div className="border-b border-gray-200 pb-4">
                  <h4 className="text-lg font-semibold text-gray-900 mb-3">
                    Projetos já realizados
                  </h4>
                  <div className="relative w-full" style={{ paddingBottom: "56.25%" }}>
                    <iframe
                      className="absolute inset-0 w-full h-full rounded-lg"
                      src={(() => {
                        const match = selectedProfessional.youtubeUrl.match(
                          /(?:youtube\.com\/watch\?v=|youtu\.be\/)([^&\n?#]+)/
                        );
                        return match
                          ? `https://www.youtube.com/embed/${match[1]}`
                          : selectedProfessional.youtubeUrl;
                      })()}
                      title="Vídeo do profissional"
                      allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                      allowFullScreen
                    />
                  </div>
                </div>
              )}

              <div className="flex gap-3 pt-2">
                <button
                  onClick={() => {
                    setModalInfoProfissional(false);
                    setShowExpandedImage(false);
                    openQuoteForm(selectedProfessional, false);
                  }}
                  className={`${
                    selectedProfessional.telefone &&
                    selectedProfessional.telefone.trim() !== ""
                      ? "flex-1"
                      : "w-full"
                  } bg-[#FF6B35] hover:bg-[#E55A2B] text-white py-3 px-4 rounded-lg transition-colors font-semibold`}
                >
                  Solicitar Orçamento
                </button>
              </div>
            </div>
          </div>
        </div>
      )}

      {showExpandedImage &&
        selectedProfessional &&
        getImageSrc(selectedProfessional.image) && (
          <div
            className="fixed inset-0 bg-black bg-opacity-90 flex items-center justify-center z-[60] p-4"
            onClick={() => setShowExpandedImage(false)}
          >
            <div className="relative max-w-4xl max-h-[90vh]">
              <button
                onClick={() => setShowExpandedImage(false)}
                className="absolute -top-12 right-0 text-white hover:text-gray-300 transition-colors"
              >
                <X className="w-8 h-8" />
              </button>
              <img
                src={getImageSrc(selectedProfessional.image)!}
                alt={selectedProfessional.nome}
                className="max-w-full max-h-[85vh] object-contain rounded-lg"
                onClick={(e) => e.stopPropagation()}
              />
            </div>
          </div>
        )}
    </div>
  );
}

export default SearchResults;