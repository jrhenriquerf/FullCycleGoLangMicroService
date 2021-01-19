package service_test

import (
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/jrhenriquerf/FullCycleGoLangMicroService/application/repositories"
	service "github.com/jrhenriquerf/FullCycleGoLangMicroService/application/services"
	"github.com/jrhenriquerf/FullCycleGoLangMicroService/domain"
	"github.com/jrhenriquerf/FullCycleGoLangMicroService/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file!")
	}
}

func prepare() (*domain.Video, repositories.VideoRepositoryDb) {
	db := database.NewDbTest()
	defer db.Close() //Fecha a conexão quando percebe que não está mais sendo utilizada

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "convite.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	return video, repo
}

func TestVideoServiceDownload(t *testing.T) {

	video, repo := prepare()

	videoService := service.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("codeeducationtest")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)
}
