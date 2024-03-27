package core

func isAmericanExp(cardnumber string) bool {
	return cardnumber[:2] == "34" ||
		cardnumber[:2] == "37"
}

func isChinaTUnion(cardnumber string) bool {
	return cardnumber[:2] == "31"
}

func isChinaUnionPay(cardnumber string) bool {
	return cardnumber[:2] == "62"
}

func isDinersClubInter(cardnumber string) bool {
	return cardnumber[:2] == "36"
}

func isDinersClubUSCanada(cardnumber string) bool {
	return cardnumber[:2] == "55"
}

func isDiscoveCard(cardnumber string) bool {
	return cardnumber[:4] == "6011" ||
		cardnumber[:2] == "65" ||
		cardnumber[:3] >= "644" && cardnumber[:3] <= "649" ||
		cardnumber[:6] >= "622126" && cardnumber[:6] <= "622925"
}

func isUkrCard(cardnumber string) bool {
	return cardnumber[:8] >= "60400100" && cardnumber[:8] <= "60420099"
}

func isRuPay(cardnumber string) bool {
	return cardnumber[:2] == "60" ||
		cardnumber[:2] == "65" ||
		cardnumber[:2] == "81" ||
		cardnumber[:2] == "82" ||
		cardnumber[:3] == "508" ||
		cardnumber[:3] == "353" ||
		cardnumber[:3] == "356"
}

func isInterPayment(cardnumber string) bool {
	return cardnumber[:3] == "636"
}

func isInstaPayment(cardnumber string) bool {
	return cardnumber[:3] >= "637" && cardnumber[:3] <= "639"
}

func isJCB(cardnumber string) bool {
	return cardnumber[:4] >= "3528" && cardnumber[:4] <= "3589"
}

func isLaser(cardnumber string) bool {
	return cardnumber[:4] == "6304" ||
		cardnumber[:4] == "6706" ||
		cardnumber[:4] == "6771" ||
		cardnumber[:4] == "6709"
}

func isMaestroUK(cardnumber string) bool {
	return cardnumber[:4] == "6759" ||
		cardnumber[:6] == "676770" ||
		cardnumber[:6] == "676774"
}

func isMaestro(cardnumber string) bool {
	return cardnumber[:4] == "5018" ||
		cardnumber[:4] == "5020" ||
		cardnumber[:4] == "5038" ||
		cardnumber[:4] == "5893" ||
		cardnumber[:4] == "6304" ||
		cardnumber[:4] == "6759" ||
		cardnumber[:4] == "6761" ||
		cardnumber[:4] == "6762" ||
		cardnumber[:4] == "6763"
}

func isDankort(cardnumber string) bool {
	return cardnumber[:4] == "5019" ||
		cardnumber[:4] == "4571"
}

func isMir(cardnumber string) bool {
	return cardnumber[:4] >= "2200" && cardnumber[:4] <= "2204"
}

func isBorica(cardnumber string) bool {
	return cardnumber[:4] == "2205"
}

func isNPSPridnestrovie(cardnumber string) bool {
	return cardnumber[:7] >= "6054740" && cardnumber[:4] <= "6054744"
}

func isMastercard(cardnumber string) bool {
	return cardnumber[:2] >= "51" && cardnumber[:2] <= "55" ||
		cardnumber[:4] >= "2221" && cardnumber[:4] <= "2720"
}

func isSolo(cardnumber string) bool {
	return cardnumber[:4] == "6334" ||
		cardnumber[:4] == "6767"
}

func isSwitch(cardnumber string) bool {
	return cardnumber[:4] == "4903" ||
		cardnumber[:4] == "4905" ||
		cardnumber[:4] == "4911" ||
		cardnumber[:4] == "4936" ||
		cardnumber[:6] == "564182" ||
		cardnumber[:6] == "633110" ||
		cardnumber[:4] == "6333" ||
		cardnumber[:4] == "6759"
}

func isTroy(cardnumber string) bool {
	return cardnumber[:2] == "9792"
}

func isVisa(cardnumber string) bool {
	return string(cardnumber[0]) == "4"
}

func isVisaElectron(cardnumber string) bool {
	return cardnumber[:4] == "4026" ||
		cardnumber[:6] == "417500" ||
		cardnumber[:4] == "4508" ||
		cardnumber[:4] == "4844" ||
		cardnumber[:4] == "4913" ||
		cardnumber[:4] == "4917"
}

func isUATP(cardnumber string) bool {
	return string(cardnumber[0]) == "1"
}

func isVerve(cardnumber string) bool {
	return cardnumber[:6] >= "506099" && cardnumber[:6] <= "506198" ||
		cardnumber[:6] >= "650002" && cardnumber[:6] <= "650027" ||
		cardnumber[:6] >= "507865" && cardnumber[:6] <= "507964"
}

func isLankaPay(cardnumber string) bool {
	return cardnumber[:6] == "357111"
}

func isUzCard(cardnumber string) bool {
	return cardnumber[:4] == "8600"
}

func isHumo(cardnumber string) bool {
	return cardnumber[:4] == "9860"
}

func isGPN(cardnumber string) bool {
	return cardnumber[:4] == "1946" ||
		cardnumber[:2] == "50" ||
		cardnumber[:2] == "56" ||
		cardnumber[:2] == "58" ||
		cardnumber[:2] >= "60" && cardnumber[:2] <= "63"
}

func isNapas(cardnumber string) bool {
	return cardnumber[:4] == "9704"
}
