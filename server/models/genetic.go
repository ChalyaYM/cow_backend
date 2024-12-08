package models

type Genetic struct {
	ID                        uint             `gorm:"primaryKey"` // ID записи о генотипировании
	CowID                     uint             // ID коровы
	ProbeNumber               string           // Номер пробы
	BloodDate                 *DateOnly         // Дата взятия пробы крови
	ResultDate                *DateOnly         // Дата получения  результата
	InbrindingCoeffByGenotype *float64          // Коэф. инбриндинга по генотипу
	GeneticIllnesses          []GeneticIllness `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;many2many:genetic_genetic_illnesses;"` // Список генетических заболеваний, пустой если нет

	GtcFilePath *string
}
type GeneticIllness struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  // имя генетического заболевания
	Description string  // описание генетического заболевания
	OMIA        *string // Какой-то там ОМИЯ номер	
	Status 		*GeneticIllnessStatus // статус этой фигни (болен, не болен ...)
	StatusID    *uint
}

type GeneticIllnessStatus struct {
	ID uint `gorm:"primaryKey"`
	Status string 
}
