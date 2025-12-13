package unlockprice

import (
	"log"

	"gorm.io/gorm"
)

func SeedUnlockPrices(db *gorm.DB) {
	var count int64
	db.Model(&UnlockPrice{}).Count(&count)

	if count > 0 {
		log.Println("UnlockPrices already seeded")
		return
	}

	prices := []UnlockPrice{
		{
			UserType:    UserTypeProfessional,
			Price:       30.00,
			Description: "Desbloqueio avulso de orçamento para profissionais",
			IsActive:    true,
		},
		{
			UserType:    UserTypeStore,
			Price:       35.00,
			Description: "Desbloqueio avulso de orçamento para lojistas",
			IsActive:    true,
		},
	}

	for _, price := range prices {
		if err := db.Create(&price).Error; err != nil {
			log.Printf("Error seeding unlock price for %s: %v", price.UserType, err)
		}
	}

	log.Println("UnlockPrices seeded successfully")
}
