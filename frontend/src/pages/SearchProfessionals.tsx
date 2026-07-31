import React, { useState, useEffect, useMemo, useRef } from "react";
import { MapPin, HardHat } from "lucide-react";
import { CityService } from "../services/CityService";
import { ProfessionalService } from "../services/ProfessionalService";
import { ProfessionService } from "../services/ProfessionService";
import { useNavigate } from "react-router-dom";
import AsyncSelect from "react-select/async";
import { ICitySearchProfessionals, IProfissional } from "../interfaces";
import { IProfessionSearchProfessionals } from "../interfaces/IProfession";
import { ArrowLeft } from "lucide-react";

interface SearchProfessionalsProps {
  onNavigate?: (page: string) => void;
}

interface CityOption {
  value: number;
  label: string;
}

// Minimo de caracteres e intervalo de espera antes de consultar o backend.
const CITY_SEARCH_MIN_LENGTH = 2;
const CITY_SEARCH_DEBOUNCE_MS = 400;

function SearchProfessionals({ onNavigate }: SearchProfessionalsProps) {
  const [selectedCity, setSelectedCity] = useState<string>("");
  const [selectedCityOption, setSelectedCityOption] =
    useState<CityOption | null>(null);
  const [selectedProfessional, setSelectedProfessional] = useState<string>("");
  const [professionals, setProfessionals] = useState<IProfissional[]>([]);
  const [professions, setProfessions] = useState<
    IProfessionSearchProfessionals[]
  >([]);
  const navigate = useNavigate();
  const cityDebounceRef = useRef<ReturnType<typeof setTimeout>>();
  const cityAbortRef = useRef<AbortController>();

  // Busca as profissões ao montar a página. As cidades não são mais carregadas
  // em lote: vêm do autocomplete conforme o usuário digita.
  useEffect(() => {
    const fetchProfessions = async () => {
      try {
        const professionsRes = await ProfessionService.getProfessionsPublic();
        if (professionsRes.status === 200) setProfessions(professionsRes.data);
      } catch (error) {
        console.error("Erro ao buscar profissões:", error);
      }
    };

    fetchProfessions();
  }, []);

  // Cancela requisições/timers pendentes do autocomplete ao desmontar
  useEffect(() => {
    return () => {
      if (cityDebounceRef.current) clearTimeout(cityDebounceRef.current);
      cityAbortRef.current?.abort();
    };
  }, []);

  // Função de busca de profissionais
  const handleSearch = async () => {
    localStorage.setItem("search_city", selectedCity);
    try {
      const return_professionals =
        await ProfessionalService.getProfessionalByCityAndProfession({
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

  // Autocomplete de cidades: aguarda o usuário parar de digitar, cancela a
  // requisição anterior e devolve as sugestões no formato "Cidade - UF".
  const loadCityOptions = useMemo(
    () =>
      (inputValue: string): Promise<CityOption[]> => {
        const term = inputValue.trim();

        if (cityDebounceRef.current) clearTimeout(cityDebounceRef.current);
        cityAbortRef.current?.abort();

        if (term.length < CITY_SEARCH_MIN_LENGTH) return Promise.resolve([]);

        return new Promise<CityOption[]>((resolve) => {
          cityDebounceRef.current = setTimeout(async () => {
            const controller = new AbortController();
            cityAbortRef.current = controller;

            try {
              const response = await CityService.searchCitiesPublic(
                term,
                8,
                controller.signal
              );
              resolve(
                response.data.map((city: ICitySearchProfessionals) => ({
                  value: city.id,
                  label: `${city.name} - ${city.uf}`,
                }))
              );
            } catch (error) {
              // Requisição cancelada por uma digitação mais recente não é erro
              if (!controller.signal.aborted) {
                console.error("Erro ao buscar cidades:", error);
              }
              resolve([]);
            }
          }, CITY_SEARCH_DEBOUNCE_MS);
        });
      },
    []
  );

  return (
    <div className="min-h-screen bg-gray-50">
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
            Encontrar Profissional
          </h1>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {/* Cidade */}
            <div className="relative">
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Cidade
              </label>
              <div className="relative">
                <MapPin className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5 z-10" />
                <AsyncSelect
                  loadOptions={loadCityOptions}
                  cacheOptions
                  placeholder="Digite sua cidade"
                  loadingMessage={() => "Buscando cidades..."}
                  noOptionsMessage={({ inputValue }) =>
                    inputValue.trim().length < CITY_SEARCH_MIN_LENGTH
                      ? "Digite ao menos 2 letras da cidade"
                      : "Cidade não encontrada"
                  }
                  value={selectedCityOption}
                  onChange={(option) => {
                    setSelectedCityOption(option);
                    setSelectedCity(option ? option.value.toString() : "");
                  }}
                  className="text-black"
                  styles={{
                    control: (base) => ({
                      ...base,
                      minHeight: "46px",
                      borderColor: "#d1d5db",
                      paddingLeft: "30px",
                      borderRadius: "0.5rem",
                    }),
                  }}
                />
              </div>
            </div>

            {/* Profissional */}
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
            onClick={handleSearch}
            disabled={!selectedCity || !selectedProfessional}
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
