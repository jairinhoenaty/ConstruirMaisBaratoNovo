import React, { useState } from "react";
import { Mail, Lock, ArrowRight, Building2, EyeOff, Eye } from "lucide-react";
import { LoginService, UserService } from "../services";
import ErrorAlert from "../components/ErrorAlert";
import { Link, useNavigate } from "react-router-dom";

interface LoginProps {
  onNavigate?: (page: string) => void;
}

function Login({ onNavigate }: LoginProps) {
  const navigate = useNavigate();
  const [showForgotPassword, setShowForgotPassword] = useState(false);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
    const [showPassword, setShowPassword] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    // Simulate login and redirect to professional panel
    if (!showForgotPassword) {
      LoginService.login({
        email: email,
        password: password,
      })
        .then((response) => {
          if (response.data.isLoged == true) {
            localStorage.setItem("isLoggedIn", response.data.isLoged);
            localStorage.setItem("token", response.data.token);
            localStorage.setItem("user", response.data.user);
            localStorage.setItem("id", response.data.user.id);
            localStorage.setItem("name", response.data.user.name);
            localStorage.setItem("profile", response.data.user.profile);
            localStorage.setItem("email", response.data.user.email);
            //const profile = response.data.user.profile;
            /*if (profile == "profissional") {
            } else if (profile == "client") {
            } else if (profile == "store") {
            }*/

            if (response.data.user.profile === "admin") {
              navigate("/dashboard");
              //onNavigate && onNavigate("dashboard");
            } else {
              navigate("/professional-panel");
              //onNavigate && onNavigate("professional-panel");
            }
          } else {
            setError("Login Inválido!!");
          }
        })
        .catch((err) => {
          setError("Login Inválido!!");
          console.error("ops! ocorreu um erro" + err);
        });

      //localStorage.setItem('isLoggedIn', 'true');
      //onNavigate && onNavigate('professional-panel');
    } else {
      try {
        const result = await UserService.sendMail({
          email: email,
        });
        if (result.status==200) {
          setError("");
          alert("Instruções enviadas para o seu email.");
          setShowForgotPassword(false)
        } else {
          setError("Erro ao enviar instruções. Tente novamente.");
        }
        console.log("Reset password for:", email);
      } catch (err) {	
        console.error("Erro ao enviar email de recuperação:", err);
        setError("Erro ao enviar instruções. Tente novamente.");
      }
    }
  };

  const closeError = () => {
    setError("");
  };

  return (
    <div className="min-h-screen bg-gray-50 flex flex-col justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <div className="flex justify-center">
          <Building2 className="w-12 h-12 text-blue-600" />
        </div>
        <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
          {showForgotPassword ? "Recuperar Senha" : "Acessar sua conta"}
        </h2>
        <p className="mt-2 text-center text-sm text-gray-600">
          {showForgotPassword
            ? "Digite seu email para receber as instruções"
            : "Entre com suas credenciais para acessar"}
        </p>
      </div>

      <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
        <div className="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
          {error && <ErrorAlert message={error} onClose={closeError} />}
          <form className="space-y-6" onSubmit={handleSubmit}>
            <div>
              <label
                htmlFor="email"
                className="block text-sm font-medium text-gray-700"
              >
                Email
              </label>
              <div className="mt-1 relative rounded-md shadow-sm">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <Mail className="h-5 w-5 text-gray-400" />
                </div>
                <input
                  id="email"
                  name="email"
                  type="email"
                  autoComplete="email"
                  required
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  placeholder="seu@email.com"
                />
              </div>
            </div>

            {!showForgotPassword && (
              <div>
                <label
                  htmlFor="password"
                  className="block text-sm font-medium text-gray-700"
                >
                  Senha
                </label>
                <div className="mt-1 relative rounded-md shadow-sm">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <Lock className="h-5 w-5 text-gray-400" />
                  </div>
                  <input
                    id="password"
                    name="password"
                    type={showPassword ? "text" : "password"}
                    autoComplete="current-password"
                    required
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    placeholder="••••••••"
                  />
                  <button
                    type="button"
                    onClick={() => setShowPassword(!showPassword)}
                    className="absolute inset-y-0 right-0 pr-3 flex items-center"
                  >
                    {showPassword ? (
                      <EyeOff className="h-5 w-5 text-gray-400 hover:text-gray-500" />
                    ) : (
                      <Eye className="h-5 w-5 text-gray-400 hover:text-gray-500" />
                    )}
                  </button>
                </div>
              </div>
            )}

            {!showForgotPassword && (
              <div className="flex items-center justify-between">
                <div className="flex items-center">
                  <input
                    id="remember_me"
                    name="remember_me"
                    type="checkbox"
                    className="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label
                    htmlFor="remember_me"
                    className="ml-2 block text-sm text-gray-900"
                  >
                    Lembrar-me
                  </label>
                </div>

                <div className="text-sm">
                  <button
                    type="button"
                    onClick={() => setShowForgotPassword(true)}
                    className="font-medium text-blue-600 hover:text-blue-500"
                  >
                    Esqueceu sua senha?
                  </button>
                </div>
              </div>
            )}

            <div>
              <button
                type="submit"
                className="w-full flex justify-center items-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              >
                {showForgotPassword ? "Enviar instruções" : "Entrar"}
                <ArrowRight className="ml-2 h-4 w-4" />
              </button>
            </div>
          </form>

          <div className="mt-6">
            <div className="relative">
              <div className="absolute inset-0 flex items-center">
                <div className="w-full border-t border-gray-300" />
              </div>
              <div className="relative flex justify-center text-sm">
                <span className="px-2 bg-white text-gray-500">
                  Não tem uma conta?
                </span>
              </div>
            </div>

            <div className="mt-6">
              <Link to="/register">
                <button
                  onClick={() => onNavigate && onNavigate("register")}
                  className="w-full flex justify-center py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                >
                  Cadastre-se
                </button>
              </Link>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Login;
