export interface IJob {
  id?: number;

  cargo: string;
  contratacao: string;
  salario: string;
  local: string;
  descricao: string;
  horario: string;
  requisitos: string;
  beneficios: string;
  contato: string;
  empresa: string;
  quantidade_vagas: number;

  created_at?: string;
}