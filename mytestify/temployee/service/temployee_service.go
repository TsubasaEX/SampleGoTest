package serivce

import "github.com/TsubasaEX/SampleGoModTest/mytestify/temployee/repo"

type EmployeeService interface {
	GetSrEmployeeNumbers() int
}

type EmployeeSerivceImpl struct {
	EmpRepo repo.EmployeeRepo
}

func (es *EmployeeSerivceImpl) GetSrEmployeeNumbers(age int) int {
	srEmps := es.EmpRepo.FindEmployeesAgeGreaterThan(age)
	return len(srEmps)
}
