package mysql

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mooc "github.com/jorgeAM/goHexagonal/internal/platform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCourseRepositoryWithSuccess(t *testing.T) {
	courseID, courseName, courseDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Course", "10 months"
	course, err := mooc.NewCourse(courseID, courseName, courseDuration)

	require.NoError(t, err)

	db, mock, err := sqlmock.New()
	defer db.Close()

	require.NoError(t, err)

	mock.ExpectExec("INSERT INTO courses").WithArgs(
		courseID,
		courseName,
		courseDuration,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	repository := NewCourseRepository(db)
	err = repository.Save(context.Background(), *course)

	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Nil(t, err)
	assert.NoError(t, err)
}

func TestCourseRepositoryWithError(t *testing.T) {
	courseID, courseName, courseDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Course", "10 months"
	course, err := mooc.NewCourse(courseID, courseName, courseDuration)

	require.NoError(t, err)

	db, mock, err := sqlmock.New()
	defer db.Close()

	require.NoError(t, err)

	mock.ExpectExec("INSERT INTO courses").WithArgs(
		courseID,
		courseName,
		courseDuration,
	).WillReturnError(errors.New("something got wrong"))

	repository := NewCourseRepository(db)
	err = repository.Save(context.Background(), *course)

	assert.NoError(t, mock.ExpectationsWereMet())
	assert.NotNil(t, err)
	assert.Error(t, err)
}
