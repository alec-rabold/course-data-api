package test

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/alec-rabold/course-data-api/pkg/client/ellucian"
	"github.com/alec-rabold/course-data-api/pkg/crawler"
)

var ctx context.Context

// Shared pool of supported clients
var (
	ellucianClient *ellucian.Client
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2E Suite Test")
}

var _ = BeforeSuite(func() {
	ctx = context.Background()
	ellucianClient = ellucian.NewClient(crawler.NewCrawler(
		crawler.WithContentType("application/x-www-form-urlencoded"),
	))
})
