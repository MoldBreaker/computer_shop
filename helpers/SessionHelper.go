package helpers

import (
	"computer_shop/models"
	"encoding/gob"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func initSession() {
	gob.Register(models.UserModel{})
}

func SetSession(key string, value interface{}, c echo.Context) error {
	initSession()
	var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	session, err := store.Get(c.Request(), "session")
	if err != nil {
		return fmt.Errorf("error when getting session store: %w", err)
	}
	session.Values[key] = value
	errSave := session.Save(c.Request(), c.Response())
	if errSave != nil {
		return fmt.Errorf("error when saving session: %w", errSave)
	}
	return nil
}

func GetSession(key string, c echo.Context) (interface{}, error) {
	initSession()
	var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	session, err := store.Get(c.Request(), "session")
	if err != nil {
		return nil, fmt.Errorf("error when getting session store: %w", err)
	}
	value := session.Values[key]
	return value, nil
}

func RemoveSession(key string, c echo.Context) error {
	//initSession()
	//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	//session, err := store.Get(c.Request(), "session")
	//if err != nil {
	//	return fmt.Errorf("error when getting session store: %w", err)
	//}
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(c.Response(), cookie)
	//delete(session.Values, key)
	//errSave := session.Save(c.Request(), c.Response())
	//if errSave != nil {
	//	return fmt.Errorf("error when saving session: %w", errSave)
	//}
	return nil
}
