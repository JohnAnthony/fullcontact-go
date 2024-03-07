package fullcontact

type PersonRequestOption func(pr *PersonRequest)

type PersonRequest struct {
	Emails     []string    `json:"emails,omitempty"`
	Phones     []string    `json:"phones,omitempty"`
	DataFilter []string    `json:"dataFilter,omitempty"`
	Maid       []string    `json:"maids,omitempty"`
	Location   *Location   `json:"location,omitempty"`
	Name       *PersonName `json:"name,omitempty"`
	Profiles   []*Profile  `json:"profiles,omitempty"`
	WebhookUrl string      `json:"webhookUrl,omitempty"`
	RecordId   string      `json:"recordId,omitempty"`
	PersonId   string      `json:"personId,omitempty"`
	PartnerId  string      `json:"partnerId,omitempty"`
	LiNonId    string      `json:"li_nonid,omitempty"`
	Confidence string      `json:"confidence,omitempty"`
	Infer      *bool       `json:"infer,omitempty"`
	Placekey   string      `json:"placekey,omitempty"`
	MaxMaids   int         `json:"maxMaids,omitempty"`
	PanoramaId string      `json:"panoramaId,omitempty"`
	HemType    string      `json:"hemType,omitempty"`
	MaxEmails  int         `json:"maxEmails,omitempty"`
}

func NewPersonRequest(option ...PersonRequestOption) (*PersonRequest, error) {
	pr := &PersonRequest{}
	for _, opt := range option {
		opt(pr)
	}
	return pr, nil
}

func (pr *PersonRequest) isQueryable() bool {
	return pr.Emails != nil ||
		pr.Phones != nil ||
		pr.Profiles != nil ||
		pr.Maid != nil ||
		isPopulated(pr.RecordId) ||
		isPopulated(pr.PersonId) ||
		isPopulated(pr.PartnerId) ||
		isPopulated(pr.LiNonId)
}

