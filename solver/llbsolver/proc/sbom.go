package proc

import (
	"context"

	"github.com/pkg/errors"
	"github.com/preminger/buildkit/client/llb"
	"github.com/preminger/buildkit/exporter/containerimage/exptypes"
	"github.com/preminger/buildkit/frontend"
	"github.com/preminger/buildkit/frontend/attestations/sbom"
	"github.com/preminger/buildkit/solver"
	"github.com/preminger/buildkit/solver/llbsolver"
	"github.com/preminger/buildkit/solver/result"
)

func SBOMProcessor(scannerRef string, useCache bool) llbsolver.Processor {
	return func(ctx context.Context, res *llbsolver.Result, s *llbsolver.Solver, j *solver.Job) (*llbsolver.Result, error) {
		// skip sbom generation if we already have an sbom
		if sbom.HasSBOM(res.Result) {
			return res, nil
		}

		ps, err := exptypes.ParsePlatforms(res.Metadata)
		if err != nil {
			return nil, err
		}

		scanner, err := sbom.CreateSBOMScanner(ctx, s.Bridge(j), scannerRef)
		if err != nil {
			return nil, err
		}
		if scanner == nil {
			return res, nil
		}

		for _, p := range ps.Platforms {
			ref, ok := res.FindRef(p.ID)
			if !ok {
				return nil, errors.Errorf("could not find ref %s", p.ID)
			}
			if ref == nil {
				continue
			}

			defop, err := llb.NewDefinitionOp(ref.Definition())
			if err != nil {
				return nil, err
			}
			st := llb.NewState(defop)

			var opts []llb.ConstraintsOpt
			if !useCache {
				opts = append(opts, llb.IgnoreCache)
			}
			att, err := scanner(ctx, p.ID, st, nil, opts...)
			if err != nil {
				return nil, err
			}
			attSolve, err := result.ConvertAttestation(&att, func(st llb.State) (solver.ResultProxy, error) {
				def, err := st.Marshal(ctx)
				if err != nil {
					return nil, err
				}

				r, err := s.Bridge(j).Solve(ctx, frontend.SolveRequest{ // TODO: buildinfo
					Definition: def.ToPB(),
				}, j.SessionID)
				if err != nil {
					return nil, err
				}
				return r.Ref, nil
			})
			if err != nil {
				return nil, err
			}
			res.AddAttestation(p.ID, *attSolve)
		}
		return res, nil
	}
}
