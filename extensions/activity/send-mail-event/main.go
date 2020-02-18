package sendmailevent

import (
	"fmt"
	"time"
)


func main(){
	fmt.Println("Building ICS test")
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now().UTC().Format("20060102T150405Z"))
	
	/*c := ical.New()
	c.AddProperty("X-Foo-Bar-Baz", "value")
	tz := ical.NewTimezone()
	tz.AddProperty("TZID", "Europe/Berlin")
	tz.AddProperty("SUMMARY", "Titulo")
	tz.AddProperty("DESCRIPTION", "body")
	c.AddEntry(tz)
	e := ical.Event{}

	ical.NewEncoder(os.Stdout).Encode(c)*/
}
