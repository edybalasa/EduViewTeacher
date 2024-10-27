package app

import (
	"EduViewTeacher/SQLite3"
	"EduViewTeacher/networking"
)

func StartApplication() {
	sg := networking.SignalSender{}
	SQLite3.HandleHostname(sg.ListenAndHandle())
}
