export interface IProfessionCategory {
  id: number;
  name: string;
  description: string;
  icon: string;
  position: number;
  active: boolean;
  /** Quando true, é o cliente quem se desloca até o profissional. */
  client_travels: boolean;
  profession_count: number;
}

export interface IProfessionCategoryWithProfessions extends IProfessionCategory {
  professions: IProfessionOfCategory[];
}

export interface IProfessionOfCategory {
  id: number;
  name: string;
  description: string;
  icon: string;
  category_id?: number;
  /** Já resolvido pelo backend: override da profissão ou padrão da categoria. */
  client_travels: boolean;
  /** Nulo significa "herda da categoria"; só a administração usa este campo. */
  client_travels_override: boolean | null;
}
