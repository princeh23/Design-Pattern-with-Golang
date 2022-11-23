package main

import "fmt"

// Mario
// 参考文章：https://juejin.cn/post/6992067357170466852
/*
马里奥状态有小马里奥（Small Mario）、超级马里奥（Super Mario）、斗篷马里奥（Cape Mario）;
小马里奥吃了蘑菇变为超级马里奥;
小马里奥和超级马里奥获得斗篷变成斗篷马里奥;
超级马里奥和斗篷马里奥碰到怪物变成小马里奥。
*/
// 改进：避免创建过多临时对象，Mario中包含所有！状态值！
type Mario struct {
	score           int64
	smallMarioState MarioState
	superMarioState MarioState
	capeMarioState  MarioState

	currentState MarioState
}

func newMario() *Mario {
	m := &Mario{}
	m.score = 0
	m.smallMarioState = &SmallMarioState{
		mario: m,
	}
	m.superMarioState = &SuperMarioState{
		mario: m,
	}
	m.capeMarioState = &CapeMarioState{
		mario: m,
	}
	m.currentState = m.smallMarioState
	return m
}

func (m *Mario) setState(s MarioState) {
	m.currentState = s
}

type MarioState interface {
	Name()
	ObtainMushroom()
	ObtainCape()
	MeetMonster()
	SetMario(mario *Mario)
}

type SmallMarioState struct {
	mario *Mario
}

func (s *SmallMarioState) SetMario(mario *Mario) {
	s.mario = mario
}

func (s *SmallMarioState) Name() {
	fmt.Println("小马里奥")
}

func (s *SmallMarioState) ObtainMushroom() {
	s.mario.setState(s.mario.superMarioState)
	s.mario.score += 100
}

func (s *SmallMarioState) ObtainCape() {
	s.mario.setState(s.mario.capeMarioState)
	s.mario.score += 200
}

func (s *SmallMarioState) MeetMonster() {
	s.mario.score -= 100
}

type SuperMarioState struct {
	mario *Mario
}

func (s *SuperMarioState) SetMario(mario *Mario) {
	s.mario = mario
}

func (s *SuperMarioState) Name() {
	fmt.Println("超级马里奥")
}

func (s *SuperMarioState) ObtainMushroom() {

}

func (s *SuperMarioState) ObtainCape() {
	s.mario.setState(s.mario.capeMarioState)
	s.mario.score += 200
}

func (s *SuperMarioState) MeetMonster() {
	s.mario.setState(s.mario.smallMarioState)
	s.mario.score -= 200
}

type CapeMarioState struct {
	mario *Mario
}

func (c *CapeMarioState) SetMario(mario *Mario) {
	c.mario = mario
}

func (c *CapeMarioState) Name() {
	fmt.Println("斗篷马里奥")
}

func (c *CapeMarioState) ObtainMushroom() {

}

func (c *CapeMarioState) ObtainCape() {

}

func (c *CapeMarioState) MeetMonster() {
	c.mario.setState(c.mario.smallMarioState)
	c.mario.score -= 200
}

func main() {
	mario := newMario()
	mario.currentState.SetMario(mario)

	mario.currentState.Name()
	fmt.Println("-------------------获得蘑菇")
	mario.currentState.ObtainMushroom()

	mario.currentState.Name()
	fmt.Println("-------------------获得斗篷")
	mario.currentState.ObtainCape()

	mario.currentState.Name()
	fmt.Println("-------------------遇到怪兽")
	mario.currentState.MeetMonster()

	mario.currentState.Name()
}
