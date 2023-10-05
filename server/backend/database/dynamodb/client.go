package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/yorkie-team/yorkie/api/types"
	"github.com/yorkie-team/yorkie/pkg/document"
	"github.com/yorkie-team/yorkie/pkg/document/change"
	"github.com/yorkie-team/yorkie/pkg/document/key"
	"github.com/yorkie-team/yorkie/pkg/document/time"
	"github.com/yorkie-team/yorkie/server/backend/database"
)

type Client struct {
	client *dynamodb.Client
}

var _ database.Database = (*Client)(nil)

func NewClient() *Client {
	return &Client{
		client: nil,
	}
}

// Close all resources of this database.
func (c *Client) Close() error {
	return nil
}

// FindProjectInfoByPublicKey returns a project by public key.
func (c *Client) FindProjectInfoByPublicKey(
	ctx context.Context,
	publicKey string,
) (*database.ProjectInfo, error) {
	return nil, nil
}

// FindProjectInfoByName returns a project by the given name.
func (c *Client) FindProjectInfoByName(
	ctx context.Context,
	owner types.ID,
	name string,
) (*database.ProjectInfo, error) {
	return nil, nil
}

// FindProjectInfoByID returns a project by the given id. It should not be
// used directly by clients because it is not checked if the project is
// permitted to be accessed by the admin client.
func (c *Client) FindProjectInfoByID(ctx context.Context, id types.ID) (*database.ProjectInfo, error) {
	return nil, nil
}

// EnsureDefaultUserAndProject ensures that the default user and project
// exists.
func (c *Client) EnsureDefaultUserAndProject(
	ctx context.Context,
	username,
	password string,
	clientDeactivateThreshold string,
) (*database.UserInfo, *database.ProjectInfo, error) {
	return nil, nil, nil
}

// CreateProjectInfo creates a new project.
func (c *Client) CreateProjectInfo(
	ctx context.Context,
	name string,
	owner types.ID,
	clientDeactivateThreshold string,
) (*database.ProjectInfo, error) {
	return nil, nil
}

// ListProjectInfos returns all project infos owned by owner.
func (c *Client) ListProjectInfos(ctx context.Context, owner types.ID) ([]*database.ProjectInfo, error) {
	return nil, nil
}

// UpdateProjectInfo updates the project.
func (c *Client) UpdateProjectInfo(
	ctx context.Context,
	owner types.ID,
	id types.ID,
	fields *types.UpdatableProjectFields,
) (*database.ProjectInfo, error) {
	return nil, nil
}

// CreateUserInfo creates a new user.
func (c *Client) CreateUserInfo(
	ctx context.Context,
	username string,
	hashedPassword string,
) (*database.UserInfo, error) {
	return nil, nil
}

// FindUserInfo returns a user by the given username.
func (c *Client) FindUserInfo(ctx context.Context, username string) (*database.UserInfo, error) {
	return nil, nil
}

// ListUserInfos returns all users.
func (c *Client) ListUserInfos(ctx context.Context) ([]*database.UserInfo, error) {
	return nil, nil
}

// ActivateClient activates the client of the given key.
func (c *Client) ActivateClient(ctx context.Context, projectID types.ID, key string) (*database.ClientInfo, error) {
	return nil, nil
}

// DeactivateClient deactivates the client of the given ID.
func (c *Client) DeactivateClient(ctx context.Context, projectID, clientID types.ID) (*database.ClientInfo, error) {
	return nil, nil
}

// FindClientInfoByID finds the client of the given ID.
func (c *Client) FindClientInfoByID(ctx context.Context, projectID, clientID types.ID) (*database.ClientInfo, error) {
	return nil, nil
}

// UpdateClientInfoAfterPushPull updates the client from the given clientInfo
// after handling PushPull.
func (c *Client) UpdateClientInfoAfterPushPull(ctx context.Context, clientInfo *database.ClientInfo, docInfo *database.DocInfo) error {
	return nil
}

// FindDeactivateCandidates finds the housekeeping candidates.
func (c *Client) FindDeactivateCandidates(
	ctx context.Context,
	candidatesLimitPerProject int,
	projectFetchSize int,
	lastProjectID types.ID,
) (types.ID, []*database.ClientInfo, error) {
	return "", nil, nil
}

// FindDocInfoByKey finds the document of the given key.
func (c *Client) FindDocInfoByKey(
	ctx context.Context,
	projectID types.ID,
	docKey key.Key,
) (*database.DocInfo, error) {
	return nil, nil
}

