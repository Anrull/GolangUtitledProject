package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var ProfilePanel = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔴 Настроить цвет расписания 🟢", "profile;choice-color;main"),
	),
)

var ProfileColorsPanel = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🟢", "profile;choice-color;green"),
		tgbotapi.NewInlineKeyboardButtonData("🔵", "profile;choice-color;blue"),
		tgbotapi.NewInlineKeyboardButtonData("⚪️", "profile;choice-color;white"),
		tgbotapi.NewInlineKeyboardButtonData("🟡", "profile;choice-color;yellow"),
		tgbotapi.NewInlineKeyboardButtonData("🟣", "profile;choice-color;purple"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "profile;escape")),
)

var ProfilePanelEscape = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "profile;escape")),
)

var AdminPanel = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Database", "admin;get_db"),
		tgbotapi.NewInlineKeyboardButtonData("Logs", "admin;get_logs"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Кол-во юзеров", "admin;count"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Оценка уроков", "admin;fb")),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Прочее", "admin;other")),
)

var AdminFB = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Эта неделя", "admin;fbHandler;nowWeek"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Прошлая неделя", "admin;fbHandler;lastWeek"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Этот месяц", "admin;fbHandler;nowMonth"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Прошлый месяц", "admin;fbHandler;lastMonth"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "admin;escape")),
)

var AdminPanelShutdown = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Shutdown", "admin;shutdown"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "admin;escape")))

var AdminPanelXLSX = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(".db", "admin;mode;.db"),
		tgbotapi.NewInlineKeyboardButtonData(".xlsx", "admin;mode;.xlsx")),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "admin;escape")),
)

var AdminPanelDB = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("All database", "admin;all"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Users", "admin;users"),
		tgbotapi.NewInlineKeyboardButtonData("Records", "admin;records"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Tracker", "admin;tracker"),
		tgbotapi.NewInlineKeyboardButtonData("Students", "admin;students"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("FeedBack", "admin;fb"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "admin;escape")),
)

var AdminPanelEscape = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "admin;escape")),
)

var MenuScheduleBotKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Получить расписание", "menu;schedule;Получить расписание",
		)),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Прочее", "menu;schedule;Прочее",
		),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Назад", "menu;schedule;Сменить бота",
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
			"Без фильтров", "menu;filter;Без фильтров",
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

func CopyInlineKeyboard(kb tgbotapi.InlineKeyboardMarkup) tgbotapi.InlineKeyboardMarkup {
	will := tgbotapi.InlineKeyboardMarkup{}
	willCopy_ := make([][]tgbotapi.InlineKeyboardButton, len(kb.InlineKeyboard))
	copy(willCopy_, kb.InlineKeyboard)
	will.InlineKeyboard = willCopy_
	return will
}
