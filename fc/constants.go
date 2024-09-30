package fullcontact

const (
	version                    = "1.4.0"
	userAgent                  = "FullContact_Go_Client_V" + version
	FcApiKey                   = "FC_API_KEY"
	FCGoClientTestType         = "FCGoClientTestType"
	defaultBaseUrl             = "https://api.fullcontact.com/v3/"
	personEnrichUrl            = "person.enrich"
	companyEnrichUrl           = "company.enrich"
	identityMapUrl             = "identity.map"
	identityResolveUrl         = "identity.resolve"
	identityResolveWithTagsUrl = "identity.resolve?tags=true"
	identityMapResolveUrl      = "identity.mapResolve"
	identityDeleteUrl          = "identity.delete"
	tagsCreateUrl              = "tags.create"
	tagsGetUrl                 = "tags.get"
	tagsDeleteUrl              = "tags.delete"
	audienceCreateUrl          = "audience.create"
	audienceDownloadUrl        = "audience.download"
	permissionCreateUrl        = "permission.create"
	permissionDeleteUrl        = "permission.delete"
	permissionFindUrl          = "permission.find"
	permissionCurrentUrl       = "permission.current"
	permissionVerifyUrl        = "permission.verify"
	verifySignalsUrl           = "verify.signals"
	verifyMatchUrl             = "verify.match"
	verifyActivityUrl          = "verify.activity"
)
