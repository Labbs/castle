package repository

import (
	pb "github.com/labbs/castle/gen/repository"
	"gorm.io/gorm"
)

type RepositoryRepository struct {
	database *gorm.DB
}

func NewRepostioryRepository(database *gorm.DB) RepositoryRepository {
	return RepositoryRepository{database: database}
}

func (d *RepositoryRepository) GetRepositoryByName(name string) (*pb.Repository, error) {
	u := pb.Repository{}
	r := d.database.Where("name = ?", name).First(&u)
	return &u, r.Error
}

func (d *RepositoryRepository) GetRepositoryById(id string) (*pb.Repository, error) {
	u := pb.Repository{}
	r := d.database.Where("id = ?", id).First(&u)
	return &u, r.Error
}

func (d *RepositoryRepository) UpdateRepository(repo *pb.Repository) error {
	r := d.database.Save(&repo)
	return r.Error
}

func (d *RepositoryRepository) CreateRepository(repository *pb.Repository) error {
	r := d.database.Create(&repository)
	return r.Error
}

func (d *RepositoryRepository) GetAllRepositories() ([]pb.Repository, error) {
	var repositories []pb.Repository
	r := d.database.Find(&repositories)
	return repositories, r.Error
}

func (d *RepositoryRepository) EditRepository(repository *pb.Repository) error {
	r := d.database.Save(&repository)
	return r.Error
}

func (d *RepositoryRepository) DeleteRepository(id string) error {
	r := d.database.Where("id = ?", id).Delete(&pb.Repository{})
	return r.Error
}
