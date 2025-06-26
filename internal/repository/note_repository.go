package repository

import (
	"github.com/dassajib/prohor-api/internal/model"
	"gorm.io/gorm"
)

// interface defines all required methods for note operations
type NoteRepository interface {
	Create(note *model.Note) error
	Update(note *model.Note) error
	FindByID(id uint) (*model.Note, error)
	FindByUser(userID uint) ([]model.Note, error)
	// marks deleted_at but doesn't remove
	DeleteSoft(id uint) error
	RestoreDeleted(id uint) error
	DeletePermanent(id uint) error
}

type noteRepository struct {
	// gorm DB instance injected from outside
	db *gorm.DB
}

// constructor returns a new noteRepository instance
func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepository{db}
}

// inserts a new note into the notes table
func (r *noteRepository) Create(note *model.Note) error {
	return r.db.Create(note).Error
}

// update note using primary key
func (r *noteRepository) Update(note *model.Note) error {
	return r.db.Save(note).Error
}

// findByID returns a note by its ID, including soft-deleted ones
func (r *noteRepository) FindByID(id uint) (*model.Note, error) {
	var note model.Note
	// Unscoped includes soft-deleted notes
	err := r.db.Unscoped().First(&note, id).Error
	return &note, err
}

// returns all notes created by a specific user
func (r *noteRepository) FindByUser(userID uint) ([]model.Note, error) {
	var notes []model.Note
	// soft-deleted notes using Unscoped() to response with soft-deleted notes also
	err := r.db.Unscoped().Where("user_id = ?", userID).Find(&notes).Error
	return notes, err
}

// deleteSoft marks the note as deleted
// soft delete using GORM's DeletedAt
func (r *noteRepository) DeleteSoft(id uint) error {
	return r.db.Delete(&model.Note{}, id).Error
}

// restoreDeleted resets the deleted_at field to NULL (restores the note)
func (r *noteRepository) RestoreDeleted(id uint) error {
	return r.db.Unscoped().Model(&model.Note{}).Where("id = ?", id).Update("deleted_at", nil).Error
}

// permanently delete from db
func (r * noteRepository) DeletePermanent (id uint) error {
	return r.db.Unscoped().Delete(&model.Note{}, id).Error
}
