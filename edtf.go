package edtf

import (
	"fmt"
	_edtf "github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/common"
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

	layout := "%04d-%02d"

	if yyyy < 0 {
		layout = "%05d-%02d"
	}

	edtf_str := fmt.Sprintf(layout, yyyy, mm)

	if circa != "" {

		layout := "%04d-%s%02d"

		if yyyy < 0 {
			layout = "%05d-%s%02d"
		}

		edtf_str = fmt.Sprintf(layout, yyyy, _edtf.APPROXIMATE, mm)
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

	layout := "%04d-%02d-%02d"

	if yyyy < 0 {
		layout = "%05d-%02d-%02d"
	}

	edtf_str := fmt.Sprintf(layout, yyyy, mm, dd)

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

	return formatWithSemestral(yyyy, 1, 4, circa)
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

	return formatWithSemestral(yyyy, 5, 8, circa)
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

	return formatWithSemestral(yyyy, 9, 12, circa)
}

func formatWithSemestral(yyyy int, start int, end int, circa string) (string, error) {

	layout := "%04d-%02d/%04d-%02d"

	if yyyy < 0 {
		layout = "%05d-%02d/%05d-%02d"
	}

	edtf_str := fmt.Sprintf(layout, yyyy, start, yyyy, end)

	if circa != "" {

		layout = "%04d-%s%02d/%04d-%s%02d"

		if yyyy < 0 {
			layout = "%05d-%s%02d/%05d-%s%02d"
		}

		edtf_str = fmt.Sprintf(layout, yyyy, _edtf.APPROXIMATE, start, yyyy, _edtf.APPROXIMATE, end)
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

	// because "~198X" is not a valid EDTF string

	if circa != "" {
		edtf_str = fmt.Sprintf("%s%s-01-01/%s%s-12-31", _edtf.APPROXIMATE, edtf_str, _edtf.APPROXIMATE, edtf_str)
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
		edtf_str = fmt.Sprintf("%s%s/%s%s", _edtf.APPROXIMATE, start, _edtf.APPROXIMATE, end)
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

	layout := "%04d"

	if yyyy < 0 {
		layout = "%05d"
	}

	edtf_str := fmt.Sprintf(layout, yyyy)

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

	yyyy, err := strconv.Atoi(str_yyyy)

	if err != nil {
		return "", errors.Invalid()
	}

	is_bce := false

	if yyyy < 0 {
		is_bce = true
	}

	if is_bce {
		yyyy = common.FlipYear(yyyy)
	}

	str_date := fmt.Sprintf("%s %s, %04d", str_m, str_d, yyyy)

	t_fmt := "Jan 2, 2006"

	t, err := time.Parse(t_fmt, str_date)

	if err != nil {
		return "", err
	}

	if is_bce {
		t = common.TimeToBCE(t)
	}

	edtf_str := t.Format("2006-01-02")
	return edtf_str, nil
}
