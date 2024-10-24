package middlewares

import (
	"fmt"
	"gilangaryap/gym-buddy/pkg"
)

func Register(email string) (string, error) {
    emailSender := pkg.NewEmailSender()

    to := email
    subject := "Selamat Bergabung di Gym Buddy!"
    body := `Halo,

			Terima kasih telah mendaftar di Gym Buddy! Kami senang menyambut Anda di komunitas kami.
			
			Dengan akun ini, Anda dapat mengakses gym dan mendapatkan notifikasi jika langganan Anda habis.
			
			Pastikan untuk memverifikasi email Anda dan mulai perjalanan kebugaran Anda bersama kami. Jika Anda memiliki pertanyaan atau memerlukan bantuan, jangan ragu untuk menghubungi tim dukungan kami.
			
			Selamat berlatih!
			
			Salam hangat,
			Tim Gym Buddy
			`

    if err := emailSender.Send(to, subject, body); err != nil {
        fmt.Println("Error sending email:", err)
        return "", err
    } else {
        fmt.Println("Email sent successfully!")
    }
    return email, nil
}
