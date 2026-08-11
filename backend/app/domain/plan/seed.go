package plan

import (
	"gorm.io/gorm"
)

// SeedPlans popula a tabela de planos com os dados iniciais
func SeedPlans(db *gorm.DB) error {
	plans := []Plan{
		{
			UserType:    UserTypeProfessional,
			Name:        "Premium Profissional",
			Price:       9.90,
			Description: "Plano premium para profissionais da construção civil",
			Features:    `["Destaque nos resultados de busca","Badge Premium no perfil","Acesso ilimitado a orçamentos","Suporte prioritário","Verificação profissional"]`,
			IsActive:    true,
		},
		{
			UserType:    UserTypeSolicitation,
			Name:        "Taxa de Solicitação",
			Price:       9.90,
			Description: "Taxa de deslocamento cobrada por solicitação de profissional no app",
			Features:    `[]`,
			IsActive:    true,
		},
		{
			UserType:    UserTypeStore,
			Name:        "Premium Lojista",
			Price:       19.90,
			Description: "Plano premium para lojistas parceiros",
			Features:    `["Destaque na lista de fornecedores","Badge Premium na loja","Produtos em destaque","Suporte prioritário","Verificação comercial"]`,
			IsActive:    true,
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
		}
		// Se já existe, não faz nada (mantém o plano existente)
	}

	return nil
}
