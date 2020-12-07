package fullcontact

import (
	"encoding/json"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestMarshallNewResolveRequest(t *testing.T) {
	emails := []string{"test1@gmail.com", "test2@outlook.com"}
	profile1, err := NewProfile(WithUrl("https://twitter.com/mcreedy"))
	assert.NoError(t, err)
	profile2, err := NewProfile(WithUrl("https://twitter.com/mcreedytest"))
	assert.NoError(t, err)
	requestJson := "{\"emails\":[\"marianrd97@outlook.com\",\"test1@gmail.com\",\"test2@outlook.com\"],\"phones\":[\"123-4567890\"],\"maids\":[\"abcd-123-abcd-1234-abcdlkjhasdfgh\",\"1234-snbk-lkldiemvmruixp-2kdp-vdm\"],\"location\":{\"addressLine1\":\"123/23\",\"addressLine2\":\"Some Street\",\"city\":\"Denver\",\"region\":\"Denver\",\"regionCode\":\"123123\",\"postalCode\":\"23124\"},\"name\":{\"full\":\"Marian C Reed\"},\"profiles\":[{\"url\":\"https://twitter.com/mcreedy\"},{\"url\":\"https://twitter.com/mcreedytest\"}],\"recordId\":\"customer123\",\"personId\":\"VS1OPPPPvxHcCNPezUbvYBCDEAOdSj5AI0adsA2bLmh12345\"}"
	pr, err := NewResolveRequest(
		WithNameForResolve(&PersonName{Full: "Marian C Reed"}),
		WithEmailForResolve("marianrd97@outlook.com"),
		WithEmailsForResolve(emails),
		WithPhoneForResolve("123-4567890"),
		WithLocationForResolve(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))),
		WithProfileForResolve(profile1),
		WithProfileForResolve(profile2),
		WithMaidForResolve("abcd-123-abcd-1234-abcdlkjhasdfgh"),
		WithMaidForResolve("1234-snbk-lkldiemvmruixp-2kdp-vdm"),
		WithRecordIdForResolve("customer123"),
		WithPersonIdForResolve("VS1OPPPPvxHcCNPezUbvYBCDEAOdSj5AI0adsA2bLmh12345"))
	assert.NoError(t, err)
	reqBytes, err := json.Marshal(pr)
	assert.NoError(t, err)
	assert.Equal(t, requestJson, string(reqBytes))
}

func TestNewResolveRequestWithoutNameAndLocation(t *testing.T) {
	_, err := NewResolveRequest(WithEmailForResolve("marianrd97@outlook.com"))
	assert.NoError(t, err)
}

func TestNewResolveRequestWithNameOnly(t *testing.T) {
	_, err := NewResolveRequest(WithEmailForResolve("marianrd97@outlook.com"),
		WithNameForResolve(&PersonName{
			Full: "Marian C Reed",
		}))
	assert.EqualError(t, err, "FullContactError: If you want to use 'location' or 'name' as an input, both must be present and they must have non-blank values")
}

func TestNewResolveRequestWithLocationOnly(t *testing.T) {
	_, err := NewResolveRequest(
		WithEmailForResolve("marianrd97@outlook.com"),
		WithLocationForResolve(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))))
	assert.EqualError(t, err, "FullContactError: If you want to use 'location' or 'name' as an input, both must be present and they must have non-blank values")
}

