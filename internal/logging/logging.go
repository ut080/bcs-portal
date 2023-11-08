package logging

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogging(loglevel string, console bool) {
	if console {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	}

	lvl, err := zerolog.ParseLevel(loglevel)
	if err != nil {
		log.Error().Err(err).Str("loglevel", loglevel).Msg("failed to parse loglevel, defaulting to INFO")
		lvl = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(lvl)
}

// Panic is a wrapper around log.Panic() from github.com/rs/zerolog/log.
func Panic() *Event {
	return &Event{log.Panic()}
}

// Fatal is a wrapper around log.Fatal() from github.com/rs/zerolog/log.
func Fatal() *Event {
	return &Event{log.Fatal()}
}

// Error is a wrapper around log.Error() from github.com/rs/zerolog/log.
func Error() *Event {
	return &Event{log.Error()}
}

// Warn is a wrapper around log.Warn() from github.com/rs/zerolog/log.
func Warn() *Event {
	return &Event{log.Warn()}
}

// Info is a wrapper around log.Info() from github.com/rs/zerolog/log.
func Info() *Event {
	return &Event{log.Info()}
}

// Debug is a wrapper around log.Debug() from github.com/rs/zerolog/log.
func Debug() *Event {
	return &Event{log.Debug()}
}

// Trace is a wrapper around log.Trace() from github.com/rs/zerolog/log.
func Trace() *Event {
	return &Event{log.Trace()}
}

// DefaultLogger is a not-so-great solution to passing the logger on to libraries such as github.com/ag7if/go-files.
// TODO: Find a better way to do this.
func DefaultLogger() *zerolog.Logger {
	return &log.Logger
}

// Logger is, currently, a placeholder object. For now, there is no difference between Logger.LvlFunc()
// and logging.LvlFunc(). This might change if we start configuring subloggers.
type Logger struct{}

// DefaultLogger is a not-so-great solution to passing the logger on to libraries such as github.com/ag7if/go-files.
// TODO: Find a better way to do this.
func (l *Logger) DefaultLogger() *zerolog.Logger {
	return &log.Logger
}

func (l *Logger) Panic() *Event {
	return &Event{log.Panic()}
}

func (l *Logger) Fatal() *Event {
	return &Event{log.Fatal()}
}

func (l *Logger) Error() *Event {
	return &Event{log.Error()}
}

func (l *Logger) Warn() *Event {
	return &Event{log.Warn()}
}

func (l *Logger) Info() *Event {
	return &Event{log.Info()}
}

func (l *Logger) Debug() *Event {
	return &Event{log.Debug()}
}

func (l *Logger) Trace() *Event {
	return &Event{log.Trace()}
}

// Event is a wrapper around Event from github.com/rs/zerolog.
// The methods implemented for this wrapper represent the subset of Zerolog methods that we anticipate actually using.
// This might be expanded in the future.
type Event struct {
	*zerolog.Event
}

func (e *Event) AnErr(key string, err error) *Event {
	return &Event{e.Event.AnErr(key, err)}
}

func (e *Event) Bool(key string, b bool) *Event {
	return &Event{e.Event.Bool(key, b)}
}

func (e *Event) Bools(key string, b []bool) *Event {
	return &Event{e.Event.Bools(key, b)}
}

func (e *Event) Bytes(key string, val []byte) *Event {
	return &Event{e.Event.Bytes(key, val)}
}

func (e *Event) Dur(key string, d time.Duration) *Event {
	return &Event{e.Event.Dur(key, d)}
}

func (e *Event) Durs(key string, d []time.Duration) *Event {
	return &Event{e.Event.Durs(key, d)}
}

func (e *Event) Err(err error) *Event {
	return &Event{e.Event.Err(err)}
}

func (e *Event) Errs(key string, err []error) *Event {
	return &Event{e.Event.Errs(key, err)}
}

func (e *Event) Float32(key string, f float32) *Event {
	return &Event{e.Event.Float32(key, f)}
}

func (e *Event) Float64(key string, f float64) *Event {
	return &Event{e.Event.Float64(key, f)}
}

func (e *Event) Floats32(key string, f []float32) *Event {
	return &Event{e.Event.Floats32(key, f)}
}

func (e *Event) Floats64(key string, f []float64) *Event {
	return &Event{e.Event.Floats64(key, f)}
}

func (e *Event) Hex(key string, val []byte) *Event {
	return &Event{e.Event.Hex(key, val)}
}

func (e *Event) IPAddr(key string, ip net.IP) *Event {
	return &Event{e.Event.IPAddr(key, ip)}
}

func (e *Event) IPPrefix(key string, pfx net.IPNet) *Event {
	return &Event{e.Event.IPPrefix(key, pfx)}
}

func (e *Event) Int(key string, i int) *Event {
	return &Event{e.Event.Int(key, i)}
}

func (e *Event) Int8(key string, i int8) *Event {
	return &Event{e.Event.Int8(key, i)}
}

func (e *Event) Int16(key string, i int16) *Event {
	return &Event{e.Event.Int16(key, i)}
}

func (e *Event) Int32(key string, i int32) *Event {
	return &Event{e.Event.Int32(key, i)}
}

func (e *Event) Int64(key string, i int64) *Event {
	return &Event{e.Event.Int64(key, i)}
}

func (e *Event) Ints(key string, i []int) *Event {
	return &Event{e.Event.Ints(key, i)}
}

func (e *Event) Ints8(key string, i []int8) *Event {
	return &Event{e.Event.Ints8(key, i)}
}

func (e *Event) Ints16(key string, i []int16) *Event {
	return &Event{e.Event.Ints16(key, i)}
}

func (e *Event) Ints32(key string, i []int32) *Event {
	return &Event{e.Event.Ints32(key, i)}
}

func (e *Event) Ints64(key string, i []int64) *Event {
	return &Event{e.Event.Ints64(key, i)}
}

func (e *Event) MACAddr(key string, ha net.HardwareAddr) *Event {
	return &Event{e.Event.MACAddr(key, ha)}
}

func (e *Event) Msg(msg string) {
	e.Event.Msg(msg)
}

func (e *Event) Msgf(format string, v ...interface{}) {
	e.Event.Msgf(format, v...)
}

func (e *Event) RawJSON(key string, b []byte) *Event {
	return &Event{e.Event.RawJSON(key, b)}
}

func (e *Event) Send() {
	e.Event.Send()
}

func (e *Event) Str(key string, s string) *Event {
	return &Event{e.Event.Str(key, s)}
}

func (e *Event) Stringer(key string, val fmt.Stringer) *Event {
	return &Event{e.Event.Stringer(key, val)}
}

func (e *Event) Stringers(key string, vals []fmt.Stringer) *Event {
	return &Event{e.Event.Stringers(key, vals)}
}

func (e *Event) Strs(key string, vals []string) *Event {
	return &Event{e.Event.Strs(key, vals)}
}

func (e *Event) Time(key string, t time.Time) *Event {
	return &Event{e.Event.Time(key, t)}
}

func (e *Event) TimeDiff(key string, t time.Time, start time.Time) *Event {
	return &Event{e.Event.TimeDiff(key, t, start)}
}

func (e *Event) Times(key string, t []time.Time) *Event {
	return &Event{e.Event.Times(key, t)}
}

func (e *Event) Timestamp() *Event {
	return &Event{e.Event.Timestamp()}
}

func (e *Event) Uint(key string, i uint) *Event {
	return &Event{e.Event.Uint(key, i)}
}

func (e *Event) Uint8(key string, i uint8) *Event {
	return &Event{e.Event.Uint8(key, i)}
}

func (e *Event) Uint16(key string, i uint16) *Event {
	return &Event{e.Event.Uint16(key, i)}
}

func (e *Event) Uint32(key string, i uint32) *Event {
	return &Event{e.Event.Uint32(key, i)}
}

func (e *Event) Uint64(key string, i uint64) *Event {
	return &Event{e.Event.Uint64(key, i)}
}

func (e *Event) Uints(key string, i []uint) *Event {
	return &Event{e.Event.Uints(key, i)}
}

func (e *Event) Uints8(key string, i []uint8) *Event {
	return &Event{e.Event.Uints8(key, i)}
}

func (e *Event) Uints16(key string, i []uint16) *Event {
	return &Event{e.Event.Uints16(key, i)}
}

func (e *Event) Uints32(key string, i []uint32) *Event {
	return &Event{e.Event.Uints32(key, i)}
}

func (e *Event) Uints64(key string, i []uint64) *Event {
	return &Event{e.Event.Uints64(key, i)}
}
