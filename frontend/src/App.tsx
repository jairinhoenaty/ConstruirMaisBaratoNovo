import SearchProfessionals from "./pages/SearchProfessionals";
import Marketplace from "./pages/Marketplace";
import Login from "./pages/Login";
import Register from "./pages/Register";
import ProfessionalPanel from "./pages/ProfessionalPanel";
import Checkout from "./pages/Checkout";
import CheckoutUnlock from "./pages/CheckoutUnlock";
import Dashboard from "./pages/Dashboard";
import Footer from "./components/Footer";
import PrivacyPolicy from "./pages/PrivacyPolicy";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Home from "./pages/Home";
import Header from "./components/Header";
import BannerHeader from "./components/BannerHeader";
import SearchResults from "./pages/SearchResults";
import ResetPassword from "./pages/ResetPassword";
import QuoteProducts from "./pages/QuoteProducts";
import ProductCategoriesAdmin from "./pages/ProductCategoriesAdmin";

function App() {
  return (
    <>
      <div className="min-h-screen bg-gray-50">
        <BrowserRouter>
          <Header></Header>
          <BannerHeader></BannerHeader>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/marketplace" element={<Marketplace />} />
            <Route path="/search" element={<SearchProfessionals />} />
            <Route path="/quote-product" element={<QuoteProducts />} />
            <Route path="/privacy" element={<PrivacyPolicy />} />
            <Route path="/checkout" element={<Checkout />} />
            <Route path="/checkout-unlock" element={<CheckoutUnlock />} />

            <Route path="/dashboard" element={<Dashboard />} />
            <Route path="/search-results" element={<SearchResults />} />
            <Route path="/professional-panel" element={<ProfessionalPanel />} />
            <Route path="/product-categories-admin" element={<ProductCategoriesAdmin />} />
            <Route path="/confirmar-senha/:token" element={<ResetPassword />} />
          </Routes>
        </BrowserRouter>
      </div>
      <div>
        <Footer></Footer>
      </div>
    </>
  );
}

export default App;
