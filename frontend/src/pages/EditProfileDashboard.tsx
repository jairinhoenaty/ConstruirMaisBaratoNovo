import {
  Building2,
  Check,
  ChevronDown,
  HardHat,
  Mail,
  MapPin,
  Phone,
  Save,
  User,
  X,
} from "lucide-react";
import InputMask from "react-input-mask";
import { states } from "../data"; // Assuming states is imported from here
import { useState } from "react";
import Swal from "sweetalert2";
import {
  CityService,
  ProfessionalService,
  ProfessionService,
} from "../services";
import { ClientService } from "../services/ClientService";
import { StoreService } from "../services/StoreService";
import React from "react";

interface EditProps {
  id: number;
  profile: string;
  onClose: any;
}

function EditProfileDashboard({ id, profile, onClose }: EditProps) {
  const [showProfessions, setShowProfessions] = useState(false);
  const [citiesByState, setcitiesByState] = useState([{ id: 0, name: "" }]);
  const [professions, setProfessions] = useState([
    { id: "", name: "", description: "" },
  ]);
  const [isProfessional, setIsProfessional] = useState(false);
  const [isClient, setIsClient] = useState(false);

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
    verified: false, // Added the verified field
  });

  React.useEffect(() => {
    const fetchData = async () => {
      console.log(profile);
      if (profile === "client") {
        const response = await ClientService.getClientbyID(id);
        const json = await response.data;
        if (response.status === 200) {
          localStorage.setItem("post_id", response.data.oid);
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
            verified: json.verified || false, // Assuming 'verified' field exists in client data
          }));
          setIsClient(true);
          const citiesResponse = await CityService.citiesByState({
            uf: json.cidade.uf,
          });

          const json_cities = await citiesResponse.data;
          if (citiesResponse.status === 200) {
            setcitiesByState(json_cities);
          }
          setFormData((prev) => ({
            ...prev,
            city: json.cidade.oid,
          }));
        }
      } else if (profile === "profissional" || profile === "admin") {
        const professionsResponse = await ProfessionService.getProfessions(
          10000,
          0
        );
        console.log(professionsResponse);
        const json_professions = await professionsResponse.data.profissoes;
        if (professionsResponse.status === 200) {
          setProfessions(json_professions);
        }

        const response = await ProfessionalService.getProfessionalbyID(id);
        const json = await response.data;
        if (response.status === 200) {
          localStorage.setItem("post_id", response.data.oid);
          console.log(json);
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
            verified: json.verified || false, // Assuming 'verified' field exists in professional data
          }));
          setIsProfessional(true);
          const citiesResponse = await CityService.citiesByState({
            uf: json.cidade.uf,
          });

          const json_cities = await citiesResponse.data;
          if (citiesResponse.status === 200) {
            setcitiesByState(json_cities);
          }
          setFormData((prev) => ({
            ...prev,
            city: json.cidade.oid,
          }));
        }
      } else if (profile === "store") {
        const response = await StoreService.getStorebyID(id);
        const json = await response.data;
        if (response.status === 200) {
          localStorage.setItem("post_id", response.data.oid);
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
            verified: json.verified || false, // Assuming 'verified' field exists in store data
          }));
          const citiesResponse = await CityService.citiesByState({
            uf: json.cidade.uf,
          });

          const json_cities = await citiesResponse.data;
          if (citiesResponse.status === 200) {
            setcitiesByState(json_cities);
          }
          setFormData((prev) => ({
            ...prev,
            city: json.cidade.oid,
          }));
        }
      }
    };
    fetchData();
  }, [id, profile]); // Added id and profile to dependency array

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

  const handleChange = async (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
    const { name, value, type, checked } = e.target;

    if (name === "state") {
      const citiesResponse = await CityService.citiesByState({
        uf: value,
      });

      const json_cities = await citiesResponse.data;
      if (citiesResponse.status === 200) {
        setcitiesByState(json_cities);
      }
    }

    setFormData((prev) => ({
      ...prev,
      [name]: type === "checkbox" ? checked : value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const currentProfile = localStorage.getItem("profile"); // Use a different variable name to avoid conflict with prop
    const post_id = localStorage.getItem("post_id") ?? "";
    console.log(post_id);
    let postReturn: any;
    console.log(currentProfile);

    const commonData = {
      oid: parseInt(post_id),
      Name: formData.fullName,
      Email: formData.email,
      Telephone: formData.whatsapp,
      cep: formData.cep,
      street: formData.street,
      neighborhood: formData.neighborhood,
      cityId: parseInt(formData.city),
      LgpdAceito: "S",
      Password: null,
      image: null,
      verified: formData.verified, // Include the verified status
    };

    if (currentProfile === "client") {
      postReturn = await ClientService.postClient(commonData);
    } else if (
      currentProfile === "profissional" ||
      currentProfile === "admin"
    ) {
      postReturn = await ProfessionalService.postProfessional({
        ...commonData,
        professionIds: formData.professions,
      });
    } else if (currentProfile === "store") {
      postReturn = await StoreService.postStore(commonData);
    }

    if (postReturn && postReturn.status === 200) {
      Swal.fire({
        position: "center",
        icon: "success",
        title: "Conta Atualizada!!!",
        showConfirmButton: false,
        timer: 1500,
      });
      onClose();
    } else {
      Swal.fire({
        position: "center",
        icon: "error",
        title: "Erro ao Atualizar Conta",
        text: "Por favor, tente novamente.",
        showConfirmButton: true,
      });
    }
    console.log("Form submitted:", formData);
  };

  return (
    <div className="bg-white rounded-lg shadow-md p-8">
      <h2 className="text-2xl font-bold text-gray-900 mb-6">Dados da Conta</h2>
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
                onChange={handleChange} // Use the unified handleChange
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
                onChange={handleChange} // Use the unified handleChange
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
              <span className="block truncate">{selectedProfessionsText}</span>
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

        {/* Verified Checkbox */}
        {(profile === "admin" || profile === "profissional") && ( // Only show for admin or professional profiles
          <div className="flex items-center mt-4">
            <input
              id="verified"
              name="verified"
              type="checkbox"
              checked={formData.verified}
              onChange={handleChange}
              className="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label
              htmlFor="verified"
              className="ml-2 block text-sm text-gray-900"
            >
              Verificado
            </label>
          </div>
        )}

        <div>
          <button
            type="submit"
            className="w-full flex items-center justify-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
          >
            <Save className="w-5 h-5" />
            Salvar
          </button>
        </div>
        <div>
          <button
            onClick={() => {
              onClose(false);
            }}
            className="w-full flex items-center justify-center gap-2 px-4 py-2 bg-gray-500 text-white rounded-lg hover:bg-gray-600 transition-colors"
          >
            <X className="w-5 h-5" />
            Cancelar
          </button>
        </div>
      </form>
    </div>
  );
}
export default EditProfileDashboard;
