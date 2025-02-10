package load

import (
	"cow_backend/models"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const GR_COW_SELECS_COL = "CowSelecs"

const GR_GENERAL_VALUE_HOZ_COL = "GeneralValueHoz"
const GR_EBV_MILK_HOZ_COL = "ebv_milkHoz"
const GR_EBV_FAT_HOZ_COL = "evb_fatHoz"
const GR_EBV_PROTEIN_HOZ_COL = "ebv_proteinHoz"
const GR_EBV_INSEMENATION_HOZ_COL = "ebv_insemenationHoz"
const GR_EBV_SERVICE_HOZ_COL = "ebv_serviceHoz"

const GR_GENERAL_VALUE_REG_COL = "GeneralValueReg"
const GR_EBV_MILK_REG_COL = "ebv_milkReg"
const GR_EBV_FAT_REG_COL = "evb_fatReg"
const GR_EBV_PROTEIN_REG_COL = "ebv_proteinReg"
const GR_EBV_INSEMENATION_REG_COL = "ebv_insemenationReg"
const GR_EBV_SERVICE_REG_COL = "ebv_serviceReg"
const GR_EBV_SOMATIC_NUC_REG_COL = "ebv_somaticNucReg"
const GR_EBV_PRODUCTIVE_LONGLIVITY_COL = "ebv_productiveLongLivityReg"
const GR_EBV_MASTIT_COL = "ebv_mastitReg"
const GR_EBV_PROTEIN_PERCENTS_COL = "ebv_proteinPercentsReg"
const GR_EBV_FAT_PERCENTS_COL = "ebv_fatPercentsReg"

const GR_GENERAL_VALUE_REG_REL_COL = "GeneralValueReg_rel"
const GR_EBV_MILK_REG_REL_COL = "ebv_milkReg_rel"
const GR_EBV_FAT_REG_REL_COL = "ebv_fatReg_rel"
const GR_EBV_PROTEIN_REG_REL_COL = "ebv_proteinReg_rel"
const GR_EBV_INSEMENATION_REG_REL_COL = "ebv_insemenationReg_rel"
const GR_EBV_SERVICE_REG_REL_COL = "ebv_serviceReg_rel"
const GR_EBV_SOMATIC_NUC_REG_REL_COL = "ebv_somaticNucReg_rel"
const GR_EBV_PRODUCTIVE_LONGLIVITY_REG_REL_COL = "ebv_productiveLongLivityReg_rel"
const GR_EBV_MASTIT_REG_REL_COL = "ebv_mastitReg_rel"
const GR_EBV_PROTEIN_PERCENTS_REG_REL_COL = "ebv_proteinPercentsReg_rel"
const GR_EBV_FAT_PERCENTS_REG_REL_COL = "ebv_fatPercentsReg_rel"

type gradeRecord struct {
	CowSelecs uint

	EbvMilkHoz *float64
	EbvMilkReg *float64

	EbvFatHoz *float64
	EbvFatReg *float64

	EbvProteinHoz *float64
	EbvProteinReg *float64

	EbvInsemenationHoz    *float64
	EbvInsemenationReg    *float64
	EbvInsemenationRegRel *float64

	EbvServiceHoz    *float64
	EbvServiceReg    *float64
	EbvServiceRegRel *float64

	GeneralValueHoz *float64
	GeneralValueReg *float64
	HeaderIndexes   map[string]int

	EbvSomaticNucsReg         *float64
	EbvProductiveLongevityReg *float64
	EbvMastitReg              *float64
	EbvProteinPercentsReg     *float64
	EbvFatPercentsReg         *float64

	EbvSomaticNucsRegRel         *float64
	EbvProductiveLongevityRegRel *float64
	EbvMastitRegRel              *float64
	EbvProteinPercentsRegRel     *float64
	EbvFatPercentsRegRel         *float64

	GeneralValueRegRel *float64
	EbvMilkRegRel      *float64
	EbvFatRegRel       *float64
	EbvProteinRegRel   *float64
}

var gradeRecordParsers = map[string]func(*gradeRecord, []string) error{
	GR_COW_SELECS_COL: func(gr *gradeRecord, rec []string) error {
		if rec[gr.HeaderIndexes[GR_COW_SELECS_COL]] == "" {
			return errors.New("колонка CowSelecs не может быть пустой")
		}
		selecs, err := strconv.ParseUint(rec[gr.HeaderIndexes[GR_COW_SELECS_COL]], 10, 64)
		if err != nil {
			return err
		}
		gr.CowSelecs = uint(selecs)
		return nil
	},
	GR_GENERAL_VALUE_HOZ_COL: func(gr *gradeRecord, rec []string) error {
		generalValue := rec[gr.HeaderIndexes[GR_GENERAL_VALUE_HOZ_COL]]
		if generalValue == "" {
			gr.GeneralValueHoz = nil
			return nil
		}
		generalValueFloat, err := strconv.ParseFloat(generalValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + generalValue + " " + err.Error())
		}
		gr.GeneralValueHoz = &generalValueFloat
		return nil
	},
	GR_GENERAL_VALUE_REG_COL: func(gr *gradeRecord, rec []string) error {
		generalValue := rec[gr.HeaderIndexes[GR_GENERAL_VALUE_REG_COL]]
		if generalValue == "" {
			gr.GeneralValueReg = nil
			return nil
		}
		generalValueFloat, err := strconv.ParseFloat(generalValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + generalValue + " " + err.Error())
		}
		gr.GeneralValueReg = &generalValueFloat
		return nil
	},

	GR_EBV_FAT_HOZ_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_FAT_HOZ_COL]]
		if ebvValue == "" {
			gr.EbvFatHoz = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvFatHoz = &ebvFloat
		return nil
	},
	GR_EBV_FAT_REG_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_FAT_REG_COL]]
		if ebvValue == "" {
			gr.EbvFatReg = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvFatReg = &ebvFloat
		return nil
	},

	GR_EBV_MILK_HOZ_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_MILK_HOZ_COL]]
		if ebvValue == "" {
			gr.EbvMilkHoz = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvMilkHoz = &ebvFloat
		return nil
	},
	GR_EBV_MILK_REG_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_MILK_REG_COL]]
		if ebvValue == "" {
			gr.EbvMilkReg = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvMilkReg = &ebvFloat
		return nil
	},

	GR_EBV_PROTEIN_HOZ_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_PROTEIN_HOZ_COL]]
		if ebvValue == "" {
			gr.EbvProteinHoz = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvProteinHoz = &ebvFloat
		return nil
	},
	GR_EBV_PROTEIN_REG_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_PROTEIN_REG_COL]]
		if ebvValue == "" {
			gr.EbvProteinReg = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvProteinReg = &ebvFloat
		return nil
	},

	GR_EBV_INSEMENATION_HOZ_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_PROTEIN_HOZ_COL]]
		if ebvValue == "" {
			gr.EbvInsemenationHoz = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvInsemenationHoz = &ebvFloat
		return nil
	},
	GR_EBV_INSEMENATION_REG_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_INSEMENATION_REG_COL]]
		if ebvValue == "" {
			gr.EbvInsemenationReg = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvInsemenationReg = &ebvFloat
		return nil
	},

	GR_EBV_SERVICE_HOZ_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_SERVICE_HOZ_COL]]
		if ebvValue == "" {
			gr.EbvServiceHoz = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvServiceHoz = &ebvFloat
		return nil
	},
	GR_EBV_SERVICE_REG_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_SERVICE_REG_COL]]
		if ebvValue == "" {
			gr.EbvServiceReg = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvServiceReg = &ebvFloat
		return nil
	},
	GR_EBV_MILK_REG_REL_COL: func(gr *gradeRecord, rec []string) error {
		milkRel := rec[gr.HeaderIndexes[GR_EBV_MILK_REG_REL_COL]]
		if milkRel == "" {
			gr.EbvMilkRegRel = nil
			return nil
		}
		milkRelFloat, err := strconv.ParseFloat(milkRel, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + milkRel + " " + err.Error())
		}
		gr.EbvMilkRegRel = &milkRelFloat
		return nil
	},
	GR_EBV_FAT_REG_REL_COL: func(gr *gradeRecord, rec []string) error {
		fatRel := rec[gr.HeaderIndexes[GR_EBV_FAT_REG_REL_COL]]
		if fatRel == "" {
			gr.EbvFatRegRel = nil
			return nil
		}
		fatRelFloat, err := strconv.ParseFloat(fatRel, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + fatRel + " " + err.Error())
		}
		gr.EbvFatRegRel = &fatRelFloat
		return nil
	},
	GR_EBV_PROTEIN_REG_REL_COL: func(gr *gradeRecord, rec []string) error {
		proteinRel := rec[gr.HeaderIndexes[GR_EBV_PROTEIN_REG_REL_COL]]
		if proteinRel == "" {
			gr.EbvProteinRegRel = nil
			return nil
		}
		proteinRelFloat, err := strconv.ParseFloat(proteinRel, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + proteinRel + " " + err.Error())
		}
		gr.EbvProteinRegRel = &proteinRelFloat
		return nil
	},

	GR_GENERAL_VALUE_REG_REL_COL: func(gr *gradeRecord, rec []string) error {
		strVal := rec[gr.HeaderIndexes[GR_GENERAL_VALUE_REG_REL_COL]]
		if strVal == "" {
			gr.GeneralValueRegRel = nil
			return nil
		}
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return errors.New("не удалось распарсить колонку " + GR_GENERAL_VALUE_REG_REL_COL + "значение: " + strVal)
		}
		gr.GeneralValueRegRel = &floatVal
		return nil
	},

	GR_EBV_SOMATIC_NUC_REG_COL: func(gr *gradeRecord, rec []string) error {
		strVal := rec[gr.HeaderIndexes[GR_EBV_SOMATIC_NUC_REG_COL]]
		if strVal == "" {
			gr.EbvSomaticNucsReg = nil
			return nil
		}
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return errors.New("не удалось распарсить колонку " + GR_EBV_SOMATIC_NUC_REG_COL + "значение: " + strVal)
		}
		gr.EbvSomaticNucsReg = &floatVal
		return nil
	},

	GR_EBV_SOMATIC_NUC_REG_REL_COL: func(gr *gradeRecord, rec []string) error {
		strVal := rec[gr.HeaderIndexes[GR_EBV_SOMATIC_NUC_REG_REL_COL]]
		if strVal == "" {
			gr.EbvSomaticNucsRegRel = nil
			return nil
		}
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return errors.New("не удалось распарсить колонку " + GR_EBV_SOMATIC_NUC_REG_REL_COL + "значение: " + strVal)
		}
		gr.EbvSomaticNucsRegRel = &floatVal
		return nil
	},

	GR_EBV_PRODUCTIVE_LONGLIVITY_COL: func(gr *gradeRecord, rec []string) error {
		strVal := rec[gr.HeaderIndexes[GR_EBV_PRODUCTIVE_LONGLIVITY_COL]]
		if strVal == "" {
			gr.EbvProductiveLongevityReg = nil
			return nil
		}
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return errors.New("не удалось распарсить колонку " + GR_EBV_PRODUCTIVE_LONGLIVITY_COL + "значение: " + strVal)
		}
		gr.EbvProductiveLongevityReg = &floatVal
		return nil
	},

	GR_EBV_PRODUCTIVE_LONGLIVITY_REG_REL_COL: func(gr *gradeRecord, rec []string) error {
		strVal := rec[gr.HeaderIndexes[GR_EBV_PRODUCTIVE_LONGLIVITY_REG_REL_COL]]
		if strVal == "" {
			gr.EbvProductiveLongevityRegRel = nil
			return nil
		}
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return errors.New("не удалось распарсить колонку " + GR_EBV_PRODUCTIVE_LONGLIVITY_REG_REL_COL + "значение: " + strVal)
		}
		gr.EbvProductiveLongevityRegRel = &floatVal
		return nil
	},

	GR_EBV_MASTIT_COL: func(gr *gradeRecord, rec []string) error {
		strVal := rec[gr.HeaderIndexes[GR_EBV_MASTIT_COL]]
		if strVal == "" {
			gr.EbvMastitReg = nil
			return nil
		}
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return errors.New("не удалось распарсить колонку " + GR_EBV_MASTIT_COL + "значение: " + strVal)
		}
		gr.EbvMastitReg = &floatVal
		return nil
	},

	GR_EBV_MASTIT_REG_REL_COL: func(gr *gradeRecord, rec []string) error {
		strVal := rec[gr.HeaderIndexes[GR_EBV_MASTIT_REG_REL_COL]]
		if strVal == "" {
			gr.EbvMastitRegRel = nil
			return nil
		}
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return errors.New("не удалось распарсить колонку " + GR_EBV_MASTIT_REG_REL_COL + "значение: " + strVal)
		}
		gr.EbvMastitRegRel = &floatVal
		return nil
	},

	GR_EBV_SERVICE_REG_REL_COL: func(gr *gradeRecord, rec []string) error {
		strVal := rec[gr.HeaderIndexes[GR_EBV_SERVICE_REG_REL_COL]]
		if strVal == "" {
			gr.EbvServiceRegRel = nil
			return nil
		}
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return errors.New("не удалось распарсить колонку " + GR_EBV_SERVICE_REG_REL_COL + "значение: " + strVal)
		}
		gr.EbvServiceRegRel = &floatVal
		return nil
	},

	GR_EBV_INSEMENATION_REG_REL_COL: func(gr *gradeRecord, rec []string) error {
		strVal := rec[gr.HeaderIndexes[GR_EBV_INSEMENATION_REG_REL_COL]]
		if strVal == "" {
			gr.EbvInsemenationRegRel = nil
			return nil
		}
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return errors.New("не удалось распарсить колонку " + GR_EBV_INSEMENATION_REG_REL_COL + "значение: " + strVal)
		}
		gr.EbvInsemenationRegRel = &floatVal
		return nil
	},

	GR_EBV_PROTEIN_PERCENTS_COL: func(gr *gradeRecord, rec []string) error {
		strVal := rec[gr.HeaderIndexes[GR_EBV_PROTEIN_PERCENTS_COL]]
		if strVal == "" {
			gr.EbvProteinPercentsReg = nil
			return nil
		}
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return errors.New("не удалось распарсить колонку " + GR_EBV_PROTEIN_PERCENTS_COL + "значение: " + strVal)
		}
		gr.EbvProteinPercentsReg = &floatVal
		return nil
	},
	GR_EBV_PROTEIN_PERCENTS_REG_REL_COL: func(gr *gradeRecord, rec []string) error {
		strVal := rec[gr.HeaderIndexes[GR_EBV_PROTEIN_PERCENTS_REG_REL_COL]]
		if strVal == "" {
			gr.EbvProteinPercentsRegRel = nil
			return nil
		}
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return errors.New("не удалось распарсить колонку " + GR_EBV_PROTEIN_PERCENTS_REG_REL_COL + "значение: " + strVal)
		}
		gr.EbvProteinPercentsRegRel = &floatVal
		return nil
	},

	GR_EBV_FAT_PERCENTS_COL: func(gr *gradeRecord, rec []string) error {
		strVal := rec[gr.HeaderIndexes[GR_EBV_FAT_PERCENTS_COL]]
		if strVal == "" {
			gr.EbvFatPercentsReg = nil
			return nil
		}
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return errors.New("не удалось распарсить колонку " + GR_EBV_FAT_PERCENTS_COL + "значение: " + strVal)
		}
		gr.EbvFatPercentsReg = &floatVal
		return nil
	},
	GR_EBV_FAT_PERCENTS_REG_REL_COL: func(gr *gradeRecord, rec []string) error {
		strVal := rec[gr.HeaderIndexes[GR_EBV_FAT_PERCENTS_REG_REL_COL]]
		if strVal == "" {
			gr.EbvFatPercentsRegRel = nil
			return nil
		}
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return errors.New("не удалось распарсить колонку " + GR_EBV_FAT_PERCENTS_REG_REL_COL + "значение: " + strVal)
		}
		gr.EbvFatPercentsRegRel = &floatVal
		return nil
	},
}

