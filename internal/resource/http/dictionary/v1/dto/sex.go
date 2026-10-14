package dto

import (
	"strings"

	rcmodel "github.com/web-rabis/reader-client/model"
)

type DSex struct {
	rcmodel.Dictionary
	Translations LangOptions `json:"translations"`
}

func NewSexList(sexList []*rcmodel.Dictionary) []DSex {
	newSexList := make([]DSex, 2)
	for _, sex := range sexList {
		sexName := strings.ToLower(sex.Name)
		if sexName == "мужской" ||
			sexName == "муж" ||
			sexName == "м" ||
			sexName == "male" ||
			sexName == "m" ||
			sexName == "е" ||
			sexName == "еркек" ||
			sexName == "ер" {
			newSexList[0] = DSex{
				Dictionary: *sex,
				Translations: LangOptions{
					EN: "Male",
					RU: "Мужской",
					KZ: "Еркек",
				},
			}
		}
		if sexName == "женский" ||
			sexName == "жен" ||
			sexName == "ж" ||
			sexName == "female" ||
			sexName == "f" ||
			sexName == "ә" ||
			sexName == "әйел" {
			newSexList[1] = DSex{
				Dictionary: *sex,
				Translations: LangOptions{
					EN: "Female",
					RU: "Женский",
					KZ: "Әйел",
				},
			}
		}
	}
	return newSexList
}
