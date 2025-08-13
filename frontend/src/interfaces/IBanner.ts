export interface IBanner {
  oid: number;
  image: string;
  title: string;
  page: string;
  accessLink: string;
  bannerId: number;
  cityId: number | null;
  regionId: number | null;
}
