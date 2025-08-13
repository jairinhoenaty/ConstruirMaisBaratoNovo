import React, { useState } from "react";
import { Search, Edit2, Trash2, User, Phone } from "lucide-react";
import { states } from "../data";
import {
  ProfessionalService,
  ProfessionService,
  UserService,
} from "../services";
import Swal from "sweetalert2";
import { StoreService } from "../services/StoreService";
import { ClientService } from "../services/ClientService";
import Pagination from "../components/Pagination";
import EditProfileDashboard from "./EditProfileDashboard";

function DashboardProfessional() {
  const [professionalsList, setProfessionalList] = useState([]);

  const [searchTerm, setSearchTerm] = useState("");
  const [selectedState, setSelectedState] = useState("");
  const [selectedProfession, setSelectedProfession] = useState("");
  const [showEditModal, setShowEditModal] = useState(false);
  const [professions, setProfessions] = useState([]);
  const [isUpdate, setIsUpdate] = useState(false);
  const [editDados, setEditDados] = useState({ id: 0, profile: "" });

  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);

  const [selectedTipo, setSelectedTipo] = useState("profissional");

  React.useEffect(() => {
    const fetchData = async () => {
      if (professions.length == 0) {
        const professions = await ProfessionService.getProfessionsPublic();
        const json_professions = await professions.data;
        if ((professions.status == 200)) {
          setProfessions(json_professions);
        }
      }
      setProfessionalList(undefined);
      if (selectedTipo == "profissional") {
        console.log(selectedProfession);
        const response = await ProfessionalService.getProfessionals(
          10,
          (page - 1) * 10,
          searchTerm,
          selectedState,
          selectedProfession == "" ? 0 : parseInt(selectedProfession),
          null
        );
        const total = await response.data.total;
        setTotalPage(Math.ceil(total / 10));
        const json = await response.data.profissionais;

        if ((response.status == 200)) {
          setProfessionalList(json);
        }
      } else if (selectedTipo == "client") {
        const response = await ClientService.getClients(10, (page - 1) * 10);
        if ((response.status == 200)) {
          const total = await response.data.total;
          setTotalPage(Math.ceil(total / 10));
          const json = await response.data.clients;

          setProfessionalList(json);
        }
      } else if (selectedTipo == "store") {
        const response = await StoreService.getStores(10, (page - 1) * 10);
        if ((response.status == 200)) {
          const total = await response.data.total;
          setTotalPage(Math.ceil(total / 10));
          const json = await response.data.stores;

          setProfessionalList(json);
        }
      }
    };
    fetchData();
  }, [
    selectedTipo,
    page,
    searchTerm,
    selectedState,
    selectedProfession,
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
    let response: any;
    const profile = selectedTipo;
    if (window.confirm("Tem certeza que deseja excluir?" + profile)) {
      if (profile == "profissional") {
        response = await ProfessionalService.deleteProfessional(parseInt(id));
      } else if (profile == "client") {
        response = await ClientService.deleteClient(parseInt(id));
      } else if (profile == "store") {
        response = await StoreService.deleteStore(parseInt(id));
      }
      console.log(response);

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

  const handleEdit = async (email: string) => {
    console.log(email);
    const profile = selectedTipo;
    const response = await UserService.findbyemail({ email: email });
    if ((response.status == 200)) {
      const json = await response.data;
      console.log(json);
      const id = json.oid;
      setEditDados({ id, profile });
      setShowEditModal(true);
      console.log("Editing professional:", id);
    }
  };

  return (
    <div className="space-y-6">
      {/* Success Modal */}
      {showEditModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 overflow-y-auto">
          <div className="flex flex-col items-center">
            <EditProfileDashboard
              id={editDados.id}
              profile={editDados.profile}
              onClose={setShowEditModal}
            ></EditProfileDashboard>
          </div>
        </div>
      )}

      {/* Search and Filters */}
      <div className="bg-white rounded-lg shadow-md p-6">
        <div className="flex flex-col md:flex-row gap-4">
          <div className="w-full md:w-48">
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Tipo Cadastro
            </label>
            <select
              value={selectedTipo}
              onChange={(e) => {
                setSelectedTipo(e.target.value);
                setPage(1);
              }}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option key="profissional" value="profissional">
                Profissional
              </option>
              <option key="client" value="client">
                Cliente
              </option>
              <option key="store" value="store">
                Lojista
              </option>
            </select>
          </div>

          <div className="flex-1">
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Pesquisar Profissionais
            </label>
            <div className="relative">
              <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
              <input
                type="text"
                value={searchTerm}
                onChange={(e) => {
                  setPage(1);
                  setSearchTerm(e.target.value);
                }}
                placeholder="Buscar por nome..."
                className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
          <div className="w-full md:w-48">
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Estado
            </label>
            <select
              value={selectedState}
              onChange={(e) => setSelectedState(e.target.value)}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">Todos</option>
              {states.map((state) => (
                <option key={state.id} value={state.id}>
                  {state.name}
                </option>
              ))}
            </select>
          </div>
          <div className="w-full md:w-48">
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Profissão
            </label>
            <select
              value={selectedProfession}
              onChange={(e) => setSelectedProfession(e.target.value)}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">Todas</option>
              {professions.map((prof) => (
                <option key={prof.id} value={prof.id}>
                  {prof.name}
                </option>
              ))}
            </select>
          </div>
        </div>
      </div>
      {/* All Professionals Table */}
      {professionalsList == null && (
        <div className="bg-white rounded-lg shadow-md p-6">
          <h3 className="text-lg font-semibold text-gray-900 mb-6">
            Nenhum cadastro encontrado
          </h3>
        </div>
      )}

      {/* All Professionals Table */}
      {professionalsList != null && (
        <div className="bg-white rounded-lg shadow-md p-6">
          <h3 className="text-lg font-semibold text-gray-900 mb-6">
            Todos os Profissionais
          </h3>
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
          <div className="overflow-x-auto">
            <table className="w-full responsive-table">
              <thead>
                <tr className="border-b border-gray-200">
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Foto
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Nome
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Telefone
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    E-mail
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Estado
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Cidade
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Profissão
                  </th>
                  {/*<th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Status
                  </th>
                  */}
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Ações
                  </th>
                </tr>
              </thead>
              <tbody>
                {professionalsList.map((prof) => (
                  <tr key={prof.oid} className="border-b border-gray-100">
                    {prof.name}
                    <td className="px-4 py-3">
                      <div className="w-10 h-10 rounded-full overflow-hidden bg-gray-100">
                        {prof.image ? (
                          <img
                            src={"data:image;base64," + prof.image}
                            alt={prof.oid}
                            className="w-full h-full object-cover"
                          />
                        ) : (
                          <User className="w-full h-full p-2 text-gray-400" />
                        )}
                      </div>
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-900">
                      {prof.nome}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      <div className="flex items-center gap-1">
                        <Phone className="w-4 h-4" />
                        {prof.telefone}
                      </div>
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {prof.email}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {prof.cidade.uf}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {prof.cidade.nome}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {
                        /*professionals.find((p) => p.id === prof.profissoes)?.nome*/
                        prof.profissoes != null ? (
                          prof.profissoes.map((prof1) => (
                            <span>
                              {prof1.nome}
                              <br></br>
                            </span>
                          ))
                        ) : (
                          <span></span>
                        )
                      }
                    </td>
                    {/*<td className="px-4 py-3">
                      <span
                        className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                          prof.status === "active"
                            ? "bg-green-100 text-green-800"
                            : prof.status === "pending"
                            ? "bg-yellow-100 text-yellow-800"
                            : "bg-red-100 text-red-800"
                        }`}
                      >
                        {prof.status === "active"
                          ? "Ativo"
                          : prof.status === "pending"
                          ? "Pendente"
                          : "Inativo"}
                      </span>
                    </td>*/}
                    <td className="px-4 py-3">
                      <div className="flex items-center gap-2">
                        <button
                          onClick={() => handleEdit(prof.email)}
                          className="p-1 text-blue-600 hover:bg-blue-50 rounded"
                        >
                          <Edit2 className="w-4 h-4" />
                        </button>
                        <button
                          onClick={() => handleDelete(prof.oid)}
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

export default DashboardProfessional;
