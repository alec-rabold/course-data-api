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
	AlabamaUniversityMontgomery             string = "AUM"
	AmericanUniversityInCairo               string = "AUCEgypt"
	AppalachianStateUniversity              string = "AppState"
	ArkansasTechUniversity                  string = "ATU"
	BatesCollege                            string = "Bates"
	BridgewaterStateUniversity              string = "Bridgew"
	BrownUniversity                         string = "Brown"
	BryantUniversity                        string = "Bryant"
	CanisiusCollege                         string = "Canisius"
	CentralNewMexicoCommunityCollege        string = "CNM"
	ClarkAtlantaUniversity                  string = "CAU"
	ColoradoCommunityCollegeSystem          string = "CCCS"
	ColoradoSchoolofMines                   string = "Mines"
	ConnecticutCollege                      string = "ConnColl"
	CreightonUniversity                     string = "Creighton"
	DeVryUniversity                         string = "DeVry"
	DrexelUniversity                        string = "Drexel"
	EastCarolinaUniversity                  string = "ECU"
	EasternKentuckyUniversity               string = "EKU"
	EasternOregonUniversity                 string = "EOU"
	EmporiaStateUniversity                  string = "Emporia"
	FashionInstituteofTechnology            string = "FITNYC"
	FerrisStateUniversity                   string = "Ferris"
	FloridaInstituteofTechnology            string = "FIT"
	GeorgeMasonUniversity                   string = "GMU"
	GeorgiaStateUniversity                  string = "GSU"
	GeorgiaTech                             string = "GeorgiaTech"
	HarperCollege                           string = "Harper"
	HawaiiPacificUniversity                 string = "HPU"
	HofstraUniversity                       string = "Hofstra"
	HowardUniversity                        string = "Howard"
	IllinoisInstituteofTechnology           string = "IIT"
	ImperialValleyCollege                   string = "Imperial"
	JohnCarrollUniversity                   string = "JCU"
	KennesawStateUniversityinGeorgia        string = "Kennesaw"
	KentuckyStateUniversity                 string = "KYSU"
	LakeSuperiorStateUniversity             string = "LSSU"
	LewisUniverisity                        string = "Lewis"
	LongwoodUniversity                      string = "Longwood"
	LouisianasCommunityAndTechnicalColleges string = "LCTCS"
	MontclairStateUniversity                string = "Montclair"
	MontgomeryCollege                       string = "Montgomery"
	MorehouseCollege                        string = "Morehouse"
	MorganStateUniversity                   string = "Morgan"
	NewCollegeofFlorida                     string = "NCF"
	NorthCarolinaCentralUniversity          string = "NCCU"
	NorthernMichiganUniversity              string = "NMU"
	NorthwesternStateUniversity             string = "NSULA"
	NovaSoutheasternUniversity              string = "Nova"
	OaklandUniversity                       string = "Oakland"
	OklahomaBaptistUniversity               string = "OKBU"
	OklahomaStateUniversity                 string = "OKState"
	OralRobertsUniversity                   string = "ORU"
	PaceUniversity                          string = "Pace"
	PlymouthStateUniversity                 string = "Plymouth"
	PurdueUniversity                        string = "Purdue"
	PurdueUniversityNorthwest               string = "PurdueNW"
	RamapoCollegeofNewJersey                string = "Ramapo"
	RhodesCollege                           string = "Rhodes"
	RockhurstUniversity                     string = "Rockhurst"
	RowanUniversity                         string = "Rowan"
	RutgersUniversity                       string = "Rutgers"
	SanDiegoStateUniversity                 string = "SDSU"
	SouthCarolinaStateUniversity            string = "SCSU"
	SouthernIllinoisUniversityEdwardsville  string = "SIUE"
	SouthernUtahUniversity                  string = "SUU"
	SpelmanCollege                          string = "Spelman"
	StEdwardsUniversityInAustin             string = "StEdwards"
	StocktonUniversity                      string = "Stockton"
	TempleUniversity                        string = "Temple"
	TennesseeStateUniversity                string = "TNState"
	TexasSouthernUniversity                 string = "TSU"
	TexasTechUniversitySystem               string = "TexasTech"
	ThomasJeffersonUniversity               string = "Jefferson"
	UniversityofAlabama                     string = "UA"
	UniversityOfAlabamaAtBirmingham         string = "UAB"
	UniversityofAlabamainHuntsville         string = "UAH"
	UniversityOfCaliforniaBerkeley          string = "Berkeley"
	UniversityOfCaliforniaSantaBarbara      string = "UCSB"
	UniversityofCentralArkansas             string = "UCA"
	UniversityofDallas                      string = "UDallas"
	UniversityofDetroitMercy                string = "UDMercy"
	UniversityofGeorgia                     string = "UGA"
	UniversityofIllinoisSystem              string = "UIllinois"
	UniversityofMemphis                     string = "Memphis"
	UniversityOfMontevallo                  string = "Montevallo"
	UniversityofNewEnglandinMaine           string = "UNE"
	UniversityofNewMexico                   string = "UNM"
	UniversityOfOklahoma                    string = "OU"
	UniversityofSanFrancisco                string = "UFCA"
	UniversityOfSouthFlorida                string = "USF"
	UniversityOfTennesseeAtMartin           string = "UTM"
	UniversityOfTennesseeKnoxville          string = "UTK"
	UniversityofTennesseeKnoxville          string = "UTK"
	UniversityofTexasatElPaso               string = "UTEP"
	UniversityOfTheDistrictOfColumbia       string = "UDC"
	UniversityofthePacific                  string = "Pacific"
	UniversityOfWashington                  string = "UW"
	UniversitySystemOfGeorgia               string = "USG"
	UniversitySystemofGeorgia               string = "USG"
	VirginiaCommonwealthUniversity          string = "VCU"
	VirginiaStateUniversity                 string = "VSU"
	WakeForestUniversity                    string = "WFU"
	WashingtonStateUniversity               string = "WSU"
	WeberStateUniversity                    string = "Weber"
	WesternNewMexicoUniversity              string = "WNMU"
	WestGeorgiaTechnicalCollege             string = "WestGATech"
	WestVirginiaNetwork                     string = "WVNet"
	WichitaStateUniversity                  string = "Wichita"
	WilliamAndMary                          string = "WM"
	WilliamPatersonUniversity               string = "WPUNJ"
	WinthropUniversity                      string = "Winthrop"
	WrightStateUniversity                   string = "Wright"

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
	EllucianDataAlabamaUniversityMontgomery             string = "https://senator.aum.edu/prod"
	EllucianDataAmericanUniversityInCairo               string = "https://ssb-prod.ec.aucegypt.edu:9100/PROD"
	EllucianDataAppalachianStateUniversity              string = "https://bannerweb.appstate.edu/pls/PROD"
	EllucianDataArkansasTechUniversity                  string = "https://ssbprod.atu.edu/pls/PROD"
	EllucianDataBatesCollege                            string = "https://gg-bprod.bates.edu/bprod"
	EllucianDataBridgewaterStateUniversity              string = "https://infobear.bridgew.edu/BANP"
	EllucianDataBrownUniversity                         string = "https://selfservice.brown.edu/ss"
	EllucianDataBryantUniversity                        string = "https://ssb-prod.bryantec.bryant.edu/PROD"
	EllucianDataCanisiusCollege                         string = "https://asterope.canisius.edu/pls/prod"
	EllucianDataCentralNewMexicoCommunityCollege        string = "https://banner.cnm.edu/proddad"
	EllucianDataClarkAtlantaUniversity                  string = "https://ssb-prod.ec.cau.edu/PROD"
	EllucianDataColoradoCommunityCollegeSystem          string = "https://erpdnssb.cccs.edu/PRODFRCC"
	EllucianDataColoradoSchoolOfMines                   string = "https://banner.mines.edu/prod/owa"
	EllucianDataConnecticutCollege                      string = "https://ssbprod.conncoll.edu:9112/CONN"
	EllucianDataCreightonUniversity                     string = "https://thenest.creighton.edu/PROD"
	EllucianDataDeVryUniversity                         string = "https://ssb-prod.ec.devry.edu/PROD"
	EllucianDataDrexelUniversity                        string = "https://banner.drexel.edu/pls/duprod"
	EllucianDataEastCarolinaUniversity                  string = "https://ssbprd-banner.ecu.edu/DAD_PROD"
	EllucianDataEasternKentuckyUniversity               string = "https://web4s.eku.edu/prod"
	EllucianDataEasternOregonUniversity                 string = "https://ssb-prod.ec.eou.edu/PROD"
	EllucianDataEmporiaStateUniversity                  string = "https://ssb.emporia.edu/pls/prod"
	EllucianDataFashionInstituteOfTechnology            string = "https://banweb02.fitnyc.edu/prodssb"
	EllucianDataFerrisStateUniversity                   string = "https://banner.ferris.edu:9000/pls/GOLD"
	EllucianDataFloridaInstituteOfTechnology            string = "https://nssb-p.adm.fit.edu/prod"
	EllucianDataGeorgeMasonUniversity                   string = "https://patriotweb.gmu.edu/pls/prod"
	EllucianDataGeorgiaStateUniversity                  string = "https://www.gosolar.gsu.edu/bprod"
	EllucianDataGeorgiaTech                             string = "https://oscar.gatech.edu/pls/bprod"
	EllucianDataHarperCollege                           string = "https://student-self-service.harpercollege.edu/prod"
	EllucianDataHawaiiPacificUniversity                 string = "https://bweb.hpu.edu:4443/hpud"
	EllucianDataHofstraUniversity                       string = "https://hofstraonline.hofstra.edu/pls/HPRO"
	EllucianDataHowardUniversity                        string = "https://ssb-prod.ec.howard.edu/PROD"
	EllucianDataIllinoisInstituteOfTechnology           string = "https://ssb.iit.edu/bnrprd"
	EllucianDataImperialValleyCollege                   string = "https://webstar.imperial.edu/pls/ivc01"
	EllucianDataJohnCarrollUniversity                   string = "https://web4.jcu.edu/pjcu"
	EllucianDataKennesawStateUniversityinGeorgia        string = "https://owlexpress.kennesaw.edu/prodban"
	EllucianDataKentuckyStateUniversity                 string = "https://ssb-prod.ec.kysu.edu/PROD"
	EllucianDataLakeSuperiorStateUniversity             string = "https://bssmain.lssu.edu:9060/pls/PROD8"
	EllucianDataLewisUniverisity                        string = "http://ssb.lewisu.edu:9000/PROD"
	EllucianDataLongwoodUniversity                      string = "https://bannerssb.longwood.edu/bnr8prod"
	EllucianDataLouisianasCommunityAndTechnicalColleges string = "https://ssb-prod.ec.lctcs.edu/PROD_NSHORE"
	EllucianDataMontclairStateUniversity                string = "https://ssb.montclair.edu/PROD"
	EllucianDataMorehouseCollege                        string = "https://ssb-prod.ec.morehouse.edu/MC"
	EllucianDataMorganStateUniversity                   string = "https://lbssbnprod.morgan.edu/nprod"
	EllucianDataNewCollegeOfFlorida                     string = "https://banner.ncf.edu/pls/ncpo"
	EllucianDataNorthCarolinaCentralUniversity          string = "https://nccussbprod.nccu.edu/pls/NCCUPROD"
	EllucianDataNorthernMichiganUniversity              string = "https://bssbprod.nmu.edu:8443/PROD"
	EllucianDataNorthwesternStateUniversity             string = "https://connect.nsula.edu/prod"
	EllucianDataNovaSoutheasternUniversity              string = "https://ssb.nova.edu/pls/PROD"
	EllucianDataOaklandUniversity                       string = "https://sail.oakland.edu/PROD"
	EllucianDataOklahomaBaptistUniversity               string = "https://self-service.okbu.edu/PROD"
	EllucianDataOklahomaStateUniversity                 string = "https://ssb.okstate.edu/OKM"
	EllucianDataOralRobertsUniversity                   string = "https://vision.oru.edu/prod"
	EllucianDataPaceUniversity                          string = "https://ban8prodss.pace.edu/prod"
	EllucianDataPlymouthStateUniversity                 string = "https://ssb.plymouth.edu/psc1"
	EllucianDataPurdueUniversity                        string = "https://selfservice.mypurdue.purdue.edu/prod"
	EllucianDataPurdueUniversityNorthwest               string = "https://ssb-prod.pnw.edu/dbServer_prod"
	EllucianDataRamapoCollegeOfNewJersey                string = "https://ssb.ramapo.edu/pls/RCNJ"
	EllucianDataRhodesCollege                           string = "https://banweb.rhodes.edu/prod"
	EllucianDataRockhurstUniversity                     string = "https://rockweb.rockhurst.edu/pls/dbp"
	EllucianDataRowanUniversity                         string = "https://banner9.rowan.edu/ords/ssb"
	EllucianDataRutgersUniversity                       string = "https://banweb.rutgers.edu/PROD"
	EllucianDataSouthCarolinaStateUniversity            string = "https://bannerweb.scsu.edu/prod"
	EllucianDataSouthernIllinoisUniversityEdwardsville  string = "https://banner.siue.edu/ssb"
	EllucianDataSouthernUtahUniversity                  string = "https://bannerxe.suu.edu/prod"
	EllucianDataSpelmanCollege                          string = "https://ssb-prod.ec.spelman.edu/PROD"
	EllucianDataStEdwardsUniversityInAustin             string = "https://ssb.stedwards.edu/prod"
	EllucianDataStocktonUniversity                      string = "https://pssb.stockton.edu/prod"
	EllucianDataTempleUniversity                        string = "https://prd-wlssb.temple.edu/prod8"
	EllucianDataTennesseeStateUniversity                string = "https://bannerssb.tnstate.edu/pls/PROD"
	EllucianDataTexasSouthernUniversity                 string = "https://banner.tsu.edu:7777/pls/orasso"
	EllucianDataTexasTechUniversitySystem               string = "https://ssb.texastech.edu/pls/TTUSPRD"
	EllucianDataThomasJeffersonUniversity               string = "https://banner.jefferson.edu/pls/tju"
	EllucianDataUniversityOfAlabama                     string = "https://ssb.ua.edu/pls/PROD"
	EllucianDataUniversityOfAlabamaAtBirmingham         string = "https://ssb.it.uab.edu/pls/sctprod"
	EllucianDataUniversityOfAlabamainHuntsville         string = "https://sierra.uah.edu:9021/PROD"
	EllucianDataUniversityOfCentralArkansas             string = "https://ssbprod.uca.edu/PROD"
	EllucianDataUniversityOfDallas                      string = "https://ssb-prod.ec.udallas.edu/PROD"
	EllucianDataUniversityOfDetroitMercy                string = "https://selfserv.udmercy.edu/PROD"
	EllucianDataUniversityOfGeorgia                     string = "https://sis-ssb-prod.uga.edu/PROD"
	EllucianDataUniversityOfIllinoisSystem              string = "https://ui2web1.apps.uillinois.edu/BANPROD1"
	EllucianDataUniversityOfMemphis                     string = "https://ssb.bannerprod.memphis.edu/prod"
	EllucianDataUniversityOfMontevallo                  string = "https://linuxss.montevallo.edu:9010/PROD"
	EllucianDataUniversityOfNewEnglandinMaine           string = "https://ssb1.une.edu:4443/pls/prod"
	EllucianDataUniversityOfNewMexico                   string = "https://www8.unm.edu/pls/banp"
	EllucianDataUniversityOfOklahoma                    string = "https://ssb.ou.edu/pls/PROD"
	EllucianDataUniversityOfSanFrancisco                string = "https://hebe.usfca.edu/prod"
	EllucianDataUniversityOfSouthFlorida                string = "https://usfonline.admin.usf.edu/pls/prod"
	EllucianDataUniversityOfTennesseeAtMartin           string = "https://banner8.utm.edu/prod"
	EllucianDataUniversityOfTennesseeKnowxville         string = "https://bannerssb.utk.edu/kbanpr"
	EllucianDataUniversityOfTennesseeKnoxville          string = "https://bannerssb.utk.edu/kbanpr"
	EllucianDataUniversityOfTexasatElPaso               string = "https://www.goldmine.utep.edu/prod/owa"
	EllucianDataUniversityOfTheDistrictOfColumbia       string = "https://ssb-prod.ec.udc.edu/PROD"
	EllucianDataUniversityOfthePacific                  string = "https://bssprd.ec.pacific.edu/PROD"
	EllucianDataUniversitySystemOfGeorgia               string = "https://gsw.gabest.usg.edu/pls/B420"
	EllucianDataVirginiaCommonwealthUniversity          string = "https://ssb.vcu.edu/proddad"
	EllucianDataVirginiaStateUniversity                 string = "https://ssb-prod.ec.vsu.edu/BNPROD"
	EllucianDataWakeForestUniversity                    string = "https://ssb.win.wfu.edu/BANPROD"
	EllucianDataWeberStateUniversity                    string = "https://selfservice.weber.edu/pls/proddad"
	EllucianDataWesternNewMexicoUniversity              string = "https://sapphire.wnmu.edu/wnmudad"
	EllucianDataWestGeorgiaTechnicalCollege             string = "https://bannersso.westgatech.edu/pls/ban8"
	EllucianDataWestVirginiaNetwork                     string = "https://rand.wvnet.edu:9780/MCTCPROD"
	EllucianDataWichitaStateUniversity                  string = "https://ssbprod.wichita.edu/PROD"
	EllucianDataWilliamAndMary                          string = "https://banweb.wm.edu/pls/PROD"
	EllucianDataWilliamPatersonUniversity               string = "https://selfservice.wpunj.edu/PROD"
	EllucianDataWinthropUniversity                      string = "https://ssb.winthrop.edu/prod"
	EllucianDataWrightStateUniversity                   string = "https://wingsexpress.wright.edu/pls/PROD"
)

