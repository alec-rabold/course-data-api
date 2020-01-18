package service

import (
	"github.com/alec-rabold/EllucianBannerApi-go/model/request"
	"github.com/anaskhan96/soup"
	log "github.com/sirupsen/logrus"

	"golang.org/x/net/html"
)

// EllucianDataAPIService defines the available api operations
type EllucianDataAPIService interface {
	GetTerms(request.TermsRequestModel) error
	GetSubjects(request.SubjectsRequestModel) (bool, error)
	GetCourses(request.CoursesRequestModel) (bool, error)
	GetSections(request.SectionsRequestModel) (int64, error)
}

// GetTerms returns a list of all academic terms and respective term codes
func GetTerms(request.TermsRequestModel) {

}

// GetSubjects returns a list of all departments / subjects offered for a specific term
func GetSubjects(request.SubjectsRequestModel) {

}

// GetCourses returns a list of all courses within the specified department/subject
func GetCourses(request.CoursesRequestModel) {

}

// GetSections returns a list of all sections and meeting times for a specific course
func GetSections(request.SectionsRequestModel) {

}

func getDocumentModel(collegeName, relativePath, referrerPath string, entries []string) soup.Root {
	
}
