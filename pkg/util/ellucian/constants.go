package ellucian

import (
	"net/url"
)

// College is a school supported by this service
type College string

// Supported Ellucian colleges and universities
const (
	AlabamaUniversityMontgomery             College = "AUM"
	AmericanUniversityInCairo               College = "AUCEgypt"
	AppalachianStateUniversity              College = "AppState"
	ArkansasTechUniversity                  College = "ATU"
	BatesCollege                            College = "Bates"
	BridgewaterStateUniversity              College = "Bridgew"
	BrownUniversity                         College = "Brown"
	BryantUniversity                        College = "Bryant"
	CanisiusCollege                         College = "Canisius"
	CentralNewMexicoCommunityCollege        College = "CNM"
	ClarkAtlantaUniversity                  College = "CAU"
	ColoradoCommunityCollegeSystem          College = "CCCS"
	ColoradoSchoolOfMines                   College = "Mines"
	ConnecticutCollege                      College = "ConnColl"
	CreightonUniversity                     College = "Creighton"
	DeVryUniversity                         College = "DeVry"
	DrexelUniversity                        College = "Drexel"
	EastCarolinaUniversity                  College = "ECU"
	EasternKentuckyUniversity               College = "EKU"
	EasternOregonUniversity                 College = "EOU"
	EmporiaStateUniversity                  College = "Emporia"
	FashionInstituteOfTechnology            College = "FITNYC"
	FerrisStateUniversity                   College = "Ferris"
	FloridaInstituteOfTechnology            College = "FIT"
	GeorgeMasonUniversity                   College = "GMU"
	GeorgiaStateUniversity                  College = "GSU"
	GeorgiaTech                             College = "GeorgiaTech"
	HarperCollege                           College = "Harper"
	HawaiiPacificUniversity                 College = "HPU"
	HofstraUniversity                       College = "Hofstra"
	HowardUniversity                        College = "Howard"
	IllinoisInstituteOfTechnology           College = "IIT"
	ImperialValleyCollege                   College = "Imperial"
	JohnCarrollUniversity                   College = "JCU"
	KennesawStateUniversityinGeorgia        College = "Kennesaw"
	KentuckyStateUniversity                 College = "KYSU"
	LakeSuperiorStateUniversity             College = "LSSU"
	LewisUniverisity                        College = "Lewis"
	LongwoodUniversity                      College = "Longwood"
	LouisianasCommunityAndTechnicalColleges College = "LCTCS"
	MontclairStateUniversity                College = "Montclair"
	MorehouseCollege                        College = "Morehouse"
	MorganStateUniversity                   College = "Morgan"
	NewCollegeOfFlorida                     College = "NCF"
	NorthCarolinaCentralUniversity          College = "NCCU"
	NorthernMichiganUniversity              College = "NMU"
	NorthwesternStateUniversity             College = "NSULA"
	NovaSoutheasternUniversity              College = "Nova"
	OaklandUniversity                       College = "Oakland"
	OklahomaBaptistUniversity               College = "OKBU"
	OklahomaStateUniversity                 College = "OKState"
	OralRobertsUniversity                   College = "ORU"
	PaceUniversity                          College = "Pace"
	PlymouthStateUniversity                 College = "Plymouth"
	PurdueUniversity                        College = "Purdue"
	PurdueUniversityNorthwest               College = "PurdueNW"
	RamapoCollegeOfNewJersey                College = "Ramapo"
	RhodesCollege                           College = "Rhodes"
	RockhurstUniversity                     College = "Rockhurst"
	RowanUniversity                         College = "Rowan"
	RutgersUniversity                       College = "Rutgers"
	SanDiegoStateUniversity                 College = "SDSU"
	SouthCarolinaStateUniversity            College = "SCSU"
	SouthernIllinoisUniversityEdwardsville  College = "SIUE"
	SouthernUtahUniversity                  College = "SUU"
	SpelmanCollege                          College = "Spelman"
	StEdwardsUniversityInAustin             College = "StEdwards"
	StocktonUniversity                      College = "Stockton"
	TempleUniversity                        College = "Temple"
	TennesseeStateUniversity                College = "TNState"
	TexasSouthernUniversity                 College = "TSU"
	TexasTechUniversitySystem               College = "TexasTech"
	ThomasJeffersonUniversity               College = "Jefferson"
	UniversityOfAlabama                     College = "UA"
	UniversityOfAlabamaAtBirmingham         College = "UAB"
	UniversityOfAlabamainHuntsville         College = "UAH"
	UniversityOfCaliforniaBerkeley          College = "Berkeley"
	UniversityOfCaliforniaSantaBarbara      College = "UCSB"
	UniversityOfCentralArkansas             College = "UCA"
	UniversityOfDallas                      College = "UDallas"
	UniversityOfDetroitMercy                College = "UDMercy"
	UniversityOfGeorgia                     College = "UGA"
	UniversityOfIllinoisSystem              College = "UIllinois"
	UniversityOfMemphis                     College = "Memphis"
	UniversityOfMontevallo                  College = "Montevallo"
	UniversityOfNewEnglandinMaine           College = "UNE"
	UniversityOfNewMexico                   College = "UNM"
	UniversityOfOklahoma                    College = "OU"
	UniversityOfSanFrancisco                College = "UFCA"
	UniversityOfSouthFlorida                College = "USF"
	UniversityOfTennesseeAtMartin           College = "UTM"
	UniversityOfTennesseeKnoxville          College = "UTK"
	UniversityOfTexasatElPaso               College = "UTEP"
	UniversityOfTheDistrictOfColumbia       College = "UDC"
	UniversityOfthePacific                  College = "Pacific"
	UniversityOfWashington                  College = "UW"
	UniversitySystemOfGeorgia               College = "USG"
	VirginiaCommonwealthUniversity          College = "VCU"
	VirginiaStateUniversity                 College = "VSU"
	WakeForestUniversity                    College = "WFU"
	WashingtonStateUniversity               College = "WSU"
	WeberStateUniversity                    College = "Weber"
	WesternNewMexicoUniversity              College = "WNMU"
	WestGeorgiaTechnicalCollege             College = "WestGATech"
	WestVirginiaNetwork                     College = "WVNet"
	WichitaStateUniversity                  College = "Wichita"
	WilliamAndMary                          College = "WM"
	WilliamPatersonUniversity               College = "WPUNJ"
	WinthropUniversity                      College = "Winthrop"
	WrightStateUniversity                   College = "Wright"
)

