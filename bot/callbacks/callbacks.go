package callbacks

import (
	"awesomeProject/bot/lexicon"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mymmrac/telego"
)
var BuilderChoiceBot = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Бот-расписание", "main;choice;bot-schedule"),
		tgbotapi.NewInlineKeyboardButtonData("Бот-трекер", "main;choice;bot-treker"),
	),
)

var BuilderChoiceBotWithMiniApp = telego.InlineKeyboardMarkup{
	InlineKeyboard: [][]telego.InlineKeyboardButton{
		{
			telego.InlineKeyboardButton{
				Text: "Бот-расписание",
				CallbackData: "main;choice;bot-schedule",
			},
			telego.InlineKeyboardButton{
				Text: "Бот-трекер",
				CallbackData: "main;choice;bot-treker",
			},
		},
		{
			telego.InlineKeyboardButton{
				Text: "Mini-App",
				WebApp: &telego.WebAppInfo{
					URL: fmt.Sprintf("%s/home", lexicon.Link),
				},
			},
		},
	},
}


// func init() {
// 	BuilderChoiceBotWithMiniApp = telego.InlineKeyboardMarkup{
// 		InlineKeyboard: [][]telego.InlineKeyboardButton{
// 			{
// 				telego.InlineKeyboardButton{
// 					Text: "Бот-расписание",
// 					CallbackData: "main;choice;bot-schedule",
// 				},
// 				telego.InlineKeyboardButton{
// 					Text: "Бот-трекер",
// 					CallbackData: "main;choice;bot-treker",
// 				},
// 			},
// 			{
// 				telego.InlineKeyboardButton{
// 					Text: "Mini-App",
// 					WebApp: &telego.WebAppInfo{
// 						URL: fmt.Sprintf("%s/home", lexicon.Link),
// 					},
// 				},
// 			},
// 		},
// 	}
// }
