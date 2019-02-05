package movies

import(

)


type Movie struct{
	Name string `json:"name"`
	Year string `json:"year"`
	Nacionality string `json:"nacionality"`
	Review string `json:"review,omitempty"`
}