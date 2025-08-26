import React, { useState } from "react";
import { Upload, Tag, DollarSign, FileText, Package, Info } from "lucide-react";
import {
  ProductCategoryService,
  ProductService,
  ProfessionService,
} from "../services";
import { IProduct } from "../interfaces/IProduct";
import CurrencyInput from "react-currency-input-field";

function NewProduct() {
  const [formData, setFormData] = useState({
    title: "",
    mainCategory: "",
    subcategory: "",
    price: "",
    description: "",
    photo: "",
  });
  const [previewUrl, setPreviewUrl] = useState("");
  const [professions, setProfessions] = useState([{}]);
  const [subcategories, setSubcategories] = useState([{}]);

  const handleChange = async (
    e: React.ChangeEvent<
      HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement
    >
  ) => {
    const { name, value } = e.target;

    if (name == "mainCategory") {
      const result = await ProductCategoryService.productCategoriesByProfession(
        parseInt(value)
      );
      if (result.status == 200) {
        const json_categories = await result.data;

        setSubcategories(json_categories);
      }
    }
    setFormData((prev) => ({
      ...prev,
      [name]: value,
      ...(name === "mainCategory" ? { subcategory: "" } : {}),
    }));
  };

  React.useEffect(() => {
    const fetchData = async () => {
      const result = await ProfessionService.getProfessionsPublic();
      const json_professions = await result.data;
      if (result.status == 200) {
        setProfessions(json_professions);
      }
    };
    fetchData();
  }),
    [];

  const handleChangePrice = (value: any) => {
    setFormData((prev) => ({
      ...prev,
      price: value,
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

      setFormData((prev) => ({ ...prev, photo: URL.createObjectURL(file) }));
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const postID = parseInt(localStorage.getItem("post_id") ?? "");
    const profile = localStorage.getItem("profile");
    let base64image = previewUrl;
    base64image = base64image
      .replace("data:image/png;base64,", "")
      .replace("data:image/jpg;base64,", "")
      .replace("data:image/jpeg;base64,", "");
    const product: IProduct = {
      id: 0,
      name: formData.title,
      description: formData.description,
      image: base64image,
      price: parseInt(formData.price),
      originalprice: 0,
      discount: 0,
      approved: false,
      dayoffer: false,
      professionID: 0,
      categoryID: parseInt(formData.subcategory),
      professionalID: profile == "profissional" ? postID : null,
      storeID: profile == "store" ? postID : null,
    };

    const result = await ProductService.saveProduct(product);
    if (result.status == 200) {
      alert("Produto cadastrado com sucesso!");
      setFormData((prev) => ({
        ...prev,
        title: "",
        mainCategory: "",
        subcategory: "",
        price: "",
        description: "",
        photo: "",
      }));
    }
  };

  return (
    <div className="bg-white rounded-lg shadow-md p-6">
      <div className="flex items-center gap-3 mb-6">
        <Package className="w-6 h-6 text-blue-600" />
        <h2 className="text-xl font-bold text-gray-900">Novo Produto</h2>
      </div>

      <form onSubmit={handleSubmit} className="space-y-6">
        {/* Photo Upload */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Foto do Produto
          </label>
          <div className="flex items-center justify-center">
            <div className="relative">
              <div
                className={`w-full h-64 rounded-lg overflow-hidden border-2 border-gray-300 flex items-center justify-center bg-gray-50 ${
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
                  <Upload className="w-8 h-8 text-gray-400" />
                )}
              </div>
              <label
                htmlFor="photo-upload"
                className="absolute bottom-4 right-4 bg-blue-600 text-white p-2 rounded-full cursor-pointer hover:bg-blue-700 transition-colors"
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
            </div>
          </div>
        </div>

        {/* Title */}
        <div>
          <label
            htmlFor="title"
            className="block text-sm font-medium text-gray-700 mb-1"
          >
            Título do Produto
          </label>
          <div className="relative">
            <Package className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
            <input
              type="text"
              id="title"
              name="title"
              value={formData.title}
              onChange={handleChange}
              required
              className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="Ex: Kit Ferramentas Profissional"
            />
          </div>
        </div>

        {/* Main Category */}
        <div>
          <label
            htmlFor="mainCategory"
            className="block text-sm font-medium text-gray-700 mb-1"
          >
            Categoria Principal
          </label>
          <div className="relative">
            <Tag className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
            <select
              id="mainCategory"
              name="mainCategory"
              value={formData.mainCategory}
              onChange={handleChange}
              required
              className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none"
            >
              <option value="">Selecione a categoria principal</option>
              {professions.map((category) => (
                <option key={category.id} value={category.id}>
                  {"Produtos para " + category.name}
                </option>
              ))}
            </select>
          </div>
        </div>

        {/* Subcategory */}
        <div>
          <label
            htmlFor="subcategory"
            className="block text-sm font-medium text-gray-700 mb-1"
          >
            Subcategoria
          </label>
          <div className="relative">
            <Tag className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
            <select
              id="subcategory"
              name="subcategory"
              value={formData.subcategory}
              onChange={handleChange}
              required
              disabled={!formData.mainCategory}
              className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none disabled:bg-gray-100"
            >
              <option value="">Selecione a subcategoria</option>
              {subcategories.map((subcategory) => (
                <option key={subcategory.id} value={subcategory.id}>
                  {subcategory.name}
                </option>
              ))}
            </select>
          </div>
        </div>

        {/* Price */}
        <div>
          <label
            htmlFor="price"
            className="block text-sm font-medium text-gray-700 mb-1"
          >
            Preço
          </label>
          <div className="relative">
            <DollarSign className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" />
            <CurrencyInput
              //mask="999999.99"
              //intlConfig={{ locale: 'pt-BR', currency: 'BRL' }}
              decimalsLimit={2}
              decimalScale={2}
              allowDecimals={true}
              decimalSeparator={","}
              groupSeparator={"."}
              //type="number"
              id="price"
              name="price"
              value={formData.price}
              //onChange={handleChange}
              onValueChange={(value, name, values) => handleChangePrice(value)} //(value, name, values) => console.log(value, name, values)}
              required
              className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="Entre com o preço"
            />
          </div>
        </div>

        {/* Description */}
        <div>
          <label
            htmlFor="description"
            className="block text-sm font-medium text-gray-700 mb-1"
          >
            Descrição do Produto
          </label>
          <div className="relative">
            <FileText className="absolute left-3 top-3 text-gray-400 w-5 h-5" />
            <textarea
              id="description"
              name="description"
              value={formData.description}
              onChange={handleChange}
              required
              rows={4}
              className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="Descreva as características, especificações e diferenciais do produto"
            />
          </div>
        </div>

        <div className="flex items-start gap-2 p-4 bg-yellow-50 rounded-lg">
          <Info className="w-5 h-5 text-yellow-600 flex-shrink-0 mt-0.5" />
          <p className="text-sm text-yellow-700">
            Ao cadastrar um produto, você concorda com os termos de uso e
            política de privacidade da plataforma. Todas as informações
            fornecidas são de sua responsabilidade.
          </p>
        </div>

        <button
          type="submit"
          className="w-full bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 transition-colors"
        >
          Cadastrar Produto
        </button>
      </form>
    </div>
  );
}

export default NewProduct;
