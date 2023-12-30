package helpers

func GetShioSupport(text string) string {
	if text >= "23:00:00" && text <= "00:59:00" {
		return "Tikus"
	} else if text >= "01:00:00" && text <= "02:59:00" {
		return "Kerbau"
	} else if text >= "03:00:00" && text <= "04:59:00" {
		return "Macan"
	} else if text >= "05:00:00" && text <= "06:59:00" {
		return "Kelinci"
	} else if text >= "07:00:00" && text <= "08:59:00" {
		return "Naga"
	} else if text >= "09:00:00" && text <= "10:59:00" {
		return "Ular"
	} else if text >= "11:00:00" && text <= "12:59:00" {
		return "Kuda"
	} else if text >= "13:00:00" && text <= "14:59:00" {
		return "Kambing"
	} else if text >= "15:00:00" && text <= "16:59:00" {
		return "Kera"
	} else if text >= "17:00:00" && text <= "18:59:00" {
		return "Ayam"
	} else if text >= "19:00:00" && text <= "20:59:00" {
		return "Anjing"
	} else {
		return "Babi"
	}
}
