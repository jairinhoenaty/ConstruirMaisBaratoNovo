import React, { useEffect, useState } from "react";
import {
  ArrowLeft,
  // Star,
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
} from "lucide-react";
import InputMask from "react-input-mask";
import { BudgetService } from "../services/Budget";
import Swal from "sweetalert2";
import Navigation from "../components/Navigation";
import { ProfessionalService } from "../services/ProfessionalService";
import { useLocation, useNavigate } from "react-router-dom";
import PrivacyPolicyCheckbox from "../components/PrivacyPolicyCheckbox";
import {
  IBudget,
  IProfissional,
} from "../interfaces";
//import { useNavigate } from "react-router-dom";

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

// interface SearchResultsProps {
//   profession: string;
//   professionals: IProfissional[];
//   onNewSearch: () => void;
// }

// Helper function para processar imagem (base64 ou URL)
const getImageSrc = (imageData: string | null | undefined): string | null => {
  if (!imageData || imageData.trim() === "") return null;

  // Se já é uma data URL (base64), retorna direto
  if (imageData.startsWith("data:image")) {
    return imageData;
  }

  // Se é uma URL normal, retorna direto
  if (imageData.startsWith("http://") || imageData.startsWith("https://")) {
    return imageData;
  }

  // Se é só o base64 sem o prefixo, adiciona o prefixo
  // Assume JPEG por padrão, mas pode ser ajustado
  if (imageData.length > 100) {
    return `data:image/jpeg;base64,${imageData}`;
  }

  return null;
};

