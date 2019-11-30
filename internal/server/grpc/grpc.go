package grpc

import (
	"context"
	"fmt"
	zlog "github.com/rs/zerolog/log"
	gameenginepb "github.com/sorborail/m-apis/game-enginepb/v1"
	"google.golang.org/grpc"
	"m-game-engine/internal/server/logic"
	"net"
)

type Server struct {
	address string
	srv *grpc.Server
	lis net.Listener
}

func NewServer(addr string) *Server {
	return &Server{address: addr}
}

func (*Server) GetSize(ctx context.Context, req *gameenginepb.GetSizeRequest) (*gameenginepb.GetSizeResponse, error) {
	zlog.Info().Msg("GetSize method in m-game-engine service called")
	return &gameenginepb.GetSizeResponse{Size: logic.GetSize()}, nil
}

func (*Server) SetScore(ctx context.Context, req *gameenginepb.SetScoreRequest) (*gameenginepb.SetScoreResponse, error) {
	zlog.Info().Msg("SetScore method in m-game-engine service called")
	return &gameenginepb.SetScoreResponse{Result: logic.SetScore(req.GetScore())}, nil
}

func (s *Server) DoServe() error {
	var err error
	s.lis, err = net.Listen("tcp", s.address)
	if err != nil {
		zlog.Error().Msg("Failed to listen service")
		return fmt.Errorf("failed to listen server port %w", err)
	}
	s.srv = grpc.NewServer()
	gameenginepb.RegisterGameEngineServer(s.srv, s)
	zlog.Info().Str("address", s.address).Msg("gRPC gameengine microservice is started.")
	if err = s.srv.Serve(s.lis); err != nil {
		zlog.Error().Msg("Failed to serve server for gameengine microservice")
		return fmt.Errorf("failed to serve server for gameengine microservice %w", err)
	}
	return nil
}

func (s *Server) StopServer() {
	zlog.Info().Msg("Stopping the Gameengine service Server...")
	s.srv.Stop()
	zlog.Info().Msg("Closing the Listener...")
	_ = s.lis.Close()
	zlog.Info().Msg("End of program.")
}


