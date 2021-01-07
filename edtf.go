package edtf

import (
	"fmt"
	_edtf "github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/parser"
	"github.com/sfomuseum/go-sfomuseum-edtf/errors"
	"github.com/sfomuseum/go-sfomuseum-edtf/re"
	"strconv"
	"time"
)

func ToEDTFDate(raw string) (*_edtf.EDTFDate, error) {

	edtf_str, err := ToEDTFString(raw)

	if err != nil {
		return nil, err
	}

	return parser.ParseString(edtf_str)
}

func ToEDTFString(raw string) (string, error) {

	if re.MY.MatchString(raw) {
		return EDTFStringFromMY(raw)
	}

	if re.MDY.MatchString(raw) {
		return EDTFStringFromMDY(raw)
	}

	if re.EARLY.MatchString(raw) {
		return EDTFStringFromEARLY(raw)
	}

	if re.MID.MatchString(raw) {
		return EDTFStringFromMID(raw)
	}

	if re.LATE.MatchString(raw) {
		return EDTFStringFromLATE(raw)
	}

	if re.DECADE.MatchString(raw) {
		return EDTFStringFromDECADE(raw)
	}

	if re.RANGE.MatchString(raw) {
		return EDTFStringFromRANGE(raw)
	}

	if re.YYYY.MatchString(raw) {
		return EDTFStringFromYYYY(raw)
	}

	if re.MDYHUMAN.MatchString(raw) {
		return EDTFStringFromMDYHUMAN(raw)
	}

	return "", errors.Invalid()
}

func EDTFStringFromMY(raw string) (string, error) {

	m := re.MY.FindStringSubmatch(raw)

	if len(m) != 4 {
		return "", errors.Invalid()
	}

	circa := m[1]
	str_mm := m[2]
	str_yyyy := m[3]

	mm, err := strconv.Atoi(str_mm)

	if err != nil {
		return "", errors.Invalid()
	}

	yyyy, err := strconv.Atoi(str_yyyy)

	if err != nil {
		return "", errors.Invalid()
	}

	edtf_str := fmt.Sprintf("%04d-%02d", yyyy, mm)

	if circa != "" {
		edtf_str = edtf_str + _edtf.APPROXIMATE
	}

	return edtf_str, nil
}

func EDTFStringFromMDY(raw string) (string, error) {

	m := re.MDY.FindStringSubmatch(raw)

	if len(m) != 5 {
		return "", errors.Invalid()
	}

	circa := m[1]

	str_mm := m[2]
	str_dd := m[3]
	str_yyyy := m[4]

	dd, err := strconv.Atoi(str_dd)

	if err != nil {
		return "", errors.Invalid()
	}

	mm, err := strconv.Atoi(str_mm)

	if err != nil {
		return "", errors.Invalid()
	}

	yyyy, err := strconv.Atoi(str_yyyy)

	if err != nil {
		return "", errors.Invalid()
	}

	edtf_str := fmt.Sprintf("%04d-%02d-%02d", yyyy, mm, dd)

	if circa != "" {

		edtf_str = edtf_str + _edtf.APPROXIMATE
	}

	return edtf_str, nil
}

func EDTFStringFromEARLY(raw string) (string, error) {

	m := re.EARLY.FindStringSubmatch(raw)

	if len(m) != 3 {
		return "", errors.Invalid()
	}

	circa := m[1]
	str_yyyy := m[2]

	yyyy, err := strconv.Atoi(str_yyyy)

	if err != nil {
		return "", errors.Invalid()
	}

	edtf_str := fmt.Sprintf("%04d-01/%04d-04", yyyy, yyyy)

	if circa != "" {
		edtf_str = fmt.Sprintf("%04d-01%s/%04d-04%s", yyyy, _edtf.APPROXIMATE, yyyy, _edtf.APPROXIMATE)
	}

	return edtf_str, nil
}

func EDTFStringFromMID(raw string) (string, error) {

	m := re.MID.FindStringSubmatch(raw)

	if len(m) != 3 {
		return "", errors.Invalid()
	}

	circa := m[1]
	str_yyyy := m[2]

	yyyy, err := strconv.Atoi(str_yyyy)

	if err != nil {
		return "", errors.Invalid()
	}

	edtf_str := fmt.Sprintf("%04d-05/%04d-08", yyyy, yyyy)

	if circa != "" {
		edtf_str = fmt.Sprintf("%04d-05%s/%04d-08%s", yyyy, _edtf.APPROXIMATE, yyyy, _edtf.APPROXIMATE)
	}

	return edtf_str, nil
}

func EDTFStringFromLATE(raw string) (string, error) {

	m := re.LATE.FindStringSubmatch(raw)

	if len(m) != 3 {
		return "", errors.Invalid()
	}

	circa := m[1]
	str_yyyy := m[2]

	yyyy, err := strconv.Atoi(str_yyyy)

	if err != nil {
		return "", errors.Invalid()
	}

	edtf_str := fmt.Sprintf("%04d-09/%04d-12", yyyy, yyyy)

	if circa != "" {
		edtf_str = fmt.Sprintf("%04d-09%s/%04d-12%s", yyyy, _edtf.APPROXIMATE, yyyy, _edtf.APPROXIMATE)
	}

	return edtf_str, nil
}

func EDTFStringFromDECADE(raw string) (string, error) {

	m := re.DECADE.FindStringSubmatch(raw)

	if len(m) != 3 {
		return "", errors.Invalid()
	}

	circa := m[1]
	str_yyy := m[2]

	edtf_str := fmt.Sprintf("%sX", str_yyy)

	if circa != "" {
		edtf_str = edtf_str + _edtf.APPROXIMATE
	}

	return edtf_str, nil
}

func EDTFStringFromRANGE(raw string) (string, error) {

	m := re.RANGE.FindStringSubmatch(raw)

	if len(m) != 4 {
		return "", errors.Invalid()
	}

	circa := m[1]
	start := m[2]
	end := m[3]

	edtf_str := fmt.Sprintf("%s/%s", start, end)

	if circa != "" {
		edtf_str = fmt.Sprintf("%s%s/%s%s", start, _edtf.APPROXIMATE, end, _edtf.APPROXIMATE)
	}

	return edtf_str, nil
}

func EDTFStringFromYYYY(raw string) (string, error) {

	m := re.YYYY.FindStringSubmatch(raw)

	if len(m) != 4 {
		return "", errors.Invalid()
	}

	circa := m[1]

	before := m[2]
	str_yyyy := m[3]

	yyyy, err := strconv.Atoi(str_yyyy)

	if err != nil {
		return "", errors.Invalid()
	}

	edtf_str := fmt.Sprintf("%04d", yyyy)

	if before != "" {
		edtf_str = fmt.Sprintf("%s/%s", _edtf.UNKNOWN, edtf_str)
	}

	if circa != "" {
		edtf_str = edtf_str + _edtf.APPROXIMATE
	}

	return edtf_str, nil
}

func EDTFStringFromMDYHUMAN(raw string) (string, error) {

	m := re.MDYHUMAN.FindStringSubmatch(raw)

	if len(m) != 4 {
		return "", errors.Invalid()
	}

	str_m := m[1]
	str_d := m[2]
	str_yyyy := m[3]

	str_date := fmt.Sprintf("%s %s, %s", str_m, str_d, str_yyyy)

	t_fmt := "Jan 2, 2006"
	
	t, err := time.Parse(t_fmt, str_date)

	if err != nil {
		return "", err
	}

	edtf_str := t.Format("2006-01-02")

	return edtf_str, nil
}
