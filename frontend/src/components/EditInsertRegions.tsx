import Select from "react-select";
import { useEffect, useState } from "react";
import {
  RegionService,
  CityService,
  //StateService, // supondo que exista
} from "../services";
import { Save, X } from "lucide-react";
import Swal from "sweetalert2";
import { states } from "../data";

interface EditProps {
  id: number;
  onClose: any;
}

function EditInsertRegions({ id, onClose }: EditProps) {
  const [formData, setFormData] = useState({
    name: "",
    description: "",
    icon: "",
    cityIds: [] as number[],
  });

  const [UFs, setUFs] = useState<{ label: string; value: string }[]>([]);
  const [selectedUf, setSelectedUf] = useState<string>("");
  const [cities, setCities] = useState<{ label: string; value: number }[]>([]);

  useEffect(() => {
    // Carrega estados
    const loadStates = async () => {
      //const res = await StateService.getAll();
      
      const data = states.map((uf: any) => ({
        label: uf.name,
        value: uf.id,
      }));
      setUFs(data);
    };
    loadStates();
  }, []);

  useEffect(() => {
    // Carrega dados da região se for edição
    const fetchData = async () => {
      if (id !== 0) {
        const response = await RegionService.getRegionbyID(id);
        if (response.status === 200) {
          const json = response.data;
          setSelectedUf(json.cities[0].uf);

          setFormData({
            name: json.name,
            description: json.description,
            icon: json.icon,
            cityIds: json.cities.map((x: any) => x.oid), //json.cities || [],
          });
        }
      }
    };
    fetchData();
  }, [id]);

  // Carrega cidades quando UF muda
  useEffect(() => {
    const loadCities = async () => {
      if (!selectedUf) return;
      const res = await CityService.citiesByState({uf:selectedUf});
      const data = res.data.map((city: any) => ({
        label: city.name,
        value: city.id,
      }));
      setCities(data);
    };
    loadCities();
  }, [selectedUf]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const postReturn = await RegionService.postRegion({
      id,
      name: formData.name,
      description: formData.description,
      icon: formData.icon,
      cityIds: formData.cityIds,
      uf: selectedUf,
    });

    if (postReturn.status === 200) {
      Swal.fire({
        position: "center",
        icon: "success",
        title: id === 0 ? "Região inserida!" : "Região atualizada!",
        showConfirmButton: false,
        timer: 1500,
      });
      onClose();
    }
  };

  return (
    <div className="bg-white rounded-lg shadow-md p-8">
      <h2 className="text-2xl font-bold text-gray-900 mb-6">Dados da Região</h2>
      <form onSubmit={handleSubmit} className="space-y-6">
        {/* Nome */}
        <div>
          <label
            htmlFor="name"
            className="block text-sm font-medium text-gray-700"
          >
            Nome
          </label>
          <input
            type="text"
            name="name"
            id="name"
            value={formData.name}
            onChange={handleChange}
            className="block w-full px-3 py-2 border border-gray-300 rounded-md"
          />
        </div>

        {/* Descrição */}
        <div>
          <label
            htmlFor="description"
            className="block text-sm font-medium text-gray-700"
          >
            Descrição
          </label>
          <input
            type="text"
            name="description"
            id="description"
            value={formData.description}
            onChange={handleChange}
            className="block w-full px-3 py-2 border border-gray-300 rounded-md"
          />
        </div>

        {/* Estado (UF) */}
        <div>
          <label className="block text-sm font-medium text-gray-700">
            Estado (UF)
          </label>
          <Select
            options={UFs}
            isDisabled={selectedUf != "" && formData.cityIds.length!=0}
            value={UFs.find((uf) => uf.value === selectedUf)} // valor inicial aqui
            onChange={(option) => setSelectedUf(option?.value || "")}
            placeholder="Selecione o estado"
          />
        </div>

        {/* Cidades */}
        <div>
          <label className="block text-sm font-medium text-gray-700">
            Cidades
          </label>
          <Select
            isMulti
            required
            options={cities}
            value={cities.filter((city) =>
              formData.cityIds.includes(city.value)
            )}
            onChange={(selectedOptions) =>
              setFormData((prev) => ({
                ...prev,
                cityIds: selectedOptions.map((opt) => opt.value),
              }))
            }
            placeholder="Selecione as cidades"
          />
        </div>

        {/* Botões */}
        <div className="flex gap-4">
          <button
            type="submit"
            className="flex-1 flex items-center justify-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
          >
            <Save className="w-5 h-5" />
            Salvar
          </button>
          <button
            type="button"
            onClick={() => onClose(false)}
            className="flex-1 flex items-center justify-center gap-2 px-4 py-2 bg-gray-400 text-white rounded-lg hover:bg-gray-500"
          >
            <X className="w-5 h-5" />
            Cancelar
          </button>
        </div>
      </form>
    </div>
  );
}
export default EditInsertRegions;
