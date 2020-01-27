package service

import (
	"github.com/alec-rabold/EllucianBannerApi-go/model/entity"
	"github.com/alec-rabold/EllucianBannerApi-go/model/request"
)

// APIService defines the available api operations
type APIService interface {
	GetColleges() []entity.College
	GetTerms(request.TermsRequestModel) []entity.Term
	GetSubjects(request.SubjectsRequestModel) []entity.Subject
	GetCourses(request.CoursesRequestModel) []entity.Course
	GetSections(request.SectionsRequestModel) []entity.Section
}

// NewAPIService creates new service instantiation to respond to API requests
func NewAPIService() (APIService, error) {
	return NewEllucianAPIService(), nil
}
