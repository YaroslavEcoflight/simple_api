package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	adapterrepo "simple_api/internal/adapter/repository"
	"simple_api/internal/domain"
	"simple_api/internal/domain/entity"
	"simple_api/internal/infrastructure/model"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&model.Book{}))
	t.Cleanup(func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	})
	return db
}

// Базовые интеграционные тесты
func TestBookRepositoryCreate(t *testing.T) {
	db := setupTestDB(t)
	repo := adapterrepo.NewBookRepository(db)

	book, err := repo.Create(
		entity.Book{
			Title:  "test",
			Author: "test_author",
			Rating: 5,
		},
	)
	require.NoError(t, err)
	assert.NotZero(t, book.Id)
	assert.Equal(t, "test", book.Title)
	assert.Equal(t, "test_author", book.Author)
	assert.Equal(t, 5, book.Rating)
}

func TestBookGetById(t *testing.T) {
	db := setupTestDB(t)
	repo := adapterrepo.NewBookRepository(db)

	created, err := repo.Create(
		entity.Book{
			Title:  "test",
			Author: "test_author",
		},
	)
	book, err := repo.GetById(created.Id)
	require.NoError(t, err)
	assert.Equal(t, "test", book.Title)
}

func TestBookUpdate(t *testing.T) {
	db := setupTestDB(t)
	repo := adapterrepo.NewBookRepository(db)

	created, _ := repo.Create(
		entity.Book{
			Title:  "test",
			Author: "test_author",
			Rating: 5,
		},
	)

	book, err := repo.Update(entity.Book{
		Id:    created.Id,
		Title: "test_updated",
	})

	require.NoError(t, err)
	assert.Equal(t, "test_updated", book.Title)
	assert.Equal(t, created.Id, book.Id)
}

func TestBookDelete(t *testing.T) {
	db := setupTestDB(t)
	repo := adapterrepo.NewBookRepository(db)

	created, _ := repo.Create(
		entity.Book{
			Title:  "test",
			Author: "test_author",
			Rating: 5,
		},
	)

	err := repo.Delete(created.Id)

	require.NoError(t, err)
}

func TestBookGetById_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := adapterrepo.NewBookRepository(db)

	_, err := repo.GetById(999)

	assert.ErrorIs(t, err, domain.ErrNotFound)
}

func TestBookUpdate_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := adapterrepo.NewBookRepository(db)

	_, err := repo.Update(entity.Book{
		Id:    999,
		Title: "ghost",
	})

	assert.ErrorIs(t, err, domain.ErrNotFound)
}

func TestBookDelete_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := adapterrepo.NewBookRepository(db)

	err := repo.Delete(999)

	assert.ErrorIs(t, err, domain.ErrNotFound)
}
