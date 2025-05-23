package renderer

import (
	"fmt"

	"github.com/needydetain/gust/internal/config"
	"github.com/needydetain/gust/internal/models"
)

func (r *TerminalRenderer) RenderFullWeather(city *models.City, weather *models.OneCallResponse, cfg *config.Config) {
	r.RenderCurrentWeather(city, weather, cfg)
	fmt.Println()

	if len(weather.Alerts) > 0 {
		r.RenderAlerts(city, weather, cfg)
		fmt.Println()
	}

	r.RenderDailyForecast(city, weather, cfg)
	fmt.Println()
}
