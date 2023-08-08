package utils

func GetLang(lang, routeName string) string {
	if routeName == "index" {
		if lang == "es" {
			return "4164da5e-cacd-4827-8891-26945019a5be"
		} else if lang == "en" {
			return "5da9ff07-3a11-4cc9-8d0a-f9a0e6daf925"
		} else if lang == "fr" {
			return "076d9ef9-bc59-40c7-ae7f-6a9e90cae856"
		} else {
			return "4164da5e-cacd-4827-8891-26945019a5be"
		}
	} else if routeName == "header" {
		if lang == "es" {
			return "76ff0c4a-ca1c-4d62-9304-6e3a71565ff4"
		} else if lang == "en" {
			return "a6b3948a-9f61-4c20-b1b6-e60628c70958"
		} else if lang == "fr" {
			return "f46c2a6c-54cb-4b41-b97f-7de4469b4df1"
		} else {
			return "76ff0c4a-ca1c-4d62-9304-6e3a71565ff4"
		}
	} else if routeName == "footer" {
		if lang == "es" {
			return "f93746d6-8b27-481f-ad1f-e888f7ef6d0f"
		} else if lang == "en" {
			return "4601db41-e3be-4f45-9809-5e3f67d43222"
		} else if lang == "fr" {
			return "f21aa619-b746-4b50-8381-de8956eba4df"
		} else {
			return "f93746d6-8b27-481f-ad1f-e888f7ef6d0f"
		}
	} else if routeName == "contact" {
		if lang == "es" {
			return "a310cb1b-21ca-42f8-936c-4c0fe117a2f0"
		} else if lang == "en" {
			return "5c240292-e280-4f49-8f6c-342295160271"
		} else if lang == "fr" {
			return "2c00ab61-d7d7-4ef9-b7d9-da0898ee0a6e"
		} else {
			return "a310cb1b-21ca-42f8-936c-4c0fe117a2f0"
		}
	} else if routeName == "about" {
		if lang == "es" {
			return "a85f6d25-63ff-4b88-af9a-24671923e884"
		} else if lang == "en" {
			return "3638b512-2c34-4d69-84dd-6668443af742"
		} else if lang == "fr" {
			return "9dd9e63a-cfe5-4806-8afe-fdbc102d612a"
		} else {
			return "a85f6d25-63ff-4b88-af9a-24671923e884"
		}
	} else if routeName == "courses" {
		if lang == "es" {
			return "ab0c5aa3-a761-4e19-a452-79bc02b4ccf3"
		} else if lang == "en" {
			return "7a6610c6-5aa6-4e44-8e9a-39cae82a79e8"
		} else if lang == "fr" {
			return "4b5ee0ba-9376-4e6d-aa21-14f28a74df2d"
		} else {
			return "ab0c5aa3-a761-4e19-a452-79bc02b4ccf3"
		}
	} else if routeName == "tours" {
		if lang == "es" {
			return "ffdbd06c-91b8-4e36-b3c8-26dff1aa39f6"
		} else if lang == "en" {
			return "501b3474-f7fb-4b7f-b110-309082d7a5f2"
		} else if lang == "fr" {
			return "2b702a78-e175-435a-9a49-a0a7f7d406d7"
		} else {
			return "ffdbd06c-91b8-4e36-b3c8-26dff1aa39f6"
		}
	} else if routeName == "coursesInfo" {
		if lang == "es" {
			return "b4d59541-0e6f-475d-a5ee-1c90b188f29f"
		} else if lang == "en" {
			return "2fd19aa7-97f9-4ef7-be32-1e4ed4f22ddb"
		} else if lang == "fr" {
			return "6b89154c-0221-4271-8aa1-9666368e8aed"
		} else {
			return "b4d59541-0e6f-475d-a5ee-1c90b188f29f"
		}
	} else if routeName == "toursInfo" {
		if lang == "es" {
			return "u8u79541-0e6f-475d-a5ee-1c90b18l9hdt"
		} else if lang == "en" {
			return "s4h79aa7-1458-4ef7-be32-1e4ed4f22ccc"
		} else if lang == "fr" {
			return "o09u154c-0221-4271-8aa1-9666368e8khe"
		} else {
			return "u8u79541-0e6f-475d-a5ee-1c90b18l9hdt"
		}
	} else {
		return ""
	}
}
