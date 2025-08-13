import React, { useState } from "react";
import { Mail, Phone, MapPin, MessageSquare, User, Trash2 } from "lucide-react";
import { ContactService } from "../services/ContactService";
import Pagination from "../components/Pagination";
import { IContact } from "../interfaces/IContact";
import Swal from "sweetalert2";


function ClientMessages() {
  const [messages, setMessages] = useState([]);
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);
  const [isUpdate, setIsUpdate] = useState(false);
  const profile = localStorage.getItem("profile");


  React.useEffect(() => {
    const fetchData = async () => {
      const profile = localStorage.getItem("profile") ?? "";
      //const id = localStorage.getItem("id") ?? "";
      const post_id = localStorage.getItem("post_id") ?? "";      
      let response;
      if (profile == "admin") {
        response = await ContactService.getContacts(10, (page - 1) * 10);
      } else if (profile == "profissional") {
        response = await ContactService.getContactsUser({
          limit: 10,
          offset: (page - 1) * 10,
          professional_id: parseInt(post_id),
          client_id: 0,
          store_id: 0,
        });
      } else if (profile == "client") {
        response = await ContactService.getContactsUser({
          limit: 10,
          offset: (page - 1) * 10,
          professional_id: 0,
          client_id: parseInt(post_id),
          store_id: 0,
        });
      } else if (profile == "store") {
        response = await ContactService.getContactsUser({
          limit: 10,
          offset: (page - 1) * 10,
          professional_id: 0,
          client_id: 0,
          store_id: parseInt(post_id),
        });
      }
      const json = response?.data;

      if ((response?.status == 200)) {
        setTotalPage(Math.ceil(json.total / 10));
        setMessages(json.contacts);
      }
    };
    fetchData();
  }, [page,isUpdate]);


  const handleStatusChange = async (
    contact_approved: IContact,
    approved: boolean
  ) => {
    contact_approved.approved = !approved;
    console.log(contact_approved);

    const response = await ContactService.saveContact(contact_approved);
    if (response.status == 200) {
      Swal.fire({
        icon: "success",
        text: "Mensagem alterada",
        showConfirmButton: false,
        timer: 1500,
      });
      setIsUpdate(!isUpdate);
    }
    
  };

  const handleDeleteMessage = async (id:number) => {
    if (window.confirm("Tem certeza que deseja excluir esta mensagem?")) {
      console.log(id);
      const response = await ContactService.deleteContact(id);
      if ((response.status == 200)) {
        Swal.fire({
          icon: "success",
          text: "Mensagem excluída",
          showConfirmButton: false,
          timer: 1500,
        });
        setIsUpdate(!isUpdate);
      }
    }
  };


  const renderClientInfo = (message) => {
    return (
      <>
        <div className="border-t border-gray-100 pt-4"></div>
        <div className="space-y-4">
          <h2 className="text-lg font-semibold text-gray-700">
            Dados do Comprador
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
                      {message.telefone}
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
                    {message.client_id != 0
                      ? message.client.City.Name + "/" + message.client.City.UF
                      : message.city.Name + "/" + message.city.UF}
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
    return (
      <>
        <div className="border-t border-gray-100 pt-4"></div>
        <div className="space-y-4">
          <h2 className="text-lg font-semibold text-gray-700">
            Dados do Vendedor
          </h2>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {/* Coluna 1 */}
            <div className="space-y-4">
              <div className="flex items-start gap-3">
                <User className="w-5 h-5 mt-1 text-gray-400" />
                <div>
                  <span className="block text-sm text-gray-500">Nome</span>
                  <p className="text-gray-900 font-medium">
                    {message.professional.Name}
                  </p>
                </div>
              </div>

              {profile === "admin" && (
                <div className="flex items-start gap-3">
                  <Mail className="w-5 h-5 mt-1 text-gray-400" />
                  <div>
                    <span className="block text-sm text-gray-500">Email:</span>
                    <p className="text-gray-900 font-medium">
                      {message.professional.Email}
                    </p>
                  </div>
                </div>
              )}
            </div>

            {/* Coluna 2 */}
            <div className="space-y-4">
              {profile === "admin" && (
                <div className="flex items-start gap-3">
                  <Phone className="w-5 h-5 mt-1 text-gray-400" />
                  <div>
                    <span className="block text-sm text-gray-500">
                      WhatsApp:
                    </span>
                    <p className="text-gray-900 font-medium">
                      {message.professional.Telephone}
                    </p>
                  </div>
                </div>
              )}

              <div className="flex items-start gap-3">
                <MapPin className="w-5 h-5 mt-1 text-gray-400" />
                <div>
                  <span className="block text-sm text-gray-500">
                    UF/Cidade:
                  </span>
                  <p className="text-gray-900 font-medium">
                    {message.professional.City.Name +
                      "/" +
                      message.professional.City.UF}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </>
    );
  };


  const renderStoreInfo = (message) => {
    console.log("message");
    console.log(message);
    return (
      <>
        <div className="border-t border-gray-100 pt-4"></div>
        <div className="space-y-4">
          <h2 className="text-lg font-semibold text-gray-700">
            Dados do Lojista
          </h2>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {/* Coluna 1 */}
            <div className="space-y-4">
              <div className="flex items-start gap-3">
                <User className="w-5 h-5 mt-1 text-gray-400" />
                <div>
                  <span className="block text-sm text-gray-500">Nome</span>
                  <p className="text-gray-900 font-medium">
                    {message.store.Name}
                  </p>
                </div>
              </div>

                <div className="flex items-start gap-3">
                  <Mail className="w-5 h-5 mt-1 text-gray-400" />
                  <div>
                    <span className="block text-sm text-gray-500">Email:</span>
                    <p className="text-gray-900 font-medium">
                      {message.store.Email}
                    </p>
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
                      {message.store.Telephone}
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
                    {message.store.City.Name +
                      "/" +
                      message.store.City.UF}
                  </p>
                </div>
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
        <MessageSquare className="w-8 h-8 text-blue-600" />
        <h2 className="text-2xl font-bold text-gray-900">Mensagens</h2>
      </div>
      {totalPage == 0 && <h1>Nenhuma mensagem</h1>}
      {totalPage > 0 && (
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
        {messages.map((message: IContact) => (
          <div
            key={message.id}
            className="border border-gray-200 rounded-lg p-6 hover:border-blue-500 transition-colors"
          >
            <div className="flex justify-between items-center text-sm text-gray-500">
              <span className="flex items-center gap-2">
                📅 {new Date(message.created_at).toLocaleString("pt-BR")}
              </span>

              <div className="flex items-center gap-3">
                {profile === "admin" && (
                  <label className="inline-flex items-center gap-2 text-sm text-gray-700">
                    <input
                      type="checkbox"
                      className="form-checkbox text-blue-600 h-5 w-5"
                      checked={message.approved}
                      onChange={() =>
                        handleStatusChange(message, message.approved)
                      }
                    />
                    Aprovar
                  </label>
                )}
                <button
                  onClick={() => handleDeleteMessage(message.id)}
                  className="text-red-600 hover:text-red-800 p-2 rounded-full hover:bg-red-100 transition"
                  title="Excluir mensagem"
                >
                  <Trash2 className="w-5 h-5" />
                </button>
              </div>
            </div>

            {/* Produto */}
            {message.product_id !== 0 && (
              <div className="flex items-center mb-4">
                <h3 className="text-lg font-semibold">
                  Produto:{" "}
                  <span className="text-lg font-semibold text-blue-600">
                    {message.product.Name}
                  </span>
                </h3>
              </div>
            )}

            {(profile == "admin" || profile == "client") &&
              message.professional_id !== 0 &&
              message.professional_id !== null &&
              renderProfessionalInfo(message)}
            {(profile == "admin" ||
              profile == "client") &&
              message.store_id !== 0 &&
              message.store_id !== null &&
              renderStoreInfo(message)}
            {(profile == "admin" || profile == "professional" || profile == "store") &&
              ((message.client_id !== 0  && message.client_id !== null )|| 
               (message.product_id !== 0 && message.product_id !== null ) 
              ) &&
              renderClientInfo(message)}
            <div className="border-t border-gray-100 pt-4">
              <span className="text-sm text-gray-500">Descrição:</span>
              <p className="mt-2 text-gray-900">{message.mensagem}</p>
            </div>
            {profile == "admin" && (
              <div className="flex justify-end gap-3 mt-4">
                <a
                  href={`https://wa.me/${message.telefone.replace(/\D/g, "")}`}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="flex items-center gap-2 px-4 py-2 text-green-600 hover:bg-green-50 rounded-lg transition-colors"
                >
                  <Phone className="w-5 h-5" />
                  Responder via WhatsApp
                </a>
                <a
                  href={`mailto:${message.email}`}
                  className="flex items-center gap-2 px-4 py-2 text-blue-600 hover:bg-blue-50 rounded-lg transition-colors"
                >
                  <Mail className="w-5 h-5" />
                  Responder via Email
                </a>
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}

export default ClientMessages;
