package service

import (
	"github.com/dassajib/prohor-api/internal/model"
	"github.com/dassajib/prohor-api/internal/repository"
)

// interface defines what functionalities the note service must provide
type NoteService interface {
	Create(note *model.Note) error
	Update(note *model.Note) error
	GetNoteByID(id uint) (*model.Note, error)
	GetUserNotes(userID uint) ([]model.Note, error)
	// soft delete a note
	Delete(id uint) error
	Restore(id uint) error
	DeletePermanent(id uint) error
}

// struct that implements NoteService interface
type noteService struct {
	// uses repository layer to access DB
	repo repository.NoteRepository
}

// constructor returns a new noteService instance
func NewNoteService(repo repository.NoteRepository) NoteService {
	return &noteService{repo}
}

// calls the repository to insert a new note into the db
func (s *noteService) Create(note *model.Note) error {
	return s.repo.Create(note)
}

// calls the repository to modify an existing note
func (s *noteService) Update(note *model.Note) error {
	return s.repo.Update(note)
}

// fetches a single note by its ID (can include soft-deleted ones)
func (s *noteService) GetNoteByID(id uint) (*model.Note, error) {
	return s.repo.FindByID(id)
}

// fetches all notes belonging to a particular user
func (s *noteService) GetUserNotes(userID uint) ([]model.Note, error) {
	return s.repo.FindByUser(userID)
}

// soft delete by setting deleted_at field
func (s *noteService) Delete(id uint) error {
	return s.repo.DeleteSoft(id)
}

// restore brings back a soft-deleted note by nullifying deleted_at
func (s *noteService) Restore(id uint) error {
	return s.repo.RestoreDeleted(id)
}

// delete permanently
func (s *noteService) DeletePermanent(id uint) error {
	return s.repo.DeletePermanent(id)
}
