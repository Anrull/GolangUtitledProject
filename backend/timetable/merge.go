package timetable

import (
	"fmt"
	"reflect"
)

func Merge(lessons, extraLessons [][]string) [][]string {
	if reflect.DeepEqual(extraLessons, [][]string{{"", ""}, {"", ""}, {"", ""}}) {
		return lessons
	}
	if len(extraLessons) != 3 {
		return lessons
	}
	mergedLessons := [][]string{}
	mergedLessons = append(mergedLessons, lessons...)

	if len(lessons) == 6 {
		mergedLessons = append(mergedLessons, delete(extraLessons)...)
	} else if len(lessons) == 7 {
		if reflect.DeepEqual(extraLessons[0], []string{"", ""}) {
			mergedLessons = append(mergedLessons, delete(extraLessons[1:])...)
		} else {
			mergedLessons[6] = []string{
				fmt.Sprintf("%s/%s", lessons[6][0], extraLessons[0][0]),
				fmt.Sprintf("%s/%s", lessons[6][1], extraLessons[0][1]),}
			mergedLessons = append(mergedLessons, delete(extraLessons[1:])...)
		}
	} else if len(lessons) == 8 {
		if reflect.DeepEqual(extraLessons[:2], [][]string{{"", ""}, {"", ""}}) {
			mergedLessons = append(mergedLessons, delete(extraLessons[2:])...)
		} else {
			mergedLessons[6] = []string{
				fmt.Sprintf("%s/%s", lessons[6][0], extraLessons[0][0]),
				fmt.Sprintf("%s/%s", lessons[6][1], extraLessons[0][1]),}
			mergedLessons = append(mergedLessons, delete(extraLessons[1:])...)
		}
	} else {
		mergedLessons = append(mergedLessons, extraLessons...)
	}

	return mergedLessons
}


func delete(data [][]string) [][]string {
	var result [][]string
	nonEmptyFound := false

	for i := len(data) - 1; i >= 0; i-- {
		if data[i][0] != "" || data[i][1] != "" {
			nonEmptyFound = true
		}
		if nonEmptyFound {
			result = append([][]string{data[i]}, result...)
		}
	}

	return result
}