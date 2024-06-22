package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var MenuScheduleBotKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Получить расписание", "menu;schedule;Получить расписание",
		)),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Сменить бота", "menu;schedule;Сменить бота",
		),
		tgbotapi.NewInlineKeyboardButtonData(
			"Прочее", "menu;schedule;Прочее",
		),
	),
)

var MenuChoiceModeScheduleKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Сегодня", "menu;schedule;Сегодня",
		),
		tgbotapi.NewInlineKeyboardButtonData(
			"Завтра", "menu;schedule;Завтра",
		)),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"По дням", "menu;schedule;По дням",
		),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Назад", "menu;schedule;Назад",
		),
	),
)

var MenuScheduleOtherBotKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Посмотреть неделю", "menu;schedule;Посмотреть неделю",
		)),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Расписание звонков", "menu;schedule;Расписание звонков",
		)),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Назад", "menu;schedule;Назад",
		),
	),
)

var BuilderMenuTracker = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Добавить запись", "menu;tracker;Добавить запись",
		)),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Просмотр записей", "menu;tracker;Просмотр записей",
		)),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Назад", "menu;tracker;Назад",
		),
	),
)

var BuilderChoiceTrackerFilter = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Без фильтров", "menu;filter;Добавить запись",
		)),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Несколько фильтров", "menu;filter;Несколько фильтров",
		)),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Отфильтровать по олимпиаде", "menu;filter;Отфильтровать по олимпиаде",
		)),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Отфильтровать по предмету", "menu;filter;Отфильтровать по предмету",
		)),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Отфильтровать по этапу", "menu;filter;Отфильтровать по этапу",
		)),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Отфильтровать по наставнику", "menu;filter;Отфильтровать по наставнику",
		)),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Назад", "menu;filter;Назад",
		),
	),
)
