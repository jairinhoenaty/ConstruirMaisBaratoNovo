import React, { useState } from "react";
import {
  User,
  Mail,
  Phone,
  MapPin,
  Building2,
  Key,
  MessageSquare,
  Package,
  //LogOut,
  Save,
  CreditCard,
  FileText,
  UserCircle,
  PlusCircle,
  Edit2,
  Tag,
  Clock,
  Smartphone,
  X,
  HelpCircle,
  Send,
  HardHat,
  ChevronDown,
  Check,
  Trash2,
} from "lucide-react";
import InputMask from "react-input-mask";
//import { states, citiesByState } from '../data';
//import { professionals_data } from "../data";
import { states } from "../data";
import QuotesPanel from "./QuotesPanel";
import ClientMessages from "./ClientMessages";
import PasswordPanel from "./PasswordPanel";
import { ProfessionalService } from "../services/ProfessionalService";
import { CityService, ProductService, ProfessionService } from "../services";
import Swal from "sweetalert2";
import NewProduct from "./NewProduct";
import { StoreService } from "../services/StoreService";
import { ClientService } from "../services/ClientService";
import { IContact } from "../interfaces/IContact";
import { ContactService } from "../services/ContactService";
import Pagination from "../components/Pagination";
import EditProduct from "./EditProduct";

interface Product {
  id: string;
  name: string;
  price: number;
  image: string;
  status: "pending" | "approved";
}

