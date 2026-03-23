package application

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"time"
)

const dateFormat = "2006-01-02"

func parseDates(dateFromStr, dateToStr string) (from, to time.Time, status int, err error) {
	switch {
	case dateFromStr == "" && dateToStr == "":
		today := time.Now()

		return today, today.AddDate(0, 0, 1), http.StatusOK, nil

	case dateFromStr == "" && dateToStr != "":
		return time.Time{}, time.Time{}, http.StatusBadRequest,
			fmt.Errorf("date_from is required when date_to is provided")

	case dateFromStr != "" && dateToStr == "":
		return time.Time{}, time.Time{}, http.StatusBadRequest,
			fmt.Errorf("date_to is required when date_from is provided")
	}

	from, err = time.Parse(dateFormat, dateFromStr)
	if err != nil {
		return time.Time{}, time.Time{}, http.StatusBadRequest,
			fmt.Errorf("invalid date_from format, expected YYYY-MM-DD")
	}

	to, err = time.Parse(dateFormat, dateToStr)
	if err != nil {
		return time.Time{}, time.Time{}, http.StatusBadRequest,
			fmt.Errorf("invalid date_to format, expected YYYY-MM-DD")
	}

	if from.After(to) {
		return time.Time{}, time.Time{}, http.StatusBadRequest,
			fmt.Errorf("date_from cannot be after date_to")
	}

	return from, to, http.StatusOK, nil
}

func buildWeatherURL(city, from, to, apiKey string) string {
	encodedCity := url.QueryEscape(city)
	base := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline"
	path := fmt.Sprintf("%s/%s/%s", encodedCity, from, to)
	query := fmt.Sprintf("unitGroup=metric&include=days&key=%s", apiKey)

	return fmt.Sprintf("%s/%s?%s", base, path, query)
}

func fetchWeatherData(url string) (*weatherAPIResponse, int, error) {
	client := &http.Client{Timeout: 100 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to call weather API: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))

		return nil, resp.StatusCode,
			fmt.Errorf("weather API returned error (status %d): %s", resp.StatusCode, string(body))
	}

	var apiResp weatherAPIResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&apiResp); err != nil {
		return nil, http.StatusInternalServerError,
			fmt.Errorf("failed to parse weather API response: %v", err)
	}

	return &apiResp, http.StatusOK, nil
}

func parseWeatherResponse(apiResp *weatherAPIResponse) (temps, tempsMin,
	tempsMax []float64, status int, err error) {
	if len(apiResp.Days) == 0 {
		return nil, nil, nil, http.StatusNotFound,
			fmt.Errorf("no weather data for given dates")
	}

	temps = make([]float64, len(apiResp.Days))
	tempsMin = make([]float64, len(apiResp.Days))
	tempsMax = make([]float64, len(apiResp.Days))

	for i, day := range apiResp.Days {
		temps[i] = day.Temp
		tempsMin[i] = day.TempMin
		tempsMax[i] = day.TempMax
	}

	return temps, tempsMin, tempsMax, http.StatusOK, nil
}

func (app *App) WeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)

		return
	}

	dateFromStr := r.URL.Query().Get("date_from")
	dateToStr := r.URL.Query().Get("date_to")

	fromDate, toDate, status, err := parseDates(dateFromStr, dateToStr)
	if err != nil {
		http.Error(w, err.Error(), status)

		return
	}

	from := fromDate.Format(dateFormat)
	to := toDate.Format(dateFormat)

	apiURL := buildWeatherURL(city, from, to, app.Config.ApiKey)

	apiResp, status, err := fetchWeatherData(apiURL)
	if err != nil {
		http.Error(w, err.Error(), status)

		return
	}

	temps, tempsMin, tempsMax, status, err := parseWeatherResponse(apiResp)
	if err != nil {
		http.Error(w, err.Error(), status)

		return
	}

	avg := calcAverage(temps)
	med := calcMedian(temps)
	min := slices.Min(tempsMin)
	max := slices.Max(tempsMax)

	response := weatherResponse{Service: "weather"}
	response.Data.TemperatureC.Average = avg
	response.Data.TemperatureC.Median = med
	response.Data.TemperatureC.Min = min
	response.Data.TemperatureC.Max = max

	writeJSONResponse(w, response)
}
