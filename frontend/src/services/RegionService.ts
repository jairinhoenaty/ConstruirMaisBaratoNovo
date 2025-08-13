import { IRegion } from "../interfaces/IRegion";
import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

const getRegions = (limit: number, offset: number,uf:string) =>
  Api.get("/regions?limit=" + limit + "&offset=" + offset+"&uf=" + uf);
const getRegionsPublic = () => ApiPublica.get("/regions/all");
const postRegion = (data: IRegion) => Api.post("/region", data);
const getRegionbyID = (id: number) => Api.get("/region/" + id);
const deleteRegion = (id: number) => Api.delete("/region/" + id);
const getRegionbyCity = (id: number) =>
  ApiPublica.get("/regions/findbycity?cityId=" + id);

export const RegionService = {
  getRegions,
  getRegionsPublic,
  postRegion,
  getRegionbyID,
  deleteRegion,
  getRegionbyCity,
};
