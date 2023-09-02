package handlers

type RateRequest struct {
	Base  string
	Quote string
}

type RateResponse struct {
	Rate float32
}
