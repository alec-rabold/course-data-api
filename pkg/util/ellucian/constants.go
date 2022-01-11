package ellucian

import (
	"net/url"
)

// College is a school supported by this service
type College string

// Supported Ellucian colleges and universities
const (
	AlabamaUniversityMontgomery            College = "AUM"
	AmericanUniversityInCairo              College = "AUCEgypt"
	ArkansasTechUniversity                 College = "ATU"
	BrownUniversity                        College = "Brown"
	CentralNewMexicoCommunityCollege       College = "CNM"
	ClarkAtlantaUniversity                 College = "CAU"
	ColoradoCommunityCollegeSystem         College = "CCCS"
	CreightonUniversity                    College = "Creighton"
	DeVryUniversity                        College = "DeVry"
	DrexelUniversity                       College = "Drexel"
	EastCarolinaUniversity                 College = "ECU"
	EasternKentuckyUniversity              College = "EKU"
	EasternOregonUniversity                College = "EOU"
	EmporiaStateUniversity                 College = "Emporia"
	FerrisStateUniversity                  College = "Ferris"
	FloridaInstituteOfTechnology           College = "FIT"
	GeorgeMasonUniversity                  College = "GMU"
	GeorgiaTech                            College = "GeorgiaTech"
	HawaiiPacificUniversity                College = "HPU"
	HowardUniversity                       College = "Howard"
	IllinoisInstituteOfTechnology          College = "IIT"
	ImperialValleyCollege                  College = "Imperial"
	JohnCarrollUniversity                  College = "JCU"
	KentuckyStateUniversity                College = "KYSU"
	LakeSuperiorStateUniversity            College = "LSSU"
	LewisUniverisity                       College = "Lewis"
	MontclairStateUniversity               College = "Montclair"
	MorehouseCollege                       College = "Morehouse"
	NorthCarolinaCentralUniversity         College = "NCCU"
	NorthernMichiganUniversity             College = "NMU"
	NorthwesternStateUniversity            College = "NSULA"
	NovaSoutheasternUniversity             College = "Nova"
	OaklandUniversity                      College = "Oakland"
	OklahomaBaptistUniversity              College = "OKBU"
	OklahomaStateUniversity                College = "OKState"
	PaceUniversity                         College = "Pace"
	PlymouthStateUniversity                College = "Plymouth"
	RamapoCollegeOfNewJersey               College = "Ramapo"
	RhodesCollege                          College = "Rhodes"
	RockhurstUniversity                    College = "Rockhurst"
	RowanUniversity                        College = "Rowan"
	RutgersUniversity                      College = "Rutgers"
	SouthCarolinaStateUniversity           College = "SCSU"
	SouthernIllinoisUniversityEdwardsville College = "SIUE"
	SouthernUtahUniversity                 College = "SUU"
	SpelmanCollege                         College = "Spelman"
	StocktonUniversity                     College = "Stockton"
	TennesseeStateUniversity               College = "TNState"
	TexasSouthernUniversity                College = "TSU"
	ThomasJeffersonUniversity              College = "Jefferson"
	UniversityOfAlabama                    College = "UA"
	UniversityOfAlabamaAtBirmingham        College = "UAB"
	UniversityOfAlabamainHuntsville        College = "UAH"
	UniversityOfCentralArkansas            College = "UCA"
	UniversityOfDallas                     College = "UDallas"
	UniversityOfDetroitMercy               College = "UDMercy"
	UniversityOfGeorgia                    College = "UGA"
	UniversityOfIllinoisSystem             College = "UIllinois"
	UniversityOfMemphis                    College = "Memphis"
	UniversityOfMontevallo                 College = "Montevallo"
	UniversityOfNewEnglandinMaine          College = "UNE"
	UniversityOfNewMexico                  College = "UNM"
	UniversityOfOklahoma                   College = "OU"
	UniversityOfSanFrancisco               College = "UFCA"
	UniversityOfSouthFlorida               College = "USF"
	UniversityOfTennesseeAtMartin          College = "UTM"
	UniversityOfTennesseeKnoxville         College = "UTK"
	UniversityOfTexasatElPaso              College = "UTEP"
	UniversityOfTheDistrictOfColumbia      College = "UDC"
	UniversityOfthePacific                 College = "Pacific"
	UniversitySystemOfGeorgia              College = "USG"
	VirginiaCommonwealthUniversity         College = "VCU"
	VirginiaStateUniversity                College = "VSU"
	WakeForestUniversity                   College = "WFU"
	WesternNewMexicoUniversity             College = "WNMU"
	WestGeorgiaTechnicalCollege            College = "WestGATech"
	WestVirginiaNetwork                    College = "WVNet"
	WichitaStateUniversity                 College = "Wichita"
	WilliamAndMary                         College = "WM"
	WilliamPatersonUniversity              College = "WPUNJ"
	WinthropUniversity                     College = "Winthrop"
	WrightStateUniversity                  College = "Wright"
)

