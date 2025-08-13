import { useEffect, useState } from "react";
import {
  Users,
  MapPin,
  HardHat,
  Edit2,
  Trash2,
  ArrowRight,
  Calendar,
  User,
  Phone,
  ArrowLeft,
  Save,
} from "lucide-react";
import MarketplaceProducts from "./MarketplaceProducts";
import Banner from "./Banner";
import {
  ProfessionalService,
  UserService,
} from "../services";
import Swal from "sweetalert2";
import DashboardProfessional from "./DashboardProfissionais";
import { StoreService } from "../services/StoreService";
import { ClientService } from "../services/ClientService";
import EditProfileDashboard from "./EditProfileDashboard";

import SearchProfessionalsByUF from "./SearchProfessionalsByUF";
import Pagination from "../components/Pagination";
import LoadingText from "../components/LoadingText";
import DashboardProfessions from "./DashboardProfissions";
import DashboardRegions from "./DashboardRegions";

function Dashboard() {

  const [selectedDashboardSection, setSelectedDashboardSection] =
    useState("dashboard");
  const [professionCount, setProfessionCount] = useState([{}]);
  const [professionCountTotal, setProfessionCountTotal] = useState(0);
  const [stateCount, setStateCount] = useState([{}]);
  const [stateCountTotal, setStateCountTotal] = useState(0);
  const [showEditModal, setShowEditModal] = useState(false);
  const [showAll, setShowAll] = useState(false);
  const [isUpdate, setIsUpdate] = useState(false);
  const [editDados, setEditDados] = useState({ id: 0, profile: "" });
  const [recentProfessionals, setRecentProfessionals] = useState([
    {
      oid: 0,
      cidade: { nome: "", uf: "" },
      profissoes: [{ nome: "" }, { nome: "" }],
    },
  ]);
  const [recentClients, setRecentClients] = useState([
    {
      oid: 0,
      cidade: { nome: "", uf: "" },
    },
  ]);
  const [recentStores, setRecentStores] = useState([
    {
      oid: 0,
      cidade: { nome: "", uf: "" },
    },
  ]);

  const [pageProfessional, setPageProfessional] = useState(1);
  const [totalPageProfessional, setTotalPageProfessional] = useState(1);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {

    const fetchData = async () => {
      setIsLoading(true);
      
        const responseLastProfessional =
          await ProfessionalService.getProfessionals(
            50,
            (pageProfessional - 1) * 50,
            "",
            "",
            0,
            "created_at DESC"
          );
        const jsonLast = await responseLastProfessional.data;
        const total = await responseLastProfessional.data.total;
        setTotalPageProfessional(Math.ceil(total / 50));

        if (responseLastProfessional.status == 200) {
          setRecentProfessionals(jsonLast.profissionais);
        }

        if (pageProfessional== 1) {
          const response =
            await ProfessionalService.countProfessionalByProfession();
          const json = response.data;

          if (response.status == 200) {
            setProfessionCount(json);
            const Total = json.reduce((a, v) => (a = a + v.quantity), 0);
            setProfessionCountTotal(Total);
          }

          const responseUF = await ProfessionalService.countProfessionalByState(
            {
              limit: 1000000,
              offset: 0,
            }
          );
          const jsonUF = responseUF.data.result;

          if (responseUF.status == 200) {
            setStateCount(jsonUF);
            const TotalUF = await jsonUF.reduce(
              (a, v) => (a = a + v.quantity),
              0
            );
            /*await jsonUF.map(
          (uf: any) => (uf.name = states.find((x: any) => x.id == uf.UF))
        );*/

            //json.profissoes.map((x: any) => x.oid);
            //TotalUF;
            setStateCountTotal(TotalUF);
          }

          const responseLastClient = await ClientService.lastClients({
            quantity: 50,
          });
          const jsonLastClient = responseLastClient.data;

          if (responseLastClient.status == 200) {
            setRecentClients(jsonLastClient);
          }

          const responseLastStore = await StoreService.lastStores({
            quantity: 50,
          });
          const jsonLastStore = responseLastStore.data;

          if (responseLastStore.status == 200) {
            setRecentStores(jsonLastStore);
          }
        }
      setIsLoading(false);
    };

    fetchData();

  }, [isUpdate, selectedDashboardSection, showEditModal,pageProfessional]);

  const handleDelete = async (id: string, profile: string) => {
    let response: any;
    if (window.confirm("Tem certeza que deseja excluir?" + profile)) {
      if (profile == "profissional") {
        response = await ProfessionalService.deleteProfessional(parseInt(id));
      } else if (profile == "client") {
        response = await ClientService.deleteClient(parseInt(id));
      } else if (profile == "store") {
        response = await StoreService.deleteStore(parseInt(id));
      }

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

  const handleEdit = async (email: string, profile: string) => {
    const response = await UserService.findbyemail({ email: email });
    if ((response.status == 200)) {
      const json = await response.data;
      const id = json.oid;
      setEditDados({ id, profile });
      setShowEditModal(true);
      console.log("Editing professional:", id);
    }
  };


  const DownloadProfessionalXLSX = async () => {
    ProfessionalService.postProfessionalXLSX();    
  };

  const DownloadClientXLSX = async () => {
    ClientService.postClientXLSX();
  };

  const DownloadStoreXLSX = async () => {
    StoreService.postStoreXLSX();
  };

  const DashboardSections = [
    { id: "dashboard", label: "Dashboard" },
    { id: "carrousel", label: "Gerenciar Imagens do Carrossel" },
    { id: "products", label: "Produtos Recebidos" },
    { id: "professionalbyCiity", label: "Profissionais por Estado" },
    { id: "professions", label: "Profissões" },
    { id: "regions", label: "Regiões" },
  ];

  const renderDashboard = () => {
    return (
      <>
        {/* Loading */}
        {isLoading && <LoadingText />}

        {/* Statistics Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center gap-3 mb-4">
              <div className="p-3 bg-blue-100 rounded-lg">
                <Users className="w-6 h-6 text-blue-600" />
              </div>
              <h3 className="text-lg font-semibold text-gray-900">
                Total de Profissionais
              </h3>
            </div>
            <p className="text-3xl font-bold text-gray-900">
              {stateCountTotal}
            </p>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center gap-3 mb-4">
              <div className="p-3 bg-green-100 rounded-lg">
                <Calendar className="w-6 h-6 text-green-600" />
              </div>
              <h3 className="text-lg font-semibold text-gray-900">
                Cadastros Hoje
              </h3>
            </div>
            <p className="text-3xl font-bold text-gray-900">
              {!isLoading &&
                recentProfessionals.length != 0 &&
                recentProfessionals.filter(
                  (p) =>
                    new Date(p.created_at).toDateString() ===
                    new Date().toDateString()
                ).length +
                  recentClients.filter(
                    (p) =>
                      new Date(p.created_at).toDateString() ===
                      new Date().toDateString()
                  ).length +
                  recentStores.filter(
                    (p) =>
                      new Date(p.created_at).toDateString() ===
                      new Date().toDateString()
                  ).length}
            </p>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center gap-3 mb-4">
              <div className="p-3 bg-purple-100 rounded-lg">
                <MapPin className="w-6 h-6 text-purple-600" />
              </div>
              <h3 className="text-lg font-semibold text-gray-900">
                Estados Atendidos
              </h3>
            </div>
            <p className="text-3xl font-bold text-gray-900">
              {stateCount.length}
            </p>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center gap-3 mb-4">
              <div className="p-3 bg-orange-100 rounded-lg">
                <HardHat className="w-6 h-6 text-orange-600" />
              </div>
              <h3 className="text-lg font-semibold text-gray-900">
                Profissões
              </h3>
            </div>
            <p className="text-3xl font-bold text-gray-900">
              {professionCount.length}
            </p>
          </div>
        </div>
        {/* Recent Professionals */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="flex items-center justify-between mb-6">
            <h3 className="text-lg font-semibold text-gray-900">
              Últimos Profissionais Cadastrados
            </h3>
            <button
              onClick={() => {
                DownloadProfessionalXLSX();
              }}
              className="flex items-center gap-2 text-blue-600 hover:text-blue-700"
            >
              Download XLSX
              <Save className="w-4 h-4" />
            </button>
            <button
              onClick={() => {
                setShowAll(true);
              }}
              className="flex items-center gap-2 text-blue-600 hover:text-blue-700"
            >
              Ver todos
              <ArrowRight className="w-4 h-4" />
            </button>
          </div>
          {/* Paginacao */}
          <Pagination
            currentPage={pageProfessional}
            totalPages={totalPageProfessional}
            handleNextPage={() => {
              setPageProfessional(pageProfessional + 1);
            }}
            handlePrevPage={() => {
              setPageProfessional(pageProfessional - 1);
            }}
          />

          <div className="mt-6 overflow-x-auto bg-white shadow-md rounded-xl">
            <table className="w-full table-auto text-sm">
              <thead className="bg-gray-50">
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
                    Estado
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Cidade
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Profissão
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Data
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Ações
                  </th>
                </tr>
              </thead>
              <tbody>
                {recentProfessionals.map((prof) => (
                  <tr className="border-b border-gray-100">
                    <td className="px-4 py-3">
                      <div className="w-10 h-10 rounded-full overflow-hidden bg-gray-100">
                        {prof.image ? (
                          <img
                            src={"data:image;base64," + prof.image}
                            alt={prof.nome}
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
                      {prof.cidade.uf}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {prof.cidade.nome}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {prof.profissoes.map((prof1) => (
                        <span>
                          {prof1.nome}
                          <br></br>
                        </span>
                      ))}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {new Date(prof.created_at).toLocaleDateString("pt-BR")}
                    </td>
                    <td className="px-4 py-3">
                      <div className="flex items-center gap-2">
                        <button
                          onClick={() => handleEdit(prof.email, "profissional")}
                          className="p-1 text-blue-600 hover:bg-blue-50 rounded"
                        >
                          <Edit2 className="w-4 h-4" />
                        </button>
                        <button
                          onClick={() => handleDelete(prof.oid, "profissional")}
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
        </div>
        {/* Recent Clients */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="flex items-center justify-between mb-6">
            <h3 className="text-lg font-semibold text-gray-900">
              Últimos Clientes Cadastrados
            </h3>
            <button
              onClick={() => {
                DownloadClientXLSX();
              }}
              className="flex items-center gap-2 text-blue-600 hover:text-blue-700"
            >
              Download XLSX
              <Save className="w-4 h-4" />
            </button>

            <button
              onClick={() => {
                setShowAll(true);
              }}
              className="flex items-center gap-2 text-blue-600 hover:text-blue-700"
            >
              Ver todos
              <ArrowRight className="w-4 h-4" />
            </button>
          </div>
          <div className="mt-6 overflow-x-auto bg-white shadow-md rounded-xl">
            <table className="w-full table-auto text-sm">
              <thead className="bg-gray-50">
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
                    Estado
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Cidade
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Profissão
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Data
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Ações
                  </th>
                </tr>
              </thead>
              <tbody>
                {recentClients.map((prof) => (
                  <tr key={prof.oid} className="border-b border-gray-100">
                    <td className="px-4 py-3">
                      <div className="w-10 h-10 rounded-full overflow-hidden bg-gray-100">
                        {prof.image ? (
                          <img
                            src={"data:image;base64," + prof.image}
                            alt={prof.nome}
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
                      {prof.cidade.uf}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {prof.cidade.nome}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      <span></span>
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {new Date(prof.created_at).toLocaleDateString("pt-BR")}
                    </td>
                    <td className="px-4 py-3">
                      <div className="flex items-center gap-2">
                        <button
                          onClick={() => handleEdit(prof.email, "client")}
                          className="p-1 text-blue-600 hover:bg-blue-50 rounded"
                        >
                          <Edit2 className="w-4 h-4" />
                        </button>
                        <button
                          onClick={() => handleDelete(prof.oid, "client")}
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
        </div>
        {/* Recent Stores */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="flex items-center justify-between mb-6">
            <h3 className="text-lg font-semibold text-gray-900">
              Últimos Lojistas Cadastrados
            </h3>
            <button
              onClick={() => {
                DownloadStoreXLSX();
              }}
              className="flex items-center gap-2 text-blue-600 hover:text-blue-700"
            >
              Download XLSX
              <Save className="w-4 h-4" />
            </button>

            <button
              onClick={() => {
                setShowAll(true);
              }}
              className="flex items-center gap-2 text-blue-600 hover:text-blue-700"
            >
              Ver todos
              <ArrowRight className="w-4 h-4" />
            </button>
          </div>
          <div className="mt-6 overflow-x-auto bg-white shadow-md rounded-xl">
            <table className="w-full table-auto text-sm">
              <thead className="bg-gray-50">
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
                    Estado
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Cidade
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Profissão
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Data
                  </th>
                  <th className="px-4 py-3 text-left text-sm font-medium text-gray-500">
                    Ações
                  </th>
                </tr>
              </thead>
              <tbody>
                {recentStores.map((prof) => (
                  <tr key={prof.oid} className="border-b border-gray-100">
                    <td className="px-4 py-3">
                      <div className="w-10 h-10 rounded-full overflow-hidden bg-gray-100">
                        {prof.image ? (
                          <img
                            src={"data:image;base64," + prof.image}
                            alt={prof.nome}
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
                      {prof.cidade.uf}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {prof.cidade.nome}
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      <span></span>
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {new Date(prof.created_at).toLocaleDateString("pt-BR")}
                    </td>
                    <td className="px-4 py-3">
                      <div className="flex items-center gap-2">
                        <button
                          onClick={() => handleEdit(prof.email, "store")}
                          className="p-1 text-blue-600 hover:bg-blue-50 rounded"
                        >
                          <Edit2 className="w-4 h-4" />
                        </button>
                        <button
                          onClick={() => handleDelete(prof.oid, "store")}
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
        </div>
        {/* Professionals by State */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div className="bg-white rounded-lg shadow-md p-6">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">
              Profissionais por Estado
            </h3>
            <div className="space-y-4">
              {stateCount
                .filter((s) => s.quantity > 0)
                .sort((a, b) => b.quantity - a.quantity)
                .map((state) => (
                  <div
                    key={state.UF}
                    className="flex items-center justify-between"
                  >
                    <span className="text-gray-600">{state.UF}</span>
                    <div className="flex items-center gap-2">
                      <span className="font-semibold text-gray-900">
                        {state.quantity}
                      </span>
                      <div className="w-24 h-2 bg-gray-100 rounded-full overflow-hidden">
                        <div
                          className="h-full bg-blue-600 rounded-full"
                          style={{
                            width: `${
                              (state.quantity / stateCountTotal) * 100
                            }%`,
                          }}
                        />
                      </div>
                    </div>
                  </div>
                ))}
            </div>
          </div>

          {/* Professionals by Profession */}
          <div className="bg-white rounded-lg shadow-md p-6">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">
              Profissionais por Profissão
            </h3>
            <div className="space-y-4">
              {professionCount
                .filter((p) => p.quantity > 0)
                .sort((a, b) => b.quantity - a.quantity)
                .map((prof) => (
                  <div
                    key={prof.professionName}
                    className="flex items-center justify-between"
                  >
                    <span className="text-gray-600">{prof.professionName}</span>
                    <div className="flex items-center gap-2">
                      <span className="font-semibold text-gray-900">
                        {prof.quantity}
                      </span>
                      <div className="w-24 h-2 bg-gray-100 rounded-full overflow-hidden">
                        <div
                          className="h-full bg-green-600 rounded-full"
                          style={{
                            width: `${
                              (prof.quantity / professionCountTotal) * 100
                            }%`,
                          }}
                        />
                      </div>
                    </div>
                  </div>
                ))}
            </div>
          </div>
        </div>
      </>
    );};

  const renderMain = () => {
    return (
      <>
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
        {/* Carousel Image Management */}
        <div className="space-y-8">
          {/* Section Selection */}
          <div className="bg-white rounded-lg shadow-lg p-2">
            <div className="flex flex-wrap gap-3">
              {DashboardSections.map((section) => (
                <button
                  key={section.id}
                  onClick={() => setSelectedDashboardSection(section.id)}
                  className={`px-6 py-3 rounded-lg text-sm font-medium transition-colors ${
                    selectedDashboardSection === section.id
                      ? "bg-blue-600 text-white shadow-lg"
                      : "bg-gray-100 text-gray-700 hover:bg-gray-200"
                  }`}
                >
                  {section.label}
                </button>
              ))}
            </div>
          </div>

          {/* Section Content */}
          <div className="bg-white rounded-lg shadow-lg p-6">
            {selectedDashboardSection === "products" && <MarketplaceProducts />}
            {selectedDashboardSection === "carrousel" && <Banner />}
            {selectedDashboardSection === "professionalbyCiity" && (
              <SearchProfessionalsByUF />
            )}
            {selectedDashboardSection === "dashboard" && renderDashboard()}
            {selectedDashboardSection === "professions" && (
              <DashboardProfessions />
            )}
            {selectedDashboardSection === "regions" && (
              <DashboardRegions />
            )}
          </div>
        </div>
      </>
    );
  };

  return (
    <div className="space-y-6">
      {/* Success Modal */}
      {showAll ? (
        <div className="space-y-6">
          {" "}
          <button
            onClick={() => {
              setShowAll(false);
            }}
            className="flex items-center gap-2 text-gray-600 hover:text-blue-600 mb-8 transition-colors"
          >
            <ArrowLeft className="w-5 h-5" />
            Voltar para página inicial
          </button>
          <DashboardProfessional />
        </div>
      ) : (
        renderMain()
      )}
    </div>
  );
}

export default Dashboard;
