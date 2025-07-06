package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dassajib/prohor-api/internal/model"
	"github.com/dassajib/prohor-api/internal/service"
	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	// uses the note service layer
	service service.NoteService
}

// constructor for NoteHandler
func NewNoteHandler(service service.NoteService) *NoteHandler {
	return &NoteHandler{service}
}

// saves a new note with title, content, tag, and auto-filled user and date
func (h *NoteHandler) CreateNote(c *gin.Context) {
	// get logged-in user's ID from token (set by auth middleware)
	userID := c.MustGet("user_id").(uint)
	var note model.Note

	// bind JSON input to note model
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// set required fields manually
	note.UserID = userID
	note.Date = time.Now()

	if err := h.service.Create(&note); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create note"})
		return
	}

	c.JSON(http.StatusCreated, note)
}

// allows partial update (title, content, tag), also updates date automatically
func (h *NoteHandler) UpdateNote(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	// read note ID from URL param
	var noteID uint
	if _, err := fmt.Sscan(c.Param("id"), &noteID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid note ID"})
		return
	}

	// fetch note from DB
	existingNote, err := h.service.GetNoteByID(noteID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
		return
	}

	// only allow editing user's own note
	if existingNote.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized access"})
		return
	}

	// allow partial update (title, content, tag)
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if title, ok := updateData["title"].(string); ok {
		existingNote.Title = title
	}
	if content, ok := updateData["content"].(string); ok {
		existingNote.Content = content
	}
	if tag, ok := updateData["tag"].(string); ok {
		existingNote.Tag = tag
	}

	// update the date field on every update
	existingNote.Date = time.Now()

	if err := h.service.Update(existingNote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update note"})
		return
	}

	c.JSON(http.StatusOK, existingNote)
}

// returns all notes created by the logged-in user
func (h *NoteHandler) GetUserNotes(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	notes, err := h.service.GetUserNotes(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch notes"})
		return
	}

	c.JSON(http.StatusOK, notes)
}

// performs a soft delete (sets deleted_at) for safety
func (h *NoteHandler) DeleteNote(c *gin.Context) {
	var id uint
	fmt.Sscan(c.Param("id"), &id)

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "note soft deleted"})
}

// restores a soft-deleted note (if deleted less than 30 days ago)
func (h *NoteHandler) RestoreNote(c *gin.Context) {
	var id uint
	fmt.Sscan(c.Param("id"), &id)

	if err := h.service.Restore(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "note restored"})
}

// del permanently
func (h *NoteHandler) DeleteNotePermanent(c *gin.Context) {
	var id uint
	if _, err := fmt.Sscan(c.Param("id"), &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid note ID"})
		return
	}

	if err := h.service.DeletePermanent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not permanently delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "note permanently deleted"})
}

// search note
func (h *NoteHandler) SearchNotes(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	query := c.Query("q")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "search query cannot be empty"})
		return
	}

	notes, err := h.service.SearchUserNotes(userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "search failed"})
		return
	}

	c.JSON(http.StatusOK, notes)
}

// to handle toggle pinned
func (h *NoteHandler) TogglePin(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	// read note id from URL parameter and convert to uint
	var noteID uint
	if _, err := fmt.Sscan(c.Param("id"), &noteID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid note ID"})
		return
	}

	// fetch the note from the service
	note, err := h.service.GetNoteByID(noteID)
	if err != nil || note.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized access"})
		return
	}

	// bind request body JSON to a struct to extract the "pinned" field
	var payload struct {
		Pinned bool `json:"pinned"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid input",
			"details": err.Error(),
		})
		return
	}

	// call service to toggle the pinned status
	if err := h.service.TogglePin(noteID, payload.Pinned); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update pin status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "pin status updated"})
}
