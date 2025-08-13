package profession_repository_impl

import (
	"gorm.io/gorm"

	pkgprofession "construir_mais_barato/app/domain/profession"
)

type repository struct {
	DB *gorm.DB
}

func NewProfessionRepositoryImpl(db *gorm.DB) pkgprofession.ProfessionRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindAllWithoutPagination() ([]*pkgprofession.Profession, error) {
	var professions []*pkgprofession.Profession
	if err := r.DB.Order("name").Find(&professions).Error; err != nil {
		return nil, err
	}
	return professions, nil
}

func (r *repository) Find(quantityProfessions uint) ([]*pkgprofession.Profession, error) {
	var professions []*pkgprofession.Profession
	if err := r.DB.Limit(int(quantityProfessions)).Find(&professions).Error; err != nil {
		return nil, err
	}
	return professions, nil
}

func (r *repository) FindAll(limit, offset int) ([]*pkgprofession.Profession, int64, error) {

	var total int64
	// Contagem total de profissionais
	if err := r.DB.Model(&pkgprofession.Profession{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var professions []*pkgprofession.Profession
	if err := r.DB.Limit(limit).Offset(offset).Order("name").Find(&professions).Error; err != nil {
		return nil, 0, err
	}
	return professions, total, nil
}

func (r *repository) FindProfessionsWithCount() ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	// Realiza a consulta usando uma junção entre as tabelas Professional e professional_professions
	// e depois agrupa os resultados pela profissão para contar quantos profissionais estão associados a cada uma
	if err := r.DB.Table("professional_professions").
		Select("professions.name AS profession, COUNT(*) AS count").
		Joins("JOIN professions ON professional_professions.profession_id = professions.id").
		Group("professions.name").
		Order("professions.name").
		Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) FindById(id uint) (*pkgprofession.Profession, error) {
	profession := pkgprofession.Profession{}
	if err := r.DB.First(&profession, id).Error; err != nil {
		return nil, err
	}
	return &profession, nil
}

func (r *repository) Save(profession pkgprofession.Profession) (*pkgprofession.Profession, error) {
	if err := r.DB.Save(&profession).Error; err != nil {
		return nil, err
	}
	return &profession, nil
}

func (r *repository) Remove(id uint) error {

	if err := r.DB.Delete(&pkgprofession.Profession{}, id).Error; err != nil {
		return err
	}
	return nil
}
