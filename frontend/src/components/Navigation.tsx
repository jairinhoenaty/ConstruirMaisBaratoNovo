import {
  Home,
  LayoutDashboard,
  LogIn,
  Menu,
  Search,
  Shield,
  ShoppingBag,
  UserPlus,
  X,
} from "lucide-react";
import { useState } from "react";
import { Link, useLocation, useNavigate } from "react-router-dom";

function Navigation({
  currentPage,
  setCurrentPage,
}: {
  currentPage: string;
  setCurrentPage: (page: string) => void;
}) {
  const navigate = useNavigate();
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
  //  const [isAppPopupOpen, setIsAppPopupOpen] = useState(false);
  const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";
  const isAdmin = localStorage.getItem("profile") === "admin";
  const isClient = localStorage.getItem("profile") === "client";  

  const menuItems = [
    { id: "", label: "Início", icon: Home, page: "H" },
    ...(!isLoggedIn || (isLoggedIn && isClient)
      ? [
          {
            id: "marketplace",
            label: "Loja",
            icon: ShoppingBag,
            page: "M",
          },
        ]
      : []),
    ...(!isLoggedIn || (isLoggedIn && isClient)
      ? [
          {
            id: "search",
            label: "Pesquisar Profissional",
            icon: Search,
            page: "S",
          },
        ]
      : []),
    ...(!isLoggedIn || (isLoggedIn && isClient)
      ? [
          {
            id: "quote-product",
            label: "Consultar Preços",
            icon: Search,
            page: "CP",
          },
        ]
      : []),
   
    ...(isLoggedIn
      ? isAdmin
        ? [
            {
              id: "dashboard",
              label: "Dashboard",
              icon: LayoutDashboard,
              page: "D",
            },
            {
              id: "professional-panel",
              label: "Área de Trabalho",
              icon: UserPlus,
              page: "N",
            },
            { id: "logout", label: "Sair", icon: LogIn },
          ]
        : [
            {
              id: "professional-panel",
              label: "Área do Trabalho",
              icon: UserPlus,
              page: "N",
            },
            { id: "logout", label: "Sair", icon: LogIn },
          ]
      : [
          { id: "register", label: "Cadastre-se", icon: UserPlus, page: "R" },
          { id: "login", label: "Login", icon: LogIn, page: "L" },
        ]),
  ];
  const location = useLocation();
  setCurrentPage(location.pathname.substring(1, 100));

  const handleMenuClick = (pageId: string) => {
    if (pageId === "logout") {
      localStorage.removeItem("isLoggedIn");
      localStorage.removeItem("token");
      localStorage.removeItem("user");
      localStorage.removeItem("id");
      localStorage.removeItem("name");
      localStorage.removeItem("profile");
      setCurrentPage(pageId);
      navigate("/");
      window.location.reload();
    } else {
      setCurrentPage(pageId);
    }
    setIsMobileMenuOpen(false);
  };

  return (
    <>
      <nav className="fixed top-0 left-0 right-0 bg-blue-600 text-white shadow-lg z-50">
        <div className="max-w-7xl mx-auto px-4">
          <div className="flex justify-between items-center h-16 md:h-20 lg:h-16">
            {/* Logo */}
            <div className="flex items-center space-x-2">
            <Link to="/">
              <img
                src="../../public/images/Hassiscompany.png"
                alt="HASSIS CONECTA"
                className="h-10 md:h-12 lg:h-14 w-auto object-contain rounded-full"
                //style={{ width: 50 }}
              />
            </Link>

              <span className="block font-bold text-lg">
               HASSIS CONECTA
              </span>
            </div>

            {/* Desktop Menu */}
            <div className="hidden lg:flex items-center">
              {menuItems.map((item) => {
                const Icon = item.icon;
                return (
                  <Link to={item.id == "logout" ? "/" : "/"+item.id}>
                    <button
                      key={item.id}
                      onClick={() => handleMenuClick(item.id)}
                      className={`flex items-center h-full px-3 text-sm font-medium hover:bg-blue-700 transition-colors ${
                      currentPage === item.id ? "bg-blue-700" : ""
                    }`}
                    >
                      <Icon className="w-4 h-4 mr-1.5" />
                      <span>{item.label}</span>
                    </button>
                  </Link>
                );
              })}
              {/*
              <button 
                onClick={() => setIsAppPopupOpen(true)}
                className="flex items-center space-x-1.5 ml-3 px-4 h-8 bg-white text-blue-600 rounded-full text-sm font-medium hover:bg-blue-50 transition-colors"
              >
                <Smartphone className="w-4 h-4" />
                <span>Profissional Já</span>
              </button>*/}
            </div>

           {/* Mobile / Tablet Visible Buttons */}
             <div className="hidden md:flex lg:hidden items-center space-x-2">
              {menuItems
                .filter(
                  (item) =>
                    item.id === "search" ||
                    item.id === "login" ||
                    item.id === "register"
                )
                .map((item) => {
                  const Icon = item.icon;
                  return (
                    <Link key={item.id} to={"/" + item.id}>
                      <button
                        onClick={() => handleMenuClick(item.id)}
                        className="flex items-center px-2 py-1 text-xs font-medium hover:bg-blue-700 rounded-md transition-colors"
                      >
                        <Icon className="w-4 h-4 mr-1" />
                        <span>{item.label}</span>
                      </button>
                    </Link>
                  );
                })}
            </div>

            {/* Mobile Menu Button */}
            <div className="lg:hidden">
              <button
                onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
                className="p-2 rounded-md hover:bg-blue-700 transition-colors"
              >
                {isMobileMenuOpen ? (
                  <X className="w-6 h-6" />
                ) : (
                  <Menu className="w-6 h-6" />
                )}
              </button>
            </div>
          </div>
        </div>

        {/* Mobile Menu */}
        {isMobileMenuOpen && (
          <div className="lg:hidden">
            <div className="px-2 pt-2 pb-3 space-y-1">
              {menuItems.map((item) => {
                const Icon = item.icon;
                return (
                  <Link to={item.id == "logout" ? "/" : "/" + item.id}>
                    <button
                      key={item.id}
                      onClick={() => handleMenuClick(item.id)}
                      className={`flex items-center space-x-2 w-full px-3 py-2 rounded-md hover:bg-blue-700 transition-colors ${
                        currentPage === item.id ? "bg-blue-700" : ""
                      }`}
                    >
                      <Icon className="w-4 h-4" />
                      <span>{item.label}</span>
                    </button>
                  </Link>
                );
              })}
              {/*<button 
                onClick={() => setIsAppPopupOpen(true)}
                className="flex items-center space-x-2 w-full px-3 py-2 bg-white text-blue-600 rounded-md hover:bg-blue-50 transition-colors"
              >
                <Smartphone className="w-4 h-4" />
                <span>Profissional Já</span>
              </button>*/}
            </div>
          </div>
        )}
      </nav>

      {/*
      Botão de Download APP
      <AppDownloadPopup 
        isOpen={isAppPopupOpen}
        onClose={() => setIsAppPopupOpen(false)}
      />*/}
    </>
  );
}

export default Navigation;