func NewGradeRecord(header []string) (*gradeRecord, error) {
	gr := gradeRecord{}
	columns := []string{
		GR_COW_SELECS_COL,

		GR_GENERAL_VALUE_HOZ_COL,
		GR_EBV_MILK_HOZ_COL,
		GR_EBV_FAT_HOZ_COL,
		GR_EBV_PROTEIN_HOZ_COL,
		GR_EBV_INSEMENATION_HOZ_COL,
		GR_EBV_SERVICE_HOZ_COL,

		GR_GENERAL_VALUE_REG_COL,
		GR_EBV_MILK_REG_COL,
		GR_EBV_FAT_REG_COL,
		GR_EBV_PROTEIN_REG_COL,
		GR_EBV_INSEMENATION_REG_COL,
		GR_EBV_SERVICE_REG_COL,
		GR_EBV_SOMATIC_NUC_REG_COL,
		GR_EBV_PRODUCTIVE_LONGLIVITY_COL,
		GR_EBV_MASTIT_COL,
		GR_EBV_PROTEIN_PERCENTS_COL,
		GR_EBV_FAT_PERCENTS_COL,

		GR_GENERAL_VALUE_REG_REL_COL,
		GR_EBV_MILK_REG_REL_COL,
		GR_EBV_FAT_REG_REL_COL,
		GR_EBV_PROTEIN_REG_REL_COL,
		GR_EBV_INSEMENATION_REG_REL_COL,
		GR_EBV_SERVICE_REG_REL_COL,
		GR_EBV_SOMATIC_NUC_REG_REL_COL,
		GR_EBV_PRODUCTIVE_LONGLIVITY_REG_REL_COL,
		GR_EBV_MASTIT_REG_REL_COL,
		GR_EBV_PROTEIN_PERCENTS_REG_REL_COL,
		GR_EBV_FAT_PERCENTS_REG_REL_COL,
	}
	gr.HeaderIndexes = make(map[string]int)
	for idx, col := range header {
		gr.HeaderIndexes[col] = idx
	}
	for _, col := range columns {
		if _, ok := gr.HeaderIndexes[col]; !ok {
			return nil, errors.New("не найдена колонка " + col)
		}
	}
	return &gr, nil
}

