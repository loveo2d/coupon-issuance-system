package main

import (
	"log"
	"net/http"

	campaign_handler "github.com/loveo2d/CouponIssuanceSystem/internal/api/rpc/campaign"
	coupon_handler "github.com/loveo2d/CouponIssuanceSystem/internal/api/rpc/coupon"
)

func main() {
	campaignPath, campaignHandler := campaign_handler.New()
	couponPath, couponHandler := coupon_handler.New()

	mux := http.NewServeMux()
	mux.Handle(campaignPath, campaignHandler)
	mux.Handle(couponPath, couponHandler)

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	log.Println("Starting server on port 8000\nRouting services to:\n- ", campaignPath, "\n- ", couponPath)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
