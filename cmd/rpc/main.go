package main

import (
	"log"
	"net/http"

	campaign_handler "github.com/loveo2d/CouponIssuanceSystem/internal/api/rpc/campaign"
	coupon_handler "github.com/loveo2d/CouponIssuanceSystem/internal/api/rpc/coupon"
	"github.com/loveo2d/CouponIssuanceSystem/internal/infra/db"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	campaignPath, campaignHandler := campaign_handler.New(db)
	couponPath, couponHandler := coupon_handler.New(db)

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
