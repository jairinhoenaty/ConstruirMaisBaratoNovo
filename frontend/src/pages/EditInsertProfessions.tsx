import {
  Save,
  User,
  X,
} from "lucide-react";
import { useState } from "react";
import Swal from "sweetalert2";
import {
  ProfessionService,
} from "../services";
import React from "react";

interface EditProps {
  id: number;
  onClose: any;
}

function EditInsertProfessions({ id, onClose }: EditProps) {
  const [formData, setFormData] = useState({
    name: "",
    description: "",
    icon: "",
  });

  React.useEffect(() => {
    const fetchData = async () => {
      if (id!==0) {
        const response = await ProfessionService.getProfessionbyID(id);
        const json = response.data;
        if (response.status == 200) {
          setFormData((prev) => ({
            ...prev,
            name: json.name,
            description: json.description,
            icon: json.icon,
          }));
        }
      }
    };
    fetchData();
  }, []);

  
  const handleChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
      const postReturn = await ProfessionService.postProfession({
        id: id,
        name: formData.name,
        description: formData.description,
        icon: formData.icon,
      });

    if (postReturn.status == 200) {
      Swal.fire({
        position: "center",
        icon: "success",
        title: id==0?"Profissão inserida!!!":"Profissão atualizada!!!",
        showConfirmButton: false,
        timer: 1500,
      });
      onClose();
    }
    console.log("Form submitted:", formData);
  };

  return (
    <div className="bg-white rounded-lg shadow-md p-8">
      <h2 className="text-2xl font-bold text-gray-900 mb-6">Dados da Profissão</h2>
      <form onSubmit={handleSubmit} className="space-y-6">
        <div>
          <label
            htmlFor="name"
            className="block text-sm font-medium text-gray-700"
          >
            Nome
          </label>
          <div className="mt-1 relative rounded-md shadow-sm">
            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <User className="h-5 w-5 text-gray-400" />
            </div>
            <input
              type="text"
              name="name"
              id="name"
              value={formData.name}
              onChange={handleChange}
              className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        </div>
        <div>
          <label
            htmlFor="description"
            className="block text-sm font-medium text-gray-700"
          >
            Descrição
          </label>
          <div className="mt-1 relative rounded-md shadow-sm">
            <input
              type="text"
              name="description"
              id="description"
              value={formData.description}
              onChange={handleChange}
              className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        </div>
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
            className="w-full flex items-center justify-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
          >
            <X className="w-5 h-5" />
            Cancelar
          </button>
        </div>
      </form>
    </div>
  );
}
export default EditInsertProfessions;
