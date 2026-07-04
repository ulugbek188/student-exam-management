package main

import "fmt"

const (
	SchoolName              = "PDP University"
	Country                 = "Uzbekistan"
	AcademicYear            = "2026-2027"
	PassMark                = 60
	ExcellentScore          = 90
	VeryGoodScore           = 80
	GoodScore               = 70
	AverageScore            = 60
	FullScholarshipScore    = 95
	PartialScholarshipScore = 85
)

// Returns basic student information
func studentInfo() (string, string, int, string, string, string) {

	studentName := "Ulug'bek Zokirov"
	studentID := "ST250022"
	age := 18
	faculty := "Computer Science"
	course := "Backend Engineering"
	subject := "Go Programming"

	return studentName, studentID, age, faculty, course, subject
}

// Returns student exam information
func studentExam() (int, int, int, int, float64) {
	attendance := 98
	assignmentScore := 90
	midtermScore := 99
	finalExamScore := 95
	examScore := 99.8

	return attendance, assignmentScore, midtermScore, finalExamScore, examScore
}

// Calculates student's grade
func calculateGrade(examScore float64) string {

	if examScore >= ExcellentScore {
		return "A"
	} else if examScore >= VeryGoodScore {
		return "B"
	} else if examScore >= GoodScore {
		return "C"
	} else if examScore >= AverageScore {
		return "D"
	}
	return "F"
}

// Returns student's exam status
func studentStatus(examScore float64) string {

	if examScore >= PassMark {
		return "PASSED"
	}
	return "FAILED"
}

// Returns student's performance
func studentPerformance(grade string) string {

	switch grade {

	case "A":
		return "Excellent"

	case "B":
		return "Very Good"

	case "C":
		return "Good"

	case "D":
		return "Average"

	default:
		return "Poor"
	}
}

// Returns student's scholarship
func studentScholarship(examScore float64) string {

	switch {

	case examScore >= FullScholarshipScore:
		return "Full Scholarship"

	case examScore >= PartialScholarshipScore:
		return "50% Scholarship"

	case examScore >= GoodScore:
		return "25% Scholarship"

	default:
		return "No Scholarship"
	}
}

// Prints current subjects
func printSubjects() {
	fmt.Println("=============================")
	fmt.Println("CURRENT SUBJECTS")
	fmt.Println("=============================")

	for i := 1; i <= 5; i++ {

		switch i {

		case 1:
			fmt.Println("Subject 1: Go Programming")

		case 2:
			fmt.Println("Subject 2: Database Systems")

		case 3:
			fmt.Println("Subject 3: Computer Networks")

		case 4:
			fmt.Println("Subject 4: Operating Systems")

		case 5:
			fmt.Println("Subject 5: Full Stack Development")
		}
	}
}

// Prints final student report
func printReport(
	studentName string,
	studentID string,
	age int,
	faculty string,
	course string,
	subject string,
	attendance int,
	assignmentScore int,
	midtermScore int,
	finalExamScore int,
	examScore float64,
	grade string,
	status string,
	performance string,
	scholarship string,
) {

	fmt.Println("===================================================")
	fmt.Println("          STUDENT EXAM MANAGEMENT SYSTEM")
	fmt.Println("===================================================")

	fmt.Println("School Name       :", SchoolName)
	fmt.Println("Country           :", Country)
	fmt.Println("Academic Year     :", AcademicYear)

	fmt.Println("---------------------------------------------------")

	fmt.Println("Student Name      :", studentName)
	fmt.Println("Student ID        :", studentID)
	fmt.Println("Age               :", age)
	fmt.Println("Faculty           :", faculty)
	fmt.Println("Course            :", course)
	fmt.Println("Subject           :", subject)

	fmt.Println("---------------------------------------------------")

	fmt.Println("Attendance        :", attendance, "%")
	fmt.Println("Assignment Score  :", assignmentScore)
	fmt.Println("Midterm Score     :", midtermScore)
	fmt.Println("Final Exam Score  :", finalExamScore)
	fmt.Println("Exam Score        :", examScore)

	fmt.Println("---------------------------------------------------")

	fmt.Println("Grade             :", grade)
	fmt.Println("Status            :", status)
	fmt.Println("Performance       :", performance)
	fmt.Println("Scholarship       :", scholarship)

	fmt.Println("===================================================")

}

func main() {

	defer fmt.Println("\nProgram Finished Successfully.")

	// Get student information
	studentName, studentID, age, faculty, course, subject := studentInfo()

	// Get exam information
	attendance, assignmentScore, midtermScore, finalExamScore, examScore := studentExam()

	// Calculate grade
	grade := calculateGrade(examScore)

	// Check exam status
	status := studentStatus(examScore)

	// Check student performance
	performance := studentPerformance(grade)

	// Check scholarship
	scholarship := studentScholarship(examScore)

	// Print subjects
	printSubjects()

	fmt.Println()

	// Print final report
	printReport(
		studentName,
		studentID,
		age,
		faculty,
		course,
		subject,
		attendance,
		assignmentScore,
		midtermScore,
		finalExamScore,
		examScore,
		grade,
		status,
		performance,
		scholarship,
	)
}
