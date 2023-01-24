package storage

import (
	"fmt"
	"github.com/deevins/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)

	if err != nil {
		return nil, err
	}
	s.files[newFile.ID] = newFile

	return newFile, nil
	//fmt.Println(newFile)
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {

	found, ok := s.files[fileID]

	if !ok {
		return nil, fmt.Errorf("file with id %v not found", fileID)
	}
	return found, nil

}