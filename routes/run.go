package routes

import (
	"context"
	"diveEvolution/db"
	"diveEvolution/models"
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
	router.HandleFunc("/api/getIndexImg", IndexImgHandler).Methods("GET")
	router.HandleFunc("/api/getAboutImg", AboutImgHandler).Methods("GET")
	router.HandleFunc("/api/getCoursesImg", CoursesImgHandler).Methods("GET")
	router.HandleFunc("/api/getContactImg", ContactImgHandler).Methods("GET")
	router.HandleFunc("/api/getToursImg", ToursImgHandler).Methods("GET")
	router.HandleFunc("/api/getHeaderImg", HeaderImgHandler).Methods("GET")
	router.HandleFunc("/api/getFooterImg", FooterImgHandler).Methods("GET")


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

	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origin := handlers.AllowedOrigins([]string{"*"})
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, handlers.CORS(methods, origin)(router))
}
func Home (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("/api/getIndex/es\n/api/getHeader/es\n/api/getFooter/es\n" +
		"/api/getAbout/es\n/api/getCourses/es\n/api/getContact/es\n/api/getTours/fr/sc\n" +
		"/api/getCourseInfo/{lang}/{ref}\n\n\n"+
		"/api/getIndexImg\n/api/getHeaderImg\n/api/getFooterImg\n"+
		"/api/getAboutImg\n/api/getCoursesImg\n/api/getContactImg\n/api/getToursImg\n"))
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