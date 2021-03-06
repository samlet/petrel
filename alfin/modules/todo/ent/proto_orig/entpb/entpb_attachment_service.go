// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	ent "github.com/samlet/petrel/alfin/modules/todo/ent"
	runtime "entgo.io/contrib/entproto/runtime"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// AttachmentService implements AttachmentServiceServer
type AttachmentService struct {
	client *ent.Client
	UnimplementedAttachmentServiceServer
}

// NewAttachmentService returns a new AttachmentService
func NewAttachmentService(client *ent.Client) *AttachmentService {
	return &AttachmentService{
		client: client,
	}
}

// toProtoAttachment transforms the ent type to the pb type
func toProtoAttachment(e *ent.Attachment) *Attachment {
	return &Attachment{
		Id: runtime.MustExtractUUIDBytes(e.ID),
	}
}

// validateAttachment validates that all fields are encoded properly and are safe to pass
// to the ent entity builder.
func validateAttachment(x *Attachment, checkId bool) error {
	if err := runtime.ValidateUUID(x.GetId()); err != nil && checkId {
		return err
	}
	return nil
}

// Create implements AttachmentServiceServer.Create
func (svc *AttachmentService) Create(ctx context.Context, req *CreateAttachmentRequest) (*Attachment, error) {
	attachment := req.GetAttachment()
	if err := validateAttachment(attachment, true); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	res, err := svc.client.Attachment.Create().
		SetUserID(int(attachment.GetUser().GetId())).
		Save(ctx)

	switch {
	case err == nil:
		return toProtoAttachment(res), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal: %s", err)
	}
}

// Get implements AttachmentServiceServer.Get
func (svc *AttachmentService) Get(ctx context.Context, req *GetAttachmentRequest) (*Attachment, error) {
	if err := runtime.ValidateUUID(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	get, err := svc.client.Attachment.Get(ctx, runtime.MustBytesToUUID(req.GetId()))
	switch {
	case err == nil:
		return toProtoAttachment(get), nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
}

// Update implements AttachmentServiceServer.Update
func (svc *AttachmentService) Update(ctx context.Context, req *UpdateAttachmentRequest) (*Attachment, error) {
	attachment := req.GetAttachment()
	if err := validateAttachment(attachment, false); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	res, err := svc.client.Attachment.UpdateOneID(runtime.MustBytesToUUID(attachment.GetId())).
		SetUserID(int(attachment.GetUser().GetId())).
		Save(ctx)

	switch {
	case err == nil:
		return toProtoAttachment(res), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal: %s", err)
	}
}

// Delete implements AttachmentServiceServer.Delete
func (svc *AttachmentService) Delete(ctx context.Context, req *DeleteAttachmentRequest) (*emptypb.Empty, error) {
	if err := runtime.ValidateUUID(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	err := svc.client.Attachment.DeleteOneID(runtime.MustBytesToUUID(req.GetId())).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
}
