import React, { useState } from "react";
import InputMask from "react-input-mask";
import {
  Mail,
  Lock,
  User,
  Phone,
  MapPin,
  HardHat,
  ArrowRight,
  FileText,
  Check,
  ChevronDown,
  Eye,
  EyeOff,
  X,
  UserCircle,
  ShoppingBag,
  Youtube,
} from "lucide-react";
import { states } from "../data";
import {
  CityService,
  ProfessionalService,
  ProfessionService,
  UserService,
} from "../services";
import { StoreService } from "../services/StoreService";
import { ClientService } from "../services/ClientService";
import ErrorAlert from "../components/ErrorAlert";
import Swal from "sweetalert2";
import { useNavigate } from "react-router-dom";
import LoadingText from "../components/LoadingText";
import VideoPopup from "../components/VideoPopup";

type UserRole = "client" | "professional" | "store";

function Register() {
  const [selectedRole, setSelectedRole] = useState<UserRole>("professional");
  const [isVideoPopupOpen, setIsVideoPopupOpen] = useState(false);
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    phone: "",
    password: "",
    confirmPassword: "",
    state: "",
    city: "",
    professions: [] as string[],
    acceptTerms: false,
    photo: "",
    company: "",
  });
  const [showProfessions, setShowProfessions] = useState(false);
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [showSuccessModal, setShowSuccessModal] = useState(false);
  const [previewUrl, setPreviewUrl] = useState("");
  const [showPrivacyPolicy, setShowPrivacyPolicy] = useState(false);
  const [citiesByState, setcitiesByState] = useState([{}]);
  const [professions, setProfessions] = useState([
    { id: "", name: "", description: "" },
  ]);
  const [error, setError] = useState("");
  const [errorPass, setErrorPass] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const navigate = useNavigate();

  const roles = [
    {
      id: "professional",
      name: "Profissional",
      icon: HardHat,
      description: "Ofereço serviços profissionais",
    },
    {
      id: "client",
      name: "Cliente",
      icon: UserCircle,
      description: "Procuro profissionais e produtos",
    },
    {
      id: "store",
      name: "Lojista Parceiro",
      icon: ShoppingBag,
      description: "Vendo produtos e materiais",
    },
  ];

  React.useEffect(() => {
    const fetchData = async () => {
      const result = await ProfessionService.getProfessionsPublic();
      const json_professions = await result.data;
      if (result.status == 200) {
        setProfessions(json_professions);
      }
    };

    /* const handleClickOutside = (e) => {
      console.info(showProfessions);
      console.log(e.target);
      if (chevronRef.current && !chevronRef.current.contains(e.target)) {

        setShowProfessions(false);
      }
    };
    const handleEscapeKey = (e) => {
      if (e.key === "Escape") {
     //   setShowProfessions(false);
      }
    };

//    window.addEventListener("click", handleClickOutside, true);
 //   window.addEventListener("keydown", handleEscapeKey);

    //return () => {
    //  window.removeEventListener("click", handleClickOutside, true);
    //  window.removeEventListener("keydown", handleEscapeKey);
    //};    */

    fetchData();
  }, [showProfessions]);

  /*const handleClickOutside = (e) => {
    if (chevronRef.current && !chevronRef.current.contains(e.target)) {
      setShowProfessions(false);
    }
  };*/

  const handleChange = async (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
    const { name, value, type } = e.target;
    const checked = (e.target as HTMLInputElement).checked;

    if (name == "state") {
      const citiesByState = await CityService.citiesByStatePublic({
        uf: value,
      });

      const json_cities = await citiesByState.data;
      if (citiesByState.status == 200) {
        setcitiesByState(json_cities);
      }
    }

    setFormData((prev) => ({
      ...prev,
      [name]: type === "checkbox" ? checked : value,
      ...(name === "state" ? { city: "" } : {}),
    }));
  };

  const handlePhotoChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        setPreviewUrl(reader.result as string);
      };
      reader.readAsDataURL(file);
      const photoUrl = URL.createObjectURL(file);
      setFormData((prev) => ({ ...prev, photo: photoUrl }));
    }
  };

  const toggleProfession = (professionId: string) => {
    setFormData((prev) => ({
      ...prev,
      professions: prev.professions.includes(professionId)
        ? prev.professions.filter((id) => id !== professionId)
        : [...prev.professions, professionId],
    }));
  };

  const closeError = () => {
    setError("");
  };
  const closeErrorPass = () => {
    setErrorPass("");
  };

  const validateForm = () => {
    let isValid = true;
    setErrorPass("");
    setError("");

    // Validate password
    if (formData.password != formData.confirmPassword) {
      setErrorPass("Senhas não estão iguais!!!!");
      isValid = false;
    }

    return isValid;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    setIsLoading(true);
    e.preventDefault();

    if (validateForm()) {
      console.info("Validade" + formData.email);
      try {
        const emailReturn = await UserService.findbyemailPublic({
          email: formData.email,
        });
        if (emailReturn.status != 200) {
          console.log("status != 200");
          Swal.fire({
            position: "center",
            icon: "error",
            title: "Erro Verificação Cliente",
            showConfirmButton: false,
            timer: 1500,
          });
          setIsLoading(false);
        } else if (emailReturn.status == 200) {
          console.log("status == 200");
          const userData = await emailReturn.data;
          console.log("userData");
          console.log(userData);
          if (userData) {
            console.log("ifuserData");

            Swal.fire({
              position: "center",
              icon: "error",
              title: "Já existe uma conta <br> com esse e-mail!!",
              showConfirmButton: false,
              timer: 1500,
            });
            setIsLoading(false);
          } else {
            console.log("Passou tudo");

            let base64image = previewUrl;
            console.log(base64image);
            base64image = base64image
              .replace("data:image/png;base64,", "")
              .replace("data:image/jpg;base64,", "")
              .replace("data:image/jpeg;base64,", "");

            let postReturn: any;
            if (selectedRole == "professional") {
              try {
                postReturn = await ProfessionalService.postProfessionalPublic({
                  oid: parseInt(localStorage.getItem("id") ?? "0"),
                  Name: formData.name,
                  Email: formData.email,
                  Telephone: formData.phone,
                  Password: formData.password,
                  //LgpdAceito: "S",
                  //created_at:  "time.Date(2025, time.March, 16, 19, 41, 30, 309000000, time.Local)",
                  cep: "",
                  street: "",
                  neighborhood: "",
                  cityId: parseInt(formData.city),
                  professionIds: formData.professions,
                  image: base64image,
                });
              } catch (error) {
                Swal.fire({
                  position: "center",
                  icon: "error",
                  title: "Erro ao cadastrar profissional",
                  text: "Por favor, tente novamente mais tarde." + error,
                  showConfirmButton: true,
                });
                setIsLoading(false);
              }
            } else if (selectedRole == "client") {
              try {
                postReturn = await ClientService.postClientPublic({
                  oid: parseInt(localStorage.getItem("id") ?? "0"),
                  Name: formData.name,
                  Email: formData.email,
                  Telephone: formData.phone,
                  LgpdAceito: "S",
                  //created_at:  "time.Date(2025, time.March, 16, 19, 41, 30, 309000000, time.Local)",
                  Password: formData.password,
                  cep: "",
                  street: "",
                  neighborhood: "",
                  cityId: parseInt(formData.city),
                  image: base64image,
                });
              } catch (error) {
                Swal.fire({
                  position: "center",
                  icon: "error",
                  title: "Erro ao cadastrar cliente",
                  text: "Por favor, tente novamente mais tarde." + error,
                  showConfirmButton: true,
                });
                setIsLoading(false);
              }
            } else if (selectedRole == "store") {
              try {
                postReturn = await StoreService.postStorePublic({
                  oid: parseInt(localStorage.getItem("id") ?? "0"),
                  Name: formData.name,
                  Email: formData.email,
                  Telephone: formData.phone,
                  LgpdAceito: "S",
                  Password: formData.password,
                  //created_at:  "time.Date(2025, time.March, 16, 19, 41, 30, 309000000, time.Local)",
                  cep: "",
                  street: "",
                  neighborhood: "",
                  cityId: parseInt(formData.city),
                  image: base64image,
                });
              } catch (error) {
                setIsLoading(false);
                Swal.fire({
                  position: "center",
                  icon: "error",
                  title: "Erro ao cadastrar logista",
                  text: "Por favor, tente novamente mais tarde." + error,
                  showConfirmButton: true,
                });

              }
            }

            if (postReturn.status == 200) {
              //setShowSuccessModal(true);
              Swal.fire({
                position: "center",
                icon: "success",
                title: "Cadastro Realizado!",
                text: "Redirecionado para o login...",
                showConfirmButton: false,
                timer: 3000,
              });
              setTimeout(() => {
                setShowSuccessModal(false);
                navigate("/login");
              }, 2000);
              setIsLoading(false);
            } else {
              // if (postReturn.status == 412) {
              setIsLoading(false);
              Swal.fire({
                position: "center",
                icon: "error",
                title: "Ocorreu um erro na inclusão",
                text: postReturn.status,
                showConfirmButton: false,
                timer: 1500,
              });
            }

            setIsLoading(false);
          }
          setIsLoading(false);
        }
      } catch (error) {
        Swal.fire({
          position: "center",
          icon: "error",
          title: "Erro ao verificar e-mail",
          text: "Por favor, tente novamente mais tarde." + error,
          showConfirmButton: true,
        });
      }
      setIsLoading(false);
    }
  };

  const selectedProfessionsText =
    formData.professions.length > 0
      ? professions
          .filter((prof) => formData.professions.includes(prof.id))
          .map((prof) => prof.name)
          .join(", ")
      : "Selecione suas profissões";


  return (
    <div className="min-h-screen bg-gray-50 flex flex-col justify-center py-12 px-4 sm:px-6 lg:px-8">
      {/* loading */}
      {isLoading && <LoadingText />}

      {/* Privacy Policy Modal */}
      {showPrivacyPolicy && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
          <div className="bg-white rounded-lg shadow-xl w-full max-w-4xl max-h-[90vh] overflow-y-auto relative">
            <button
              onClick={() => setShowPrivacyPolicy(false)}
              className="absolute top-4 right-4 text-gray-400 hover:text-gray-600 transition-colors"
            >
              <X className="w-6 h-6" />
            </button>

            <div className="p-8">
              <h2 className="text-2xl font-bold text-gray-900 mb-6">
                Política de Privacidade
              </h2>

              <div className="prose prose-sm max-w-none text-gray-600 space-y-6">
                <p>
                  Sua privacidade é muito importante para nós! Esta Política de
                  Privacidade esclarece como é feito o tratamento dos seus dados
                  pessoais a partir da nossa ferramenta. Assim, prezamos pela
                  transparência entre nossa equipe e você, nosso usuário,
                  fortalecendo nossa parceria e relação de confiança. Nesse
                  sentido, gostaríamos de tranquilizá-los, pois estamos
                  totalmente adequados à Lei Geral de Proteção de Dados do
                  Brasil – LGPD (Lei n° 13.709/2018), conforme podem conferir os
                  termos abaixo estipulados.
                </p>

                <h3 className="text-xl font-bold text-gray-900">Quem somos?</h3>
                <p>
                  Mais que um site, a C + B é uma plataforma online que busca
                  reunir prestadores de serviços e clientes de uma forma rápida
                  e barata, facilitando o encontro entre profissional e sua
                  obra.
                </p>
                <p>
                  O nosso contato é realizado por meio do e-mail:
                  atendimento@construirmaisbarato.com.br
                </p>
                <p>
                  Nós temos também um responsável pela proteção de dados,
                  portanto, quaisquer dúvidas ou solicitações sobre o uso de
                  seus dados pessoais devem ser encaminhadas para o nosso
                  encarregado de dados:
                </p>
                <p>
                  Jairo Assis lgpd@construirmaisbarato.com.br (14) 98835-0791
                </p>

                <h3 className="text-xl font-bold text-gray-900">
                  COMO USAMOS OS SEUS DADOS:
                </h3>
                <p>
                  Nosso site pode ser utilizado para áreas como construção,
                  pintura, elétrica e reparos hidráulicos. Podem oferecer
                  serviços em nosso site profissionais com CNPJ, MEI ou
                  autônomos. Os usuários (cliente final) poderão ser pessoas
                  jurídicas ou físicas. Ao fazer o cadastro em nossa plataforma
                  (site/aplicativo), coletaremos algumas informações que serão
                  fornecidas exclusivamente pelo usuário. Todavia, esclarecemos
                  que essas informações são basicamente cadastrais, como as
                  seguintes informações: nome, e-mail, CPF, endereço e telefone.
                  Quando solicitado o endereço, este se refere ao local da
                  prestação de serviço a ser realizado. Menores de idade não
                  poderão utilizar nossos serviços. Ressaltamos que a exclusão
                  dos dados de nossa ferramenta é perfeitamente possível.
                </p>
                <p>
                  Usamos essas informações exclusivamente para a funcionalidade
                  de nosso sistema. Também podemos lhe enviar e-mails. Faremos
                  isso com base em nosso interesse legítimo em fornecer
                  informações precisas e um serviço de qualidade. Caso não
                  queira receber nossos e-mails, basta realizar o
                  descadastramento em nosso site.
                </p>
                <p>
                  Suas informações são armazenadas em nosso servidor e será
                  tratada apenas em decorrência da nossa prestação de serviços.
                  Não comercial
                </p>

                <h3 className="text-xl font-bold text-gray-900">COOKIES</h3>
                <p>
                  Quando você usa nosso site para navegar em nossos serviços,
                  vários cookies são usados por nós e por terceiros para
                  permitir que o site funcione, para coletar informações úteis
                  sobre os visitantes, ajudando a tornar sua experiência de
                  usuário melhor.
                </p>
                <p>
                  Alguns dos cookies que usamos são estritamente necessários
                  para o funcionamento do nosso site, e não pedimos o seu
                  consentimento para colocá-los no seu computador. No entanto,
                  para os cookies que são úteis, mas não estritamente
                  necessários, pediremos sempre o seu consentimento antes de os
                  colocar.
                </p>

                <h3 className="text-xl font-bold text-gray-900">
                  Do Compartilhamento
                </h3>
                <p>
                  Seus dados são armazenados em nosso banco de dados, mas não
                  serão compartilhados com terceiros, a não ser nos casos
                  previstos em Lei.
                </p>

                <h3 className="text-xl font-bold text-gray-900">
                  Dos Serviços
                </h3>
                <p>
                  A função da nossa plataforma é facilitar o encontro entre
                  profissionais e clientes, meramente informativo e consultivo,
                  no estilo "páginas amarelas" das listas telefônicas. Toda e
                  qualquer negociação realizada entre as partes é de
                  responsabilidade delas. Nosso site NÃO se responsabiliza por
                  defeitos na prestação dos serviços contratados pelo usuário.
                </p>

                <h3 className="text-xl font-bold text-gray-900">
                  Do armazenamento e segurança
                </h3>
                <p>
                  Utilizamos técnicas e softwares seguros e renomados para o
                  armazenamento de todas as informações que transitam pelo site.
                  Assim, garantimos a utilização de medidas técnicas e
                  administrativas aptas a proteger os dados pessoais de acessos
                  não autorizados e de situações acidentais ou ilícitas de
                  destruição, perda, alteração, comunicação ou difusão de seus
                  dados.
                </p>

                <h3 className="text-xl font-bold text-gray-900">
                  Seus direitos como titular de dados
                </h3>
                <p>
                  Por lei, qualquer indivíduo poderá nos perguntar quais são as
                  informações que temos sobre ele em nosso banco de dados, além
                  de ser garantido o direito de correção, se as informações
                  estiverem imprecisas, por meio do e-mail
                  lgpd@construirmaisbarato.com.br. Se solicitarmos o seu
                  consentimento para processar seus dados, você poderá retirar
                  esse consentimento a qualquer momento, bem como solicitar a
                  exclusão de dados. Caso queira enviar uma solicitação sobre a
                  utilização de seus dados pessoais (informações, correções e
                  exclusão), use o endereço eletrônico fornecido nesta política.
                </p>

                <h3 className="text-xl font-bold text-gray-900">
                  Atualizações para esta política de privacidade
                </h3>
                <p>
                  Revisamos regularmente e, se apropriado, atualizaremos esta
                  política de privacidade de tempos em tempos, e conforme nossos
                  serviços e uso de dados sejam alterados. Se, eventualmente,
                  usarmos seus dados pessoais de uma forma que não identificada
                  ou descrita anteriormente, entraremos em contato para fornecer
                  informações sobre isso e, se necessário, solicitar o seu
                  consentimento.
                </p>
              </div>

              <div className="mt-8 flex justify-end">
                <button
                  onClick={() => setShowPrivacyPolicy(false)}
                  className="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
                >
                  Fechar
                </button>
              </div>
            </div>
          </div>
        </div>
      )}

      {/* Success Modal */}
      {showSuccessModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-8 max-w-sm w-full mx-4 transform transition-all">
            <div className="flex flex-col items-center">
              <div className="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mb-4">
                <Check className="w-8 h-8 text-green-600" />
              </div>
              <h3 className="text-xl font-semibold text-gray-900 mb-2">
                Cadastro Realizado!
              </h3>
              <p className="text-gray-600 text-center">
                Redirecionando para o login...
              </p>
            </div>
          </div>
        </div>
      )}

      <div className="sm:mx-auto sm:w-full sm:max-w-2xl">
        <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
          Cadastre-se
        </h2>
        <p className="mt-2 text-center text-sm text-gray-600">
          Preencha seus dados para criar sua conta
        </p>
      </div>

      <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-2xl">
        <div className="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
          {error && <ErrorAlert message={error} onClose={closeError} />}
          {/* Role Selection */}
          <div className="mb-8">
            <label className="block text-sm font-medium text-gray-700 mb-4">
              Como você deseja se cadastrar?
            </label>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              {roles.map((role) => {
                const Icon = role.icon;
                return (
                  <button
                    key={role.id}
                    type="button"
                    onClick={() => setSelectedRole(role.id as UserRole)}
                    className={`flex flex-col items-center p-4 rounded-lg border-2 transition-colors ${
                      selectedRole === role.id
                        ? "border-blue-600 bg-blue-50"
                        : "border-gray-200 hover:border-blue-300"
                    }`}
                  >
                    <Icon
                      className={`w-8 h-8 mb-2 ${
                        selectedRole === role.id
                          ? "text-blue-600"
                          : "text-gray-400"
                      }`}
                    />
                    <span
                      className={`font-medium ${
                        selectedRole === role.id
                          ? "text-blue-600"
                          : "text-gray-900"
                      }`}
                    >
                      {role.name}
                    </span>
                    <span className="text-xs text-gray-500 text-center mt-1">
                      {role.description}
                    </span>
                  </button>
                );
              })}
            </div>
          </div>

          <form className="space-y-6" onSubmit={handleSubmit}>
            {/* Foto */}
            {/*
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Foto de Perfil
              </label>
              <div className="flex items-center justify-center">
                <div className="relative">
                  <div
                    className={`w-32 h-32 rounded-full overflow-hidden border-2 border-gray-300 flex items-center justify-center bg-gray-50 ${
                      previewUrl ? "" : "border-dashed"
                    }`}
                  >
                    {previewUrl ? (
                      <img
                        src={previewUrl}
                        alt="Preview"
                        className="w-full h-full object-cover"
                      />
                    ) : (
                      <Camera className="w-8 h-8 text-gray-400" />
                    )}
                  </div>
                  <label
                    htmlFor="photo-upload"
                    className="absolute bottom-0 right-0 bg-blue-600 text-white p-2 rounded-full cursor-pointer hover:bg-blue-700 transition-colors"
                  >
                    <Upload className="w-4 h-4" />
                  </label>
                  <input
                    id="photo-upload"
                    name="photo"
                    type="file"
                    accept="image/*"
                    onChange={handlePhotoChange}
                    className="hidden"
                  />
                  <button
                    type="button"
                    onClick={() => {
                      setPreviewUrl("");
                      setFormData((prev) => ({ ...prev, photo: "" }));
                    }}
                    className="p-2 bg-red-600 text-white rounded-full hover:bg-red-700 transition-colors"
                  >
                    <X className="w-5 h-5" />
                  </button>
                </div>
              </div>
            </div>
            */}
            {/* Nome */}
            <div>
              <label
                htmlFor="name"
                className="block text-sm font-medium text-gray-700"
              >
                Nome Completo
              </label>
              <div className="mt-1 relative rounded-md shadow-sm">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <User className="h-5 w-5 text-gray-400" />
                </div>
                <input
                  id="name"
                  name="name"
                  type="text"
                  required
                  value={formData.name}
                  onChange={handleChange}
                  className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  placeholder="João da Silva"
                />
              </div>
            </div>

            {/* Empresa - Only show not client role */}
            {selectedRole !== "client" && (
              <div>
                <label
                  htmlFor="name"
                  className="block text-sm font-medium text-gray-700"
                >
                  Empresa
                </label>
                <div className="mt-1 relative rounded-md shadow-sm">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <User className="h-5 w-5 text-gray-400" />
                  </div>
                  <input
                    id="company"
                    name="company"
                    type="text"
                    required
                    value={formData.company}
                    onChange={handleChange}
                    minLength={10}
                    className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    placeholder="Construir LTDA"
                  />
                </div>
              </div>
            )}

            {/* Email */}
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
                  required
                  value={formData.email}
                  onChange={handleChange}
                  className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  placeholder="seu@email.com"
                />
              </div>
            </div>

            {/* Telefone */}
            <div>
              <label
                htmlFor="phone"
                className="block text-sm font-medium text-gray-700"
              >
                Telefone Comercial
              </label>
              <div className="mt-1 relative rounded-md shadow-sm">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <Phone className="h-5 w-5 text-gray-400" />
                </div>
                <InputMask
                  mask="(99) 99999-9999"
                  id="phone"
                  name="phone"
                  type="tel"
                  required
                  value={formData.phone}
                  onChange={handleChange}
                  className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  placeholder="(11) 99999-9999"
                />
              </div>
            </div>

            {/* Senha */}
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
                  required
                  value={formData.password}
                  onChange={handleChange}
                  className="appearance-none block w-full pl-10 pr-10 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
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
              {errorPass && (
                <ErrorAlert message={errorPass} onClose={closeErrorPass} />
              )}
            </div>

            {/* Confirmar Senha */}
            <div>
              <label
                htmlFor="confirmPassword"
                className="block text-sm font-medium text-gray-700"
              >
                Confirmar Senha
              </label>
              <div className="mt-1 relative rounded-md shadow-sm">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <Lock className="h-5 w-5 text-gray-400" />
                </div>
                <input
                  id="confirmPassword"
                  name="confirmPassword"
                  type={showConfirmPassword ? "text" : "password"}
                  required
                  value={formData.confirmPassword}
                  onChange={handleChange}
                  className="appearance-none block w-full pl-10 pr-10 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  placeholder="••••••••"
                />
                <button
                  type="button"
                  onClick={() => setShowConfirmPassword(!showConfirmPassword)}
                  className="absolute inset-y-0 right-0 pr-3 flex items-center"
                >
                  {showConfirmPassword ? (
                    <EyeOff className="h-5 w-5 text-gray-400 hover:text-gray-500" />
                  ) : (
                    <Eye className="h-5 w-5 text-gray-400 hover:text-gray-500" />
                  )}
                </button>
              </div>
            </div>

            {/* Estado e Cidade */}
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label
                  htmlFor="state"
                  className="block text-sm font-medium text-gray-700"
                >
                  Estado
                </label>
                <div className="mt-1 relative rounded-md shadow-sm">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <MapPin className="h-5 w-5 text-gray-400" />
                  </div>
                  <select
                    id="state"
                    name="state"
                    required
                    value={formData.state}
                    onChange={handleChange}
                    className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  >
                    <option value="">Selecione o estado</option>
                    {states.map((state) => (
                      <option key={state.id} value={state.id}>
                        {state.name}
                      </option>
                    ))}
                  </select>
                </div>
              </div>

              <div>
                <label
                  htmlFor="city"
                  className="block text-sm font-medium text-gray-700"
                >
                  Cidade
                </label>
                <div className="mt-1 relative rounded-md shadow-sm">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <MapPin className="h-5 w-5 text-gray-400" />
                  </div>
                  <select
                    id="city"
                    name="city"
                    required
                    value={formData.city}
                    onChange={handleChange}
                    disabled={!formData.state}
                    className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100"
                  >
                    <option value="">Selecione a cidade</option>
                    {formData.state &&
                      citiesByState.map((city: any) => (
                        <option key={city.id} value={city.id}>
                          {city.name}
                        </option>
                      ))}
                  </select>
                </div>
              </div>
            </div>

            {/* Profissões - Only show for professional role */}
            {selectedRole === "professional" && (
              <div className="relative">
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Profissões
                </label>
                <button
                  type="button"
                  onClick={() => setShowProfessions(!showProfessions)}
                  className="relative w-full bg-white border border-gray-300 rounded-md shadow-sm pl-10 pr-10 py-2 text-left cursor-pointer focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500"
                >
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <HardHat className="h-5 w-5 text-gray-400" />
                  </div>
                  <span className="block truncate">
                    {selectedProfessionsText}
                  </span>
                  <span className="absolute inset-y-0 right-0 flex items-center pr-2">
                    <ChevronDown
                      className={`h-5 w-5 text-gray-400 transition-transform ${
                        showProfessions ? "transform rotate-180" : ""
                      }`}
                    />
                  </span>
                </button>

                {showProfessions && (
                  <div className="absolute z-10 mt-1 w-full bg-white shadow-lg max-h-60 rounded-md py-1 text-base overflow-auto focus:outline-none sm:text-sm">
                    {professions != null &&
                      professions.map((profession) => (
                        <div
                          key={profession.id}
                          className="relative cursor-pointer select-none py-2 pl-10 pr-4 hover:bg-blue-50"
                          onClick={() => toggleProfession(profession.id)}
                        >
                          <span
                            className={`block truncate ${
                              formData.professions.includes(profession.id)
                                ? "font-medium text-blue-600"
                                : "font-normal"
                            }`}
                          >
                            {profession.name}
                          </span>
                          {formData.professions.includes(profession.id) && (
                            <span className="absolute inset-y-0 left-0 flex items-center pl-3 text-blue-600">
                              <Check className="h-5 w-5" />
                            </span>
                          )}
                        </div>
                      ))}
                  </div>
                )}
              </div>
            )}

            {/* Termos de Uso */}
            <div className="flex items-center">
              <input
                id="acceptTerms"
                name="acceptTerms"
                type="checkbox"
                required
                checked={formData.acceptTerms}
                onChange={handleChange}
                className="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label
                htmlFor="acceptTerms"
                className="ml-2 block text-sm text-gray-900"
              >
                Li e aceito a{" "}
                <button
                  type="button"
                  onClick={() => setShowPrivacyPolicy(true)}
                  className="text-blue-600 hover:text-blue-500 font-medium"
                >
                  política de privacidade
                </button>
              </label>
              <FileText className="ml-2 h-4 w-4 text-gray-400" />
            </div>

            {/* Botão de Cadastro */}
            <div>
              <button
                type="submit"
                disabled={
                  !formData.acceptTerms ||
                  (selectedRole === "professional" &&
                    formData.professions.length === 0)
                }
                className="w-full flex justify-center items-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:bg-gray-400 disabled:cursor-not-allowed"
              >
                Cadastrar
                <ArrowRight className="ml-2 h-4 w-4" />
              </button>
            </div>
          </form>
        </div>
      </div>
      {/* Floating YouTube Button */}
      <button
        onClick={() => setIsVideoPopupOpen(true)}
        className="fixed bottom-24 right-1 z-40 flex flex-col items-center"
      >
        <span className="text-xs font-medium text-white bg-red-600 px-3 py-1 rounded-full mb-2 shadow-md">
          COMO SE <br></br>
          CADASTRAR
        </span>
        <div className="w-16 h-16 bg-red-600 rounded-full shadow-lg hover:bg-red-700 transition-colors flex items-center justify-center">
          <Youtube className="w-8 h-8 text-white" />
        </div>
      </button>

      <VideoPopup
        isOpen={isVideoPopupOpen}
        onClose={() => setIsVideoPopupOpen(false)}
        url="https://www.youtube.com/embed/a5Orf5iu9EQ"
      />
    </div>
  );
}

export default Register;
