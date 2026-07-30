package city_repository_impl

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	pkgcity "construir_mais_barato/app/domain/city"
)

type repository struct {
	DB *gorm.DB
}

func NewCityRepositoryImpl(db *gorm.DB) pkgcity.CityRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindAll() ([]*pkgcity.City, error) {
	var citys []*pkgcity.City
	if err := r.DB.Find(&citys).Error; err != nil {
		return nil, err
	}
	return citys, nil
}

func (r *repository) FindById(id uint) (*pkgcity.City, error) {
	city := pkgcity.City{}
	if err := r.DB.First(&city, id).Error; err != nil {
		return nil, err
	}
	return &city, nil
}

func (r *repository) FindByUF(uf string) ([]*pkgcity.City, error) {
	cities := make([]*pkgcity.City, 0)
	if err := r.DB.Where("uf = ?", uf).Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}

func (r *repository) SearchByName(term string, limit int) ([]*pkgcity.City, error) {
	cities := make([]*pkgcity.City, 0)

	sanitizedTerm := stripLikeWildcards(term)
	// Um termo formado somente por curingas ficaria vazio e casaria com a base
	// inteira, portanto nao chega a consultar o banco.
	if sanitizedTerm == "" {
		return cities, nil
	}

	contains := "%" + sanitizedTerm + "%"
	startsWith := sanitizedTerm + "%"

	if err := r.DB.
		Where("name LIKE ?", contains).
		Clauses(clause.OrderBy{
			Expression: clause.Expr{
				SQL:                "CASE WHEN name LIKE ? THEN 0 ELSE 1 END, name ASC",
				Vars:               []interface{}{startsWith},
				WithoutParentheses: true,
			},
		}).
		Limit(limit).
		Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}

// stripLikeWildcards remove os curingas do LIKE do termo digitado pelo usuario.
func stripLikeWildcards(term string) string {
	replacer := strings.NewReplacer("%", "", "_", "", "\\", "")
	return replacer.Replace(term)
}

func (r *repository) Save(city pkgcity.City) (*pkgcity.City, error) {
	if err := r.DB.Save(&city).Error; err != nil {
		return nil, err
	}
	return &city, nil
}

func (r *repository) Remove(id uint) error {

	if err := r.DB.Delete(&pkgcity.City{}, id).Error; err != nil {
		return err
	}
	return nil
}
