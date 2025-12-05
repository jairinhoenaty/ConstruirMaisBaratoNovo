export interface ICategoryProduct {
  id: number;
  name: string;
  professional_id?: number;
  parent_id?: number | null;
  children?: ICategoryProduct[];
}

// Alias para compatibilidade
export interface IProductCategory extends ICategoryProduct {}

// Interface simplificada para uso em dropdowns/selects
export interface ICategoryProductSimple {
  id: number;
  name: string;
  parent_id?: number | null;
}

export interface ICategoryAndSubCategories{
  categoryId:number;
  subCategoriesId: number[] | null;
}