// Ellucian Banner self-service portal procedures
const (
	RegistrationTermsProcedure    string = "bwckschd.p_disp_dyn_sched"
	RegistrationCoursesProcedure  string = "bwckschd.p_get_crse_unsec"
	RegistrationSubjectsProcedure string = "bwckgens.p_proc_term_date"
)

// A subset of form inputs used to request data from the self-service portal
const (
	DataFormSubject string = "sel_subj"
	DataFormCourse  string = "sel_crse"
	DataFormTerm    string = "term_in"
)

// Various data keys and values to aid in parsing
const (
	DataNameKey                 string = "name"
	DataNameValueTerms          string = "p_term"
	DataNameValueSubjects       string = "sel_subj"
	DataClassKey                string = "class"
	DataClassValueCourses       string = "ddtitle"
	DataClassTableKey           string = "SUMMARY"
	DataClassTableValueSections string = "This table lists the scheduled meeting times and assigned instructors for this class.."
	DataClassTableRowTag        string = "tr"
	DataClassTableDataColumnTag string = "td"
	DataDummyNode               string = "dummy"
)

var SelfServicePages = map[College]string{
	AlabamaUniversityMontgomery:            "https://senator.aum.edu/prod",
	AmericanUniversityInCairo:              "https://ssb-prod.ec.aucegypt.edu/PROD",
	ArkansasTechUniversity:                 "https://ssbprod.atu.edu/pls/PROD",
	BrownUniversity:                        "https://selfservice.brown.edu/ss",
	ClarkAtlantaUniversity:                 "https://ssb-prod.ec.cau.edu/PROD",
	ColoradoCommunityCollegeSystem:         "https://erpdnssb.cccs.edu/PRODFRCC",
	CreightonUniversity:                    "https://thenest.creighton.edu/prod",
	DeVryUniversity:                        "https://ssb-prod.ec.devry.edu/PROD",
	DrexelUniversity:                       "https://banner.drexel.edu/duprod/",
	EastCarolinaUniversity:                 "https://ssbprd-banner.ecu.edu/DAD_PROD",
	EasternKentuckyUniversity:              "https://web4s.eku.edu/prod",
	EasternOregonUniversity:                "https://ssb-prod.ec.eou.edu/PROD",
	EmporiaStateUniversity:                 "https://ssb.emporia.edu/pls/prod",
	FerrisStateUniversity:                  "https://banner.ferris.edu:9000/pls/GOLD",
	FloridaInstituteOfTechnology:           "https://nssb-p.adm.fit.edu/prod",
	GeorgeMasonUniversity:                  "https://patriotweb.gmu.edu/pls/prod",
	GeorgiaTech:                            "https://oscar.gatech.edu/pls/bprod",
	HawaiiPacificUniversity:                "https://bweb.hpu.edu:4443/hpud",
	HowardUniversity:                       "https://ssb-prod.ec.howard.edu/PROD",
	IllinoisInstituteOfTechnology:          "https://ssb.iit.edu/bnrprd",
	ImperialValleyCollege:                  "https://webstar.imperial.edu/pls/ivc01",
	JohnCarrollUniversity:                  "https://web4.jcu.edu/pjcu",
	KentuckyStateUniversity:                "https://ssb-prod.ec.kysu.edu/PROD",
	LakeSuperiorStateUniversity:            "https://bssmain.lssu.edu:9060/pls/PROD8",
	LewisUniverisity:                       "http://ssb.lewisu.edu:9000/PROD",
	MontclairStateUniversity:               "https://ssb.montclair.edu/PROD",
	MorehouseCollege:                       "https://ssb-prod.ec.morehouse.edu/MC",
	NorthCarolinaCentralUniversity:         "https://nccussbprod.nccu.edu/pls/NCCUPROD",
	NorthernMichiganUniversity:             "https://bssbprod.nmu.edu:8443/PROD",
	NorthwesternStateUniversity:            "https://connect.nsula.edu/prod",
	NovaSoutheasternUniversity:             "https://ssb.nova.edu/pls/PROD",
	OaklandUniversity:                      "https://sail.oakland.edu/PROD",
	OklahomaBaptistUniversity:              "https://self-service.okbu.edu/PROD",
	OklahomaStateUniversity:                "https://ssb.okstate.edu/PROD/OKC",
	PaceUniversity:                         "https://ban8prodss.pace.edu/prod",
	PlymouthStateUniversity:                "https://ssb.plymouth.edu/psc1",
	RamapoCollegeOfNewJersey:               "https://ssba.ramapo.edu:8443/myssb",
	RhodesCollege:                          "https://banweb.rhodes.edu/prod",
	RockhurstUniversity:                    "https://ssbp.rockhurst.edu:8101/PROD",
	RowanUniversity:                        "https://banner9.rowan.edu/ords/ssb",
	RutgersUniversity:                      "https://banweb.rutgers.edu/PROD",
	SouthCarolinaStateUniversity:           "https://bannerweb.scsu.edu/prod",
	SouthernIllinoisUniversityEdwardsville: "https://banner.siue.edu/ssb",
	SouthernUtahUniversity:                 "https://bannerxe.suu.edu/prod",
	SpelmanCollege:                         "https://ssb-prod.ec.spelman.edu/PROD",
	StocktonUniversity:                     "https://pssb.stockton.edu/prod",
	TennesseeStateUniversity:               "https://bannerssb.tnstate.edu/pls/PROD",
	TexasSouthernUniversity:                "https://ssb-prod.ec.tsu.edu/PROD/",
	ThomasJeffersonUniversity:              "https://banner.jefferson.edu/pls/tju",
	UniversityOfAlabama:                    "https://ssb.ua.edu/pls/PROD",
	UniversityOfAlabamaAtBirmingham:        "https://ssb.it.uab.edu/pls/sctprod",
	UniversityOfAlabamainHuntsville:        "https://ssbprod.uah.edu/PROD",
	UniversityOfCentralArkansas:            "https://ssbprod.uca.edu/PROD",
	UniversityOfDallas:                     "https://ssb-prod.ec.udallas.edu/PROD",
	UniversityOfDetroitMercy:               "https://selfserv.udmercy.edu/PROD",
	UniversityOfGeorgia:                    "https://sis-ssb-prod.uga.edu/PROD",
	UniversityOfIllinoisSystem:             "https://ui2web1.apps.uillinois.edu/BANPROD1",
	UniversityOfMemphis:                    "https://ssb.bannerprod.memphis.edu/prod",
	UniversityOfMontevallo:                 "https://linuxss.montevallo.edu:9010/PROD",
	UniversityOfNewEnglandinMaine:          "https://ssb1.une.edu:4443/pls/prod",
	UniversityOfNewMexico:                  "https://lobowebapp.unm.edu/ban_ssb/",
	UniversityOfOklahoma:                   "https://ssb.ou.edu/pls/PROD",
	UniversityOfSanFrancisco:               "https://ssb-prod.ec.usfca.edu/PROD",
	UniversityOfSouthFlorida:               "https://usfonline.admin.usf.edu/pls/prod",
	UniversityOfTennesseeAtMartin:          "https://banssb.utm.edu/prod/",
	UniversityOfTennesseeKnoxville:         "https://bannerssb.utk.edu/kbanpr",
	UniversityOfTexasatElPaso:              "https://www.goldmine.utep.edu/prod/owa",
	UniversityOfTheDistrictOfColumbia:      "https://ssb-prod.ec.udc.edu/PROD",
	UniversityOfthePacific:                 "https://bssprd.ec.pacific.edu/PROD",
	UniversitySystemOfGeorgia:              "https://gsw.gabest.usg.edu/pls/B420",
	VirginiaCommonwealthUniversity:         "https://ssb.vcu.edu/proddad",
	VirginiaStateUniversity:                "https://ssb-prod.ec.vsu.edu/BNPROD",
	WakeForestUniversity:                   "https://ssb.win.wfu.edu/BANPROD",
	WesternNewMexicoUniversity:             "https://sapphire.wnmu.edu/wnmudad",
	WestGeorgiaTechnicalCollege:            "https://bannerss.westgatech.edu/pls/ban8",
	WestVirginiaNetwork:                    "https://xesccprod.wvnet.edu/ords/",
	WichitaStateUniversity:                 "https://ssbprod.wichita.edu/PROD",
	WilliamAndMary:                         "https://prod.banner.wm.edu/ssb8",
	WilliamPatersonUniversity:              "https://selfservice.wpunj.edu/PROD",
	WinthropUniversity:                     "https://ssb.winthrop.edu/prod",
	WrightStateUniversity:                  "https://wingsexpress.wright.edu/pls/PROD",
}

