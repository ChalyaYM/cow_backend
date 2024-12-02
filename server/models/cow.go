package models

import "time"

type Cow struct {
	ID        uint      `gorm:"primaryKey" example:"1"` // ID коровы
	CreatedAt time.Time `example:"2007-01-01"`          // Время создания коровы в базе данных
	Farm      *Farm     `json:"-"`
	FarmID    *uint     `example:"1"` // ID фермы, которой корова принадлежит

	FarmGroup   Farm `json:"-"`
	FarmGroupId uint `example:"1"` // ID хозяйства, которому корова принадлежит

	Breed   Breed `json:"-"`
	BreedId uint  `example:"1"` // ID породы коровы

	Sex   Sex  `json:"-"`
	SexId uint `example:"1"` // ID пола коровы

	Father   *Cow  `json:"-"`
	FatherId *uint `example:"1"` // ID коровы отца коровы

	Mother   *Cow  `json:"-"`
	MotherId *uint `example:"1"` // ID коровы матери коровы

	Genetic *Genetic `json:",omitempty"`
	Exterior *Exterior `json:",omitempty"`
	Lactation []Lactation `json:"-"`

	InventoryNumber *string `example:"1213321"`    // Инвентарный номер коровы
	SelecsNumber    *string `example:"98989"`      // Селекс номер коровы
	RSHNNumber      *string `example:"1323323232"` // РСХН номер коровы
	Name            string  `example:"Дима"`       // Кличка коровы

	// Exterior                float64  `example:"3.14"` // Оценка экстерьера коровы, будет переделано в ID экстерьера коровы
	InbrindingCoeffByFamily *float64 `example:"3.14"` // Коэф. инбриндинга по роду

	Approved    int       `example:"1"` // Целое число, что-то для админов, чтобы подтверждать коров
	BirthDate   DateOnly  // День рождения
	DepartDate  *DateOnly // День отбытия из коровника
	DeathDate   *DateOnly // Дата смерти
	BirkingDate *DateOnly // Дата перебирковки
}
