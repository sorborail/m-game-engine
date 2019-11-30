package main

import (
	"context"
	"flag"
	zlog "github.com/rs/zerolog/log"
	gameenginepb "github.com/sorborail/m-apis/game-enginepb/v1"
	"google.golang.org/grpc"
	"time"
)

func main() {
	zlog.Info().Msg("Begin starting gameengine client...")
	var addrPtr = flag.String("address", "localhost:60051", "address to connect gameengine service")
	flag.Parse()
	zlog.Debug().Msgf("Value of addrPtr - %s", *addrPtr)
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(*addrPtr, opts)
	if err != nil {
		zlog.Fatal().Err(err).Str("address", *addrPtr).Msg("could not connect to the gameengine server")
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			zlog.Error().Err(err).Str("address", *addrPtr).Msg("Failed to close connection")
		}
	}()
	cl := gameenginepb.NewGameEngineClient(conn)
	if cl == nil {
		zlog.Error().Msg("Cannot connection to the gameengine service")
	} else {
		zlog.Info().Msg("Gameengine client is started")
		doGetSize(cl, addrPtr)
		doSetScore(cl, addrPtr)
	}
}

func doGetSize(cl gameenginepb.GameEngineClient, addr *string) {
	zlog.Info().Msg("Begin get size request...")
	req := &gameenginepb.GetSizeRequest{}
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()
	res, err := cl.GetSize(ctx, req)
	if err != nil {
		zlog.Fatal().Err(err).Str("address", *addr).Msg("error happened while get size response")
	}
	zlog.Info().Interface("get size", res.GetSize()).Msg("Value from GetSize method gameengine service")
}

func doSetScore(cl gameenginepb.GameEngineClient, addr *string) {
	zlog.Info().Msg("Begin set score request...")
	req := &gameenginepb.SetScoreRequest{Score: 1.1}
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()
	res, err := cl.SetScore(ctx, req)
	if err != nil {
		zlog.Fatal().Err(err).Str("address", *addr).Msg("error happened while set score response")
	}
	zlog.Info().Interface("set score", res.GetResult()).Msg("Value from SetScore method gameengine service")
}