// Ellucian Banner self-service portal endpoints
const (
	RegistrationTermsRelativePath    string = "/bwckschd.p_disp_dyn_sched"
	RegistrationCoursesRelativePath  string = "/bwckschd.p_get_crse_unsec"
	RegistrationSubjectsRelativePath string = "/bwckgens.p_proc_term_date"
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

// SelfServicePages contains a mapping of universities to self-service registration pages
var SelfServicePages = map[College]string{
	AlabamaUniversityMontgomery:             "https://senator.aum.edu/prod",
	AmericanUniversityInCairo:               "https://ssb-prod.ec.aucegypt.edu:9100/PROD",
	AppalachianStateUniversity:              "https://bannerweb.appstate.edu/pls/PROD",
	ArkansasTechUniversity:                  "https://ssbprod.atu.edu/pls/PROD",
	BatesCollege:                            "https://gg-bprod.bates.edu/bprod",
	BridgewaterStateUniversity:              "https://infobear.bridgew.edu/BANP",
	BrownUniversity:                         "https://selfservice.brown.edu/ss",
	BryantUniversity:                        "https://ssb-prod.bryantec.bryant.edu/PROD",
	CanisiusCollege:                         "https://asterope.canisius.edu/pls/prod",
	CentralNewMexicoCommunityCollege:        "https://banner.cnm.edu/proddad",
	ClarkAtlantaUniversity:                  "https://ssb-prod.ec.cau.edu/PROD",
	ColoradoCommunityCollegeSystem:          "https://erpdnssb.cccs.edu/PRODFRCC",
	ColoradoSchoolOfMines:                   "https://banner.mines.edu/prod/owa",
	ConnecticutCollege:                      "https://ssbprod.conncoll.edu:9112/CONN",
	CreightonUniversity:                     "https://thenest.creighton.edu/PROD",
	DeVryUniversity:                         "https://ssb-prod.ec.devry.edu/PROD",
	DrexelUniversity:                        "https://banner.drexel.edu/pls/duprod",
	EastCarolinaUniversity:                  "https://ssbprd-banner.ecu.edu/DAD_PROD",
	EasternKentuckyUniversity:               "https://web4s.eku.edu/prod",
	EasternOregonUniversity:                 "https://ssb-prod.ec.eou.edu/PROD",
	EmporiaStateUniversity:                  "https://ssb.emporia.edu/pls/prod",
	FashionInstituteOfTechnology:            "https://banweb02.fitnyc.edu/prodssb",
	FerrisStateUniversity:                   "https://banner.ferris.edu:9000/pls/GOLD",
	FloridaInstituteOfTechnology:            "https://nssb-p.adm.fit.edu/prod",
	GeorgeMasonUniversity:                   "https://patriotweb.gmu.edu/pls/prod",
	GeorgiaStateUniversity:                  "https://www.gosolar.gsu.edu/bprod",
	GeorgiaTech:                             "https://oscar.gatech.edu/pls/bprod",
	HarperCollege:                           "https://student-self-service.harpercollege.edu/prod",
	HawaiiPacificUniversity:                 "https://bweb.hpu.edu:4443/hpud",
	HofstraUniversity:                       "https://hofstraonline.hofstra.edu/pls/HPRO",
	HowardUniversity:                        "https://ssb-prod.ec.howard.edu/PROD",
	IllinoisInstituteOfTechnology:           "https://ssb.iit.edu/bnrprd",
	ImperialValleyCollege:                   "https://webstar.imperial.edu/pls/ivc01",
	JohnCarrollUniversity:                   "https://web4.jcu.edu/pjcu",
	KennesawStateUniversityinGeorgia:        "https://owlexpress.kennesaw.edu/prodban",
	KentuckyStateUniversity:                 "https://ssb-prod.ec.kysu.edu/PROD",
	LakeSuperiorStateUniversity:             "https://bssmain.lssu.edu:9060/pls/PROD8",
	LewisUniverisity:                        "http://ssb.lewisu.edu:9000/PROD",
	LongwoodUniversity:                      "https://bannerssb.longwood.edu/bnr8prod",
	LouisianasCommunityAndTechnicalColleges: "https://ssb-prod.ec.lctcs.edu/PROD_NSHORE",
	MontclairStateUniversity:                "https://ssb.montclair.edu/PROD",
	MorehouseCollege:                        "https://ssb-prod.ec.morehouse.edu/MC",
	MorganStateUniversity:                   "https://lbssbnprod.morgan.edu/nprod",
	NewCollegeOfFlorida:                     "https://banner.ncf.edu/pls/ncpo",
	NorthCarolinaCentralUniversity:          "https://nccussbprod.nccu.edu/pls/NCCUPROD",
	NorthernMichiganUniversity:              "https://bssbprod.nmu.edu:8443/PROD",
	NorthwesternStateUniversity:             "https://connect.nsula.edu/prod",
	NovaSoutheasternUniversity:              "https://ssb.nova.edu/pls/PROD",
	OaklandUniversity:                       "https://sail.oakland.edu/PROD",
	OklahomaBaptistUniversity:               "https://self-service.okbu.edu/PROD",
	OklahomaStateUniversity:                 "https://ssb.okstate.edu/OKM",
	OralRobertsUniversity:                   "https://vision.oru.edu/prod",
	PaceUniversity:                          "https://ban8prodss.pace.edu/prod",
	PlymouthStateUniversity:                 "https://ssb.plymouth.edu/psc1",
	PurdueUniversity:                        "https://selfservice.mypurdue.purdue.edu/prod",
	PurdueUniversityNorthwest:               "https://ssb-prod.pnw.edu/dbServer_prod",
	RamapoCollegeOfNewJersey:                "https://ssb.ramapo.edu/pls/RCNJ",
	RhodesCollege:                           "https://banweb.rhodes.edu/prod",
	RockhurstUniversity:                     "https://rockweb.rockhurst.edu/pls/dbp",
	RowanUniversity:                         "https://banner9.rowan.edu/ords/ssb",
	RutgersUniversity:                       "https://banweb.rutgers.edu/PROD",
	SouthCarolinaStateUniversity:            "https://bannerweb.scsu.edu/prod",
	SouthernIllinoisUniversityEdwardsville:  "https://banner.siue.edu/ssb",
	SouthernUtahUniversity:                  "https://bannerxe.suu.edu/prod",
	SpelmanCollege:                          "https://ssb-prod.ec.spelman.edu/PROD",
	StEdwardsUniversityInAustin:             "https://ssb.stedwards.edu/prod",
	StocktonUniversity:                      "https://pssb.stockton.edu/prod",
	TempleUniversity:                        "https://prd-wlssb.temple.edu/prod8",
	TennesseeStateUniversity:                "https://bannerssb.tnstate.edu/pls/PROD",
	TexasSouthernUniversity:                 "https://banner.tsu.edu:7777/pls/orasso",
	TexasTechUniversitySystem:               "https://ssb.texastech.edu/pls/TTUSPRD",
	ThomasJeffersonUniversity:               "https://banner.jefferson.edu/pls/tju",
	UniversityOfAlabama:                     "https://ssb.ua.edu/pls/PROD",
	UniversityOfAlabamaAtBirmingham:         "https://ssb.it.uab.edu/pls/sctprod",
	UniversityOfAlabamainHuntsville:         "https://sierra.uah.edu:9021/PROD",
	UniversityOfCentralArkansas:             "https://ssbprod.uca.edu/PROD",
	UniversityOfDallas:                      "https://ssb-prod.ec.udallas.edu/PROD",
	UniversityOfDetroitMercy:                "https://selfserv.udmercy.edu/PROD",
	UniversityOfGeorgia:                     "https://sis-ssb-prod.uga.edu/PROD",
	UniversityOfIllinoisSystem:              "https://ui2web1.apps.uillinois.edu/BANPROD1",
	UniversityOfMemphis:                     "https://ssb.bannerprod.memphis.edu/prod",
	UniversityOfMontevallo:                  "https://linuxss.montevallo.edu:9010/PROD",
	UniversityOfNewEnglandinMaine:           "https://ssb1.une.edu:4443/pls/prod",
	UniversityOfNewMexico:                   "https://www8.unm.edu/pls/banp",
	UniversityOfOklahoma:                    "https://ssb.ou.edu/pls/PROD",
	UniversityOfSanFrancisco:                "https://hebe.usfca.edu/prod",
	UniversityOfSouthFlorida:                "https://usfonline.admin.usf.edu/pls/prod",
	UniversityOfTennesseeAtMartin:           "https://banner8.utm.edu/prod",
	UniversityOfTennesseeKnoxville:          "https://bannerssb.utk.edu/kbanpr",
	UniversityOfTexasatElPaso:               "https://www.goldmine.utep.edu/prod/owa",
	UniversityOfTheDistrictOfColumbia:       "https://ssb-prod.ec.udc.edu/PROD",
	UniversityOfthePacific:                  "https://bssprd.ec.pacific.edu/PROD",
	UniversitySystemOfGeorgia:               "https://gsw.gabest.usg.edu/pls/B420",
	VirginiaCommonwealthUniversity:          "https://ssb.vcu.edu/proddad",
	VirginiaStateUniversity:                 "https://ssb-prod.ec.vsu.edu/BNPROD",
	WakeForestUniversity:                    "https://ssb.win.wfu.edu/BANPROD",
	WeberStateUniversity:                    "https://selfservice.weber.edu/pls/proddad",
	WesternNewMexicoUniversity:              "https://sapphire.wnmu.edu/wnmudad",
	WestGeorgiaTechnicalCollege:             "https://bannersso.westgatech.edu/pls/ban8",
	WestVirginiaNetwork:                     "https://rand.wvnet.edu:9780/MCTCPROD",
	WichitaStateUniversity:                  "https://ssbprod.wichita.edu/PROD",
	WilliamAndMary:                          "https://banweb.wm.edu/pls/PROD",
	WilliamPatersonUniversity:               "https://selfservice.wpunj.edu/PROD",
	WinthropUniversity:                      "https://ssb.winthrop.edu/prod",
	WrightStateUniversity:                   "https://wingsexpress.wright.edu/pls/PROD",
}

// SupportedColleges is a list of all supported Ellucian colleges
var SupportedColleges = map[College]string{
	AlabamaUniversityMontgomery:             "Alabama University Montgomery",
	AmericanUniversityInCairo:               "American University in Cairo",
	AppalachianStateUniversity:              "Appalachian State University",
	ArkansasTechUniversity:                  "Arkansas Tech University",
	BatesCollege:                            "Bates College",
	BridgewaterStateUniversity:              "Bridgewater State University",
	BrownUniversity:                         "Brown University",
	BryantUniversity:                        "Bryant University",
	CanisiusCollege:                         "Canisius College",
	CentralNewMexicoCommunityCollege:        "Central New Mexico Community College",
	ClarkAtlantaUniversity:                  "Clark Atlanta University",
	ColoradoCommunityCollegeSystem:          "Colorado Community College System",
	ColoradoSchoolOfMines:                   "Colorado School of Mines",
	ConnecticutCollege:                      "Connecticut College",
	CreightonUniversity:                     "Creighton University",
	DeVryUniversity:                         "De Vry University",
	DrexelUniversity:                        "Drexel University",
	EastCarolinaUniversity:                  "East Carolina University",
	EasternKentuckyUniversity:               "Eastern Kentucky University",
	EasternOregonUniversity:                 "Eastern Oregon University",
	EmporiaStateUniversity:                  "Emporia State University",
	FashionInstituteOfTechnology:            "Fashion institute of Technology",
	FerrisStateUniversity:                   "Ferris State University",
	FloridaInstituteOfTechnology:            "Florida institute of Technology",
	GeorgeMasonUniversity:                   "George Mason University",
	GeorgiaStateUniversity:                  "Georgia State University",
	GeorgiaTech:                             "Georgia Tech",
	HarperCollege:                           "Harper College",
	HawaiiPacificUniversity:                 "Hawaii Pacific University",
	HofstraUniversity:                       "Hofstra University",
	HowardUniversity:                        "Howard University",
	IllinoisInstituteOfTechnology:           "Illinois institute of Technology",
	ImperialValleyCollege:                   "Imperial Valley College",
	JohnCarrollUniversity:                   "John Carroll University",
	KennesawStateUniversityinGeorgia:        "Kennesaw State University in Georgia",
	KentuckyStateUniversity:                 "Kentucky State University",
	LakeSuperiorStateUniversity:             "Lake Superior State University",
	LewisUniverisity:                        "Lewis Univerisity",
	LongwoodUniversity:                      "Longwood University",
	LouisianasCommunityAndTechnicalColleges: "Louisianas Community and Technical Colleges",
	MontclairStateUniversity:                "Montclair State University",
	MorehouseCollege:                        "Morehouse College",
	MorganStateUniversity:                   "Morgan State University",
	NewCollegeOfFlorida:                     "New College of Florida",
	NorthCarolinaCentralUniversity:          "North Carolina Central University",
	NorthernMichiganUniversity:              "Northern Michigan University",
	NorthwesternStateUniversity:             "Northwestern State University",
	NovaSoutheasternUniversity:              "Nova Southeastern University",
	OaklandUniversity:                       "Oakland University",
	OklahomaBaptistUniversity:               "Oklahoma Baptist University",
	OklahomaStateUniversity:                 "Oklahoma State University",
	OralRobertsUniversity:                   "Oral Roberts University",
	PaceUniversity:                          "Pace University",
	PlymouthStateUniversity:                 "Plymouth State University",
	PurdueUniversity:                        "Purdue University",
	PurdueUniversityNorthwest:               "Purdue University Northwest",
	RamapoCollegeOfNewJersey:                "Ramapo College of New Jersey",
	RhodesCollege:                           "Rhodes College",
	RockhurstUniversity:                     "Rockhurst University",
	RowanUniversity:                         "Rowan University",
	RutgersUniversity:                       "Rutgers University",
	SanDiegoStateUniversity:                 "San Diego State University",
	SouthCarolinaStateUniversity:            "South Carolina State University",
	SouthernIllinoisUniversityEdwardsville:  "Southern Illinois University Edwardsville",
	SouthernUtahUniversity:                  "Southern Utah University",
	SpelmanCollege:                          "Spelman College",
	StEdwardsUniversityInAustin:             "St Edwards University in Austin",
	StocktonUniversity:                      "Stockton University",
	TempleUniversity:                        "Temple University",
	TennesseeStateUniversity:                "Tennessee State University",
	TexasSouthernUniversity:                 "Texas Southern University",
	TexasTechUniversitySystem:               "Texas Tech University System",
	ThomasJeffersonUniversity:               "Thomas Jefferson University",
	UniversityOfAlabama:                     "University of Alabama",
	UniversityOfAlabamaAtBirmingham:         "University of Alabama At Birmingham",
	UniversityOfAlabamainHuntsville:         "University of Alabamain Huntsville",
	UniversityOfCaliforniaBerkeley:          "University of California Berkeley",
	UniversityOfCaliforniaSantaBarbara:      "University of California Santa Barbara",
	UniversityOfCentralArkansas:             "University of Central Arkansas",
	UniversityOfDallas:                      "University of Dallas",
	UniversityOfDetroitMercy:                "University of Detroit Mercy",
	UniversityOfGeorgia:                     "University of Georgia",
	UniversityOfIllinoisSystem:              "University of Illinois System",
	UniversityOfMemphis:                     "University of Memphis",
	UniversityOfMontevallo:                  "University of Montevallo",
	UniversityOfNewEnglandinMaine:           "University of New England in Maine",
	UniversityOfNewMexico:                   "University of New Mexico",
	UniversityOfOklahoma:                    "University of Oklahoma",
	UniversityOfSanFrancisco:                "University of San Francisco",
	UniversityOfSouthFlorida:                "University of South Florida",
	UniversityOfTennesseeAtMartin:           "University of Tennessee At Martin",
	UniversityOfTennesseeKnoxville:          "University of Tennessee Knoxville",
	UniversityOfTexasatElPaso:               "University of Texasat El Paso",
	UniversityOfTheDistrictOfColumbia:       "University of The District of Columbia",
	UniversityOfthePacific:                  "University of the Pacific",
	UniversityOfWashington:                  "University of Washington",
	UniversitySystemOfGeorgia:               "University System of Georgia",
	VirginiaCommonwealthUniversity:          "Virginia Commonwealth University",
	VirginiaStateUniversity:                 "Virginia State University",
	WakeForestUniversity:                    "Wake Forest University",
	WashingtonStateUniversity:               "Washington State University",
	WeberStateUniversity:                    "Weber State University",
	WesternNewMexicoUniversity:              "Western New Mexico University",
	WestGeorgiaTechnicalCollege:             "West Georgia Technical College",
	WestVirginiaNetwork:                     "West Virginia Network",
	WichitaStateUniversity:                  "Wichita State University",
	WilliamAndMary:                          "William and Mary",
	WilliamPatersonUniversity:               "William Paterson University",
	WinthropUniversity:                      "Winthrop University",
	WrightStateUniversity:                   "Wright State University",
}

// DefaultCourseDataFormValues returns a mutable copy of the original map
func DefaultCourseDataFormValues() url.Values {
	copy := make(url.Values, len(defaultCourseDataFormValues))
	for key, values := range defaultCourseDataFormValues {
		copy[key] = values
	}
	return copy
}

var defaultCourseDataFormValues = url.Values{
	"sel_subj":      []string{"dummy"},
	"sel_day":       []string{"dummy"},
	"sel_schd":      []string{"dummy", "%25"},
	"sel_insm":      []string{"dummy"},
	"sel_camp":      []string{"dummy"},
	"sel_levl":      []string{"dummy", "%25"},
	"sel_sess":      []string{"dummy"},
	"sel_instr":     []string{"dummy", "%25"},
	"sel_ptrm":      []string{"dummy"},
	"sel_attr":      []string{"dummy", "%25"},
	"sel_title":     []string{""},
	"sel_from_cred": []string{""},
	"sel_to_cred":   []string{""},
	"begin_hh":      []string{"0"},
	"begin_mi":      []string{"0"},
	"begin_ap":      []string{"a"},
	"end_hh":        []string{"0"},
	"end_mi":        []string{"0"},
	"end_ap":        []string{"a"},
	"term_in":       []string{""},
}

// InvalidTerms are terms that should be ingored when parsing
var InvalidTerms = []string{"(", ")", "None", "Language", "Continuing"}
