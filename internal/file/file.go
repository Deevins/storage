package file

import (
	"fmt"
	"github.com/google/uuid"
)

type File struct {
	Name string
	ID   uuid.UUID
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	fileID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	return &File{
		Name: filename,
		ID:   fileID,
		Data: blob,
	}, nil
}

func (f *File) String() string {
	return fmt.Sprintf(`File name "%s" with id {%v}`, f.Name, f.ID)
}
