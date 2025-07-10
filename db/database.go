package db

import "FSA/internal/domain"

type media interface {
	SaveMediaItem(media domain.MediaItem) error
	FindMediaItemByID(ID string) (*domain.MediaItem, error)
	DeleteMedia(media domain.MediaItem) error
	ListAllMediaItems() ([]domain.MediaItem, error)
}

type config interface {
}

type Database interface {
	config
	media
}
