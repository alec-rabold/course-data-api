package service

import (
	// "github.com/PuerkitoBio/goquery"
	"github.com/alec-rabold/EllucianBannerApi-go/model/entity"
	"github.com/alec-rabold/EllucianBannerApi-go/model/request"
	// "github.com/anaskhan96/soup"
	// log "github.com/sirupsen/logrus"
	// "golang.org/x/net/html"
)

// APIService defines the available api operations
type APIService interface {
	GetTerms(request.TermsRequestModel)
	GetSubjects(request.SubjectsRequestModel)
	GetCourses(request.CoursesRequestModel)
	GetSections(request.SectionsRequestModel) []entity.Section
}

// NewAPIService creates new service instantiation to respond to API requests
func NewAPIService() (APIService, error) {
	return NewEllucianAPIService(), nil
}
