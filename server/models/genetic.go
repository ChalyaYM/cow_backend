package models

type Genetic struct {
	ID                        uint                 `gorm:"primaryKey"` // ID записи о генотипировании
	CowID                     uint                 // ID коровы
	ProbeNumber               string               // Номер пробы
	BloodDate                 *DateOnly            // Дата взятия пробы крови
	ResultDate                *DateOnly            // Дата получения  результата
	InbrindingCoeffByGenotype *float64             // Коэф. инбриндинга по генотипу
	GeneticIllnessesData      []GeneticIllnessData `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Список генетических заболеваний, пустой если нет

	GtcFilePath *string
}

type GeneticIllnessData struct {
	ID        uint `gorm:"primaryKey"`
	GeneticID uint
	Status    *GeneticIllnessStatus
	StatusID  *uint // статус заболевания
	Illness   GeneticIllness
	IllnessID uint // заболевание
}

type GeneticIllness struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  // имя генетического заболевания
	Description string  // описание генетического заболевания
	OMIA        *string // Какой-то там ОМИЯ номер
}

type GeneticIllnessStatus struct {
	ID     uint `gorm:"primaryKey"`
	Status string
}
