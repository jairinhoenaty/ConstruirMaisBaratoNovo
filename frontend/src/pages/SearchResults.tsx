import React, { useEffect, useState } from "react";
import {
  ArrowLeft,
  // Star,
  Phone,
  Mail,
  User,
  MessageSquare,
  X,
  Users,
  Info,
  ShieldCheck,
  Building2,
  MapPin,
  Brain,
  Check,
  BadgeCheck,
  // Crown,
  // Trophy,
  // Diamond,
  // Award,
  // BadgeCheck,
} from "lucide-react";
import InputMask from "react-input-mask";
import { BudgetService } from "../services/Budget";
import Swal from "sweetalert2";
import Login from "./Login";
import { states } from "../data";
import Navigation from "../components/Navigation";
import { ClientService } from "../services/ClientService";
import { CityService } from "../services/CityService";
import { ProfessionalService } from "../services/ProfessionalService";
import { useLocation, useNavigate } from "react-router-dom";
import Select from "react-select";
import {
  IBudget,
  ICitySearchProfessionals,
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
  email: string;
  phone: string;
  message: string;
  clientId: number;
  cityId: number;
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
  const [selectedState, setSelectedState] = useState<string>("");
  const [selectedCity, setSelectedCity] = useState<string>("");
  const [citiesByState, setcitiesByState] = useState<
    ICitySearchProfessionals[]
  >([]);
  const [showLGPDTerms, setShowLGPDTerms] = useState(false);
  const [modalInfoProfissional, setModalInfoProfissional] = useState(false);
  const [showExpandedImage, setShowExpandedImage] = useState(false);
  // const [selectedProfessional, setProfessionalForModal] = useState<IProfissional | null>(null);
  // const [showLogin, setShowLogin] = useState(false);
  const [showContactForm, setShowContactForm] = useState(false);
  const [showProfessionalSearch, setShowProfessionalSearch] = useState(true);
  const [selectedProfessional, setSelectedProfessional] =
    useState<IProfissional | null>(null);
  const [isBulkRequest, setIsBulkRequest] = useState(false);
  const [formData, setFormData] = useState<FormData>({
    name: "",
    email: "",
    phone: "",
    message: "",
    clientId: 0,
    cityId: 0,
  });

  const [showPhoneNumbers, setShowPhoneNumbers] = useState<boolean>(false);
  const [showErrorMessage, setShowErrorMessage] = useState<boolean>(false);
  const [currentPage, setCurrentPage] = useState<string>("search-results");
  // const [isClient, setIsClient] = useState<boolean>(false);
  const [professionals, setProfessionals] = useState<IProfissional[]>([]);
  // const [profession, setProfession] = useState<string>("");
  const isPodeTodos = false;
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
      const selectedCity = location.state.selectedCity;
      const selectedProfessional = location.state.selectedProfessional;

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

  // Busca cidades e profissões ao mudar estado
  useEffect(() => {
    setSelectedCity("");
    const fetchData = async () => {
      if (!selectedState) return;

      try {
        const citiesRes = await CityService.citiesByStatePublic({
          uf: selectedState,
        });
        if (citiesRes.status === 200) setcitiesByState(citiesRes.data);
      } catch (error) {
        console.error("Erro ao buscar cidades ou profissões:", error);
      }
    };

    fetchData();
  }, [selectedState]);

  const handleRequestQuote = async (
    professional: IProfissional | null = null,
    bulk = false
  ) => {
    // if (localStorage.getItem("id") != null) {
    setSelectedProfessional(professional as IProfissional | null);
    setIsBulkRequest(bulk);
    // setShowLGPDTerms(true);
    // const result = await ClientService.getClientbyID(
    //   parseInt(localStorage.getItem("id") ?? "0")
    // );
    // console.log(result);

    // if ((result.status == 200)) {
    //   const json = await result.data;
    //   formData.name = json.nome;
    //   formData.email = json.email;
    //   formData.phone = json.telefone;
    //   formData.clientId = json.oid;
    // }
    // setIsClient(true);
    // } else {
    setShowProfessionalSearch(false);
    // setShowLogin(true);
    //      setCurrentPage("login");
    //onNavigate && onNavigate("login");
  };

  const handleAcceptTerms = () => {
    setShowLGPDTerms(false);
    setShowContactForm(true);
    // console.log("Selected_Professionals");
    // console.log(selectedProfessional);
    // console.log("------");
  };

  const handleRejectTerms = () => {
    setShowLGPDTerms(false);
    setShowErrorMessage(true);
    setSelectedProfessional(null);
    setIsBulkRequest(false);
    setTimeout(() => {
      setShowErrorMessage(false);
    }, 3000);
  };

  const handleFormSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    // console.log("Form submitted:", formData);
    // console.log(selectedProfessional);
    let profs: string[];
    if (selectedProfessional == null) {
      // console.log(professionals);
      profs = professionals.map((prof) => prof.oid.toString());
    } else {
      profs = [selectedProfessional.oid.toString()];
    }

    const budget: IBudget = {
      name: formData.name,
      email: formData.email,
      telephone: formData.phone,
      // clientId: formData.clientId,
      description: formData.message,
      termResponsabilityAccepted: true,
      cityId: parseInt(selectedCity),
      professionalsId: [],
    };
    // adicionar na lista o id do profissional selecionado
    //  profs é um array de string, preciso converter para number
    budget.professionalsId = profs.map(Number);

    const postReturn = await BudgetService.saveBudget(budget);

    if (postReturn.status == 200) {
      setShowContactForm(false);
      setShowProfessionalSearch(true);
      //setShowPhoneNumbers(true);

      // Limpar os campos do formulário após envio bem-sucedido
      setFormData({
        name: "",
        email: "",
        phone: "",
        message: "",
        clientId: 0,
        cityId: 0,
      });
      setSelectedState("");
      setSelectedCity("");
      setcitiesByState([]);

      // Mostrar mensagem de sucesso com SweetAlert2
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

    const cityOptions = citiesByState.map(
    (city: ICitySearchProfessionals) => ({
      value: city.id,
      label: city.name,
    })
  );

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
      {/* Success Message - Removido pois agora usa SweetAlert2 */}
      {/* Error Message */}
      {showErrorMessage && (
        <div className="fixed top-4 right-4 bg-red-600 text-white px-6 py-3 rounded-lg shadow-lg z-50 animate-slide-up">
          <div className="flex items-center gap-2">
            <Info className="w-5 h-5" />
            <p>
              É necessário aceitar os termos para prosseguir com a solicitação.
            </p>
          </div>
        </div>
      )}
      {/* LGPD Terms Modal */}
      {/* {showLogin && (
        /*
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
          <div className="bg-white rounded-lg shadow-xl w-full max-w-2xl max-h-[90vh] overflow-y-auto">
            <div className="p-6">
              <button
                onClick={() => {
                  setShowLogin(false);
                }}
                className="text-gray-400 hover:text-gray-600"
              >
                <X className="w-5 h-5" />
              </button>
              <Login onNavigate="home"></Login>
            </div>
          </div>
        </div>

        // <Login onNavigate={setCurrentPage}></Login>
      )} 
      */}
      {/* LGPD Terms Modal */}
      {showLGPDTerms && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
          <div className="bg-white rounded-lg shadow-xl w-full max-w-2xl max-h-[90vh] overflow-y-auto">
            <div className="p-6">
              <div className="flex justify-between items-start mb-4">
                <h3 className="text-xl font-bold text-gray-900">
                  Termo de Responsabilidade e Isenção de Responsabilidade por
                  Dados Fornecidos - LGPD
                </h3>
                <button
                  onClick={handleRejectTerms}
                  className="text-gray-400 hover:text-gray-600"
                >
                  <X className="w-5 h-5" />
                </button>
              </div>

              <div className="prose prose-sm max-w-none text-gray-600 space-y-4">
                <p>
                  Por favor, leia atentamente o seguinte termo antes de
                  prosseguir:
                </p>

                <p>
                  Eu, o Cliente, ao utilizar os serviços fornecidos pela plataforma digital HASSIS CONECTA, reconheço e concordo com os termos e condições estabelecidos neste documento.
                </p>

                <h4 className="font-semibold text-gray-900">
                  Responsabilidade pelos Dados Fornecidos:
                </h4>
                <p>
                  Eu reconheço e concordo que sou totalmente responsável por quaisquer dados pessoais, informações ou conteúdos que eu solicite, receba ou de qualquer forma obtenha através da plataforma HASSIS CONECTA.
                </p>

                <h4 className="font-semibold text-gray-900">
                  Isenção de Responsabilidade da Plataforma:
                </h4>
                <p>
                  A responsabilidade pela proteção e tratamento adequado dos
                  dados pessoais é exclusivamente do Cliente.
                </p>

                <h4 className="font-semibold text-gray-900">
                  Finalidade e Consentimento:
                </h4>
                <p>
                  Eu reconheço que a HASSIS CONECTA pode coletar, armazenar e utilizar meus dados pessoais conforme necessário para a prestação de serviços ou cumprimento de obrigações contratuais, desde que em conformidade com as disposições da LGPD e mediante consentimento explícito do titular dos dados, quando aplicável.
                </p>

                <h4 className="font-semibold text-gray-900">
                  Segurança dos Dados:
                </h4>
                <p>
                  A HASSIS CONECTA adota medidas técnicas e organizacionais adequadas para proteger os dados pessoais contra acesso não autorizado, uso indevido, divulgação, alteração e destruição não autorizados, em conformidade com as disposições da LGPD.
                </p>

                <h4 className="font-semibold text-gray-900">
                  Direitos dos Titulares dos Dados:
                </h4>
                <p>
                  Eu reconheço e concordo em respeitar os direitos dos titulares dos dados, conforme previsto na LGPD, incluindo o direito de acesso, retificação, exclusão, anonimização, portabilidade e revogação do consentimento.
                </p>

                <h4 className="font-semibold text-gray-900">Indenização:</h4>
                <p>
                  Eu concordo em indenizar e isentar a HASSIS CONECTA, seus diretores, funcionários e agentes de qualquer responsabilidade, perda, reclamação ou despesa (incluindo honorários advocatícios razoáveis) decorrentes ou relacionados com o tratamento de dados pessoais pelo Cliente ou com o uso da plataforma.
                </p>

                <p>
                  Ao clicar no botão "Concordo", eu confirmo que li, entendi e concordo com os termos e condições estabelecidos neste Termo de Responsabilidade e Isenção de Responsabilidade por Dados Fornecidos.
                </p>
              </div>

              <div className="mt-6 flex justify-end gap-4">
                <button
                  onClick={handleRejectTerms}
                  className="px-4 py-2 text-gray-600 hover:text-gray-900"
                >
                  Recusar
                </button>
                <button
                  onClick={() => {
                    handleRequestQuote(selectedProfessional, false);
                    handleAcceptTerms();
                  }}
                  className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
                >
                  Concordo
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
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
                      {selectedProfessional && selectedProfessional.nome
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
                    htmlFor="email"
                    className="block text-sm font-medium text-gray-700 mb-1"
                  >
                    Email
                  </label>
                  <div className="relative">
                    <Mail className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                    <input
                      type="email"
                      id="email"
                      name="email"
                      value={formData.email}
                      onChange={handleInputChange}
                      // disabled={isClient}
                      required
                      className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-50"
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
                  {/* Estado */}
                  <div>
                    <label
                      htmlFor="state"
                      className="block text-sm font-medium text-gray-700 mb-1"
                    >
                      Estado
                    </label>
                    <div className="relative">
                      <Building2 className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                      <select
                        id="state"
                        value={selectedState}
                        onChange={(e) => setSelectedState(e.target.value)}
                        className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none bg-white mb-4"
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

                  {/* Cidade */}
                  <div>
                    <label
                      htmlFor="city"
                      className="block text-sm font-medium text-gray-700 mb-1"
                    >
                      Cidade
                    </label>
                    <div className="relative">
                       <MapPin className="absolute left-3 top-3 text-gray-400 w-5 h-5 z-10" />
                      <Select
                          options={cityOptions}
                          value={
                            cityOptions.find(
                              (option) => option.value.toString() === selectedCity
                            ) || null
                          }
                          onChange={(selectedOption) =>
                            setSelectedCity(selectedOption?.value.toString() || "")
                          }
                          isDisabled={!selectedState}
                          placeholder="Digite ou selecione a cidade"
                          isSearchable
                          className="text-sm"
                          styles={{
                            control: (base) => ({
                              ...base,
                              minHeight: "42px",
                              paddingLeft: "30px",
                            }),
                          }}
                        />
                    </div>
                  </div>
                </div>

                <div>
                  <label
                    htmlFor="message"
                    className="block text-sm font-medium text-gray-700 mb-1"
                  >
                    Descreva com detalhes o serviço
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

                <button
                  type="submit"
                  className="w-full bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 transition-colors"
                >
                  Enviar Mensagem
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
              {isPodeTodos && (
                <button
                  onClick={() => handleRequestQuote(null, true)}
                  className="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
                >
                  Solicitar Orçamento para Todos
                </button>
              )}
            </div>
          </div>

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
                        <span className="text-xs font-semibold text-white bg-green-500 py-1 px-3 rounded-full">
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
                      onClick={() => {
                        setShowLGPDTerms(true);
                        setSelectedProfessional(professional);
                      }}
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
                    onClick={() => {
                      setShowLGPDTerms(true);
                      setSelectedProfessional(professional);
                    }}
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
                    setShowLGPDTerms(true);
                    setSelectedProfessional(selectedProfessional);
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
