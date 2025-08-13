import React, { useState } from "react";
import {
  ArrowLeft,
  Tag,
  Truck,
  Shield,
  Info,
  Phone,
  X,
  User,
  Mail,
  MapPin,
  MessageSquare,
  Check,
} from "lucide-react";
import InputMask from "react-input-mask";
import { states } from "../data";
import { ContactService } from "../services/ContactService";
import { IContact } from "../interfaces";
import Swal from "sweetalert2";
import { CityService } from "../services/CityService";
import { ClientService } from "../services/ClientService";

interface ProductDetailsProps {
  onBack: () => void;
  product: {
    id: number;
    title: string;
    price: number;
    originalPrice: number;
    image: string;
    discount: number;
    description?: string;
    category?: string;
    professionalID: number;
    storeID: number;
  };
}

function ProductDetails({ product, onBack }: ProductDetailsProps) {
  const [showContactPopup, setShowContactPopup] = useState(false);
  const [showLGPDTerms, setShowLGPDTerms] = useState(false);
  const [showSuccessMessage, setShowSuccessMessage] = useState(false);
  const [showErrorMessage, setShowErrorMessage] = useState(false);
  const [selectedState, setSelectedState] = useState("");
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    phone: "",
    city: "",
    message: "",
    acceptTerms: false,
  });
  const [citiesByState, setCitiesByState] = useState([{}]);
  const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";

  //console.log(product);

  React.useEffect(() => {
    const fetchData = async () => {
      if (isLoggedIn) {
        const result = await ClientService.getClientbyID(
          parseInt(localStorage.getItem("id") ?? "0")
        );
        console.log(result);

        if ((result.status == 200)) {
          const json = await result.data;
          formData.name = json.nome;
          formData.email = json.email;
          formData.phone = json.telefone;
          formData.city = json.cidade.oid;
          //formData.clientId = json.oid;
          setSelectedState(json.cidade.uf);
        }
      }
      
      const citiesByState = await CityService.citiesByStatePublic({
        uf: selectedState,
      });

      const json_cities = await citiesByState.data;
      console.log(json_cities);
      if ((citiesByState.status == 200)) {
        setCitiesByState(json_cities);
      }
      
    };
    fetchData();
  }),
    [selectedState];

  const handleChange = (
    e: React.ChangeEvent<
      HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement
    >
  ) => {
    const { name, value, type } = e.target;
    const checked = (e.target as HTMLInputElement).checked;

    setFormData((prev) => ({
      ...prev,
      [name]: type === "checkbox" ? checked : value,
    }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (isLoggedIn)  {
      handleAcceptTerms();
    }
    else {
      setShowLGPDTerms(true);
    }
  };

  const handleAcceptTerms = async () => {
    setFormData((prev) => ({ ...prev, acceptTerms: true }));
    setShowLGPDTerms(false);
    setShowContactPopup(false);
    setShowSuccessMessage(true);
    const contact: IContact = {
      id: 0,
      name: formData.name,
      telefone: formData.phone,
      email: formData.email,
      mensagem: formData.message,
      status: "ativo",
      city_id: parseInt(formData.city),
      professional_id: product.professionalID,
      client_id: isLoggedIn ? (parseInt(localStorage.getItem("post_id") ?? "")) : null,
      store_id: product.storeID,
      product_id: product.id,
    };
    //Send Message
    console.log(contact);
    try {
      const response = await ContactService.saveContact(contact);
      //const json = await response.data;
      if ((response.status == 200)) {
        Swal.fire({
          position: "center",
          icon: "success",
          title: "Mensagem enviada!!!",
          text: "O vendedor logo entrará em contato",
          confirmButtonColor: "#2563eb",
          //showConfirmButton: false,
          //timer: 2000,
        });
      }

      setTimeout(() => {
        setShowSuccessMessage(false);
      }, 3000);
    }
    catch (error) {
      Swal.fire({
        position: "center",
        icon: "error",
        title:
          "Erro ao enviar mensagem"+error,
        showConfirmButton: true,
      });
    }    
  };

  const handleRejectTerms = () => {
    setShowLGPDTerms(false);
    setShowErrorMessage(true);
    setTimeout(() => {
      setShowErrorMessage(false);
    }, 3000);
  };

  return (
    <div className="min-h-screen bg-gray-50 py-12">
      {/* Success Message Popup */}
      {showSuccessMessage && (
        <div className="fixed top-4 right-4 bg-green-600 text-white px-6 py-3 rounded-lg shadow-lg z-50 animate-slide-up">
          <div className="flex items-center gap-2">
            <Check className="w-5 h-5" />
            <p>Sua mensagem foi enviada ao vendedor. Aguarde o contato!</p>
          </div>
        </div>
      )}

      {/* Error Message Popup */}
      {showErrorMessage && (
        <div className="fixed top-4 right-4 bg-red-600 text-white px-6 py-3 rounded-lg shadow-lg z-50 animate-slide-up">
          <div className="flex items-center gap-2">
            <Info className="w-5 h-5" />
            <p>É necessário aceitar os termos para enviar a mensagem.</p>
          </div>
        </div>
      )}

      {/* Contact Popup */}
      {showContactPopup && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
          <div className="bg-white rounded-lg shadow-xl w-full max-w-md relative">
            <button
              onClick={() => setShowContactPopup(false)}
              className="absolute top-2 right-2 text-gray-400 hover:text-gray-600 transition-colors"
            >
              <X className="w-5 h-5" />
            </button>

            <div className="p-4">
              <div className="flex items-center gap-2 mb-4">
                <User className="w-6 h-6 text-blue-600" />
                <div>
                  <h3 className="text-lg font-bold text-gray-900">
                    {
                      "" /*(product.professionalID == 0 && "Construir Mais Barato") ||
                      product.professional.Name*/
                    }
                  </h3>
                  <p className="text-xs text-gray-600">Vendedor</p>
                </div>
              </div>

              <form onSubmit={handleSubmit} className="space-y-3">
                {/* Nome Completo */}
                <div>
                  <div className="relative">
                    <User className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-4 h-4" />
                    <input
                      type="text"
                      id="name"
                      name="name"
                      value={formData.name}
                      onChange={handleChange}
                      disabled={isLoggedIn}
                      required
                      className="block w-full pl-9 pr-3 py-2 text-sm border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100"
                      placeholder="Nome Completo"
                    />
                  </div>
                </div>

                {/* Email */}
                <div>
                  <div className="relative">
                    <Mail className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-4 h-4" />
                    <input
                      type="email"
                      id="email"
                      name="email"
                      value={formData.email}
                      onChange={handleChange}
                      disabled={isLoggedIn}
                      required
                      className="block w-full pl-9 pr-3 py-2 text-sm border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100"
                      placeholder="Email"
                    />
                  </div>
                </div>

                {/* WhatsApp */}
                <div>
                  <div className="relative">
                    <Phone className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-4 h-4" />
                    <InputMask
                      mask="(99) 99999-9999"
                      type="tel"
                      id="phone"
                      name="phone"
                      value={formData.phone}
                      onChange={handleChange}
                      disabled={isLoggedIn}
                      required
                      className="block w-full pl-9 pr-3 py-2 text-sm border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100"
                      placeholder="WhatsApp"
                    />
                  </div>
                </div>

                {/* Estado/Cidade */}
                <div className="grid grid-cols-2 gap-2">
                  <div>
                    <div className="relative">
                      <MapPin className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-4 h-4" />
                      <select
                        id="state"
                        name="state"
                        value={selectedState}
                        disabled={isLoggedIn}
                        onChange={(e) => setSelectedState(e.target.value)}
                        required
                        className="block w-full pl-9 pr-3 py-2 text-sm border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none disabled:bg-gray-100"
                      >
                        <option value="">Estado</option>
                        {states.map((state) => (
                          <option key={state.id} value={state.id}>
                            {state.name}
                          </option>
                        ))}
                      </select>
                    </div>
                  </div>

                  <div>
                    <div className="relative">
                      <MapPin className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-4 h-4" />
                      <select
                        id="city"
                        name="city"
                        value={formData.city}
                        onChange={handleChange}
                        required
                        disabled={!selectedState || isLoggedIn}
                        className="block w-full pl-9 pr-3 py-2 text-sm border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none disabled:bg-gray-100"
                      >
                        <option value="">Cidade</option>
                        {selectedState &&
                          citiesByState.map((city: any) => (
                            <option key={city.id} value={city.id}>
                              {city.name}
                            </option>
                          ))}
                      </select>
                    </div>
                  </div>
                </div>

                {/* Mensagem */}
                <div>
                  <div className="relative">
                    <MessageSquare className="absolute left-3 top-3 text-gray-400 w-4 h-4" />
                    <textarea
                      id="message"
                      name="message"
                      value={formData.message}
                      onChange={handleChange}
                      required
                      rows={3}
                      className="block w-full pl-9 pr-3 py-2 text-sm border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 resize-none"
                      placeholder="Descreva sua mensagem sobre o produto"
                    />
                  </div>
                </div>

                <button
                  type="submit"
                  className="w-full bg-blue-600 text-white py-2 px-4 rounded-md text-sm font-medium hover:bg-blue-700 transition-colors"
                >
                  Enviar Mensagem
                </button>

                <div className="flex items-start gap-2 text-xs text-gray-500 border-t border-gray-200 pt-3">
                  <Info className="w-4 h-4 flex-shrink-0 mt-0.5" />
                  <p className="italic">
                    *A Plataforma Marketplace Comprar mais barato é uma mera
                    plataforma de divulgação, sem nenhum vínculo com os
                    anunciantes.
                  </p>
                </div>
              </form>
            </div>
          </div>
        </div>
      )}

      {/* LGPD Terms Popup */}
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
                  Responsabilidade pelas Negociações no Marketplace:
                </h4>
                <p>
                  A CONSTRUIR MAIS BARATO é uma plataforma que permite que
                  usuários anunciem e vendam seus produtos diretamente entre si.
                  A plataforma não possui qualquer responsabilidade sobre a
                  qualidade, garantias, defeitos ou qualquer outro aspecto
                  relacionado aos produtos anunciados. Toda e qualquer
                  negociação realizada entre compradores e vendedores é de
                  exclusiva responsabilidade das partes envolvidas.
                </p>

                <h4 className="font-semibold text-gray-900">
                  Responsabilidade pelos Dados Fornecidos:
                </h4>
                <p>
                  Eu reconheço e concordo que sou totalmente responsável por
                  quaisquer dados pessoais, informações ou conteúdos que eu
                  solicite, receba ou de qualquer forma obtenha através da
                  plataforma.
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
                  contratuais.
                </p>

                <h4 className="font-semibold text-gray-900">
                  Segurança dos Dados:
                </h4>
                <p>
                  A CONSTRUIR MAIS BARATO adota medidas técnicas e
                  organizacionais adequadas para proteger os dados pessoais
                  contra acesso não autorizado, uso indevido, divulgação,
                  alteração e destruição não autorizados.
                </p>

                <h4 className="font-semibold text-gray-900">
                  Direitos dos Titulares dos Dados:
                </h4>
                <p>
                  Eu reconheço e concordo em respeitar os direitos dos titulares
                  dos dados, conforme previsto na LGPD.
                </p>

                <h4 className="font-semibold text-gray-900">Indenização:</h4>
                <p>
                  Eu concordo em indenizar e isentar a CONSTRUIR MAIS BARATO de
                  qualquer responsabilidade, perda, reclamação ou despesa
                  decorrentes do tratamento de dados pessoais pelo Cliente.
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
                  onClick={handleAcceptTerms}
                  className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
                >
                  Li e Concordo
                </button>
              </div>
            </div>
          </div>
        </div>
      )}

      <div className="max-w-7xl mx-auto px-4">
        <button
          onClick={onBack}
          className="flex items-center gap-2 text-gray-600 hover:text-blue-600 mb-8 transition-colors"
        >
          <ArrowLeft className="w-5 h-5" />
          Voltar para o Marketplace
        </button>

        <div className="bg-white rounded-xl shadow-lg overflow-hidden">
          <div className="grid md:grid-cols-2 gap-8 p-8">
            {/* Product Image */}
            <div className="relative">
              <div className="aspect-square rounded-lg overflow-hidden">
                <img
                  src={
                    product.image
                      ? "data:image;base64," + product.image
                      : "/images/default_product.png"
                  }
                  alt={product.title}
                  className="w-full h-full object-cover"
                />
              </div>
              {product.discount > 0 && (
                <div className="absolute top-4 right-4 bg-red-600 text-white px-3 py-1 rounded-full text-sm font-semibold">
                  -{product.discount}%
                </div>
              )}
            </div>

            {/* Product Info */}
            <div className="space-y-6">
              <div>
                <h1 className="text-3xl font-bold text-gray-900 mb-2">
                  {product.name}
                </h1>
                <div className="flex items-center gap-2">
                  <Tag className="w-5 h-5 text-blue-600" />
                  <span className="text-gray-600">
                    {product.category.Name || "Ferramentas"}
                  </span>
                </div>
              </div>

              <div className="space-y-2">
                <div className="flex items-center gap-3">
                  {product.price !== 0 &&
                    product.originalPrice > product.price && (
                      <span className="text-gray-500 line-through text-lg">
                        R${" "}
                        {new Intl.NumberFormat("pt-BR", {
                          style: "decimal",
                          minimumFractionDigits: 2,
                        }).format(product.originalPrice)}
                      </span>
                    )}
                  <span className="text-3xl font-bold text-blue-600">
                    {product.price !== 0
                      ? product.price.toLocaleString("pt-BR", {
                          style: "currency",
                          currency: "BRL",
                          minimumFractionDigits: 2,
                        })
                      : ""}
                  </span>
                </div>
                {/*<p className="text-sm text-gray-600">Em até 12x sem juros</p>*/}
              </div>

              <div className="space-y-4">
                <h2 className="text-xl font-semibold text-gray-900">
                  Descrição do Produto
                </h2>
                <p className="text-gray-600">
                  {product.description ||
                    "Produto de alta qualidade, ideal para profissionais da construção civil. Fabricado com materiais resistentes e duráveis, garantindo maior vida útil e melhor custo-benefício."}
                </p>
              </div>

              {/*<div className="space-y-4 pt-6 border-t border-gray-200">
                <div className="flex items-center gap-3">
                  <Truck className="w-5 h-5 text-green-600" />
                  <div>
                    <h3 className="font-medium text-gray-900">Frete Grátis</h3>
                    <p className="text-sm text-gray-600">
                      Para compras acima de R$ 299,00
                    </p>
                  </div>
                </div>
                <div className="flex items-center gap-3">
                  <Shield className="w-5 h-5 text-blue-600" />
                  <div>
                    <h3 className="font-medium text-gray-900">
                      Garantia de 12 meses
                    </h3>
                    <p className="text-sm text-gray-600">
                      Contra defeitos de fabricação
                    </p>
                  </div>
                </div>
              </div>
              */}

              <button
                onClick={() => setShowContactPopup(true)}
                className="w-full bg-blue-600 text-white py-3 px-6 rounded-lg font-medium hover:bg-blue-700 transition-colors"
              >
                Entrar em Contato com Vendedor
              </button>

              <div className="pt-6 border-t border-gray-200">
                <div className="flex items-start gap-2 text-sm text-gray-500">
                  <Info className="w-5 h-5 flex-shrink-0 mt-0.5" />
                  <p className="italic">
                    *A Plataforma Marketplace Comprar mais barato é uma mera
                    plataforma de divulgação, sem nenhum vínculo com os
                    anunciantes.
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default ProductDetails;
