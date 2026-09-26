import React, { useCallback, useEffect, useState } from "react";
import { Edit2, Trash2, PlusCircle, HardHat } from "lucide-react";
import Swal from "sweetalert2";

import { ProfessionCategoryService } from "../services/ProfessionCategoryService";
import { IProfessionCategory } from "../interfaces/IProfessionCategory";
import Pagination from "../components/Pagination";
import LoadingText from "../components/LoadingText";
import EditInsertProfessionCategory from "./EditInsertProfessionCategory";

const PAGE_SIZE = 10;

function DashboardProfessionCategories() {
  const [categories, setCategories] = useState<IProfessionCategory[]>([]);
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);
  const [isLoading, setIsLoading] = useState(true);
  const [editingId, setEditingId] = useState<number | null>(null);

  const fetchCategories = useCallback(async () => {
    setIsLoading(true);
    try {
      const response = await ProfessionCategoryService.getCategories(
        PAGE_SIZE,
        (page - 1) * PAGE_SIZE
      );
      if (response.status === 200) {
        setTotalPage(Math.ceil((response.data.total || 0) / PAGE_SIZE));
        setCategories(response.data.categorias || []);
      }
    } catch (error) {
      console.error("Erro ao buscar categorias:", error);
    } finally {
      setIsLoading(false);
    }
  }, [page]);

  useEffect(() => {
    fetchCategories();
  }, [fetchCategories]);

  const handleDelete = async (category: IProfessionCategory) => {
    const confirmation = await Swal.fire({
      title: "Tem certeza?",
      text: `Deseja excluir a categoria "${category.name}"?`,
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#d33",
      confirmButtonText: "Sim, excluir!",
      cancelButtonText: "Cancelar",
    });
    if (!confirmation.isConfirmed) return;

    try {
      await ProfessionCategoryService.deleteCategory(category.id);
      Swal.fire("Excluído!", "Categoria excluída com sucesso.", "success");
      fetchCategories();
    } catch (error: any) {
      Swal.fire(
        "Erro",
        error?.response?.data?.error || "Falha ao excluir a categoria",
        "error"
      );
    }
  };

  const handleClose = (shouldReload: boolean) => {
    setEditingId(null);
    if (shouldReload) fetchCategories();
  };

  return (
    <div className="space-y-6">
      {isLoading && <LoadingText />}

      {editingId !== null && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 overflow-y-auto p-4">
          <EditInsertProfessionCategory id={editingId} onClose={handleClose} />
        </div>
      )}

      <div className="bg-white rounded-lg shadow-md p-6">
        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-6">
          <h3 className="text-2xl font-bold text-gray-900">
            Categorias de Profissões
          </h3>

          <Pagination
            currentPage={page}
            totalPages={totalPage}
            handleNextPage={() => setPage(page + 1)}
            handlePrevPage={() => setPage(page - 1)}
          />

          <button
            className="flex items-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors shadow-md"
            onClick={() => setEditingId(0)}
          >
            <PlusCircle className="w-5 h-5" />
            Nova Categoria
          </button>
        </div>

        {!isLoading && categories.length === 0 && (
          <p className="text-gray-500">Nenhuma categoria cadastrada.</p>
        )}

        {categories.length > 0 && (
          <div className="overflow-x-auto">
            <table className="w-full responsive-table">
              <thead>
                <tr className="border-b border-gray-200">
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Ícone
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Nome
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Profissões
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Deslocamento
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Situação
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Ações
                  </th>
                </tr>
              </thead>
              <tbody>
                {categories.map((category) => (
                  <tr key={category.id} className="border-b border-gray-100">
                    <td className="px-4 py-3">
                      <div className="w-10 h-10 rounded-full overflow-hidden bg-gray-100">
                        {category.icon ? (
                          <img
                            src={category.icon}
                            alt=""
                            className="w-full h-full object-cover"
                          />
                        ) : (
                          <HardHat className="w-full h-full p-2 text-gray-400" />
                        )}
                      </div>
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-900">
                      {category.name}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {category.profession_count}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {category.client_travels
                        ? "Cliente vai até o profissional"
                        : "Profissional vai até o cliente"}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {category.active ? "Ativa" : "Inativa"}
                    </td>
                    <td className="px-4 py-3">
                      <div className="flex items-center gap-2">
                        <button
                          onClick={() => setEditingId(category.id)}
                          className="p-1 text-blue-600 hover:bg-blue-50 rounded"
                        >
                          <Edit2 className="w-4 h-4" />
                        </button>
                        <button
                          onClick={() => handleDelete(category)}
                          className="p-1 text-red-600 hover:bg-red-50 rounded"
                        >
                          <Trash2 className="w-4 h-4" />
                        </button>
                      </div>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}

        <Pagination
          currentPage={page}
          totalPages={totalPage}
          handleNextPage={() => setPage(page + 1)}
          handlePrevPage={() => setPage(page - 1)}
        />
      </div>
    </div>
  );
}

export default DashboardProfessionCategories;