func validatePersonRequest(pr *PersonRequest) error {
	if isPopulated(pr.HemType) &&
		pr.HemType != "md5" &&
		pr.HemType != "sha1" &&
		pr.HemType != "sha256" {
		return NewFullContactError("HemType value can only be 'md5', 'sha1', 'sha256'")
	}

	if isPopulated(pr.Confidence) &&
		pr.Confidence != "LOW" &&
		pr.Confidence != "MED" &&
		pr.Confidence != "HIGH" &&
		pr.Confidence != "MAX" {
		return NewFullContactError("Confidence value can only be 'LOW', 'MED', 'HIGH', 'MAX'")
	}
	if !pr.isQueryable() {
		if (pr.Location == nil && pr.Name == nil && !isPopulated(pr.Placekey)) || (isPopulated(pr.Placekey) && pr.Name != nil) {
			return nil
		} else if pr.Location != nil && pr.Name != nil {
			// Validating Location fields
			if isPopulated(pr.Location.AddressLine1) &&
				((isPopulated(pr.Location.City) &&
					(isPopulated(pr.Location.Region) || isPopulated(pr.Location.RegionCode))) ||
					(isPopulated(pr.Location.PostalCode))) {
				// Validating Name fields
				if (isPopulated(pr.Name.Full)) ||
					(isPopulated(pr.Name.Given) && isPopulated(pr.Name.Family)) {
					return nil
				} else {
					return NewFullContactError("Name data requires full name or given and family name")
				}
			} else {
				return NewFullContactError(
					"Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
			}
		}
		return NewFullContactError(
			"If you want to use 'location'(or placekey) or 'name' as an input, both must be present and they must have non-blank values")
	}

	return nil
}

func WithEmail(email string) PersonRequestOption {
	return func(pr *PersonRequest) {
		if pr.Emails == nil {
			pr.Emails = make([]string, 0)
		}
		pr.Emails = append(pr.Emails, email)
	}
}

func WithEmails(emails []string) PersonRequestOption {
	return func(pr *PersonRequest) {
		if pr.Emails == nil {
			pr.Emails = make([]string, 0)
		}
		pr.Emails = append(pr.Emails, emails...)
	}
}

func WithPhone(phone string) PersonRequestOption {
	return func(pr *PersonRequest) {
		if pr.Phones == nil {
			pr.Phones = make([]string, 0)
		}
		pr.Phones = append(pr.Phones, phone)
	}
}

func WithPhones(phones []string) PersonRequestOption {
	return func(pr *PersonRequest) {
		if pr.Phones == nil {
			pr.Phones = make([]string, 0)
		}
		pr.Phones = append(pr.Phones, phones...)
	}
}

func WithDataFilter(dataFilter string) PersonRequestOption {
	return func(pr *PersonRequest) {
		if pr.DataFilter == nil {
			pr.DataFilter = make([]string, 0)
		}
		pr.DataFilter = append(pr.DataFilter, dataFilter)
	}
}

func WithDataFilters(dataFilters []string) PersonRequestOption {
	return func(pr *PersonRequest) {
		if pr.DataFilter == nil {
			pr.DataFilter = make([]string, 0)
		}
		pr.DataFilter = append(pr.DataFilter, dataFilters...)
	}
}

func WithMaid(maid string) PersonRequestOption {
	return func(pr *PersonRequest) {
		if pr.Maid == nil {
			pr.Maid = make([]string, 0)
		}
		pr.Maid = append(pr.Maid, maid)
	}
}

func WithMaids(maids []string) PersonRequestOption {
	return func(pr *PersonRequest) {
		if pr.Maid == nil {
			pr.Maid = make([]string, 0)
		}
		pr.Maid = append(pr.Maid, maids...)
	}
}

func WithLocation(location *Location) PersonRequestOption {
	return func(pr *PersonRequest) {
		pr.Location = location
	}
}

func WithWebhookUrl(webhookUrl string) PersonRequestOption {
	return func(pr *PersonRequest) {
		pr.WebhookUrl = webhookUrl
	}
}

func WithRecordId(recordId string) PersonRequestOption {
	return func(pr *PersonRequest) {
		pr.RecordId = recordId
	}
}

func WithPersonId(personId string) PersonRequestOption {
	return func(pr *PersonRequest) {
		pr.PersonId = personId
	}
}

func WithPartnerId(partnerId string) PersonRequestOption {
	return func(pr *PersonRequest) {
		pr.PartnerId = partnerId
	}
}

func WithLiNonId(liNonId string) PersonRequestOption {
	return func(pr *PersonRequest) {
		pr.LiNonId = liNonId
	}
}

func WithName(name *PersonName) PersonRequestOption {
	return func(pr *PersonRequest) {
		pr.Name = name
	}
}

func WithProfile(profile *Profile) PersonRequestOption {
	return func(pr *PersonRequest) {
		if pr.Profiles == nil {
			pr.Profiles = make([]*Profile, 0)
		}
		pr.Profiles = append(pr.Profiles, profile)
	}
}

func WithProfiles(profiles []*Profile) PersonRequestOption {
	return func(pr *PersonRequest) {
		if pr.Profiles == nil {
			pr.Profiles = make([]*Profile, 0)
		}
		pr.Profiles = append(pr.Profiles, profiles...)
	}
}

func WithConfidence(confidence string) PersonRequestOption {
	return func(pr *PersonRequest) {
		pr.Confidence = confidence
	}
}

func WithInfer(infer bool) PersonRequestOption {
	return func(pr *PersonRequest) {
		pr.Infer = &infer
	}
}

func WithPlacekey(placekey string) PersonRequestOption {
	return func(pr *PersonRequest) {
		pr.Placekey = placekey
	}
}

func WithMaxMaids(maxMaids int) PersonRequestOption {
	return func(pr *PersonRequest) {
		pr.MaxMaids = maxMaids
	}
}

func WithPanoramaID(panoramaId string) PersonRequestOption {
	return func(pr *PersonRequest) {
		pr.PanoramaId = panoramaId
	}
}

func WithHemType(hemType string) PersonRequestOption {
	return func(pr *PersonRequest) {
		pr.HemType = hemType
	}
}

func WIthMaxEmails(maxEmails int) PersonRequestOption {
	return func(pr *PersonRequest) {
		pr.MaxEmails = maxEmails
	}
}
