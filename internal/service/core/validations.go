package core

import (
	"github.com/tandesko/cards_validator/resources"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strconv"
	"strings"
	"time"
)

func ValidateAllMethods(data resources.ValidateCardRequestAttributes, log *logan.Entry) (bool, error) {
	cardnumber := strings.Replace(data.Cardnumber, "-", "", -1)
	luhnResult := luhnMethodCheck(cardnumber)
	dateCheckResult, err := dateCheck(data.Month, data.Year, log)
	if err != nil {
		return false, err
	}
	return luhnResult && dateCheckResult && iinCheck(cardnumber), nil
}

func dateCheck(monthStr string, yearStr string, log *logan.Entry) (bool, error) {
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		log.WithError(err).Info("wrong month")
		return false, errors.Wrap(err, "wrong month")
	}

	year := 0
	if len(yearStr) < 3 {
		year, err = strconv.Atoi(strconv.Itoa(time.Now().Year())[:2] + yearStr)
	} else {
		year, err = strconv.Atoi(yearStr)
	}
	if err != nil {
		log.WithError(err).Info("wrong year")
		return false, errors.Wrap(err, "wrong year")
	}

	validDate := false
	if time.Now().Year() < year || int(time.Now().Month()) < month && time.Now().Year() == year {
		validDate = true
	}

	return validDate, nil
}

func luhnMethodCheck(cardnumber string) bool {
	var sum int
	parity := len(cardnumber) % 2
	for i := len(cardnumber) - 2; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(cardnumber[i]))
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	lastDigit, _ := strconv.Atoi(string(cardnumber[len(cardnumber)-1]))
	return 10-(sum%10) == lastDigit
}

func iinCheck(cardnumber string) bool {
	return isAmericanExp(cardnumber) ||
		isChinaTUnion(cardnumber) ||
		isChinaUnionPay(cardnumber) ||
		isDinersClubInter(cardnumber) ||
		isDinersClubUSCanada(cardnumber) ||
		isDiscoveCard(cardnumber) ||
		isUkrCard(cardnumber) ||
		isRuPay(cardnumber) ||
		isInterPayment(cardnumber) ||
		isInstaPayment(cardnumber) ||
		isJCB(cardnumber) ||
		isLaser(cardnumber) ||
		isMaestroUK(cardnumber) ||
		isMaestro(cardnumber) ||
		isMaestroUK(cardnumber) ||
		isDankort(cardnumber) ||
		isMir(cardnumber) ||
		isBorica(cardnumber) ||
		isNPSPridnestrovie(cardnumber) ||
		isMastercard(cardnumber) ||
		isSolo(cardnumber) ||
		isSwitch(cardnumber) ||
		isTroy(cardnumber) ||
		isVisa(cardnumber) ||
		isVisaElectron(cardnumber) ||
		isUATP(cardnumber) ||
		isVerve(cardnumber) ||
		isLankaPay(cardnumber) ||
		isUzCard(cardnumber) ||
		isHumo(cardnumber) ||
		isGPN(cardnumber) ||
		isNapas(cardnumber)
}
