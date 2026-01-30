package repo

import (
	"fmt"

	"gorm.io/gorm"
)

type GormDBCli struct {
	db *gorm.DB
}

type InterGormDBCli interface {
	Create(table, value interface{}) error
	Update(value Update) error
	Updates(value Updates) error
	Delete(value Delete) error
}

func NewInterGormDBCli(db *gorm.DB) InterGormDBCli {
	return &GormDBCli{
		db: db,
	}
}

// Insert record into the specified table.
func (g GormDBCli) Create(table, value interface{}) error {
	return g.executeTransaction(func(tx *gorm.DB) error {
		return tx.Model(table).Create(value).Error
	}, "data write failed")
}

// Update record into the specified table.
func (g GormDBCli) Update(value Update) error {
	return g.executeTransaction(func(tx *gorm.DB) error {
		tx = tx.Model(value.Table)
		for column, val := range value.Where {
			tx = tx.Where(column, val)
		}
		return tx.Update(value.Update[0], value.Update[1:]).Error
	}, "update data failed")
}

// Updates multi records into the specified table.
func (g GormDBCli) Updates(value Updates) error {
	return g.executeTransaction(func(tx *gorm.DB) error {
		tx = tx.Model(value.Table)
		for column, val := range value.Where {
			tx = tx.Where(column, val)
		}
		return tx.Updates(value.Updates).Error
	}, "update bulk data failed")
}

// Delete record in the specified table.
func (g GormDBCli) Delete(value Delete) error {
	return g.executeTransaction(func(tx *gorm.DB) error {
		tx = tx.Model(value.Table)
		for column, val := range value.Where {
			tx = tx.Where(column, val)
		}
		return tx.Delete(value.Table).Error
	}, "delete data failed")
}

func (g GormDBCli) executeTransaction(operation func(tx *gorm.DB) error, errorMessage string) error {
	tx := g.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("start transaction failed, err: %s", tx.Error)
	}

	if err := operation(tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("%s -> %s", errorMessage, err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("commit transaction failed, err: %s", err)
	}

	return nil
}

// Update defines the single update data structure
type Update struct {
	Table  interface{}
	Where  map[string]interface{}
	Update []string
}

// Update defines the multiple update data structure
type Updates struct {
	Table   interface{}
	Where   map[string]interface{}
	Updates interface{}
}

// Delete defines the delete data structure
type Delete struct {
	Table interface{}
	Where map[string]interface{}
}
