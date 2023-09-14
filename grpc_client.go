package backendplugin

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/marcinwyszynski/backendplugin/proto"
)

type GRPCClient struct {
	client proto.BackendClient
}

func (c *GRPCClient) Configure(ctx context.Context, config map[string]string) error {
	_, err := c.client.ConfigureBackend(ctx, &proto.ConfigureBackend_Request{
		Config: config,
	})

	if err != nil {
		return fmt.Errorf("failed to configure backend: %w", err)
	}

	return nil
}

func (c *GRPCClient) ListWorkspaces(ctx context.Context) ([]string, error) {
	response, err := c.client.ListWorkspaces(ctx, &proto.ListWorkspaces_Request{})
	if err != nil {
		return nil, fmt.Errorf("failed to list workspaces: %w", err)
	}

	return response.Workspaces, nil
}

func (c *GRPCClient) DeleteWorkspace(ctx context.Context, workspace string, force bool) error {
	_, err := c.client.DeleteWorkspace(ctx, &proto.DeleteWorkspace_Request{
		Workspace: workspace,
		Force:     force,
	})

	if err != nil {
		return fmt.Errorf("failed to delete workspace: %w", err)
	}

	return nil
}

func (c *GRPCClient) GetStatePayload(ctx context.Context, workspace string) (*Payload, error) {
	response, err := c.client.GetStatePayload(ctx, &proto.GetStatePayload_Request{
		Workspace: workspace,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get state: %w", err)
	} else if response.Payload == nil {
		return nil, nil
	}

	return &Payload{
		Data: response.Payload.Data,
		MD5:  response.Payload.Md5,
	}, nil
}

func (c *GRPCClient) PutState(ctx context.Context, workspace string, data []byte) error {
	_, err := c.client.PutState(ctx, &proto.PutState_Request{
		Workspace: workspace,
		Data:      data,
	})

	if err != nil {
		return fmt.Errorf("failed to put state: %w", err)
	}

	return nil
}

func (c *GRPCClient) DeleteState(ctx context.Context, workspace string) error {
	_, err := c.client.DeleteState(ctx, &proto.DeleteState_Request{
		Workspace: workspace,
	})

	if err != nil {
		return fmt.Errorf("failed to delete state: %w", err)
	}

	return nil
}

func (c *GRPCClient) LockState(ctx context.Context, workspace string, info *LockInfo) (string, error) {
	response, err := c.client.LockState(ctx, &proto.StateLock_Request{
		Workspace: workspace,
		Info: &proto.StateLockInfo{
			Id:        info.ID,
			Operation: info.Operation,
			Info:      info.Info,
			Who:       info.Who,
			Version:   info.Version,
			Created:   timestamppb.New(info.Created),
			Path:      info.Path,
		},
	})

	if err != nil {
		return "", fmt.Errorf("failed to lock state: %w", err)
	}

	return response.Id, nil
}

func (c *GRPCClient) UnlockState(ctx context.Context, workspace, id string) error {
	_, err := c.client.UnlockState(ctx, &proto.StateUnlock_Request{
		Workspace: workspace,
		Id:        id,
	})

	if err != nil {
		return fmt.Errorf("failed to unlock state: %w", err)
	}

	return nil
}