func (gr *gradeRecord) FromCsvRecord(rec []string) (CsvToDbLoader, error) {
	for col, parser := range gradeRecordParsers {
		if err := parser(gr, rec); err != nil {
			return nil, errors.New("ошибка парсинга колонки " + col + " значения " + rec[gr.HeaderIndexes[col]] + ": " + err.Error())
		}
	}
	return gr, nil
}

func (gr *gradeRecord) ToDbModel(tx *gorm.DB) (any, error) {
	grCow := models.Cow{}

	db := models.GetDb()
	if err := db.Preload("GradeRegion").Preload("GradeHoz").First(&grCow, map[string]any{"selecs_number": gr.CowSelecs}).Error; err != nil {
		return nil, errors.New("Не найдена корова с селексом " + strconv.FormatUint(uint64(gr.CowSelecs), 10))
	}

	if grCow.GradeHoz != nil {
		return nil, errors.New("Для коровы с селексом " + strconv.FormatUint(uint64(gr.CowSelecs), 10) + " Уже существует оценка по хозяйству")
	}
	if grCow.GradeRegion != nil {
		return nil, errors.New("Для коровы с селексом " + strconv.FormatUint(uint64(gr.CowSelecs), 10) + " Уже существует оценка по региону")
	}

	grCow.GradeHoz = new(models.GradeHoz)
	grCow.GradeRegion = new(models.GradeRegion)
	grCow.GradeCountry = new(models.GradeCountry)

	grCow.GradeHoz.EbvFat = gr.EbvFatHoz
	grCow.GradeHoz.EbvInsemenation = gr.EbvInsemenationHoz
	grCow.GradeHoz.EbvMilk = gr.EbvMilkHoz
	grCow.GradeHoz.EbvProtein = gr.EbvProteinHoz
	grCow.GradeHoz.EbvService = gr.EbvServiceHoz
	grCow.GradeHoz.GeneralValue = gr.GeneralValueHoz

	grCow.GradeRegion.GeneralValueReliability = gr.GeneralValueRegRel
	grCow.GradeRegion.EbvFatReliability = gr.EbvFatRegRel
	grCow.GradeRegion.EbvMilkReliability = gr.EbvMilkRegRel
	grCow.GradeRegion.EbvProteinReliability = gr.EbvProteinRegRel
	grCow.GradeRegion.EbvSomaticNucsReliability = gr.EbvSomaticNucsRegRel
	grCow.GradeRegion.EbvProductiveLongevityReliability = gr.EbvProductiveLongevityRegRel
	grCow.GradeRegion.EbvMastitReliability = gr.EbvMastitRegRel
	grCow.GradeRegion.EbvServiceReliability = gr.EbvServiceRegRel
	grCow.GradeRegion.EbvInsemenationReliability = gr.EbvInsemenationRegRel
	grCow.GradeRegion.EbvProteinPercentsReliability = gr.EbvProteinPercentsRegRel
	grCow.GradeRegion.EbvFatPercentsReliability = gr.EbvFatPercentsRegRel

	grCow.GradeRegion.GeneralValue = gr.GeneralValueReg
	grCow.GradeRegion.EbvService = gr.EbvServiceReg
	grCow.GradeRegion.EbvProductiveLongevity = gr.EbvProductiveLongevityReg
	grCow.GradeRegion.EbvFat = gr.EbvFatReg
	grCow.GradeRegion.EbvInsemenation = gr.EbvInsemenationReg
	grCow.GradeRegion.EbvMilk = gr.EbvMilkReg
	grCow.GradeRegion.EbvProtein = gr.EbvProteinReg
	grCow.GradeRegion.EbvSomaticNucs = gr.EbvSomaticNucsReg
	grCow.GradeRegion.EbvMastit = gr.EbvMastitReg
	grCow.GradeRegion.EbvProteinPercents = gr.EbvProteinPercentsReg
	grCow.GradeRegion.EbvFatPercents = gr.EbvFatPercentsReg

	return grCow, nil
}

