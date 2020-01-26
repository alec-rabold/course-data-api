package util

// Global constants for EllucianDataLambda
const (
	/* API Paths */
	APIPathColleges string = "/v1/colleges"
	APIPathTerms    string = "/v1/terms"
	APIPathSubjects string = "/v1/subjects"
	APIPathCourses  string = "/v1/courses"
	APIPathSections string = "/v1/sections"
	/* Supported colleges */
	UniversityOfWashington             string = "UW"
	WashingtonStateUniversity          string = "WSU"
	SanDiegoStateUniversity            string = "SDSU"
	UniversityOfCaliforniaBerkeley     string = "Berkeley"
	UniversityOfCaliforniaSantaBarbara string = "UCSB"
	GeorgiaStateUniversity             string = "GSU"
	WeberStateUniversity               string = "Weber"
	DrexelUniversity                   string = "Drexel"
	PurdueUniversity                   string = "Purdue"
	PurdueUniversityNorthwest          string = "PurdueNW"
	GeorgeMasonUniversity              string = "GMU"
	UniversityOfTennesseeKnoxville     string = "UTK"
	HarperCollege                      string = "Harper"
	BrownUniversity                    string = "Brown"
	GeorgiaTech                        string = "GeorgiaTech"
	/** Ellucian */
	EllucianCourseDataFormDataDefault        string = "sel_subj=dummy&sel_day=dummy&sel_schd=dummy&sel_insm=dummy&sel_camp=dummy&sel_levl=dummy&sel_sess=dummy&sel_instr=dummy&sel_ptrm=dummy&sel_attr=dummy&sel_title=&sel_schd=%25&sel_from_cred=&sel_to_cred=&sel_levl=%25&sel_instr=%25&sel_attr=%25&begin_hh=0&begin_mi=0&begin_ap=a&end_hh=0&end_mi=0&end_ap=a&term_in="
	EllucianRegistrationTermsRelativePath    string = "/bwckschd.p_disp_dyn_sched"
	EllucianRegistrationCoursesRelativePath  string = "/bwckschd.p_get_crse_unsec"
	EllucianRegistrationSubjectsRelativePath string = "/bwckgens.p_proc_term_date"
	EllucianDataNameKey                      string = "name"
	EllucianDataNameValueTerms               string = "p_term"
	EllucianDataNameValueSubjects            string = "sel_subj"
	EllucianDataClassKey                     string = "class"
	EllucianDataClassValueCourses            string = "ddtitle"
	EllucianDataClassTableKey                string = "SUMMARY"
	EllucianDataClassTableValueSections      string = "This table lists the scheduled meeting times and assigned instructors for this class.."
	EllucianDataClassTableRowTag             string = "tr"
	EllucianDataClassTableDataColumnTag      string = "td"
	EllucianDataFormSubject                  string = "&sel_subj="
	EllucianDataFormCourse                   string = "&sel_crse="
	EllucianDataDummyNode                    string = "dummy"
	/* Self-Service Data Pages */
	EllucianDataGeorgiaStateUniversity          string = "https://www.gosolar.gsu.edu/bprod"
	EllucianDataWeberStateUniversity            string = "https://selfservice.weber.edu/pls/proddad"
	EllucianDataDrexelUniversity                string = "https://banner.drexel.edu/pls/duprod"
	EllucianDataPurdueUniversity                string = "https://selfservice.mypurdue.purdue.edu/prod"
	EllucianDataPurdueUniversityNorthwest       string = "https://ssb-prod.pnw.edu/dbServer_prod"
	EllucianDataGeorgeMasonUniversity           string = "https://patriotweb.gmu.edu/pls/prod"
	EllucianDataUniversityOfTennesseeKnowxville string = "https://bannerssb.utk.edu/kbanpr"
	EllucianDataHarperCollege                   string = "https://student-self-service.harpercollege.edu/prod"
	EllucianDataBrownUniversity                 string = "https://selfservice.brown.edu/ss"
	EllucianDataGeorgiaTech                     string = "https://oscar.gatech.edu/pls/bprod"
)

// EllucianSupportedColleges is a list of all supported Ellucian colleges
var EllucianSupportedColleges = map[string]string{
	BrownUniversity:                "Brown University",
	DrexelUniversity:               "Drexel University",
	GeorgeMasonUniversity:          "George Mason University",
	GeorgiaStateUniversity:         "Georgia State University",
	GeorgiaTech:                    "Georgia Tech",
	HarperCollege:                  "Harper College",
	PurdueUniversity:               "Purdue University",
	PurdueUniversityNorthwest:      "Purdue University Northwest",
	UniversityOfTennesseeKnoxville: "University of Tennessee Knoxville",
	WeberStateUniversity:           "Weber State University",
}

// EllucianUniversitiesDataPages maps university names to their respective Ellucian data pages
var EllucianUniversitiesDataPages = map[string]string{
	GeorgiaStateUniversity:         EllucianDataGeorgiaStateUniversity,
	WeberStateUniversity:           EllucianDataWeberStateUniversity,
	DrexelUniversity:               EllucianDataDrexelUniversity,
	PurdueUniversity:               EllucianDataPurdueUniversity,
	PurdueUniversityNorthwest:      EllucianDataPurdueUniversityNorthwest,
	GeorgeMasonUniversity:          EllucianDataGeorgeMasonUniversity,
	UniversityOfTennesseeKnoxville: EllucianDataUniversityOfTennesseeKnowxville,
	HarperCollege:                  EllucianDataHarperCollege,
	BrownUniversity:                EllucianDataBrownUniversity,
	GeorgiaTech:                    EllucianDataGeorgiaTech,
}

// InvalidTerms are terms that should be ingored when parsing
var InvalidTerms = []string{"(", ")", "None", "Language", "Continuing"}
