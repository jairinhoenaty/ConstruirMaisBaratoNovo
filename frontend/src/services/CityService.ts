import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

const citiesByState = (data: any) => Api.post('/city-uf', data);
const citiesByStatePublic = (data:any) => ApiPublica.post('/cities-by-state',data);

export const CityService= {
   citiesByState,
   citiesByStatePublic
}