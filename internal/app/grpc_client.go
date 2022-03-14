package app

import (
	"context"
	"fmt"
	"playlist-server/models"

	api "github.com/kostyasolovev/youtube_pb_api"

	"google.golang.org/grpc"
)

func (server *Server) dialGRPC(ctx context.Context, port string) (err error) {
	server.grpcConn, err = grpc.DialContext(ctx, fmt.Sprintf(":%s", port), grpc.WithInsecure())
	return err
}

func (server *Server) registerGRPCClient(ctx context.Context) (err error) {
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
			return nil, err
		}

		return resp, nil
	}
}
