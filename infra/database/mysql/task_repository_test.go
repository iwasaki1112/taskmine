package mysql

import (
	"taskmine/domain/entity"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setup() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn:                      mockDB,
			SkipInitializeWithVersion: true,
		}),
		&gorm.Config{
			NowFunc: func() time.Time {
				return time.Now().Truncate(time.Millisecond)
			},
		},
	)

	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

func TestStore(t *testing.T) {
	DB, mock, err := setup()
	if err != nil {
		t.Fatalf("faild to setup: %s", err)
	}

	r := NewTaskRepository(DB)

	task := &entity.Task{
		ID:          0,
		Title:       "テストタスク",
		Description: "テストタスクディスクリプション",
		Status:      entity.TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `tasks`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = r.Store(task)

	if err != nil {
		t.Errorf("error was not expected while storing task %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdate(t *testing.T) {
	DB, mock, err := setup()
	if err != nil {
		t.Fatalf("failed to setup: %s", err)
	}

	r := NewTaskRepository(DB)

	testTime := time.Now().Round(time.Millisecond)

	task := &entity.Task{
		ID:          1,
		Title:       "テストタスク",
		Description: "テストタスクディスクリプション",
		Status:      entity.TODO,
		CreatedAt:   testTime,
		UpdatedAt:   testTime,
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE (.+)").WithArgs(task.UpdatedAt, task.Title, task.Description, task.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err = r.Update(task)
	if err != nil {
		t.Fatal(err)
	}
}
