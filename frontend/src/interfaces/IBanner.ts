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


export interface Cidade {
  nome: string;
  uf: string;
}

export interface Region {
  oid: number;
  nome: string;
  uf: string;
}

export interface IBannerSearchProfessionals {
  id: number;
  link: string;
  cidade: Cidade;
  professions: string[] | null;      // pode ser lista ou null
  professionsIds: number[] | null;   // pode ser lista ou null
  page: string;
  image: string;                     // caminho em base64 (não a imagem em si)
  region: Region;
}