// Package vault provides access to the G Suite Vault API.
//
// See https://developers.google.com/vault
//
// Usage example:
//
//   import "google.golang.org/api/vault/v1"
//   ...
//   vaultService, err := vault.New(oauthHttpClient)
package vault // import "google.golang.org/api/vault/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "vault:v1"
const apiName = "vault"
const apiVersion = "v1"
const basePath = "https://vault.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// Manage your eDiscovery data
	EdiscoveryScope = "https://www.googleapis.com/auth/ediscovery"

	// View your eDiscovery data
	EdiscoveryReadonlyScope = "https://www.googleapis.com/auth/ediscovery.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Matters = NewMattersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Matters *MattersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewMattersService(s *Service) *MattersService {
	rs := &MattersService{s: s}
	rs.Exports = NewMattersExportsService(s)
	rs.Holds = NewMattersHoldsService(s)
	return rs
}

type MattersService struct {
	s *Service

	Exports *MattersExportsService

	Holds *MattersHoldsService
}

func NewMattersExportsService(s *Service) *MattersExportsService {
	rs := &MattersExportsService{s: s}
	return rs
}

type MattersExportsService struct {
	s *Service
}

func NewMattersHoldsService(s *Service) *MattersHoldsService {
	rs := &MattersHoldsService{s: s}
	rs.Accounts = NewMattersHoldsAccountsService(s)
	return rs
}

type MattersHoldsService struct {
	s *Service

	Accounts *MattersHoldsAccountsService
}

func NewMattersHoldsAccountsService(s *Service) *MattersHoldsAccountsService {
	rs := &MattersHoldsAccountsService{s: s}
	return rs
}

type MattersHoldsAccountsService struct {
	s *Service
}

