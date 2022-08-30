package sap

import (
	"encoding/json"
	"strconv"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) MarshalSchema() string {
	return d.Time.Format("2006-01-02")
}

func (d Date) IsEmpty() bool {
	return d.Time.IsZero()
}

type DateTime struct {
	time.Time
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	f := d.Time.Format("20060102")
	i, err := strconv.Atoi(f)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(i)
}

func (d *Date) UnmarshalJSON(text []byte) (err error) {
	var i int
	err = json.Unmarshal(text, &i)
	if err == nil && i != 0 {
		d.Time, err = time.Parse("20060102", strconv.Itoa(i))
		return err
	}

	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02T15:04:05", value)
	return err
}

func (d DateTime) MarshalSchema() string {
	return d.Time.Format("2006-01-02T15:04:05.000-07:00")
}

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	if dt.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(dt.Time.Format("2006-01-02T15:04:05.000-0700"))
}

func (dt *DateTime) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	dt.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	dt.Time, err = time.Parse("2006-01-02T15:04:05.000-0700", value)
	return err
}

type Number float64

func (i *Number) UnmarshalJSON(data []byte) error {
	realNumber := 0.0
	err := json.Unmarshal(data, &realNumber)
	if err == nil {
		*i = Number(realNumber)
		return nil
	}

	// error, so maybe it isn't an int
	str := ""
	err = json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	if str == "" {
		*i = 0
		return nil
	}

	realNumber, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}

	i2 := Number(realNumber)
	*i = i2
	return nil
}

func (i Number) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(i))
}
