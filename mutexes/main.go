package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type member string

const (
	_father = "father"
	_mother = "mother"
	_son    = "son"
)

type Person struct {
	Name string
	Age  uint64

	mu sync.Mutex
}

type Family struct {
	Father *Person
	Mother *Person
	Son    *Person

	mu sync.Mutex
}

func main() {
	jurandir := Person{Name: "Jurandir", Age: 32}
	jecineia := Person{Name: "Jecineia", Age: 35}
	julyscleydy := Person{Name: "Juslyscleydy", Age: 12}

	robersons := Family{
		Father: &jurandir,
		Mother: &jecineia,
		Son:    &julyscleydy,
	}

	// fmt.Printf("%+v\n", robersons.Father)
	// fmt.Printf("%s\t-\t%s\t-\t%s\n", robersons.Father.Name, robersons.Mother.Name, robersons.Son.Name)
	fmt.Printf("%s: %d\t-\t%s: %d\t-\t%s: %d\n",
		robersons.Father.Name, robersons.Father.Age,
		robersons.Mother.Name, robersons.Mother.Age,
		robersons.Son.Name, robersons.Son.Age,
	)

	// testFatherReplacement(robersons)
	// testFamilyMemberUpdateByPersonMutex(&robersons)
	testFamilyMemberUpdateByAtomic(&robersons)
	// testFamilyMemberUpdateByFamilyMutex(&robersons)

	// fmt.Printf("%s\t-\t%s\t-\t%s\n", robersons.Father.Name, robersons.Mother.Name, robersons.Son.Name)
	fmt.Printf("%s: %d\t-\t%s: %d\t-\t%s: %d\n",
		robersons.Father.Name, robersons.Father.Age,
		robersons.Mother.Name, robersons.Mother.Age,
		robersons.Son.Name, robersons.Son.Age,
	)

}

func (f *Family) updateFamilyPersonAgeByAtomic(person member, ageDelta uint64) {
	switch person {
	case _father:
		fmt.Println("atomic operation by father")
		fmt.Printf("Updating father's age from %d to add %d\n", f.Father.Age, ageDelta)
		atomic.AddUint64(&f.Father.Age, ageDelta)
	case _mother:
		fmt.Println("atomic operation by mother")
		fmt.Printf("Updating father's age from %d to add %d\n", f.Mother.Age, ageDelta)
		atomic.AddUint64(&f.Mother.Age, ageDelta)
	case _son:
		fmt.Println("atomic operation by son")
		fmt.Printf("Updating father's age from %d to add %d\n", f.Son.Age, ageDelta)
		atomic.AddUint64(&f.Son.Age, ageDelta)
		fmt.Println("son is holding for 5sec")
		time.Sleep(5 * time.Second)
	}
	fmt.Printf("leaving operation entered by %s\n", person)
}

func (f *Family) updateFamilyPersonNameByPersonMutex(person member, newName string) {
	switch person {
	case _father:
		f.Father.mu.Lock()
		fmt.Println("mutex locked by father")
		defer f.Father.mu.Unlock()
		fmt.Printf("Updating father's name from %s to %s\n", f.Father.Name, newName)
		f.Father.Name = newName
	case _mother:
		f.Mother.mu.Lock()
		fmt.Println("mutex locked by mother")
		defer f.Mother.mu.Unlock()
		fmt.Printf("Updating mother's name from %s to %s\n", f.Mother.Name, newName)
		f.Mother.Name = newName
	case _son:
		f.Son.mu.Lock()
		fmt.Println("mutex locked by son")
		defer f.Son.mu.Unlock()
		fmt.Printf("Updating son's name from %s to %s\n", f.Son.Name, newName)
		f.Son.Name = newName
		fmt.Println("son is holding mutex for 5sec")
		time.Sleep(5 * time.Second)
	}
	fmt.Printf("unlocking mutex acquired by %s\n", person)
}

func (f *Family) updateFamilyPersonNameByFamilyMutex(person member, newName string) {
	switch person {
	case _father:
		f.mu.Lock()
		fmt.Println("mutex locked by father")
		defer f.mu.Unlock()
		fmt.Printf("Updating father's name from %s to %s\n", f.Father.Name, newName)
		f.Father.Name = newName
	case _mother:
		f.mu.Lock()
		fmt.Println("mutex locked by mother")
		defer f.mu.Unlock()
		fmt.Printf("Updating mother's name from %s to %s\n", f.Mother.Name, newName)
		f.Mother.Name = newName
	case _son:
		f.mu.Lock()
		fmt.Println("mutex locked by son")
		defer f.mu.Unlock()
		fmt.Printf("Updating son's name from %s to %s\n", f.Son.Name, newName)
		f.Son.Name = newName
		fmt.Println("son is holding mutex for 5sec")
		time.Sleep(5 * time.Second)
	}
	fmt.Printf("unlocking mutex acquired by %s\n", person)
}

func (f *Family) changeFamilyPerson(person member, newPerson *Person) {
	fmt.Printf("Updating %s from %s to %s\n", person, f.Father.Name, newPerson.Name)
	switch person {
	case _father:
		f.Father = newPerson
	case _mother:
		f.Mother = newPerson
	case _son:
		f.Son = newPerson
	}
}

func testFamilyMemberUpdateByAtomic(family *Family) {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		family.updateFamilyPersonAgeByAtomic(_father, 1)
	}()

	go func() {
		defer wg.Done()
		family.updateFamilyPersonAgeByAtomic(_mother, 2)
	}()

	go func() {
		defer wg.Done()
		family.updateFamilyPersonAgeByAtomic(_son, 3)
	}()
	wg.Wait()
}

func testFamilyMemberUpdateByFamilyMutex(family *Family) {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		family.updateFamilyPersonNameByFamilyMutex(_father, "Josenval")
	}()

	go func() {
		defer wg.Done()
		family.updateFamilyPersonNameByFamilyMutex(_mother, "Dorotéia")
	}()

	go func() {
		defer wg.Done()
		family.updateFamilyPersonNameByFamilyMutex(_son, "Rogismar")
	}()
	wg.Wait()
}

func testFamilyMemberUpdateByPersonMutex(family *Family) {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		family.updateFamilyPersonNameByPersonMutex(_father, "Josenval")
	}()

	go func() {
		defer wg.Done()
		family.updateFamilyPersonNameByPersonMutex(_mother, "Dorotéia")
	}()

	go func() {
		defer wg.Done()
		family.updateFamilyPersonNameByPersonMutex(_son, "Rogismar")
	}()
	wg.Wait()
}

func testFatherReplacement(family *Family) {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		family.changeFamilyPerson(_father, &Person{
			Name: "Josenval",
			Age:  18,
		})
	}()

	go func() {
		defer wg.Done()
		family.changeFamilyPerson(_father, &Person{
			Name: "Robério",
			Age:  22,
		})
	}()

	go func() {
		defer wg.Done()
		family.changeFamilyPerson(_father, &Person{
			Name: "Rogismar",
			Age:  22,
		})
	}()
	wg.Wait()
}
