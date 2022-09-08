package db

import (
	"context"
	"diveEvolution/models"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectDB() *mongo.Client {
	opt := options.Client().ApplyURI("mongodb+srv://doadmin:Z3d87ni4E91g05aX@logiciel-applab-dab57134.mongo.ondigitalocean.com/admin?authSource=admin&replicaSet=logiciel-applab&tls=true&tlsCAFile=db/ca-certificate.crt")
	client, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func GetDocument(collection *mongo.Collection, id string) bson.M {
	opts := options.Find().SetProjection(bson.D{{"_id", 0}})
	cur, err := collection.Find(context.TODO(), bson.D{{"_id", id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	var result bson.M
	for cur.Next(context.TODO()) {
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
	}
	return result
}
func GetDocumentDest(collection *mongo.Collection, id string, destination string) models.TourType {
	opts := options.Find().SetProjection(bson.D{{"_id", 0}})
	cur, err := collection.Find(context.TODO(), bson.D{{"_id", id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	var result models.Tours
	for cur.Next(context.TODO()) {
		err := cur.Decode(&result)
		if err != nil {
			fmt.Println("ERROR AL DECODIFICAR")
			log.Fatal(err)
		}
	}
	if destination == "sc" {
		return result.Sc
	} else if destination == "sx" {
		return result.Sx
	} else if destination == "ib" {
		return result.Ib
	} else {
		return result.Sc
	}
}
func GetDocumentCourse(collection *mongo.Collection, lang string, ref string) models.CourseInfo {
	opts := options.Find().SetProjection(bson.D{{"_id", 0}})
	cur, err := collection.Find(context.TODO(), bson.D{{"ref", ref}, {"lang", lang}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	var result models.CourseInfo
	for cur.Next(context.TODO()) {
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(result)
	return result
}
func UpdateDocumment(collections []*mongo.Collection) {
	navBar := models.NavBar{
		NavBarItems: []models.NavBarItem{
			models.NavBarItem{
				Title:     "Inicio",
				Accordion: []string{},
			},
			models.NavBarItem{
				Title:     "Nosotros",
				Accordion: []string{},
			},
			models.NavBarItem{
				Title: "Tours",
				Accordion: []string{
					"San Cristóbal",
					"Santa Cruz",
					"Isabela",
				},
			},
			models.NavBarItem{
				Title:     "Cursos de Buceo",
				Accordion: []string{},
			},
			models.NavBarItem{
				Title:     "Contacto",
				Accordion: []string{},
			},
		},
	}
	//header := models.Header{
	//	Id: uuid.NewV4().String(),
	//	NavBar: navBar,
	//	Title: "Lorem Ipsum is simply dummy text",
	//	Button: "Lorem Ipsum",
	//}
	footer := models.Footer{
		Id:     uuid.NewV4().String(),
		NavBar: navBar,
		Phone:  "+593982291894",
		SocialMedia: models.SocialMedia{
			Facebook:  "https://www.facebook.com/Logiciel-AppLab-115115559920070",
			Instagram: "https://www.instagram.com/logicielapplab/",
			Twitter:   "",
		},
	}
	data := models.Index{
		Id: uuid.NewV4().String(),
		Body: models.Body{
			Section1: models.Section1{
				Calidad: []string{"Lorem Ipsum", "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s"},
				Precio:  []string{"Lorem Ipsum", "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s"},
			},
			Section2: models.Section2{
				Items: []models.ItemsInfo{
					models.ItemsInfo{
						Title:       "Lorem Ipsum",
						Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
						Button:      "Lorem Ipsum",
					},
					models.ItemsInfo{
						Title:       "Lorem Ipsum",
						Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
						Button:      "Lorem Ipsum",
					},
					models.ItemsInfo{
						Title:       "Lorem Ipsum",
						Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
						Button:      "Lorem Ipsum",
					},
				},
			},
		},
	}
	collections[0].InsertOne(context.TODO(), data)
	collections[1].InsertOne(context.TODO(), footer)

}
