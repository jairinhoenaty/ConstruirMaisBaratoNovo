package domain_test

// import (
// 	"testing"

// 	pkgprofession "construir_mais_barato/app/domain/profession"
// 	pkgmock "construir_mais_barato/tests/app_test/domain_test/mocks"

// 	"github.com/stretchr/testify/assert"
// )

// func TestProfessionService(t *testing.T) {

// 	mockRepo := new(pkgmock.MockProfessionRepository)
// 	professionService := pkgprofession.NewProfessionService(mockRepo)

// 	// t.Run("Test method findAll", func(t *testing.T) {
// 	// 	mockProfessions := []*pkgprofession.Profession{
// 	// 		{Name: "Arquiteto", Description: "Descrição da profissão de arquiteto", Icon: "icon.png"},
// 	// 		{Name: "Engenheiro", Description: "Descrição da profissão de engenheiro", Icon: "icon.png"},
// 	// 	}

// 	// 	mockRepo.On("FindAll").Return(mockProfessions, nil)

// 	// 	professions, _, err := professionService.FindAll()

// 	// 	assert.NoError(t, err)
// 	// 	assert.Equal(t, mockProfessions, professions)
// 	// 	mockRepo.AssertExpectations(t)
// 	// })

// 	t.Run("Test method findById", func(t *testing.T) {
// 		mockProfession := &pkgprofession.Profession{
// 			Name: "Arquiteto", Description: "Descrição da profissão de arquiteto", Icon: "icon.png",
// 		}

// 		mockRepo.On("FindById", uint(1)).Return(mockProfession, nil)

// 		profession, err := professionService.FindById(1)

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockProfession, profession)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("Test method Save (insert)", func(t *testing.T) {
// 		newProfession := pkgprofession.Profession{
// 			Name: "Arquiteto", Description: "Descrição da profissão de arquiteto", Icon: "icon.png",
// 		}
// 		savedProfession := &pkgprofession.Profession{
// 			Name: "Arquiteto", Description: "Descrição da profissão de arquiteto", Icon: "icon.png",
// 		}

// 		mockRepo.On("Save", newProfession).Return(savedProfession, nil)

// 		profession, err := professionService.Save(newProfession)

// 		assert.NoError(t, err)
// 		assert.Equal(t, savedProfession, profession)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("Test method Remove", func(t *testing.T) {
// 		mockRepo.On("Remove", uint(1)).Return(nil)

// 		err := professionService.Remove(1)

// 		assert.NoError(t, err)
// 		mockRepo.AssertExpectations(t)
// 	})

// }
