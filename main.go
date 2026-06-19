package main

import (
	"log"
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
	err := godotenv.Load("token.env")
	if err != nil {
		log.Fatal("error happenne")
	}
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is not set")
	}

	// insertin datas in struct
	Data := About{
		Skills:     []string{"Python\t", "Golang\n", "Mysql\t", "Sql server", "c++\n"}, //skil haro bishtar kon to switch
		Educations: []string{"B.Sc . computer engineering at Azad university of tabriz\n"},
		GPA:        []string{"bachelor(untill now):16.31", "Diploma: 17.62"},
		Projects:   []string{"pharmacy manegment system : used language: c++\n", "Taskmanager : used language: Golang\n", "Student managment real life database : used skill : Mysql\n"},
		Info:       []string{"Email: amirsepehr265@gmail.com", " -- ", "Whatsapp: +989148884564"},
	}

	keboard := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Skills")), tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Educations")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Gpa")), tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Projects")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Contact")))
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("authorized on", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		message := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		message.ReplyMarkup = keboard
		switch update.Message.Text {
		case "Skills":
			responseText := "🧤 Skills that i have currently 🧤\n"
			for _, skills := range Data.Skills {
				responseText += "-" + skills
			}
			message.Text = responseText
		case "Gpa":
			responseText := "◻ All my Grades currently ◻:"
			for _, gpa := range Data.GPA {
				responseText += "--" + gpa
			}
			message.Text = responseText
		case "Contact":
			responseText := "☎ Contact info ☎:"
			for _, contact := range Data.Info {
				responseText += " - " + contact
			}
			message.Text = responseText
		case "Educations":
			responseText := " 👨‍🎓 My degrees 👨‍🎓 :"
			for _, edu := range Data.Educations {
				responseText += " -- " + edu
			}
			message.Text = responseText
		case "Projects":
			responseText := " 👨‍💻 Here are the projects that i've done :"
			for _, pro := range Data.Projects {
				responseText += "-" + pro
			}
			message.Text = responseText
		default:
			message.Text = "choose one of the options below "
		}
		if _, err := bot.Send(message); err != nil {
			log.Println("error in sending !!!", err)
		}

	}
}
