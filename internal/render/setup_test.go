package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/thiruthanikaiarasu/udemy-go/bookings/internal/config"
	"github.com/thiruthanikaiarasu/udemy-go/bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

// TestMain runs before the test file and end it execute the test file 
func TestMain(m *testing.M) {

	//What am I going to to put in the session 
	gob.Register(models.Reservation{})

	// change this to true when in production 
	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour 
	session.Cookie.Persist = true 
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (mh *myWriter) Header() http.Header {
	var h http.Header

	return h
}

func (mh *myWriter) WriteHeader(i int) {

}

func (mh *myWriter) Write(b []byte) (int, error) {
	length := len(b)

	return length, nil 
}