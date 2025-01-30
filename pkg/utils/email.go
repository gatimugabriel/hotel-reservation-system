package utils

import (
	"fmt"
	"net/smtp"
	"os"
	"time"
)

type ReservationEmailData struct {
	ID            string    `json:"id"`
	CheckInDate   time.Time `json:"check_in_date"`
	CheckOutDate  time.Time `json:"check_out_date"`
	RoomNumber    int       `json:"room_number"`
	RoomType      string    `json:"room_type"`
	HotelName     string    `json:"hotel_name"`
	GuestName     string    `json:"guest_name"`
	TotalPrice    float64   `json:"total_price"`
	PaymentStatus string    `json:"payment_status"`

	GuestEmail     string `json:"guest_email"`
	SpecialRequest string `json:"special_request,omitempty"`
	NumGuests      int    `json:"num_guests"`
	Currency       string `json:"currency"`
	PaymentMethod  string `json:"payment_method"`
}

func formatDate(date time.Time) string {
	return date.Format("Monday, January 2, 2006")
}

func SendEmailNotification(email string, reservationData ReservationEmailData) error {
	from := os.Getenv("EMAIL_SENDER")
	pass := os.Getenv("EMAIL_PASSWORD")
	ApiUrl := os.Getenv("SERVER_BASE_URL")
	to := email

	reservationDetailsUrl := fmt.Sprintf("%s/reservation/reservation-details/%s", ApiUrl, reservationData.ID)

	htmlTemplate := `
    <!DOCTYPE html>
    <html>
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
    </head>
    <body style="margin: 0; padding: 0; font-family: Arial, sans-serif; background-color: #f4f4f4;">
        <div style="max-width: 600px; margin: 0 auto; background-color: #ffffff; padding: 20px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
            <div style="text-align: center; padding: 20px 0; border-bottom: 2px solid #f0f0f0;">
                <h1 style="color: #2e6c80; margin: 0;">Booking Confirmation</h1>
            </div>
            
            <div style="padding: 20px 0;">
                <p style="font-size: 16px; color: #333;">Dear %s,</p>
                <p style="font-size: 16px; color: #333;">Thank you for choosing %s. Your reservation has been confirmed!</p>
                
                <div style="background-color: #f8f9fa; padding: 20px; border-radius: 4px; margin: 20px 0;">
                    <h2 style="color: #2e6c80; font-size: 18px; margin-top: 0;">Reservation Details</h2>
                    <table style="width: 100%%; border-collapse: collapse;">
                        <tr>
                            <td style="padding: 8px 0; color: #666;">Confirmation Number:</td>
                            <td style="padding: 8px 0; color: #333; font-weight: bold;">%s</td>
                        </tr>
                        <tr>
                            <td style="padding: 8px 0; color: #666;">Check-in Date:</td>
                            <td style="padding: 8px 0; color: #333;">%s</td>
                        </tr>
                        <tr>
                            <td style="padding: 8px 0; color: #666;">Check-out Date:</td>
                            <td style="padding: 8px 0; color: #333;">%s</td>
                        </tr>
                        <tr>
                            <td style="padding: 8px 0; color: #666;">Room Type:</td>
                            <td style="padding: 8px 0; color: #333;">%s</td>
                        </tr>
                        <tr>
                            <td style="padding: 8px 0; color: #666;">Room Number:</td>
                            <td style="padding: 8px 0; color: #333;">%d</td>
                        </tr>
                        <tr>
                            <td style="padding: 8px 0; color: #666;">Total Price:</td>
                            <td style="padding: 8px 0; color: #333;">$%.2f</td>
                        </tr>
                        <tr>
                            <td style="padding: 8px 0; color: #666;">Payment Status:</td>
                            <td style="padding: 8px 0; color: %s; font-weight: bold;">%s</td>
                        </tr>
                    </table>
                </div>
                
                <div style="text-align: center; margin: 30px 0;">
                    <a href="%s" style="background-color: #2e6c80; color: white; padding: 12px 24px; text-decoration: none; border-radius: 4px; display: inline-block;">View Reservation Details</a>
                </div>

                <div style="background-color: #fff3cd; padding: 15px; border-radius: 4px; margin: 20px 0;">
                    <p style="color: #856404; margin: 0;">
                        <strong>Important:</strong> Please keep this email for your records. You may need to present it upon check-in.
                    </p>
                </div>
            </div>
            
            <div style="border-top: 2px solid #f0f0f0; padding-top: 20px; margin-top: 20px;">
                <p style="margin: 0; color: #666; font-size: 14px;">Best Regards,</p>
                <p style="margin: 5px 0; color: #2e6c80; font-weight: bold;">%s Team</p>
            </div>
            
            <div style="text-align: center; margin-top: 30px; padding-top: 20px; border-top: 1px solid #f0f0f0;">
                <p style="color: #999; font-size: 12px;">If you have any questions, please contact us at:</p>
                <p style="color: #666; font-size: 14px;">📞 Contact: <a href="tel:%s" style="color: #2e6c80; text-decoration: none;">%s</a></p>
                <p style="color: #666; font-size: 14px;">✉️ Email: <a href="mailto:%s" style="color: #2e6c80; text-decoration: none;">%s</a></p>
            </div>
        </div>
    </body>
    </html>
    `

	paymentStatusColor := "#28a745" //  success
	if reservationData.PaymentStatus != "COMPLETED" {
		paymentStatusColor = "#ffc107" // pending
	}

	msg := fmt.Sprintf("From: %s\n"+
		"To: %s\n"+
		"Subject: Booking Confirmation - %s\n"+
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"+
		htmlTemplate,
		from,
		to,
		reservationData.HotelName,
		reservationData.GuestName,
		reservationData.HotelName,
		reservationData.ID,
		formatDate(reservationData.CheckInDate),
		formatDate(reservationData.CheckOutDate),
		reservationData.RoomType,
		reservationData.RoomNumber,
		reservationData.TotalPrice,
		paymentStatusColor,
		reservationData.PaymentStatus,
		reservationDetailsUrl,
		reservationData.HotelName,
		os.Getenv("HOTEL_CONTACT"),
		os.Getenv("HOTEL_CONTACT"),
		os.Getenv("HOTEL_EMAIL"),
		os.Getenv("HOTEL_EMAIL"))

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from,
		[]string{to},
		[]byte(msg))

	if err != nil {
		return fmt.Errorf("smtp error: %s", err)
	}

	return nil
}