package model
import(
	"time"
)
type Start struct {
	UUID    string
	Mins    int32
	EndTime time.Time
}
type Config struct {
	Category int32
	UUID    string
	Maxcur int32
	Minvol int32
	Ratevol int32
	Servicefee float32
	Parkfee float32
	Version int32
	Timefee []FeeRate
	Light []Period
}
type   Period struct {
	Start         time.Time
	End         time.Time

}
type   FeeRate struct {
	Fee              *float32
	Start            *uint32
	End              *uint32

}