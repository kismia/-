package logruspflag

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type value logrus.Level

func newValue(val logrus.Level, p *logrus.Level) *value {
	*p = val
	return (*value)(p)
}

func (i *value) String() string {
	return logrus.Level(*i).String()
}

func (i *value) Set(s string) error {
	level, err := logrus.ParseLevel(strings.TrimSpace(s))
	if err != nil {
		return fmt.Errorf("failed to parse logrus level: %q", s)
	}
	*i = value(level)
	return nil
}

func (i *value) Type() string {
	return "string"
}

func LevelVar(f *pflag.FlagSet, p *logrus.Level, name string, value logrus.Level, usage string) {
	f.VarP(newValue(value, p), name, "", usage)
}

func LevelVarP(f *pflag.FlagSet, p *logrus.Level, name, shorthand string, value logrus.Level, usage string) {
	f.VarP(newValue(value, p), name, shorthand, usage)
}
