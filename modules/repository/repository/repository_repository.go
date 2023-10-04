package repository

import (
	"github.com/labbs/castle/modules/repository/domain"
	"gorm.io/gorm"
)

type RepositoryRepository struct {
	database *gorm.DB
}

func NewRepostioryRepository(database *gorm.DB) RepositoryRepository {
	return RepositoryRepository{database: database}
}

func (d *RepositoryRepository) GetRepositoryByName(name string) (domain.Repository, error) {
	u := domain.Repository{}
	r := d.database.Where("name = ?", name).First(&u)
	return u, r.Error
}

func (d *RepositoryRepository) GetRepositoryById(id string) (domain.Repository, error) {
	u := domain.Repository{}
	r := d.database.Where("id = ?", id).First(&u)
	return u, r.Error
}

func (d *RepositoryRepository) UpdateRepository(repo domain.Repository) error {
	r := d.database.Save(&repo)
	return r.Error
}

func (d *RepositoryRepository) CreateRepository(repository domain.Repository) error {
	r := d.database.Create(&repository)
	return r.Error
}

func (d *RepositoryRepository) GetAllRepositories() ([]domain.Repository, error) {
	var repositories []domain.Repository
	r := d.database.Find(&repositories)
	return repositories, r.Error
}

func (d *RepositoryRepository) EditRepository(repository domain.Repository) error {
	r := d.database.Save(&repository)
	return r.Error
}
