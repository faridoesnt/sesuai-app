package entities

type MobileLogAdmin struct {
	AdminId    string `db:"admin_id"`
	Endpoint   string `db:"endpoint"`
	EventType  string `db:"eventType"`
	Result     string `db:"result"`
	Message    string `db:"message"`
	Header     string `db:"header"`
	Params     string `db:"params"`
	Response   string `db:"response"`
	Device     string `db:"device"`
	DeviceTime string `db:"device_time"`
}

type MobileLogUser struct {
	UserId     string `db:"user_id"`
	Endpoint   string `db:"endpoint"`
	EventType  string `db:"eventType"`
	Result     string `db:"result"`
	Message    string `db:"message"`
	Header     string `db:"header"`
	Params     string `db:"params"`
	Response   string `db:"response"`
	Device     string `db:"device"`
	DeviceTime string `db:"device_time"`
}
