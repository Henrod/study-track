/*
Copyright © 2020 henrod henrique.93.rodrigues@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Henrod/study-track/internal/controller"
	"github.com/Henrod/study-track/internal/storage/db"
	"github.com/Henrod/study-track/pkg/studytrack"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "starts grpc and http api",
	Long:  `starts grpc and http api`,
	Run: func(cmd *cobra.Command, args []string) {
		grpcServer := grpc.NewServer()
		wait := make(chan struct{})
		grpcEndpoint := ":8080"
		httpEndpoint := ":8081"

		// http
		go func() {
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			mux := runtime.NewServeMux()
			opts := []grpc.DialOption{grpc.WithInsecure()}
			err := studytrack.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
			if err != nil {
				log.Fatalf("failed to register: %v", err)
			}

			if err = http.ListenAndServe(httpEndpoint, mux); err != nil {
				log.Fatalf("failed to serve http: %v", err)
			}
		}()

		// grpc
		go func() {
			lis, err := net.Listen("tcp", grpcEndpoint)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}

			logger := logrus.New()
			pg, err := sql.Open(
				"postgres",
				"postgres://postgres:@localhost:9000/studytrack?sslmode=disable")
			if err != nil {
				log.Fatalf("failed to connect to db: %v", err)
			}

			storage := db.New(pg)
			studytrack.RegisterUserServiceServer(grpcServer, controller.NewUser(storage, logger))
			studytrack.RegisterSubjectServiceServer(grpcServer, controller.NewSubject(storage, logger))
			if err = grpcServer.Serve(lis); err != nil {
				log.Fatalf("failed to serve grpc: %v", err)
			}
		}()

		<-wait
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
