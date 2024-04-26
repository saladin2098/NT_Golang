package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	Nmae   string
	Age    int
	Events []Event
}
type Event struct {
	Name string
	Date time.Time
}

func (user *User) deleteEvent(name string) (string, error) {
	for i := 0; i < len(user.Events); i++ {
		if user.Events[i].Name == name {
			user.Events = append(user.Events[:i], user.Events[i+1:]...)
			return "event deleted", nil
		}
	}
	return "Event not found", errors.New("Event not found")
}
func (usr *User) addEvent(ev Event) (string,error) {
	dif := ev.Date.Sub(time.Now()).Seconds()
	if dif >0 {
		usr.Events = append(usr.Events, ev)
		return "event added",nil
	} 
	return "event with unproper date",errors.New("not able to add event!")
}
func (user *User) modifyEvent(name string,new_date time.Time) (string,error) {
	for i:=0;i < len(user.Events);i++ {
		if user.Events[i].Name == name {
			user.Events[i].Date = new_date
			return "event modified",nil
		}
	}
	return "even not found",errors.New("event not found1")
}
func (usr *User) checkForDate() User {
	for i:=0;i<len(usr.Events); i++ {
		diff := usr.Events[i].Date.Sub(time.Now()).Seconds()
		if diff <0 {
			usr.Events = append(usr.Events[:i],usr.Events[i+1:]... )
		}
	}
	return *usr
}

func main() {
	// -------Homework 3.2---------
	// t1 := time.Now()
	// t2 := time.Date(2004, 10, 03, 16, 20, 20, 0, time.UTC)

	// res := t1.Sub(t2).Hours() / 24
	// fmt.Println(int(res))

	// hours := t1.Sub(t2).Hours()
	// fmt.Println(int(hours))

	// minutes := t1.Sub(t2).Minutes()
	// fmt.Println(int(minutes))

	// seconds := t1.Sub(t2).Minutes() * 60
	// fmt.Println(int(seconds))

	//----homework 3.3-----
	ev1 := Event{"Exam res", time.Date(2024, 04, 19, 9, 30, 0, 0, time.UTC)}
	evs := []Event{ev1}
	user1 := User{
		"Shamsiddin",
		19,
		evs,
	}

	ev2 := Event{"meeting",time.Date(2024,05,01,15,20,0,0,time.UTC)}
	// fmt.Println(user1)
	user1.addEvent(ev2)
	// fmt.Println(user1)

	sana := time.Date(2022,12,31,00,0,0,0,time.UTC)

	user1.modifyEvent("meeting",sana)
	// fmt.Println(user1)

	// user1.deleteEvent("meeting")
	// fmt.Println(user1)

	// ev3 := Event{"lesson",time.Date(2024,04,16,14,0,0,0,time.UTC)}
	// user1.addEvent(ev3)
	fmt.Println(user1.checkForDate())
        // my change
}
