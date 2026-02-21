import React, { useState, useEffect } from "react";
import {
  Building2,
  MapPin,
  HardHat,
  MessageSquare,
  User,
  Mail,
  Phone,
} from "lucide-react";
import { states } from "../data";
import { CityService } from "../services/CityService";
import { ProfessionalService } from "../services/ProfessionalService";
import { ProfessionService } from "../services/ProfessionService";
import { BannerService } from "../services/BannerService";
import InputMask from "react-input-mask";
import { ProductCategoryService, RegionService } from "../services";
import { StoreService } from "../services/StoreService";
import { useNavigate } from "react-router-dom";
import {
  IBannerSearchProfessionals,
  IBudget,
  ICitySearchProfessionals,
  IProfissional,
  IStore,
} from "../interfaces";
import { ICategoryProduct } from "../interfaces/ICategoryProduct";
import Swal from "sweetalert2";
import { BudgetService } from "../services/Budget";
import { ArrowLeft } from "lucide-react";

interface SearchProductsProps {
  onNavigate?: (page: string) => void;
}
interface FormData {
  name: string;
  email: string;
  phone: string;
  message: string;
  clientId: number;
  cityId: number;
}

// URL_IMAGES_WEB do .env
const URL_IMAGES_WEB = import.meta.env.VITE_URL_IMAGES_WEB;

