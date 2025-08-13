package domain_test

import (
	"testing"

	pkgcontact "construir_mais_barato/app/domain/contact"
	pkgmock "construir_mais_barato/tests/app_test/domain_test/mocks"

	"github.com/stretchr/testify/assert"
)

func TestContactService(t *testing.T) {

	mockRepo := new(pkgmock.MockContactRepository)
	contactService := pkgcontact.NewContactService(mockRepo)

	t.Run("Test method findAll", func(t *testing.T) {
		mockContacts := []*pkgcontact.Contact{
			{Name: "Alfredinho", Telephone: "12999987899", Email: "alfredo@teste.com.br", Message: "Testando contato 1"},
			{Name: "Mariazinha", Telephone: "12999987878", Email: "mariazinha@teste.com.br", Message: "Testando contato 2"},
		}

		mockRepo.On("FindAll").Return(mockContacts, nil)

		contacts, err := contactService.FindAll()

		assert.NoError(t, err)
		assert.Equal(t, mockContacts, contacts)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test method findById", func(t *testing.T) {
		mockContact := &pkgcontact.Contact{
			Name: "Alfredinho", Telephone: "12999987899", Email: "alfredo@teste.com.br", Message: "Testando contato 1",
		}

		mockRepo.On("FindById", uint(1)).Return(mockContact, nil)

		contact, err := contactService.FindById(1)

		assert.NoError(t, err)
		assert.Equal(t, mockContact, contact)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test method Save(insert)", func(t *testing.T) {
		newContact := pkgcontact.Contact{
			Name: "Alfredinho", Telephone: "12999987899", Email: "alfredo@teste.com.br", Message: "Testando contato 1",
		}
		savedContact := &pkgcontact.Contact{
			Name: "Alfredinho", Telephone: "12999987899", Email: "alfredo@teste.com.br", Message: "Testando contato 1",
		}

		mockRepo.On("Save", newContact).Return(savedContact, nil)

		contact, err := contactService.Save(newContact)

		assert.NoError(t, err)
		assert.Equal(t, savedContact, contact)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test method Remove", func(t *testing.T) {
		mockRepo.On("Remove", uint(1)).Return(nil)

		err := contactService.Remove(1)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

}
