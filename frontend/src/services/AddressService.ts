export interface CepAddress {
  street: string;
  neighborhood: string;
  city: string;
  uf: string;
}

export interface Coordinates {
  latitude: number;
  longitude: number;
}

/**
 * ViaCEP é público e sem chave, diferente da API do Google, cujo billing está
 * desativado e responde REQUEST_DENIED.
 */
const lookupCep = async (cep: string): Promise<CepAddress | null> => {
  const digits = cep.replace(/\D/g, "");
  if (digits.length !== 8) return null;

  try {
    const response = await fetch(`https://viacep.com.br/ws/${digits}/json/`);
    if (!response.ok) return null;

    const data = await response.json();
    if (data.erro) return null;

    return {
      street: data.logradouro ?? "",
      neighborhood: data.bairro ?? "",
      city: data.localidade ?? "",
      uf: data.uf ?? "",
    };
  } catch (error) {
    console.error("Erro ao consultar o CEP:", error);
    return null;
  }
};

/**
 * O navegador não tem geocoder nativo como o aparelho, então o Nominatim é a
 * única via — no app ele é o fallback do geocoder do sistema.
 */
const geocodeAddress = async (address: string): Promise<Coordinates | null> => {
  if (!address.trim()) return null;

  try {
    const response = await fetch(
      "https://nominatim.openstreetmap.org/search" +
        `?q=${encodeURIComponent(address)}` +
        "&format=json&limit=1&countrycodes=br"
    );
    if (!response.ok) return null;

    const results = await response.json();
    if (!Array.isArray(results) || results.length === 0) return null;

    const latitude = parseFloat(results[0].lat);
    const longitude = parseFloat(results[0].lon);
    if (Number.isNaN(latitude) || Number.isNaN(longitude)) return null;

    return { latitude, longitude };
  } catch (error) {
    console.error("Erro ao geocodificar o endereço:", error);
    return null;
  }
};

export const AddressService = {
  lookupCep,
  geocodeAddress,
};