func (gr *gradeRecord) Copy() *gradeRecord {
	copy := gradeRecord{}
	copy.HeaderIndexes = gr.HeaderIndexes
	return &copy
}

const GRADE_CSV_PATH = "./csv/grades/"

var gradeUniqueIndex uint64 = 0

func (l *Load) Grade() func(*gin.Context) {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(500, err)
			return
		}
		csvField, ok := form.File["csv"]
		if !ok {
			c.JSON(500, "not found field csv")
			return
		}

		now := time.Now()
		uploadedName := GRADE_CSV_PATH + "grade_" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(gradeUniqueIndex, 16) + ".csv"
		if err := c.SaveUploadedFile(csvField[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		gradeUniqueIndex++

		file, err := os.Open(uploadedName)
		if err != nil {
			c.JSON(500, "error opening file")
			return
		}
		defer file.Close()
		csvReader := csv.NewReader(file)
		header, err := csvReader.Read()
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		recordWithHeader, err := NewGradeRecord(header)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}

		errors := []string{}
		errorsMtx := sync.Mutex{}
		loaderWg := sync.WaitGroup{}
		loadChannel := make(chan loaderData)
		MakeLoadingPool(loadChannel, SaveRecordToDb[models.Cow])
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		for record, err := csvReader.Read(); err != io.EOF; record, err = csvReader.Read() {
			if err != nil {
				errorsMtx.Lock()
				errors = append(errors, err.Error())
				errorsMtx.Unlock()
				continue
			}
			loaderWg.Add(1)
			loadChannel <- loaderData{
				Loader:    recordWithHeader.Copy(),
				Record:    record,
				Errors:    &errors,
				ErrorsMtx: &errorsMtx,
				WaitGroup: &loaderWg,
			}
		}
		loaderWg.Wait()
		close(loadChannel)
		c.JSON(200, errors)
	}
}
