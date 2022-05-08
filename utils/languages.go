package utils


func GetLang(lang, routeName string)string{
	if routeName == "index"{
		if lang == "es"{
			return "4164da5e-cacd-4827-8891-26945019a5be"
		}else if lang == "en"{
			return "5da9ff07-3a11-4cc9-8d0a-f9a0e6daf925"
		}else if lang == "fr"{
			return "076d9ef9-bc59-40c7-ae7f-6a9e90cae856"
		}else {
			return "4164da5e-cacd-4827-8891-26945019a5be"
		}
	}else if routeName == "header"{
		if lang == "es"{
			return "76ff0c4a-ca1c-4d62-9304-6e3a71565ff4"
		}else if lang == "en"{
			return "a6b3948a-9f61-4c20-b1b6-e60628c70958"
		}else if lang == "fr"{
			return "f46c2a6c-54cb-4b41-b97f-7de4469b4df1"
		}else {
			return "76ff0c4a-ca1c-4d62-9304-6e3a71565ff4"
		}
	}else if routeName == "footer" {
		if lang == "es"{
			return "f93746d6-8b27-481f-ad1f-e888f7ef6d0f"
		}else if lang == "en"{
			return "4601db41-e3be-4f45-9809-5e3f67d43222"
		}else if lang == "fr"{
			return "f21aa619-b746-4b50-8381-de8956eba4df"
		}else {
			return "f93746d6-8b27-481f-ad1f-e888f7ef6d0f"
		}
	}else if routeName == "contact"{
		if lang == "es"{
			return "a310cb1b-21ca-42f8-936c-4c0fe117a2f0"
		}else if lang == "en"{
			return "5c240292-e280-4f49-8f6c-342295160271"
		}else if lang == "fr"{
			return "2c00ab61-d7d7-4ef9-b7d9-da0898ee0a6e"
		}else {
			return "a310cb1b-21ca-42f8-936c-4c0fe117a2f0"
		}
	}else if routeName == "about"{
		if lang == "es"{
			return "a85f6d25-63ff-4b88-af9a-24671923e884"
		}else if lang == "en"{
			return "3638b512-2c34-4d69-84dd-6668443af742"
		}else if lang == "fr"{
			return "9dd9e63a-cfe5-4806-8afe-fdbc102d612a"
		}else {
			return "a85f6d25-63ff-4b88-af9a-24671923e884"
		}
	}else {
		return ""
	}
}
