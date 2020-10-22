package store

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	"log"
	"time"
)

const (
	host     = ""
	port     = 5432
	user     = ""
	password = ""
	dbname   = ""
)

var db *gorm.DB
func InitDbConn()  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbase, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("error : ", err.Error())
		panic(err)
	}
	db=dbase
	log.Println("database connection created")
	if !db.HasTable(&Image{}){
		db = db.CreateTable(&Image{})
		if db.Error!= nil{
			panic(err)
		}
	}
}

func init() {
	InitDbConn()
}
type Repository interface {
	AddImage(imageDetails Image) (err error)
	GetImages(qParams QueryDetails)([]Image,error)
}
type Image struct {
	ImageID     int   `gorm:"primary_key;autoIncrement:true:image_id" json:"image_id"`
	ImageName	string	`gorm:"size:75;not null" json:"image_name"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}
type QueryDetails struct {
	ImageId int
	OffSet int
	Limit int
	SortOrder string
	SortBy string
}
func AddImage(imageDetails Image) (err error) {
	if db == nil {
		err = errors.New("unable to connect to database")
		log.Println("error : ", err.Error())
		return
	}
	var q *gorm.DB
	q=db
	q = q.Create(&imageDetails)
	if q.Error!=nil{
		return err
	}
	return
}

func CloseDbConn() (err error) {
	err = db.Close()
	return
}

func GetImages(qParams QueryDetails)([]Image,error){
	var q *gorm.DB
	q=db
	var images []Image
	if qParams.ImageId>0{
		q=q.Where("image_id= ?",qParams.ImageId)
	}
	if qParams.Limit>0{
		q=q.Limit(qParams.Limit)
	}
	if qParams.OffSet>0{
		q=q.Offset(qParams.OffSet)
	}
	if qParams.SortBy != "" && qParams.SortOrder !=""{
		q=q.Order(qParams.SortBy +" " +qParams.SortOrder)
	}
	q=q.Find(&images)
	if errors.Is(q.Error, gorm.ErrRecordNotFound){
		return images,errors.New("no record found")
	}
	return images,nil
}