function SearchResults() {
  /*{
  profession,
  professionals,
  onNewSearch,
}: SearchResultsProps*/
  const [showContactForm, setShowContactForm] = useState(false);
  const [showProfessionalSearch, setShowProfessionalSearch] = useState(true);
  const [privacyAccepted, setPrivacyAccepted] = useState(false);
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
  // const [profession, setProfession] = useState<string>("");
  const navigate = useNavigate();
  const location = useLocation();

  //  console.log(professionals);
  // Mock data for professionals
  /*const mockProfessionals: Professional[] = [
    {
      id: "1",
      name: "João Silva",
      rating: 4.8,
      reviews: 156,
      city: "São Paulo",
      state: "SP",
      phone: "(11) 99999-9999",
      photo: "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d",
    },
    {
      id: "2",
      name: "Maria Santos",
      rating: 4.9,
      reviews: 203,
      city: "Rio de Janeiro",
      state: "RJ",
      phone: "(21) 98888-8888",
      photo: "https://images.unsplash.com/photo-1494790108377-be9c29b29330",
    },
    {
      id: "3",
      name: "Pedro Oliveira",
      rating: 4.7,
      reviews: 89,
      city: "Belo Horizonte",
      state: "MG",
      phone: "(31) 97777-7777",
      photo: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e",
    },
  ];
  */

  // useEffect(() => {
  //   const fetchData = async () => {
  //     // console.log("Effect");
  //     // console.log(location);
  //     // console.log(location.state.selectedProfessional);
  //     const selectedCity = location.state.selectedCity;
  //     const selectedProfessional = location.state.selectedProfessional;
  //     const return_professionals =
  //       await ProfessionalService.getProfessionalByCityAndProfession({
  //         cityID: parseInt(selectedCity),
  //         professionID: parseInt(selectedProfessional),
  //         limit: 1000,
  //         offset: 0,
  //       });

  //     const json_professionals = await return_professionals.data.profissionais;
  //     setProfessionals(json_professionals);
  //     // setProfession(selectedProfessional);
  //   };

  //   fetchData();
  // }, []);

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

  const openQuoteForm = (
    professional: IProfissional | null = null,
    bulk = false
  ) => {
    setSelectedProfessional(professional);
    setIsBulkRequest(bulk);
    setPrivacyAccepted(false);
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
      setShowContactForm(false);
      setShowProfessionalSearch(true);

      setFormData({
        name: "",
        phone: "",
        message: "",
      });
      setPrivacyAccepted(false);

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
      {/* Contact Form Modal */}
      {showContactForm && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
          <div className="bg-white rounded-lg shadow-xl w-full max-w-md">
            <div className="p-6">
              <div className="flex justify-between items-start mb-6">
                <div className="flex items-center gap-3">
                  {isBulkRequest ? (
                    <div className="w-12 h-12 rounded-full bg-blue-100 flex items-center justify-center">
                      <Users className="w-6 h-6 text-blue-600" />
                    </div>
                  ) : (
                    /*                    <img
                      src={selectedProfessional?.photo}
                      alt={selectedProfessional?.nome}
                      className="w-12 h-12 rounded-full object-cover"
                    />*/
                    <div className="w-12 h-12 rounded-full bg-blue-100 flex items-center justify-center">
                      <User className="w-6 h-6 text-blue-600" />
                    </div>
                  )}
                  <div>
                    <h3 className="text-lg font-semibold text-gray-900">
                      {/* {isBulkRequest
                        ? "Solicitar Orçamento"
                        : selectedProfessional?.nome} */}
                      Solicitar Orçamento
                    </h3>
                    <p className="text-sm text-gray-600">
                      {isBulkRequest
                        ? `Enviar para ${professionals.length} profissionais`
                        : selectedProfessional && selectedProfessional.nome
                        ? `${selectedProfessional.nome}`
                        : `${selectedProfessional?.cidade.nome}, ${selectedProfessional?.cidade.uf}`}
                    </p>
                  </div>
                </div>
                <button
                  onClick={() => {
                    setShowContactForm(false);
                    setSelectedProfessional(null);
                    setShowProfessionalSearch(true);
                  }}
                  className="text-gray-400 hover:text-gray-600"
                >
                  <X className="w-5 h-5" />
                </button>
              </div>

              <form onSubmit={handleFormSubmit} className="space-y-4">
                <div>
                  <label
                    htmlFor="name"
                    className="block text-sm font-medium text-gray-700 mb-1"
                  >
                    Nome Completo
                    {/* {isClient} */}
                  </label>

                  <div className="relative">
                    <User className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                    <input
                      type="text"
                      id="name"
                      name="name"
                      value={formData.name}
                      onChange={handleInputChange}
                      // disabled={isClient}
                      required
                      className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500  disabled:bg-gray-50"
                    />
                  </div>
                </div>

                <div>
                  <label
                    htmlFor="phone"
                    className="block text-sm font-medium text-gray-700 mb-1"
                  >
                    Telefone
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
                      // disabled={isClient}
                      required
                      className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-50"
                    />
                  </div>
                </div>

                <div>
                  <label
                    htmlFor="message"
                    className="block text-sm font-medium text-gray-700 mb-1"
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
                      rows={4}
                      className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      placeholder="Descreva o serviço que você precisa..."
                    />
                  </div>
                </div>

                <PrivacyPolicyCheckbox
                  checked={privacyAccepted}
                  onChange={setPrivacyAccepted}
                />

                <button
                  type="submit"
                  disabled={!privacyAccepted}
                  className="w-full bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 transition-colors disabled:bg-gray-300 disabled:cursor-not-allowed"
                >
                  Solicitar Orçamento
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

          {/* Botão pra solicitar orçamento para todos os profissionais que veio do backend (com os filtros)*/}
          {professionals.length > 0 && (
            // <div className="mb-8 rounded-2xl bg-gradient-to-r from-blue-50 to-purple-50 border border-purple-100 p-4 sm:p-5 flex flex-col sm:flex-row items-center gap-3 sm:gap-4">
            <div className="mb-8 rounded-2xl p-4 sm:p-5 flex flex-col sm:flex-row items-center gap-3 sm:gap-4">
              {/* <div className="flex flex-col sm:flex-row items-center gap-3 flex-1 w-full text-center sm:text-left">
                <div className="bg-white rounded-full p-3 shadow-sm flex-shrink-0">
                  <Users className="w-7 h-7 text-blue-600" />
                </div>
                <div>
                  <p className="text-sm sm:text-base font-semibold text-gray-800 leading-snug">
                    Solicite uma única vez e receba orçamentos de todos os
                    profissionais.
                  </p>
                  <div className="flex items-center justify-center sm:justify-start gap-1 text-xs text-gray-500 mt-1">
                    <Check className="w-4 h-4 text-green-500 flex-shrink-0" />
                    <span>
                      Envie para os {professionals.length} profissionais
                      encontrados
                    </span>
                  </div>
                </div>
              </div> */}
              <button
                onClick={() => openQuoteForm(null, true)}
                className="w-full sm:w-auto flex items-center justify-center gap-2 bg-blue-600 hover:bg-blue-700 text-white font-semibold py-3 px-6 rounded-xl transition-colors whitespace-nowrap shadow-sm"
              >
                <Send className="w-5 h-5" />
                Solicitar para TODOS DE UMA VEZ
              </button>
            </div>
          )}

          {/* Professional List */}
          <div className="space-y-4">
            {professionals.map((professional: IProfissional) => (
              <div
                key={professional.oid}
                className="p-4 bg-white border border-gray-200 rounded-lg hover:border-blue-500 transition-colors"
              >
                <div className="flex items-start gap-4">
                  <div className="flex-1 min-w-0">
                    {/* Badge Premium - Mobile (acima) */}
                    {professional.isPremium && (
                      <div className="mb-3 sm:hidden">
                        <span className="whitespace-nowrap text-xs font-semibold text-white bg-green-500 py-1 px-3 rounded-full">
                          Profissional Premium
                        </span>
                      </div>
                    )}

                    {/* Desktop: Badge à esquerda, Nome e Cidade à direita */}
                    <div className="hidden sm:flex items-center gap-3">
                      {/* Badge Premium - Desktop (à esquerda, centralizado) */}
                      {professional.isPremium && (
                        <span className="text-xs font-semibold text-white bg-green-500 py-1 px-3 rounded-full self-center flex-shrink-0">
                          Profissional Premium
                        </span>
                      )}

                      {/* Nome e Cidade empilhados */}
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

                    {/* Mobile: Nome e Cidade sem badge */}
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
                    {/* Botão Saiba Mais - Mobile e Desktop */}
                    {professional.isPremium && (
                      <button
                        onClick={() => {
                          setSelectedProfessional(professional);
                          setModalInfoProfissional(true);
                        }}
                        className="flex-shrink-0 rounded-lg hover:bg-[#fd7c4c1a] transition-colors"
                      >
                        <img
                          className="h-36 w-3h-36 -mt-3 sm:mt-0 sm:h-14 sm:w-14"
                          src="images/saibaMais.png"
                          alt="Saiba Mais"
                        />
                      </button>
                    )}

                    {/* Botão Solicitar Orçamento - Desktop */}
                    <button
                      onClick={() => openQuoteForm(professional, false)}
                      className={`hidden sm:flex items-center self-start ${
                        professional.isPremium
                          ? "bg-[#FF6B35] hover:bg-[#E55A2B]"
                          : "bg-blue-600 hover:bg-blue-700"
                      } py-1 px-4 text-white rounded-lg transition-colors whitespace-nowrap h-14`}
                    >
                      Solicitar Orçamento
                    </button>
                  </div>
                </div>

                {/* Botão Solicitar Orçamento - Mobile */}
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
            {/* Header do Modal */}
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
                    // setProfessionalForModal(null);
                  }}
                  className="text-white hover:text-white/80 transition-colors"
                >
                  <X className="w-6 h-6" />
                </button>
              </div>
            </div>

            {/* Conteúdo do Modal */}
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
          {/* </div> */}
        </div>
      )}

      {/* Modal de Imagem Expandida */}
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
