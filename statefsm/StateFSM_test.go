package statefsm

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// 电风扇
type ElectricFan struct {
	*FSM
}

// 实例化电风扇
func NewElectricFan(initState string) *ElectricFan {
	return &ElectricFan{
		FSM: NewFSM(initState),
	}
}

var (
	PowerOff   = string("关闭")
	FirstGear  = string("1档")
	SecondGear = string("2档")
	ThirdGear  = string("3档")

	//PowerOffEvent   = FSMEvent{EventName: "按下关闭按钮", EventState: PowerOff}
	//FirstGearEvent  = FSMEvent{EventName: "按下1档按钮", EventState: FirstGear}
	//SecondGearEvent = FSMEvent{EventName: "按下2档按钮", EventState: SecondGear}
	//ThirdGearEvent  = FSMEvent{EventName: "按下3档按钮", EventState: ThirdGear}

	PowerOffEvent   = "按下关闭按钮"
	FirstGearEvent  = "按下1档按钮"
	SecondGearEvent = "按下2档按钮"
	ThirdGearEvent  = "按下3档按钮"
)

func TestOut(t *testing.T) {
	log.Info(time.Now().Unix())
}
func TestFSM(t *testing.T) {
	efan := NewElectricFan(PowerOff)                                                                      // 初始状态是关闭的
	efan.AddHandler(PowerOff, []string{PowerOffEvent, FirstGearEvent, SecondGearEvent, ThirdGearEvent})   // 关闭状态
	efan.AddHandler(FirstGear, []string{PowerOffEvent, FirstGearEvent, SecondGearEvent, ThirdGearEvent})  // 1档状态
	efan.AddHandler(SecondGear, []string{PowerOffEvent, FirstGearEvent, SecondGearEvent, ThirdGearEvent}) // 2档状态
	efan.AddHandler(ThirdGear, []string{PowerOffEvent, FirstGearEvent, SecondGearEvent, ThirdGearEvent})  // 3档状态
	// 开始测试状态变化
	err := efan.Call(ThirdGearEvent, FSMHandler(func() (string, error) {
		log.Println("电风扇开启3档，发型被吹乱了！")
		return ThirdGear, nil
	})) // 按下3档按钮
	assert.Nil(t, err)
	err = efan.Call(FirstGearEvent, FSMHandler(func() (string, error) {
		log.Println("电风扇开启1档，微风徐来！")
		return FirstGear, nil
	})) // 按下1档按钮
	assert.Nil(t, err)
	err = efan.Call(PowerOffEvent, FSMHandler(func() (string, error) {
		log.Println("电风扇已关闭")
		return PowerOff, nil
	})) // 按下关闭按钮
	assert.Nil(t, err)
	err = efan.Call(SecondGearEvent, FSMHandler(func() (string, error) {
		log.Println("电风扇开启2档，凉飕飕！")
		return SecondGear, nil
	})) // 按下2档按钮
	assert.Nil(t, err)
	err = efan.Call(PowerOffEvent, FSMHandler(func() (string, error) {
		log.Println("电风扇已关闭")
		return PowerOff, nil
	})) // 按下关闭按钮
	assert.Nil(t, err)
}

func TestFSM2(t *testing.T) {
	efan := NewElectricFan(PowerOff)                                                     // 初始状态是关闭的
	efan.AddHandler(PowerOff, []string{FirstGearEvent})                                  // 关闭状态
	efan.AddHandler(FirstGear, []string{PowerOffEvent, SecondGearEvent})                 // 1档状态
	efan.AddHandler(SecondGear, []string{PowerOffEvent, FirstGearEvent, ThirdGearEvent}) // 2档状态
	efan.AddHandler(ThirdGear, []string{PowerOffEvent, SecondGearEvent})                 // 3档状态
	// 开始测试状态变化
	err := efan.Call(ThirdGearEvent, FSMHandler(func() (string, error) {
		log.Println("电风扇开启3档，发型被吹乱了！")
		return ThirdGear, nil
	})) // 按下3档按钮
	assert.NotNil(t, err)
	assert.EqualValues(t, efan.currentState, PowerOff)

	err = efan.Call(FirstGearEvent, FSMHandler(func() (string, error) {
		log.Println("电风扇开启1档，微风徐来！")
		return FirstGear, nil
	})) // 按下1档按钮
	assert.Nil(t, err)

	err = efan.Call(ThirdGearEvent, FSMHandler(func() (string, error) {
		log.Println("电风扇开启3档，发型被吹乱了！")
		return ThirdGear, nil
	})) // 按下3档按钮
	assert.NotNil(t, err)

	err = efan.Call(SecondGearEvent, FSMHandler(func() (string, error) {
		log.Println("电风扇开启2档，凉飕飕！")
		return SecondGear, nil
	})) // 按下2档按钮
	assert.Nil(t, err)

	err = efan.Call(ThirdGearEvent, FSMHandler(func() (string, error) {
		log.Println("电风扇开启3档，发型被吹乱了！")
		return ThirdGear, nil
	})) // 按下3档按钮
	assert.Nil(t, err)
	assert.EqualValues(t, efan.currentState, ThirdGear)

	err = efan.Call(PowerOffEvent, FSMHandler(func() (string, error) {
		log.Println("电风扇已关闭")
		return PowerOff, nil
	})) // 按下关闭按钮
	assert.Nil(t, err)

	err = efan.Call(PowerOffEvent, FSMHandler(func() (string, error) {
		log.Println("电风扇已关闭")
		return PowerOff, nil
	})) // 按下关闭按钮
	assert.NotNil(t, err)
}