// AccountInfo: Accounts to search
type AccountInfo struct {
	// Emails: A set of accounts to search.
	Emails []string `json:"emails,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Emails") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Emails") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AccountInfo) MarshalJSON() ([]byte, error) {
	type NoMethod AccountInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AddHeldAccountResult: A status detailing the status of each account
// creation, and the
// HeldAccount, if successful.
type AddHeldAccountResult struct {
	// Account: If present, this account was successfully created.
	Account *HeldAccount `json:"account,omitempty"`

	// Status: This represents the success status. If failed, check message.
	Status *Status `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Account") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Account") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AddHeldAccountResult) MarshalJSON() ([]byte, error) {
	type NoMethod AddHeldAccountResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AddHeldAccountsRequest: Add a list of accounts to a hold.
type AddHeldAccountsRequest struct {
	// AccountIds: Account ids to identify which accounts to add. Only
	// account_ids or only
	// emails should be specified, but not both.
	AccountIds []string `json:"accountIds,omitempty"`

	// Emails: Emails to identify which accounts to add. Only emails or only
	// account_ids
	// should be specified, but not both.
	Emails []string `json:"emails,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AccountIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccountIds") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AddHeldAccountsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod AddHeldAccountsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AddHeldAccountsResponse: Response for batch create held accounts.
type AddHeldAccountsResponse struct {
	// Responses: The list of responses, in the same order as the batch
	// request.
	Responses []*AddHeldAccountResult `json:"responses,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Responses") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Responses") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AddHeldAccountsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod AddHeldAccountsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AddMatterPermissionsRequest: Add an account with the permission
// specified. The role cannot be owner.
// If an account already has a role in the matter, it will
// be
// overwritten.
type AddMatterPermissionsRequest struct {
	// CcMe: Only relevant if send_emails is true.
	// True to CC requestor in the email message.
	// False to not CC requestor.
	CcMe bool `json:"ccMe,omitempty"`

	// MatterPermission: The MatterPermission to add.
	MatterPermission *MatterPermission `json:"matterPermission,omitempty"`

	// SendEmails: True to send notification email to the added
	// account.
	// False to not send notification email.
	SendEmails bool `json:"sendEmails,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CcMe") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CcMe") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AddMatterPermissionsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod AddMatterPermissionsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CloseMatterRequest: Close a matter by ID.
type CloseMatterRequest struct {
}

// CloseMatterResponse: Response to a CloseMatterRequest.
type CloseMatterResponse struct {
	// Matter: The updated matter, with state CLOSED.
	Matter *Matter `json:"matter,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Matter") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Matter") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CloseMatterResponse) MarshalJSON() ([]byte, error) {
	type NoMethod CloseMatterResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CloudStorageFile: An export file on cloud storage
type CloudStorageFile struct {
	// BucketName: The cloud storage bucket name of this export file.
	// Can be used in cloud storage JSON/XML API.
	BucketName string `json:"bucketName,omitempty"`

	// Md5Hash: The md5 hash of the file.
	Md5Hash string `json:"md5Hash,omitempty"`

	// ObjectName: The cloud storage object name of this export file.
	// Can be used in cloud storage JSON/XML API.
	ObjectName string `json:"objectName,omitempty"`

	// Size: The size of the export file.
	Size int64 `json:"size,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "BucketName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BucketName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CloudStorageFile) MarshalJSON() ([]byte, error) {
	type NoMethod CloudStorageFile
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CloudStorageSink: Export sink for cloud storage files.
type CloudStorageSink struct {
	// Files: Output only. The exported files on cloud storage.
	Files []*CloudStorageFile `json:"files,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Files") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Files") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CloudStorageSink) MarshalJSON() ([]byte, error) {
	type NoMethod CloudStorageSink
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CorpusQuery: Corpus specific queries.
type CorpusQuery struct {
	// DriveQuery: Details pertaining to Drive holds. If set, corpus must be
	// Drive.
	DriveQuery *HeldDriveQuery `json:"driveQuery,omitempty"`

	// GroupsQuery: Details pertaining to Groups holds. If set, corpus must
	// be Groups.
	GroupsQuery *HeldGroupsQuery `json:"groupsQuery,omitempty"`

	// HangoutsChatQuery: Details pertaining to Hangouts Chat holds. If set,
	// corpus must be
	// Hangouts Chat.
	HangoutsChatQuery *HeldHangoutsChatQuery `json:"hangoutsChatQuery,omitempty"`

	// MailQuery: Details pertaining to mail holds. If set, corpus must be
	// mail.
	MailQuery *HeldMailQuery `json:"mailQuery,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DriveQuery") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DriveQuery") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CorpusQuery) MarshalJSON() ([]byte, error) {
	type NoMethod CorpusQuery
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DriveExportOptions: The options for Drive export.
type DriveExportOptions struct {
	// IncludeAccessInfo: Set to true to include access level information
	// for users
	// with <a
	// href="https://support.google.com/vault/answer/6099459#metadata">indire
	// ct access</a>
	// to files.
	IncludeAccessInfo bool `json:"includeAccessInfo,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IncludeAccessInfo")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IncludeAccessInfo") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DriveExportOptions) MarshalJSON() ([]byte, error) {
	type NoMethod DriveExportOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DriveOptions: Drive search advanced options
type DriveOptions struct {
	// IncludeTeamDrives: Set to true to include Team Drive.
	IncludeTeamDrives bool `json:"includeTeamDrives,omitempty"`

	// VersionDate: Search the versions of the Drive file
	// as of the reference date. These timestamps are in GMT and
	// rounded down to the given date.
	VersionDate string `json:"versionDate,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IncludeTeamDrives")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IncludeTeamDrives") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DriveOptions) MarshalJSON() ([]byte, error) {
	type NoMethod DriveOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// Export: An export
type Export struct {
	// CloudStorageSink: Output only. Export sink for cloud storage files.
	CloudStorageSink *CloudStorageSink `json:"cloudStorageSink,omitempty"`

	// CreateTime: Output only. The time when the export was created.
	CreateTime string `json:"createTime,omitempty"`

	// ExportOptions: Advanced options of the export.
	ExportOptions *ExportOptions `json:"exportOptions,omitempty"`

	// Id: Output only. The generated export ID.
	Id string `json:"id,omitempty"`

	// MatterId: Output only. The matter ID.
	MatterId string `json:"matterId,omitempty"`

	// Name: The export name.
	Name string `json:"name,omitempty"`

	// Query: The search query being exported.
	Query *Query `json:"query,omitempty"`

	// Requester: Output only. The requester of the export.
	Requester *UserInfo `json:"requester,omitempty"`

	// Stats: Output only. Export statistics.
	Stats *ExportStats `json:"stats,omitempty"`

	// Status: Output only. The export status.
	//
	// Possible values:
	//   "EXPORT_STATUS_UNSPECIFIED" - The status is unspecified.
	//   "COMPLETED" - The export completed.
	//   "FAILED" - The export failed.
	//   "IN_PROGRESS" - The export is still being executed.
	Status string `json:"status,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CloudStorageSink") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CloudStorageSink") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Export) MarshalJSON() ([]byte, error) {
	type NoMethod Export
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ExportOptions: Export advanced options
type ExportOptions struct {
	// DriveOptions: Option available for Drive export.
	DriveOptions *DriveExportOptions `json:"driveOptions,omitempty"`

	// GroupsOptions: Option available for groups export.
	GroupsOptions *GroupsExportOptions `json:"groupsOptions,omitempty"`

	// HangoutsChatOptions: Option available for hangouts chat export.
	HangoutsChatOptions *HangoutsChatExportOptions `json:"hangoutsChatOptions,omitempty"`

	// MailOptions: Option available for mail export.
	MailOptions *MailExportOptions `json:"mailOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DriveOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DriveOptions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExportOptions) MarshalJSON() ([]byte, error) {
	type NoMethod ExportOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ExportStats: Stats of an export.
type ExportStats struct {
	// ExportedArtifactCount: The number of documents already processed by
	// the export.
	ExportedArtifactCount int64 `json:"exportedArtifactCount,omitempty,string"`

	// SizeInBytes: The size of export in bytes.
	SizeInBytes int64 `json:"sizeInBytes,omitempty,string"`

	// TotalArtifactCount: The number of documents to be exported.
	TotalArtifactCount int64 `json:"totalArtifactCount,omitempty,string"`

	// ForceSendFields is a list of field names (e.g.
	// "ExportedArtifactCount") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExportedArtifactCount") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ExportStats) MarshalJSON() ([]byte, error) {
	type NoMethod ExportStats
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GroupsExportOptions: The options for groups export.
type GroupsExportOptions struct {
	// ExportFormat: The export format for groups export.
	//
	// Possible values:
	//   "EXPORT_FORMAT_UNSPECIFIED" - No export format specified.
	//   "MBOX" - MBOX as export format.
	//   "PST" - PST as export format
	ExportFormat string `json:"exportFormat,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExportFormat") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExportFormat") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GroupsExportOptions) MarshalJSON() ([]byte, error) {
	type NoMethod GroupsExportOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HangoutsChatExportOptions: The options for hangouts chat export.
type HangoutsChatExportOptions struct {
	// ExportFormat: The export format for hangouts chat export.
	//
	// Possible values:
	//   "EXPORT_FORMAT_UNSPECIFIED" - No export format specified.
	//   "MBOX" - MBOX as export format.
	//   "PST" - PST as export format
	ExportFormat string `json:"exportFormat,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExportFormat") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExportFormat") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HangoutsChatExportOptions) MarshalJSON() ([]byte, error) {
	type NoMethod HangoutsChatExportOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HangoutsChatInfo: Accounts to search
type HangoutsChatInfo struct {
	// RoomId: A set of rooms to search.
	RoomId []string `json:"roomId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "RoomId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "RoomId") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HangoutsChatInfo) MarshalJSON() ([]byte, error) {
	type NoMethod HangoutsChatInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HangoutsChatOptions: Hangouts chat search advanced options
type HangoutsChatOptions struct {
	// IncludeRooms: Set to true to include rooms.
	IncludeRooms bool `json:"includeRooms,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IncludeRooms") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IncludeRooms") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HangoutsChatOptions) MarshalJSON() ([]byte, error) {
	type NoMethod HangoutsChatOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HeldAccount: An account being held in a particular hold. This
// structure is immutable.
// This can be either a single user or a google group, depending on the
// corpus.
type HeldAccount struct {
	// AccountId: The account's ID as provided by the
	// <a href="https://developers.google.com/admin-sdk/">Admin SDK</a>.
	AccountId string `json:"accountId,omitempty"`

	// HoldTime: When the account was put on hold.
	HoldTime string `json:"holdTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AccountId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccountId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HeldAccount) MarshalJSON() ([]byte, error) {
	type NoMethod HeldAccount
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HeldDriveQuery: Query options for Drive holds.
type HeldDriveQuery struct {
	// IncludeTeamDriveFiles: If true, include files in Team Drives in the
	// hold.
	IncludeTeamDriveFiles bool `json:"includeTeamDriveFiles,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "IncludeTeamDriveFiles") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IncludeTeamDriveFiles") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *HeldDriveQuery) MarshalJSON() ([]byte, error) {
	type NoMethod HeldDriveQuery
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HeldGroupsQuery: Query options for group holds.
type HeldGroupsQuery struct {
	// EndTime: The end time range for the search query. These timestamps
	// are in GMT and
	// rounded down to the start of the given date.
	EndTime string `json:"endTime,omitempty"`

	// StartTime: The start time range for the search query. These
	// timestamps are in GMT and
	// rounded down to the start of the given date.
	StartTime string `json:"startTime,omitempty"`

	// Terms: The search terms for the hold.
	Terms string `json:"terms,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HeldGroupsQuery) MarshalJSON() ([]byte, error) {
	type NoMethod HeldGroupsQuery
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HeldHangoutsChatQuery: Query options for hangouts chat holds.
type HeldHangoutsChatQuery struct {
	// IncludeRooms: If true, include rooms the user has participated in.
	IncludeRooms bool `json:"includeRooms,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IncludeRooms") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IncludeRooms") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HeldHangoutsChatQuery) MarshalJSON() ([]byte, error) {
	type NoMethod HeldHangoutsChatQuery
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HeldMailQuery: Query options for mail holds.
type HeldMailQuery struct {
	// EndTime: The end time range for the search query. These timestamps
	// are in GMT and
	// rounded down to the start of the given date.
	EndTime string `json:"endTime,omitempty"`

	// StartTime: The start time range for the search query. These
	// timestamps are in GMT and
	// rounded down to the start of the given date.
	StartTime string `json:"startTime,omitempty"`

	// Terms: The search terms for the hold.
	Terms string `json:"terms,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HeldMailQuery) MarshalJSON() ([]byte, error) {
	type NoMethod HeldMailQuery
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HeldOrgUnit: A organizational unit being held in a particular
// hold.
// This structure is immutable.
type HeldOrgUnit struct {
	// HoldTime: When the org unit was put on hold. This property is
	// immutable.
	HoldTime string `json:"holdTime,omitempty"`

	// OrgUnitId: The org unit's immutable ID as provided by the Admin SDK.
	OrgUnitId string `json:"orgUnitId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "HoldTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "HoldTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HeldOrgUnit) MarshalJSON() ([]byte, error) {
	type NoMethod HeldOrgUnit
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Hold: Represents a hold within Vault. A hold restricts purging
// of
// artifacts based on the combination of the query and accounts
// restrictions.
// A hold can be configured to either apply to an explicitly configured
// set
// of accounts, or can be applied to all members of an organizational
// unit.
type Hold struct {
	// Accounts: If set, the hold applies to the enumerated accounts and
	// org_unit must be
	// empty.
	Accounts []*HeldAccount `json:"accounts,omitempty"`

	// Corpus: The corpus to be searched.
	//
	// Possible values:
	//   "CORPUS_TYPE_UNSPECIFIED" - No corpus specified.
	//   "DRIVE" - Drive.
	//   "MAIL" - Mail.
	//   "GROUPS" - Groups.
	//   "HANGOUTS_CHAT" - Hangouts Chat.
	Corpus string `json:"corpus,omitempty"`

	// HoldId: The unique immutable ID of the hold. Assigned during
	// creation.
	HoldId string `json:"holdId,omitempty"`

	// Name: The name of the hold.
	Name string `json:"name,omitempty"`

	// OrgUnit: If set, the hold applies to all members of the
	// organizational unit and
	// accounts must be empty. This property is mutable. For groups
	// holds,
	// set the accounts field.
	OrgUnit *HeldOrgUnit `json:"orgUnit,omitempty"`

	// Query: The corpus-specific query. If set, the corpusQuery must match
	// corpus
	// type.
	Query *CorpusQuery `json:"query,omitempty"`

	// UpdateTime: The last time this hold was modified.
	UpdateTime string `json:"updateTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Accounts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Accounts") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Hold) MarshalJSON() ([]byte, error) {
	type NoMethod Hold
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListExportsResponse: The holds for a matter.
type ListExportsResponse struct {
	// Exports: The list of exports.
	Exports []*Export `json:"exports,omitempty"`

	// NextPageToken: Page token to retrieve the next page of results in the
	// list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Exports") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Exports") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListExportsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListExportsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListHeldAccountsResponse: Returns a list of held accounts for a hold.
type ListHeldAccountsResponse struct {
	// Accounts: The held accounts on a hold.
	Accounts []*HeldAccount `json:"accounts,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Accounts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Accounts") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListHeldAccountsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListHeldAccountsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListHoldsResponse: The holds for a matter.
type ListHoldsResponse struct {
	// Holds: The list of holds.
	Holds []*Hold `json:"holds,omitempty"`

	// NextPageToken: Page token to retrieve the next page of results in the
	// list.
	// If this is empty, then there are no more holds to list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Holds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Holds") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListHoldsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListHoldsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListMattersResponse: Provides the list of matters.
type ListMattersResponse struct {
	// Matters: List of matters.
	Matters []*Matter `json:"matters,omitempty"`

	// NextPageToken: Page token to retrieve the next page of results in the
	// list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Matters") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Matters") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListMattersResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListMattersResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MailExportOptions: The options for mail export.
type MailExportOptions struct {
	// ExportFormat: The export file format.
	//
	// Possible values:
	//   "EXPORT_FORMAT_UNSPECIFIED" - No export format specified.
	//   "MBOX" - MBOX as export format.
	//   "PST" - PST as export format
	ExportFormat string `json:"exportFormat,omitempty"`

	// ShowConfidentialModeContent: Set to true to export confidential mode
	// content
	ShowConfidentialModeContent bool `json:"showConfidentialModeContent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExportFormat") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExportFormat") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MailExportOptions) MarshalJSON() ([]byte, error) {
	type NoMethod MailExportOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MailOptions: Mail search advanced options
type MailOptions struct {
	// ExcludeDrafts: Set to true to exclude drafts.
	ExcludeDrafts bool `json:"excludeDrafts,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExcludeDrafts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExcludeDrafts") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MailOptions) MarshalJSON() ([]byte, error) {
	type NoMethod MailOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Matter: Represents a matter.
type Matter struct {
	// Description: The description of the matter.
	Description string `json:"description,omitempty"`

	// MatterId: The matter ID which is generated by the server.
	// Should be blank when creating a new matter.
	MatterId string `json:"matterId,omitempty"`

	// MatterPermissions: List of users and access to the matter. Currently
	// there is no programmer
	// defined limit on the number of permissions a matter can have.
	MatterPermissions []*MatterPermission `json:"matterPermissions,omitempty"`

	// Name: The name of the matter.
	Name string `json:"name,omitempty"`

	// State: The state of the matter.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - The matter has no specified state.
	//   "OPEN" - This matter is open.
	//   "CLOSED" - This matter is closed.
	//   "DELETED" - This matter is deleted.
	State string `json:"state,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Matter) MarshalJSON() ([]byte, error) {
	type NoMethod Matter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MatterPermission: Currently each matter only has one owner, and all
// others are collaborators.
// When an account is purged, its corresponding MatterPermission
// resources
// cease to exist.
type MatterPermission struct {
	// AccountId: The account id, as provided by <a
	// href="https://developers.google.com/admin-sdk/">Admin SDK</a>.
	AccountId string `json:"accountId,omitempty"`

	// Role: The user's role in this matter.
	//
	// Possible values:
	//   "ROLE_UNSPECIFIED" - No role assigned.
	//   "COLLABORATOR" - A collaborator to the matter.
	//   "OWNER" - The owner of the matter.
	Role string `json:"role,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AccountId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccountId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MatterPermission) MarshalJSON() ([]byte, error) {
	type NoMethod MatterPermission
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OrgUnitInfo: Org Unit to search
type OrgUnitInfo struct {
	// OrgUnitId: Org unit to search, as provided by the
	// <a href="https://developers.google.com/admin-sdk/directory/">Admin
	// SDK Directory API</a>.
	OrgUnitId string `json:"orgUnitId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OrgUnitId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OrgUnitId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OrgUnitInfo) MarshalJSON() ([]byte, error) {
	type NoMethod OrgUnitInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Query: A query definition relevant for search & export.
type Query struct {
	// AccountInfo: When 'ACCOUNT' is chosen as search method,
	// account_info needs to be specified.
	AccountInfo *AccountInfo `json:"accountInfo,omitempty"`

	// Corpus: The corpus to search.
	//
	// Possible values:
	//   "CORPUS_TYPE_UNSPECIFIED" - No corpus specified.
	//   "DRIVE" - Drive.
	//   "MAIL" - Mail.
	//   "GROUPS" - Groups.
	//   "HANGOUTS_CHAT" - Hangouts Chat.
	Corpus string `json:"corpus,omitempty"`

	// DataScope: The data source to search from.
	//
	// Possible values:
	//   "DATA_SCOPE_UNSPECIFIED" - No data scope specified.
	//   "ALL_DATA" - All available data.
	//   "HELD_DATA" - Data on hold.
	//   "UNPROCESSED_DATA" - Data not processed.
	DataScope string `json:"dataScope,omitempty"`

	// DriveOptions: For Drive search, specify more options in this field.
	DriveOptions *DriveOptions `json:"driveOptions,omitempty"`

	// EndTime: The end time range for the search query. These timestamps
	// are in GMT and
	// rounded down to the start of the given date.
	EndTime string `json:"endTime,omitempty"`

	// HangoutsChatInfo: When 'ROOM' is chosen as search method,
	// hangout_chats_info needs to be
	// specified. (read-only)
	HangoutsChatInfo *HangoutsChatInfo `json:"hangoutsChatInfo,omitempty"`

	// HangoutsChatOptions: For hangouts chat search, specify more options
	// in this field. (read-only)
	HangoutsChatOptions *HangoutsChatOptions `json:"hangoutsChatOptions,omitempty"`

	// MailOptions: For mail search, specify more options in this field.
	MailOptions *MailOptions `json:"mailOptions,omitempty"`

	// OrgUnitInfo: When 'ORG_UNIT' is chosen as as search method,
	// org_unit_info needs
	// to be specified.
	OrgUnitInfo *OrgUnitInfo `json:"orgUnitInfo,omitempty"`

	// SearchMethod: The search method to use.
	//
	// Possible values:
	//   "SEARCH_METHOD_UNSPECIFIED" - A search method must be specified. If
	// a request does not specify a
	// search method, it will be rejected.
	//   "ACCOUNT" - Will search all accounts provided in account_info.
	//   "ORG_UNIT" - Will search all accounts in the OU specified in
	// org_unit_info.
	//   "TEAM_DRIVE" - Will search for all accounts in the Team Drive
	// specified in
	// team_drive_info.
	//   "ENTIRE_ORG" - Will search for all accounts in the organization.
	// No need to set account_info or org_unit_info.
	//   "ROOM" - Will search in the Room specified in
	// hangout_chats_info. (read-only)
	SearchMethod string `json:"searchMethod,omitempty"`

	// StartTime: The start time range for the search query. These
	// timestamps are in GMT and
	// rounded down to the start of the given date.
	StartTime string `json:"startTime,omitempty"`

	// TeamDriveInfo: When 'TEAM_DRIVE' is chosen as search method,
	// team_drive_info needs to be
	// specified.
	TeamDriveInfo *TeamDriveInfo `json:"teamDriveInfo,omitempty"`

	// Terms: The corpus-specific
	// <a href="https://support.google.com/vault/answer/2474474">search
	// operators</a>
	// used to generate search results.
	Terms string `json:"terms,omitempty"`

	// TimeZone: The time zone name.
	// It should be an IANA TZ name, such as "America/Los_Angeles".
	// For more information, see
	// <a
	// href="https://en.wikipedia.org/wiki/List_of_tz_database_time_zones">Ti
	// me
	// Zone</a>.
	TimeZone string `json:"timeZone,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AccountInfo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccountInfo") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Query) MarshalJSON() ([]byte, error) {
	type NoMethod Query
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RemoveHeldAccountsRequest: Remove a list of accounts from a hold.
type RemoveHeldAccountsRequest struct {
	// AccountIds: Account ids to identify HeldAccounts to remove.
	AccountIds []string `json:"accountIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AccountIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccountIds") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RemoveHeldAccountsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod RemoveHeldAccountsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RemoveHeldAccountsResponse: Response for batch delete held accounts.
type RemoveHeldAccountsResponse struct {
	// Statuses: A list of statuses for deleted accounts. Results have
	// the
	// same order as the request.
	Statuses []*Status `json:"statuses,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Statuses") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Statuses") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RemoveHeldAccountsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod RemoveHeldAccountsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RemoveMatterPermissionsRequest: Remove an account as a matter
// collaborator.
type RemoveMatterPermissionsRequest struct {
	// AccountId: The account ID.
	AccountId string `json:"accountId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AccountId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccountId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RemoveMatterPermissionsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod RemoveMatterPermissionsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ReopenMatterRequest: Reopen a matter by ID.
type ReopenMatterRequest struct {
}

// ReopenMatterResponse: Response to a ReopenMatterRequest.
type ReopenMatterResponse struct {
	// Matter: The updated matter, with state OPEN.
	Matter *Matter `json:"matter,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Matter") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Matter") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ReopenMatterResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ReopenMatterResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TeamDriveInfo: Team Drives to search
type TeamDriveInfo struct {
	// TeamDriveIds: List of Team Drive ids, as provided by
	// <a
	// href="https://developers.google.com/drive">Drive API</a>.
	TeamDriveIds []string `json:"teamDriveIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "TeamDriveIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "TeamDriveIds") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TeamDriveInfo) MarshalJSON() ([]byte, error) {
	type NoMethod TeamDriveInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UndeleteMatterRequest: Undelete a matter by ID.
type UndeleteMatterRequest struct {
}

// UserInfo: User's information.
type UserInfo struct {
	// DisplayName: The displayed name of the user.
	DisplayName string `json:"displayName,omitempty"`

	// Email: The email address of the user.
	Email string `json:"email,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UserInfo) MarshalJSON() ([]byte, error) {
	type NoMethod UserInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "vault.matters.addPermissions":

type MattersAddPermissionsCall struct {
	s                           *Service
	matterId                    string
	addmatterpermissionsrequest *AddMatterPermissionsRequest
	urlParams_                  gensupport.URLParams
	ctx_                        context.Context
	header_                     http.Header
}

// AddPermissions: Adds an account as a matter collaborator.
func (r *MattersService) AddPermissions(matterId string, addmatterpermissionsrequest *AddMatterPermissionsRequest) *MattersAddPermissionsCall {
	c := &MattersAddPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.addmatterpermissionsrequest = addmatterpermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersAddPermissionsCall) Fields(s ...googleapi.Field) *MattersAddPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersAddPermissionsCall) Context(ctx context.Context) *MattersAddPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersAddPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersAddPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.addmatterpermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}:addPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.addPermissions" call.
// Exactly one of *MatterPermission or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *MatterPermission.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MattersAddPermissionsCall) Do(opts ...googleapi.CallOption) (*MatterPermission, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &MatterPermission{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Adds an account as a matter collaborator.",
	//   "flatPath": "v1/matters/{matterId}:addPermissions",
	//   "httpMethod": "POST",
	//   "id": "vault.matters.addPermissions",
	//   "parameterOrder": [
	//     "matterId"
	//   ],
	//   "parameters": {
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}:addPermissions",
	//   "request": {
	//     "$ref": "AddMatterPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "MatterPermission"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.close":

type MattersCloseCall struct {
	s                  *Service
	matterId           string
	closematterrequest *CloseMatterRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
	header_            http.Header
}

// Close: Closes the specified matter. Returns matter with updated
// state.
func (r *MattersService) Close(matterId string, closematterrequest *CloseMatterRequest) *MattersCloseCall {
	c := &MattersCloseCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.closematterrequest = closematterrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersCloseCall) Fields(s ...googleapi.Field) *MattersCloseCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersCloseCall) Context(ctx context.Context) *MattersCloseCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersCloseCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersCloseCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.closematterrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}:close")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.close" call.
// Exactly one of *CloseMatterResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *CloseMatterResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MattersCloseCall) Do(opts ...googleapi.CallOption) (*CloseMatterResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CloseMatterResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Closes the specified matter. Returns matter with updated state.",
	//   "flatPath": "v1/matters/{matterId}:close",
	//   "httpMethod": "POST",
	//   "id": "vault.matters.close",
	//   "parameterOrder": [
	//     "matterId"
	//   ],
	//   "parameters": {
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}:close",
	//   "request": {
	//     "$ref": "CloseMatterRequest"
	//   },
	//   "response": {
	//     "$ref": "CloseMatterResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.create":

type MattersCreateCall struct {
	s          *Service
	matter     *Matter
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a new matter with the given name and description. The
// initial state
// is open, and the owner is the method caller. Returns the created
// matter
// with default view.
func (r *MattersService) Create(matter *Matter) *MattersCreateCall {
	c := &MattersCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matter = matter
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersCreateCall) Fields(s ...googleapi.Field) *MattersCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersCreateCall) Context(ctx context.Context) *MattersCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.matter)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.create" call.
// Exactly one of *Matter or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Matter.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *MattersCreateCall) Do(opts ...googleapi.CallOption) (*Matter, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Matter{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new matter with the given name and description. The initial state\nis open, and the owner is the method caller. Returns the created matter\nwith default view.",
	//   "flatPath": "v1/matters",
	//   "httpMethod": "POST",
	//   "id": "vault.matters.create",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/matters",
	//   "request": {
	//     "$ref": "Matter"
	//   },
	//   "response": {
	//     "$ref": "Matter"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.delete":

type MattersDeleteCall struct {
	s          *Service
	matterId   string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified matter. Returns matter with updated
// state.
func (r *MattersService) Delete(matterId string) *MattersDeleteCall {
	c := &MattersDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersDeleteCall) Fields(s ...googleapi.Field) *MattersDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersDeleteCall) Context(ctx context.Context) *MattersDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.delete" call.
// Exactly one of *Matter or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Matter.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *MattersDeleteCall) Do(opts ...googleapi.CallOption) (*Matter, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Matter{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified matter. Returns matter with updated state.",
	//   "flatPath": "v1/matters/{matterId}",
	//   "httpMethod": "DELETE",
	//   "id": "vault.matters.delete",
	//   "parameterOrder": [
	//     "matterId"
	//   ],
	//   "parameters": {
	//     "matterId": {
	//       "description": "The matter ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}",
	//   "response": {
	//     "$ref": "Matter"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.get":

type MattersGetCall struct {
	s            *Service
	matterId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the specified matter.
func (r *MattersService) Get(matterId string) *MattersGetCall {
	c := &MattersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	return c
}

// View sets the optional parameter "view": Specifies which parts of the
// Matter to return in the response.
//
// Possible values:
//   "VIEW_UNSPECIFIED"
//   "BASIC"
//   "FULL"
func (c *MattersGetCall) View(view string) *MattersGetCall {
	c.urlParams_.Set("view", view)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersGetCall) Fields(s ...googleapi.Field) *MattersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *MattersGetCall) IfNoneMatch(entityTag string) *MattersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersGetCall) Context(ctx context.Context) *MattersGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.get" call.
// Exactly one of *Matter or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Matter.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *MattersGetCall) Do(opts ...googleapi.CallOption) (*Matter, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Matter{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the specified matter.",
	//   "flatPath": "v1/matters/{matterId}",
	//   "httpMethod": "GET",
	//   "id": "vault.matters.get",
	//   "parameterOrder": [
	//     "matterId"
	//   ],
	//   "parameters": {
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Specifies which parts of the Matter to return in the response.",
	//       "enum": [
	//         "VIEW_UNSPECIFIED",
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}",
	//   "response": {
	//     "$ref": "Matter"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery",
	//     "https://www.googleapis.com/auth/ediscovery.readonly"
	//   ]
	// }

}

// method id "vault.matters.list":

type MattersListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists matters the user has access to.
func (r *MattersService) List() *MattersListCall {
	c := &MattersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// PageSize sets the optional parameter "pageSize": The number of
// matters to return in the response.
// Default and maximum are 100.
func (c *MattersListCall) PageSize(pageSize int64) *MattersListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The pagination
// token as returned in the response.
func (c *MattersListCall) PageToken(pageToken string) *MattersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// State sets the optional parameter "state": If set, list only matters
// with that specific state. The default is listing
// matters of all states.
//
// Possible values:
//   "STATE_UNSPECIFIED"
//   "OPEN"
//   "CLOSED"
//   "DELETED"
func (c *MattersListCall) State(state string) *MattersListCall {
	c.urlParams_.Set("state", state)
	return c
}

// View sets the optional parameter "view": Specifies which parts of the
// matter to return in response.
//
// Possible values:
//   "VIEW_UNSPECIFIED"
//   "BASIC"
//   "FULL"
func (c *MattersListCall) View(view string) *MattersListCall {
	c.urlParams_.Set("view", view)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersListCall) Fields(s ...googleapi.Field) *MattersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *MattersListCall) IfNoneMatch(entityTag string) *MattersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersListCall) Context(ctx context.Context) *MattersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.list" call.
// Exactly one of *ListMattersResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListMattersResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MattersListCall) Do(opts ...googleapi.CallOption) (*ListMattersResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListMattersResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists matters the user has access to.",
	//   "flatPath": "v1/matters",
	//   "httpMethod": "GET",
	//   "id": "vault.matters.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The number of matters to return in the response.\nDefault and maximum are 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The pagination token as returned in the response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "state": {
	//       "description": "If set, list only matters with that specific state. The default is listing\nmatters of all states.",
	//       "enum": [
	//         "STATE_UNSPECIFIED",
	//         "OPEN",
	//         "CLOSED",
	//         "DELETED"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Specifies which parts of the matter to return in response.",
	//       "enum": [
	//         "VIEW_UNSPECIFIED",
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters",
	//   "response": {
	//     "$ref": "ListMattersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery",
	//     "https://www.googleapis.com/auth/ediscovery.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *MattersListCall) Pages(ctx context.Context, f func(*ListMattersResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "vault.matters.removePermissions":

type MattersRemovePermissionsCall struct {
	s                              *Service
	matterId                       string
	removematterpermissionsrequest *RemoveMatterPermissionsRequest
	urlParams_                     gensupport.URLParams
	ctx_                           context.Context
	header_                        http.Header
}

// RemovePermissions: Removes an account as a matter collaborator.
func (r *MattersService) RemovePermissions(matterId string, removematterpermissionsrequest *RemoveMatterPermissionsRequest) *MattersRemovePermissionsCall {
	c := &MattersRemovePermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.removematterpermissionsrequest = removematterpermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersRemovePermissionsCall) Fields(s ...googleapi.Field) *MattersRemovePermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersRemovePermissionsCall) Context(ctx context.Context) *MattersRemovePermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersRemovePermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersRemovePermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.removematterpermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}:removePermissions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.removePermissions" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *MattersRemovePermissionsCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Removes an account as a matter collaborator.",
	//   "flatPath": "v1/matters/{matterId}:removePermissions",
	//   "httpMethod": "POST",
	//   "id": "vault.matters.removePermissions",
	//   "parameterOrder": [
	//     "matterId"
	//   ],
	//   "parameters": {
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}:removePermissions",
	//   "request": {
	//     "$ref": "RemoveMatterPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.reopen":

type MattersReopenCall struct {
	s                   *Service
	matterId            string
	reopenmatterrequest *ReopenMatterRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Reopen: Reopens the specified matter. Returns matter with updated
// state.
func (r *MattersService) Reopen(matterId string, reopenmatterrequest *ReopenMatterRequest) *MattersReopenCall {
	c := &MattersReopenCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.reopenmatterrequest = reopenmatterrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersReopenCall) Fields(s ...googleapi.Field) *MattersReopenCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersReopenCall) Context(ctx context.Context) *MattersReopenCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersReopenCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersReopenCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.reopenmatterrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}:reopen")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.reopen" call.
// Exactly one of *ReopenMatterResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ReopenMatterResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MattersReopenCall) Do(opts ...googleapi.CallOption) (*ReopenMatterResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ReopenMatterResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Reopens the specified matter. Returns matter with updated state.",
	//   "flatPath": "v1/matters/{matterId}:reopen",
	//   "httpMethod": "POST",
	//   "id": "vault.matters.reopen",
	//   "parameterOrder": [
	//     "matterId"
	//   ],
	//   "parameters": {
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}:reopen",
	//   "request": {
	//     "$ref": "ReopenMatterRequest"
	//   },
	//   "response": {
	//     "$ref": "ReopenMatterResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.undelete":

type MattersUndeleteCall struct {
	s                     *Service
	matterId              string
	undeletematterrequest *UndeleteMatterRequest
	urlParams_            gensupport.URLParams
	ctx_                  context.Context
	header_               http.Header
}

// Undelete: Undeletes the specified matter. Returns matter with updated
// state.
func (r *MattersService) Undelete(matterId string, undeletematterrequest *UndeleteMatterRequest) *MattersUndeleteCall {
	c := &MattersUndeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.undeletematterrequest = undeletematterrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersUndeleteCall) Fields(s ...googleapi.Field) *MattersUndeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersUndeleteCall) Context(ctx context.Context) *MattersUndeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersUndeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersUndeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.undeletematterrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}:undelete")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.undelete" call.
// Exactly one of *Matter or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Matter.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *MattersUndeleteCall) Do(opts ...googleapi.CallOption) (*Matter, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Matter{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Undeletes the specified matter. Returns matter with updated state.",
	//   "flatPath": "v1/matters/{matterId}:undelete",
	//   "httpMethod": "POST",
	//   "id": "vault.matters.undelete",
	//   "parameterOrder": [
	//     "matterId"
	//   ],
	//   "parameters": {
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}:undelete",
	//   "request": {
	//     "$ref": "UndeleteMatterRequest"
	//   },
	//   "response": {
	//     "$ref": "Matter"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.update":

type MattersUpdateCall struct {
	s          *Service
	matterId   string
	matter     *Matter
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Updates the specified matter.
// This updates only the name and description of the matter, identified
// by
// matter id. Changes to any other fields are ignored.
// Returns the default view of the matter.
func (r *MattersService) Update(matterId string, matter *Matter) *MattersUpdateCall {
	c := &MattersUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.matter = matter
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersUpdateCall) Fields(s ...googleapi.Field) *MattersUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersUpdateCall) Context(ctx context.Context) *MattersUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.matter)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PUT", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.update" call.
// Exactly one of *Matter or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Matter.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *MattersUpdateCall) Do(opts ...googleapi.CallOption) (*Matter, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Matter{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified matter.\nThis updates only the name and description of the matter, identified by\nmatter id. Changes to any other fields are ignored.\nReturns the default view of the matter.",
	//   "flatPath": "v1/matters/{matterId}",
	//   "httpMethod": "PUT",
	//   "id": "vault.matters.update",
	//   "parameterOrder": [
	//     "matterId"
	//   ],
	//   "parameters": {
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}",
	//   "request": {
	//     "$ref": "Matter"
	//   },
	//   "response": {
	//     "$ref": "Matter"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.exports.create":

type MattersExportsCreateCall struct {
	s          *Service
	matterId   string
	export     *Export
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates an Export.
func (r *MattersExportsService) Create(matterId string, export *Export) *MattersExportsCreateCall {
	c := &MattersExportsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.export = export
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersExportsCreateCall) Fields(s ...googleapi.Field) *MattersExportsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersExportsCreateCall) Context(ctx context.Context) *MattersExportsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersExportsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersExportsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.export)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}/exports")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.exports.create" call.
// Exactly one of *Export or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Export.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *MattersExportsCreateCall) Do(opts ...googleapi.CallOption) (*Export, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Export{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates an Export.",
	//   "flatPath": "v1/matters/{matterId}/exports",
	//   "httpMethod": "POST",
	//   "id": "vault.matters.exports.create",
	//   "parameterOrder": [
	//     "matterId"
	//   ],
	//   "parameters": {
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}/exports",
	//   "request": {
	//     "$ref": "Export"
	//   },
	//   "response": {
	//     "$ref": "Export"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.exports.delete":

type MattersExportsDeleteCall struct {
	s          *Service
	matterId   string
	exportId   string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes an Export.
func (r *MattersExportsService) Delete(matterId string, exportId string) *MattersExportsDeleteCall {
	c := &MattersExportsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.exportId = exportId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersExportsDeleteCall) Fields(s ...googleapi.Field) *MattersExportsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersExportsDeleteCall) Context(ctx context.Context) *MattersExportsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersExportsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersExportsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}/exports/{exportId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
		"exportId": c.exportId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.exports.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *MattersExportsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes an Export.",
	//   "flatPath": "v1/matters/{matterId}/exports/{exportId}",
	//   "httpMethod": "DELETE",
	//   "id": "vault.matters.exports.delete",
	//   "parameterOrder": [
	//     "matterId",
	//     "exportId"
	//   ],
	//   "parameters": {
	//     "exportId": {
	//       "description": "The export ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}/exports/{exportId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.exports.get":

type MattersExportsGetCall struct {
	s            *Service
	matterId     string
	exportId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets an Export.
func (r *MattersExportsService) Get(matterId string, exportId string) *MattersExportsGetCall {
	c := &MattersExportsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.exportId = exportId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersExportsGetCall) Fields(s ...googleapi.Field) *MattersExportsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *MattersExportsGetCall) IfNoneMatch(entityTag string) *MattersExportsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersExportsGetCall) Context(ctx context.Context) *MattersExportsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersExportsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersExportsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}/exports/{exportId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
		"exportId": c.exportId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.exports.get" call.
// Exactly one of *Export or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Export.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *MattersExportsGetCall) Do(opts ...googleapi.CallOption) (*Export, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Export{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets an Export.",
	//   "flatPath": "v1/matters/{matterId}/exports/{exportId}",
	//   "httpMethod": "GET",
	//   "id": "vault.matters.exports.get",
	//   "parameterOrder": [
	//     "matterId",
	//     "exportId"
	//   ],
	//   "parameters": {
	//     "exportId": {
	//       "description": "The export ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}/exports/{exportId}",
	//   "response": {
	//     "$ref": "Export"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery",
	//     "https://www.googleapis.com/auth/ediscovery.readonly"
	//   ]
	// }

}

// method id "vault.matters.exports.list":

type MattersExportsListCall struct {
	s            *Service
	matterId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists Exports.
func (r *MattersExportsService) List(matterId string) *MattersExportsListCall {
	c := &MattersExportsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	return c
}

// PageSize sets the optional parameter "pageSize": The number of
// exports to return in the response.
func (c *MattersExportsListCall) PageSize(pageSize int64) *MattersExportsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The pagination
// token as returned in the response.
func (c *MattersExportsListCall) PageToken(pageToken string) *MattersExportsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersExportsListCall) Fields(s ...googleapi.Field) *MattersExportsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *MattersExportsListCall) IfNoneMatch(entityTag string) *MattersExportsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersExportsListCall) Context(ctx context.Context) *MattersExportsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersExportsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersExportsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}/exports")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.exports.list" call.
// Exactly one of *ListExportsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListExportsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MattersExportsListCall) Do(opts ...googleapi.CallOption) (*ListExportsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListExportsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists Exports.",
	//   "flatPath": "v1/matters/{matterId}/exports",
	//   "httpMethod": "GET",
	//   "id": "vault.matters.exports.list",
	//   "parameterOrder": [
	//     "matterId"
	//   ],
	//   "parameters": {
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The number of exports to return in the response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The pagination token as returned in the response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}/exports",
	//   "response": {
	//     "$ref": "ListExportsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery",
	//     "https://www.googleapis.com/auth/ediscovery.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *MattersExportsListCall) Pages(ctx context.Context, f func(*ListExportsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "vault.matters.holds.addHeldAccounts":

type MattersHoldsAddHeldAccountsCall struct {
	s                      *Service
	matterId               string
	holdId                 string
	addheldaccountsrequest *AddHeldAccountsRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// AddHeldAccounts: Adds HeldAccounts to a hold. Returns a list of
// accounts that have been
// successfully added. Accounts can only be added to an existing
// account-based
// hold.
func (r *MattersHoldsService) AddHeldAccounts(matterId string, holdId string, addheldaccountsrequest *AddHeldAccountsRequest) *MattersHoldsAddHeldAccountsCall {
	c := &MattersHoldsAddHeldAccountsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.holdId = holdId
	c.addheldaccountsrequest = addheldaccountsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersHoldsAddHeldAccountsCall) Fields(s ...googleapi.Field) *MattersHoldsAddHeldAccountsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersHoldsAddHeldAccountsCall) Context(ctx context.Context) *MattersHoldsAddHeldAccountsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersHoldsAddHeldAccountsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersHoldsAddHeldAccountsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.addheldaccountsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}/holds/{holdId}:addHeldAccounts")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
		"holdId":   c.holdId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.holds.addHeldAccounts" call.
// Exactly one of *AddHeldAccountsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AddHeldAccountsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MattersHoldsAddHeldAccountsCall) Do(opts ...googleapi.CallOption) (*AddHeldAccountsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AddHeldAccountsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Adds HeldAccounts to a hold. Returns a list of accounts that have been\nsuccessfully added. Accounts can only be added to an existing account-based\nhold.",
	//   "flatPath": "v1/matters/{matterId}/holds/{holdId}:addHeldAccounts",
	//   "httpMethod": "POST",
	//   "id": "vault.matters.holds.addHeldAccounts",
	//   "parameterOrder": [
	//     "matterId",
	//     "holdId"
	//   ],
	//   "parameters": {
	//     "holdId": {
	//       "description": "The hold ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}/holds/{holdId}:addHeldAccounts",
	//   "request": {
	//     "$ref": "AddHeldAccountsRequest"
	//   },
	//   "response": {
	//     "$ref": "AddHeldAccountsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.holds.create":

type MattersHoldsCreateCall struct {
	s          *Service
	matterId   string
	hold       *Hold
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a hold in the given matter.
func (r *MattersHoldsService) Create(matterId string, hold *Hold) *MattersHoldsCreateCall {
	c := &MattersHoldsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.hold = hold
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersHoldsCreateCall) Fields(s ...googleapi.Field) *MattersHoldsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersHoldsCreateCall) Context(ctx context.Context) *MattersHoldsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersHoldsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersHoldsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.hold)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}/holds")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.holds.create" call.
// Exactly one of *Hold or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Hold.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *MattersHoldsCreateCall) Do(opts ...googleapi.CallOption) (*Hold, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Hold{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a hold in the given matter.",
	//   "flatPath": "v1/matters/{matterId}/holds",
	//   "httpMethod": "POST",
	//   "id": "vault.matters.holds.create",
	//   "parameterOrder": [
	//     "matterId"
	//   ],
	//   "parameters": {
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}/holds",
	//   "request": {
	//     "$ref": "Hold"
	//   },
	//   "response": {
	//     "$ref": "Hold"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.holds.delete":

type MattersHoldsDeleteCall struct {
	s          *Service
	matterId   string
	holdId     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Removes a hold by ID. This will release any HeldAccounts on
// this Hold.
func (r *MattersHoldsService) Delete(matterId string, holdId string) *MattersHoldsDeleteCall {
	c := &MattersHoldsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.holdId = holdId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersHoldsDeleteCall) Fields(s ...googleapi.Field) *MattersHoldsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersHoldsDeleteCall) Context(ctx context.Context) *MattersHoldsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersHoldsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersHoldsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}/holds/{holdId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
		"holdId":   c.holdId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.holds.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *MattersHoldsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Removes a hold by ID. This will release any HeldAccounts on this Hold.",
	//   "flatPath": "v1/matters/{matterId}/holds/{holdId}",
	//   "httpMethod": "DELETE",
	//   "id": "vault.matters.holds.delete",
	//   "parameterOrder": [
	//     "matterId",
	//     "holdId"
	//   ],
	//   "parameters": {
	//     "holdId": {
	//       "description": "The hold ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}/holds/{holdId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.holds.get":

type MattersHoldsGetCall struct {
	s            *Service
	matterId     string
	holdId       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a hold by ID.
func (r *MattersHoldsService) Get(matterId string, holdId string) *MattersHoldsGetCall {
	c := &MattersHoldsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.holdId = holdId
	return c
}

// View sets the optional parameter "view": Specifies which parts of the
// Hold to return.
//
// Possible values:
//   "HOLD_VIEW_UNSPECIFIED"
//   "BASIC_HOLD"
//   "FULL_HOLD"
func (c *MattersHoldsGetCall) View(view string) *MattersHoldsGetCall {
	c.urlParams_.Set("view", view)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersHoldsGetCall) Fields(s ...googleapi.Field) *MattersHoldsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *MattersHoldsGetCall) IfNoneMatch(entityTag string) *MattersHoldsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersHoldsGetCall) Context(ctx context.Context) *MattersHoldsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersHoldsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersHoldsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}/holds/{holdId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
		"holdId":   c.holdId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.holds.get" call.
// Exactly one of *Hold or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Hold.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *MattersHoldsGetCall) Do(opts ...googleapi.CallOption) (*Hold, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Hold{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a hold by ID.",
	//   "flatPath": "v1/matters/{matterId}/holds/{holdId}",
	//   "httpMethod": "GET",
	//   "id": "vault.matters.holds.get",
	//   "parameterOrder": [
	//     "matterId",
	//     "holdId"
	//   ],
	//   "parameters": {
	//     "holdId": {
	//       "description": "The hold ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Specifies which parts of the Hold to return.",
	//       "enum": [
	//         "HOLD_VIEW_UNSPECIFIED",
	//         "BASIC_HOLD",
	//         "FULL_HOLD"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}/holds/{holdId}",
	//   "response": {
	//     "$ref": "Hold"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery",
	//     "https://www.googleapis.com/auth/ediscovery.readonly"
	//   ]
	// }

}

// method id "vault.matters.holds.list":

type MattersHoldsListCall struct {
	s            *Service
	matterId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists holds within a matter. An empty page token in
// ListHoldsResponse
// denotes no more holds to list.
func (r *MattersHoldsService) List(matterId string) *MattersHoldsListCall {
	c := &MattersHoldsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	return c
}

// PageSize sets the optional parameter "pageSize": The number of holds
// to return in the response, between 0 and 100 inclusive.
// Leaving this empty, or as 0, is the same as page_size = 100.
func (c *MattersHoldsListCall) PageSize(pageSize int64) *MattersHoldsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The pagination
// token as returned in the response.
// An empty token means start from the beginning.
func (c *MattersHoldsListCall) PageToken(pageToken string) *MattersHoldsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// View sets the optional parameter "view": Specifies which parts of the
// Hold to return.
//
// Possible values:
//   "HOLD_VIEW_UNSPECIFIED"
//   "BASIC_HOLD"
//   "FULL_HOLD"
func (c *MattersHoldsListCall) View(view string) *MattersHoldsListCall {
	c.urlParams_.Set("view", view)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersHoldsListCall) Fields(s ...googleapi.Field) *MattersHoldsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *MattersHoldsListCall) IfNoneMatch(entityTag string) *MattersHoldsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersHoldsListCall) Context(ctx context.Context) *MattersHoldsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersHoldsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersHoldsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}/holds")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.holds.list" call.
// Exactly one of *ListHoldsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListHoldsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MattersHoldsListCall) Do(opts ...googleapi.CallOption) (*ListHoldsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListHoldsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists holds within a matter. An empty page token in ListHoldsResponse\ndenotes no more holds to list.",
	//   "flatPath": "v1/matters/{matterId}/holds",
	//   "httpMethod": "GET",
	//   "id": "vault.matters.holds.list",
	//   "parameterOrder": [
	//     "matterId"
	//   ],
	//   "parameters": {
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The number of holds to return in the response, between 0 and 100 inclusive.\nLeaving this empty, or as 0, is the same as page_size = 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The pagination token as returned in the response.\nAn empty token means start from the beginning.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Specifies which parts of the Hold to return.",
	//       "enum": [
	//         "HOLD_VIEW_UNSPECIFIED",
	//         "BASIC_HOLD",
	//         "FULL_HOLD"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}/holds",
	//   "response": {
	//     "$ref": "ListHoldsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery",
	//     "https://www.googleapis.com/auth/ediscovery.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *MattersHoldsListCall) Pages(ctx context.Context, f func(*ListHoldsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "vault.matters.holds.removeHeldAccounts":

type MattersHoldsRemoveHeldAccountsCall struct {
	s                         *Service
	matterId                  string
	holdId                    string
	removeheldaccountsrequest *RemoveHeldAccountsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// RemoveHeldAccounts: Removes HeldAccounts from a hold. Returns a list
// of statuses in the same
// order as the request. If this request leaves the hold with no
// held
// accounts, the hold will not apply to any accounts.
func (r *MattersHoldsService) RemoveHeldAccounts(matterId string, holdId string, removeheldaccountsrequest *RemoveHeldAccountsRequest) *MattersHoldsRemoveHeldAccountsCall {
	c := &MattersHoldsRemoveHeldAccountsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.holdId = holdId
	c.removeheldaccountsrequest = removeheldaccountsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersHoldsRemoveHeldAccountsCall) Fields(s ...googleapi.Field) *MattersHoldsRemoveHeldAccountsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersHoldsRemoveHeldAccountsCall) Context(ctx context.Context) *MattersHoldsRemoveHeldAccountsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersHoldsRemoveHeldAccountsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersHoldsRemoveHeldAccountsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.removeheldaccountsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}/holds/{holdId}:removeHeldAccounts")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
		"holdId":   c.holdId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.holds.removeHeldAccounts" call.
// Exactly one of *RemoveHeldAccountsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *RemoveHeldAccountsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MattersHoldsRemoveHeldAccountsCall) Do(opts ...googleapi.CallOption) (*RemoveHeldAccountsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &RemoveHeldAccountsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Removes HeldAccounts from a hold. Returns a list of statuses in the same\norder as the request. If this request leaves the hold with no held\naccounts, the hold will not apply to any accounts.",
	//   "flatPath": "v1/matters/{matterId}/holds/{holdId}:removeHeldAccounts",
	//   "httpMethod": "POST",
	//   "id": "vault.matters.holds.removeHeldAccounts",
	//   "parameterOrder": [
	//     "matterId",
	//     "holdId"
	//   ],
	//   "parameters": {
	//     "holdId": {
	//       "description": "The hold ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}/holds/{holdId}:removeHeldAccounts",
	//   "request": {
	//     "$ref": "RemoveHeldAccountsRequest"
	//   },
	//   "response": {
	//     "$ref": "RemoveHeldAccountsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.holds.update":

type MattersHoldsUpdateCall struct {
	s          *Service
	matterId   string
	holdId     string
	hold       *Hold
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Updates the OU and/or query parameters of a hold. You cannot
// add accounts
// to a hold that covers an OU, nor can you add OUs to a hold that
// covers
// individual accounts. Accounts listed in the hold will be ignored.
func (r *MattersHoldsService) Update(matterId string, holdId string, hold *Hold) *MattersHoldsUpdateCall {
	c := &MattersHoldsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.holdId = holdId
	c.hold = hold
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersHoldsUpdateCall) Fields(s ...googleapi.Field) *MattersHoldsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersHoldsUpdateCall) Context(ctx context.Context) *MattersHoldsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersHoldsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersHoldsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.hold)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}/holds/{holdId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PUT", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
		"holdId":   c.holdId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.holds.update" call.
// Exactly one of *Hold or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Hold.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *MattersHoldsUpdateCall) Do(opts ...googleapi.CallOption) (*Hold, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Hold{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the OU and/or query parameters of a hold. You cannot add accounts\nto a hold that covers an OU, nor can you add OUs to a hold that covers\nindividual accounts. Accounts listed in the hold will be ignored.",
	//   "flatPath": "v1/matters/{matterId}/holds/{holdId}",
	//   "httpMethod": "PUT",
	//   "id": "vault.matters.holds.update",
	//   "parameterOrder": [
	//     "matterId",
	//     "holdId"
	//   ],
	//   "parameters": {
	//     "holdId": {
	//       "description": "The ID of the hold.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}/holds/{holdId}",
	//   "request": {
	//     "$ref": "Hold"
	//   },
	//   "response": {
	//     "$ref": "Hold"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.holds.accounts.create":

type MattersHoldsAccountsCreateCall struct {
	s           *Service
	matterId    string
	holdId      string
	heldaccount *HeldAccount
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Create: Adds a HeldAccount to a hold. Accounts can only be added to a
// hold that
// has no held_org_unit set. Attempting to add an account to an
// OU-based
// hold will result in an error.
func (r *MattersHoldsAccountsService) Create(matterId string, holdId string, heldaccount *HeldAccount) *MattersHoldsAccountsCreateCall {
	c := &MattersHoldsAccountsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.holdId = holdId
	c.heldaccount = heldaccount
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersHoldsAccountsCreateCall) Fields(s ...googleapi.Field) *MattersHoldsAccountsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersHoldsAccountsCreateCall) Context(ctx context.Context) *MattersHoldsAccountsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersHoldsAccountsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersHoldsAccountsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.heldaccount)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}/holds/{holdId}/accounts")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
		"holdId":   c.holdId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.holds.accounts.create" call.
// Exactly one of *HeldAccount or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *HeldAccount.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *MattersHoldsAccountsCreateCall) Do(opts ...googleapi.CallOption) (*HeldAccount, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &HeldAccount{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Adds a HeldAccount to a hold. Accounts can only be added to a hold that\nhas no held_org_unit set. Attempting to add an account to an OU-based\nhold will result in an error.",
	//   "flatPath": "v1/matters/{matterId}/holds/{holdId}/accounts",
	//   "httpMethod": "POST",
	//   "id": "vault.matters.holds.accounts.create",
	//   "parameterOrder": [
	//     "matterId",
	//     "holdId"
	//   ],
	//   "parameters": {
	//     "holdId": {
	//       "description": "The hold ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}/holds/{holdId}/accounts",
	//   "request": {
	//     "$ref": "HeldAccount"
	//   },
	//   "response": {
	//     "$ref": "HeldAccount"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.holds.accounts.delete":

type MattersHoldsAccountsDeleteCall struct {
	s          *Service
	matterId   string
	holdId     string
	accountId  string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Removes a HeldAccount from a hold. If this request leaves the
// hold with
// no held accounts, the hold will not apply to any accounts.
func (r *MattersHoldsAccountsService) Delete(matterId string, holdId string, accountId string) *MattersHoldsAccountsDeleteCall {
	c := &MattersHoldsAccountsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.holdId = holdId
	c.accountId = accountId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersHoldsAccountsDeleteCall) Fields(s ...googleapi.Field) *MattersHoldsAccountsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersHoldsAccountsDeleteCall) Context(ctx context.Context) *MattersHoldsAccountsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersHoldsAccountsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersHoldsAccountsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}/holds/{holdId}/accounts/{accountId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId":  c.matterId,
		"holdId":    c.holdId,
		"accountId": c.accountId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.holds.accounts.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *MattersHoldsAccountsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Removes a HeldAccount from a hold. If this request leaves the hold with\nno held accounts, the hold will not apply to any accounts.",
	//   "flatPath": "v1/matters/{matterId}/holds/{holdId}/accounts/{accountId}",
	//   "httpMethod": "DELETE",
	//   "id": "vault.matters.holds.accounts.delete",
	//   "parameterOrder": [
	//     "matterId",
	//     "holdId",
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The ID of the account to remove from the hold.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "holdId": {
	//       "description": "The hold ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}/holds/{holdId}/accounts/{accountId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery"
	//   ]
	// }

}

// method id "vault.matters.holds.accounts.list":

type MattersHoldsAccountsListCall struct {
	s            *Service
	matterId     string
	holdId       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists HeldAccounts for a hold. This will only list individually
// specified
// held accounts. If the hold is on an OU, then use
// <a href="https://developers.google.com/admin-sdk/">Admin SDK</a>
// to enumerate its members.
func (r *MattersHoldsAccountsService) List(matterId string, holdId string) *MattersHoldsAccountsListCall {
	c := &MattersHoldsAccountsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.matterId = matterId
	c.holdId = holdId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MattersHoldsAccountsListCall) Fields(s ...googleapi.Field) *MattersHoldsAccountsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *MattersHoldsAccountsListCall) IfNoneMatch(entityTag string) *MattersHoldsAccountsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MattersHoldsAccountsListCall) Context(ctx context.Context) *MattersHoldsAccountsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MattersHoldsAccountsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MattersHoldsAccountsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/matters/{matterId}/holds/{holdId}/accounts")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"matterId": c.matterId,
		"holdId":   c.holdId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "vault.matters.holds.accounts.list" call.
// Exactly one of *ListHeldAccountsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListHeldAccountsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MattersHoldsAccountsListCall) Do(opts ...googleapi.CallOption) (*ListHeldAccountsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListHeldAccountsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists HeldAccounts for a hold. This will only list individually specified\nheld accounts. If the hold is on an OU, then use\n\u003ca href=\"https://developers.google.com/admin-sdk/\"\u003eAdmin SDK\u003c/a\u003e\nto enumerate its members.",
	//   "flatPath": "v1/matters/{matterId}/holds/{holdId}/accounts",
	//   "httpMethod": "GET",
	//   "id": "vault.matters.holds.accounts.list",
	//   "parameterOrder": [
	//     "matterId",
	//     "holdId"
	//   ],
	//   "parameters": {
	//     "holdId": {
	//       "description": "The hold ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "matterId": {
	//       "description": "The matter ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/matters/{matterId}/holds/{holdId}/accounts",
	//   "response": {
	//     "$ref": "ListHeldAccountsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/ediscovery",
	//     "https://www.googleapis.com/auth/ediscovery.readonly"
	//   ]
	// }

}