// SupportedColleges is a list of all supported Ellucian colleges
var SupportedColleges = map[College]string{
	AlabamaUniversityMontgomery:            "Alabama University Montgomery",
	AmericanUniversityInCairo:              "American University in Cairo",
	BrownUniversity:                        "Brown University",
	CentralNewMexicoCommunityCollege:       "Central New Mexico Community College",
	ClarkAtlantaUniversity:                 "Clark Atlanta University",
	ColoradoCommunityCollegeSystem:         "Colorado Community College System",
	CreightonUniversity:                    "Creighton University",
	DeVryUniversity:                        "De Vry University",
	DrexelUniversity:                       "Drexel University",
	EastCarolinaUniversity:                 "East Carolina University",
	EasternKentuckyUniversity:              "Eastern Kentucky University",
	EasternOregonUniversity:                "Eastern Oregon University",
	EmporiaStateUniversity:                 "Emporia State University",
	FerrisStateUniversity:                  "Ferris State University",
	FloridaInstituteOfTechnology:           "Florida institute of Technology",
	GeorgeMasonUniversity:                  "George Mason University",
	GeorgiaTech:                            "Georgia Tech",
	HawaiiPacificUniversity:                "Hawaii Pacific University",
	HowardUniversity:                       "Howard University",
	IllinoisInstituteOfTechnology:          "Illinois institute of Technology",
	ImperialValleyCollege:                  "Imperial Valley College",
	JohnCarrollUniversity:                  "John Carroll University",
	KentuckyStateUniversity:                "Kentucky State University",
	LakeSuperiorStateUniversity:            "Lake Superior State University",
	LewisUniverisity:                       "Lewis Univerisity",
	MontclairStateUniversity:               "Montclair State University",
	MorehouseCollege:                       "Morehouse College",
	NorthCarolinaCentralUniversity:         "North Carolina Central University",
	NorthernMichiganUniversity:             "Northern Michigan University",
	NorthwesternStateUniversity:            "Northwestern State University",
	NovaSoutheasternUniversity:             "Nova Southeastern University",
	OaklandUniversity:                      "Oakland University",
	OklahomaBaptistUniversity:              "Oklahoma Baptist University",
	OklahomaStateUniversity:                "Oklahoma State University",
	PaceUniversity:                         "Pace University",
	PlymouthStateUniversity:                "Plymouth State University",
	RamapoCollegeOfNewJersey:               "Ramapo College of New Jersey",
	RhodesCollege:                          "Rhodes College",
	RockhurstUniversity:                    "Rockhurst University",
	RowanUniversity:                        "Rowan University",
	RutgersUniversity:                      "Rutgers University",
	SouthCarolinaStateUniversity:           "South Carolina State University",
	SouthernIllinoisUniversityEdwardsville: "Southern Illinois University Edwardsville",
	SouthernUtahUniversity:                 "Southern Utah University",
	SpelmanCollege:                         "Spelman College",
	StocktonUniversity:                     "Stockton University",
	TennesseeStateUniversity:               "Tennessee State University",
	TexasSouthernUniversity:                "Texas Southern University",
	ThomasJeffersonUniversity:              "Thomas Jefferson University",
	UniversityOfAlabama:                    "University of Alabama",
	UniversityOfAlabamaAtBirmingham:        "University of Alabama At Birmingham",
	UniversityOfAlabamainHuntsville:        "University of Alabamain Huntsville",
	UniversityOfCentralArkansas:            "University of Central Arkansas",
	UniversityOfDallas:                     "University of Dallas",
	UniversityOfDetroitMercy:               "University of Detroit Mercy",
	UniversityOfGeorgia:                    "University of Georgia",
	UniversityOfIllinoisSystem:             "University of Illinois System",
	UniversityOfMemphis:                    "University of Memphis",
	UniversityOfMontevallo:                 "University of Montevallo",
	UniversityOfNewEnglandinMaine:          "University of New England in Maine",
	UniversityOfNewMexico:                  "University of New Mexico",
	UniversityOfOklahoma:                   "University of Oklahoma",
	UniversityOfSanFrancisco:               "University of San Francisco",
	UniversityOfSouthFlorida:               "University of South Florida",
	UniversityOfTennesseeAtMartin:          "University of Tennessee At Martin",
	UniversityOfTennesseeKnoxville:         "University of Tennessee Knoxville",
	UniversityOfTexasatElPaso:              "University of Texasat El Paso",
	UniversityOfTheDistrictOfColumbia:      "University of The District of Columbia",
	UniversityOfthePacific:                 "University of the Pacific",
	UniversitySystemOfGeorgia:              "University System of Georgia",
	VirginiaCommonwealthUniversity:         "Virginia Commonwealth University",
	VirginiaStateUniversity:                "Virginia State University",
	WakeForestUniversity:                   "Wake Forest University",
	WesternNewMexicoUniversity:             "Western New Mexico University",
	WestGeorgiaTechnicalCollege:            "West Georgia Technical College",
	WestVirginiaNetwork:                    "West Virginia Network",
	WichitaStateUniversity:                 "Wichita State University",
	WilliamAndMary:                         "William and Mary",
	WilliamPatersonUniversity:              "William Paterson University",
	WinthropUniversity:                     "Winthrop University",
	WrightStateUniversity:                  "Wright State University",
}