// FindDocInfoByKeyAndOwner finds the document of the given key. If the
// createDocIfNotExist condition is true, create the document if it does not
// exist.
func (c *Client) FindDocInfoByKeyAndOwner(
	ctx context.Context,
	projectID types.ID,
	clientID types.ID,
	docKey key.Key,
	createDocIfNotExist bool,
) (*database.DocInfo, error) {
	return nil, nil
}

// FindDocInfoByID finds the document of the given ID.
func (c *Client) FindDocInfoByID(
	ctx context.Context,
	projectID types.ID,
	id types.ID,
) (*database.DocInfo, error) {
	return nil, nil
}

// UpdateDocInfoStatusToRemoved updates the document status to removed.
func (c *Client) UpdateDocInfoStatusToRemoved(
	ctx context.Context,
	projectID types.ID,
	docID types.ID,
) error {
	return nil
}

// CreateChangeInfos stores the given changes then updates the given docInfo.
func (c *Client) CreateChangeInfos(
	ctx context.Context,
	projectID types.ID,
	docInfo *database.DocInfo,
	initialServerSeq int64,
	changes []*change.Change,
	isRemoved bool,
) error {
	return nil
}

// PurgeStaleChanges delete changes before the smallest in `syncedseqs` to
// save storage.
func (c *Client) PurgeStaleChanges(
	ctx context.Context,
	docID types.ID,
) error {
	return nil
}

// FindChangesBetweenServerSeqs returns the changes between two server sequences.
func (c *Client) FindChangesBetweenServerSeqs(
	ctx context.Context,
	docID types.ID,
	from int64,
	to int64,
) ([]*change.Change, error) {
	return nil, nil
}

// FindChangeInfosBetweenServerSeqs returns the changeInfos between two server sequences.
func (c *Client) FindChangeInfosBetweenServerSeqs(
	ctx context.Context,
	docID types.ID,
	from int64,
	to int64,
) ([]*database.ChangeInfo, error) {
	return nil, nil
}

// CreateSnapshotInfo stores the snapshot of the given document.
func (c *Client) CreateSnapshotInfo(ctx context.Context, docID types.ID, doc *document.InternalDocument) error {
	return nil
}

// FindSnapshotInfoByID returns the snapshot by the given id.
func (c *Client) FindSnapshotInfoByID(ctx context.Context, id types.ID) (*database.SnapshotInfo, error) {
	return nil, nil
}

// FindClosestSnapshotInfo finds the closest snapshot info in a given serverSeq.
func (c *Client) FindClosestSnapshotInfo(
	ctx context.Context,
	docID types.ID,
	serverSeq int64,
	includeSnapshot bool,
) (*database.SnapshotInfo, error) {
	return nil, nil
}

// FindMinSyncedSeqInfo finds the minimum synced sequence info.
func (c *Client) FindMinSyncedSeqInfo(ctx context.Context, docID types.ID) (*database.SyncedSeqInfo, error) {
	return nil, nil
}

// UpdateAndFindMinSyncedTicket updates the given serverSeq of the given client
// and returns the min synced ticket.
func (c *Client) UpdateAndFindMinSyncedTicket(
	ctx context.Context,
	clientInfo *database.ClientInfo,
	docID types.ID,
	serverSeq int64,
) (*time.Ticket, error) {
	return nil, nil
}

// UpdateSyncedSeq updates the syncedSeq of the given client.
func (c *Client) UpdateSyncedSeq(
	ctx context.Context,
	clientInfo *database.ClientInfo,
	docID types.ID,
	serverSeq int64,
) error {
	return nil
}

// FindDocInfosByPaging returns the documentInfos of the given paging.
func (c *Client) FindDocInfosByPaging(
	ctx context.Context,
	projectID types.ID,
	paging types.Paging[types.ID],
) ([]*database.DocInfo, error) {
	return nil, nil
}

// FindDocInfosByQuery returns the documentInfos which match the given query.
func (c *Client) FindDocInfosByQuery(
	ctx context.Context,
	projectID types.ID,
	query string,
	pageSize int,
) (*types.SearchResult[*database.DocInfo], error) {
	return nil, nil
}

// IsDocumentAttached returns true if the document is attached to clients.
func (c *Client) IsDocumentAttached(
	ctx context.Context,
	projectID types.ID,
	docID types.ID,
	excludeClientID types.ID,
) (bool, error) {
	return false, nil
}
