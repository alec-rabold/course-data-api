package api

import (
	"context"

	"github.com/alec-rabold/course-data-api/pkg/model/entity"
	"github.com/alec-rabold/course-data-api/pkg/model/request"
)

type provider int

const (
	ProviderEllucianBanner provider = iota
)

type ErrUnsupportedProvider struct{}

func (e ErrUnsupportedProvider) Error() string {
	return "unsupported provider"
}

// API defines the available API operations
type API interface {
	GetColleges(ctx context.Context) ([]entity.College, error)
	GetTerms(ctx context.Context, request request.Terms) ([]entity.Term, error)
	GetSubjects(ctx context.Context, request request.Subjects) ([]entity.Subject, error)
	GetCourses(ctx context.Context, request request.Courses) ([]entity.Course, error)
	GetSections(ctx context.Context, request request.Sections) ([]entity.Section, error)
}
