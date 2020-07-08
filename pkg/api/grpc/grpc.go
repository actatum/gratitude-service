package grpc

import ()

// Run starts grpc server
/*
func Run(ctx context.Context, service gratitude.GratitudeServiceServer, port string) error {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		return errs.Wrap(err, "api.grpc.Run")
	}

	server := grpc.NewServer()
	service.RegisterServiceServer(server, service)

	log.Println("starting gRPC server on port " + port)
	return server.Serve(listen)
}
*/
