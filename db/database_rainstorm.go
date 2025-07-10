package db

import (
	"FSA/internal/domain"
	"errors"
	"fmt"

	"github.com/AndersonBargas/rainstorm/v5"
)

type DatabaseRainstorm struct {
	instance rainstorm.DB
}

func NewDatabaseRainstorm(dbInstance *rainstorm.DB) (*DatabaseRainstorm, error) {
	if dbInstance == nil {
		return nil, errors.New("The DB Instance should not be null")
	}
	db := &DatabaseRainstorm{
		instance: *dbInstance,
	}
	return db, nil
}

func (db *DatabaseRainstorm) SaveMediaItem(media domain.MediaItem) error {
	if err := db.instance.Save(&media); err != nil {
		return fmt.Errorf("Error saving the Media Item: %w", err)
	}
	return nil
}

func (db *DatabaseRainstorm) FindMediaItemByID(ID string) (*domain.MediaItem, error) {
	var mediaItem *domain.MediaItem
	if err := db.instance.One("ID", ID, &mediaItem); err != nil {
		return nil, fmt.Errorf("Error finding the Media Item by ID: %w", err)
	}
	return mediaItem, nil
}

func (db *DatabaseRainstorm) DeleteMedia(media domain.MediaItem) error {
	// if media == nil {
	// 	return errors.New("The MediaItem Struct should not be null")
	// }
	if err := db.instance.DeleteStruct(&media); err != nil {
		return fmt.Errorf("Error deleting the Struct: %w", err)
	}
	return nil
}

func (db *DatabaseRainstorm) ListAllMediaItems() ([]domain.MediaItem, error) {
	var mediaItems []domain.MediaItem
	if err := db.instance.All(&mediaItems); err != nil {
		return mediaItems, fmt.Errorf("Error listing all Media Items: %w", err)
	}
	return mediaItems, nil
}

func (db *DatabaseRainstorm) Close() error {
	return db.instance.Close()
}
