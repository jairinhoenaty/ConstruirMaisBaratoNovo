import React, { useState, useEffect, useRef } from "react";
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
  Camera,
  Upload,
  Shield,
  Building,
  ExternalLink,
} from "lucide-react";
import { states } from "../data";
import {
  CityService,
  ProfessionalService,
  ProfessionService,
  UserService,
  LoginService,
  ProductCategoryService,
} from "../services";
import { StoreService } from "../services/StoreService";
import { ClientService } from "../services/ClientService";
import { PlanService } from "../services/PlanService";
import { CheckoutState, Payer } from "../interfaces";
import { Plan } from "../interfaces/IPlan";
import ErrorAlert from "../components/ErrorAlert";
import Swal from "sweetalert2";
import { useNavigate, useLocation, Form } from "react-router-dom";
import LoadingText from "../components/LoadingText";
import VideoPopup from "../components/VideoPopup";
import { ICategoryProduct } from "../interfaces/ICategoryProduct";

type UserRole = "client" | "professional" | "store";

function Register() {
  const [categoryProducts, setCategoryProducts] = useState<ICategoryProduct[]>(
    []
  );
  const [selectedSubcategories, setSelectedSubcategories] = useState<number[]>(
    []
  );
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
    // Campos Premium
    dateOfBirth: "",
    experience: "",
    codeVerification: "",
    meiCnpj: "",
    negativeCertificateNumber: "",
    isPremium: false,
    isPremiumStore: false,
  });
  const [showProfessions, setShowProfessions] = useState(false);
  const [selectedCategoryProduct, setSelectedCategoryProduct] =
    useState<string>("");
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [showSuccessModal, setShowSuccessModal] = useState(false);
  const [showPremiumModal, setShowPremiumModal] = useState(false);
  const [showStoreModal, setShowStoreModal] = useState(false);
  const [premiumStore, setPremiumStore] = useState(false);
  const [premiumForm, setPremiumForm] = useState(false);
  const [clickedCertificateButton, setclickedCertificateButton] =
    useState(false);
  const [previewUrl, setPreviewUrl] = useState("");
  const [showPrivacyPolicy, setShowPrivacyPolicy] = useState(false);
  const [citiesByState, setcitiesByState] = useState([{}]);
  const [professions, setProfessions] = useState([
    { id: "", name: "", description: "" },
  ]);
  const [error, setError] = useState("");
  const [errorPass, setErrorPass] = useState("");
  const [errorAge, setErrorAge] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [codigoEnviado, setCodigoEnviado] = useState(false);
  const [loadingSMS, setLoadingSMS] = useState(false);
  const [choose, setChoose] = useState(false);
  const [isUpgradeMode, setIsUpgradeMode] = useState(false);
  const [professionalPlan, setProfessionalPlan] = useState<Plan | null>(null);
  const [storePlan, setStorePlan] = useState<Plan | null>(null);

  const navigate = useNavigate();
  const location = useLocation();
  const formRef = useRef<HTMLFormElement>(null);

  const roles = [
    {
      id: "professional",
      name: "Profissional",
      icon: HardHat,
      description: "Ofereço serviços profissionais",
    },

    // {
    //   id: "asdasdasdas",
    //   name: "Balcão de Vagas",
    //   icon: UserCircle,
    //   description: "descrição de balcão de vagas",
    // },
    {
      id: "store",
      name: "Lojista Parceiro",
      icon: ShoppingBag,
      description: "Vendo produtos e materiais",
    },
  ];

  React.useEffect(() => {
    const fetchData = async () => {
      // Buscar profissões
      const result = await ProfessionService.getProfessionsPublic();
      const json_professions = await result.data;
      if (result.status == 200) {
        setProfessions(json_professions);
      }
      // Buscar Categorias de produtos
      handleSearchCategoriesProduct();

      // Buscar planos premium
      try {
        const professionalPlanResult = await PlanService.getPlanByUserType(
          "professional"
        );
        if (professionalPlanResult.data) {
          setProfessionalPlan(professionalPlanResult.data);
        }
      } catch (error) {
        console.error("Erro ao buscar plano profissional:", error);
      }

      try {
        const storePlanResult = await PlanService.getPlanByUserType("store");
        if (storePlanResult.data) {
          setStorePlan(storePlanResult.data);
        }
      } catch (error) {
        console.error("Erro ao buscar plano lojista:", error);
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

  // Detectar modo de upgrade e pré-popular dados do usuário logado
  React.useEffect(() => {
    const upgradeState = location.state as { upgradeToPremium?: boolean };

    if (upgradeState?.upgradeToPremium) {
      console.log("Modo upgrade detectado!");
      setIsUpgradeMode(true);

      // Buscar dados do usuário logado do localStorage
      const userName = localStorage.getItem("name") || "";
      const userEmail = localStorage.getItem("email") || "";
      const userId = localStorage.getItem("id") || "";

      // Pre-popular formulário com dados existentes
      setFormData((prev) => ({
        ...prev,
        name: userName,
        email: userEmail,
      }));

      // Automaticamente mostrar modal premium na primeira vez
      setChoose(false);
      setShowPremiumModal(true);
    }
  }, [location]);

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

  const handleGerarCertidao = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
    e.preventDefault();
    window.open("https://servicos.pf.gov.br/epol-sinic-publico/", "_blank");
    setclickedCertificateButton(true);
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
  const closeErrorAge = () => {
    setErrorAge("");
  };

  // Função para calcular idade
  const calculateAge = (birthDate: string): number => {
    const today = new Date();
    const birth = new Date(birthDate);
    let age = today.getFullYear() - birth.getFullYear();
    const monthDiff = today.getMonth() - birth.getMonth();

    if (
      monthDiff < 0 ||
      (monthDiff === 0 && today.getDate() < birth.getDate())
    ) {
      age--;
    }

    return age;
  };

  // Função para validar CPF
  const validateCPF = (cpf: string): boolean => {
    const cleanCPF = cpf.replace(/[^\d]/g, "");

    if (cleanCPF.length !== 11) return false;
    if (/^(\d)\1{10}$/.test(cleanCPF)) return false;

    let sum = 0;
    for (let i = 0; i < 9; i++) {
      sum += parseInt(cleanCPF.charAt(i)) * (10 - i);
    }
    let checkDigit = 11 - (sum % 11);
    if (checkDigit === 10 || checkDigit === 11) checkDigit = 0;
    if (checkDigit !== parseInt(cleanCPF.charAt(9))) return false;

    sum = 0;
    for (let i = 0; i < 10; i++) {
      sum += parseInt(cleanCPF.charAt(i)) * (11 - i);
    }
    checkDigit = 11 - (sum % 11);
    if (checkDigit === 10 || checkDigit === 11) checkDigit = 0;
    if (checkDigit !== parseInt(cleanCPF.charAt(10))) return false;

    return true;
  };

  // Função para validar CNPJ/MEI
  const validateCNPJ = (cnpj: string): boolean => {
    const cleanCNPJ = cnpj.replace(/[^\d]/g, "");

    if (cleanCNPJ.length !== 14) return false;
    if (/^(\d)\1{13}$/.test(cleanCNPJ)) return false;

    // Validação dos dígitos verificadores do CNPJ
    let length = cleanCNPJ.length - 2;
    let numbers = cleanCNPJ.substring(0, length);
    let digits = cleanCNPJ.substring(length);
    let sum = 0;
    let pos = length - 7;

    for (let i = length; i >= 1; i--) {
      sum += parseInt(numbers.charAt(length - i)) * pos--;
      if (pos < 2) pos = 9;
    }

    let result = sum % 11 < 2 ? 0 : 11 - (sum % 11);
    if (result !== parseInt(digits.charAt(0))) return false;

    length = length + 1;
    numbers = cleanCNPJ.substring(0, length);
    sum = 0;
    pos = length - 7;

    for (let i = length; i >= 1; i--) {
      sum += parseInt(numbers.charAt(length - i)) * pos--;
      if (pos < 2) pos = 9;
    }

    result = sum % 11 < 2 ? 0 : 11 - (sum % 11);
    return result === parseInt(digits.charAt(1));
  };

  const handleCloseStoreModal = async () => {
    setShowStoreModal(false);
    setChoose(true);

    setFormData((prev) => ({
      ...prev,
      isPremiumStore: false,
    }));
    setCodigoEnviado(false);

    const registrationSuccess = await processBasicRegistration();

    if (registrationSuccess) {
      LoginService.login({
        email: formData.email,
        password: formData.password,
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

            if (response.data.user.profile === "admin") {
              navigate("/dashboard");
            } else {
              navigate("/professional-panel");
            }
          } else {
            setError("Login Inválido!!");
          }
        })
        .catch((err) => {
          setError("Login Inválido!!");
          console.error("ops! ocorreu um erro" + err);
        });
    }
  };

  const handleSearchCategoriesProduct = async () => {
    // localStorage.setItem("search_city", selectedCity);
    try {
      const response = await ProductCategoryService.findTopLevelCategory();
      setCategoryProducts(response.data || []);
    } catch (error) {
      console.error("Erro ao buscar categorias:", error);
    }
  };

  const subcategories = selectedCategoryProduct
    ? categoryProducts.find(
        (cat) => cat.id === parseInt(selectedCategoryProduct)
      )?.children || []
    : [];

  const toggleSubcategory = (subcategoryId: number) => {
    setSelectedSubcategories((prev) => {
      if (prev.includes(subcategoryId)) {
        return prev.filter((id) => id !== subcategoryId);
      } else {
        return [...prev, subcategoryId];
      }
    });
  };

  const toggleAllSubcategories = () => {
    if (subcategories.length > 0) {
        setSelectedSubcategories(subcategories.map((sub) => sub.id));
    }
  };

  // useEffect(() => {
  //   if (subcategories.length > 0) {
  //     toggleAllSubcategories();
  //   }
  // }, [subcategories]);
  
  // useEffect(() => {
  //   setSelectedSubcategories([]); // limpa seleção
  // }, [selectedCategoryProduct]);
  

  const handleClosePremiumModal = async () => {
    setShowPremiumModal(false);
    setChoose(true);
    setPremiumForm(false);

    setFormData((prev) => ({
      ...prev,
      dateOfBirth: "",
      experience: "",
      codeVerification: "",
      meiCnpj: "",
      negativeCertificate: "",
      isPremium: false,
    }));
    setCodigoEnviado(false);

    const registrationSuccess = await processBasicRegistration();

    if (registrationSuccess) {
      LoginService.login({
        email: formData.email,
        password: formData.password,
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
    }
  };

  const enviarCodigoSMS = async () => {
    if (!formData.phone) {
      setError("Preencha o telefone antes de solicitar o código");
      return;
    }

    setLoadingSMS(true);
    try {
      // Simular envio de SMS (substituir pela API real)
      await new Promise((resolve) => setTimeout(resolve, 2000));
      setCodigoEnviado(true);
      setError("");
      // Aqui você faria a chamada real para sua API de SMS
      console.log("Código enviado para:", formData.phone);
    } catch (error) {
      setError("Erro ao enviar código SMS. Tente novamente.");
    } finally {
      setLoadingSMS(false);
    }
  };

  const validateBasicForm = () => {
    let isValid = true;
    setErrorPass("");
    setError("");

    // Validate password
    if (formData.password != formData.confirmPassword) {
      setErrorPass("Senhas não estão iguais!!!!");
      isValid = false;
    }

    // Validações básicas apenas (sem campos premium)
    return isValid;
  };

  const validateForm = () => {
    let isValid = true;
    setErrorPass("");
    setError("");
    setErrorAge("");

    // Validate password - pular em modo upgrade
    if (!isUpgradeMode) {
      if (formData.password != formData.confirmPassword) {
        setErrorPass("Senhas não coincidem");
        isValid = false;
      }
    }

    // Validações específicas para premium
    if (premiumForm) {
      // Validar telefone em modo upgrade (obrigatório para SMS)
      if (isUpgradeMode && !formData.phone) {
        setError("Telefone é obrigatório para receber o código SMS");
        isValid = false;
      }

      // Validar idade (maior de 18 anos)
      if (formData.dateOfBirth) {
        const age = calculateAge(formData.dateOfBirth);
        if (age < 18) {
          setErrorAge("Você deve ser maior de 18 anos para o cadastro premium");
          isValid = false;
        }
      }

      // Validar código de verificação SMS
      if (!formData.codeVerification) {
        setError("Código de verificação SMS é obrigatório");
        isValid = false;
      } else if (
        formData.codeVerification.length < 4 ||
        formData.codeVerification.length > 6
      ) {
        setError("Código deve ter entre 4 e 6 dígitos");
        isValid = false;
      }

      // Validar CNPJ/MEI (opcional, mas se preenchido deve ser válido)
      if (formData.meiCnpj && !validateCNPJ(formData.meiCnpj)) {
        setError("CNPJ/MEI inválido");
        isValid = false;
      }

      // Validar foto (obrigatória para premium)
      if (!formData.photo) {
        setError("Foto é obrigatória para cadastro premium");
        isValid = false;
      }

      // Validar certificado negativo (obrigatória para premium)
      if (!formData.negativeCertificateNumber) {
        setError("Certificado negativo é obrigatório");
        isValid = false;
      }
    }

    return isValid;
  };

  const processBasicRegistration = async () => {
    // Validar apenas campos básicos
    if (!validateBasicForm()) {
      return false;
    }

    // Configurar como usuário básico
    setFormData((prev) => ({ ...prev, isPremium: false }));

    setIsLoading(true);

    try {
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
          return false;
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
            return false;
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
                const professionalData = {
                  oid: parseInt(localStorage.getItem("id") ?? "0"),
                  Name: formData.name,
                  Email: formData.email,
                  Telephone: formData.phone,
                  Password: formData.password,
                  cep: "",
                  street: "",
                  neighborhood: "",
                  cityId: parseInt(formData.city),
                  professionIds: formData.professions,
                  image: base64image,
                  isPremium: false,
                };

                console.log("Dados do profissional:", professionalData);
                postReturn = await ProfessionalService.postProfessionalPublic(
                  professionalData
                );
              } catch (error) {
                Swal.fire({
                  position: "center",
                  icon: "error",
                  title: "Erro ao cadastrar profissional",
                  text: "Por favor, tente novamente mais tarde." + error,
                  showConfirmButton: true,
                });
                return false;
              }
            } else if (selectedRole == "client") {
              try {
                postReturn = await ClientService.postClientPublic({
                  oid: parseInt(localStorage.getItem("id") ?? "0"),
                  Name: formData.name,
                  Email: formData.email,
                  Telephone: formData.phone,
                  LgpdAceito: "S",
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
                return false;
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
                  cep: "",
                  street: "",
                  neighborhood: "",
                  cityId: parseInt(formData.city),
                  image: base64image,
                  categoryProductID: parseInt(selectedCategoryProduct),
                  subCategories: selectedSubcategories,
                  // isPremiumStore: formData.isPremiumStore,
                });
              } catch (error) {
                Swal.fire({
                  position: "center",
                  icon: "error",
                  title: "Erro ao cadastrar logista",
                  text: "Por favor, tente novamente mais tarde." + error,
                  showConfirmButton: true,
                });
                return false;
              }
            }

            if (postReturn.status == 200) {
              setShowSuccessModal(true);
              Swal.fire({
                position: "center",
                icon: "success",
                title: "Cadastro Realizado!",
                text: "Logando...",
                showConfirmButton: false,
                timer: 3000,
              });
              return true; // Cadastro bem-sucedido
            } else {
              Swal.fire({
                position: "center",
                icon: "error",
                title: "Ocorreu um erro na inclusão",
                text: `Erro: ${postReturn.status}`,
                showConfirmButton: false,
                timer: 1500,
              });
              return false; // Cadastro falhou
            }
          }
        } else {
          console.log("status != 200");
          Swal.fire({
            position: "center",
            icon: "error",
            title: "Erro Verificação Cliente",
            showConfirmButton: false,
            timer: 1500,
          });
          return false;
        }
      } catch (error) {
        console.error("Erro no cadastro:", error);
        Swal.fire({
          position: "center",
          icon: "error",
          title: "Erro ao verificar e-mail",
          text: "Por favor, tente novamente mais tarde.",
          showConfirmButton: true,
        });
        return;
      } finally {
        setIsLoading(false);
      }
    } catch (error) {
      console.error("Erro geral:", error);
      Swal.fire({
        position: "center",
        icon: "error",
        title: "Erro interno",
        text: "Por favor, tente novamente mais tarde.",
        showConfirmButton: true,
      });
      setIsLoading(false);
      return false; // General error
    }
  };

  const processFormSubmission = async () => {
    if (!choose && selectedRole == "professional") {
      setShowPremiumModal(true);
      return false;
    }
    if (!choose && selectedRole == "store") {
      setShowStoreModal(true);
      return false;
    }
    if (!validateForm()) {
      return false;
    }

    return true;
  };

  useEffect(() => {
    if (selectedCategoryProduct) {
      toggleAllSubcategories();
    }
  }, [selectedCategoryProduct]);
  

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // Se estiver em modo upgrade, pular validação de modal e ir direto pro checkout
    if (isUpgradeMode && premiumForm) {
      // Validar apenas campos premium
      if (!validateForm()) {
        return;
      }

      setIsLoading(true);

      // Buscar dados do usuário logado
      const userId = localStorage.getItem("id") || "";
      const userName = localStorage.getItem("name") || "";
      const userEmail = localStorage.getItem("email") || "";

      // TODO: Aqui deveria chamar API para atualizar dados premium do usuário
      // Por enquanto, vamos apenas redirecionar para o checkout

      setTimeout(() => {
        const payer: Payer = {
          first_name: userName.split(" ")[0],
          last_name: userName.split(" ").slice(1).join(" "),
          email: userEmail,
          identification: {
            type: "CPF",
            number: formData.meiCnpj || "",
          },
          address: {
            zip_code: "",
            street_name: "",
            street_number: "",
            neighborhood: "",
            city: formData.city || "",
            federal_unit: formData.state || "",
          },
        };

        const checkoutState: CheckoutState = {
          userId: parseInt(userId),
          userName: userName,
          userEmail: userEmail,
          planId: professionalPlan?.id || 0,
          userType: "professional",
          payer: payer,
        };

        Swal.fire({
          position: "center",
          icon: "success",
          title: "🚀 Upgrade para Premium!",
          text: "Redirecionando para o pagamento...",
          showConfirmButton: false,
          timer: 2000,
        });

        setTimeout(() => {
          navigate("/checkout", { state: checkoutState });
        }, 2000);
      }, 500);

      setIsLoading(false);
      return;
    }

    const shouldContinue = await processFormSubmission();
    if (!shouldContinue) {
      return;
    }

    setIsLoading(true);

    try {
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
                // Preparar dados base
                const professionalData = {
                  oid: parseInt(localStorage.getItem("id") ?? "0"),
                  Name: formData.name,
                  Email: formData.email,
                  Telephone: formData.phone,
                  Password: formData.password,
                  cep: "",
                  street: "",
                  neighborhood: "",
                  cityId: parseInt(formData.city),
                  professionIds: formData.professions,
                  // Campos Premium
                  ...(premiumForm && {
                    image: base64image,
                    dateOfBirth: formData.dateOfBirth,
                    experience: formData.experience,
                    meiCnpj: formData.meiCnpj,
                    telefoneVerificado: !!formData.codeVerification,
                    negativeCertificateNumber: parseInt(
                      formData.negativeCertificateNumber
                    ),
                  }),
                };

                console.log("Dados do profissional:", professionalData);
                postReturn = await ProfessionalService.postProfessionalPublic(
                  professionalData
                );
              } catch (error) {
                Swal.fire({
                  position: "center",
                  icon: "error",
                  title: "Erro ao cadastrar profissional",
                  text: "Por favor, tente novamente mais tarde." + error,
                  showConfirmButton: true,
                });
                return false;
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
                return false;
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
                  categoryProductID: parseInt(selectedCategoryProduct),
                  subCategories: selectedSubcategories,
                  // isPremiumStore: formData.isPremiumStore
                });
              } catch (error) {
                Swal.fire({
                  position: "center",
                  icon: "error",
                  title: "Erro ao cadastrar logista",
                  text: "Por favor, tente novamente mais tarde." + error,
                  showConfirmButton: true,
                });
                return false;
              }
            }

            if (postReturn.status == 200) {
              const professionalId = postReturn.data.oid;
              Swal.fire({
                position: "center",
                icon: "success",
                title: "Cadastro Realizado!",
                text: premiumForm ? "Redirecionando para o pagamento..." : "",
                showConfirmButton: false,
                timer: 3000,
              });
              if (premiumForm || formData.isPremiumStore) {
                setTimeout(() => {
                  const payer: Payer = {
                    first_name: formData.name.split(" ")[0],
                    last_name: formData.name.split(" ").slice(1).join(" "),
                    email: formData.email,
                    identification: {
                      type: "CPF",
                      number: formData.meiCnpj || "",
                    },
                    address: {
                      zip_code: "",
                      street_name: "",
                      street_number: "",
                      neighborhood: "",
                      city: formData.city,
                      federal_unit: formData.state,
                    },
                  };

                  // Selecionar plano e userType baseado no selectedRole
                  const currentPlan =
                    selectedRole === "store" ? storePlan : professionalPlan;
                  const checkoutState: CheckoutState = {
                    userId: professionalId,
                    userName: formData.name,
                    userEmail: formData.email,
                    planId: currentPlan?.id || 0,
                    userType: selectedRole as "professional" | "store",
                    payer: payer,
                  };

                  navigate("/checkout", { state: checkoutState });
                }, 2000);
              } else {
                navigate("/");
              }
            } else {
              Swal.fire({
                position: "center",
                icon: "error",
                title: "Ocorreu um erro na inclusão",
                text: `Erro: ${postReturn.status}`,
                showConfirmButton: false,
                timer: 1500,
              });
            }
          }
        } else {
          console.log("status != 200");
          Swal.fire({
            position: "center",
            icon: "error",
            title: "Erro Verificação Cliente",
            showConfirmButton: false,
            timer: 1500,
          });
          return false; // Email verification failed
        }
      } catch (error) {
        console.error("Erro no cadastro:", error);
        Swal.fire({
          position: "center",
          icon: "error",
          title: "Erro ao verificar e-mail",
          text: "Por favor, tente novamente mais tarde.",
          showConfirmButton: true,
        });
        return false; // Error during registration
      } finally {
        setIsLoading(false);
      }
    } catch (error) {
      console.error("Erro geral:", error);
      Swal.fire({
        position: "center",
        icon: "error",
        title: "Erro interno",
        text: "Por favor, tente novamente mais tarde.",
        showConfirmButton: true,
      });
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
      {showStoreModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-70 p-4">
          <div className="relative max-w-4xl w-full max-h-[95vh] rounded-lg overflow-hidden shadow-xl bg-white">
            <button
              onClick={handleCloseStoreModal}
              className="absolute top-4 right-4 z-10 text-gray-600 hover:text-gray-800 text-2xl bg-white bg-opacity-90 hover:bg-opacity-100 rounded-full w-10 h-10 flex items-center justify-center shadow-lg"
            >
              ✕
            </button>

            {/* Conteúdo do modal de lojista */}
            <div className="p-6 lg:p-8 flex flex-col justify-center bg-gradient-to-br from-blue-50 to-indigo-100">
              <div className="text-center mb-6">
                <h2 className="text-2xl lg:text-3xl font-bold text-gray-900 mb-2">
                  🚀 Lojista Parceiro Premium
                </h2>
                <div className="mb-3">
                  <span className="text-3xl font-bold text-[#FF6B35]">
                    {storePlan
                      ? `R$ ${storePlan.price.toFixed(2).replace(".", ",")}`
                      : "Carregando..."}
                  </span>
                  <span className="text-gray-500 text-sm ml-1">/mês</span>
                </div>
                <p className="text-gray-600 text-sm lg:text-base">
                  Destaque-se da concorrência e conquiste mais clientes
                </p>
              </div>

              <div className="space-y-4 mb-8">
                <div className="flex items-start space-x-3">
                  <div className="w-6 h-6 bg-green-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                    <Check className="w-3 h-3 text-white" />
                  </div>
                  <div>
                    <h4 className="font-semibold text-gray-900">
                      Possibilidade de anunciar
                    </h4>
                    <p className="text-sm text-gray-600">
                      Venda na plataforma Construir Mais Barato!
                    </p>
                  </div>
                </div>

                <div className="flex items-start space-x-3">
                  <div className="w-6 h-6 bg-green-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                    <Check className="w-3 h-3 text-white" />
                  </div>
                  <div className="flex-1">
                    <div className="flex items-center justify-between">
                      <div>
                        <h4 className="font-semibold text-gray-900">
                          Mais visibilidade para sua empresa
                        </h4>
                        <p className="text-sm text-gray-600">
                          Apareça antes das pesquisas
                        </p>
                      </div>
                      <div className="text-right">
                        <span className="text-2xl font-bold text-[#FF6B35]">
                          {storePlan
                            ? `R$ ${storePlan.price
                                .toFixed(2)
                                .replace(".", ",")}`
                            : "..."}
                        </span>
                        <span className="text-gray-500 text-sm ml-1">/mês</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div className="space-y-3">
                <button
                  type="button"
                  onClick={() => {
                    setPremiumStore(true);
                    setChoose(true);
                    setShowStoreModal(false);
                    setFormData((prev) => ({ ...prev, isPremiumStore: true }));
                    // Submeter o formulário programaticamente
                    setTimeout(() => {
                      formRef.current?.requestSubmit();
                    }, 100);
                  }}
                  className="w-full bg-gradient-to-r from-blue-600 to-indigo-600 text-white py-3 px-6 rounded-lg font-semibold hover:from-blue-700 hover:to-indigo-700 transition-all transform hover:scale-105 shadow-lg"
                >
                  ✨ Quero ser Lojista Parceiro Premium
                </button>

                <button
                  onClick={handleCloseStoreModal}
                  className="w-full bg-gray-200 text-gray-700 py-2 px-6 rounded-lg font-medium hover:bg-gray-300 transition-colors"
                >
                  Continuar com cadastro gratuito
                </button>
              </div>

              <p className="text-xs text-gray-500 text-center mt-4">
                * Você pode se tornar Lojista Parceiro Premium a qualquer
                momento
              </p>
            </div>
          </div>
        </div>
      )}
      {showPremiumModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-70 p-4">
          <div className="relative max-w-4xl w-full max-h-[95vh] rounded-lg overflow-hidden shadow-xl bg-white">
            <button
              onClick={handleClosePremiumModal}
              className="absolute top-4 right-4 z-10 text-gray-600 hover:text-gray-800 text-2xl bg-white bg-opacity-90 hover:bg-opacity-100 rounded-full w-10 h-10 flex items-center justify-center shadow-lg"
            >
              ✕
            </button>

            {/* <div className="grid grid-cols-1 lg:grid-cols-2 h-full"> */}
            {/* Lado esquerdo - Imagem */}
            {/* <div className="relative bg-black flex items-center justify-center min-h-[300px] lg:min-h-[500px]">
                <img
                  src="images/premiumBanner.jpeg"
                  alt="Construir Mais Barato Premium"
                  className="max-h-full object-cover w-full h-full lg:object-contain"
                />
                <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent lg:hidden"></div>
              </div> */}

            {/* Lado direito - Conteúdo */}
            <div className="p-6 lg:p-8 flex flex-col justify-center bg-gradient-to-br from-blue-50 to-indigo-100">
              <div className="text-center mb-6">
                <h2 className="text-2xl lg:text-3xl font-bold text-gray-900 mb-2">
                  🚀 Torne-se Premium
                  {/* Torne-se Premium */}
                </h2>
                <div className="mb-3">
                  <span className="text-3xl font-bold text-[#FF6B35]">
                    {professionalPlan
                      ? `R$ ${professionalPlan.price
                          .toFixed(2)
                          .replace(".", ",")}`
                      : "Carregando..."}
                  </span>
                  <span className="text-gray-500 text-sm ml-1">/mês</span>
                </div>
                <p className="text-gray-600 text-sm lg:text-base">
                  Destaque-se da concorrência e conquiste mais clientes
                </p>
              </div>

              <div className="space-y-4 mb-8">
                <div className="flex items-start space-x-3">
                  <div className="w-6 h-6 bg-green-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                    <Check className="w-3 h-3 text-white" />
                  </div>
                  <div>
                    <h4 className="font-semibold text-gray-900">
                      Perfil Destacado
                    </h4>
                    <p className="text-sm text-gray-600">
                      Apareça nas primeiras posições das buscas
                    </p>
                  </div>
                </div>

                <div className="flex items-start space-x-3">
                  <div className="w-6 h-6 bg-green-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                    <Check className="w-3 h-3 text-white" />
                  </div>
                  <div>
                    <h4 className="font-semibold text-gray-900">
                      Portifólio do profissional
                    </h4>
                    <p className="text-sm text-gray-600">
                      Permite que o cliente veja vídeos e fotos dos seus
                      trabalhos
                    </p>
                  </div>
                </div>

                <div className="flex items-start space-x-3">
                  <div className="w-6 h-6 bg-green-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                    <Check className="w-3 h-3 text-white" />
                  </div>
                  <div>
                    <h4 className="font-semibold text-gray-900">
                      Recebimento dos orçamentos em primeira mão
                    </h4>
                    <p className="text-sm text-gray-600">
                      Quando um cliente quiser solicitar um orçamento, o nome do
                      profissional premium apareçe primeiro
                    </p>
                  </div>
                </div>

                <div className="flex items-start space-x-3">
                  <div className="w-6 h-6 bg-green-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                    <Check className="w-3 h-3 text-white" />
                  </div>
                  {/* <div>
                    <h4 className="font-semibold text-gray-900">
                      Selo de premium
                    </h4>
                    <p className="text-sm text-gray-600">
                      Passa confiança no seu trabalho
                    </p>
                  </div> */}
                  <div className="flex-1">
                    <div className="flex items-center justify-between">
                      <div>
                        <h4 className="font-semibold text-gray-900">
                          Selo de premium
                        </h4>
                        <p className="text-sm text-gray-600">
                          Passa confiança no seu trabalho
                        </p>
                      </div>
                      <div className="text-right">
                        <span className="text-2xl font-bold text-[#FF6B35]">
                          {professionalPlan
                            ? `R$ ${professionalPlan.price
                                .toFixed(2)
                                .replace(".", ",")}`
                            : "..."}
                        </span>
                        <span className="text-gray-500 text-sm ml-1">/mês</span>
                      </div>
                    </div>
                  </div>
                </div>
                {/*                 
                <div className="flex items-start space-x-3">
                  <div className="w-6 h-6 bg-green-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                    <Check className="w-3 h-3 text-white" />
                  </div>
                  <div>
                    <span className="text-3xl font-bold text-blue-600">
                      R$ 19,90
                    </span>
                    <span className="text-gray-500 text-sm ml-1">/mês</span>
                  </div>
                </div> */}
              </div>

              <div className="space-y-3">
                <button
                  onClick={() => {
                    setPremiumForm(true);
                    setChoose(true);
                    setShowPremiumModal(false);
                  }}
                  className="w-full bg-gradient-to-r from-blue-600 to-indigo-600 text-white py-3 px-6 rounded-lg font-semibold hover:from-blue-700 hover:to-indigo-700 transition-all transform hover:scale-105 shadow-lg"
                >
                  ✨ Quero ser Premium
                </button>

                <button
                  onClick={() => {
                    handleClosePremiumModal();
                  }}
                  className="w-full bg-gray-200 text-gray-700 py-2 px-6 rounded-lg font-medium hover:bg-gray-300 transition-colors"
                >
                  Continuar com cadastro gratuito
                </button>
              </div>

              <p className="text-xs text-gray-500 text-center mt-4">
                * Você pode migrar para premium a qualquer momento
              </p>
            </div>
            {/* </div> */}
          </div>
        </div>
      )}

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
      {/* {showSuccessModal && (
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
      )} */}

      <div className="sm:mx-auto sm:w-full sm:max-w-2xl">
        <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
          Cadastre-se
        </h2>
        <p className="mt-2 text-center text-sm text-gray-600">
          Preencha seus dados para criar sua conta
        </p>
        {choose && (
          <div className="mt-4 text-center">
            <span
              className={`inline-flex items-center px-3 py-1 rounded-full text-sm font-medium ${
                premiumForm
                  ? "bg-blue-100 text-blue-800 border border-blue-300"
                  : "bg-gray-100 text-gray-800 border border-gray-300"
              }`}
            >
              {premiumForm ? "🚀 Cadastro Premium" : "👤 Cadastro Gratuito"}
            </span>
            <button
              onClick={() => {
                setChoose(false);
                setPremiumForm(false);
                // Limpar campos premium
                setFormData((prev) => ({
                  ...prev,
                  dateOfBirth: "",
                  experience: "",
                  codeVerification: "",
                  meiCnpj: "",
                  isPremium: false,
                }));
                setCodigoEnviado(false);
              }}
              className="ml-2 text-xs text-blue-600 hover:text-blue-800 underline"
            >
              Alterar
            </button>
          </div>
        )}
      </div>

      <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-2xl">
        <div className="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
          {error && <ErrorAlert message={error} onClose={closeError} />}

          {/* Role Selection - Ocultar em modo upgrade */}
          {!isUpgradeMode && (
            <div className="mb-8">
              <label className="block text-sm font-medium text-gray-700 mb-4">
                Como você deseja se cadastrar?
              </label>
              <div className="grid grid-cols-1 md:grid-cols-1 gap-4">
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
          )}

          <form ref={formRef} className="space-y-6" onSubmit={handleSubmit}>
            {/* Campos básicos - Ocultar em modo upgrade */}
            {!isUpgradeMode && (
              <>
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
                      onClick={() =>
                        setShowConfirmPassword(!showConfirmPassword)
                      }
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
                {selectedRole === "store" && (
                  <div className="relative">
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      Especialidade
                    </label>
                    <div className="relative">
                      <HardHat className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                      <select
                        value={selectedCategoryProduct}
                        onChange={(e) => {
                          setSelectedCategoryProduct(e.target.value);
                          // toggleAllSubcategories();
                        }}
                        className="block w-full pl-10 pr-4 py-2.5 text-gray-900 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none bg-white"
                      >
                        <option value="">Seleciona uma especialidade</option>
                        {categoryProducts.map((cp: ICategoryProduct) => (
                          <option key={cp.id} value={cp.id}>
                            {cp.name}
                          </option>
                        ))}
                      </select>
                    </div>
                    
                    {/* {subcategories.length > 0 && toggleAllSubcategories()} */}
                    {/* {subcategories.length > 0 && (
                      <div className="mt-6 bg-gradient-to-br from-blue-50 to-indigo-50 rounded-xl p-6 border-2 border-blue-200 shadow-lg">
                        <div className="flex items-center justify-between mb-4">
                          <h3 className="text-lg font-semibold text-gray-800 flex items-center gap-2">
                            <HardHat className="w-5 h-5 text-blue-600" />
                            Subcategorias Disponíveis
                          </h3>
                          <button
                            onClick={toggleAllSubcategories}
                            className="text-sm text-blue-600 hover:text-blue-800 font-medium transition-colors"
                          >
                            {selectedSubcategories.length ===
                            subcategories.length
                              ? "Desmarcar Todas"
                              : "Selecionar Todas"}
                          </button>
                        </div>

                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3">
                          {subcategories.map((subcategory) => (
                            <label
                              key={subcategory.id}
                              className={`
                              flex items-center gap-3 p-3 rounded-lg cursor-pointer transition-all
                              ${
                                selectedSubcategories.includes(subcategory.id)
                                  ? "bg-blue-600 text-white shadow-md scale-105"
                                  : "bg-white text-gray-700 hover:bg-blue-100 hover:shadow-sm"
                              }
                            `}
                            >
                              <input
                                type="checkbox"
                                checked={selectedSubcategories.includes(
                                  subcategory.id
                                )}
                                onChange={() =>
                                  toggleSubcategory(subcategory.id)
                                }
                                className="w-5 h-5 rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                              />
                              <span className="font-medium text-sm">
                                {subcategory.name}
                              </span>
                            </label>
                          ))}
                        </div>

                        {selectedSubcategories.length > 0 && (
                          <div className="mt-4 pt-4 border-t border-blue-200">
                            <p className="text-sm text-gray-600">
                              <span className="font-semibold text-blue-600">
                                {selectedSubcategories.length}
                              </span>{" "}
                              subcategoria
                              {selectedSubcategories.length !== 1
                                ? "s"
                                : ""}{" "}
                              selecionada
                              {selectedSubcategories.length !== 1 ? "s" : ""}
                            </p>
                          </div>
                        )}
                      </div>
                    )} */}
                  </div>
                )}
              </>
            )}

            {premiumForm && (
              <div className="bg-gradient-to-r from-blue-50 to-indigo-50 p-6 rounded-lg border border-blue-200">
                <div className="mb-6">
                  <h3 className="text-lg font-semibold text-blue-900 mb-2">
                    🚀 Cadastro Premium Verificado
                  </h3>
                  <p className="text-sm text-blue-700">
                    Complete os campos abaixo para ter um perfil verificado e
                    confiável
                  </p>
                </div>

                {/* Telefone - Mostrar em modo upgrade */}
                {isUpgradeMode && (
                  <div className="mb-6">
                    <label
                      htmlFor="phone-upgrade"
                      className="block text-sm font-medium text-gray-700 mb-2"
                    >
                      Telefone Comercial *
                    </label>
                    <div className="relative rounded-md shadow-sm">
                      <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                        <Phone className="h-5 w-5 text-gray-400" />
                      </div>
                      <InputMask
                        mask="(99) 99999-9999"
                        id="phone-upgrade"
                        name="phone"
                        type="tel"
                        required
                        value={formData.phone}
                        onChange={handleChange}
                        className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                        placeholder="(11) 99999-9999"
                      />
                    </div>
                    <p className="text-xs text-gray-500 mt-1">
                      Necessário para receber o código de verificação SMS
                    </p>
                  </div>
                )}

                {/* Foto de Perfil */}
                <div className="mb-6">
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    Foto de Perfil *
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
                      {previewUrl && (
                        <button
                          type="button"
                          onClick={() => {
                            setPreviewUrl("");
                            setFormData((prev) => ({ ...prev, photo: "" }));
                          }}
                          className="absolute -top-2 -right-2 p-1 bg-red-600 text-white rounded-full hover:bg-red-700 transition-colors"
                        >
                          <X className="w-3 h-3" />
                        </button>
                      )}
                    </div>
                  </div>
                  <p className="text-xs text-gray-500 text-center mt-2">
                    Uma foto profissional aumenta sua credibilidade
                  </p>
                </div>

                {/* Data de Nascimento */}
                <div className="mb-4">
                  <label
                    htmlFor="dateOfBirth"
                    className="block text-sm font-medium text-gray-700"
                  >
                    Data de Nascimento *
                  </label>
                  <div className="mt-1 relative rounded-md shadow-sm">
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                      <User className="h-5 w-5 text-gray-400" />
                    </div>
                    <input
                      id="dateOfBirth"
                      name="dateOfBirth"
                      type="date"
                      required
                      value={formData.dateOfBirth}
                      onChange={handleChange}
                      className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                  {errorAge && (
                    <p className="mt-1 text-sm text-red-600">{errorAge}</p>
                  )}
                  <p className="text-xs text-gray-500 mt-1">
                    Necessário para confirmar maioridade e identidade
                  </p>
                </div>

                {/* Experiência */}
                <div className="mb-4">
                  <label
                    htmlFor="experience"
                    className="block text-sm font-medium text-gray-700"
                  >
                    Qual sua experiência? (Opcional)
                  </label>
                  <textarea
                    id="experience"
                    name="experience"
                    rows={3}
                    value={formData.experience}
                    onChange={handleChange}
                    className="mt-1 appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    placeholder="Conte sobre sua experiência, projetos anteriores, certificações..."
                  />
                  <p className="text-xs text-gray-500 mt-1">
                    Ajude os clientes a conhecerem melhor seu trabalho
                  </p>
                </div>

                {/* Verificação do Telefone */}
                <div className="mb-4">
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    🔒 Verificação do Telefone *
                  </label>
                  <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-4 mb-3">
                    <div className="flex items-center">
                      <Shield className="w-5 h-5 text-yellow-600 mr-2" />
                      <p className="text-sm text-yellow-800">
                        Um código de 4 a 6 dígitos será enviado para seu
                        telefone para garantir que é real
                      </p>
                    </div>
                  </div>

                  <div className="flex gap-3">
                    <button
                      type="button"
                      onClick={enviarCodigoSMS}
                      disabled={loadingSMS || !formData.phone}
                      className="flex-shrink-0 bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:bg-gray-400 disabled:cursor-not-allowed"
                    >
                      {loadingSMS
                        ? "Enviando..."
                        : codigoEnviado
                        ? "Reenviar"
                        : "Enviar Código"}
                    </button>

                    <div className="flex-1">
                      <input
                        id="codeVerification"
                        name="codeVerification"
                        type="text"
                        required
                        maxLength={6}
                        value={formData.codeVerification}
                        onChange={handleChange}
                        disabled={!codigoEnviado}
                        className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100"
                        placeholder="Digite o código recebido"
                      />
                    </div>
                  </div>

                  {codigoEnviado && (
                    <p className="text-xs text-green-600 mt-1">
                      ✓ Código enviado para {formData.phone}
                    </p>
                  )}
                </div>

                {/* MEI ou CNPJ */}
                <div className="mb-4">
                  <label
                    htmlFor="meiCnpj"
                    className="block text-sm font-medium text-gray-700"
                  >
                    MEI ou CNPJ (Opcional)
                  </label>
                  <div className="mt-1 relative rounded-md shadow-sm">
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                      <Building className="h-5 w-5 text-gray-400" />
                    </div>
                    <InputMask
                      mask="99.999.999/9999-99"
                      id="meiCnpj"
                      name="meiCnpj"
                      type="text"
                      value={formData.meiCnpj}
                      onChange={handleChange}
                      className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                      placeholder="00.000.000/0000-00"
                    />
                  </div>
                  <p className="text-xs text-gray-500 mt-1">
                    Permite consulta pública na Receita Federal e aumenta a
                    confiança
                  </p>
                </div>

                <div className="mb-4">
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    Certidão Negativa *
                  </label>

                  {/* Card para gerar certidão */}
                  <div className="bg-blue-50 border border-blue-200 rounded-lg p-4 mb-3">
                    <p className="text-sm text-gray-700 mb-3">
                      Gere sua certidão no site oficial:
                    </p>
                    <button
                      type="button"
                      onClick={handleGerarCertidao}
                      className="inline-flex items-center bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors"
                    >
                      <ExternalLink className="w-4 h-4 mr-2" />
                      Gerar Certidão
                    </button>
                  </div>

                  {/* Input separado para número da certidão */}
                  <input
                    id="negativeCertificateNumber"
                    name="negativeCertificateNumber"
                    type="number"
                    required
                    value={formData.negativeCertificateNumber}
                    onChange={handleChange}
                    disabled={!clickedCertificateButton}
                    className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    placeholder="Digite o número da certidão"
                  />
                </div>
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
                  (!isUpgradeMode &&
                    selectedRole === "professional" &&
                    formData.professions.length === 0) ||
                  (premiumForm &&
                    (!formData.dateOfBirth ||
                      !formData.codeVerification ||
                      !formData.photo ||
                      (isUpgradeMode && !formData.phone)))
                }
                className="w-full flex justify-center items-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:bg-gray-400 disabled:cursor-not-allowed"
              >
                {premiumForm ? "🚀 Cadastrar Premium" : "Cadastrar"}
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
