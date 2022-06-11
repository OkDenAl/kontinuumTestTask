package pkg

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"strconv"
	"strings"
)

type Student struct {
	id         int
	name       string
	groups     []Group
	fullRating float64
}

type Group struct {
	id                                               int
	rating, m                                        float64
	countOfLessons                                   int
	PointsForHW, PointsForActivity, PointsForLessons float64
}

//Функция считает кол-во баллов за Домашку
func (student *Student) countPointsForHW(hwPercent, groupId int) {
	switch {
	case hwPercent == 100:
		student.groups[groupId].PointsForHW += 3
	case hwPercent >= 80 && hwPercent < 100:
		student.groups[groupId].PointsForHW += 2
	case hwPercent >= 50 && hwPercent < 80:
		student.groups[groupId].PointsForHW += 1
	case hwPercent < 30:
		student.groups[groupId].PointsForHW -= 0.5
	}
}

//Функция считает кол-во баллов за Активность на занятии
func (student *Student) countPointsForActivity(activity, groupId int) {
	switch {
	case activity == 4 || activity == 5:
		student.groups[groupId].PointsForActivity += 1
	case activity == 2 || activity == 3:
		student.groups[groupId].PointsForActivity += 0.5
	}
}

//Функция считает кол-во баллов за занятие основываясь на формуле [баллы за пробник]*0.02*[кол-во занятий]
func (student *Student) countPointsForLessons(groupId int) {
	student.groups[groupId].PointsForLessons +=
		student.groups[groupId].m * float64(student.groups[groupId].countOfLessons)
}

//Функция, сканирующая инфо о студентах: id и имя
//Вовзращает массив студентов и возможную ошибку
func scanStudents(db *sql.DB, countOfGroups int) (students []Student, error []string) {
	students = make([]Student, 0)
	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		return nil, []string{"Error: Не уадлось сосканировать студентов"}
	}
	defer rows.Close()
	for rows.Next() {
		curPerson := Student{groups: make([]Group, 0)}
		rows.Scan(&curPerson.id, &curPerson.name)
		for i := 0; i < countOfGroups; i++ {
			curPerson.groups = append(curPerson.groups, Group{id: i + 1})
		}
		students = append(students, curPerson)
	}
	return students, nil
}

//Функция, сканирующая инфо о занятиях
//Вовзращает возможную ошибку
func scanDataFromLessonsStat(db *sql.DB, students []Student, studentId, groupId int, Dates []string, errorCh chan []string) {
	rows, err := db.Query(fmt.Sprintf("SELECT activity, HW FROM data_about_lessons "+
		"WHERE idStudent = %d AND idGroup = %d AND lesson_date BETWEEN '%s' AND '%s'",
		students[studentId].id, groupId+1, Dates[0], Dates[1]))
	if err != nil {
		errorCh <- []string{"Error: Не удалось сосканировать информацию о статистике с занятий из БД"}
		return
	}
	defer rows.Close()
	for rows.Next() {
		students[studentId].groups[groupId].countOfLessons++
		var hwPercent, activity int
		rows.Scan(&activity, &hwPercent)
		students[studentId].countPointsForHW(hwPercent, groupId)
		students[studentId].countPointsForActivity(activity, groupId)
	}
	errorCh <- nil
}

//Функция, сканирующая инфо о результатах пробника
//Вовзращает возможную ошибку
func scanDataFromTestResStat(db *sql.DB, students []Student, studentId, groupId int, errorCh chan []string) {
	rows, err := db.Query(fmt.Sprintf("SELECT res FROM last_test_res "+
		"WHERE idStudent = %d AND idGroup = %d",
		students[studentId].id, groupId+1))
	if err != nil {
		errorCh <- []string{"Error: Не удалось сосканировать информацию о результатах пробников из БД"}
		return
	}
	defer rows.Close()
	for rows.Next() {
		var res float64
		rows.Scan(&res)
		students[studentId].groups[groupId].m = res * 0.02
	}
	errorCh <- nil
}

//Функция, сканирующая инфо о группах : расшифровку индексов, кол-во групп
//Вовзращает словарь [номер группы]-->[текстовая расшифровка], кол-во групп, возможную ошибку
func scanInfoAboutGroups(db *sql.DB) (subjects map[int]string, countOfGroups int, error []string) {
	subjects = make(map[int]string)
	countOfGroups = 0
	rows, err := db.Query("SELECT name FROM groups")
	if err != nil {
		return nil, -1, []string{"Error: Не удалось сосканировать информацию о группах из БД"}
	}
	defer rows.Close()
	for rows.Next() {
		countOfGroups++
		var name string
		rows.Scan(&name)
		subjects[countOfGroups] = name
	}
	return subjects, countOfGroups, nil
}

//Функция,генерирующая ответ
//Вовзращает массив строк результаты подсчёта баллов
func makeAnswer(students []Student, countOfStudents int, subjects map[int]string, targets []int, Args []string) (answer []string) {
	answer = make([]string, 0)
	if len(targets) > 0 {
		for _, target := range targets {
			answer = append(answer, fmt.Sprintf("%s с %s по %s набрал(a) %.2f баллов\n",
				students[target].name, Args[0], Args[1], students[target].fullRating))
		}
	} else {
		for i := 0; i < len(students); i++ {
			answer = append(answer, fmt.Sprintf("%s с %s по %s набрал(a) %.2f бaллов\n",
				students[i].name, Args[0], Args[1], students[i].fullRating))
		}
	}
	return answer
}

