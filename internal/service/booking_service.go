package services

import (
	"github.com/google/uuid"
	"github.com/sirgloveface/booking-service/internal/model"
	repositories "github.com/sirgloveface/booking-service/internal/repository"
)

type BookingService struct {
	repo *repositories.BookingRepository
}

func NewBookingService(repo *repositories.BookingRepository) *BookingService {
	return &BookingService{repo}
}

func (s *BookingService) CreateBooking(input *model.Booking) error {
	input.ID = uuid.New()
	input.Status = "pending"
	return s.repo.Create(input)
}

func (s *BookingService) GetBookingByID(id string) (*model.Booking, error) {
	return s.repo.GetByID(id)
}

func (s *BookingService) ListBookings() ([]model.Booking, error) {
	return s.repo.ListBookings()
}

func (s *BookingService) HasConflict(input *model.Booking) (bool, error) {
	return s.repo.HasConflict(input)
}

func (s *BookingService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *BookingService) Update(booking *model.Booking) error {
	return s.repo.Update(booking)
}
