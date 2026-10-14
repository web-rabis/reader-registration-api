package dto

import (
	"strings"

	rcmodel "github.com/web-rabis/reader-client/model"
)

type DEducation struct {
	rcmodel.Dictionary
	Translations LangOptions `json:"translations"`
}

func NewEducationList(educationList []*rcmodel.Dictionary) []DEducation {
	newEducationList := map[string]DEducation{}
	for _, edu := range educationList {
		eduName := strings.ToLower(edu.Name)
		if eduName == "среднее" {
			newEducationList["secondary"] = DEducation{
				Dictionary: *edu,
				Translations: LangOptions{
					EN: "secondary",
					RU: "среднее",
					KZ: "орта",
				},
			}
		}
		if eduName == "высшее" {
			newEducationList["higher"] = DEducation{
				Dictionary: *edu,
				Translations: LangOptions{
					EN: "higher",
					RU: "высшее",
					KZ: "жоғары",
				},
			}
		}
		if eduName == "н/высшее" {
			newEducationList["incomplete higher"] = DEducation{
				Dictionary: *edu,
				Translations: LangOptions{
					EN: "incomplete higher",
					RU: "незаконченое высшее",
					KZ: "аяқталмаған жоғары",
				},
			}
		}
		if eduName == "н/среднее" {
			newEducationList["incomplete secondary"] = DEducation{
				Dictionary: *edu,
				Translations: LangOptions{
					EN: "incomplete secondary",
					RU: "незаконченое среднее",
					KZ: "аяқталмаған орта",
				},
			}
		}
		if eduName == "сред / спец." {
			newEducationList["secondary specialized"] = DEducation{
				Dictionary: *edu,
				Translations: LangOptions{
					EN: "secondary specialized",
					RU: "среднеспециальное",
					KZ: "арнаулы орта",
				},
			}
		}
	}
	result := make([]DEducation, 0, len(newEducationList))
	for _, v := range newEducationList {
		result = append(result, v)
	}
	return result
}
