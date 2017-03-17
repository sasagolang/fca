package logic
import(
	"time"
	"fca/libs"
	protobuf "github.com/golang/protobuf/proto"
	"fca/model/proto"

)
func Start_cmd(mins int32, time time.Time, ch chan *libs.Proto) {
	start1 := proto.START{}
	start1.Mins = protobuf.Int32(30)
	start1.Time = time2TIME(time)
	p := &libs.Proto{}
	p.CMD = 4
	p.Content, _ = protobuf.Marshal(&start1)
	ch <- p

}
func Stop_cmd(ch chan *libs.Proto) {
	p := &libs.Proto{}
	p.CMD = 5
	p.Content = make([]byte, 0)
	ch <- p
}
func  DownConfig_cmd(config  *proto.CONFIG,ch chan *libs.Proto)  {

	p := &libs.Proto{}
	p.CMD = 8
	p.Content, _ = protobuf.Marshal(config)
	ch <- p

}
func StartLight_cmd(ch chan *libs.Proto){
	p := &libs.Proto{}
	p.CMD = 5
	p.Content = make([]byte, 0)
	ch <- p
}
func StopLight_cmd(ch chan *libs.Proto){
	p := &libs.Proto{}
	p.CMD = 5
	p.Content = make([]byte, 0)
	ch <- p
}