func TestNewResolveRequestWithLocationWithoutAddressLine1(t *testing.T) {
	_, err := NewResolveRequest(
		WithEmailForResolve("marianrd97@outlook.com"),
		WithNameForResolve(NewPersonName(WithFull("Test Name"))),
		WithLocationForResolve(NewLocation(
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))))
	assert.EqualError(t, err, "FullContactError: Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewResolveRequestWithLocationOnlyAddressLine1(t *testing.T) {
	_, err := NewResolveRequest(
		WithEmailForResolve("marianrd97@outlook.com"),
		WithNameForResolve(NewPersonName(WithFull("Test Name"))),
		WithLocationForResolve(NewLocation(
			WithAddressLine1("123/23"))))
	assert.EqualError(t, err, "FullContactError: Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewResolveRequestWithLocationWithAddressLine1AndCity(t *testing.T) {
	_, err := NewResolveRequest(
		WithEmailForResolve("marianrd97@outlook.com"),
		WithNameForResolve(NewPersonName(WithFull("Test Name"))),
		WithLocationForResolve(NewLocation(
			WithAddressLine1("123/23"),
			WithCity("Denver"))))
	assert.EqualError(t, err, "FullContactError: Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewResolveRequestWithLocationWithAddressLine1AndRegion(t *testing.T) {
	_, err := NewResolveRequest(
		WithEmailForResolve("marianrd97@outlook.com"),
		WithNameForResolve(NewPersonName(WithFull("Test Name"))),
		WithLocationForResolve(NewLocation(
			WithAddressLine1("123/23"),
			WithRegionCode("123123"))))
	assert.EqualError(t, err, "FullContactError: Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewResolveRequestWithValidLocation1(t *testing.T) {
	_, err := NewResolveRequest(
		WithNameForResolve(&PersonName{Full: "Marian C Reed"}),
		WithLocationForResolve(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("12343"))))
	assert.NoError(t, err)
}

func TestNewResolveRequestWithValidLocation2(t *testing.T) {
	_, err := NewResolveRequest(
		WithNameForResolve(&PersonName{Full: "Marian C Reed"}),
		WithLocationForResolve(NewLocation(
			WithAddressLine1("123/23"),
			WithCity("Denver"),
			WithRegionCode("123123"))))
	assert.NoError(t, err)
}

func TestNewResolveRequestWithValidLocation3(t *testing.T) {
	_, err := NewResolveRequest(
		WithNameForResolve(&PersonName{Full: "Marian C Reed"}),
		WithLocationForResolve(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("123123"))))
	assert.NoError(t, err)
}

func TestNewResolveRequestWithValidName(t *testing.T) {
	_, err := NewResolveRequest(
		WithNameForResolve(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForResolve(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	assert.NoError(t, err)
}

func TestNilIdentityMapRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.IdentityMap(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Resolve Request can't be nil")
}

func TestNilIdentityResolveRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.IdentityResolve(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Resolve Request can't be nil")
}

func TestNilIdentityDeleteRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.IdentityDelete(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Resolve Request can't be nil")
}

func TestInvalidIdentityMapRequest1(t *testing.T) {
	fcTestClient := fullContactClient{}
	rr, err := NewResolveRequest(WithPersonIdForResolve("personId"))
	assert.NoError(t, err)
	resp := <-fcTestClient.IdentityMap(rr)
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Invalid map request, person id must be empty")
}

func TestInvalidIdentityMapRequest2(t *testing.T) {
	fcTestClient := fullContactClient{}
	rr, err := NewResolveRequest(WithRecordIdForResolve("recordId"))
	assert.NoError(t, err)
	resp := <-fcTestClient.IdentityMap(rr)
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Invalid map request, Any of Email, Phone, SocialProfile, Name and Location must be present")
}

func TestInvalidIdentityResolveRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	rr, err := NewResolveRequest(WithRecordIdForResolve("recordId"), WithPersonIdForResolve("personId"))
	assert.NoError(t, err)
	resp := <-fcTestClient.IdentityResolve(rr)
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Both record id and person id are populated, please select one")
}

func TestInvalidIdentityDeleteRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	rr, err := NewResolveRequest(WithPersonIdForResolve("personId"))
	assert.NoError(t, err)
	resp := <-fcTestClient.IdentityDelete(rr)
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: recordId param must be specified")
}

func TestInvalidMapTagRequest1(t *testing.T) {
	fcTestClient := fullContactClient{}
	rr, err := NewResolveRequest(WithRecordIdForResolve("recordId"), WithTagForResolve(NewTag(WithTagKey("key"))))
	assert.NoError(t, err)
	resp := <-fcTestClient.IdentityMap(rr)
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Both Key and Value must be populated for adding a Tag")
}

func TestInvalidMapTagRequest2(t *testing.T) {
	fcTestClient := fullContactClient{}
	rr, err := NewResolveRequest(WithRecordIdForResolve("recordId"), WithTagForResolve(NewTag(WithTagValue("value"))))
	assert.NoError(t, err)
	resp := <-fcTestClient.IdentityMap(rr)
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Both Key and Value must be populated for adding a Tag")
}

func TestInvalidMapTagRequest3(t *testing.T) {
	fcTestClient := fullContactClient{}
	rr, err := NewResolveRequest(WithRecordIdForResolve("recordId"), WithTagForResolve(NewTag(WithTagKey("ke'y"), WithTagValue("value"))))
	assert.NoError(t, err)
	resp := <-fcTestClient.IdentityMap(rr)
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Both Key and Value must be populated for adding a Tag")
}