// Wraps URL values for form data to make key/value manipulation easier
type formValues struct {
	values url.Values
}

func (f *formValues) Values() url.Values {
	return f.values
}

func (f *formValues) SetTerm(term string) {
	f.values.Set("term_in", term)
}

func (f *formValues) SetSubject(subject string) {
	// For certain keys, a "dummy" value is always required in addition to the actual value
	f.values.Set("sel_subj", "dummy")
	f.values.Add("sel_subj", subject)
}

func (f *formValues) SetCourse(course string) {
	f.values.Set("sel_crse", course)
}

// DefaultCourseDataFormValues returns a mutable copy of the original map
func DefaultCourseDataFormValues() *formValues {
	copy := make(url.Values, len(defaultCourseDataFormValues))
	for key, values := range defaultCourseDataFormValues {
		copy[key] = values
	}
	return &formValues{copy}
}

// Unencoded form data
var defaultCourseDataFormValues = url.Values{
	"sel_subj":  []string{"dummy"},
	"sel_day":   []string{"dummy"},
	"sel_schd":  []string{"dummy", "%"},
	"sel_insm":  []string{"dummy"},
	"sel_camp":  []string{"dummy", "%"},
	"sel_levl":  []string{"dummy", "%"},
	"sel_sess":  []string{"dummy"},
	"sel_instr": []string{"dummy", "%"},
	"sel_ptrm":  []string{"dummy", "%"},
	"sel_attr":  []string{"dummy", "%"},
	// Needed for SOME colleges (currently only "Temple University")
	// "sel_divs":      []string{"dummy", "%"},
	"sel_crse":      []string{""},
	"sel_title":     []string{""},
	"sel_from_cred": []string{""},
	"sel_to_cred":   []string{""},
	"begin_hh":      []string{"0"},
	"begin_mi":      []string{"0"},
	"begin_ap":      []string{"a"},
	"end_hh":        []string{"0"},
	"end_mi":        []string{"0"},
	"end_ap":        []string{"a"},
}

