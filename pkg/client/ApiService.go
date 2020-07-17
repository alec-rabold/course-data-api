package service

import (
	"errors"

	"github.com/alec-rabold/EllucianBannerApi-go/pkg/model/entity"
	"github.com/alec-rabold/EllucianBannerApi-go/pkg/model/request"
)

var (
	// ErrUnsupportedProvider is thrown when an unsupported provider is provided
	ErrUnsupportedProvider error = errors.New("Unsupported provider")
)

// APIClient defines the available api operations
type APIClient interface {
	GetColleges() []entity.College
	GetTerms(request.TermsRequestModel) []entity.Term
	GetSubjects(request.SubjectsRequestModel) []entity.Subject
	GetCourses(request.CoursesRequestModel) []entity.Course
	GetSections(request.SectionsRequestModel) []entity.Section
}

type provider string

// Currently supported API providers
const (
	EllucianBanner provider = "EllucianBanner"
)

// NewAPIClient creates new client instantiation to respond to API requests
func NewAPIClient(plugin provider) (APIClient, error) {
	switch plugin {
	case EllucianBanner:
		return NewEllucianAPIClient(), nil
	default:
		return nil, ErrUnsupportedProvider
	}
}