function ProfessionalPanel() {
  const [activeTab, setActiveTab] = useState("profile");
  const [showPopup, setShowPopup] = useState(false); //Colocar True para aparecer o modal do Proffisional Já
  const [showProfessions, setShowProfessions] = useState(false);
  const [supportMessage, setSupportMessage] = useState("");
  const [citiesByState, setcitiesByState] = useState([{ id: 0, name: "" }]);
  const profile = localStorage.getItem("profile");

  const [professions, setProfessions] = useState([
    { id: "", name: "", description: "" },
  ]);
  const [formData, setFormData] = useState({
    fullName: "",
    email: "",
    whatsapp: "",
    cep: "",
    street: "",
    neighborhood: "",
    city: "",
    state: "",
    professions: [] as string[],
  });

  const [isProfessional, setIsProfessional] = useState(false);
  const [isClient, setIsClient] = useState(false);
  const [products, setProducts] = useState([
    {
      id: "0",
      name: "",
      price: 0,
      image: "",
      status: "",
      dayoffer: false,
      professionalID: 0,
      storeID: 0,
    },
  ]);
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);
  const [productID, setProductID] = useState(0);
  const [isUpdate, setIsUpdate] = useState(false);
  const [menuItems, setMenuItems] = useState([
    { id: "products", label: "Produtos", icon: Package },
    { id: "messages", label: "Mensagens", icon: MessageSquare },
    { id: "support", label: "Suporte", icon: HelpCircle },
    { id: "password", label: "Senha de Acesso", icon: Key },

    //{ id: "logout", label: "Sair do Painel", icon: LogOut },
  ]);

  const [verified,setVerified] = useState(false); // Variável para verificar se o profissional está verificado  

  React.useEffect(() => {
    const fetchData = async () => {
      let json;
      const id = parseInt(localStorage.getItem("id") || "");
      if (profile == "client") {
        const response = await ClientService.getClientbyID(id);
        const json = await response.data;

        if ((response.status == 200)) {
          localStorage.setItem("post_id", response.data.oid);
          localStorage.setItem("city_id", response.data.cidade.oid);
          setFormData((prev) => ({
            ...prev,
            fullName: json.nome,
            email: json.email,
            whatsapp: json.telefone,
            cep: json.cep,
            street: json.endereco,
            neighborhood: json.bairro,
            city: json.cidade.oid,
            state: json.cidade.uf,
          }));
          setIsClient(true);
          setMenuItems([
            { id: "messages", label: "Mensagens", icon: MessageSquare },
            { id: "support", label: "Suporte", icon: HelpCircle },
            { id: "password", label: "Senha de Acesso", icon: Key },
          ]);
          const citiesByState = await CityService.citiesByState({
            uf: json.cidade.uf,
          });

          const json_cities = await citiesByState.data;
          if ((citiesByState.status == 200)) {
            setcitiesByState(json_cities);
          }
          setFormData((prev) => ({
            ...prev,
            city: json.cidade.oid,

            //professions: json.profissoes.map((x: any) => x.oid),
          }));
        }
        //} else if (profile == "profissional" || profile == "admin") {
      } else if (profile == "profissional") {
        const professions = await ProfessionService.getProfessions();
        const json_professions = await professions.data.profissoes;
        if ((professions.status == 200)) {
          setProfessions(json_professions);
        }

        const response = await ProfessionalService.getProfessionalbyID(id);
        const json = await response.data;
        if ((response.status == 200)) {
          localStorage.setItem("post_id", response.data.oid);
          localStorage.setItem("city_id", response.data.cidade.oid);
          setVerified(json.verified);

          setFormData((prev) => ({
            ...prev,
            fullName: json.nome,
            email: json.email,
            whatsapp: json.telefone,
            cep: json.cep,
            street: json.endereco,
            neighborhood: json.bairro,
            city: json.cidade.oid,
            state: json.cidade.uf,
            professions: json.profissoes.map((x: any) => x.oid),
          }));
          setIsProfessional(true);
          const citiesByState = await CityService.citiesByState({
            uf: json.cidade.uf,
          });

          const json_cities = await citiesByState.data;
          if ((citiesByState.status == 200)) {
            setcitiesByState(json_cities);
          }
          setFormData((prev) => ({
            ...prev,
            city: json.cidade.oid,

            //professions: json.profissoes.map((x: any) => x.oid),
          }));
        }
      } else if (profile == "store") {
        const response = await StoreService.getStorebyID(id);
        const json = await response.data;
        if ((response.status == 200)) {
          localStorage.setItem("post_id", response.data.oid);
          localStorage.setItem("city_id", response.data.cidade.oid);
          setFormData((prev) => ({
            ...prev,
            fullName: json.nome,
            email: json.email,
            whatsapp: json.telefone,
            cep: json.cep,
            street: json.endereco,
            neighborhood: json.bairro,
            city: json.cidade.oid,
            state: json.cidade.uf,
            //professions: json.profissoes.map((x: any) => x.oid),
          }));
          const citiesByState = await CityService.citiesByState({
            uf: json.cidade.uf,
          });

          const json_cities = await citiesByState.data;
          if ((citiesByState.status == 200)) {
            setcitiesByState(json_cities);
          }
          setFormData((prev) => ({
            ...prev,
            city: json.cidade.oid,

            //professions: json.profissoes.map((x: any) => x.oid),
          }));
        }
      }

      if (activeTab === "products") {
        let limit = 9;
        let result_products: any;

        result_products = await ProductService.productsAllPublic(
          limit,
          (page - 1) * limit,
          profile=="profissional"?parseInt(localStorage.getItem("post_id") ?? ""):0,
          profile=="store"?parseInt(localStorage.getItem("post_id") ?? ""):0,
          null,
          null
        );

        if (result_products.status == 200) {
          const json_products = await result_products.data.products;
          setTotalPage(Math.ceil(result_products.data.total / limit));
          setProducts(json_products);
          renderProducts();
        }
      }
    };
    fetchData();
  }, [activeTab, page, isUpdate]);

  const handleChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    if (name == "state") {
      const citiesByState = await CityService.citiesByState({
        uf: value,
      });

      const json_cities = await citiesByState.data;
      if ((citiesByState.status == 200)) {
        setcitiesByState(json_cities);
      }
    }

    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const profile = localStorage.getItem("profile");
    const post_id = localStorage.getItem("post_id") ?? "";
    let postReturn: any;
    if (profile == "client") {
      postReturn = await ClientService.postClient({
        oid: parseInt(post_id),
        Name: formData.fullName,
        Email: formData.email,
        Telephone: formData.whatsapp,
        //LgpdAceito: "S",
        //created_at:  "time.Date(2025, time.March, 16, 19, 41, 30, 309000000, time.Local)",
        cep: formData.cep,
        street: formData.street,
        neighborhood: formData.neighborhood,
        cityId: parseInt(formData.city),
        LgpdAceito: "S",
        Password: null,
        image: null,
      });
    } else if (profile == "profissional") {
      postReturn = await ProfessionalService.postProfessional({
        oid: parseInt(post_id),
        Name: formData.fullName,
        Email: formData.email,
        Telephone: formData.whatsapp,
        //LgpdAceito: "S",
        //created_at:  "time.Date(2025, time.March, 16, 19, 41, 30, 309000000, time.Local)",
        cep: formData.cep,
        street: formData.street,
        neighborhood: formData.neighborhood,
        cityId: parseInt(formData.city),
        professionIds: formData.professions,
        Password: null,
        image: null,
      });
    } else if (profile == "store") {
      postReturn = await StoreService.postStore({
        oid: parseInt(post_id),
        Name: formData.fullName,
        Email: formData.email,
        Telephone: formData.whatsapp,
        //LgpdAceito: "S",
        //created_at:  "time.Date(2025, time.March, 16, 19, 41, 30, 309000000, time.Local)",
        cep: formData.cep,
        street: formData.street,
        neighborhood: formData.neighborhood,
        cityId: parseInt(formData.city),
        LgpdAceito: "S",
        Password: null,
        image: null,
      });
    }

    if (postReturn.status == 200) {
      Swal.fire({
        position: "center",
        icon: "success",
        title: "Conta Atualizada!!!",
        showConfirmButton: false,
        timer: 1500,
      });
    }
    console.log("Form submitted:", formData);
  };

  const handleSupportSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    console.log("Support message sent:", supportMessage);
    const profile = localStorage.getItem("profile");
    const contact: IContact = {
      id: 0,
      name: formData.fullName,
      telefone: formData.whatsapp,
      email: formData.email,
      mensagem: supportMessage,
      status: "ativo",
      professional_id:
        profile == "profissional"
          ? parseInt(localStorage.getItem("post_id") ?? "")
          : null,
      client_id:
        profile == "client"
          ? parseInt(localStorage.getItem("post_id") ?? "")
          : null,
      store_id:
        profile == "store"
          ? parseInt(localStorage.getItem("post_id") ?? "")
          : null,
      product_id: null,
      city_id: null,
    };
    //console.log(contact);
    //Send Message
    try {
    const response = await ContactService.saveContact(contact);
    //const json = await response.data;
    if ((response.status == 200)) {
      Swal.fire({
        position: "center",
        icon: "success",
        title:
          "Mensagem enviada com sucesso! <br>Em breve nossa equipe entrará em contato.",
        showConfirmButton: false,
        timer: 3000,
      });
    }
    /*alert(
      "Mensagem enviada com sucesso! Em breve nossa equipe entrará em contato."
    );*/
    setSupportMessage("");
  }
  catch (error) {
    Swal.fire({
      position: "center",
      icon: "error",
      title:
        "Erro ao enviar mensagem"+error,
      showConfirmButton: false,
      timer: 3000,
    });
  }
  };

  const selectedProfessionsText =
    formData.professions.length > 0
      ? professions
          .filter((prof) => formData.professions.includes(prof.id))
          .map((prof) => prof.name)
          .join(", ")
      : "Selecione suas profissões";

  const toggleProfession = (professionId: string) => {
    setFormData((prev) => ({
      ...prev,
      professions: prev.professions.includes(professionId)
        ? prev.professions.filter((id) => id !== professionId)
        : [...prev.professions, professionId],
    }));
  };

  const handleDelete = async (productId: string) => {
    if (window.confirm("Tem certeza que deseja excluir este produto?")) {
      const response = await ProductService.deleteProduct(productId);
      if ((response.status == 200)) {
        Swal.fire({
          icon: "success",
          text: "Produto excluído",
          showConfirmButton: false,
          timer: 1500,
        });
        setIsUpdate(!isUpdate);
      }
      //setProducts(products.filter((product) => product.id !== productId));
    }
  };

  const renderProducts = () => (
    <div className="bg-white rounded-lg shadow-md p-8">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold text-gray-900">Meus Produtos</h2>
        {profile !== "admin" && (
          <button
            className="flex items-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            onClick={() => {
              setActiveTab("newproduct");
            }}
          >
            <PlusCircle className="w-5 h-5" />
            Novo Produto
          </button>
        )}
      </div>
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
      <br></br>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {products.map((product) => (
          <div
            key={product.id}
            className="bg-white rounded-lg shadow-md overflow-hidden border border-gray-200"
          >
            <div className="relative h-48">
              <img
                src={"data:image;base64," + product.image}
                alt={product.name}
                className="w-full h-full object-cover"
              />
              <div
                className={`absolute top-2 right-2 px-3 py-1 rounded-full text-sm font-medium ${
                  product.status === "approved"
                    ? "bg-green-100 text-green-800"
                    : "bg-yellow-100 text-yellow-800"
                }`}
              >
                {product.approved ? (
                  <div className="flex items-center gap-1">
                    <Tag className="w-4 h-4" />
                    <span>Aprovado</span>
                  </div>
                ) : (
                  <div className="flex items-center gap-1">
                    <Clock className="w-4 h-4" />
                    <span>Aguardando Aprovação</span>
                  </div>
                )}
              </div>
              {product.dayoffer && (
                <div className="absolute top-10 right-2 bg-red-600 text-white px-2 py-1 rounded-md text-sm font-semibold">
                  {product.dayoffer ? "Em destaque" : product.dayoffer}
                </div>
              )}
            </div>
            <div className="p-4">
              <h3 className="text-lg font-semibold text-gray-900 mb-2">
                {product.name}
              </h3>
              <div className="flex items-center justify-between">
                <span className="text-2xl font-bold text-blue-600">
                  R${" "}
                  {new Intl.NumberFormat("pt-BR", {
                    style: "decimal",
                    minimumFractionDigits: 2,
                  }).format(product.price)}
                </span>
                <button
                  onClick={() => {
                    //console.log("EDITAR");
                    setProductID(parseInt(product.id));
                    setActiveTab("editproduct");
                  }}
                  className="flex items-center gap-1 text-gray-600 hover:text-blue-600 transition-colors"
                >
                  <Edit2 className="w-4 h-4" />
                  <span>Editar</span>
                </button>
                <button
                  onClick={() => handleDelete(product.id)}
                  className="text-red-600 hover:bg-red-50 p-2 rounded-lg transition-colors"
                >
                  <Trash2 className="w-5 h-5" />
                </button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );

  const renderSupport = () => (
    <div className="bg-white rounded-lg shadow-md p-8">
      <div className="flex items-center gap-3 mb-8">
        <HelpCircle className="w-8 h-8 text-blue-600" />
        <h2 className="text-2xl font-bold text-gray-900">Suporte</h2>
      </div>

      <div className="max-w-2xl">
        <p className="text-gray-600 mb-6">
          Precisa de ajuda? Envie sua mensagem para nossa equipe de suporte.
          Responderemos o mais breve possível.
        </p>

        <form onSubmit={handleSupportSubmit} className="space-y-6">
          <div>
            <label
              htmlFor="supportMessage"
              className="block text-sm font-medium text-gray-700 mb-2"
            >
              Sua Mensagem
            </label>
            <textarea
              id="supportMessage"
              rows={6}
              value={supportMessage}
              onChange={(e) => setSupportMessage(e.target.value)}
              className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 resize-none"
              placeholder="Descreva em detalhes como podemos ajudar..."
              required
            />
          </div>

          <div className="flex items-center justify-between">
            <div className="text-sm text-gray-500">
              Nossa equipe responderá através do seu email cadastrado:{" "}
              {formData.email}
            </div>
            <button
              type="submit"
              className="flex items-center gap-2 px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            >
              <Send className="w-5 h-5" />
              Enviar Mensagem
            </button>
          </div>
        </form>

        <div className="mt-8 pt-8 border-t border-gray-200">
          <h3 className="text-lg font-semibold text-gray-900 mb-4">
            Canais de Atendimento
          </h3>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="p-4 border border-gray-200 rounded-lg">
              <div className="flex items-center gap-2 text-gray-900 mb-2">
                <Phone className="w-5 h-5 text-blue-600" />
                <span className="font-medium">WhatsApp</span>
              </div>
              <p className="text-gray-600">
                Segunda a Sexta, 9h às 18h
                <br />
                (14) 99166-5023
              </p>
            </div>
            <div className="p-4 border border-gray-200 rounded-lg">
              <div className="flex items-center gap-2 text-gray-900 mb-2">
                <Mail className="w-5 h-5 text-blue-600" />
                <span className="font-medium">Email</span>
              </div>
              <p className="text-gray-600">
                Atendimento 24/7
                <br />
                suporte@construirmaisbarato.com.br
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );

  const renderCashback = () => (
    <div className="bg-white rounded-lg shadow-md p-8">
      <div className="flex items-center gap-3 mb-8">
        <h2 className="text-2xl font-bold text-gray-900 text-center">
          Em Breve
        </h2>
      </div>
    </div>
  );

  return (
    <div className="min-h-screen bg-gray-50 py-12">
      {showPopup && (
        <div className="fixed bottom-4 right-4 bg-white rounded-lg shadow-xl p-6 max-w-sm animate-slide-up z-50">
          <button
            onClick={() => setShowPopup(false)}
            className="absolute top-2 right-2 text-gray-400 hover:text-gray-600 transition-colors"
          >
            <X className="w-5 h-5" />
          </button>
          <div className="flex items-center gap-3 mb-3">
            <div className="bg-blue-100 rounded-full p-2">
              <Smartphone className="w-6 h-6 text-blue-600" />
            </div>
            <h3 className="text-lg font-semibold text-gray-900">
              Seja um Profissional Já
            </h3>
          </div>
          <p className="text-gray-600 mb-4">
            Entre em contato com o{" "}
            <button
              onClick={() => {
                setShowPopup(false);
                setActiveTab("support");
              }}
              className="font-bold text-blue-600 hover:text-blue-700 transition-colors"
            >
              suporte
            </button>{" "}
            ou pelo nosso WhatsApp{" "}
            <a
              href="https://wa.me/14991665023"
              target="_blank"
              rel="noopener noreferrer"
              className="text-green-600 hover:text-green-700 transition-colors font-bold"
            >
              (14) 99166-5023
            </a>
          </p>
        </div>
      )}

      <div className="max-w-7xl mx-auto px-4">
        <h1 className="text-3xl font-bold text-gray-900 mb-8">
          Painel Administrativo
        </h1>

        <div className="bg-white rounded-lg shadow-md p-4 mb-6">
          <div className="flex flex-wrap gap-2">
            {menuItems.map((item) => {
              //console.log("VERIFIED: " + verified);
              if (verified == false && item.id === "products") {
                  return null; // Não renderiza o item se o profissional não estiver verificado
              }                  
              else {    
              const Icon = item.icon;
              return (
                <button
                  key={item.id}
                  onClick={() => setActiveTab(item.id)}
                  className={`flex items-center gap-2 px-4 py-2 rounded-lg transition-colors ${
                    activeTab === item.id
                      ? "bg-blue-600 text-white"
                      : "text-gray-700 hover:bg-gray-100"
                  }`}
                >
                  <Icon className="w-5 h-5" />
                  <span>{item.label}</span>
                </button>
              );
            }
            })}
          </div>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
          <button
            onClick={() => setActiveTab("profile")}
            className={`flex items-center justify-center gap-2 p-4 rounded-lg shadow-md transition-colors ${
              activeTab === "profile"
                ? "bg-blue-600 text-white"
                : "bg-white text-gray-700 hover:bg-gray-50"
            }`}
          >
            <UserCircle className="w-6 h-6" />
            <span className="font-medium">Perfil</span>
          </button>
          <button
            onClick={() => setActiveTab("quotes")}
            className={`flex items-center justify-center gap-2 p-4 rounded-lg shadow-md transition-colors ${
              activeTab === "quotes"
                ? "bg-blue-600 text-white"
                : "bg-white text-gray-700 hover:bg-gray-50"
            }`}
          >
            <FileText className="w-6 h-6" />
            <span className="font-medium">Orçamentos</span>
          </button>
          {!isClient && (
            <button
              onClick={() => setActiveTab("cashback")}
              className={`flex items-center justify-center gap-2 p-4 rounded-lg shadow-md transition-colors ${
                activeTab === "cashback"
                  ? "bg-blue-600 text-white"
                  : "bg-white text-gray-700 hover:bg-gray-50"
              }`}
            >
              <CreditCard className="w-6 h-6" />
              <span className="font-medium">Cashback</span>
            </button>
          )}
        </div>

        {activeTab === "profile" && (
          <div className="bg-white rounded-lg shadow-md p-8">
            <h2 className="text-2xl font-bold text-gray-900 mb-6">
              Dados da Conta
            </h2>
            <form onSubmit={handleSubmit} className="space-y-6">
              <div>
                <label
                  htmlFor="fullName"
                  className="block text-sm font-medium text-gray-700"
                >
                  Nome Completo
                </label>
                <div className="mt-1 relative rounded-md shadow-sm">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <User className="h-5 w-5 text-gray-400" />
                  </div>
                  <input
                    type="text"
                    name="fullName"
                    id="fullName"
                    value={formData.fullName}
                    onChange={handleChange}
                    className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>
              </div>
              <div>
                <label
                  htmlFor="email"
                  className="block text-sm font-medium text-gray-700"
                >
                  Email
                </label>
                <div className="mt-1 relative rounded-md shadow-sm">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <Mail className="h-5 w-5 text-gray-400" />
                  </div>
                  <input
                    type="email"
                    name="email"
                    id="email"
                    value={formData.email}
                    onChange={handleChange}
                    className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>
              </div>
              <div>
                <label
                  htmlFor="whatsapp"
                  className="block text-sm font-medium text-gray-700"
                >
                  WhatsApp
                </label>
                <div className="mt-1 relative rounded-md shadow-sm">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <Phone className="h-5 w-5 text-gray-400" />
                  </div>
                  <InputMask
                    mask="(99) 99999-9999"
                    type="tel"
                    name="whatsapp"
                    id="whatsapp"
                    value={formData.whatsapp}
                    onChange={handleChange}
                    className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>
              </div>
              <div>
                <label
                  htmlFor="cep"
                  className="block text-sm font-medium text-gray-700"
                >
                  CEP
                </label>
                <div className="mt-1 relative rounded-md shadow-sm">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <MapPin className="h-5 w-5 text-gray-400" />
                  </div>
                  <InputMask
                    mask="99999-999"
                    type="text"
                    name="cep"
                    id="cep"
                    value={formData.cep}
                    onChange={handleChange}
                    className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>
              </div>
              <div>
                <label
                  htmlFor="street"
                  className="block text-sm font-medium text-gray-700"
                >
                  Rua
                </label>
                <div className="mt-1 relative rounded-md shadow-sm">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <Building2 className="h-5 w-5 text-gray-400" />
                  </div>
                  <input
                    type="text"
                    name="street"
                    id="street"
                    value={formData.street}
                    onChange={handleChange}
                    className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>
              </div>
              <div>
                <label
                  htmlFor="neighborhood"
                  className="block text-sm font-medium text-gray-700"
                >
                  Bairro
                </label>
                <div className="mt-1 relative rounded-md shadow-sm">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <Building2 className="h-5 w-5 text-gray-400" />
                  </div>
                  <input
                    type="text"
                    name="neighborhood"
                    id="neighborhood"
                    value={formData.neighborhood}
                    onChange={handleChange}
                    className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>
              </div>
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <label
                    htmlFor="state"
                    className="block text-sm font-medium text-gray-700"
                  >
                    Estado
                  </label>
                  <div className="mt-1 relative rounded-md shadow-sm">
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                      <Building2 className="h-5 w-5 text-gray-400" />
                    </div>
                    <select
                      id="state"
                      name="state"
                      value={formData.state}
                      onChange={(e) => handleChange(e as any)}
                      className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
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

                <div>
                  <label
                    htmlFor="city"
                    className="block text-sm font-medium text-gray-700"
                  >
                    Cidade
                  </label>
                  <div className="mt-1 relative rounded-md shadow-sm">
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                      <Building2 className="h-5 w-5 text-gray-400" />
                    </div>
                    <select
                      id="city"
                      name="city"
                      value={formData.city}
                      onChange={(e) => handleChange(e as any)}
                      disabled={!formData.state}
                      className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100"
                    >
                      <option value="">Selecione a cidade</option>
                      {citiesByState &&
                        citiesByState.map((city) => (
                          <option key={city.id} value={city.id}>
                            {city.name}
                          </option>
                        ))}
                    </select>
                  </div>
                </div>
              </div>
              {/* Profissões */}
              {isProfessional && (
                <div className="relative">
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    Profissões
                  </label>
                  <button
                    type="button"
                    onClick={() => setShowProfessions(!showProfessions)}
                    className="relative w-full bg-white border border-gray-300 rounded-md shadow-sm pl-10 pr-10 py-2 text-left cursor-pointer focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500"
                  >
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                      <HardHat className="h-5 w-5 text-gray-400" />
                    </div>
                    <span className="block truncate">
                      {selectedProfessionsText}
                    </span>
                    <span className="absolute inset-y-0 right-0 flex items-center pr-2">
                      <ChevronDown
                        className={`h-5 w-5 text-gray-400 transition-transform ${
                          showProfessions ? "transform rotate-180" : ""
                        }`}
                      />
                    </span>
                  </button>

                  {showProfessions && (
                    <div className="absolute z-10 mt-1 w-full bg-white shadow-lg max-h-60 rounded-md py-1 text-base overflow-auto focus:outline-none sm:text-sm">
                      {professions.map((profession) => (
                        <div
                          key={profession.id}
                          className="relative cursor-pointer select-none py-2 pl-10 pr-4 hover:bg-blue-50"
                          onClick={() => toggleProfession(profession.id)}
                        >
                          <span
                            className={`block truncate ${
                              formData.professions.includes(profession.id)
                                ? "font-medium text-blue-600"
                                : "font-normal"
                            }`}
                          >
                            {profession.name}
                          </span>
                          {formData.professions.includes(profession.id) && (
                            <span className="absolute inset-y-0 left-0 flex items-center pl-3 text-blue-600">
                              <Check className="h-5 w-5" />
                            </span>
                          )}
                        </div>
                      ))}
                    </div>
                  )}
                </div>
              )}
              <div>
                <button
                  type="submit"
                  className="w-full flex items-center justify-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
                >
                  <Save className="w-5 h-5" />
                  Atualizar Conta
                </button>
              </div>
            </form>
          </div>
        )}
        {activeTab === "products" && renderProducts()}
        {activeTab === "support" && renderSupport()}
        {activeTab === "quotes" && <QuotesPanel />}
        {activeTab === "messages" && <ClientMessages />}
        {activeTab === "password" && <PasswordPanel />}
        {activeTab === "newproduct" && <NewProduct />}
        {activeTab === "editproduct" && <EditProduct id={productID} />}
        {activeTab === "cashback" && renderCashback()}
      </div>
    </div>
  );
}

export default ProfessionalPanel;
