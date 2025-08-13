import React, { useState } from "react";
import { Lock, Save } from "lucide-react";
import { UserService } from "../services/UserService";
import Swal from "sweetalert2";
import { useNavigate, useParams } from "react-router-dom";

function ResetPassword() {
    const navigate = useNavigate();

    const { token } = useParams() as { token: string };
  const [formData, setFormData] = useState({
    newPassword: "",
    confirmPassword: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

      if (formData.newPassword !== "" || formData.confirmPassword !== "") {

          if (formData.newPassword !== formData.confirmPassword) {
              Swal.fire({
                  position: "center",
                  icon: "error",
                  title: "As senhas não coincidem!",
                  showConfirmButton: true,
              });
              return;
          }
          try {
              const result = await UserService.resetPassword({
                  email: token,
                  password: formData.newPassword,
              });

              if (result.status == 200) {
                  Swal.fire({
                      position: "center",
                      icon: "success",
                      title: "Sua senha foi alterada",
                      showConfirmButton: false,
                      timer: 1500,
                  });
                  console.log("Password updated:", formData);
                  navigate("/login");

              } else {
                  Swal.fire({
                      position: "center",
                      icon: "error",
                      title: "Erro ao atualizar a senha. Tente novamente.",
                      showConfirmButton: false,
                      timer: 1500,
                  });
                  console.log("error update!!");
              }
          } catch (error) {
              Swal.fire({
                  position: "center",
                  icon: "error",
                  title: "Erro ao atualizar a senha. Tente novamente." + error,
                  showConfirmButton: true,
              });
              console.error("Error updating password:", error);
          }
      }
      else {
        Swal.fire({
            position: "center",
            icon: "info",
            title: "As senhas não podem estar vazias!",
            text: "Por favor, preencha os campos de senha.",
            showConfirmButton: true,
        });

      }
  };

  return (
    
    <div className="min-h-screen bg-gray-50 flex flex-col justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <div className="bg-white rounded-lg shadow-md p-8">
          <div className="flex items-center gap-3 mb-8">
            <Lock className="w-8 h-8 text-blue-600" />
            <h2 className="text-2xl font-bold text-gray-900">Alterar Senha</h2>
          </div>

          <form onSubmit={handleSubmit} className="max-w-md space-y-6">
            {/* Nova Senha */}
            <div>
              <label
                htmlFor="newPassword"
                className="block text-sm font-medium text-gray-700"
              >
                Nova Senha
              </label>
              <div className="mt-1 relative rounded-md shadow-sm">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <Lock className="h-5 w-5 text-gray-400" />
                </div>
                <input
                  type="password"
                  name="newPassword"
                  id="newPassword"
                  required
                  value={formData.newPassword}
                  onChange={handleChange}
                  className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
                  placeholder="••••••••"
                />
              </div>
            </div>

            {/* Confirmar Nova Senha */}
            <div>
              <label
                htmlFor="confirmPassword"
                className="block text-sm font-medium text-gray-700"
              >
                Confirmar Nova Senha
              </label>
              <div className="mt-1 relative rounded-md shadow-sm">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <Lock className="h-5 w-5 text-gray-400" />
                </div>
                <input
                  type="password"
                  name="confirmPassword"
                  id="confirmPassword"
                  required
                  value={formData.confirmPassword}
                  onChange={handleChange}
                  className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
                  placeholder="••••••••"
                />
              </div>
            </div>

            {/* Botão Atualizar */}
            <div>
              <button
                type="submit"
                className="w-full flex items-center justify-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
              >
                <Save className="w-5 h-5" />
                Atualizar Senha
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}

export default ResetPassword;
