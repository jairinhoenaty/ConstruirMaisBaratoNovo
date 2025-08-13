import React, { useState } from "react";
import {
  Calendar,
  User,
  Phone,
  MessageSquare,
  HardHat,
  MapPin,
  Trash2,
  Mail,
} from "lucide-react";
import { states} from "../data";
import { BudgetService } from "../services/Budget";
import Pagination from "../components/Pagination";
import { IBudget } from "../interfaces";
import Swal from "sweetalert2";

interface Quote {
  id: string;
  date: string;
  clientName: string;
  telephone: string;
  description: string;
  profession: string;
  state: string;
  city: string;
}

function QuotesPanel() {
  const [quotes, setQuotes] = useState<Quote[]>([
      
  ]);
  const [page,setPage]=useState(1);
  const [totalPage,setTotalPage]=useState(0)
  const [isUpdate, setIsUpdate] = useState(false);  
  const profile = localStorage.getItem("profile");      

  const [professionals,setProfessionals] = useState([]);
  const citiesByState = [];

  React.useEffect(() => {

    console.log("useefect");
    const fetchData = async () => {
      const limit=10;
      let json;
      let total=0;
      const id = parseInt(localStorage.getItem("id") || "");
      const post_id = parseInt(localStorage.getItem("post_id") || "");
      let response;
      if (profile=="admin") {
      response = await BudgetService.getBudgetsAll(limit, ((page - 1) * limit)
      );
      console.log(response);
      json = response.data.budgets;
      total = response.data.total;

    }
      else {
      response = await BudgetService.getBudgetsbyMonth({
        month: "April",
        professionalID: profile == "profissional" ? id : 0,
        clientID: profile == "client" ? post_id : 0,
        page: 1,
        pagesize: 10,
      });
      json = response.data;
      //total = response.total

    }
      console.log(json);
      setProfessionals([]);

      if ((response.status == 200)) {
        setQuotes(json);
        setTotalPage(Math.ceil(total / limit));

      }
    };

    fetchData();
  }, [page,isUpdate]);

  const handleDeleteQuote = async (id: string) => {
    if (window.confirm("Tem certeza que deseja excluir este orçamento?")) {
      const response = await BudgetService.deleteBudget(id);
      if ((response.status == 200)) {
        Swal.fire({
          icon: "success",
          text: "Mensagem excluída",
          showConfirmButton: false,
          timer: 1500,
        });
        setIsUpdate(!isUpdate);
      };
    }
  };

  const handleStatusChange = async (
    budget_approved: IBudget,
    approved: boolean
  ) => {
    budget_approved.approved = !approved;
    console.log(budget_approved);

    const response = await BudgetService.saveBudget(budget_approved);
    if (response.status == 200) {
      Swal.fire({
        icon: "success",
        text: "Orçamento alterado",
        showConfirmButton: false,
        timer: 2000,
      });
      setIsUpdate(!isUpdate);
    }
  };

  const renderClientInfo = (message) => {
    return (
      <>
        <div className="border-t border-gray-100 pt-4"></div>
        <div className="space-y-4">
          <h2 className="text-lg font-semibold text-gray-700">
            Dados do Cliente
          </h2>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {/* Coluna 1 */}
            <div className="space-y-4">
              <div className="flex items-start gap-3">
                <User className="w-5 h-5 mt-1 text-gray-400" />
                <div>
                  <span className="block text-sm text-gray-500">Nome</span>
                  <p className="text-gray-900 font-medium">{message.name}</p>
                </div>
              </div>

              
                <div className="flex items-start gap-3">
                  <Mail className="w-5 h-5 mt-1 text-gray-400" />
                  <div>
                    <span className="block text-sm text-gray-500">Email:</span>
                    <p className="text-gray-900 font-medium">{message.email}</p>
                  </div>
                </div>
            </div>

            {/* Coluna 2 */}
            <div className="space-y-4">
                <div className="flex items-start gap-3">
                  <Phone className="w-5 h-5 mt-1 text-gray-400" />
                  <div>
                    <span className="block text-sm text-gray-500">
                      WhatsApp:
                    </span>
                    <p className="text-gray-900 font-medium">
                      {message.telephone}
                    </p>
                  </div>
                </div>

              <div className="flex items-start gap-3">
                <MapPin className="w-5 h-5 mt-1 text-gray-400" />
                <div>
                  <span className="block text-sm text-gray-500">
                    UF/Cidade:
                  </span>
                  <p className="text-gray-900 font-medium">
                    {message.client.city.name + "/" + message.client.city.uf}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </>
    );
  };

  const renderProfessionalInfo = (message) => {
    console.log("message");
    console.log(message);
    console.log("Professions");
    console.log(message.professionals[0]?.professions);   
    
      message.professionals[0]?.professions.map((profession) => {
        console.log(profession.name)
      });
    
    return (
      <>
        <div className="border-t border-gray-100 pt-4"></div>
        <div className="space-y-4">
          <h2 className="text-lg font-semibold text-gray-700">
            Dados do Profissional
          </h2>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {/* Coluna 1 */}
            <div className="space-y-4">
              <div className="flex items-start gap-3">
                <User className="w-5 h-5 mt-1 text-gray-400" />
                <div>
                  <span className="block text-sm text-gray-500">Nome</span>
                  <p className="text-gray-900 font-medium">
                    {message.professionals[0]?.name}
                  </p>
                </div>
              </div>

              {/*profile === "admin" && (*/}
                <div className="flex items-start gap-3">
                  <Mail className="w-5 h-5 mt-1 text-gray-400" />
                  <div>
                    <span className="block text-sm text-gray-500">Email:</span>
                    <p className="text-gray-900 font-medium">
                      {message.professionals[0]?.email}
                    </p>
                  </div>
                </div>
              {/*)*/}
            </div>

            {/* Coluna 2 */}
            <div className="space-y-4">
                <div className="flex items-start gap-3">
                  <Phone className="w-5 h-5 mt-1 text-gray-400" />
                  <div>
                    <span className="block text-sm text-gray-500">
                      WhatsApp:
                    </span>
                    <p className="text-gray-900 font-medium">
                      {message.professionals[0]?.telephone}
                    </p>
                  </div>
                </div>

              <div className="flex items-start gap-3">
                <MapPin className="w-5 h-5 mt-1 text-gray-400" />
                <div>
                  <span className="block text-sm text-gray-500">
                    UF/Cidade:
                  </span>
                  <p className="text-gray-900 font-medium">
                    {message.professionals[0]?.city.name +
                      "/" +
                      message.professionals[0]?.city.uf}
                  </p>
                </div>
              </div>
            </div>
            {/* Profissões - em nova linha mesmo em telas pequenas */}
            <div className="md:col-span-2 flex flex-wrap items-center gap-2">
              <HardHat className="w-5 h-5 text-gray-400" />
              <div className="flex flex-wrap gap-2">
                <span className="text-sm text-gray-500">Profissão:</span>
                {message.professionals[0]?.professions.map(
                  (element: any, index: number) => (
                    <span
                      key={index}
                      className="text-gray-900 bg-blue-100 text-blue-800 text-xs font-medium px-2.5 py-0.5 rounded dark:bg-blue-900 dark:text-blue-300"
                    >
                      {element.name}
                    </span>
                  )
                )}
              </div>
            </div>
          </div>
        </div>
      </>
    );
  };



  return (
    <div className="bg-white rounded-lg shadow-md p-8">
      <div className="flex items-center gap-3 mb-8">
        <Calendar className="w-8 h-8 text-blue-600" />
        <h2 className="text-2xl font-bold text-gray-900">Orçamentos</h2>
      </div>

      {quotes.length == 0 && <div>Nenhum orçamento</div>}
      {quotes.length != 0 && totalPage!==0 &&(
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
      )}

      <div className="space-y-6">
        {quotes.map((quote) => {
          const profession = professionals.find(
            (p) => p.id === quote.profession
          );
          const state = states.find((s) => s.id === quote.state);
          const city = citiesByState[quote.state]?.find(
            (c) => c.id === quote.city
          );

          return (
            <div
              key={quote.id}
              className="border border-gray-200 rounded-lg p-6 hover:border-blue-500 transition-colors"
            >
              <div className="flex justify-between items-start mb-4">
                <div className="flex items-center gap-2">
                  <Calendar className="w-5 h-5 text-gray-400" />
                  <span className="text-gray-900">
                    {new Date(quote.created_at).toLocaleDateString("pt-BR")}
                  </span>
                </div>
                <div className="flex justify-between items-center mb-4">
                  {profile == "admin" && (
                    <label className="inline-flex items-center cursor-pointer">
                      <input
                        type="checkbox"
                        className="form-checkbox h-5 w-5 text-blue-600"
                        checked={quote.approved}
                        onChange={() =>
                          handleStatusChange(quote, quote.approved)
                        }
                      />
                      <span className="ml-2 text-sm text-gray-700">
                        Aprovar
                      </span>
                    </label>
                  )}
                  <button
                    onClick={() => handleDeleteQuote(quote.id)}
                    className="text-red-600 hover:bg-red-50 p-2 rounded-lg transition-colors"
                  >
                    <Trash2 className="w-5 h-5" />
                  </button>
                </div>
              </div>
              {(profile == "admin" || profile == "client") &&
                renderProfessionalInfo(quote)}
              {(profile == "admin" || profile == "profissional") &&
                renderClientInfo(quote)}

              <div className="border-t border-gray-100 pt-4">
                <div className="flex items-start gap-2">
                  <MessageSquare className="w-5 h-5 text-gray-400 mt-1" />
                  <div>
                    <span className="text-sm text-gray-500">Descrição:</span>
                    <p className="mt-2 text-gray-900">{quote.description}</p>
                  </div>
                </div>
              </div>
              {profile == "admin" && (
                <div className="flex justify-end gap-3 mt-4">
                  <a
                    href={`https://wa.me/${quote.telephone.replace(/\D/g, "")}`}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="flex items-center gap-2 px-4 py-2 text-green-600 hover:bg-green-50 rounded-lg transition-colors"
                  >
                    <Phone className="w-5 h-5" />
                    Responder via WhatsApp
                  </a>
                </div>
              )}
            </div>
          );
        })}
      </div>
    </div>
  );
}

export default QuotesPanel;