//Функция-диспетчер, управляющая обработкой информации из баз данных, вызывает другие функции
//Вовзращает возможную ошибку
func processInfo(db *sql.DB, students []Student, countOfGroups, studentId int, Dates []string) (err []string) {
	errorCh1 := make(chan []string)
	errorCh2 := make(chan []string)
	for j := 0; j < countOfGroups; j++ {
		go scanDataFromLessonsStat(db, students, studentId, j, Dates, errorCh1) //Ассинхронка :)
		go scanDataFromTestResStat(db, students, studentId, j, errorCh2)        //Ассинхронка :)
		err1, err2 := <-errorCh1, <-errorCh2
		if err1 != nil || err2 != nil {
			return err1
		}
		students[studentId].countPointsForLessons(j)
		if students[studentId].groups[j].countOfLessons != 0 {
			sum := students[studentId].groups[j].PointsForActivity +
				students[studentId].groups[j].PointsForHW + students[studentId].groups[j].PointsForLessons
			students[studentId].groups[j].rating = sum / float64(students[studentId].groups[j].countOfLessons)
			students[studentId].fullRating += students[studentId].groups[j].rating
		}
	}
	return nil
}

func strToInt(s string) int {
	ch, _ := strconv.Atoi(s)
	return ch
}

//Функция, проверяющая то, что даты идут в нужном хронологическом порядке
func checkDateChronology(date1, date2 string) bool {
	date1Arr := strings.Split(date1, ".")
	date2Arr := strings.Split(date2, ".")
	if strToInt(date1Arr[0]) < strToInt(date2Arr[0]) {
		return true
	} else if strToInt(date1Arr[1]) < strToInt(date2Arr[1]) &&
		strToInt(date1Arr[0]) == strToInt(date2Arr[0]) {
		return true
	} else if strToInt(date1Arr[2]) < strToInt(date2Arr[2]) &&
		strToInt(date1Arr[1]) == strToInt(date2Arr[1]) &&
		strToInt(date1Arr[0]) == strToInt(date2Arr[0]) {
		return true
	}
	return false
}

//Функция, проверяющая есть ли даты в базе данных
func isThisDateExistInDB(date1, date2 string, db *sql.DB) bool {
	rows, err := db.Query("SELECT lesson_date FROM data_about_lessons LIMIT 1")
	if err != nil {
		return false
	}
	var firstDateInBD string
	for rows.Next() {
		rows.Scan(&firstDateInBD)
	}
	firstDateInBD = strings.Replace(firstDateInBD, "-", ".", 2)
	rows, err = db.Query("SELECT lesson_date FROM data_about_lessons ORDER BY lesson_date DESC LIMIT 1")
	if err != nil {
		return false
	}
	var lastDateInBD string
	for rows.Next() {
		rows.Scan(&lastDateInBD)
	}
	lastDateInBD = strings.Replace(lastDateInBD, "-", ".", 2)
	if checkDateChronology(date1, firstDateInBD) {
		return false
	} else if checkDateChronology(lastDateInBD, date2) {
		return false
	}
	return true
}

//Основная функция кода - типа ядро
func SqlHandler(Args []string) []string {
	//start := time.Now()
	if len(Args) < 2 {
		return ([]string{"Error: Вы не указали даты"})
	} else {
		matched1, _ := regexp.MatchString(`^\d{4}\.\d{2}\.\d{2}\b`, Args[0])
		matched2, _ := regexp.MatchString(`^\d{4}\.\d{2}\.\d{2}\b`, Args[1])
		if !matched1 || !matched2 {
			return ([]string{"Error: Вы не указали даты в нужном формате"})
		}
	}

	if checkDateChronology(Args[1], Args[0]) {
		return ([]string{"Error: Вы записали в даты не в хронологическом порядке"})
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/continuum")
	fmt.Println(db)
	if err != nil {
		return ([]string{"Error: Не удалось подключиться к базе данных"})
	}
	defer db.Close()

	if !isThisDateExistInDB(Args[0], Args[1], db) {
		return ([]string{"Error: таких дат нет в БД"})
	}

	subjects, countOfGroups, customErr := scanInfoAboutGroups(db)
	if customErr != nil {
		return customErr
	}

	students, customErr := scanStudents(db, countOfGroups)
	if customErr != nil {
		return customErr
	}

	targets := make([]int, 0)
	if len(Args) > 2 {
		for i := 2; i < len(Args); i++ {
			if Args[i] == "" {
				continue
			}
			target, err := strconv.Atoi(Args[i])
			if err != nil {
				return ([]string{"Error: корректно укажите id учеников (целое число)"})
			}
			if target > len(students) || target < 1 {
				return ([]string{"Error: Такого id нет в базе"})
			}
			targets = append(targets, target-1)
		}
	}

	if len(targets) > 0 {
		for _, target := range targets {
			customErr = processInfo(db, students, countOfGroups, target, Args)
			if customErr != nil {
				return customErr
			}
		}
	} else {
		for i := 0; i < len(students); i++ {
			customErr = processInfo(db, students, countOfGroups, i, Args)
			if customErr != nil {
				return customErr
			}
		}
	}
	/*elapsedTime := time.Since(start)
	fmt.Println("Total Time For Execution: " + elapsedTime.String())*/
	return makeAnswer(students, countOfGroups, subjects, targets, Args)
}
