package usecase

import (
	"booking-room-app/entity/dto"
	"booking-room-app/repository"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type ReportUseCase interface {
	PrintAllReports(startDate, endDate string) ([]dto.ReportDto, error)
}

type reportUseCase struct {
	repo repository.ReportRepository
}

// FindAllReports implements ReportUseCase.
func (r *reportUseCase) PrintAllReports(startDate string, endDate string) ([]dto.ReportDto, error) {
	// Generate file
	file, err := os.Create("public/transaction.csv")
	if err != nil {
		return []dto.ReportDto{}, fmt.Errorf("failed to create reports file")
	}
	defer file.Close()

	// Write data to file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the file headers
	writer.Write([]string{"ID", "ID Pegawai", "Nama Pegawai", "Username Akun Pegawai", "Divisi", "Jabatan", "Kontak Pegawai", "ID Ruangan", "Nama Ruangan", "Jenis Ruangan", "Kapasitas", "Daftar Fasilitas", "Catatan Pemesanan", "Status Pemesanan", "Jam Mulai Peminjaman Ruangan", "Jam Akhir Peminjaman Ruangan", "Waktu Pemesanan Dibuat", "Terakhir Diperbarui"})

	reports, err := r.repo.List(startDate, endDate)
	if err != nil {
		return []dto.ReportDto{}, fmt.Errorf("oopps, failed to get transactions data")
	}

	// Write transaction data to csv file
	for _, report := range reports {
		var roomFacilityString string
		for _, v := range report.RoomFacilities {
			roomFacilityString += fmt.Sprintf("- %s, %d buah (facility_id: %s)\n", v.Name, v.Quantity, v.FacilityID)
		}

		row := []string{
			report.ID,
			report.EmployeeId,
			report.Employee.Name,
			report.Employee.Username,
			report.Employee.Division,
			report.Employee.Position,
			report.Employee.Contact,
			report.RoomId,
			report.Room.Name,
			report.Room.RoomType,
			strconv.Itoa(report.Room.Capacity),
			roomFacilityString,
			report.Description,
			report.Status,
			report.StartTime.Format("15:04"),
			report.EndTime.Format("15:04"),
			report.CreatedAt.Format("02-01-2006 15:04"),
			report.UpdatedAt.Format("02-01-2006 15:04"),
		}
		writer.Write(row)
	}

	return reports, nil
}

func NewReportUseCase(repo repository.ReportRepository) ReportUseCase {
	return &reportUseCase{repo: repo}
}
