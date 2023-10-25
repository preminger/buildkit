package worker

import (
	"context"
	"io"

	"github.com/containerd/containerd/content"
	"github.com/containerd/containerd/leases"
	digest "github.com/opencontainers/go-digest"
	ocispecs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/preminger/buildkit/cache"
	"github.com/preminger/buildkit/client"
	"github.com/preminger/buildkit/client/llb"
	"github.com/preminger/buildkit/executor"
	"github.com/preminger/buildkit/exporter"
	"github.com/preminger/buildkit/frontend"
	"github.com/preminger/buildkit/session"
	"github.com/preminger/buildkit/solver"
)

type Worker interface {
	io.Closer
	// ID needs to be unique in the cluster
	ID() string
	Labels() map[string]string
	Platforms(noCache bool) []ocispecs.Platform
	BuildkitVersion() client.BuildkitVersion

	GCPolicy() []client.PruneInfo
	LoadRef(ctx context.Context, id string, hidden bool) (cache.ImmutableRef, error)
	// ResolveOp resolves Vertex.Sys() to Op implementation.
	ResolveOp(v solver.Vertex, s frontend.FrontendLLBBridge, sm *session.Manager) (solver.Op, error)
	ResolveImageConfig(ctx context.Context, ref string, opt llb.ResolveImageConfigOpt, sm *session.Manager, g session.Group) (digest.Digest, []byte, error)
	DiskUsage(ctx context.Context, opt client.DiskUsageInfo) ([]*client.UsageInfo, error)
	Exporter(name string, sm *session.Manager) (exporter.Exporter, error)
	Prune(ctx context.Context, ch chan client.UsageInfo, opt ...client.PruneInfo) error
	FromRemote(ctx context.Context, remote *solver.Remote) (cache.ImmutableRef, error)
	PruneCacheMounts(ctx context.Context, ids []string) error
	ContentStore() content.Store
	Executor() executor.Executor
	CacheManager() cache.Manager
	LeaseManager() leases.Manager
}

type Infos interface {
	GetDefault() (Worker, error)
	WorkerInfos() []client.WorkerInfo
}
