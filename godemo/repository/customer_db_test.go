package repository

import (
	"errors"
	"testing"

	gmock "godemo/repository/mock_repository"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type mockDB struct {
	q string
}

func (md *mockDB) Select(dest interface{}, query string, args ...interface{}) error {
	md.q = query
	return nil
}

func (md *mockDB) Get(dest interface{}, query string, args ...interface{}) error {
	md.q = query
	return nil
}

func TestGetCustomer(t *testing.T) {
	mock := &mockDB{}
	query := NewCustomerRepository(mock)
	results, err := query.GetAll()
	assert.NotNil(t, results)
	assert.NoError(t, err)

	query2 := `select customer_id, name, date_of_birth, city, zipcode, status 
		from customers`

	if mock.q != query2 {
		t.Errorf(` %v\n`, mock.q)
	}
}

func TestGetCustomer_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSv := gmock.NewMockDB(ctrl)
	mockSv.EXPECT().Select(gomock.Any(), gomock.Any()).Return(errors.New("database error"))

	// mock := &mockDB{}
	query := NewCustomerRepository(mockSv)

	results, err := query.GetAll()
	assert.Nil(t, results)
	assert.Error(t, err)
}

func TestGetCustomerById(t *testing.T) {
	mock := &mockDB{}
	query := NewCustomerRepository(mock)
	results, err := query.GetById(1)
	assert.NotNil(t, results)
	assert.NoError(t, err)
}

func TestGetCustomerById_Fail(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSv := gmock.NewMockDB(ctrl)
	mockSv.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("database error"))

	query := NewCustomerRepository(mockSv)
	results, err := query.GetById(1)
	assert.Nil(t, results)
	assert.Error(t, err)
}