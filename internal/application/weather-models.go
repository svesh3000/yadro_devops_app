package application

type weatherDay struct {
	Temp    float64 `json:"temp"`
	TempMin float64 `json:"tempmin"`
	TempMax float64 `json:"tempmax"`
}

type weatherAPIResponse struct {
	Days []weatherDay `json:"days"`
}

type weatherResponse struct {
	Service string `json:"service"`
	Data    struct {
		TemperatureC struct {
			Average float64 `json:"average"`
			Median  float64 `json:"median"`
			Min     float64 `json:"min"`
			Max     float64 `json:"max"`
		} `json:"temperature_c"`
	} `json:"data"`
}
