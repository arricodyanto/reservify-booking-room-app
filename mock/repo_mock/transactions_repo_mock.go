package repo_mock

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"

	"github.com/stretchr/testify/mock"
)

type TransactionsRepoMock struct {
	mock.Mock
}
func (t *TransactionsRepoMock) List(page, size int) ([]entity.Transaction, model.Paging, error) {
	args := t.Called(page, size)
	return args.Get(0).([]entity.Transaction), args.Get(1).(model.Paging), args.Error(2)

}

func (t *TransactionsRepoMock) Create(payload entity.Transaction) (entity.Transaction, error) {
	args := t.Called(payload)
	return args.Get(0).(entity.Transaction), args.Error(1)
}

func (m *TransactionsRepoMock) GetById(id string) (entity.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Transaction), args.Error(1)
}

func (m *TransactionsRepoMock) GetByEmployeeId(EmployeeId string) (entity.Transaction, error) {
	args := m.Called(EmployeeId)
	return args.Get(0).(entity.Transaction), args.Error(1)
}

func (m *TransactionsRepoMock) UpdatePermission(payload entity.Transaction) (entity.Transaction, error) {
	args := m.Called(payload)
	return args.Get(0).(entity.Transaction), args.Error(1)
}


