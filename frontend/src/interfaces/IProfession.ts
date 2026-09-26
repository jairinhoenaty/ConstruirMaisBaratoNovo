export interface IProfession {
  id: number;
  name: string;
  description: string;
  icon: string;
  category_id?: number | null;
  /** Já resolvido pelo backend: override da profissão ou padrão da categoria. */
  client_travels?: boolean;
  /** Nulo significa "herda da categoria"; só a administração usa este campo. */
  client_travels_override?: boolean | null;
}

export interface IProfessionSearchProfessionals {
  id: number;
  name: string;
  icon: string;
  description: string;
}
