package repository

import (
	"github.com/labbs/castle/modules/repository/domain"
	"gorm.io/gorm"
)

type repositoryRepository struct {
	database *gorm.DB
}

func NewRepostioryRepository(database *gorm.DB) *repositoryRepository {
	return &repositoryRepository{database: database}
}

func (d *repositoryRepository) GetRepositoryByName(name string) (domain.Repository, error) {
	u := domain.Repository{}
	r := d.database.Where("name = ?", name).First(&u)
	return u, r.Error
}

func (d *repositoryRepository) GetRepositoryById(id string) (domain.Repository, error) {
	u := domain.Repository{}
	r := d.database.Where("id = ?", id).First(&u)
	return u, r.Error
}

func (d *repositoryRepository) UpdateRepository(repo domain.Repository) error {
	r := d.database.Save(&repo)
	return r.Error
}

func (d *repositoryRepository) CreateRepository(repository domain.Repository) error {
	r := d.database.Create(&repository)
	return r.Error
}

func (d *repositoryRepository) GetAllRepositories() ([]domain.Repository, error) {
	var repositories []domain.Repository
	r := d.database.Find(&repositories)
	return repositories, r.Error
}

func (d *repositoryRepository) EditRepository(repository domain.Repository) error {
	r := d.database.Save(&repository)
	return r.Error
}

func (d *repositoryRepository) DeleteRepository(id string) error {
	r := d.database.Where("id = ?", id).Delete(&domain.Repository{})
	return r.Error
}
