import { useState, useEffect } from "react";
import {
  X,
  ShoppingBag,
  Search,
  Shield,
  ChevronRight,
  LayoutDashboard,
  Youtube,
} from "lucide-react";
import SearchProfessionals from "./SearchProfessionals";
import Marketplace from "./Marketplace";
import Login from "./Login";
import Register from "./Register";
import ProfessionalPanel from "./ProfessionalPanel";
import Dashboard from "./Dashboard";

import PrivacyPolicy from "./PrivacyPolicy";
import CookieBanner from "../components/CookieBanner";
import VideoPopup from "../components/VideoPopup";
import Carousel from "../components/Carousel";

function Home() {
  const [currentPage, setCurrentPage] = useState(() => {
    const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";
    return isLoggedIn ? "professional-panel" : "home";
  });
  const [carouselPage, setCarouselPage] = useState("H");
  const [showDashboardPopup, setShowDashboardPopup] = useState(false);
  const [isVideoPopupOpen, setIsVideoPopupOpen] = useState(false);

  useEffect(() => {
    //    if (currentPage === 'professional-panel') {
    //      localStorage.setItem('isLoggedIn', 'true');
    //    }
  }, [currentPage, carouselPage]);

  const renderContent = () => {
    switch (currentPage) {
      case "search":
        return <SearchProfessionals onNavigate={setCurrentPage} />;
      case "marketplace":
        return <Marketplace onNavigate={setCurrentPage} />;
      case "login":
        return <Login onNavigate={setCurrentPage} />;
      case "register":
        return <Register onNavigate={setCurrentPage} />;
      case "professional-panel":
        return <ProfessionalPanel />;
      case "dashboard":
        return <Dashboard />;
      //case 'dashboard':
      //  return <DashboardNovo/>;
      /*case 'marketplace-products':
        return <MarketplaceProducts/>;      */
      case "privacy":
        return <PrivacyPolicy onBack={() => setCurrentPage("home")} />;
      default:
        return (
          <>
            {/* Home Page Content */}
            <main className="max-w-7xl mx-auto px-4 py-8">
              <div className="grid md:grid-cols-2 gap-8 items-center">
                <div>
                  <h1 className="text-4xl font-bold text-gray-900 mb-4">
                    Encontre os Melhores Profissionais da Construção Civil
                  </h1>
                  <p className="text-lg text-gray-600 mb-6">
                    Conectamos você aos melhores profissionais do mercado, com
                    preços justos e qualidade garantida.
                  </p>
                  <button
                    onClick={() => {
                      setCurrentPage("search");
                      setCarouselPage("S");
                    }}
                    className="bg-blue-600 text-white px-6 py-3 rounded-lg font-medium hover:bg-blue-700 transition-colors"
                  >
                    Encontrar Profissionais
                  </button>
                </div>
                <div className="relative h-[300px] md:h-[400px]">
                  <img
                    src="images/photo-1504307651254-35680f356dfd.jpeg"
                    alt="Construção"
                    className="w-full h-full object-cover rounded-lg shadow-xl"
                  />
                </div>
              </div>

              {/* Features Section */}
              <div className="grid md:grid-cols-3 gap-8 mt-16">
                <div className="bg-white p-6 rounded-lg shadow-md">
                  <ShoppingBag className="w-12 h-12 text-blue-600 mb-4" />
                  <h3 className="text-xl font-semibold mb-2">Marketplace</h3>
                  <p className="text-gray-600">
                    Encontre materiais de construção com os melhores preços do
                    mercado.
                  </p>
                </div>
                <div className="bg-white p-6 rounded-lg shadow-md">
                  <Search className="w-12 h-12 text-blue-600 mb-4" />
                  <h3 className="text-xl font-semibold mb-2">
                    Profissionais da Construção Civil
                  </h3>
                  <p className="text-gray-600">
                    Todos os profissionais da construção civil , você encontra
                    aqui.
                  </p>
                </div>
                <div className="bg-white p-6 rounded-lg shadow-md">
                  <Shield className="w-12 h-12 text-blue-600 mb-4" />
                  <h3 className="text-xl font-semibold mb-2">
                    Segurança Garantida
                  </h3>
                  <p className="text-gray-600">
                    Sua satisfação e segurança são nossas principais
                    prioridades.
                  </p>
                </div>
              </div>
            </main>
            {/* Floating YouTube Button */}
            <button
              onClick={() => setIsVideoPopupOpen(true)}
              className="fixed bottom-24 right-1 z-40 flex flex-col items-center"
            >
              <span className="text-xs font-medium text-white bg-red-600 px-3 py-1 rounded-full mb-2 shadow-md">
                QUEM SOMOS
              </span>
              <div className="w-16 h-16 bg-red-600 rounded-full shadow-lg hover:bg-red-700 transition-colors flex items-center justify-center">
                <Youtube className="w-8 h-8 text-white" />
              </div>
            </button>
          </>
        );
    }
  };
  return (
    <>
      <div className="min-h-screen bg-gray-50">
        {/*<Navigation
          currentPage={currentPage}
          setCurrentPage={setCurrentPage}
          carouselPage={""}
          setCarouselPage={function (page: string): void {
            throw new Error("Function not implemented.");
          }}
        />*/}
        <div className="pt-16">{renderContent()}</div>

        {/* Floating Dashboard Popup */}
        {showDashboardPopup && currentPage === "home" && (
          <div className="fixed bottom-4 right-4 bg-white rounded-lg shadow-xl p-6 max-w-sm animate-slide-up">
            <button
              onClick={() => setShowDashboardPopup(false)}
              className="absolute top-2 right-2 text-gray-400 hover:text-gray-600 transition-colors"
            >
              <X className="w-5 h-5" />
            </button>
            <div className="flex items-center gap-3 mb-3">
              <div className="bg-blue-100 rounded-full p-2">
                <LayoutDashboard className="w-6 h-6 text-blue-600" />
              </div>
              <h3 className="text-lg font-semibold text-gray-900">
                Acesse o Dashboard
              </h3>
            </div>
            <p className="text-gray-600 mb-4">
              Visualize estatísticas, gerencie profissionais e acompanhe o
              crescimento da plataforma.
            </p>
            <button
              onClick={() => {
                setCurrentPage("dashboard");
                setCarouselPage("D");
                setShowDashboardPopup(false);
              }}
              className="w-full flex items-center justify-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            >
              <span>Acessar Dashboard</span>
              <ChevronRight className="w-5 h-5" />
            </button>
          </div>
        )}
      </div>
      <CookieBanner onNavigate={setCurrentPage} />
      <VideoPopup
        isOpen={isVideoPopupOpen}
        onClose={() => setIsVideoPopupOpen(false)}
        url="https://www.youtube.com/embed/EpOykD8vDRU"
      />
      {/*}
      <Footer></Footer>
          */}
    </>
  );
}

export default Home;
