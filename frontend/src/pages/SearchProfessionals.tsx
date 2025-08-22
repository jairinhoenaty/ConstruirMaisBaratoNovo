

import React, { useState, useEffect } from "react";
import { Building2, MapPin, HardHat } from "lucide-react";
import { states } from "../data";
import { CityService } from "../services/CityService";
import { ProfessionalService } from "../services/ProfessionalService";
import { ProfessionService } from "../services/ProfessionService";
import { BannerService } from "../services/BannerService";
import { RegionService } from "../services";
import { useNavigate } from "react-router-dom";
import { IBannerSearchProfessionals, ICitySearchProfessionals, IProfissional } from "../interfaces";
import { IProfessionSearchProfessionals } from "../interfaces/IProfession";

interface SearchProfessionalsProps {
  onNavigate?: (page: string) => void;
}

// URL_IMAGES_WEB do .env
const URL_IMAGES_WEB = import.meta.env.VITE_URL_IMAGES_WEB;

function SearchProfessionals({ onNavigate }: SearchProfessionalsProps) {
  const [selectedState, setSelectedState] = useState<string>("");
  const [selectedCity, setSelectedCity] = useState<string>("");
  const [selectedProfessional, setSelectedProfessional] = useState<string>("");
  const [citiesByState, setcitiesByState] = useState<ICitySearchProfessionals[]>([]);
  const [professionals, setProfessionals] = useState<IProfissional[]>([]);
  const [professions, setProfessions] = useState<IProfessionSearchProfessionals[]>([]);
  const [showModal, setShowModal] = useState(false);
  const [imageModal, setImageModal] = useState<IBannerSearchProfessionals | null>(null);
  const navigate = useNavigate();

  // Função para gerar número aleatório
  const gerarNumeroAleatorio = (min: number, max: number): number => {
    return Math.floor(Math.random() * (max - min + 1)) + min;
  };

  // Função para construir URL da imagem
  function getImageUrl(encodedPath: string) {
    try {
      // decodifica o Base64 para obter o path real
      const decodedPath = atob(encodedPath); // ex: "/images/upload/upload-3341225764.png"
      const baseUrl = URL_IMAGES_WEB?.replace(/\/$/, ""); // remove barra final
      return `${baseUrl}${decodedPath.startsWith("/") ? "" : "/"}${decodedPath}`;
    } catch (error) {
      console.error("Erro ao decodificar path da imagem:", error);
      return "";
    }
  }

  // Busca cidades e profissões ao mudar estado
  useEffect(() => {
    setSelectedCity("");
    const fetchData = async () => {
      if (!selectedState) return;

      try {
        const citiesRes = await CityService.citiesByStatePublic({ uf: selectedState });
        if (citiesRes.status === 200) setcitiesByState(citiesRes.data);

        const professionsRes = await ProfessionService.getProfessionsPublic();
        if (professionsRes.status === 200) setProfessions(professionsRes.data);


      } catch (error) {
        console.error("Erro ao buscar cidades ou profissões:", error);
      }
    };

    fetchData();
  }, [selectedState]);

  // Função para abrir modal e buscar imagem
  const handleOpenModal = async () => {
    if (!selectedCity) return;
  
    setShowModal(true); // abre o modal imediatamente
  
    try {
      const regionRes = await RegionService.getRegionbyCity(parseInt(selectedCity));
      if (regionRes.status !== 200) return handleSearch();
  
      const bannerRes = await BannerService.getBannerByPagePublic({
        page: "B",
        cityId: 0,
        regionId: regionRes.data.id,
      });
  
      if (bannerRes.status === 200 && bannerRes.data.length > 0) {
        const randomIndex = gerarNumeroAleatorio(0, bannerRes.data.length - 1);
        setImageModal(bannerRes.data[randomIndex]);
      } else {
        setImageModal(null); // modal ainda abre, mas sem imagem
      }
    } catch (error) {
      console.error("Erro ao abrir modal:", error);
      setImageModal(null); // modal ainda abre
    }
  };

  // Função de busca de profissionais
  const handleSearch = async () => {
    localStorage.setItem("search_city", selectedCity);
    try {
      const return_professionals = await ProfessionalService.getProfessionalByCityAndProfession({
        cityID: parseInt(selectedCity),
        professionID: parseInt(selectedProfessional),
        limit: 1000,
        offset: 0,
      });

      setProfessionals(return_professionals.data.profissionais || []);

      navigate("/search-results", {
        state: {
          selectedCity,
          selectedProfessional,
        },
      });
    } catch (error) {
      console.error("Erro ao buscar profissionais:", error);
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
                handleSearch();
              }}
              className="absolute top-4 right-4 z-10 text-white text-2xl bg-black bg-opacity-50 hover:bg-opacity-70 rounded-full w-10 h-10 flex items-center justify-center"
            >
              ✕
            </button>

            <img
              src={getImageUrl(imageModal?.image || "")}
              alt="Imagem de Construção"
              onClick={() =>
                imageModal?.link ? window.open(imageModal.link, "_blank") : undefined
              }
              className="max-h-full object-contain max-w-[100vw] sm:max-w-[90vw] lg:max-w-[80vw] xl:max-w-[70vw] w-auto sm:h-[95vh] xs:h-auto"
              style={{ cursor: imageModal ? "pointer" : undefined }}
            />
          </div>
        </div>
      )}

      {/* Formulário de seleção */}
      <div className="max-w-4xl mx-auto px-4 py-12">
        <div className="bg-white rounded-xl shadow-lg p-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-8">
            Encontre Profissionais da Construção Civil
          </h1>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            {/* Estado */}
            <div className="relative">
              <label className="block text-sm font-medium text-gray-700 mb-2">Estado</label>
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
              <label className="block text-sm font-medium text-gray-700 mb-2">Cidade</label>
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

            {/* Profissional */}
            <div className="relative">
              <label className="block text-sm font-medium text-gray-700 mb-2">Profissional</label>
              <div className="relative">
                <HardHat className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                <select
                  value={selectedProfessional}
                  onChange={(e) => setSelectedProfessional(e.target.value)}
                  className="block w-full pl-10 pr-4 py-2.5 text-gray-900 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none bg-white"
                >
                  <option value="">Selecione o profissional</option>
                  {professions.map((prof: IProfessionSearchProfessionals) => (
                    <option key={prof.id} value={prof.id}>
                      {prof.name}
                    </option>
                  ))}
                </select>
              </div>
            </div>
          </div>

          <button
            onClick={handleOpenModal}
            disabled={!selectedState || !selectedCity || !selectedProfessional}
            className="mt-8 w-full bg-blue-600 text-white py-3 px-6 rounded-lg font-medium hover:bg-blue-700 transition-colors disabled:bg-gray-300 disabled:cursor-not-allowed"
          >
            Buscar Profissionais
          </button>
        </div>
      </div>
    </div>
  );
}

export default SearchProfessionals;
