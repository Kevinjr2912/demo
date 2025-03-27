package models

type DataIoT struct {
	IdPlot      int       `json:"id_parcel"`
	Temperature float32   `json:"temp"`
	AirQuality  float32   `json:"air"`
	Humidity    float32   `json:"humedity"`
	Sun         float32	  `json:"sun"`	
}