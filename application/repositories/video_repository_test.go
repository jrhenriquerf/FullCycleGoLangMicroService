package repositories_test

import (
	"testing"
	"time"

	"github.com/jrhenriquerf/FullCycleGoLangMicroService/application/repositories"
	"github.com/jrhenriquerf/FullCycleGoLangMicroService/domain"
	"github.com/jrhenriquerf/FullCycleGoLangMicroService/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close() //Fecha a conexão quando percebe que não está mais sendo utilizada

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}
