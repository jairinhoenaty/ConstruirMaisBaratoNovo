import React, { useState } from "react";
import { Edit2, Trash2, User, PlusCircle } from "lucide-react";

import {
  ProfessionService,
} from "../services";
import Swal from "sweetalert2";
import Pagination from "../components/Pagination";
import LoadingText from "../components/LoadingText";
import EditInsertProfessions from "./EditInsertProfessions";

function DashboardProfessions() {
  const [professionsList, setProfessionsList] = useState([]);

  const [showEditModal, setShowEditModal] = useState(false);
  const [isUpdate, setIsUpdate] = useState(false);
  const [editDados, setEditDados] = useState({ id: 0 });

  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);

  const [isLoading, setIsLoading] = useState(true);

  React.useEffect(() => {
    const fetchData = async () => {
        setIsLoading(true);
        const professions = await ProfessionService.getProfessions(10,(page - 1) * 10);
        const json_professions = await professions.data;
        if ((professions.status == 200)) {
          setTotalPage(Math.ceil(professions.data.total / 10));
          setProfessionsList(json_professions.profissoes);
        }
        setIsLoading(false);
    };
    fetchData();
  }, [
    page,
    isUpdate,
  ]);

  /*const handleDelete = async (id: string) => {
    if (window.confirm("Tem certeza que deseja excluir este profissional?")) {
      const response = await ProfessionalService.deleteProfessional(
        parseInt(id)
      );
      const json = response.data;

      if ((response.status == 200)) {
        Swal.fire({
          position: "center",
          icon: "success",
          title: "Profissional Excluido",
          showConfirmButton: false,
          timer: 1500,
        });

        console.log("Deleting professional:", id);
      }
    }
  };

  const handleEdit = (id: string) => {
    setShowEditModal(true);
    console.log("Editing professional:", id);
  };
  */
  const handleDelete = async (id: string) => {
    if (window.confirm("Tem certeza que deseja excluir?")) {
      const response = await ProfessionService.deleteProfession(parseInt(id));

      if ((response.status == 200)) {
        Swal.fire({
          position: "center",
          icon: "success",
          title: "Registro Excluido",
          showConfirmButton: false,
          timer: 1500,
        });

        console.log("Deleting professional:", id);
        setIsUpdate(!isUpdate);
      }
    }
  };

  const handleEdit = async (id: number) => {
      setEditDados({id});
      setShowEditModal(true);
      console.log("Editing professional:", id);
  };

  return (
    <div className="space-y-6">
      {/* Loading */}
      {isLoading && <LoadingText />}

      {/* Success Modal */}
      {showEditModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 overflow-y-auto">
          <div className="flex flex-col items-center">
            <EditInsertProfessions
              id={editDados.id}
              onClose={setShowEditModal}
            ></EditInsertProfessions>
          </div>
        </div>
      )}

      {/* All Professionals Table */}
      {!isLoading && professionsList == null && (
        <div className="bg-white rounded-lg shadow-md p-6">
          <h3 className="text-lg font-semibold text-gray-900 mb-6">
            Nenhum cadastro encontrado
          </h3>
        </div>
      )}

      {/* All Professionals Table */}
      {!isLoading && professionsList != null && (
        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-6">
            <h3 className="text-2xl font-bold text-gray-900">
              Todas as Profissões
            </h3>

              <Pagination
                currentPage={page}
                totalPages={totalPage}
                handleNextPage={() => setPage(page + 1)}
                handlePrevPage={() => setPage(page - 1)}
              />

              <button
                className="flex items-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors shadow-md"
                onClick={() => handleEdit(0)} // Open modal for new profession
              >
                <PlusCircle className="w-5 h-5" />
                Nova Profissão
              </button>
          </div>
          <div className="overflow-x-auto">
            <table className="w-full responsive-table">
              <thead>
                <tr className="border-b border-gray-200">
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Icone
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Nome
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Descrição
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Ações
                  </th>
                </tr>
              </thead>
              <tbody>
                {professionsList.map((prof) => (
                  <tr key={prof.id} className="border-b border-gray-100">
                    <td className="px-4 py-3">
                      <div className="w-10 h-10 rounded-full overflow-hidden bg-gray-100">
                        {prof.icon ? (
                          <img
                            src={prof.icon}
                            alt={prof.id}
                            className="w-full h-full object-cover"
                          />
                        ) : (
                          <User className="w-full h-full p-2 text-gray-400" />
                        )}
                      </div>
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-900">
                      {prof.name}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {prof.description}
                    </td>
                    <td className="px-4 py-3">
                      <div className="flex items-center gap-2">
                        <button
                          onClick={() => handleEdit(prof.id)}
                          className="p-1 text-blue-600 hover:bg-blue-50 rounded"
                        >
                          <Edit2 className="w-4 h-4" />
                        </button>
                        <button
                          onClick={() => handleDelete(prof.id)}
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
          <Pagination
            currentPage={page}
            totalPages={totalPage}
            handleNextPage={() => {
              setPage(page + 1);
            }}
            handlePrevPage={() => {
              setPage(page - 1);
            }}
          />
        </div>
      )}
    </div>
  );
}

export default DashboardProfessions;
