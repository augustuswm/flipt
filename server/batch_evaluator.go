package server

import (
	"context"
	// "fmt"
	// "hash/crc32"
	// "sort"
	// "strconv"
	// "strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes"
	// "github.com/markphelps/flipt/errors"
	flipt "github.com/markphelps/flipt/rpc"
	// "github.com/markphelps/flipt/storage"
)

// Evaluate evaluates a request for a batch of flags and a single entity
func (s *Server) BatchEvaluate(ctx context.Context, r *flipt.BatchEvaluationRequest) (*flipt.BatchEvaluationResponse, error) {
	s.logger.WithField("request", r).Debug("evaluate")
	startTime := time.Now()

	// set request ID if not present
	if r.RequestId == "" {
		r.RequestId = uuid.Must(uuid.NewV4()).String()
	}

	resp, err := s.batchEvaluate(ctx, r)
	if resp != nil {
		resp.RequestDurationMillis = float64(time.Since(startTime)) / float64(time.Millisecond)
	}

	if err != nil {
		return resp, err
	}

	s.logger.WithField("response", resp).Debug("evaluate")
	return resp, nil
}


func (s *Server) batchEvaluate(ctx context.Context, r *flipt.BatchEvaluationRequest) (*flipt.BatchEvaluationResponse, error) {
	var (
		ts, _ = ptypes.TimestampProto(time.Now().UTC())
		resp  = &flipt.BatchEvaluationResponse{
			RequestId:      r.RequestId,
			EntityId:       r.EntityId,
			RequestContext: r.Context,
			Evaluations: 		nil,
			Timestamp:      ts,
		}
	)

	var results []*flipt.FlagEvaluation

	flags, err := s.FlagStore.GetFlags(ctx, r.FlagKeys)

	if err != nil {
		return resp, err
	}

	for _, flag := range flags {
		if flag.Enabled {
			s.logger.WithField("flag", flag).Debug("Eval Flag")
			eval, err := s.evaluateFlag(ctx, r.EntityId, r.Context, flag)

			if err != nil {
				return resp, err
			}

			results = append(results, eval)
		} else {
			s.logger.WithField("key", flag.Key).Debug("Flag not enabled")
		}
	}

	resp.Evaluations = results

	return resp, nil
}