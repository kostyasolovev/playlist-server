package app

import (
	"context"
	"fmt"

	api "github.com/kostyasolovev/youtube_pb_api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"playlist-server/models"
)

func (server *Server) dialGRPC(ctx context.Context, port string) (err error) {
	server.grpcConn, err = grpc.DialContext(ctx, fmt.Sprintf(":%s", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	return errors.Wrap(err, "dialing grpc error")
}

func (server *Server) registerGRPCClient(context.Context) (err error) {
	server.grpcCli = api.NewYoutubePlaylistClient(server.grpcConn)

	return err
}

func (server *Server) registerDefaultListGRPCFunc() func(context.Context, string) (models.Playlist, error) {
	return func(ctx context.Context, playlistId string) (models.Playlist, error) {
		resp, err := server.grpcCli.List(
			ctx, &api.PlaylistRequest{
				Id: playlistId,
			},
		)
		if err != nil {
			return nil, errors.Wrap(err, "youtubeService response error")
		}

		return resp, nil
	}
}
