package db

import (
	"cred-alert/sniff"
	"time"

	"github.com/jinzhu/gorm"
)

//go:generate counterfeiter . RepositoryRepository

type RepositoryRepository interface {
	FindOrCreate(*Repository) error
	Create(*Repository) error

	Find(owner string, name string) (Repository, error)

	All() ([]Repository, error)
	NotFetchedSince(time.Time) ([]Repository, error)
	NotScannedWithVersion(int) ([]Repository, error)

	MarkAsCloned(string, string, string) error
}

type repositoryRepository struct {
	db *gorm.DB
}

func NewRepositoryRepository(db *gorm.DB) *repositoryRepository {
	return &repositoryRepository{db: db}
}

func (r *repositoryRepository) Find(owner, name string) (Repository, error) {
	var repository Repository
	err := r.db.Where(Repository{Owner: owner, Name: name}).First(&repository).Error
	if err != nil {
		return Repository{}, err
	}
	return repository, nil
}

func (r *repositoryRepository) FindOrCreate(repository *Repository) error {
	r2 := Repository{Name: repository.Name, Owner: repository.Owner}
	err := r.db.Where(r2).FirstOrCreate(repository).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryRepository) Create(repository *Repository) error {
	return r.db.Create(repository).Error
}

func (r *repositoryRepository) All() ([]Repository, error) {
	var existingRepositories []Repository
	err := r.db.Find(&existingRepositories).Error
	if err != nil {
		return nil, err
	}

	return existingRepositories, nil
}

func (r *repositoryRepository) MarkAsCloned(owner, name, path string) error {
	return r.db.Model(&Repository{}).Where(
		Repository{Name: name, Owner: owner},
	).Updates(
		map[string]interface{}{"cloned": true, "path": path},
	).Error
}

func (r *repositoryRepository) NotFetchedSince(since time.Time) ([]Repository, error) {
	rows, err := r.db.Raw(`
    SELECT r.id
    FROM   fetches f
           JOIN repositories r
             ON r.id = f.repository_id
           JOIN (SELECT repository_id   AS r_id,
                        MAX(created_at) AS created_at
                 FROM   fetches
                 GROUP  BY repository_id
                ) latest_fetches
             ON f.created_at = latest_fetches.created_at
                AND f.repository_id = latest_fetches.r_id
    WHERE  r.cloned = true
      AND  latest_fetches.created_at < ?`, since).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		scanErr := rows.Scan(&id)
		if scanErr != nil {
			return nil, scanErr
		}
		ids = append(ids, id)
	}

	var repositories []Repository
	err = r.db.Model(&Repository{}).Where("id IN (?)", ids).Find(&repositories).Error
	if err != nil {
		return nil, err
	}

	return repositories, nil
}

func (r *repositoryRepository) NotScannedWithVersion(version int) ([]Repository, error) {
	rows, err := r.db.Raw(`
    SELECT r.id
    FROM   scans s
           JOIN repositories r
             ON r.id = s.repository_id
           JOIN (SELECT repository_id   AS r_id,
                        MAX(rules_version) AS rules_version
                 FROM   scans
                 GROUP  BY repository_id
                ) latest_scans
             ON s.rules_version = latest_scans.rules_version
                AND s.repository_id = latest_scans.r_id
    WHERE  r.cloned = true
      AND  latest_scans.rules_version != ?`, sniff.RulesVersion).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		scanErr := rows.Scan(&id)
		if scanErr != nil {
			return nil, scanErr
		}
		ids = append(ids, id)
	}

	var repositories []Repository
	err = r.db.Model(&Repository{}).Where("id IN (?)", ids).Find(&repositories).Error
	if err != nil {
		return nil, err
	}

	return repositories, nil
}