function QuoteProducts({ onNavigate }: SearchProductsProps) {
  const [formData, setFormData] = useState<FormData>({
    name: "",
    email: "",
    phone: "",
    message: "",
    clientId: 0,
    cityId: 0,
  });
  const [selectedState, setSelectedState] = useState<string>("");
  const [selectedCity, setSelectedCity] = useState<string>("");
  const [selectedCategoryProduct, setSelectedCategoryProduct] =
    useState<string>("");
  const [citiesByState, setcitiesByState] = useState<
    ICitySearchProfessionals[]
  >([]);
  const [products, SetProducts] = useState<IProfissional[]>([]);
  const [categoryProducts, setCategoryProducts] = useState<ICategoryProduct[]>(
    []
  );
  const [storeList, setStores] = useState<IStore[]>([]);
  const [selectedSubcategories, setSelectedSubcategories] = useState<number[]>(
    []
  );
  const [showModal, setShowModal] = useState(false);
  const [imageModal, setImageModal] =
    useState<IBannerSearchProfessionals | null>(null);
  const navigate = useNavigate();

  // Função para gerar número aleatório
  const gerarNumeroAleatorio = (min: number, max: number): number => {
    return Math.floor(Math.random() * (max - min + 1)) + min;
  };

  // Função para construir URL da imagem
  function getImageUrl(encodedPath: string) {
    try {
      if (!encodedPath || encodedPath.trim() === "") return "";

      // decodifica o Base64 para obter o path real
      const decodedPath = atob(encodedPath); // ex: "/images/upload/upload-3341225764.png"

      if (!decodedPath || decodedPath.trim() === "") return "";

      const baseUrl = URL_IMAGES_WEB?.replace(/\/$/, ""); // remove barra final

      if (!baseUrl) return "";

      return `${baseUrl}${
        decodedPath.startsWith("/") ? "" : "/"
      }${decodedPath}`;
    } catch (error) {
      console.error("Erro ao decodificar path da imagem:", error);
      return "";
    }
  }

  // Função para validar se a imagem existe
  const validateImageExists = (imageUrl: string): Promise<boolean> => {
    return new Promise((resolve) => {
      if (!imageUrl) {
        resolve(false);
        return;
      }

      const img = new Image();
      img.onload = () => resolve(true);
      img.onerror = () => resolve(false);
      img.src = imageUrl;
    });
  };

  // Busca cidades e categorias de produtos ao mudar estado
  useEffect(() => {
    setSelectedCity("");
    const fetchData = async () => {
      if (!selectedState) return;

      try {
        const citiesRes = await CityService.citiesByStatePublic({
          uf: selectedState,
        });
        if (citiesRes.status === 200) setcitiesByState(citiesRes.data);

        handleSearchCategoriesProduct();

        // const professionsRes = await ProfessionService.getProfessionsPublic();
        // if (professionsRes.status === 200) setCategoryProducts(professionsRes.data);
      } catch (error) {
        console.error("Erro ao buscar cidades ou profissões:", error);
      }
    };

    fetchData();
  }, [selectedState]);

  // Limpa subcategorias selecionadas quando mudar a categoria
  useEffect(() => {
    setSelectedSubcategories([]);
  }, [selectedCategoryProduct]);

  // Pega subcategorias da categoria selecionada (vem do children)
  const subcategories = selectedCategoryProduct
    ? categoryProducts.find(
        (cat) => cat.id === parseInt(selectedCategoryProduct)
      )?.children || []
    : [];

  // Função para alternar seleção de subcategoria
  const toggleSubcategory = (subcategoryId: number) => {
    setSelectedSubcategories((prev) => {
      if (prev.includes(subcategoryId)) {
        return prev.filter((id) => id !== subcategoryId);
      } else {
        return [...prev, subcategoryId];
      }
    });
  };

  // Função para selecionar/desselecionar todas as subcategorias
  const toggleAllSubcategories = () => {
    if (selectedSubcategories.length === subcategories.length) {
      setSelectedSubcategories([]);
    } else {
      setSelectedSubcategories(subcategories.map((sub) => sub.id));
    }
  };
  
  const cleanForm = () => {
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
    setCategoryProducts([]);
    setSelectedCategoryProduct("standardCategory")
    setSelectedSubcategories([]);
  }

  const handleFormSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const quoteMessage= "Preciso dos preços destes produtos, os itens selecionados: "
    const currentFormData = { ...formData };
    const currentCities = [...citiesByState];
    const currentSelectedCity = selectedCity;
    const currentSelectedCategory = selectedCategoryProduct;
    const currentSelectedSubCategories = subcategories.filter(item => selectedSubcategories.includes(item.id)).map(item => item.name)
    const resultSubCategories = currentSelectedSubCategories.join(', ');

    const loadedStores = await handleOpenModal();

    // Verificar se realmente temos lojas carregadas
    if (loadedStores.length === 0) {
      Swal.fire({
        position: "center",
        icon: "warning",
        title: "Nenhuma loja encontrada",
        text: "Não há lojas disponíveis para as subcategorias selecionadas.",
        showConfirmButton: true,
      });
      return;
    }

    let stores: string[];
    stores = loadedStores.map((store) => store.oid.toString());

    const budget: IBudget = {
      name: currentFormData.name,
      email: currentFormData.email,
      telephone: currentFormData.phone,
      storesId: [],
      description: quoteMessage+resultSubCategories,
      // description: "Preciso dos preços destes produtos, os itens selecionados",
      termResponsabilityAccepted: true,
      cityId: parseInt(currentSelectedCity),
      professionalsId: [],
      approved: false,
    };
    budget.storesId = stores.map(Number);

    const postReturn = await BudgetService.saveBudget(budget);

    if (postReturn.status == 200) {
      Swal.fire({
        position: "center",
        icon: "success",
        title: "Solicitação enviada com sucesso!",
        text: "Os lojistas entrarão em contato em breve.",
        showConfirmButton: false,
        timer: 1500,
        
      });
      cleanForm()
      // Filtrar lojistas premium usando loadedStores
      //       const premiumStores = loadedStores.filter(
      //         (store) => store.isPremiumStore === true
      //       );

      //       // Preparar mensagem de orçamento ANTES de limpar
      //       const cityName =
      //         currentCities.find((c) => c.id === parseInt(currentSelectedCity))
      //           ?.name || "";
      //       const categoryName =
      //         categoryProducts.find((c) => c.id === parseInt(currentSelectedCategory))
      //           ?.name || "";

      //       const budgetMessage = `Olá! Gostaria de solicitar um orçamento.

      // *Dados do Cliente:*
      // Nome: ${currentFormData.name}
      // WhatsApp: ${currentFormData.phone}
      // ${currentFormData.email ? `Email: ${currentFormData.email}` : ""}
      // Cidade: ${cityName}

      // *Categoria:* ${categoryName}

      // ${currentFormData.message ? `*Detalhes:*\n${currentFormData.message}` : ""}

      // Orçamento enviado através de https://construirmaisbarato.com.br/`;

      //       const encodedMessage = encodeURIComponent(budgetMessage);

      //       // Limpar os campos do formulário DEPOIS de salvar os dados
      //       setFormData({
      //         name: "",
      //         email: "",
      //         phone: "",
      //         message: "",
      //         clientId: 0,
      //         cityId: 0,
      //       });
      //       setSelectedState("");
      //       setSelectedCity("");
      //       setcitiesByState([]);

      //       // Se tem lojistas premium, redirecionar para WhatsApp
      //       if (premiumStores.length > 0) {
      //         // Se só tem 1 lojista premium, redireciona direto
      //         if (premiumStores.length === 1) {
      //           Swal.fire({
      //             position: "center",
      //             icon: "success",
      //             title: "Orçamento enviado!",
      //             text: "Você será redirecionado para o WhatsApp do lojista premium.",
      //             showConfirmButton: false,
      //             timer: 2000,
      //           });

      //           const phone = (
      //             premiumStores[0].telefone || premiumStores[0].Telephone
      //           ).replace(/\D/g, "");
      //           window.open(
      //             `https://wa.me/55${phone}?text=${encodedMessage}`,
      //             "_blank"
      //           );
      //         } else {
      //           // Se tem múltiplos lojistas premium, mostra opções
      //           const storeOptions: any = {};
      //           premiumStores.forEach((store) => {
      //             const storeName = store.nome || store.Name;
      //             const storePhone = store.telefone || store.Telephone;
      //             storeOptions[store.oid] = `${storeName} - ${storePhone}`;
      //           });

      //           Swal.fire({
      //             title: "✅ Orçamento Enviado!",
      //             html: "<p>Escolha um lojista premium para contato direto via WhatsApp:</p>",
      //             input: "select",
      //             inputOptions: storeOptions,
      //             inputPlaceholder: "Selecione um lojista",
      //             showCancelButton: true,
      //             confirmButtonText: "Abrir WhatsApp",
      //             cancelButtonText: "Fechar",
      //             preConfirm: (selectedStoreId) => {
      //               if (!selectedStoreId) {
      //                 Swal.showValidationMessage("Selecione um lojista!");
      //               }
      //               return selectedStoreId;
      //             },
      //           }).then((result) => {
      //             if (result.isConfirmed) {
      //               const selectedStore = premiumStores.find(
      //                 (s) => s.oid === parseInt(result.value)
      //               );
      //               if (selectedStore) {
      //                 const phone = (
      //                   selectedStore.telefone || selectedStore.Telephone
      //                 ).replace(/\D/g, "");
      //                 window.open(
      //                   `https://wa.me/55${phone}?text=${encodedMessage}`,
      //                   "_blank"
      //                 );
      //               }
      //             }
      //           });
      //         }
      //       } else {
      //         // Não tem lojistas premium - mensagem padrão
      //         Swal.fire({
      //           position: "center",
      //           icon: "success",
      //           title: "Solicitação enviada com sucesso!",
      //           text: "Os lojistas entrarão em contato em breve.",
      //           showConfirmButton: false,
      //           timer: 3000,
      //         });
      //       }
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

  // Função para abrir modal e buscar imagem
  const handleOpenModal = async (): Promise<IStore[]> => {
    if (!selectedCity) return [];

    try {
      const regionRes = await RegionService.getRegionbyCity(
        parseInt(selectedCity)
      );
      if (regionRes.status !== 200) {
        return await handleSearchProducts();
      }

      const bannerRes = await BannerService.getBannerByPagePublic({
        page: "B",
        cityId: 0,
        regionId: regionRes.data.id,
      });

      if (bannerRes.status === 200 && bannerRes.data.length > 0) {
        const randomIndex = gerarNumeroAleatorio(0, bannerRes.data.length - 1);
        const selectedBanner = bannerRes.data[randomIndex];
        const imageUrl = getImageUrl(selectedBanner.image);

        if (imageUrl) {
          const imageExists = await validateImageExists(imageUrl);
          if (imageExists) {
            setShowModal(true);
            setImageModal(selectedBanner);
            // Ainda buscar produtos mesmo mostrando banner
            return await handleSearchProducts();
          }
        }
      }
    } catch (error) {
      console.error("Erro ao abrir modal:", error);
    }

    return await handleSearchProducts();
  };

  const handleSearchCategoriesProduct = async () => {
    // localStorage.setItem("search_city", selectedCity);
    try {
      const response = await ProductCategoryService.findTopLevelCategory();
      setCategoryProducts(response.data || []);
    } catch (error) {
      console.error("Erro ao buscar categorias:", error);
    }
  };
  const handleSearchProducts = async (): Promise<IStore[]> => {
    // Salva dados no localStorage
    localStorage.setItem("search_city", selectedCity);
    localStorage.setItem("search_category", selectedCategoryProduct);
    localStorage.setItem(
      "search_subcategories",
      JSON.stringify(selectedSubcategories)
    );

    try {
      const response = await StoreService.getByCategoryAndSubCategories({
        categoryId: parseInt(selectedCategoryProduct),
        subCategoriesId: selectedSubcategories,
        cityId: parseInt(selectedCity),
      });
      setStores(response.data);
      // ✅ Retornar as lojas diretamente
      return response.data || [];
    } catch (error) {
      console.error("Erro ao buscar categorias:", error);
      return [];
    }
  };

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Modal */}
      {showModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-70">
          <div className="relative h-auto max-h-[95vh] rounded-lg overflow-hidden shadow-xl bg-black flex items-center justify-center">
            <button
              onClick={() => {
                setShowModal(false);
                handleSearchProducts();
              }}
              className="absolute top-4 right-4 z-10 text-white text-2xl bg-black bg-opacity-50 hover:bg-opacity-70 rounded-full w-10 h-10 flex items-center justify-center"
            >
              ✕
            </button>

            <img
              src={getImageUrl(imageModal?.image || "")}
              alt="Imagem de Construção"
              onClick={() =>
                imageModal?.link
                  ? window.open(imageModal.link, "_blank")
                  : undefined
              }
              className="max-h-full object-contain max-w-[100vw] sm:max-w-[90vw] lg:max-w-[80vw] xl:max-w-[70vw] w-auto sm:h-[95vh] xs:h-auto"
              style={{ cursor: imageModal ? "pointer" : undefined }}
            />
          </div>
        </div>
      )}

      {/* Formulário de seleção */}
      
      <div className="max-w-4xl mx-auto px-4 py-12">
        
        <div className="bg-white rounded-xl shadow-lg p-3">
          <button 
            onClick={() => 
            navigate("/")}
            className="flex items-center gap-2 text-orange-600 hover:text-orange-900 transition-colors"
          >
            <ArrowLeft className="w-5 h-5" />
            Voltar para início
          </button>
          <h1 className="text-3xl font-bold text-gray-900 mb-2">
            Encontrar Produtos
          </h1>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            {/* Estado */}
            <div className="relative">
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Estado
              </label>
              <div className="relative">
                <Building2 className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                <select
                  value={selectedState}
                  onChange={(e) => setSelectedState(e.target.value)}
                  className="block w-full pl-10 pr-4 py-2.5 text-gray-900 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none bg-white"
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
            <div className="relative">
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Cidade
              </label>
              <div className="relative">
                <MapPin className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                <select
                  value={selectedCity}
                  onChange={(e) => setSelectedCity(e.target.value)}
                  disabled={!selectedState}
                  className="block w-full pl-10 pr-4 py-2.5 text-gray-900 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none bg-white disabled:bg-gray-100 disabled:cursor-not-allowed"
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

            {/* CategoryProducts */}
            <div className="relative">
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Categoria de Produtos
              </label>
              <div className="relative">
                <HardHat className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                <select
                  value={selectedCategoryProduct}
                  onChange={(e) => setSelectedCategoryProduct(e.target.value)}
                  className="block w-full pl-10 pr-4 py-2.5 text-gray-900 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none bg-white"
                >
                  <option value="standardCategory">Seleciona uma Categoria</option>
                  {categoryProducts.map((cp: ICategoryProduct) => (
                    <option key={cp.id} value={cp.id}>
                      {cp.name}
                    </option>
                  ))}
                </select>
              </div>
            </div>
          </div>

          {/* Card de Subcategorias */}
          {subcategories.length > 0 && (
            <div className="mt-6 bg-gradient-to-br from-blue-50 to-indigo-50 rounded-xl p-6 border-2 border-blue-200 shadow-lg">
              <div className="flex items-center justify-between mb-4">
                <h3 className="text-lg font-semibold text-gray-800 flex items-center gap-2">
                  <HardHat className="w-5 h-5 text-blue-600" />
                  Subcategorias Disponíveis
                </h3>
                <button
                  onClick={toggleAllSubcategories}
                  className="text-sm text-blue-600 hover:text-blue-800 font-medium transition-colors"
                >
                  {selectedSubcategories.length === subcategories.length
                    ? "Desmarcar Todas"
                    : "Selecionar Todas"}
                </button>
              </div>

              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3">
                {subcategories.map((subcategory) => (
                  <label
                    key={subcategory.id}
                    className={`
                      flex items-center gap-3 p-3 rounded-lg cursor-pointer transition-all
                      ${
                        selectedSubcategories.includes(subcategory.id)
                          ? "bg-blue-600 text-white shadow-md scale-105"
                          : "bg-white text-gray-700 hover:bg-blue-100 hover:shadow-sm"
                      }
                    `}
                  >
                    <input
                      type="checkbox"
                      checked={selectedSubcategories.includes(subcategory.id)}
                      onChange={() => toggleSubcategory(subcategory.id)}
                      className="w-5 h-5 rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                    />
                    <span className="font-medium text-sm">
                      {subcategory.name}
                    </span>
                  </label>
                ))}
              </div>

              {selectedSubcategories.length > 0 && (
                <div className="mt-4 pt-4 border-t border-blue-200">
                  <p className="text-sm text-gray-600">
                    <span className="font-semibold text-blue-600">
                      {selectedSubcategories.length}
                    </span>{" "}
                    subcategoria{selectedSubcategories.length !== 1 ? "s" : ""}{" "}
                    selecionada
                    {selectedSubcategories.length !== 1 ? "s" : ""}
                  </p>
                </div>
              )}
            </div>
          )}
          {selectedSubcategories.length > 0 && (
            <div className="mt-6 bg-gradient-to-br from-blue-50 to-indigo-50 rounded-xl p-6 border-2 border-blue-200 shadow-lg">
              <h3 className="text-lg font-semibold text-gray-800 flex items-center gap-2 mb-2">
                <HardHat className="w-5 h-5 text-blue-600" />
                Suas Informações de Contato
              </h3>
              <form onSubmit={handleFormSubmit} className="space-y-4">
                {/* <div>
                  <label
                    htmlFor="name"
                    className="block text-sm font-medium text-gray-700 mb-1"
                  >
                    Nome Completo *
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
                </div> */}

                <div>
                  <label
                    htmlFor="phone"
                    className="block text-sm font-medium text-gray-700 mb-1"
                  >
                    WhatsApp *
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
                    htmlFor="email"
                    className="block text-sm font-medium text-gray-700 mb-1"
                  >
                    Email (Opcional)
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
                      className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-50"
                    />
                  </div>
                </div> 
                <button
                  // onClick={handleOpenModal}
                  type="submit"
                  disabled={
                    !selectedState ||
                    !selectedCity ||
                    selectedSubcategories.length === 0 ||
                    formData.phone === ""
                  }
                  className="mt-8 w-full bg-blue-600 text-white py-3 px-6 rounded-lg font-medium hover:bg-blue-700 transition-colors disabled:bg-gray-300 disabled:cursor-not-allowed"
                >
                  Orçar com Lojistas Parceiros
                </button>
              </form>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default QuoteProducts;
