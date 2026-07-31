import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

const citiesByState = (data: any) => Api.post('/city-uf', data);
const citiesByStatePublic = (data:any) => ApiPublica.post('/cities-by-state',data);

// Autocomplete de cidades: busca por parte do nome e retorna { id, name, uf }.
const searchCitiesPublic = (term: string, limit = 8, signal?: AbortSignal) =>
   ApiPublica.get('/cities/search', { params: { q: term, limit }, signal });

export const CityService= {
   citiesByState,
   citiesByStatePublic,
   searchCitiesPublic
}