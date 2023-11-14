package data

type Grade struct {
	Name  string
	Info  string
	Price string
}

func (g *Grade) EditName(newName string) {
	g.Name = newName
}

func (g *Grade) EditInfo(newInfo string) {
	g.Info = newInfo
}

func (g *Grade) EditPrice(newPrice string) {
	g.Price = newPrice
}

func (g *Grade) Overwrite(newGrade Grade) {
	g = &newGrade
}

// Data structure funcs
func (d Data_str) GetGradesNames() []string {
	var grades []string
	for _, grade := range d.Grades {
		grades = append(grades, grade.Name)
	}
	return grades
}

func (d Data_str) FindGradeByName(gradeName string) *Grade {
	for _, grade := range d.Grades {
		if grade.Name == gradeName {
			return grade
		}
	}
	return nil
}
