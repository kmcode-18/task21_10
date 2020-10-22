package store

/*
func TestShouldUpdateStats(t *testing.T) {
	dBase, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer dBase.Close()
	dddd, err := gorm.Open("postgres", dBase)
	if err !=nil{

	}
	db = dddd
	//mock.ExpectExec("INSERT images").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO images where image_name= $1").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	// now we execute our method
	if err = AddImage(Image{
		ImageName: "test",
	}); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
*/
/*
import (
	"database/sql"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	repository Repository
	images     *[]Image
}

func (s *Suite) SetupSuite() {
	var (
		dBase  *sql.DB
		err error
	)
	dBase, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", dBase)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)
	db = s.DB
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestAddImage() {
	var (
		id   = 1
		name = "test-name"
		time =time.Now()
	)

	/*s.mock.ExpectQuery("SELECT * FROM images WHERE (image_id=$1)").
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"image_id", "image_name","created_at"}).AddRow(id, name,time))
*/
	/*mockedRow := sqlmock.NewRows([]string{"image_id", "image_name","created_at"}).AddRow(id, name,time)
	s.mock.ExpectQuery(`SELECT * FROM images WHERE (image_id= ?)`).
		WithArgs(id).
		WillReturnRows(mockedRow)
	res, err := GetImages(QueryDetails{
		ImageId:   id,
	})

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal([]Image{{
		ImageID:   id,
		ImageName: "test_name",
		CreatedAt: time,
	}}, res))
}*/