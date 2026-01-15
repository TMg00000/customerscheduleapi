package validation

import (
	"strconv"
	"strings"
	"time"

	"github.com/TMg00000/customerscheduleapi/internal/domain/models/requests"
	"github.com/TMg00000/customerscheduleapi/internal/domain/models/resources/resourceserrorsmessages"
)

func ErrorsInName(client requests.Client, errorsList *[]string) {
	if (strings.HasPrefix(client.Name, " ") || strings.HasPrefix(client.Name, "\n")) ||
		(strings.HasSuffix(client.Name, " ") || strings.HasSuffix(client.Name, "\n")) {
		*errorsList = append(*errorsList, resourceserrorsmessages.TheNameStartedOrFinishWithSpace)
	}

	if strings.TrimSpace(client.Name) == "" {
		*errorsList = append(*errorsList, resourceserrorsmessages.NameIsNil)
	}
}

func ErrorsInPhoneNumber(client requests.Client, errorsList *[]string) {
	phoneStr := strconv.Itoa(client.PhoneNumber)

	if len(phoneStr) != 11 || client.PhoneNumber == 0 {
		*errorsList = append(*errorsList, resourceserrorsmessages.TheNumberMustContainElevenDigits)
	}
}

func ErrorsInTypeServices(client requests.Client, errorsList *[]string) {
	if int(client.TypeService) < 0 || int(client.TypeService) > 8 {
		*errorsList = append(*errorsList, resourceserrorsmessages.OptionInvalid)
	}
}

func ErrorsInDateScheduling(client requests.Client, errorsList *[]string) {
	if client.DateTime.IsZero() {
		*errorsList = append(*errorsList, resourceserrorsmessages.DateTimeIsInvalid)
	}

	if client.DateTime.Before(time.Now()) {
		*errorsList = append(*errorsList, resourceserrorsmessages.ChooseADateLaterThanToday)
	}
}

func ListErrorsMessages(client requests.Client) []string {
	errorsList := []string{}

	ErrorsInName(client, &errorsList)
	ErrorsInPhoneNumber(client, &errorsList)
	ErrorsInTypeServices(client, &errorsList)
	ErrorsInDateScheduling(client, &errorsList)

	return errorsList
}
