package product_usecase

import pkgproduct "construir_mais_barato/app/domain/product"

func GenerateProduct(assembler *ProductAssembler) pkgproduct.Product {
	product := pkgproduct.Product{}
	if assembler != nil {
		product.ID = assembler.ID
		product.Name = assembler.Name
		product.Description = assembler.Description
		product.Image = assembler.Image
		product.Price = assembler.Price
		product.Discount = assembler.Discount
		product.OriginalPrice = assembler.OriginalPrice
		product.Approved = assembler.Approved
		product.Dayoffer = assembler.Dayoffer
		product.ProfessionID = assembler.ProfessionID
		product.CategoryID = assembler.CategoryID
		product.Category = assembler.Category
		product.ProfessionalID = assembler.ProfessionalID
		product.Professional = assembler.Professional
		product.StoreID = assembler.StoreID
		product.Store = assembler.Store

	}
	return product
}

func GenerateProductPresenter(product *pkgproduct.Product) ProductPresenter {

	presenter := ProductPresenter{}

	presenter.ID = product.ID
	presenter.Name = product.Name
	presenter.Description = product.Description
	presenter.Image = product.Image
	presenter.Price = product.Price
	presenter.Discount = product.Discount
	presenter.OriginalPrice = product.OriginalPrice
	presenter.Approved = product.Approved
	presenter.Dayoffer = product.Dayoffer
	presenter.ProfessionID = product.ProfessionID
	presenter.CategoryID = product.CategoryID
	presenter.Category = product.Category
	presenter.ProfessionalID = product.ProfessionalID
	presenter.Professional = product.Professional
	presenter.StoreID = product.StoreID
	presenter.Store = product.Store

	return presenter
}

func GenerateProductsPresenter(products []*pkgproduct.Product) *[]ProductPresenter {
	list := make([]ProductPresenter, 0)
	if products != nil && len(products) > 0 {
		for _, product := range products {

			//professions := getProfessionsPresenter(professional.Professions)

			presenter := ProductPresenter{}

			presenter.ID = product.ID
			presenter.Name = product.Name
			presenter.Description = product.Description
			presenter.Image = product.Image
			presenter.Price = product.Price
			presenter.OriginalPrice = product.OriginalPrice
			presenter.Discount = product.Discount
			presenter.Approved = product.Approved
			presenter.Dayoffer = product.Dayoffer
			presenter.ProfessionID = product.ProfessionID
			presenter.CategoryID = product.CategoryID
			presenter.Category = product.Category
			presenter.ProfessionalID = product.ProfessionalID
			presenter.Professional = product.Professional
			presenter.StoreID = product.StoreID
			presenter.Store = product.Store

			list = append(list, presenter)
		}
	}

	return &list
}
