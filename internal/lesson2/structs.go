package lesson2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/jmoiron/sqlx"
)

type PetSound interface {
	MakeSound()
}

func Sound(ps PetSound) {
	ps.MakeSound()
}

type Pet struct {
	Name string
	Age  int
}

func (p Pet) printName() {
	fmt.Println(p.Name)
}

func (p *Pet) ChangeName(newName string) {
	p.Name = newName
}

func (p *Pet) GetName() string {
	return p.Name
}

type Dog struct {
	Pet
	Breed string
	dog   *Dog
}

// MakeSound Dog реализует интерфейс PetSound
func (d Dog) MakeSound() {
	fmt.Println("Gav")
}

// func (d Dog) PrintName() {
// 	fmt.Println("MY NAME IS ", d.Name)
// }

type App struct {
	Db     sqlx.DB
	client *http.Client
}

type DayOfWeek int

func HandleDayOfWeek(dw DayOfWeek) {

}

// Объявляем свой тип Dogs, который представляет из себя срез Dog
type Dogs []Dog

// Создаем функцию объекта Dogs, которая вернет срез объектов Dog, подходящих по фильтру age
func (d Dogs) GetDogsByAge(age int) Dogs {
	result := make(Dogs, 0)
	for _, dog := range d {
		if dog.Pet.Age == age {
			result = append(result, dog)
		}
	}
	return result
}

// Реализация интерфейса сортировки
func (d Dogs) Len() int           { return len(d) }
func (d Dogs) Less(i, j int) bool { return d[i].Pet.Age < d[j].Pet.Age }
func (d Dogs) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func Run() {
	t := int(64)
	HandleDayOfWeek(DayOfWeek(t))

	// создание объекта структуры
	// p := Pet{"Sharik", 15}
	p := Pet{Name: "Sharik", Age: 15}

	sharikDog := &Dog{
		Breed: "Немецкая овчарка",
	}
	sharikDog.Pet.Age = p.Age
	sharikDog.Pet.Name = p.Name

	vasyaDog := &Dog{
		Breed: "Немецкая овчарка",
	}
	vasyaDog.Pet.Age = 21
	vasyaDog.Pet.Name = p.Name

	vasyaDog.Pet.printName()
	fmt.Println(sharikDog)

	dogs := Dogs{*sharikDog, *vasyaDog}
	fmt.Println(dogs)
	// sort.Sort(dogs)

	// Сортировка среза без реализации интерфейса
	sort.Slice(dogs, func(i, j int) bool {
		return dogs[i].Pet.Age > dogs[j].Pet.Age
	})
	fmt.Println(dogs)

	fmt.Println(dogs.GetDogsByAge(1))
}

type KeyValueStruct struct {
	// Экспортируемое поле Key.
	// Если напишем с маленькой буквы, то поле будет недоступно
	Key string `json:"key"`
}

// RunInterface interface argument
func RunInterface(v interface{}) {

	// попытка приведения v к int
	if d, ok := v.(int); ok {
		fmt.Println(d)
	} else {
		fmt.Printf("%v is not int\n", v)
	}

	kvstruct := KeyValueStruct{}

	js := `{"key": "value"}`
	mp := make(map[string]interface{})
	if err := json.Unmarshal([]byte(js), &mp); err != nil {
		fmt.Println(err)
		return
	}

	if err := json.Unmarshal([]byte(js), &kvstruct); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(mp)
	fmt.Printf("%v", kvstruct)

}
