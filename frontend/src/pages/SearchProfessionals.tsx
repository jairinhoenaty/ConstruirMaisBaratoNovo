import React, { useState, useEffect } from "react";
import { Building2, MapPin, HardHat } from "lucide-react";
import { states } from "../data";
// import SearchResults from "./SearchResults";
import { CityService } from "../services/CityService";
import { ProfessionalService } from "../services/ProfessionalService";
import { ProfessionService } from "../services/ProfessionService";
import { useNavigate } from "react-router-dom";
import { BannerService } from "../services/BannerService";
import { RegionService } from "../services";
import { ICity, IProfissional } from "../interfaces";
import { IProfession } from "../interfaces/IProfession";

interface SearchProfessionalsProps {
  onNavigate?: (page: string) => void;
}

function SearchProfessionals({onNavigate}: SearchProfessionalsProps){
  const [selectedState, setSelectedState] = useState<string>("");
  const [selectedCity, setSelectedCity] = useState<string>("");
  const [selectedProfessional, setSelectedProfessional] = useState<string>("");
  const [showResults ] = useState<boolean>(false); 
  // const [selectedProfessionalName, setSelectedProfessionalName] = useState<string>("");
  const [citiesByState, setcitiesByState] = useState<ICity[]>([]);
  const [professionals, setProfessionals] = useState<IProfissional[]>([]);
  const [professions, setProfessions] = useState<IProfession[]>([]);
  const [showModal, setShowModal] = useState(false);
  const [imageModal, setImageModal] = useState<string>("");
  const navigate = useNavigate();

  function gerarNumeroAleatorio(min: number, max: number): number {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1)) + min;
  }


  // function getScreenSize() {
  //   const width = window.innerWidth;

  //   if (width < 768) {
  //     return "xs";
  //   } else if (width < 992) {
  //     return "sm";
  //   } else if (width < 1200) {
  //     return "md";
  //   } else {
  //     return "lg"; // ou xl, dependendo da sua necessidade
  //   }
  // }

  //const screenSize = getScreenSize();
  //console.log("Tamanho da tela:", screenSize);

  useEffect(() => {
    setSelectedCity("");

    const fetchData = async () => {
      const citiesByState = await CityService.citiesByStatePublic({
        uf: selectedState,
      });

      const json_cities = await citiesByState.data;
      if (citiesByState.status == 200) {
        setcitiesByState(json_cities as ICity[]);
      }

      const professions = await ProfessionService.getProfessionsPublic();
      const json_professions = await professions.data;
      if (professions.status == 200) {
        // console.log("json_prof:", json_professions);
        setProfessions(json_professions as IProfession[]);
      }
      //setSelectedCity("3648");
    };

    fetchData();
  }, [selectedState]);

  useEffect(() => {
    const fetchData = async () => {
      // console.log("selectedCity:", selectedCity);
      if (!selectedCity) {
        setcitiesByState([] as ICity[]);
        setShowModal(false);
        setImageModal("");
        return;
      }
      try {
        const resu_region = await RegionService.getRegionbyCity(
          parseInt(selectedCity)
        );
        if (resu_region.status == 200) {
          const region_id = await resu_region.data.id;
          //setSelectedRegion(parseInt(resu_region.data.id));
          // console.log("Região selecionada:", region_id);

          const response_modal = await BannerService.getBannerByPagePublic({
            page: "B",
            cityId: 0,
            regionId: region_id,
          });
          if (response_modal.status == 200) {
            // console.log("response_modal:", response_modal);
            const json_home = await response_modal.data;
            if (json_home.length > 0) {
              const img_number = gerarNumeroAleatorio(0, json_home.length - 1);
                // console.log("img_number:", img_number);
                // console.log("Imagem selecionada:", json_home[img_number]);
              setImageModal(json_home[img_number] || "");
            }
            else {
              // console.log("Nenhuma imagem encontrada para a região selecionada.");
              setImageModal("");
              handleSearch();
            }
          } else {
            handleSearch();
          }
        } else {
          console.log("Região não encontrada para a cidade selecionada.");
          handleSearch();
        }
      } catch (error) {
        console.error("Erro ao obter região da cidade:", error);
        handleSearch();
      }
    };

    fetchData();
  }, [showModal]);

  const handleSearch = async () => {
    localStorage.setItem("search_city", selectedCity);
    const return_professionals =
      await ProfessionalService.getProfessionalByCityAndProfession({
        cityID: parseInt(selectedCity),
        professionID: parseInt(selectedProfessional),
        limit: 1000,
        offset: 0,
      });
    // console.log(return_professionals);

    const json_professionals = await return_professionals.data.profissionais;

    setProfessionals(json_professionals as IProfissional[]);
    //navigate(`/search-results?${selectedCity}&${selectedProfessional}`, {
    navigate("/search-results", {
      state: {
        selectedCity: selectedCity,
        selectedProfessional: selectedProfessional,
      },
    });
    // const professional = professionals.find(p => p.id === selectedProfessional);
    // setSelectedProfessionalName(professional?.name || '');
    // setShowResults(true);
  };

  if (showResults) {
    //return (
      // <SearchResults
      // profession={selectedProfessional}
      // professionals={professionals}
      // onNewSearch={() => setShowResults(false)}
      // onNavigate={onNavigate}
      // />
    // );
  };

  return (
    <div className="min-h-screen bg-gray-50">
      {showModal && imageModal !== "" && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-70">
          {/* O modal agora não tem largura máxima fixa. Ele se ajustará à largura da imagem. */}
          <div className="relative h-auto max-h-[95vh] rounded-lg overflow-hidden shadow-xl bg-black flex items-center justify-center">
            {/* Botão X */}
            <button
              onClick={() => {
                setShowModal(false);
                handleSearch();
              }}
              className="absolute top-4 right-4 z-10 text-white text-2xl bg-black bg-opacity-50 hover:bg-opacity-70 rounded-full w-10 h-10 flex items-center justify-center"
            >
              ✕
            </button>

            {/* Imagem */}
            <img
              src={atob(imageModal)} // corrigir esse erro de tipo de imagem imageModal é do tipo string 
              alt="Imagem de Construção"
              onClick={
                imageModal != ""
                  ? () => window.open(imageModal, "_blank")
                  : () => {}
              }
              // Garante que a imagem se ajuste à altura máxima do modal e preencha a largura
              className="max-h-full object-contain
                   max-w-[100vw] sm:max-w-[90vw] lg:max-w-[80vw] xl:max-w-[70vw]
                   w-auto sm:h-[95vh] xs:h-auto" //Adicionamos w-auto h-auto diretamente aqui"
              //style={{ width: "auto", height: "95vh" }} // Adicionado estilo inline para largura máxima da viewport
              style={{
                cursor: imageModal != "" ? "pointer" : "",
              }}
            />
          </div>
        </div>
      )}

      {/* Selection Form */}
      <div className="max-w-4xl mx-auto px-4 py-12">
        <div className="bg-white rounded-xl shadow-lg p-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-8">
            Encontre Profissionais da Construção Civil
          </h1>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            {/* Estado Select */}
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

            {/* Cidade Select */}
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
                  {selectedState &&
                    citiesByState.map((city: ICity) => (
                      <option key={city.id} value={city.id}>
                        {city.name}
                      </option>
                    ))}
                </select>
              </div>
            </div>

            {/* Profissional Select */}
            <div className="relative">
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Profissional
              </label>
              <div className="relative">
                <HardHat className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                <select
                  value={selectedProfessional}
                  onChange={(e) => setSelectedProfessional(e.target.value)}
                  className="block w-full pl-10 pr-4 py-2.5 text-gray-900 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none bg-white"
                >
                  <option value="">Selecione o profissional</option>
                  {professions.map((prof: IProfession) => (
                    <option key={prof.id} value={prof.id}>
                      {prof.name}
                    </option>
                  ))}
                </select>
              </div>
            </div>
          </div>

          <button
            onClick={ () => {
              setShowModal(true);
            }}
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
