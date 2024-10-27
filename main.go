package main

import "EduViewTeacher/networking"

func main() {
	sg := networking.SignalSender{}
	sg.ListenAndHandle()
}
