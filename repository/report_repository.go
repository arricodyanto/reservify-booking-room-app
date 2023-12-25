package repository

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"booking-room-app/entity/dto"
	"database/sql"
	"log"
)

type ReportRepository interface {
	List(startDate, endDate string) ([]dto.ReportDto, error)
}

type reportRepository struct {
	db *sql.DB
}

// List implements ReportRepository.
func (r *reportRepository) List(startDate string, endDate string) ([]dto.ReportDto, error) {
	var reports []dto.ReportDto

	rows, err := r.db.Query(`SELECT t.id, t.employee_id, e.name, e.username, e.division, e.position, e.contact, t.room_id, r.name, r.room_type, r.capacity, t.description, t.status, t.start_time, t.end_time, t.created_at, t.updated_at FROM transactions t join employees e on e.id = t.employee_id JOIN rooms r on r.id = t.room_id order by created_at DESC`)
	if err != nil {
		log.Println("transactionsRepository.Query:", err.Error())
		return nil, err
	}
	for rows.Next() {
		var report dto.ReportDto
		err = rows.Scan(
			&report.ID,
			&report.EmployeeId,
			&report.Employee.Name,
			&report.Employee.Username,
			&report.Employee.Division,
			&report.Employee.Position,
			&report.Employee.Contact,
			&report.RoomId,
			&report.Room.Name,
			&report.Room.RoomType,
			&report.Room.Capacity,
			&report.Description,
			&report.Status,
			&report.StartTime,
			&report.EndTime,
			&report.CreatedAt,
			&report.UpdatedAt)
		if err != nil {
			log.Println("transactionsRepository.Rows.Next():", err.Error())
			return nil, err
		}

		roomFacilities, err := r.db.Query(config.SelectRoomWithFacilities, report.RoomId)
		if err != nil {
			log.Println("transactionsRepository.Query:", err.Error())
			return nil, err
		}
		for roomFacilities.Next() {
			var roomFacility entity.RoomFacility
			err = roomFacilities.Scan(
				&roomFacility.ID,
				&roomFacility.FacilityId,
				&roomFacility.Quantity,
				&roomFacility.CreatedAt,
				&roomFacility.UpdatedAt)
			if err != nil {
				log.Println("transactionsRepository.Rows.Next():", err.Error())
				return nil, err
			}
			report.RoomFacilities = append(report.RoomFacilities, roomFacility)
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepository{db: db}
}
