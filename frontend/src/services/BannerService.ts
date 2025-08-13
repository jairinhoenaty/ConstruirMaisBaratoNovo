//import { IBanner } from "../interfaces/IBanner";
import { IBanner } from "../interfaces/IBanner";
import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

const getBannerById = (id: number) => Api.get("/banner/" + id);
const getBannerByPage = (data: { page: string; cityId: number; regionId:number }) =>
  Api.post("/banners/page", data);
const saveBanner = (data: IBanner) => Api.post("/banner", data);
const deleteBanner = (id: any) => Api.delete("/banner/" + id);
const getBannerByPagePublic = (data: {
  page: string;
  cityId: number;
  regionId: number;
}) => ApiPublica.post("/banners/page", data);

export const BannerService = {
  getBannerById,
  saveBanner,
  deleteBanner,
  getBannerByPage,
  getBannerByPagePublic,
};
