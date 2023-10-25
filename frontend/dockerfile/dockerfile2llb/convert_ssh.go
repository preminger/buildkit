package dockerfile2llb

import (
	"github.com/pkg/errors"
	"github.com/preminger/buildkit/client/llb"
	"github.com/preminger/buildkit/frontend/dockerfile/instructions"
	"github.com/preminger/buildkit/frontend/dockerfile/parser"
)

func dispatchSSH(d *dispatchState, m *instructions.Mount, loc []parser.Range) (llb.RunOption, error) {
	if m.Source != "" {
		return nil, errors.Errorf("ssh does not support source")
	}

	id := m.CacheID
	if id == "" {
		id = "default"
	}
	if _, ok := d.outline.ssh[id]; !ok {
		d.outline.ssh[id] = sshInfo{
			location: loc,
			required: m.Required,
		}
	}

	opts := []llb.SSHOption{llb.SSHID(m.CacheID)}

	if m.Target != "" {
		opts = append(opts, llb.SSHSocketTarget(m.Target))
	}

	if !m.Required {
		opts = append(opts, llb.SSHOptional)
	}

	if m.UID != nil || m.GID != nil || m.Mode != nil {
		var uid, gid, mode int
		if m.UID != nil {
			uid = int(*m.UID)
		}
		if m.GID != nil {
			gid = int(*m.GID)
		}
		if m.Mode != nil {
			mode = int(*m.Mode)
		} else {
			mode = 0600
		}
		opts = append(opts, llb.SSHSocketOpt(m.Target, uid, gid, mode))
	}

	return llb.AddSSHSocket(opts...), nil
}
