package fullcontact

import "os"

type CredentialsProvider interface {
	GetApiKey() string
}

type StaticCredentialsProvider struct {
	apiKey string
}

func NewStaticCredentialsProvider(apiKey string) (StaticCredentialsProvider, error) {
	if !isPopulated(apiKey) {
		return StaticCredentialsProvider{}, NewFullContactError("API Key can't be empty")
	} else {
		return StaticCredentialsProvider{apiKey: apiKey}, nil
	}
}

func (scp StaticCredentialsProvider) GetApiKey() string {
	return scp.apiKey
}

type DefaultCredentialsProvider struct {
	apiKey string
}

func NewDefaultCredentialsProvider(envVar string) (DefaultCredentialsProvider, error) {
	apiKey := os.Getenv(envVar)
	if !isPopulated(apiKey) {
		return DefaultCredentialsProvider{}, NewFullContactError("Couldn't find valid API Key from ENV variable: " + envVar)
	} else {
		return DefaultCredentialsProvider{apiKey: apiKey}, nil
	}
}

func (dcp DefaultCredentialsProvider) GetApiKey() string {
	return dcp.apiKey
}
