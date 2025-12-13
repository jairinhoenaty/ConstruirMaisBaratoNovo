import React, { useState, useEffect } from "react";
import { ProductCategoryService } from "../services/ProductCategoryService";
import { ICategoryProduct } from "../interfaces/ICategoryProduct";
import Swal from "sweetalert2";
import { Package, Edit2, Trash2, Plus, ChevronRight, ChevronDown } from "lucide-react";

function ProductCategoriesAdmin() {
  const [categories, setCategories] = useState<ICategoryProduct[]>([]);
  const [loading, setLoading] = useState(true);
  const [expandedCategories, setExpandedCategories] = useState<Set<number>>(new Set());
  const [showModal, setShowModal] = useState(false);
  const [modalMode, setModalMode] = useState<"create" | "edit">("create");
  const [editingCategory, setEditingCategory] = useState<ICategoryProduct | null>(null);
  const [formData, setFormData] = useState({ name: "", parent_id: null as number | null });

  useEffect(() => {
    fetchCategories();
  }, []);

  const fetchCategories = async () => {
    try {
      setLoading(true);
      const response = await ProductCategoryService.findTopLevelCategory();
      setCategories(response.data || []);
    } catch (error) {
      console.error("Erro ao buscar categorias:", error);
      Swal.fire("Erro", "Falha ao carregar categorias", "error");
    } finally {
      setLoading(false);
    }
  };

  const toggleCategory = (id: number) => {
    const newExpanded = new Set(expandedCategories);
    if (newExpanded.has(id)) {
      newExpanded.delete(id);
    } else {
      newExpanded.add(id);
    }
    setExpandedCategories(newExpanded);
  };

  const handleCreate = () => {
    setModalMode("create");
    setFormData({ name: "", parent_id: null });
    setEditingCategory(null);
    setShowModal(true);
  };

  const handleCreateSubcategory = (parentId: number) => {
    setModalMode("create");
    setFormData({ name: "", parent_id: parentId });
    setEditingCategory(null);
    setShowModal(true);
  };

  const handleEdit = (category: ICategoryProduct) => {
    setModalMode("edit");
    setFormData({ name: category.name, parent_id: category.parent_id || null });
    setEditingCategory(category);
    setShowModal(true);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!formData.name.trim()) {
      Swal.fire("Atenção", "Nome é obrigatório", "warning");
      return;
    }

    try {
      if (modalMode === "create") {
        await ProductCategoryService.create(formData);
        Swal.fire("Sucesso", "Categoria criada com sucesso!", "success");
      } else if (editingCategory) {
        await ProductCategoryService.update(editingCategory.id, { name: formData.name });
        Swal.fire("Sucesso", "Categoria atualizada com sucesso!", "success");
      }
      setShowModal(false);
      fetchCategories();
    } catch (error: any) {
      Swal.fire("Erro", error?.response?.data?.error || "Falha ao salvar categoria", "error");
    }
  };

  const handleDelete = async (category: ICategoryProduct) => {
    // Verificar se tem subcategorias
    if (category.children && category.children.length > 0) {
      Swal.fire({
        icon: "warning",
        title: "Não é possível excluir",
        text: "Esta categoria possui subcategorias. Exclua as subcategorias primeiro.",
      });
      return;
    }

    const result = await Swal.fire({
      title: "Tem certeza?",
      text: `Deseja excluir a categoria "${category.name}"?`,
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#d33",
      cancelButtonColor: "#3085d6",
      confirmButtonText: "Sim, excluir!",
      cancelButtonText: "Cancelar",
    });

    if (result.isConfirmed) {
      try {
        await ProductCategoryService.deleteCategory(category.id);
        Swal.fire("Excluído!", "Categoria excluída com sucesso.", "success");
        fetchCategories();
      } catch (error: any) {
        Swal.fire("Erro", error?.response?.data?.error || "Falha ao excluir categoria", "error");
      }
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-xl text-gray-600">Carregando...</div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50 p-6">
      <div className="max-w-5xl mx-auto">
        <div className="bg-white rounded-lg shadow-lg p-6">
          {/* Header */}
          <div className="flex justify-between items-center mb-6">
            <h1 className="text-2xl font-bold text-gray-900 flex items-center gap-2">
              <Package className="w-7 h-7 text-blue-600" />
              Categorias de Produtos
            </h1>
            <button
              onClick={handleCreate}
              className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 flex items-center gap-2 transition-colors"
            >
              <Plus className="w-5 h-5" />
              Nova Categoria
            </button>
          </div>

          {/* Categories List */}
          {categories.length === 0 ? (
            <div className="text-center py-12 text-gray-500">
              <Package className="w-16 h-16 mx-auto mb-4 text-gray-300" />
              <p>Nenhuma categoria cadastrada</p>
            </div>
          ) : (
            <div className="space-y-2">
              {categories.map((category) => (
                <div key={category.id} className="border border-gray-200 rounded-lg">
                  {/* Parent Category */}
                  <div className="flex items-center justify-between p-4 hover:bg-gray-50">
                    <div className="flex items-center gap-2 flex-1">
                      {category.children && category.children.length > 0 ? (
                        <button
                          onClick={() => toggleCategory(category.id)}
                          className="text-gray-600 hover:text-gray-900"
                        >
                          {expandedCategories.has(category.id) ? (
                            <ChevronDown className="w-5 h-5" />
                          ) : (
                            <ChevronRight className="w-5 h-5" />
                          )}
                        </button>
                      ) : (
                        <div className="w-5" />
                      )}
                      <Package className="w-5 h-5 text-blue-600" />
                      <span className="font-semibold text-gray-900">{category.name}</span>
                      {category.children && category.children.length > 0 && (
                        <span className="text-sm text-gray-500 bg-gray-100 px-2 py-0.5 rounded">
                          {category.children.length} subcategorias
                        </span>
                      )}
                    </div>
                    <div className="flex items-center gap-2">
                      <button
                        onClick={() => handleCreateSubcategory(category.id)}
                        className="text-green-600 hover:text-green-700 px-2 py-1 text-sm"
                        title="Adicionar subcategoria"
                      >
                        + Sub
                      </button>
                      <button
                        onClick={() => handleEdit(category)}
                        className="text-blue-600 hover:text-blue-700 p-2"
                        title="Editar"
                      >
                        <Edit2 className="w-4 h-4" />
                      </button>
                      <button
                        onClick={() => handleDelete(category)}
                        className="text-red-600 hover:text-red-700 p-2"
                        title="Excluir"
                      >
                        <Trash2 className="w-4 h-4" />
                      </button>
                    </div>
                  </div>

                  {/* Subcategories */}
                  {expandedCategories.has(category.id) &&
                    category.children &&
                    category.children.map((subcategory) => (
                      <div
                        key={subcategory.id}
                        className="flex items-center justify-between p-3 pl-16 bg-gray-50 border-t border-gray-200"
                      >
                        <div className="flex items-center gap-2">
                          <div className="w-px h-6 bg-gray-300 mr-2" />
                          <span className="text-gray-700">{subcategory.name}</span>
                        </div>
                        <div className="flex items-center gap-2">
                          <button
                            onClick={() => handleEdit(subcategory)}
                            className="text-blue-600 hover:text-blue-700 p-2"
                            title="Editar"
                          >
                            <Edit2 className="w-4 h-4" />
                          </button>
                          <button
                            onClick={() => handleDelete(subcategory)}
                            className="text-red-600 hover:text-red-700 p-2"
                            title="Excluir"
                          >
                            <Trash2 className="w-4 h-4" />
                          </button>
                        </div>
                      </div>
                    ))}
                </div>
              ))}
            </div>
          )}
        </div>
      </div>

      {/* Modal */}
      {showModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 w-full max-w-md">
            <h2 className="text-xl font-bold mb-4">
              {modalMode === "create"
                ? formData.parent_id
                  ? "Nova Subcategoria"
                  : "Nova Categoria"
                : "Editar Categoria"}
            </h2>
            <form onSubmit={handleSubmit}>
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  Nome *
                </label>
                <input
                  type="text"
                  value={formData.name}
                  onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                  className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="Digite o nome da categoria"
                  required
                />
              </div>
              <div className="flex gap-3 justify-end">
                <button
                  type="button"
                  onClick={() => setShowModal(false)}
                  className="px-4 py-2 text-gray-700 bg-gray-200 rounded-lg hover:bg-gray-300"
                >
                  Cancelar
                </button>
                <button
                  type="submit"
                  className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
                >
                  {modalMode === "create" ? "Criar" : "Salvar"}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}

export default ProductCategoriesAdmin;
