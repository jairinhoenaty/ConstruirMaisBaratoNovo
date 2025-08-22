import React, { useEffect, useState } from "react";
import {
  ArrowLeft,
  // Star,
  Phone,
  Mail,
  User,
  MessageSquare,
  X,
  Check,
  Users,
  Info,
  ShieldCheck,
  Building2,
  MapPin,
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
import { IBudget, ICitySearchProfessionals, IProfissional } from "../interfaces";
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

function SearchResults() {
  /*{
  profession,
  professionals,
  onNewSearch,
}: SearchResultsProps*/
  const [selectedState, setSelectedState] = useState<string>("");
  const [selectedCity, setSelectedCity] = useState<string>("");
  const [citiesByState, setcitiesByState] = useState<ICitySearchProfessionals[]>([]);
  const [showLGPDTerms, setShowLGPDTerms] = useState(false);
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
  const [showSuccessMessage, setShowSuccessMessage] = useState<boolean>(false);
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

  useEffect(() => {
    const fetchData = async () => {
      // console.log("Effect");
      // console.log(location);
      // console.log(location.state.selectedProfessional);
      const selectedCity = location.state.selectedCity;
      const selectedProfessional = location.state.selectedProfessional;
      const return_professionals =
        await ProfessionalService.getProfessionalByCityAndProfession({
          cityID: parseInt(selectedCity),
          professionID: parseInt(selectedProfessional),
          limit: 1000,
          offset: 0,
        });

      const json_professionals = await return_professionals.data.profissionais;
      setProfessionals(json_professionals);
      // setProfession(selectedProfessional);
    };

    fetchData();
  }, []);



    // Busca cidades e profissões ao mudar estado
    useEffect(() => {
      setSelectedCity("");
      const fetchData = async () => {
        if (!selectedState) return;
  
        try {
          const citiesRes = await CityService.citiesByStatePublic({ uf: selectedState });
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
    console.log("Enviando orçamento");
    console.log(profs);


    const budget: IBudget = {
      name: formData.name,
      email: formData.email,
      telephone: formData.phone,
      // clientId: formData.clientId,
      description: formData.message,
      termResponsabilityAccepted: true,
      cityId: parseInt(selectedCity),
      professionalsId:[]
    };
    // adicionar na lista o id do profissional selecionado
    //  profs é um array de string, preciso converter para number
    budget.professionalsId = profs.map(Number);

    const postReturn = await BudgetService.saveBudget(budget);

    if (postReturn.status == 200) {
      console.log("Retorno do cadastro de orçamento postReturn  ===> ", postReturn);
      setShowContactForm(false);
      setShowSuccessMessage(true);
      setShowProfessionalSearch(true);
      //setShowPhoneNumbers(true);
      setTimeout(() => {
        setShowSuccessMessage(false);
      }, 3000);
    } else {
      Swal.fire({
        position: "center",
        icon: "error",
        title: "Erro ao enviar orçamentos",
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
      {/* Success Message */}
      {showSuccessMessage && (
        <div className="fixed top-4 right-4 bg-green-600 text-white px-6 py-3 rounded-lg shadow-lg z-50 animate-slide-up">
          <div className="flex items-center gap-2">
            <Check className="w-5 h-5" />
            <p>
              Solicitação enviada com sucesso! Os profissionais entrarão em
              contato.
            </p>
          </div>
        </div>
      )}
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
                  Eu, o Cliente, ao utilizar os serviços fornecidos pela
                  plataforma digital CONSTRUIR MAIS BARATO, reconheço e concordo
                  com os termos e condições estabelecidos neste documento.
                </p>

                <h4 className="font-semibold text-gray-900">
                  Responsabilidade pelos Dados Fornecidos:
                </h4>
                <p>
                  Eu reconheço e concordo que sou totalmente responsável por
                  quaisquer dados pessoais, informações ou conteúdos que eu
                  solicite, receba ou de qualquer forma obtenha através da
                  plataforma CONSTRUIR MAIS BARATO.
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
                  Eu reconheço que a CONSTRUIR MAIS BARATO pode coletar,
                  armazenar e utilizar meus dados pessoais conforme necessário
                  para a prestação de serviços ou cumprimento de obrigações
                  contratuais, desde que em conformidade com as disposições da
                  LGPD e mediante consentimento explícito do titular dos dados,
                  quando aplicável.
                </p>

                <h4 className="font-semibold text-gray-900">
                  Segurança dos Dados:
                </h4>
                <p>
                  A CONSTRUIR MAIS BARATO adota medidas técnicas e
                  organizacionais adequadas para proteger os dados pessoais
                  contra acesso não autorizado, uso indevido, divulgação,
                  alteração e destruição não autorizados, em conformidade com as
                  disposições da LGPD.
                </p>

                <h4 className="font-semibold text-gray-900">
                  Direitos dos Titulares dos Dados:
                </h4>
                <p>
                  Eu reconheço e concordo em respeitar os direitos dos titulares
                  dos dados, conforme previsto na LGPD, incluindo o direito de
                  acesso, retificação, exclusão, anonimização, portabilidade e
                  revogação do consentimento.
                </p>

                <h4 className="font-semibold text-gray-900">Indenização:</h4>
                <p>
                  Eu concordo em indenizar e isentar a CONSTRUIR MAIS BARATO,
                  seus diretores, funcionários e agentes de qualquer
                  responsabilidade, perda, reclamação ou despesa (incluindo
                  honorários advocatícios razoáveis) decorrentes ou relacionados
                  com o tratamento de dados pessoais pelo Cliente ou com o uso
                  da plataforma.
                </p>

                <p>
                  Ao clicar no botão "Concordo", eu confirmo que li, entendi e
                  concordo com os termos e condições estabelecidos neste Termo
                  de Responsabilidade e Isenção de Responsabilidade por Dados
                  Fornecidos.
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
                    setShowContactForm(false)
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
                      <MapPin className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                      <select
                        id="city"
                        value={selectedCity}
                        onChange={(e) => setSelectedCity(e.target.value)}
                        disabled={!selectedState}
                        className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none bg-white disabled:bg-gray-100 disabled:cursor-not-allowed"
                      >
                        <option value="">Selecione a cidade</option>
                        {citiesByState.map((city: ICitySearchProfessionals) => (
                          <option key={city.id} value={city.id}>
                            {city.name}
                          </option>
                        ))}
                      </select>
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
                className="flex flex-col sm:flex-row items-start sm:items-center gap-4 p-4 bg-white border border-gray-200 rounded-lg hover:border-blue-500 transition-colors"
              >
                {professional.verified && (
                  <div className="flex items-center gap-2">
                    <ShieldCheck
                      className="w-5 h-5 text-green-500"
                      title="Profissional Verificado"
                      aria-label="Profissional Verificado"
                    />
                    {/* Badge de Profissional Verificado */}
                    <span className="text-xs font-semibold text-white bg-green-500 py-1 px-3 rounded-full">
                      Profissional Verificado
                    </span>
                  </div>
                )}

                <div className="flex-1">
                  <h3 className="text-lg font-semibold text-gray-900">
                    {professional.nome}
                  </h3>
                  <div className="flex items-center gap-2 text-sm text-gray-600">
                    <div className="flex items-center"></div>
                    <span>
                      {professional.cidade.nome}, {professional.cidade.uf}
                    </span>
                  </div>
                  {showPhoneNumbers &&
                    (isBulkRequest ||
                      selectedProfessional?.oid === professional.oid) && (
                      <div className="mt-2 flex items-center gap-2 text-green-600">
                        <Phone className="w-4 h-4" />
                        <a
                          href={`https://wa.me/${professional.telephone.replace(
                            /\D/g,
                            ""
                          )}`}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="hover:text-green-700 transition-colors"
                        >
                          {professional.telephone}
                        </a>
                      </div>
                    )}
                </div>
                <button
                  onClick={() => {
                    setShowLGPDTerms(true);
                    setSelectedProfessional(professional);
                    
                  }}
                  className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
                >
                  Solicitar Orçamento
                </button>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}

export default SearchResults;
