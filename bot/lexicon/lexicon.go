package lexicon

var (
	OlimpListStep  = 10
	OlimpListLeft  = "⬅️"
	OlimpListRight = "➡️"
)

var HelpMessage = `Это бот-расписание для БЛИ № 3
Команды для бота-расписания:
♦️ /schedule - узнать расписание
♦️ /days - расписание по дням
♦️ /tomorrow - узнать расписание на завтра
♦️ /week - узнать текущую идею
♦️ /time - узнать расписание звонков
Чтобы узнать расписание у любого класса в любой день требуется написать следующее: 6A н пн, где 6A - класс (англ), н - неделя (н/ч), пн - день

Команды для трекер-бота:
♦️ /add - добавить новую запись
♦️ /my_olimps - получить свои записи в трекере
♦️ /get_treker - получение всех данных (требуется пароль)

Общие:
♦️ /newsletter - уведомления`

var AccuracyRecord = `Подтвердите правильность введенных вами данных:
%s
%s
%s
%s`

var TimetableTime = `<b>Расписание звонков</b> ⏰:

0️⃣ — <b><em>08.00 - 08.25</em></b>
1️⃣ — <b><em>08.30 - 09.10</em></b>
2️⃣ — <b><em>09.20 - 10.00</em></b>
3️⃣ — <b><em>10.20 - 11.00</em></b>
4️⃣ — <b><em>11.10 - 11.50</em></b>
5️⃣ — <b><em>12.00 - 12.40</em></b>
6️⃣ — <b><em>13.00 - 13.40</em></b>
7️⃣ — <b><em>14.00 - 14.40</em></b>
8️⃣ — <b><em>14.50 - 15.30</em></b>
`

var SubjectsForButton = []string{"Химия (19)",
	"Физика (24)",
	"Математика (27)",
	"Информатика (19)",
	"Биология (9)",
	"Астрономия (2)",
	"Искусственный интеллект (2)",
	"Инженерное дело (2)",
	"Инженерные науки (3)",
	"ОСТАЛЬНЫЕ (70)",
	"Экономика (8)",
	"Естественные науки (2)",
	"История (12)",
	"Лингвистика (2)",
	"Литература (6)",
	"Финансовая грамотность (6)",
	"Иностранный язык(15)",
	"Дизайн (2)",
	"Психология (3)",
	"Обществознание (14)",
	"География (9)",
	"Русский язык (9)",
	"Генетика (2)",
	"Геология (3)",
	"Гуманитарные и социальные науки (3)",
	"Журналистика (5)",
	"Информационная безопасность (2)",
	"Международные отношения (2)",
	"Политология (4)",
	"Право (9)",
	"Рисунок (2)",
	"Робототехника (3)",
	"Социология (2)",
	"Теория и история музыки (2)",
	"Филология (4)",
	"Философия (3)"}

var ListDays = []string{"пн", "вт", "ср", "чт", "пт"}
var Stages = []string{"11B", "11A", "10B", "10A", "9B", "9A", "8B", "8A", "7C", "7B", "7A", "6B", "6A"}

var Week = map[string]string{"0": "чет", "1": "нечет"}
var Day = map[string]string{"0": "пн", "1": "вт", "2": "ср", "3": "чт", "4": "пт"}

var StagesTracker = []string{"Прошел регистрацию", "Написал отборочные этапы (волну, тур)", "Прошел на заключительный этап (финал)",
	"Написал финал", "Победитель финала - I степени", "Призер финала - II степени", "Призер финала - III степени",
	"Сертификат участника финала", "Абсолютный участник"}

var TeacherTracker = []string{
	"Насртдинов Алмаз Касимович", "Казнабаев Ильдар Гильфанович",
	"Латыпова Альфия Файзрахмановна", "Фатхиев Нафис Назифович",
	"Гайсина Гузель Фаритовна", "Другой наставник", "Ахунова Гульнур Юсуповна",
	"Спевак Мария Владимировна", "Туктаров Фанзиль Илгамович",
	"Каипов Ильдар Ишбулдович", "Бывшева Алена Сергеевна",
	"Хамидуллина Зульфия Хурматовна", "Суяргулова Гульнур Закировна",
	"Ахметова Гузель Рафкатовна", "Аскарова Гульнар Мирсаяфовна",
	"Ахметова Гульшат Авхадиевна", "Баталлова Лилия Маратовна",
	"Бондарева Лилия Егоровна", "Галеева Наиля Рамиловна",
	"Губайдуллина Рамзия Ишмухаметовна", "Зайнетдинова Зилфера Арсланбиковна",
	"Каримова Аделина Фанилевна", "Касимова Розалия Рауфовна",
	"Конова Ольга Геннадьевна", "Коробейникова Ольга Владимировна",
	"Лапухина Зухра Сахияновна", "Латыпова Зухра Нурисламовна",
	"Максутов Айбулат Наилович", "Ташбулатова Гульнур Саматовна",
	"Халитова Милауша Нажибовна",
}

