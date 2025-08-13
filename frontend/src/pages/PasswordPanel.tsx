import React, { useState } from "react";
import { Lock, Save } from "lucide-react";
import { ProfessionalService } from "../services/ProfessionalService";
import { UserService } from "../services/UserService";
import Swal from "sweetalert2";
import { LoginService } from "../services/LoginService";

function PasswordPanel() {
  const [formData, setFormData] = useState({
    password: "",
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

    if (formData.password !== "") {
      console.info(localStorage.getItem("email")??"");
      console.info(formData.password); 
      try {
        const response_password =  await LoginService.login({
          email: localStorage.getItem("email")??"",
          password: formData.password,
        });
        console.log(response_password);
        if (response_password.status==200&&!response_password.data.isLoged) { 
          Swal.fire({
            position: "center",
            icon: "error",
            title: "Senha atual inválida",
            showConfirmButton: true,
          });
          return 
        }              
      }
      catch {
        Swal.fire({
          position: "center",
          icon: "error",
          title: "Senha atual inválida",
          showConfirmButton: true,
        });
      }
    }

    if (formData.newPassword !== formData.confirmPassword) {
      Swal.fire({
        position: "center",
        icon: "error",
        title: "As senhas não coincidem!",
        showConfirmButton: true,
      });
      return;
    }


    const result_user = await UserService.getUserById(
      parseInt(localStorage.getItem("id") || "0")
    );
    if (result_user.status == 200) {
      const user_oid = result_user.data.oid;
      const user_name = result_user.data.name;
      const user_email = result_user.data.email;
      const user_perfil = result_user.data.perfil;
      const result = await UserService.saveUser({
        oid: parseInt(user_oid),
        name: user_name,
        email: user_email,
        perfil: user_perfil,
        senha: formData.newPassword,
      });

      if (result_user.status == 200) {
        Swal.fire({
          position: "center",
          icon: "success",
          title: "Sua senha foi alterada",
          showConfirmButton: false,
          timer: 1500,
        });
        console.log("Password updated:", formData);
      } else {
        console.log("error update!!");
      }

      // Reset form
      setFormData({ password: "", newPassword: "", confirmPassword: "" });
    } else {
      console.log("Erro get user!!");
    }
  };

  return (
    <div className="bg-white rounded-lg shadow-md p-8">
      <div className="flex items-center gap-3 mb-8">
        <Lock className="w-8 h-8 text-blue-600" />
        <h2 className="text-2xl font-bold text-gray-900">Alterar Senha</h2>
      </div>

      <form onSubmit={handleSubmit} className="max-w-md space-y-6">
        {/* Senha Atual */}
        <div>
          <label
            htmlFor="password"
            className="block text-sm font-medium text-gray-700"
          >
            Senha Atual
          </label>
          <div className="mt-1 relative rounded-md shadow-sm">
            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <Lock className="h-5 w-5 text-gray-400" />
            </div>
            <input
              type="password"
              name="password"
              id="password"
              required
              value={formData.password}
              onChange={handleChange}
              className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
              placeholder="••••••••"
            />
          </div>
        </div>

        <div></div>

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
  );
}

export default PasswordPanel;
