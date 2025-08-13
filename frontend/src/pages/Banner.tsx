import React, { useState } from "react";
import { Plus, X, Image, Upload, MapPin } from "lucide-react";
import { BannerService } from "../services/BannerService";
import { CityService } from "../services/CityService";
import { states, carouselSections } from "../data";
import { IBanner } from "../interfaces/IBanner";
import Swal from "sweetalert2";
import { RegionService, UploadFileService } from "../services";
import LoadingText from "../components/LoadingText";
import { IRegion } from "../interfaces/IRegion";

function Banner() {
  const [newImageUrl, setNewImageUrl] = useState("");
  const [newImageTitle, setNewImageTitle] = useState("");
  const carouselImages = [{image:""}];
  const [carouselImagesList, setCarouselImagesList] = useState(carouselImages);
  const [selectedCarouselSection, setSelectedCarouselSection] =
    useState("home");
  const [file, setFile] = useState<File | null>(null);
  const [page, setPage] = useState("H");
  const [previewUrl, setPreviewUrl] = useState("");
  const [updateImages, setUpdateImages] = useState(false);

  const [citiesByState, setcitiesByState] = useState([{}]);

  const [formData, setFormData] = useState({
    title: "",
    url: "",
    state: "",
    city: "",
    photo: "",
    region: "",
  });
  const [isLoading, setIsLoading] = useState(false);
  const [regions, setRegions] = useState([]);


  React.useEffect(() => {
    const fetchData = async () => {
        setIsLoading(true);
        const regions = await RegionService.getRegions(1000000,0,'');
        const json_regions = await regions.data;
        if ((regions.status == 200)) {
          setRegions(json_regions.regions);
        }
        setIsLoading(false);
    };
    fetchData();
  }, []);



  React.useEffect(() => {
    const fetchData = async () => {
      const id = parseInt(localStorage.getItem("id") || "");
      const type = carouselSections.find(
        (s) => s.id === selectedCarouselSection
      )?.type;
      const response = await BannerService.getBannerByPage({
        page: page,
        cityId: 0,
        regionId: 0,
      });
      const json = response.data;
      setCarouselImagesList(json);
    };

    fetchData();
  }, [page, updateImages]);


  const handleRemoveCarouselImage = async (section: string, index: number) => {
    if (
      window.confirm("Tem certeza que deseja remover esta imagem do carrossel?")
    ) {

      const response = await BannerService.deleteBanner(index);
      if ((response.status == 200)) {
        setUpdateImages(!updateImages);
      }
    }
  };

  const handleAddCarouselImage = async (e: React.FormEvent) => {
    setIsLoading(true);
    e.preventDefault();
    //console.info(formData.photo);
    //console.info("e", e.target[0].files[0]);

    const formDataAPI = new FormData();
    formDataAPI.append("myFile", {
      uri: "file://", //Your Image File Path
      type: "image/jpeg",
      name: "imagename.jpg",
    });

    try {
    const response_image = await UploadFileService.uploadImage(
      e.target[0].files[0]
    );
    if (response_image.data.uri!=="") {
      const banner: IBanner = {
        bannerId: 0,
        image: btoa(response_image.data.uri),
        title: "", //(e.target as any).imageTitle.value,
        accessLink: (e.target as any).imageUrl.value,
        page: page,
        oid: 0,
        cityId: page == "C" ? parseInt(formData.city) : null,
        regionId: page == "U" || page == "B" ? parseInt(formData.region) : null,
      };
    
      const result = await BannerService.saveBanner(banner);
      if (result.status == 200) {
        Swal.fire({
          position: "center",
          icon: "success",
          title: "Sucesso",
          text: "Banner incluído!!",
          showConfirmButton: false,
          timer: 1500,
        });
        setPreviewUrl("");
        setNewImageUrl("");
        setUpdateImages(!updateImages);
        setIsLoading(false);
      }
      //setNewImageUrl("");
      //setNewImageTitle("");
    }
    else {
      Swal.fire({
        position: "center",
        icon: "error",
        title: "Erro ao fazer upload da imagem - url",
        text: "Por favor, tente novamente.",
        showConfirmButton: true,
      });
      setIsLoading(false);
  
    }
    } catch (error) {
    console.error("Error uploading image:", error);
    Swal.fire({
      position: "center",
      icon: "error",
      title: "Erro ao fazer upload da imagem",
      text: "Por favor, tente novamente. Error:"+ error,
      showConfirmButton: true,
    });
    setIsLoading(false);
  }
  };


  const handleChange = async (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
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
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));

    setFormData((prev) => ({
      ...prev,
      [name]: value,
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

      //setFormData((prev) => ({ ...prev, photo: URL.createObjectURL(file) }));
    }
  };

  if (isLoading) {
    return (
      <>
        <LoadingText></LoadingText>
      </>
    );
  }


  return (
    <div>
      {/* Carousel Image Management */}
      <div className="bg-white rounded-lg shadow-md p-6">
        <div className="flex items-center gap-3 mb-6">
          <Image className="w-6 h-6 text-blue-600" />
          <h2 className="text-xl font-bold text-gray-900">
            Gerenciar Imagens do Carrossel
          </h2>
        </div>

        <div className="space-y-6">
          {/* Section Selection */}
          <div className="flex flex-wrap gap-2">
            {carouselSections.map((section) => (
              <button
                key={section.id}
                onClick={() => {
                  setSelectedCarouselSection(section.id);
                  //fetchData(section.type);
                  setPage(section.type);
                }}
                className={`px-4 py-2 rounded-lg transition-colors ${
                  selectedCarouselSection === section.id
                    ? "bg-blue-600 text-white"
                    : "bg-gray-100 text-gray-700 hover:bg-gray-200"
                }`}
              >
                {section.label}
              </button>
            ))}
          </div>
        </div>
      </div>
      {/* Add New Image Form */}

      <form
        onSubmit={handleAddCarouselImage}
        className="space-y-4 border-t border-gray-200 pt-6"
      >
        {/* Photo Upload */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Nova Imagem
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
              <span>Imagens com no máximo 1mb de tamanho</span>

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
        {/*
        <div>
          <label
            htmlFor="imageTitle"
            className="block text-sm font-medium text-gray-700 mb-1"
          >
            Título da Imagem
          </label>
          <input
            type="text"
            id="imageTitle"
            name="imageTitle"
            //value={formData.title}
            onChange={handleChange}
            placeholder="Digite o título da imagem"
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            //required
          />
        </div>
*/}
        <div>
          <label
            htmlFor="imageTitle"
            className="block text-sm font-medium text-gray-700 mb-1"
          >
            Link da imagem
          </label>
          <input
            type="text"
            id="imageUrl"
            name="imageUrl"
            //value={formData.url}
            //onChange={(e) => setNewImageTitle(e.target.value)}
            placeholder="Digite o link da imagem"
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            //required
          />
        </div>
        {/* Estado e Cidade */}
        {page == "C" && (
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
        )}

        {/* Região */}
        {(page == "U" || page == "B") && (
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label
                htmlFor="city"
                className="block text-sm font-medium text-gray-700"
              >
                Região
              </label>
              <div className="mt-1 relative rounded-md shadow-sm">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <MapPin className="h-5 w-5 text-gray-400" />
                </div>
                <select
                  id="region"
                  name="region"
                  required
                  value={formData.region}
                  onChange={handleChange}
                  className="appearance-none block w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100"
                >
                  <option value="">Selecione a região</option>
                  {regions.map((region: IRegion) => (
                    <option key={region.id} value={region.id}>
                      {region.name}
                    </option>
                  ))}
                </select>
              </div>
            </div>
          </div>
        )}

        <button
          type="submit"
          className="flex items-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
        >
          <Plus className="w-5 h-5" />
          Adicionar Imagem
        </button>
      </form>

      {/* Current Images */}

      <div className="border-t border-gray-200 pt-6">
        <h3 className="text-lg font-semibold text-gray-900 mb-4">
          Imagens Atuais -{" "}
          {
            carouselSections.find((s) => s.id === selectedCarouselSection)
              ?.label
          }
        </h3>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {carouselImagesList.map((image, index) => (
            <div key={index} className="relative group">
              <img
                src={atob(image.image)}
                alt={image.title}
                className="w-full h-48 object-cover rounded-lg"
              />
              <div className="absolute inset-0 bg-black bg-opacity-50 opacity-0 group-hover:opacity-100 transition-opacity rounded-lg flex items-center justify-center">
                <button
                  onClick={() =>
                    handleRemoveCarouselImage(selectedCarouselSection, image.id)
                  }
                  className="p-2 bg-red-600 text-white rounded-full hover:bg-red-700 transition-colors"
                >
                  <X className="w-5 h-5" />
                </button>
              </div>
              {page == "C" && (
                <p className="mt-2 text-sm text-gray-600 truncate">
                  {image.cidade?.nome + "/" + image.cidade?.uf}
                </p>
              )}
              {(page == "U" || page == "B") && (
                <p className="mt-2 text-sm text-gray-600 truncate">
                  {image.region?.nome }
                </p>
              )}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
export default Banner;
