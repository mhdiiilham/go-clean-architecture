package book

import "gorm.io/gorm"

type mySQLRepo struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) *mySQLRepo {
	return &mySQLRepo{
		db: db,
	}
}

func (r *mySQLRepo) Create(e *Book) (int, error) {
	tx := r.db.Create(&e)
	if tx.Error != nil {
		return e.ID, tx.Error
	}
	return e.ID, nil
}

func (r *mySQLRepo) Remove(e *Book) error {
	tx := r.db.Delete(&e)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *mySQLRepo) Get(ID int) (*Book, error) {
	var b Book
	tx := r.db.First(&b, ID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &b, nil
}
