package repositories_test

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jrhenriquerf/FullCycleGoLangMicroService/application/repositories"
	"github.com/jrhenriquerf/FullCycleGoLangMicroService/domain"
	"github.com/jrhenriquerf/FullCycleGoLangMicroService/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func insertJobWithVideo() (*domain.Video, *domain.Job, repositories.JobRepositoryDb, *gorm.DB, error) {
	db := database.NewDbTest()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	return video, job, repoJob, db, err
}

func TestJobRepositoryInsert(t *testing.T) {
	video, job, repoJob, db, err := insertJobWithVideo()
	defer db.Close() //Fecha a conexão quando percebe que não está mais sendo utilizada

	require.Nil(t, err)
	j, err := repoJob.Find(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)
}

func TestJobRepositoryUpdate(t *testing.T) {
	_, job, repoJob, db, err := insertJobWithVideo()
	defer db.Close() //Fecha a conexão quando percebe que não está mais sendo utilizada

	require.Nil(t, err)

	job.Status = "Complete"

	repoJob.Update(job)

	j, err := repoJob.Find(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)
}
