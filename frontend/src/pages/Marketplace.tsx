import React, { useState } from "react";
import { Search, MapPin, Tag, PlusCircle, Package } from "lucide-react";
import { states } from "../data";
import { CityService } from "../services/CityService";
import { ProductService } from "../services/ProductService";
import ProductDetails from "./ProductDetails";
import { Link } from "react-router-dom";

interface MarketplaceProps {
  onNavigate?: (page: string) => void;
}

function Marketplace({ onNavigate }: MarketplaceProps) {
  const [selectedState, setSelectedState] = useState("");
  const [selectedCity, setSelectedCity] = useState("");
  const [citiesByState, setcitiesByState] = useState([]);
  const [dailyDeals, setdailyDeals] = useState([]);
  const [selectedProduct, setSelectedProduct] = useState<any>(null);
  const [listDesc, setListDesc] = useState("Ofertas do dia"); //Ofertas do Dia

  
  React.useEffect(() => {
    const fetchData = async () => {
      // const id = parseInt(localStorage.getItem("id") || "");

      const dailyDeals = await ProductService.productsDayOfferPublic();

      const json_dayoffer = await dailyDeals.data;
      if ((dailyDeals.status == 200)) {
        setdailyDeals(json_dayoffer);
      }
    };

    fetchData();
  }, []);

  const handleChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    if (name == "state") {
      const citiesByState = await CityService.citiesByStatePublic({
        uf: value,
      });

      const json_cities = await citiesByState.data;
      if ((citiesByState.status == 200)) {
        setcitiesByState(json_cities);
      }
    }
    setSelectedState(value);
  };

  //const SearchProducts = async (e: React.ChangeEvent<HTMLInputElement>) => {
  const SearchProducts = async () => {
    //const { name, value } = e.target;

    //  const id = parseInt(localStorage.getItem("id") || "");
    const cityID = selectedCity;

    const dailyDeals = await ProductService.productsByCity({
      cityID: parseInt(cityID),
    });

    const json_dayoffer = await dailyDeals.data;
    if ((dailyDeals.status == 200)) {
      setdailyDeals(json_dayoffer);
      setListDesc("Produtos");
    }
  };

  if (selectedProduct) {
    return (
      <ProductDetails
        product={selectedProduct}
        onBack={() => setSelectedProduct(null)}
      />
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Search Section */}
      <div className="max-w-7xl mx-auto px-4 py-8">
        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-2xl font-bold text-gray-900 mb-6">
            Encontre os Produtos da sua Cidade
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div className="relative">
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Estado
              </label>
              <div className="relative">
                <MapPin className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                <select
                  id="state"
                  name="state"
                  value={selectedState}
                  //onChange={(e) => setSelectedState(e.target.value)}
                  onChange={(e) => handleChange(e as any)}
                  className="block w-full pl-10 pr-4 py-2.5 text-gray-900 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none bg-white"
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

            <div className="relative">
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Cidade
              </label>
              <div className="relative">
                <MapPin className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
                <select
                  value={selectedCity}
                  onChange={(e) => setSelectedCity(e.target.value)}
                  disabled={!selectedState}
                  className="block w-full pl-10 pr-4 py-2.5 text-gray-900 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none bg-white disabled:bg-gray-100 disabled:cursor-not-allowed"
                >
                  <option value="">Selecione a cidade</option>
                  {selectedState &&
                    citiesByState.map((city: any) => (
                      <option key={city.id} value={city.id}>
                        {city.name}
                      </option>
                    ))}
                </select>
              </div>
            </div>

            <div className="flex gap-2 items-end">
              <button
                className="flex-1 h-[42px] px-6 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center justify-center gap-2"
                onClick={() => SearchProducts()}
              >
                <Search className="w-5 h-5" />
                Buscar Produtos
              </button>
              <Link to="/login">
                <button
                  onClick={() => onNavigate && onNavigate("login")}
                  className="h-[42px] px-6 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors flex items-center justify-center gap-2"
                >
                  <PlusCircle className="w-5 h-5" />
                  Vender
                </button>
              </Link>
            </div>
          </div>
        </div>

        {/* Daily Deals Section */}
        <div className="mt-12">
          <div className="flex items-center gap-2 mb-6">
            <Tag className="w-6 h-6 text-blue-600" />
            <h2 className="text-2xl font-bold text-gray-900">{listDesc}</h2>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
            {dailyDeals.map((product: any) => (
              <div
                key={product.id}
                className="bg-white rounded-lg shadow-md overflow-hidden"
              >
                <div className="relative">
                  <img
                    src={
                      product.image
                        ? "data:image;base64," + product.image
                        : "/images/default_product.png"
                    }
                    alt={product.name}
                    className="w-full h-48 object-cover"
                    onClick={() => setSelectedProduct(product)}
                  />
                  {product.dayoffer && (
                    <div className="absolute top-2 right-2 bg-red-600 text-white px-2 py-1 rounded-md text-sm font-semibold">
                      {product.dayoffer ? "Em destaque" : product.dayoffer}
                    </div>
                  )}
                </div>
                <div className="p-4">
                  <h3 className="text-gray-900 font-semibold mb-2 line-clamp-2">
                    {product.name}
                  </h3>
                  <div className="flex items-center gap-2">
                    <span className="text-gray-500 line-through text-sm">
                      {product.originalprice !== 0
                        ? product.originalprice.toLocaleString("pt-BR", {
                            style: "currency",
                            currency: "BRL",
                            minimumFractionDigits: 2,
                          })
                        : ""}
                    </span>
                    <span className="text-blue-600 font-bold">
                      {product.price!==0?product.price.toLocaleString("pt-BR", {
                        style: "currency",
                        currency: "BRL",
                        minimumFractionDigits: 2,
                      }):""}
                    </span>
                  </div>
                </div>
              </div>
            ))}
          </div>
          {dailyDeals.length === 0 && (
            <div className="text-center py-12">
              <Package className="w-12 h-12 text-gray-400 mx-auto mb-4" />
              <h3 className="text-lg font-medium text-gray-900 mb-2">
                Nenhum produto encontrado
              </h3>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default Marketplace;
