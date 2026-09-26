import { Save, User, X } from "lucide-react";
import React, { useEffect, useState } from "react";
import Swal from "sweetalert2";

import { ProfessionService } from "../services";
import { ProfessionCategoryService } from "../services/ProfessionCategoryService";
import { IProfessionCategory } from "../interfaces/IProfessionCategory";

interface EditProps {
  id: number;
  onClose: any;
}

/** "" representa herdar o flag da categoria. */
type OverrideValue = "" | "true" | "false";

function EditInsertProfessions({ id, onClose }: EditProps) {
  const [formData, setFormData] = useState({
    name: "",
    description: "",
    icon: "",
    categoryId: "",
    clientTravels: "" as OverrideValue,
  });
  const [categories, setCategories] = useState<IProfessionCategory[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const categoriesRes = await ProfessionCategoryService.getCategories(
          1000,
          0
        );
        setCategories(categoriesRes.data?.categorias || []);
      } catch (error) {
        console.error("Erro ao buscar categorias:", error);
      }

      if (id === 0) return;

      const response = await ProfessionService.getProfessionbyID(id);
      if (response.status !== 200) return;

      const json = response.data;
      setFormData({
        name: json.name ?? "",
        description: json.description ?? "",
        icon: json.icon ?? "",
        categoryId: json.category_id ? String(json.category_id) : "",
        clientTravels:
          json.client_travels_override === null ||
          json.client_travels_override === undefined
            ? ""
            : (String(json.client_travels_override) as OverrideValue),
      });
    };

    fetchData();
  }, [id]);

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const selectedCategory = categories.find(
    (category) => String(category.id) === formData.categoryId
  );

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!formData.categoryId) {
      Swal.fire("Atenção", "Escolha a categoria da profissão", "warning");
      return;
    }

    const postReturn = await ProfessionService.postProfession({
      id: id,
      name: formData.name,
      description: formData.description,
      icon: formData.icon,
      category_id: Number(formData.categoryId),
      client_travels:
        formData.clientTravels === ""
          ? null
          : formData.clientTravels === "true",
    });

    if (postReturn.status == 200) {
      Swal.fire({
        position: "center",
        icon: "success",
        title: id == 0 ? "Profissão inserida!!!" : "Profissão atualizada!!!",
        showConfirmButton: false,
        timer: 1500,
      });
      onClose();
    }
  };

  return (
    <div className="bg-white rounded-lg shadow-md p-8">
      <h2 className="text-2xl font-bold text-gray-900 mb-6">
        Dados da Profissão
      </h2>
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
          <label
            htmlFor="categoryId"
            className="block text-sm font-medium text-gray-700"
          >
            Categoria
          </label>
          <select
            name="categoryId"
            id="categoryId"
            value={formData.categoryId}
            onChange={handleChange}
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md bg-white focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">Selecione a categoria</option>
            {categories.map((category) => (
              <option key={category.id} value={category.id}>
                {category.name}
              </option>
            ))}
          </select>
        </div>

        <div>
          <label
            htmlFor="clientTravels"
            className="block text-sm font-medium text-gray-700"
          >
            Quem se desloca
          </label>
          <select
            name="clientTravels"
            id="clientTravels"
            value={formData.clientTravels}
            onChange={handleChange}
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md bg-white focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">
              Herda da categoria
              {selectedCategory
                ? selectedCategory.client_travels
                  ? " (cliente vai até o profissional)"
                  : " (profissional vai até o cliente)"
                : ""}
            </option>
            <option value="false">O profissional vai até o cliente</option>
            <option value="true">O cliente vai até o profissional</option>
          </select>
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
            type="button"
            onClick={() => onClose(false)}
            className="w-full flex items-center justify-center gap-2 px-4 py-2 bg-gray-200 text-gray-800 rounded-lg hover:bg-gray-300 transition-colors"
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
