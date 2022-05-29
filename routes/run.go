package routes

import (
	"context"
	"diveEvolution/db"
	"diveEvolution/models"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
)

var Client *mongo.Client
func init() {
	Client = db.ConnectDB()
}
func RunApp(){
	router := mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/api/getIndex/{lang}", IndexHandler).Methods("GET")
	router.HandleFunc("/api/getAbout/{lang}", AboutHandler).Methods("GET")
	router.HandleFunc("/api/getCourses/{lang}", CoursesHandler).Methods("GET")
	router.HandleFunc("/api/getContact/{lang}", ContactHandler).Methods("GET")
	router.HandleFunc("/api/getTours/{lang}/{destination}", ToursHandler).Methods("GET")
	router.HandleFunc("/api/getFooter/{lang}", FooterHandler).Methods("GET")
	router.HandleFunc("/api/getCourseInfo/{lang}/{ref}", CoursesInfoHandler).Methods("GET")
	router.HandleFunc("/api/getToursInfo/{lang}/{ref}", ToursInfoHandler).Methods("GET")
	router.HandleFunc("/api/getIndexImg", IndexImgHandler).Methods("GET")
	router.HandleFunc("/api/getAboutImg", AboutImgHandler).Methods("GET")
	router.HandleFunc("/api/getCoursesImg", CoursesImgHandler).Methods("GET")
	router.HandleFunc("/api/getContactImg", ContactImgHandler).Methods("GET")
	router.HandleFunc("/api/getToursImg", ToursImgHandler).Methods("GET")
	router.HandleFunc("/api/getHeaderImg", HeaderImgHandler).Methods("GET")
	router.HandleFunc("/api/getFooterImg", FooterImgHandler).Methods("GET")
	router.HandleFunc("/api/getCourseInfoImg/{ref}", CoursesInfoImgHandler).Methods("GET")


	router.HandleFunc("/api/updateIndex", UpdateIndexHandler).Methods("GET")
	router.HandleFunc("/api/updateHeader", writeHeader).Methods("GET")
	router.HandleFunc("/api/updateFooter", writeFooter).Methods("GET")
	router.HandleFunc("/api/updateIndexBody", writeIndexBody).Methods("GET")
	router.HandleFunc("/api/updateContactBody", writeContactBody).Methods("GET")
	router.HandleFunc("/api/updateNosotros", writeNosotrosBody).Methods("GET")
	router.HandleFunc("/api/updateCourses", writeCoursesBody).Methods("GET")
	router.HandleFunc("/api/updateTours", writeToursBody).Methods("GET")
	router.HandleFunc("/api/updateContactImg", writeContactImg).Methods("GET")
	router.HandleFunc("/api/updateNosotrosImg", writeAboutImg).Methods("GET")
	router.HandleFunc("/api/updateCoursesImg", writeCoursesImg).Methods("GET")
	router.HandleFunc("/api/updateToursImg", writeToursImg).Methods("GET")
	router.HandleFunc("/api/updateCourseInfo", writeCourseInfoBody).Methods("GET")
	router.HandleFunc("/api/updateCourseInfoImg", writeCourseInfoImg).Methods("GET")
	router.HandleFunc("/api/updateTourInfoBody", writeTourInfoBody).Methods("GET")

	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origin := handlers.AllowedOrigins([]string{"*"})
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, handlers.CORS(methods, origin)(router))
}
func Home (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("/api/getIndex/es\n/api/getHeader/es\n/api/getFooter/es\n" +
		"/api/getAbout/es\n/api/getCourses/es\n/api/getContact/es\n/api/getTours/fr/sc\n" +
		"/api/getCourseInfo/{lang}/{ref}\n/api/getToursInfo/{lang}/{ref}\n\n\n"+
		"/api/getIndexImg\n/api/getHeaderImg\n/api/getFooterImg\n"+
		"/api/getAboutImg\n/api/getCoursesImg\n/api/getContactImg\n/api/getToursImg\n"+
		"/api/getCourseInfoImg/{ref}"))
}
func writeHeader(w http.ResponseWriter, r *http.Request){
	data := models.HeaderImg{
		Id: uuid.NewV4().String(),
		Logo: "https://via.placeholder.com/150/0000FF/808080",
		LangItems: []models.LangItem{
			{
				Lang: "es",
				Img:  "https://res.cloudinary.com/logicielapplab/image/upload/v1651212913/DiveEvolution/Header/es_e6qn2t.svg",
			},
			{
				Lang: "en",
				Img:  "https://res.cloudinary.com/logicielapplab/image/upload/v1651212904/DiveEvolution/Header/gb_gxyw57.svg",
			},
			{
				Lang: "fr",
				Img:  "https://res.cloudinary.com/logicielapplab/image/upload/v1651212904/DiveEvolution/Header/fr_bta7cj.svg",
			},
		},
	}

	HeaderImg := Client.Database("DiveEvolution").Collection("HeaderImg")
	HeaderImg.InsertOne(context.TODO(),data)
}
func writeFooter(w http.ResponseWriter, r *http.Request){
	data := models.FooterImg{
		Id: uuid.NewV4().String(),
		Logo:"https://via.placeholder.com/150/0000FF/808080",
		Background: "https://via.placeholder.com/2000x800/000000/808080",
		WhatsApp: "https://res.cloudinary.com/logicielapplab/image/upload/v1651350046/DiveEvolution/Footer/3225179_app_logo_media_popular_social_icon_juvvzd.svg",
		SocialMediaIcons: models.SocialMediaIcons{
			Facebook: "https://res.cloudinary.com/logicielapplab/image/upload/v1651296519/DiveEvolution/Footer/3225194_app_facebook_logo_media_popular_icon_wrmery.svg",
			Instagram: "https://res.cloudinary.com/logicielapplab/image/upload/v1651296520/DiveEvolution/Footer/3225191_app_instagram_logo_media_popular_icon_hokrzm.svg",
			Twitter: "https://res.cloudinary.com/logicielapplab/image/upload/v1651308103/DiveEvolution/Footer/3225183_app_logo_media_popular_social_icon_hjp5gw.svg",
		},
	}

	HeaderImg := Client.Database("DiveEvolution").Collection("FooterImg")
	HeaderImg.InsertOne(context.TODO(),data)
}
func writeIndexBody(w http.ResponseWriter, r *http.Request){
	data := models.IndexImg{
		Id: uuid.NewV4().String(),
		BodyImg: models.BodyImg{
			Background: "https://via.placeholder.com/2000x900/6a5acd/fffffff",
			Section1: models.Section1{
				Calidad: []string{
					"https://via.placeholder.com/1000x900/d1d1e0/000000",
					"https://res.cloudinary.com/logicielapplab/image/upload/v1651347812/DiveEvolution/IndexBody/award_pifluk.svg",
				},
				Precio: []string{
					"https://via.placeholder.com/1000x900/d1d1e0/000000",
					"https://res.cloudinary.com/logicielapplab/image/upload/v1651347812/DiveEvolution/IndexBody/currency-dollar_koyukh.svg",
				},
			},
			Section2Img: models.Section2Img{
				ItemsImg: []string{
					"https://via.placeholder.com/2000x900/ff6699/ffffff",
					"https://via.placeholder.com/2000x900/ffad33/ffffff",
				},
			},
		},
	}

	IndexBodyImg := Client.Database("DiveEvolution").Collection("IndexImg")
	IndexBodyImg.InsertOne(context.TODO(),data)
}
func writeContactBody(w http.ResponseWriter, r *http.Request){
	data := models.Contact{
		Id: uuid.NewV4().String(),
		Title: "Lorem Ipsum",
		Introduction: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
		ContactInfo: models.ContactInfo{
			Address: "AV. Isabela y Guillermo Vega Gallegos",
			Phone: "(+593) 98 229 1894",
			Email: "jhonatan.quihuiri@gmail.com",
		},
		Button: "Lorem Ipsum",
		FormTitle: "Lorem Ipsum",
	}
	ContactBody := Client.Database("DiveEvolution").Collection("Contact")
	ContactBody.InsertOne(context.TODO(),data)
}
func writeContactImg(w http.ResponseWriter, r *http.Request){
	data := models.ContactImg{
		Id: uuid.NewV4().String(),
		Background: "https://via.placeholder.com/2000x900/6a5acd/fffffff",
		Foreground: "https://via.placeholder.com/350x500/ff0066/fffffff",
		Form: "https://via.placeholder.com/2000x900/00ff00/fffffff",
	}
	ContactImg := Client.Database("DiveEvolution").Collection("ContactImg")
	ContactImg.InsertOne(context.TODO(),data)
}
func writeNosotrosBody(w http.ResponseWriter, r *http.Request){
	data := models.About{
		Id: uuid.NewV4().String(),
		History: models.History{
			Introduction: "FRLorem Ipsum is simply dummy text",
			Brief: "FRLorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
			Question: "What is Lorem Ipsum?",
			Button: "Lorem Ipsum",
		},
		Mission: models.MVission{
			Title: "FRMision",
			Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
		},
		Vission: models.MVission{
			Title: "FRVision",
			Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
		},
		Value: models.Value{
			Title: "FRValores",
			Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
			Values: []string{"Claidad", "Profesionalismo","Responsabilidad","Honestidad"},
		},
	}
	AboutBody := Client.Database("DiveEvolution").Collection("About")
	AboutBody.InsertOne(context.TODO(),data)
}
func writeAboutImg(w http.ResponseWriter, r *http.Request){
	data := models.AboutImg{
		Id: uuid.NewV4().String(),
		HistoryImg: models.HistoryImg{
			Background: "https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			Logo: "https://res.cloudinary.com/logicielapplab/image/upload/v1652031313/DiveEvolution/About/penguin-gb2d6782e2_1920_xrirai.jpg",
		},
		VissionImg: "https://res.cloudinary.com/logicielapplab/image/upload/v1652031313/DiveEvolution/About/penguin-gb2d6782e2_1920_xrirai.jpg",
		ValueImg: []string{"","","",""},
	}
	ContactImg := Client.Database("DiveEvolution").Collection("AboutImg")
	ContactImg.InsertOne(context.TODO(),data)
}
func writeCoursesBody(w http.ResponseWriter, r *http.Request){
	data := models.Courses{
		Id: uuid.NewV4().String(),
		Methodology: models.Methodology{
			Title: "FRMetodologia",
			Methods: []models.Method{
				models.Method{
					Title: "EVO PSS",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Question1: "Como funciona",
					Works: "Lorem Ipsum is simply dummy text",
					Question2: "Ventajas",
					Advantages: []string{"Hola","Mundo","Logiciel"},
				},
				models.Method{
					Title: "E-Learning",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Question1: "Como funciona",
					Works: "Lorem Ipsum is simply dummy text",
					Question2: "Ventajas",
					Advantages: []string{"Hola","Mundo","Logiciel"},
				},
				models.Method{
					Title: "Plataforma PSS",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Question1: "Como funciona",
					Works: "Lorem Ipsum is simply dummy text",
					Question2: "Ventajas",
					Advantages: []string{"Hola","Mundo","Logiciel"},
				},
			},
		},
		CourseTypes: []models.Course{
			models.Course{
				Ref: uuid.NewV4().String(),
				Title: "Open Water",
				Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
				Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
				Price: 250,
			},
			models.Course{
				Ref: uuid.NewV4().String(),
				Title: "Intermedium",
				Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
				Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
				Price: 250,
			},
			models.Course{
				Ref: uuid.NewV4().String(),
				Title: "Advanced",
				Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
				Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
				Price: 250,
			},
			models.Course{
				Ref: uuid.NewV4().String(),
				Title: "Dive Master",
				Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
				Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
				Price: 250,
			},
		},
	}
	ContactBody := Client.Database("DiveEvolution").Collection("Courses")
	ContactBody.InsertOne(context.TODO(),data)
}
func writeCoursesImg(w http.ResponseWriter, r *http.Request){
	data := models.CoursesImg{
		Id: uuid.NewV4().String(),
		MethodsImg: []string{
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
		},
		CoursesImg: []string{
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
		},
	}
	ContactBody := Client.Database("DiveEvolution").Collection("CoursesImg")
	ContactBody.InsertOne(context.TODO(),data)
}
func writeToursBody(w http.ResponseWriter, r *http.Request){
	data := models.Tours{
		Id: uuid.NewV4().String(),
		Sc: models.TourType{
			Snorkel: []models.Course{
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "Open Water",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRIntermedium",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRAdvanced",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRDive Master",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
			},
			Dive: []models.Course{
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FROpen Water",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRIntermedium",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRAdvanced",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRDive Master",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
			},
			Land: []models.Course{
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FROpen Water",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRIntermedium",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRAdvanced",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRDive Master",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
			},
		},
		Sx: models.TourType{
			Snorkel: []models.Course{
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FROpen Water",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRIntermedium",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRAdvanced",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRDive Master",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
			},
			Dive: []models.Course{
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FROpen Water",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRIntermedium",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRAdvanced",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRDive Master",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
			},
			Land: []models.Course{
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FROpen Water",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRIntermedium",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRAdvanced",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRDive Master",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
			},
		},
		Ib: models.TourType{
			Snorkel: []models.Course{
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FROpen Water",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRIntermedium",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRAdvanced",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRDive Master",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
			},
			Dive: []models.Course{
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FROpen Water",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRIntermedium",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRAdvanced",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRDive Master",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
			},
			Land: []models.Course{
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FROpen Water",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRIntermedium",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRAdvanced",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				models.Course{
					Ref: uuid.NewV4().String(),
					Title: "FRDive Master",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
			},
		},
	}

	ToursBody := Client.Database("DiveEvolution").Collection("Tours")
	ToursBody.InsertOne(context.TODO(),data)
}
func writeToursImg(w http.ResponseWriter, r *http.Request){
	data := models.ToursImg{
		Id: uuid.NewV4().String(),
		SnorkelImg: []string{
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
		},
		DiveImg: []string{
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
		},
		LandImg: []string{
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1652031292/DiveEvolution/About/sea-g3cb441ce3_1920_i0pxzq.jpg",
		},
	}
	ToursImg := Client.Database("DiveEvolution").Collection("ToursImg")
	ToursImg.InsertOne(context.TODO(),data)
}
func writeCourseInfoBody(w http.ResponseWriter, r *http.Request){
	data := models.CourseInfo{
		Id: uuid.NewV4().String(),
		Ref: "",
		Lang: "es",
		Course: models.Course{
			Ref: uuid.NewV4().String(),
			Title: "Open Water",
			Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
			Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
			Price: 250,
		},
		LargeDescription: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
		Itinerary: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
	}
	ContactBody := Client.Database("DiveEvolution").Collection("CoursesInfo")
	ContactBody.InsertOne(context.TODO(),data)
}
func writeCourseInfoImg(w http.ResponseWriter, r *http.Request){
	data := models.CourseInfoImg{
		Id: "861da9db-74ee-4ba9-9978-585e3fb5e1c2",
		Img: []string{
			"https://res.cloudinary.com/logicielapplab/image/upload/v1653764788/DiveEvolution/CourseDetail/curso2_okynik.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1653764783/DiveEvolution/CourseDetail/courso1_xawoph.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1653764793/DiveEvolution/CourseDetail/curso4_sezzea.jpg",
			"https://res.cloudinary.com/logicielapplab/image/upload/v1653764793/DiveEvolution/CourseDetail/curso3_nm8wsz.jpg",
		},
	}
	ContactBody := Client.Database("DiveEvolution").Collection("CoursesInfoImg")
	ContactBody.InsertOne(context.TODO(),data)
}
func writeTourInfoBody(w http.ResponseWriter, r *http.Request){
	languages := []string{"es","en","fr"}
	references := []string{
		"4f1600d3-2489-4ecf-be4e-928c779f2ed3",
		"e835b382-78ec-4da7-9964-c0f5cfc6ca10",
		"e6fe1f48-529d-488f-a2fc-f18b5ce2ed5f",
		"1c151a78-a044-433f-8554-53acb9ad115b",
		"b9cc6767-fa65-4785-80a8-8c36a8fafbc8",
		"e7222bd0-9c0a-4303-8e93-d01bd687f69b",
		"40c46eb4-2288-4ee5-8890-19e3d94f1522",
		"cf80e1f7-64ec-4343-b98f-5658853915a8",
		"215352f7-d636-48e4-a67d-08b9ce297fe0",
		"9ff5fd3f-c75d-40ed-a2a7-6e2c7b3fdc7c",
		"7c009c2e-4e62-4058-b681-04edec6e7115",
		"519727bc-b15d-45a0-838f-f1cc92878abb",
		"5e61e82a-566e-48ce-a0ee-79e12c5ff7c1",
		"dd7aba04-343c-4efd-9bf8-5cc2623e3c16",
		"b3cc13c0-9d4d-4302-999a-ffc9724b99d8",
		"fb44d347-1868-4aaf-81b5-6ff3b8329c52",
		"ecdd249d-b609-4c21-b647-ba2ac37f5761",
		"4cb150c4-f131-4058-a0df-cbc70092982f",
		"eecc8aea-fb69-4f86-bd35-22bdd16ab2bc",
		"9d0fb57f-4caf-4c3c-9550-5c626bf8d9ea",
		"9f933f0b-acb3-4175-8c27-3744bc612ac1",
		"a6d73e10-cc2a-4fd3-afd1-50199fdc765c",
		"ac2c7f19-8f69-4e0a-890b-d90deab2dcc8",
		"a1f01a72-9d97-476f-8e1a-c8131e249d44",
		"551cbf9e-cbcf-49f7-848e-20ce39bbe562",
		"00661474-fd6b-4db7-92b6-7386b32c41ee",
		"2997b199-cbfd-4328-831d-3538fac49481",
		"f0f60f2b-a987-4e17-8b32-81429cb9f5e5",
		"0881c5c7-00a7-4125-b7f1-406932f2042f",
		"9d340f1d-8ba2-4f33-8b11-8287338d65f8",
		"6ff7ecec-9934-44f1-b621-0e16dc4e9328",
		"565c167b-521f-4411-ace6-aa7ef9574c2b",
		"ea831292-9cdf-49ba-aec7-9d97ce9fe271",
		"67cdb1c9-a327-415f-825f-cd5c3ac22525",
		"7cd2e9c2-35d6-4b7b-94ea-08fe04793702",
		"cabd8226-132e-4f18-a9b0-a225ec979bc7",
	}
	ContactBody := Client.Database("DiveEvolution").Collection("ToursInfo")
	for i,ref := range references {
		fmt.Println(i)
		for _, lang := range languages {
			data := models.CourseInfo{
				Id: uuid.NewV4().String(),
				Ref: ref,
				Lang: lang,
				Course: models.Course{
					Title: "Open Water",
					Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					Included: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
					Price: 250,
				},
				LargeDescription: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
				Itinerary: []string{"Hola","Mundo","Logiciel","Applab","Saludos"},
			}
			ContactBody.InsertOne(context.TODO(),data)
		}
	}

}