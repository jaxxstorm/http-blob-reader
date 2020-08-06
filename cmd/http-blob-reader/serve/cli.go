package serve

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/jaxxstorm/http-blob-reader/pkg/cloudblob"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	bucketURL     string
	blobKey       string
	listenAddress string
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "start the webserver",
		Long:  `Start a webserver`,
		RunE: func(cmd *cobra.Command, args []string) error {


			bucketURL = viper.GetString("bucket")
			blobKey = viper.GetString("blob-key")

			if bucketURL == "" {
				return fmt.Errorf("must specify a bucket URL, eg: s3://bucket | gs://bucket | azblob://bucket")
			}

			if blobKey == "" {
				return fmt.Errorf("must specify a bucket URL")
			}

			reader, err := cloudblob.Read(bucketURL, blobKey)

			if err != nil {
				return err
			}

			zerolog.SetGlobalLevel(zerolog.InfoLevel)
			if gin.IsDebugging() {
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			}

			log.Logger = log.Output(
				zerolog.ConsoleWriter{
					Out:     os.Stderr,
					NoColor: false,
				},
			)

			r := gin.New()
			r.Use(logger.SetLogger())
			r.Use(gin.Recovery())

			r.GET("/", gin.WrapF(func(w http.ResponseWriter, req *http.Request) {
				_, err = io.Copy(w, reader)
			}))

			r.Run(listenAddress)
			return nil
		},
	}

	command.Flags().StringVarP(&bucketURL, "bucket", "b", "", "the bucket URL to server")
	command.Flags().StringVarP(&blobKey, "blob-key", "k", "", "the bucket blob to serve")
	command.Flags().StringVarP(&listenAddress, "listen-address", "l", "0.0.0.0:8080", "the http ip and port to listen on")

	viper.BindEnv("bucket", "BUCKET_ADDRESS")
	viper.BindPFlag("bucket", command.Flags().Lookup("bucket"))

	viper.BindEnv("blob-key", "BLOB_KEY")
	viper.BindPFlag("blob-key", command.Flags().Lookup("blob-key"))


	return command
}
