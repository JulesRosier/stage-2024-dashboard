package handler

import (
	"Stage-2024-dashboard/pkg/demo"
	"Stage-2024-dashboard/pkg/view"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo/v4"
)

func (h *Handler) DemoHome(c echo.Context) error {
	return render(c, view.DemoHome())
}

func (h *Handler) DemoStep1(c echo.Context) error {
	return render(c, view.DemoStep1())
}

func (h *Handler) DemoStep1Post(c echo.Context) error {
	username := c.FormValue("name")
	username = strings.TrimSpace(username)

	userId := gofakeit.UUID()

	user := demo.User{
		Name: username,
		Id:   userId,
	}
	userB, err := json.Marshal(user)
	if err != nil {
		return err
	}

	ex := time.Now().Add(24 * time.Hour)
	c.SetCookie(&http.Cookie{
		Name:    "demo_user",
		Value:   base64.StdEncoding.EncodeToString(userB),
		Expires: ex,
	})

	demo.ProduceCreateUser(c.Request().Context(), user)
	return render(c, view.DemoStep2())
}

func (h *Handler) DemoStep2Post(c echo.Context) error {
	user, err := getUser(c)
	if err != nil {
		return err
	}

	bike := c.FormValue("bike")
	bike = strings.TrimSpace(bike)

	bikeS := demo.Bike{
		Company: bike,
		Id:      gofakeit.UUID(),
	}
	bikeB, err := json.Marshal(bikeS)
	if err != nil {
		return err
	}

	ex := time.Now().Add(24 * time.Hour)
	c.SetCookie(&http.Cookie{
		Name:    "demo_bike",
		Value:   base64.StdEncoding.EncodeToString(bikeB),
		Expires: ex,
	})

	err = demo.ProduceBikePickedUp(c.Request().Context(), user, bikeS)
	if err != nil {
		return err
	}

	return render(c, view.DemoStep3())
}

func (h *Handler) DemoStep3Post(c echo.Context) error {
	user, err := getUser(c)
	if err != nil {
		return err
	}

	bike, err := getBike(c)
	if err != nil {
		return err
	}

	err = demo.ProduceBikeReturned(c.Request().Context(), user, bike)
	if err != nil {
		return err
	}

	return render(c, view.DemoStep1())
}

func getUser(c echo.Context) (demo.User, error) {
	userC, err := c.Cookie("demo_user")
	if err != nil {
		return demo.User{}, err
	}
	userB, err := base64.StdEncoding.DecodeString(userC.Value)
	if err != nil {
		return demo.User{}, err
	}

	var user demo.User
	err = json.Unmarshal([]byte(userB), &user)
	if err != nil {
		return demo.User{}, err
	}

	return user, nil
}

func getBike(c echo.Context) (demo.Bike, error) {
	bikeC, err := c.Cookie("demo_bike")
	if err != nil {
		return demo.Bike{}, err
	}
	bikeB, err := base64.StdEncoding.DecodeString(bikeC.Value)
	if err != nil {
		return demo.Bike{}, err
	}

	var bike demo.Bike
	err = json.Unmarshal([]byte(bikeB), &bike)
	if err != nil {
		return demo.Bike{}, err
	}

	return bike, nil
}
