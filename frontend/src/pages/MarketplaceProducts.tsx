import React, { useState } from "react";
import { Package, Pause, Check, Trash2, Tag } from "lucide-react";
import { ProductService } from "../services/ProductService";
import Pagination from "../components/Pagination";
import { IProduct } from "../interfaces/IProduct";
import Swal from "sweetalert2";

interface Product {
  id: string;
  title: string;
  price: number;
  image: string;
  description: string;
  status: "paused" | "approved" | "daily-offer";
  dayoffer: boolean;
  approved: boolean;
}

function MarketplaceProducts() {
  const [activeTab, setActiveTab] = useState<
    "paused" | "approved" | "daily-offer"
  >("paused");
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);
  const [isUpdate, setIsUpdate] = useState(false);

  const [products, setProducts] = useState<Product[]>([]);

  
  React.useEffect(() => {
    const fetchData = async () => {
      const limit = 9;
      //let result_products: any;
      let approved: string | null = "S";
      let dayoffer: string | null = "N";
      if (activeTab == "paused") {
        approved = "N";
        dayoffer = "N";
      }
      if (activeTab == "daily-offer") {
        dayoffer = "S";
        approved = "";
      }
      console.log(activeTab);
      //console.log(parseInt(localStorage.getItem("post_id") ?? ""));
      const result_products = await ProductService.productsAll(
        limit,
        (page - 1) * limit,
        0,
        0,
        dayoffer,
        approved
      );
      console.log(result_products);
      if (result_products.status == 200) {
        const json_products = await result_products.data.products;
        console.log(json_products);
        setTotalPage(Math.ceil(result_products.data.total / limit));
        setProducts(json_products);
      }
    };
    fetchData();
  }, [activeTab, page, isUpdate]);

  const handleStatusChange = async (
    //productId: string,
    product_approved: IProduct,
    newStatus: "paused" | "approved" | "daily-offer"
  ) => {
    if (newStatus == "approved") {
      product_approved.approved = true;
    } else if (newStatus == "paused") {
      product_approved.approved = false;
      product_approved.dayoffer = false;
    } else if (newStatus == "daily-offer") {
      product_approved.dayoffer = true;
    }

    console.log(newStatus);
    console.log(product_approved);

    const response = await ProductService.saveProduct(product_approved);
    if (response.status == 200) {
      Swal.fire({
        icon: "success",
        text: "Produto alterado",
        showConfirmButton: false,
        timer: 1500,
      });
      setIsUpdate(!isUpdate);
    }
    /*setProducts(
      products.map((product) =>
        product.id === productId ? { ...product, status: newStatus } : product
      )
    );*/
  };

  const handleDelete = async (productId: string) => {
    if (window.confirm("Tem certeza que deseja excluir este produto?")) {
      const response = await ProductService.deleteProduct(productId);
      if ((response.status == 200)) {
        Swal.fire({
          icon: "success",
          text: "Produto excluído",
          showConfirmButton: false,
          timer: 1500,
        });
        setIsUpdate(!isUpdate);
      }
      //setProducts(products.filter((product) => product.id !== productId));
    }
  };

  /*const filteredProducts = products.filter(
    (product) => product.status === activeTab
  );
  */

  return (
    <div className="min-h-screen bg-gray-50 p-8">
      <div className="max-w-7xl mx-auto">
        {/* Header */}
        <div className="flex items-center justify-between mb-8">
          <div className="flex items-center gap-3">
            <Package className="w-8 h-8 text-blue-600" />
            <h1 className="text-2xl font-bold text-gray-900">
              Produtos do Marketplace
            </h1>
          </div>
        </div>

        {/* Navigation Tabs */}
        <div className="flex gap-4 mb-8">
          <button
            onClick={() => setActiveTab("paused")}
            className={`flex items-center gap-2 px-6 py-3 rounded-lg font-medium transition-colors ${
              activeTab === "paused"
                ? "bg-blue-600 text-white"
                : "bg-white text-gray-700 hover:bg-gray-50"
            }`}
          >
            <Pause className="w-5 h-5" />
            Produtos Pausados
          </button>
          <button
            onClick={() => setActiveTab("approved")}
            className={`flex items-center gap-2 px-6 py-3 rounded-lg font-medium transition-colors ${
              activeTab === "approved"
                ? "bg-blue-600 text-white"
                : "bg-white text-gray-700 hover:bg-gray-50"
            }`}
          >
            <Check className="w-5 h-5" />
            Produtos Aprovados
          </button>
          <button
            onClick={() => setActiveTab("daily-offer")}
            className={`flex items-center gap-2 px-6 py-3 rounded-lg font-medium transition-colors ${
              activeTab === "daily-offer"
                ? "bg-blue-600 text-white"
                : "bg-white text-gray-700 hover:bg-gray-50"
            }`}
          >
            <Tag className="w-5 h-5" />
            Ofertas do dia
          </button>
        </div>

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

        {/* Products Grid */}
        {products.length === 0 && (
          <div className="text-center py-12">
            <Package className="w-12 h-12 text-gray-400 mx-auto mb-4" />
            <h3 className="text-lg font-medium text-gray-900 mb-2">
              Nenhum produto encontrado
            </h3>
          </div>
        )}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {products.map((product) => (
            <div
              key={product.id}
              className="bg-white rounded-lg shadow-md overflow-hidden"
            >
              <img
                src={"data:image;base64," + product.image}
                alt={product.title}
                className="w-full h-48 object-cover"
              />
              <div className="p-4">
                <h3 className="text-lg font-semibold text-gray-900 mb-2">
                  {product.title}
                </h3>
                <p className="text-gray-600 text-sm mb-2">
                  {product.description}
                </p>
                <p className="text-blue-600 font-bold mb-4">
                  R${" "}
                  {new Intl.NumberFormat("pt-BR", {
                    style: "decimal",
                    minimumFractionDigits: 2,
                  }).format(product.price)}
                </p>

                <div className="flex items-center justify-between">
                  <div className="space-x-4">
                    {activeTab === "paused" && (
                      <label className="inline-flex items-center cursor-pointer">
                        <input
                          type="checkbox"
                          className="form-checkbox h-5 w-5 text-blue-600"
                          onChange={() =>
                            handleStatusChange(product, "approved")
                          }
                        />
                        <span className="ml-2 text-sm text-gray-700">
                          Aprovar
                        </span>
                      </label>
                    )}

                    {activeTab === "approved" && (
                      <>
                        <label className="inline-flex items-center cursor-pointer">
                          <input
                            type="checkbox"
                            className="form-checkbox h-5 w-5 text-blue-600"
                            onChange={() =>
                              handleStatusChange(product, "paused")
                            }
                          />
                          <span className="ml-2 text-sm text-gray-700">
                            Pausar
                          </span>
                        </label>
                        <label className="inline-flex items-center cursor-pointer">
                          <input
                            type="checkbox"
                            className="form-checkbox h-5 w-5 text-blue-600"
                            onChange={() =>
                              handleStatusChange(product, "daily-offer")
                            }
                          />
                          <span className="ml-2 text-sm text-gray-700">
                            Ofertas do dia
                          </span>
                        </label>
                      </>
                    )}

                    {activeTab === "daily-offer" && (
                      <label className="inline-flex items-center cursor-pointer">
                        <input
                          type="checkbox"
                          className="form-checkbox h-5 w-5 text-blue-600"
                          checked={product.approved}
                          onChange={() => handleStatusChange(product, "paused")}
                        />
                        <span className="ml-2 text-sm text-gray-700">
                          Aprovado
                        </span>
                      </label>
                    )}
                  </div>

                  <button
                    onClick={() => handleDelete(product.id)}
                    className="text-red-600 hover:bg-red-50 p-2 rounded-lg transition-colors"
                  >
                    <Trash2 className="w-5 h-5" />
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>

        {products.length === 0 && (
          <div className="text-center py-12">
            <Package className="w-12 h-12 text-gray-400 mx-auto mb-4" />
            <h3 className="text-lg font-medium text-gray-900 mb-2">
              Nenhum produto encontrado
            </h3>
            <p className="text-gray-600">
              Não há produtos{" "}
              {activeTab === "paused"
                ? "pausados"
                : activeTab === "approved"
                ? "aprovados"
                : "em oferta"}{" "}
              no momento.
            </p>
          </div>
        )}
      </div>
    </div>
  );
}

export default MarketplaceProducts;