// InvalidTerms are terms that should be ingored when parsing
var InvalidTerms = []string{"(", ")", "None", "Language", "Continuing"}

// TODO: Ellucian Ethos API, e.g.:
//  * https://reg-prod.ec.lctcs.edu/StudentRegistrationSsb/ssb/term/termSelection?mepCode=RPCC&mode=courseSearch
//  * https://bannerprod.pnw.edu/StudentRegistrationSsb/ssb/term/termSelection?mode=search
//  * https://banner.stedwards.edu/StudentRegistrationSsb/ssb/term/termSelection?mode=search
//  * https://bannerprod.weber.edu/StudentRegistrationSsb/ssb/term/termSelection?mode=search
//  * https://registration.texastech.edu/StudentRegistrationSsb/ssb/registration?mepCode=TTU

// TODO: Non-standard data layout
//   Course:
//     * BatesCollege:                     "https://gg-bprod.bates.edu/bprod",
//     * BryantUniversity:                 "https://ssb-prod.bryantec.bryant.edu/PROD",
//     * BridgewaterStateUniversity:       "https://infobear.bridgew.edu/BANP",
//     * ConnecticutCollege:               "https://ssbprod.conncoll.edu/CONN/",
//     * KennesawStateUniversityinGeorgia: "https://owlexpress.kennesaw.edu/prodban",
//   Terms:
//     * CentralNewMexicoCommunityCollege: "https://banner.cnm.edu/proddad",
//   Subjects:
//     * NewCollegeOfFlorida:            "https://newcleis.ncf.edu/pls/ncpo",
//   Etc:
//     * FashionInstituteOfTechnology: "https://banweb02.fitnyc.edu/prodssb", // "Database Log In Failed -- TNS is unable to connect to destination."
//     * GeorgiaStateUniversity:                  "https://www.gosolar.gsu.edu/bprod" // "degree" as a form input
//     * SelfServicePages contains a mapping of universities to self-service registration pages
//     *  HarperCollege:                           "https://student-self-service.harpercollege.edu/prod", unavailable page, maybe under temp maintenance
