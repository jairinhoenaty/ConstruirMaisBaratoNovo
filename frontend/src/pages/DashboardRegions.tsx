import React, { useEffect, useState } from "react";
import { Edit2, Trash2, User, PlusCircle } from "lucide-react";

import {
  RegionService,
} from "../services";
import Swal from "sweetalert2";
import Pagination from "../components/Pagination";
import LoadingText from "../components/LoadingText";
import EditInsertRegions from "../components/EditInsertRegions";
import Select from "react-select";
import { states } from "../data";


function DashboardRegions() {
  const [regionsList, setRegionsList] = useState([]);

  const [showEditModal, setShowEditModal] = useState(false);
  const [isUpdate, setIsUpdate] = useState(false);
  const [editDados, setEditDados] = useState({ id: 0 });

  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);

  const [isLoading, setIsLoading] = useState(true);
  const [UFs, setUFs] = useState<{ label: string; value: string }[]>([]);
  const [selectedUf, setSelectedUf] = useState<string>("");

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
    const fetchData = async () => {
      setIsLoading(true);
      const regions = await RegionService.getRegions(
        10,
        (page - 1) * 10,
        selectedUf
      );
      const json_regions = await regions.data;
      if (regions.status == 200) {
        setTotalPage(Math.ceil(regions.data.total / 10));
        setRegionsList(json_regions.regions);
      }
      setIsLoading(false);
    };
    fetchData();
  }, [page, isUpdate, selectedUf]);

  /*const handleDelete = async (id: string) => {
    if (window.confirm("Tem certeza que deseja excluir este profissional?")) {
      const response = await RegionalService.deleteRegional(
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

        console.log("Deleting regional:", id);
      }
    }
  };

  const handleEdit = (id: string) => {
    setShowEditModal(true);
    console.log("Editing regional:", id);
  };
  */
  const handleDelete = async (id: string) => {
    if (window.confirm("Tem certeza que deseja excluir?")) {
      const response = await RegionService.deleteRegion(parseInt(id));

      if ((response.status == 200)) {
        Swal.fire({
          position: "center",
          icon: "success",
          title: "Registro Excluido",
          showConfirmButton: false,
          timer: 1500,
        });

        console.log("Deleting regional:", id);
        setIsUpdate(!isUpdate);
      }
    }
  };

  const handleEdit = async (id: number) => {
      setEditDados({id});
      setShowEditModal(true);
      console.log("Editing regional:", id);
  };

  return (
    <div className="space-y-6">
      {/* Loading */}
      {isLoading && <LoadingText />}

      {/* Success Modal */}
      {showEditModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 overflow-y-auto">
          <div className="flex flex-col items-center">
            <EditInsertRegions
              id={editDados.id}
              onClose={setShowEditModal}
            ></EditInsertRegions>
          </div>
        </div>
      )}

      {/* All Regionals Table */}
      {!isLoading && regionsList == null && (
        <div className="bg-white rounded-lg shadow-md p-6">
          <h3 className="text-lg font-semibold text-gray-900 mb-6">
            Nenhum cadastro encontrado
          </h3>
        </div>
      )}

      {/* All Regionals Table */}
      {!isLoading && regionsList != null && (
        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-6">
            <h3 className="text-2xl font-bold text-gray-900">
              Todas as Regiões
            </h3>

            <Pagination
              currentPage={page}
              totalPages={totalPage}
              handleNextPage={() => setPage(page + 1)}
              handlePrevPage={() => setPage(page - 1)}
            />

            <button
              className="flex items-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors shadow-md"
              onClick={() => handleEdit(0)} // Open modal for new region
            >
              <PlusCircle className="w-5 h-5" />
              Nova Região
            </button>
          </div>
          {/* Estado (UF) */}
          <div className="gap-4 mb-6">
            <label className="block text-sm font-medium text-gray-700">
              Estado (UF)
            </label>
            <Select
              options={UFs}
              //isDisabled={selectedUf != "" && formData.cityIds.length!=0}
              value={UFs.find((uf) => uf.value === selectedUf)} // valor inicial aqui
              onChange={(option) => setSelectedUf(option?.value || "")}
              placeholder="Selecione o estado"
            />
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
                    Cidades
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Ações
                  </th>
                </tr>
              </thead>
              <tbody>
                {regionsList.map((prof) => (
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
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {prof.cities.map((city) => (
                        <span>
                          {city.nome}
                          <br></br>
                        </span>
                      ))}
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

export default DashboardRegions;
