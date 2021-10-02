package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"git.mrcyjanek.net/mrcyjanek/mzetczanspozniasu/mzklib"
	"github.com/google/shlex"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/format"
	"maunium.net/go/mautrix/id"
)

var (
	homeserver      = flag.String("homeserver", "https://halogen.city", "Matrix homeserver URL")
	homeserver_name = flag.String("homeservername", "halogen.city", "Matrix homeserver name")
	username        = flag.String("username", "mzetczan", "Bot username")
	accesstoken     = flag.String("accesstoken", "", "Matrix accesstoken")
)

//go:embed help.txt
var helptxt string

func init() {
	helptxt = strings.ReplaceAll(helptxt, "\n", "")
}

func main() {
	flag.Parse()
	client, err := mautrix.NewClient(*homeserver, id.NewUserID(*username, *homeserver_name), *accesstoken)
	if err != nil {
		log.Fatal(err)
	}
	syncer := client.Syncer.(*mautrix.DefaultSyncer)
	syncer.OnEventType(event.EventMessage, func(source mautrix.EventSource, evt *event.Event) {
		if !strings.HasPrefix(evt.Content.AsMessage().Body, "!mzk") {
			return
		}
		if -60*1000 > (evt.Timestamp - time.Now().Unix()*1000) {
			log.Println("Ignoring...")
			return
		}
		fmt.Printf("<%[1]s> %[4]s (%[2]s/%[3]s)\n", evt.Sender, evt.Type.String(), evt.ID, evt.Content.AsMessage().Body)

		cmd, err := shlex.Split(evt.Content.AsMessage().Body)
		if err != nil {
			client.SendText(evt.RoomID, err.Error())
			return
		}
		if len(cmd) < 2 {
			client.SendText(evt.RoomID, "Potrzeba conajmniej 2 argumentów, zobacz !mzk help")
			return
		}
		msgtosend := "??? Nieznana komenda, zobacz <code>!mzk help</code>"
		switch cmd[1] {
		case "help", "h", "?":
			msgtosend = helptxt
		case "busy", "b":
			linie := mzklib.GetLines()
			msgtosend = "<b>Lista autobusów</b>\n"
			for i := range linie.Lines {
				if i != 0 {
					msgtosend += ", "
				}
				msgtosend += linie.Lines[i].Name
			}
		case "przystanek", "p":
			if len(cmd) != 4 {
				client.SendText(evt.RoomID, "Potrzeba conajmniej 4 argumentów, zobacz !mzk help")
				return
			}
			//linia := mzklib.Linie{}
			linie := mzklib.GetLines()
			found := false
			for i := range linie.Lines {
				if linie.Lines[i].Name == cmd[2] {
					//linia = linie.Lines[i]
					found = true
					break
				}
			}
			if !found {
				client.SendText(evt.RoomID, "Nie udało mi się znaleźć tego autobusu D:")
				return
			}
			k := mzklib.GetDirectionS(cmd[2], cmd[3])
			msgtosend = "<b>Przystanki:</b>"
			for i := range k.StopPoints {
				msgtosend += "\n " + strconv.Itoa(i) + ". " + k.StopPoints[i].Name
			}
			// log.Println("Przystanek:", kierunek.Connections[id].FromSymbol, "->", kierunek.Connections[id].ToSymbol)
			// log.Println(" - Nazwa:", mzklib.GetStop(kierunek.Connections[id].FromSymbol).Name, "->", mzklib.GetStop(kierunek.Connections[id].ToSymbol).Name)
		case "rozklad", "rozkład", "r":
			if len(cmd) != 5 {
				client.SendText(evt.RoomID, "Potrzeba conajmniej 5 argumentów, zobacz !mzk help")
				return
			}
			k := mzklib.GetDirectionS(cmd[2], cmd[3])
			msgtosend = "<b>Rozkład jazdy, nr. " + cmd[2] + "</b>"
			for i := range k.StopPoints {
				if strconv.Itoa(i) == cmd[4] {
					sched := mzklib.GetAtomicSchedule(k.Connections[i].ToSymbol)
					for i := range sched.LineSchedules {
						if sched.LineSchedules[i].LineName == cmd[2] {
							for j := range sched.LineSchedules[i].Departures {
								k := sched.LineSchedules[i].Departures[j]
								msgtosend += "\n - <b>" + k.ScheduledDepartureString + "</b>  " + k.OptionalDirection
							}
						}
					}
				}
			}
		}
		client.SendMessageEvent(evt.RoomID, event.EventMessage, format.RenderMarkdown(msgtosend, false, true))
	})
	log.Println("Syncing")
	err = client.Sync()
	if err != nil {
		log.Fatal(err)
	}

	linia := mzklib.GetLines().Lines[5]
	log.Println("ID:", linia.ID, "NAME:", linia.Name, "LineType:", linia.LineType)
	kierunek := mzklib.GetDirection(int(linia.ID), true)
	//log.Fatal(len(kierunek.Connections))
	id := 2
	log.Println("Przystanek:", kierunek.Connections[id].FromSymbol, "->", kierunek.Connections[id].ToSymbol)
	log.Println(" - Nazwa:", mzklib.GetStop(kierunek.Connections[id].FromSymbol).Name, "->", mzklib.GetStop(kierunek.Connections[id].ToSymbol).Name)
	sched := mzklib.GetAtomicSchedule(kierunek.Connections[id].ToSymbol)
	log.Println("Busy:")
	for i := range sched.LineSchedules {
		if sched.LineSchedules[i].LineName == linia.Name {
			for j := range sched.LineSchedules[i].Departures {
				k := sched.LineSchedules[i].Departures[j]
				log.Println(" -", k.ScheduledDepartureString, k.OptionalDirection)
			}
		}
	}
}