// EllucianSupportedColleges is a list of all supported Ellucian colleges
var EllucianSupportedColleges = map[string]string{
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
	ColoradoSchoolofMines:                   "Colorado School of Mines",
	ConnecticutCollege:                      "Connecticut College",
	CreightonUniversity:                     "Creighton University",
	DeVryUniversity:                         "De Vry University",
	DrexelUniversity:                        "Drexel University",
	EastCarolinaUniversity:                  "East Carolina University",
	EasternKentuckyUniversity:               "Eastern Kentucky University",
	EasternOregonUniversity:                 "Eastern Oregon University",
	EmporiaStateUniversity:                  "Emporia State University",
	FashionInstituteofTechnology:            "Fashion institute of Technology",
	FerrisStateUniversity:                   "Ferris State University",
	FloridaInstituteofTechnology:            "Florida institute of Technology",
	GeorgeMasonUniversity:                   "George Mason University",
	GeorgiaStateUniversity:                  "Georgia State University",
	GeorgiaTech:                             "Georgia Tech",
	HarperCollege:                           "Harper College",
	HawaiiPacificUniversity:                 "Hawaii Pacific University",
	HofstraUniversity:                       "Hofstra University",
	HowardUniversity:                        "Howard University",
	IllinoisInstituteofTechnology:           "Illinois institute of Technology",
	ImperialValleyCollege:                   "Imperial Valley College",
	JohnCarrollUniversity:                   "John Carroll University",
	KennesawStateUniversityinGeorgia:        "Kennesaw State University in Georgia",
	KentuckyStateUniversity:                 "Kentucky State University",
	LakeSuperiorStateUniversity:             "Lake Superior State University",
	LewisUniverisity:                        "Lewis Univerisity",
	LongwoodUniversity:                      "Longwood University",
	LouisianasCommunityAndTechnicalColleges: "Louisianas Community and Technical Colleges",
	MontclairStateUniversity:                "Montclair State University",
	MontgomeryCollege:                       "Montgomery College",
	MorehouseCollege:                        "Morehouse College",
	MorganStateUniversity:                   "Morgan State University",
	NewCollegeofFlorida:                     "New Collegeof Florida",
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
	RamapoCollegeofNewJersey:                "Ramapo College of New Jersey",
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
	UniversityofAlabama:                     "University of Alabama",
	UniversityOfAlabamaAtBirmingham:         "University of Alabama At Birmingham",
	UniversityofAlabamainHuntsville:         "University of Alabamain Huntsville",
	UniversityOfCaliforniaBerkeley:          "University of California Berkeley",
	UniversityOfCaliforniaSantaBarbara:      "University of California Santa Barbara",
	UniversityofCentralArkansas:             "University of Central Arkansas",
	UniversityofDallas:                      "University of Dallas",
	UniversityofDetroitMercy:                "University of Detroit Mercy",
	UniversityofGeorgia:                     "University of Georgia",
	UniversityofIllinoisSystem:              "University of Illinois System",
	UniversityofMemphis:                     "University of Memphis",
	UniversityOfMontevallo:                  "University of Montevallo",
	UniversityofNewEnglandinMaine:           "University of New England in Maine",
	UniversityofNewMexico:                   "University of New Mexico",
	UniversityOfOklahoma:                    "University of Oklahoma",
	UniversityofSanFrancisco:                "University of San Francisco",
	UniversityOfSouthFlorida:                "University of South Florida",
	UniversityOfTennesseeAtMartin:           "University of Tennessee At Martin",
	UniversityOfTennesseeKnoxville:          "University of Tennessee Knoxville",
	UniversityofTexasatElPaso:               "University of Texasat El Paso",
	UniversityOfTheDistrictOfColumbia:       "University of The District of Columbia",
	UniversityofthePacific:                  "University of the Pacific",
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
