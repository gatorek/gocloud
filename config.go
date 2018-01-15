package main

func config(key string) string {
	switch key {
		case "dir":
			return "download/"
		case "fastrategy":
			return "local"
	}
	return ""
}