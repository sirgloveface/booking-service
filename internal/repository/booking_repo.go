package repositories

import (
	"github.com/sirgloveface/booking-service/internal/model"
	"gorm.io/gorm"
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{db}
}

func (r *BookingRepository) Create(booking *model.Booking) error {
	return r.db.Create(booking).Error
}

func (r *BookingRepository) GetByID(id string) (*model.Booking, error) {
	var booking model.Booking
	if err := r.db.First(&booking, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *BookingRepository) ListBookings() ([]model.Booking, error) {
	var bookings []model.Booking
	if err := r.db.Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *BookingRepository) HasConflict(booking *model.Booking) (bool, error) {
	var existing model.Booking
	if err := r.db.
		Where("boat_id = ? AND start_time < ? AND end_time > ?", booking.BoatID, booking.EndTime, booking.StartTime).
		First(&existing).Error; err == nil {
		return true, err
	}
	return false, nil

}

func (r *BookingRepository) Delete(id string) error {
	result := r.db.Delete(&model.Booking{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *BookingRepository) Update(booking *model.Booking) error {
	return r.db.Save(booking).Error
}
