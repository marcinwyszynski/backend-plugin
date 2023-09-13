package backendplugin

import (
	"context"

	"github.com/marcinwyszynski/backendplugin/proto"
)

type GRPCServer struct {
	proto.UnimplementedBackendServer

	// This is the real implementation
	Impl BackendPlugin
}

func (s *GRPCServer) ConfigureBackend(ctx context.Context, req *proto.ConfigureBackend_Request) (*proto.Empty, error) {
	if err := s.Impl.Configure(ctx, req.Config); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (s *GRPCServer) DeleteWorkspace(ctx context.Context, req *proto.DeleteWorkspace_Request) (*proto.Empty, error) {
	if err := s.Impl.DeleteWorkspace(ctx, req.Workspace, req.Force); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (s *GRPCServer) ListWorkspaces(ctx context.Context, req *proto.ListWorkspaces_Request) (*proto.ListWorkspaces_Response, error) {
	workspaces, err := s.Impl.ListWorkspaces(ctx)
	if err != nil {
		return nil, err
	}

	return &proto.ListWorkspaces_Response{
		Workspaces: workspaces,
	}, nil
}

func (s *GRPCServer) GetStatePayload(ctx context.Context, req *proto.GetStatePayload_Request) (*proto.GetStatePayload_Response, error) {
	payload, err := s.Impl.GetStatePayload(ctx, req.Workspace)
	if err != nil {
		return nil, err
	}

	var statePayload *proto.StatePayload
	if payload != nil {
		statePayload = &proto.StatePayload{
			Md5:  payload.MD5,
			Data: payload.Data,
		}
	}

	return &proto.GetStatePayload_Response{
		Payload: statePayload,
	}, nil
}

func (s *GRPCServer) PutState(ctx context.Context, req *proto.PutState_Request) (*proto.Empty, error) {
	if err := s.Impl.PutState(ctx, req.Workspace, req.Data); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (s *GRPCServer) DeleteState(ctx context.Context, req *proto.DeleteState_Request) (*proto.Empty, error) {
	if err := s.Impl.DeleteState(ctx, req.Workspace); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (s *GRPCServer) LockState(ctx context.Context, req *proto.StateLock_Request) (*proto.StateLock_Response, error) {
	resp, err := s.Impl.LockState(ctx, req.Workspace, &LockInfo{
		ID:        req.Info.Id,
		Operation: req.Info.Operation,
		Info:      req.Info.Info,
		Who:       req.Info.Who,
		Version:   req.Info.Version,
		Created:   req.Info.Created.AsTime(),
		Path:      req.Info.Path,
	})

	if err != nil {
		return nil, err
	}

	return &proto.StateLock_Response{Id: resp}, nil
}

func (s *GRPCServer) UnlockState(ctx context.Context, req *proto.StateUnlock_Request) (*proto.Empty, error) {
	if err := s.Impl.UnlockState(ctx, req.Workspace, req.Id); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}
