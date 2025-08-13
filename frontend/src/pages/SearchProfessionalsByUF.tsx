import { useState, useEffect } from "react";
import { Building2, Search } from "lucide-react";
import { states } from "../data";
import { ProfessionalService } from "../services/ProfessionalService";
import LoadingText from "../components/LoadingText";


function SearchProfessionalsByUF() {
  const [selectedState, setSelectedState] = useState("");
  const [selectedCity, setSelectedCity] = useState("");
  const [citiesByState, setcitiesByState] = useState([{}]);
  const [isLoading, setIsLoading] = useState(true);
  const [totalCities, setTotalCities] = useState(0);
  const [showModal, setShowModal] = useState(false);
  const [professionsByCity, setProfessionsByCity] = useState([{}]);

  useEffect(() => {
    setSelectedCity("");

    const fetchData = async () => {
      if (showModal && selectedCity.length!==0) {
        setIsLoading(true);
        const cities_professions_return =
          await ProfessionalService.CountProfessionalsByProfessionByCitie({
            cityID: selectedCity
          });
        console.log(cities_professions_return);
        if (cities_professions_return.status == 200) {
          const json = await cities_professions_return.data;
          setProfessionsByCity(json);
        }
        setIsLoading(false);
      }
      else {
        setIsLoading(true);
        const cities_return = await ProfessionalService.ProfessionalByState({ state:selectedState, limit: 1000000, offset: 0 });
        console.log(cities_return);
        if (cities_return.status == 200) {
          const json = await cities_return.data.result;
          setcitiesByState(json);
          setTotalCities(await cities_return.data.total);
        }
        setIsLoading(false);
      }
    };

    fetchData();

  }, [selectedState,showModal]);


  return (
    <div className="min-h-screen bg-gray-50">
      {/* Loading */}
      {isLoading && <LoadingText />}
      {/* Modal */}
      {!isLoading && showModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 overflow-y-auto">
          <div className="flex flex-col items-center px-4 sm:px-0">
            <div className="bg-white rounded-2xl shadow-2xl p-8 w-full max-w-4xl max-h-[85vh]">
              <h2 className="text-2xl font-semibold text-gray-800 mb-6 text-center">
                Profissões por Cidade
              </h2>
              <div className="overflow-y-auto w-full max-w-4xl max-h-[55vh]">
                <table className="w-full table-auto border-collapse">
                  <tbody>
                    {professionsByCity.map((profession) => (
                      <tr
                        key={profession.oid}
                        className="hover:bg-gray-50 transition-colors border-b border-gray-100"
                      >
                        <td className="px-6 py-4 text-sm text-gray-800">
                          {profession.professionName}
                        </td>
                        <td className="px-6 py-4 text-sm text-gray-600">
                          <div className="flex items-center gap-2">
                            {profession.quantity}
                          </div>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
              <div className="mt-6 flex justify-center">
                <button
                  onClick={() => {
                    setSelectedCity("0");
                    setShowModal(false);
                  }}
                  className="inline-flex items-center gap-2 px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-md hover:bg-blue-700 transition"
                >
                  Fechar
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
      {/* Selection Form */}
      <div className="max-w-4xl mx-auto px-4 py-0">
        <div className="bg-white rounded-2xl shadow-2xl p-8">
          <h1 className="text-3xl font-extrabold text-gray-900 mb-10 text-center">
            Profissionais por Estado
          </h1>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
            {/* Estado Select */}
            <div className="relative">
              <label className="block text-sm font-semibold text-gray-700 mb-2">
                Estado
              </label>
              <div className="relative">
                <Building2 className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                <select
                  value={selectedState}
                  onChange={(e) => setSelectedState(e.target.value)}
                  className="block w-full pl-10 pr-4 py-2.5 text-sm text-gray-900 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white shadow-sm"
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

            {/* Total Profissionais */}
            <div className="bg-gray rounded-2xl shadow-2xl p-8 flex items-end justify-center md:justify-start ">
              <div className="text-lg font-semibold text-gray-800">
                Total de Profissionais:{" "}
                <span className="text-blue-600 font-bold">{totalCities}</span>
              </div>
            </div>
          </div>
        </div>

        {/* City Selection */}
        {(!isLoading || showModal) && citiesByState.length > 0 && (
          <div className="mt-6 overflow-x-auto bg-white shadow-md rounded-xl">
            <table className="w-full table-auto text-sm">
              <thead className="bg-gray-50">
                <tr className="text-left border-b border-gray-200">
                  <th className="px-6 py-4 font-medium text-gray-600">
                    Cidade
                  </th>
                  <th className="px-6 py-4 font-medium text-gray-600">
                    Quantidade
                  </th>
                  <th className="px-6 py-4 font-medium text-gray-600">Ações</th>
                </tr>
              </thead>
              <tbody>
                {citiesByState.map((city) => (
                  <tr
                    key={city.oid}
                    className="border-b border-gray-100 hover:bg-gray-50 transition-colors"
                  >
                    <td className="px-6 py-4 text-gray-900">{city.cidade}</td>
                    <td className="px-6 py-4 text-gray-700">
                      {city.quantidadeProfissionais}
                    </td>
                    <td className="px-6 py-4">
                      <button
                        onClick={() => {
                          setSelectedCity(city.cidadeId);
                          setShowModal(true);
                        }}
                        className="inline-flex items-center justify-center p-2 text-blue-600 hover:bg-blue-50 rounded transition"
                      >
                        <Search className="w-4 h-4" />
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
}

export default SearchProfessionalsByUF;
