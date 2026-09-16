package main

import (
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type About struct {
	Skills     []string
	Educations []string
	GPA        []string
	Projects   []string
	Info       []string
}

func main() {
	_ = godotenv.Load("token.env")
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is not set")
	}
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		webhookURL = os.Getenv("RENDER_EXTERNAL_URL")
		if webhookURL == "" {
			if host := os.Getenv("RENDER_EXTERNAL_HOSTNAME"); host != "" {
				webhookURL = "https://" + host
			}
		}
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "10000"
	}
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Fatal("Error creating bot:", err)
	}

	log.Println("Authorized on account:", bot.Self.UserName)
	about := About{
		Skills: []string{
			"Python",
			"Golang",
			"Mysql",
			"Sql server",
			"c++",
		},

		Educations: []string{
			"B.Sc . computer engineering at Azad university of tabriz",
		},

		GPA: []string{
			"bachelor(untill now):16.31",
			"Diploma: 17.62",
		},

		Projects: []string{
			"pharmacy manegment system : used language: c++",
			"Taskmanager : used language: Golang",
			"Student managment real life database : used skill : Mysql",
		},

		Info: []string{
			"Email: amirsepehr265@gmail.com",
			"--",
			"Whatsapp: +989148884564",
		},
	}
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Skills"),
			tgbotapi.NewKeyboardButton("Educations"),
		),

		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Gpa"),
			tgbotapi.NewKeyboardButton("Projects"),
		),

		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Contact"),
		),
	)
	webhookPath := "/" + token

	if webhookURL != "" {
		webhook, err := tgbotapi.NewWebhook(webhookURL + webhookPath)
		if err != nil {
			log.Fatal("error in creating", err)
		}
		_, err = bot.Request(webhook)
		if err != nil {
			log.Fatal("error in setting", err)
		}
		log.Println("webhook setsuccesfully:")
		log.Println(webhookURL + webhookPath)
	} else {
		log.Println("webhook is not set yet :")
	}

	info, err := bot.GetWebhookInfo()

	if err != nil {
		log.Println("Could not get webhook information:", err)
	} else {
		log.Println("Telegram Webhook URL:", info.URL)

		if info.LastErrorDate != 0 {
			log.Println("Telegram Webhook Error:", info.LastErrorMessage)
		}
	}
	updates := bot.ListenForWebhook(webhookPath)

	go func() {

		log.Println("HTTP server running on port:", port)

		err := http.ListenAndServe(
			"0.0.0.0:"+port,
			nil,
		)

		if err != nil {
			log.Fatal("HTTP server error:", err)
		}
	}()

	for update := range updates {

		if update.Message == nil {
			continue
		}

		log.Printf(
			"[%s] %s",
			update.Message.From.UserName,
			update.Message.Text,
		)

		var response string

		switch update.Message.Text {

		case "/start":

			response = "Welcome to my portfolio bot! 👋"

		case "Skills":

			response = "Skills:\n"

			for _, skill := range about.Skills {
				response += "• " + skill + "\n"
			}

		case "Educations":

			response = "Education:\n"

			for _, education := range about.Educations {
				response += "• " + education + "\n"
			}

		case "Gpa":

			response = "GPA:\n"

			for _, gpa := range about.GPA {
				response += "• " + gpa + "\n"
			}

		case "Projects":

			response = "Projects:\n"

			for _, project := range about.Projects {
				response += "• " + project + "\n"
			}

		case "Contact":

			response = "Contact Information:\n"

			for _, info := range about.Info {
				response += info + "\n"
			}

		default:

			response = "Please choose one of the options below."

		}

		message := tgbotapi.NewMessage(
			update.Message.Chat.ID,
			response,
		)

		message.ReplyMarkup = keyboard

		_, err := bot.Send(message)

		if err != nil {
			log.Println("Error sending message:", err)
		}
	}
}
