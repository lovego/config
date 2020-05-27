package conf

import "errors"

type Duration int

func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	integer, err := ParseDuration(s)
	if err != nil {
		return err
	}
	*d = Duration(integer)
	return nil
}

func ParseDuration(s string) (int, error) {
	var result int
	var now int
	for _, b := range s {
		if b >= '0' && b <= '9' {
			now = now*10 + int(b-'0')
		} else {
			switch b {
			case 'y', 'Y':
				result += now * 365 * 86400
				now = 0
			case 'M':
				result += now * 30 * 86400
				now = 0
			case 'd', 'D':
				result += now * 86400
				now = 0
			case 'h', 'H':
				result += now * 3600
				now = 0
			case 'm':
				result += now * 60
				now = 0
			case 's', 'S':
				result += now
				now = 0
			default:
				return 0, errors.New("invalid duration: " + s)
			}
		}
	}
	return result + now, nil
}
