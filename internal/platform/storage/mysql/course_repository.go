package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/huandu/go-sqlbuilder"
	mooc "github.com/jorgeAM/goHexagonal/internal"
)

type courseRepository struct {
	db      *sql.DB
	timeout time.Duration
}

func NewCourseRepository(db *sql.DB, timeout time.Duration) mooc.CourseRepository {
	return &courseRepository{db, timeout}
}

func (r *courseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSql := new(sqlCourse)
	courseSQLStruct := sqlbuilder.NewStruct(courseSql)

	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID().String(),
		Name:     course.Name().String(),
		Duration: course.Duration().String(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)

	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}
