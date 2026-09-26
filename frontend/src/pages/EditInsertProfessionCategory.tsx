import React, { useEffect, useState } from "react";
import { Save, X } from "lucide-react";
import Swal from "sweetalert2";

import { ProfessionCategoryService } from "../services/ProfessionCategoryService";
import { ProfessionService } from "../services/ProfessionService";
import { IProfessionOfCategory } from "../interfaces/IProfessionCategory";

interface EditInsertProfessionCategoryProps {
  id: number;
  onClose: (shouldReload: boolean) => void;
}

const emptyForm = {
  name: "",
  description: "",
  icon: "",
  position: 0,
  active: true,
  client_travels: false,
};

/** "" representa herdar o flag da categoria. */
type OverrideValue = "" | "true" | "false";

const overrideOf = (profession: IProfessionOfCategory): OverrideValue =>
  profession.client_travels_override === null ||
  profession.client_travels_override === undefined
    ? ""
    : String(profession.client_travels_override) as OverrideValue;

function EditInsertProfessionCategory({
  id,
  onClose,
}: EditInsertProfessionCategoryProps) {
  const [formData, setFormData] = useState(emptyForm);
  const [professions, setProfessions] = useState<IProfessionOfCategory[]>([]);
  const [overrides, setOverrides] = useState<Record<number, OverrideValue>>({});
  const [isSaving, setIsSaving] = useState(false);

  useEffect(() => {
    if (id === 0) return;

    const fetchData = async () => {
      try {
        const response =
          await ProfessionCategoryService.getProfessionsByCategory(id);
        const category = response.data;
        setFormData({
          name: category.name ?? "",
          description: category.description ?? "",
          icon: category.icon ?? "",
          position: category.position ?? 0,
          active: category.active ?? true,
          client_travels: category.client_travels ?? false,
        });

        const list: IProfessionOfCategory[] = category.professions || [];
        setProfessions(list);
        setOverrides(
          Object.fromEntries(list.map((p) => [p.id, overrideOf(p)]))
        );
      } catch (error) {
        console.error("Erro ao carregar a categoria:", error);
      }
    };

    fetchData();
  }, [id]);

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value, type, checked } = e.target as HTMLInputElement;
    setFormData((prev) => ({
      ...prev,
      [name]:
        type === "checkbox"
          ? checked
          : type === "number"
          ? Number(value)
          : value,
    }));
  };

  // Só as profissões cujo override mudou são regravadas: evita reescrever a
  // tabela inteira a cada salvamento da categoria.
  const changedProfessions = () =>
    professions.filter((p) => overrides[p.id] !== overrideOf(p));

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!formData.name.trim()) {
      Swal.fire("Atenção", "Nome é obrigatório", "warning");
      return;
    }

    setIsSaving(true);
    try {
      await ProfessionCategoryService.postCategory({ id, ...formData });

      await Promise.all(
        changedProfessions().map((profession) => {
          const override = overrides[profession.id];
          return ProfessionService.postProfession({
            id: profession.id,
            name: profession.name,
            description: profession.description,
            icon: profession.icon,
            category_id: id,
            client_travels: override === "" ? null : override === "true",
          });
        })
      );

      Swal.fire({
        position: "center",
        icon: "success",
        title: id === 0 ? "Categoria inserida!" : "Categoria atualizada!",
        showConfirmButton: false,
        timer: 1500,
      });
      onClose(true);
    } catch (error: any) {
      Swal.fire(
        "Erro",
        error?.response?.data?.error || "Falha ao salvar a categoria",
        "error"
      );
    } finally {
      setIsSaving(false);
    }
  };

  return (
    <div className="bg-white rounded-lg shadow-md p-8 w-full max-w-2xl max-h-[85vh] overflow-y-auto">
      <h2 className="text-2xl font-bold text-gray-900 mb-6">
        Dados da Categoria
      </h2>

      <form onSubmit={handleSubmit} className="space-y-6">
        <div>
          <label htmlFor="name" className="block text-sm font-medium text-gray-700">
            Nome
          </label>
          <input
            type="text"
            name="name"
            id="name"
            value={formData.name}
            onChange={handleChange}
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

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
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div>
          <label htmlFor="icon" className="block text-sm font-medium text-gray-700">
            Ícone (URL)
          </label>
          <input
            type="text"
            name="icon"
            id="icon"
            value={formData.icon}
            onChange={handleChange}
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div className="grid grid-cols-2 gap-4">
          <div>
            <label
              htmlFor="position"
              className="block text-sm font-medium text-gray-700"
            >
              Ordem de exibição
            </label>
            <input
              type="number"
              name="position"
              id="position"
              value={formData.position}
              onChange={handleChange}
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
            />
          </div>

          <div className="flex items-end">
            <label className="flex items-center gap-2 text-sm text-gray-700">
              <input
                type="checkbox"
                name="active"
                checked={formData.active}
                onChange={handleChange}
                className="w-4 h-4"
              />
              Categoria ativa
            </label>
          </div>
        </div>

        <fieldset className="border border-gray-200 rounded-md p-4">
          <legend className="text-sm font-medium text-gray-700 px-1">
            Quem se desloca nesta categoria
          </legend>
          <label className="flex items-center gap-2 text-sm text-gray-700">
            <input
              type="radio"
              name="client_travels"
              checked={!formData.client_travels}
              onChange={() =>
                setFormData((prev) => ({ ...prev, client_travels: false }))
              }
            />
            O profissional vai até o cliente
          </label>
          <label className="flex items-center gap-2 text-sm text-gray-700 mt-2">
            <input
              type="radio"
              name="client_travels"
              checked={formData.client_travels}
              onChange={() =>
                setFormData((prev) => ({ ...prev, client_travels: true }))
              }
            />
            O cliente vai até o profissional
          </label>
        </fieldset>

        {professions.length > 0 && (
          <fieldset className="border border-gray-200 rounded-md p-4">
            <legend className="text-sm font-medium text-gray-700 px-1">
              Exceções por profissão
            </legend>
            <div className="space-y-3">
              {professions.map((profession) => (
                <div
                  key={profession.id}
                  className="flex items-center justify-between gap-4"
                >
                  <span className="text-sm text-gray-900">
                    {profession.name}
                  </span>
                  <select
                    value={overrides[profession.id] ?? ""}
                    onChange={(e) =>
                      setOverrides((prev) => ({
                        ...prev,
                        [profession.id]: e.target.value as OverrideValue,
                      }))
                    }
                    className="px-3 py-1.5 text-sm border border-gray-300 rounded-md bg-white"
                  >
                    <option value="">Herda da categoria</option>
                    <option value="false">Profissional vai até o cliente</option>
                    <option value="true">Cliente vai até o profissional</option>
                  </select>
                </div>
              ))}
            </div>
          </fieldset>
        )}

        <div className="flex gap-3">
          <button
            type="submit"
            disabled={isSaving}
            className="flex-1 flex items-center justify-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:bg-gray-300"
          >
            <Save className="w-5 h-5" />
            Salvar
          </button>
          <button
            type="button"
            onClick={() => onClose(false)}
            className="flex-1 flex items-center justify-center gap-2 px-4 py-2 bg-gray-200 text-gray-800 rounded-lg hover:bg-gray-300 transition-colors"
          >
            <X className="w-5 h-5" />
            Cancelar
          </button>
        </div>
      </form>
    </div>
  );
}

export default EditInsertProfessionCategory;
