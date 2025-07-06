package service

import (
	"github.com/dassajib/prohor-api/internal/model"
	"github.com/dassajib/prohor-api/internal/repository"
)

// defines what functionalities the note service must provide
type NoteService interface {
	Create(note *model.Note) error
	Update(note *model.Note) error
	GetNoteByID(id uint) (*model.Note, error)
	GetUserNotes(userID uint) ([]model.Note, error)
	TogglePin(id uint, pinned bool) error
	// soft delete a note
	Delete(id uint) error
	Restore(id uint) error
	DeletePermanent(id uint) error
	SearchUserNotes(userID uint, query string) ([]model.Note, error)
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

// calls repository to create
func (s *noteService) Create(note *model.Note) error {
	return s.repo.Create(note)
}

// calls repository to update note
func (s *noteService) Update(note *model.Note) error {
	return s.repo.Update(note)
}

// call repo to find a single note(can include soft-deleted ones)
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

// brings back a soft-deleted note by nullifying deleted_at
func (s *noteService) Restore(id uint) error {
	return s.repo.RestoreDeleted(id)
}

// delete permanently
func (s *noteService) DeletePermanent(id uint) error {
	return s.repo.DeletePermanent(id)
}

// search note
func (s *noteService) SearchUserNotes(userID uint, query string) ([]model.Note, error) {
	return s.repo.SearchUserNotes(userID, query)
}

// toggle pinned status
func (s *noteService) TogglePin(id uint, pinned bool) error {
	note, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	note.Pinned = pinned
	return s.repo.Update(note)
}
