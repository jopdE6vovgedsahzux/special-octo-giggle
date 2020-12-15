package config

var websiteCategories = []string{
	"ALDR",
	"ANON",
	"COMM",
	"COMT",
	"CTRL",
	"CULTR",
	"DATE",
	"ECON",
	"ENV",
	"FILE",
	"GAME",
	"GMB",
	"GOVT",
	"GRP",
	"HACK",
	"HATE",
	"HOST",
	"HUMR",
	"IGO",
	"LGBT",
	"MILX",
	"MMED",
	"NEWS",
	"POLR",
	"PORN",
	"PROV",
	"PUBH",
	"REL",
	"SRCH",
	"XED",
}

// Sharing settings
type Sharing struct {
	UploadResults bool `json:"upload_results"`
}

// Advanced settings
type Advanced struct {
	SendCrashReports bool `json:"send_crash_reports"`
}

// Nettests related settings
type Nettests struct {
	WebsitesURLLimit             int64    `json:"websites_url_limit"`
	WebsitesEnabledCategoryCodes []string `json:"websites_enabled_category_codes"`
}
