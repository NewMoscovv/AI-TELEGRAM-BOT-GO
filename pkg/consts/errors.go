package consts

// ошибки инициализации
const (
	ConfigInitialisationError = "ошибка инициализации конфигурации"
	BotStartPollingError      = "ошибка запуска бота"
)

const (
	TypingAnimationError = "ошибка анимации TYPING"
)

// ошибки open_router
const (
	ResponseBodyError    = "ошибка при чтении ответа"
	ApiRouterError       = "ошибка API"
	JSONParsingError     = "ошибка при парсинге JSON"
	EmptyAnswerByAIError = "нет ответа от ИИ"
)

// ошибки при сборке конфигурации
const (
	NoTelegramToken   = "отсутствует TELEGRAM_TOKEN"
	NoOpenRouterToken = "отсутствует токен OpenRouter"
	NoUrl             = "отсутствует API_URL"
	NoModel           = "отсутствует название Модели"
)
