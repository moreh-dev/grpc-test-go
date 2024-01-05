package sdamanager_api

import (
	"context"
	"fmt"

	proto "moreh.io/go-grpc-test/proto"
)

type SDAManagerAPI struct {
}

func NewSDAManagerAPI() *SDAManagerAPI {
	return &SDAManagerAPI{}
}
func (a *SDAManagerAPI) ListSDAModel(ctx context.Context, in *proto.Token) (*proto.SDAModelList, error) {
	fmt.Println("ListSDAModel called")
	return &proto.SDAModelList{
		Modellist: []*proto.SDAModel{
			{
				Id:   1,
				Name: "dummy",
			},
		},
	}, nil
}
func (a *SDAManagerAPI) CreateSDAModel(ctx context.Context, in *proto.CreateSDAModelRequest) (*proto.CreateSDAModelResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) DeleteSDAModel(ctx context.Context, in *proto.DeleteSDAModelRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) CreateSDA(ctx context.Context, in *proto.CreateSDARequest) (*proto.CreateSDAResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) GetSDAInfo(ctx context.Context, in *proto.Token) (*proto.GetSDAInfoResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ListSDA(ctx context.Context, in *proto.Empty) (*proto.SDAList, error) {
	return nil, nil
}
func (a *SDAManagerAPI) UpdateSDA(ctx context.Context, in *proto.UpdateSDARequest) (*proto.UpdateResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) DeleteSDA(ctx context.Context, in *proto.DeleteSDARequest) (*proto.DeleteResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) CreateToken(ctx context.Context, in *proto.CreateTokenRequest) (*proto.CreateTokenResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ListToken(ctx context.Context, in *proto.Empty) (*proto.TokenList, error) {
	return nil, nil
}
func (a *SDAManagerAPI) GetToken(ctx context.Context, in *proto.Token) (*proto.GetTokenResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) UpdateToken(ctx context.Context, in *proto.UpdateTokenRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) DeleteToken(ctx context.Context, in *proto.DeleteTokenRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ListBackend(ctx context.Context, in *proto.BackendListRequest) (*proto.BackendList, error) {
	return nil, nil
}
func (a *SDAManagerAPI) GetBackend(ctx context.Context, in *proto.GetBackendRequest) (*proto.GetBackendResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ReleaseDevice(ctx context.Context, in *proto.ReleaseDeviceRequest) (*proto.ReleaseDeviceResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ReleaseAllDevice(ctx context.Context, in *proto.Empty) (*proto.Boolean, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ListSDAUtilizations(ctx context.Context, in *proto.ListSDAUtilizationsRequest) (*proto.ListSDAUtilizationsResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ListSDAJob(ctx context.Context, in *proto.GetSDAUtilizationsRequest) (*proto.GetSDAUtilizationsResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) GetSDAUtilizations(ctx context.Context, in *proto.GetSDAUtilizationsRequest) (*proto.GetSDAUtilizationsResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ListJobQueue(ctx context.Context, in *proto.JobQueueRequest) (*proto.JobQueueList, error) {
	return nil, nil
}
func (a *SDAManagerAPI) UpdateJobPriority(ctx context.Context, in *proto.UpdateJobPriorityRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ListJobHistory(ctx context.Context, in *proto.JobHistoryRequest) (*proto.JobHistoryList, error) {
	return nil, nil
}
func (a *SDAManagerAPI) GetVersion(ctx context.Context, in *proto.Empty) (*proto.Version, error) {
	return nil, nil
}
func (a *SDAManagerAPI) GetServerVersion(ctx context.Context, in *proto.Empty) (*proto.ListServerVersion, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ListEventLog(ctx context.Context, in *proto.EventLogRequest) (*proto.EventLogList, error) {
	return nil, nil
}
func (a *SDAManagerAPI) CreateBackend(ctx context.Context, in *proto.CreateBackendRequest) (*proto.CreateBackendResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) DeleteBackend(ctx context.Context, in *proto.DeleteBackendRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) UpdateBackend(ctx context.Context, in *proto.UpdateBackendRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) UpdateDevice(ctx context.Context, in *proto.UpdateDeviceRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) CreateVersionInfo(ctx context.Context, in *proto.CreateVersionInfoRequest) (*proto.CreateVersionInfoResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) DeleteVersionInfo(ctx context.Context, in *proto.DeleteVersionInfoRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) UpdateVersionInfo(ctx context.Context, in *proto.UpdateVersionInfoRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ListVersionInfo(ctx context.Context, in *proto.Version) (*proto.VersionInfoList, error) {
	return nil, nil
}
func (a *SDAManagerAPI) CreateGroup(ctx context.Context, in *proto.CreateGroupRequest) (*proto.CreateGroupResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) DeleteGroup(ctx context.Context, in *proto.DeleteGroupRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) UpdateGroup(ctx context.Context, in *proto.UpdateGroupRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ListGroup(ctx context.Context, in *proto.ListGroupRequest) (*proto.GroupList, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ChangeGroup(ctx context.Context, in *proto.ChangeGroupRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ListRelation(ctx context.Context, in *proto.ListRelationRequest) (*proto.RelationList, error) {
	return nil, nil
}
func (a *SDAManagerAPI) GetDeviceUsage(ctx context.Context, in *proto.DeviceUsageRequest) (*proto.DeviceUsageResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ListTokenUsage(ctx context.Context, in *proto.ListTokenUsageRequest) (*proto.ListTokenUsageResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) ListConfigure(ctx context.Context, in *proto.ListConfigureRequest) (*proto.ListConfigureResponse, error) {
	return nil, nil
}
func (a *SDAManagerAPI) GetConfigure(ctx context.Context, in *proto.GetConfigureRequest) (*proto.Configure, error) {
	return nil, nil
}
func (a *SDAManagerAPI) UpdateConfigure(ctx context.Context, in *proto.UpdateConfigureRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}
