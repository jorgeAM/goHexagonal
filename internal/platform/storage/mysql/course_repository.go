package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
	mooc "github.com/jorgeAM/goHexagonal/internal"
)

type courseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) mooc.CourseRepository {
	return &courseRepository{db}
}

func (r *courseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSql := new(sqlCourse)
	courseSQLStruct := sqlbuilder.NewStruct(courseSql)

	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID().String(),
		Name:     course.Name().String(),
		Duration: course.Duration().String(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}
