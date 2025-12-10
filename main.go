package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
)

// ===== Handlers =====

func onStart(c telebot.Context) error {
	name := c.Sender().FirstName
	if name == "" {
		name = "друг"
	}
	return c.Send("👋 Привет, " + name + "!\nЯ простой Telegram-бот на Go.\nНабери /help чтобы увидеть, что я умею.")
}

func onHelp(c telebot.Context) error {
	help := `ℹ️ Доступные команды:
/start - приветствие
/help  - помощь
/echo <текст> - повторю твой текст
/id    - покажу твой Telegram ID
/time  - текущее время сервера
/menu  - показать кнопки

Также я реагирую на: текст, фото, документ, стикер.`
	return c.Send(help)
}

func onEcho(c telebot.Context) error {
	args := c.Args()
	if len(args) == 0 {
		return c.Send("Использование: /echo <текст>")
	}
	return c.Send(strings.Join(args, " "))
}

func onID(c telebot.Context) error {
	uid := c.Sender().ID
	return c.Send(fmt.Sprintf("🆔 Твой ID: %d", uid))
}

func onTime(c telebot.Context) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	return c.Send("⏰ Время сервера: " + now)
}

func onText(c telebot.Context) error {
	return c.Send("Ты написал: " + c.Text())
}

func onPhoto(c telebot.Context) error {
	p := c.Message().Photo
	if p == nil {
		return c.Send("Фото не найдено 🤔")
	}
	return c.Send(fmt.Sprintf("📷 Класс! Получил фото (%dx%d)", p.Width, p.Height))
}

func onDocument(c telebot.Context) error {
	doc := c.Message().Document
	if doc == nil {
		return c.Send("Документ не найден 🤔")
	}
	return c.Send("📄 Спасибо! Файл: " + doc.FileName)
}

func onSticker(c telebot.Context) error {
	return c.Send("😄 Классный стикер!")
}

func onMenu(c telebot.Context) error {
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	btnID := menu.Text("Мой ID")
	btnTime := menu.Text("Время")
	btnHelp := menu.Text("Help")
	menu.Reply(menu.Row(btnID, btnTime, btnHelp))

	if err := c.Send("Выбери действие:", menu); err != nil {
		return err
	}
	return nil
}

func onButtons(c telebot.Context) error {
	switch c.Text() {
	case "Мой ID":
		return onID(c)
	case "Время":
		return onTime(c)
	case "Help":
		return onHelp(c)
	default:
		return onText(c)
	}
}

// ===== Main =====

func main() {
	// 1) Пытаемся прочитать TELE_TOKEN из окружения (для Docker/K8s/прода)
	token := strings.TrimSpace(os.Getenv("TELE_TOKEN"))

	// 2) Если пусто — пробуем загрузить bot.env (для локальной разработки)
	if token == "" {
		if err := godotenv.Load("bot.env"); err != nil {
			log.Println("⚠️ bot.env не найден, TELE_TOKEN ищу только в переменных окружения")
		}
		token = strings.TrimSpace(os.Getenv("TELE_TOKEN"))
	}

	if token == "" {
		log.Fatal("❌ TELE_TOKEN пуст. Задай переменную окружения или создай bot.env")
	}

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal("❌ Ошибка при создании бота: ", err)
	}

	// Команды
	bot.Handle("/start", onStart)
	bot.Handle("/help", onHelp)
	bot.Handle("/echo", onEcho)
	bot.Handle("/id", onID)
	bot.Handle("/time", onTime)
	bot.Handle("/menu", onMenu)

	// Типы сообщений
	bot.Handle(telebot.OnText, onButtons)
	bot.Handle(telebot.OnPhoto, onPhoto)
	bot.Handle(telebot.OnDocument, onDocument)
	bot.Handle(telebot.OnSticker, onSticker)

	fmt.Println("✅ Бот запущен! Нажми Ctrl + C чтобы остановить.")
	bot.Start()
}

// TEST_SECRET: 123456:AAAbbbCCCdddEEEfffFakeTelegramToken