var TrackerOlimps = []string{"1 - «Финатлон для старшеклассников» - Всеросс.  олимпиада по финансовой грамотности, устойчивому " +
	"развитию и защите прав потребителей финансовых услуг",
	"10 - Всеросс.  олимпиада школьников «Миссия выполнима. Твое призвание - финансист!»",
	"11 - Всеросс.  Сеченовская олимпиада школьников",
	"12 - Всеросс.  Толстовская олимпиада школьников",
	"13 - Всеросс.  экономическая олимпиада школьников имени Н.Д. Кондратьева",
	"14 - Всероссийский конкурс научных работ школьников «Юниор»",
	"15 - Всесибирская открытая олимпиада школьников",
	"16 - Вузовско-академическая олимпиада по информатике",
	"17 - Герценовская олимпиада школьников",
	"18 - Городская открытая олимпиада школьников по физике",
	"19 - Государственный аудит", "2 - «Формула Единства»/«Третье тысячелетие»",
	"20 - Инженерная олимпиада школьников", "21 - Интернет-олимпиада школьников по физике",
	"22 - Кутафинская олимпиада школьников по праву",
	"23 - Междисциплинарная олимпиада школьников имени В.И. Вернадского",
	"24 - Междунар. Менделеевская олимпиада школьников по химии",
	"25 - Междунар. олимпиада «Innopolis Open»",
	"26 - Междунар. олимпиада по финансовой безопасности",
	"27 - Междунар. олимпиада школьников «Искусство графики»",
	"28 - Междунар. олимпиада школьников Уральского федерального университета «Изумруд»",
	"29 - Межрег. олимпиада по праву «ФЕМИДА»",
	"3 - XVI Южно-Российская Межрег. олимпиада школьников «Архитектура и искусство» по комплексу " +
		"предметов (рисунок, живопись, композиция, черчение)",
	"30 - Межрег. олимпиада школьников «САММАТ»",
	"31 - Межрег. олимпиада школьников «Архитектура и искусство» по комплексу предметов (рисунок, " +
		"композиция)",
	"32 - Межрег. олимпиада школьников «Будущие исследователи - будущее науки»",
	"33 - Межрег. олимпиада школьников «Евразийская лингвистическая олимпиада»",
	"34 - Межрег. олимпиада школьников имени В.Е. Татлина",
	"35 - Межрег. олимпиада школьников имени И.Я. Верченко",
	"36 - Межрег. олимпиада школьников на базе ведомственных образовательных организаций",
	"37 - Межрег. олимпиады «Казанский (Приволжский) федеральный университет»",
	"38 - Многопредметная олимпиада «Юные таланты»",
	"39 - Многопрофильная инженерная олимпиада «Звезда»",
	"4 - Всеросс.  (с международным участием) олимпиада учащихся музыкальных колледжей",
	"40 - Московская олимпиада школьников",
	"41 - Общероссийская олимпиада школьников «Основы православной культуры»",
	"42 - Объединенная межвузовская математическая олимпиада школьников",
	"43 -  НТО:  технологии виртуальной и дополненной реальности",
	"43 -  НТО:  цифровые технологии в архитектуре", "43-  НТО:  беспилотные авиационные системы",
	"43 - Океан знаний", "44 - Олимп. Курчатов",
	"44 - Олимпиада по комплексу предметов «Культура и искусство»",
	"45 - Олимп. МГИМО МИД России для школьников", "45 - Олимпиада РГГУ для школьников",
	"46 - Олимп. по архитектуре СПбГАСУ",
	"47 - Олимп. по комплексу предметов «Культура и искусство»", "48 - Олимп. РГГУ для школьников",
	"49 - Олимп. шк. «Высокие технологии и материалы будущего»",
	"5 - НТО (Национальная технологическая олимпиада)",
	"50 - Олимп. шк. «Гранит науки»", "51 - Олимп. шк. «Ломоносов»",
	"52 - Олимп. шк. «Надежда энергетики»",
	"53 - Олимп. шк. «Покори Воробьевы горы!»", "54 - Олимп. шк. «Робофест»",
	"55 - Олимп. шк. «Физтех»",
	"56 - Олимп. шк. «Шаг в будущее»",
	"56 - Олимпиада по экономике в рамках международного экономического фестиваля «Сибириада. Шаг в " +
		"мечту»",
	"57 - Олимп. шк. по географии «Земля - наш общий дом!»",
	"57 - Олимпиада РАНХиГС", "58 - Олимп. шк. по информатике и программированию",
	"59 - Олимп. шк. по программированию «ТехноКубок»", "59 - Олимпиада школьников «В мир права»",
	"6 - Всеросс.  олимпиада по искусственному интеллекту",
	"60 - Олимп. шк. по химии в ФГБОУ ВО ПСПбГМУ им. И. П. Павлова Минздрава России",
	"60 - Олимпиада Юношеской математической школы",
	"61 - Олимп. шк. по экономике в рамках международного экономического фестиваля школьников " +
		"«Сибириада. Шаг в мечту»",
	"62 - Олимп. шк. Российской академии народного хозяйства и государственной службы при Президенте " +
		"Российской Федерации",
	"63 - Олимп. шк. Санкт-Петербургского государственного университета",
	"63 - Открытая олимпиада школьников по программированию",
	"64 - Олимп. шк. федерального государственного бюджетного образовательного учреждения высшего " +
		"образования «Всероссийский государственный университет юстиции (РПА Минюста России)» «В мир " +
		"права»",
	"65 - Олимп. Юношеской математической школы", "65 - Открытая олимпиада вузов Томской области",
	"66 - Открытая межвузовская олимпиада школьников Сибирского Федерального округа «Будущее Сибири»",
	"67 - Открытая олимпиада школьников", "68 - Открытая олимпиада школьников по программированию",
	"69 - Открытая олимпиада школьников по программированию «Когнитивные технологии»",
	"69 - Плехановская олимпиада",
	"7 - Всеросс.  олимпиада по музыкально-теоретическим дисциплинам для учащихся детских музыкальных " +
		"школ и детских школ искусств",
	"70 - Открытая региональная межвузовская олимпиада вузов Томской области (ОРМО)",
	"70 - Региональный конкурс школьников Челябинского университетского образовательного округа",
	"71 - Открытая химическая олимпиада", "72 - Отраслевая олимпиада школьников «Газпром»",
	"73 - Отраслевая физико-математическая олимпиада школьников «Росатом»",
	"73 - Северо-Восточная олимпиада", "74 - Пироговская олимпиада для школьников по химии и биологии",
	"74 - Региональный конкурс школьников Челябинского университетского образовательного округа",
	"74 - Сибирская олимпиада «Архитектурно-дизайнерское творчество»",
	"75 - Олимпиады МГХПА имени С. Г. Строганова", "75 - Плехановская олимпиада школьников",
	"75 - Санкт-Петербургская астрономическая олимпиада",
	"76 - Региональный конкурс школьников Челябинского университетского образовательного округа",
	"76 - Телевизионная гуманитарная олимпиада школьников «Умницы и умники»",
	"77 - Санкт-Петербургская астрономическая олимпиада", "77 - Северо-Восточная олимпиада",
	"77 - Турнир городов", "78 - Санкт-Петербургская олимпиада школьников",
	"78 - Сибирская олимпиада «Архитектурно-дизайнерское творчество»",
	"79 - Олимпиады МГХПА имени С. Г. Строганова", "79 - Северо-Восточная олимпиада школьников",
	"8 - Всеросс.  олимпиада школьников «Юридические высоты!»",
	"80 - Олимпиада школьников «Учитель школы будущего»",
	"80 - Сибирская Межрег. олимпиада школьников «Архитектурно-дизайнерское творчество»",
	"80 - Телевизионная гуманитарная олимпиада школьников «Умницы и умники»",
	"81 - Строгановская олимпиада",
	"83 - Турнир городов", "84 - Олимпиада школьников «Учитель школы будущего»",
	"84 - Турнир имени М.В. Ломоносова", "85 - Университетская олимпиада школьников «Бельчонок»",
	"85 - Филологическая олимпиада «Юный словесник»",
	"86 - Международная олимпиада по финансовой безопасности",
	"86 - Учитель школы будущего",
	"87 - Всеросс.  олимпиада по агрогенетике для школьников старших классов «Иннагрика»",
	"9 - Всеросс.  олимпиада школьников «Высшая проба»",
}
