import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import { BannerService } from "../services/BannerService";
import { RegionService } from "../services";
import { IBannerSearchProfessionals } from "../interfaces";
import CarouselNovo from "./CarouselNovo";

// URL_IMAGES_WEB do .env
const URL_IMAGES_WEB = import.meta.env.VITE_URL_IMAGES_WEB;

// Monta a URL da imagem a partir do path que o backend devolve em base64.
function getImageUrl(encodedPath: string): string {
  try {
    if (!encodedPath || encodedPath.trim() === "") return "";

    // decodifica o Base64 para obter o path real
    const decodedPath = atob(encodedPath); // ex: "/images/upload/upload-3341225764.png"

    if (!decodedPath || decodedPath.trim() === "") return "";

    // Alguns banners ja guardam a URL completa (mesmo caso do CarouselNovo)
    if (
      decodedPath.startsWith("http://") ||
      decodedPath.startsWith("https://")
    ) {
      return decodedPath;
    }

    const baseUrl = URL_IMAGES_WEB?.replace(/\/$/, ""); // remove barra final

    if (!baseUrl) return "";

    return `${baseUrl}${decodedPath.startsWith("/") ? "" : "/"}${decodedPath}`;
  } catch (error) {
    console.error("Erro ao decodificar path da imagem:", error);
    return "";
  }
}

// Evita exibir um espaco em branco quando o arquivo do banner nao existe mais.
function validateImageExists(imageUrl: string): Promise<boolean> {
  return new Promise((resolve) => {
    if (!imageUrl) {
      resolve(false);
      return;
    }

    const img = new Image();
    img.onload = () => resolve(true);
    img.onerror = () => resolve(false);
    img.src = imageUrl;
  });
}

function gerarNumeroAleatorio(min: number, max: number): number {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

/**
 * Banner da pagina de resultados da busca.
 *
 * Antes este banner (page "B", da regiao da cidade escolhida) era exibido numa
 * modal que o cliente precisava fechar antes de ver os profissionais. Agora ele
 * aparece como banner unico no topo da pagina, economizando um clique.
 *
 * Quando a regiao nao tem banner "B" cadastrado, cai no carrossel padrao para
 * nao deixar o espaco publicitario vazio.
 */
function SearchResultsBanner() {
  const location = useLocation();
  const selectedCity = (location.state as { selectedCity?: string } | null)
    ?.selectedCity;
  const [banner, setBanner] = useState<IBannerSearchProfessionals | null>(null);
  const [imageUrl, setImageUrl] = useState<string>("");
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    // Evita atualizar o estado se a cidade mudar antes da resposta chegar.
    let isCurrent = true;

    const fetchBanner = async () => {
      setIsLoading(true);
      setBanner(null);
      setImageUrl("");

      if (!selectedCity) {
        if (isCurrent) setIsLoading(false);
        return;
      }

      try {
        const regionRes = await RegionService.getRegionbyCity(
          parseInt(selectedCity)
        );
        if (regionRes.status !== 200) {
          if (isCurrent) setIsLoading(false);
          return;
        }

        const bannerRes = await BannerService.getBannerByPagePublic({
          page: "B",
          cityId: 0,
          regionId: regionRes.data.id,
        });

        if (bannerRes.status === 200 && bannerRes.data.length > 0) {
          const randomIndex = gerarNumeroAleatorio(0, bannerRes.data.length - 1);
          const selectedBanner = bannerRes.data[randomIndex];
          const url = getImageUrl(selectedBanner.image);

          if (url && (await validateImageExists(url)) && isCurrent) {
            setBanner(selectedBanner);
            setImageUrl(url);
          }
        }
      } catch (error) {
        console.error("Erro ao buscar banner da busca:", error);
      } finally {
        if (isCurrent) setIsLoading(false);
      }
    };

    fetchBanner();

    return () => {
      isCurrent = false;
    };
  }, [selectedCity]);

  if (isLoading) return null;

  if (!banner || !imageUrl) {
    return <CarouselNovo page="/search-results" />;
  }

  return (
    <div className="max-w-[1400px] mx-auto px-4">
      <img
        src={imageUrl}
        alt="Imagem de Construção"
        onClick={
          banner.link ? () => window.open(banner.link, "_blank") : undefined
        }
        className={`w-full h-44 sm:h-64 md:h-80 lg:h-96 xl:h-[420px] object-contain md:object-cover md:rounded-lg transition-transform duration-300 hover:scale-105 ${
          banner.link ? "cursor-pointer" : ""
        }`}
      />
    </div>
  );
}

export default SearchResultsBanner;
