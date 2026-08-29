package plan

import (
	"gorm.io/gorm"
)

// SeedPlans popula a tabela de planos com os dados iniciais
func SeedPlans(db *gorm.DB) error {
	plans := []Plan{
		{
			UserType:     UserTypeProfessional,
			Name:         "Premium Profissional",
			Price:        9.90,
			Description:  "Plano premium para profissionais da construção civil",
			Features:     `["Destaque nos resultados de busca","Badge Premium no perfil","Acesso ilimitado a orçamentos","Suporte prioritário","Verificação profissional"]`,
			IsActive:     true,
			DurationDays: 30,
		},
		{
			UserType:     UserTypeSolicitation,
			Name:         "Taxa de Solicitação",
			Price:        9.90,
			Description:  "Taxa de deslocamento cobrada por solicitação de profissional no app",
			Features:     `[]`,
			IsActive:     true,
			DurationDays: 0, // cobrança avulsa, não expira
		},
		{
			UserType:     UserTypeStore,
			Name:         "Premium Lojista",
			Price:        19.90,
			Description:  "Plano premium para lojistas parceiros",
			Features:     `["Destaque na lista de fornecedores","Badge Premium na loja","Produtos em destaque","Suporte prioritário","Verificação comercial"]`,
			IsActive:     true,
			DurationDays: 30,
		},
	}

	for _, plan := range plans {
		// Verifica se o plano já existe
		var existingPlan Plan
		result := db.Where("user_type = ?", plan.UserType).First(&existingPlan)

		if result.Error == gorm.ErrRecordNotFound {
			// Plano não existe, cria um novo
			if err := db.Create(&plan).Error; err != nil {
				return err
			}
			continue
		}

		if result.Error != nil {
			return result.Error
		}

		// O plano já existe: preço, nome e descrição são gerenciados direto no
		// banco, então não são sobrescritos aqui. Só preenchemos a vigência das
		// linhas semeadas antes da coluna duration_days existir.
		if existingPlan.DurationDays == 0 && plan.DurationDays > 0 {
			if err := db.Model(&existingPlan).Update("duration_days", plan.DurationDays).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
