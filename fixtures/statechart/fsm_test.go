package main

import (
	"fmt"
	"github.com/looplab/fsm"
	"strings"
	"testing"
)

func TestFsm(t *testing.T) {
	fsmInst := fsm.NewFSM(
		"idle",
		fsm.Events{
			{Name: "scan", Src: []string{"idle"}, Dst: "scanning"},
			{Name: "working", Src: []string{"scanning"}, Dst: "scanning"},
			{Name: "situation", Src: []string{"scanning"}, Dst: "scanning"},
			{Name: "situation", Src: []string{"idle"}, Dst: "idle"},
			{Name: "finish", Src: []string{"scanning"}, Dst: "idle"},
		},
		fsm.Callbacks{
			"scan": func(e *fsm.Event) {
				fmt.Println("after_scan: " + e.FSM.Current())
			},
			"working": func(e *fsm.Event) {
				fmt.Println("working: " + e.FSM.Current())
			},
			"situation": func(e *fsm.Event) {
				fmt.Println("situation: " + e.FSM.Current())
			},
			"finish": func(e *fsm.Event) {
				fmt.Println("finish: " + e.FSM.Current())
			},
		},
	)

	fmt.Println(fsmInst.Current())

	err := fsmInst.Event("scan")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("1:" + fsmInst.Current())

	err = fsmInst.Event("working")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("2:" + fsmInst.Current())

	err = fsmInst.Event("situation")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("3:" + fsmInst.Current())

	err = fsmInst.Event("finish")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("4:" + fsmInst.Current())
}

func TestSimpleFsm(t *testing.T) {
	inst := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)

	fmt.Println(inst.Current())

	err := inst.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(inst.Current())

	err = inst.Event("close")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(inst.Current())
}


type Door struct {
	To  string
	FSM *fsm.FSM
}

func NewDoor(to string) *Door {
	d := &Door{
		To: to,
	}

	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { d.enterState(e) },
		},
	)

	return d
}

func (d *Door) enterState(e *fsm.Event) {
	fmt.Printf("The door to %s is %s\n", d.To, e.Dst)
}

func TestFsmInStruct(t *testing.T) {
	door := NewDoor("heaven")

	err := door.FSM.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	err = door.FSM.Event("close")
	if err != nil {
		fmt.Println(err)
	}
}

func TestMermaidOutput(t *testing.T) {
	fsmUnderTest := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
			{Name: "part-close", Src: []string{"intermediate"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)

	got, err := fsm.VisualizeForMermaidWithGraphType(fsmUnderTest, fsm.StateDiagram)
	if err != nil {
		t.Errorf("got error for visualizing with type MERMAID: %s", err)
	}
	wanted := `
stateDiagram
    [*] --> closed
    closed --> open: open
    intermediate --> closed: part-close
    open --> closed: close
`
	normalizedGot := strings.ReplaceAll(got, "\n", "")
	normalizedWanted := strings.ReplaceAll(wanted, "\n", "")
	if normalizedGot != normalizedWanted {
		t.Errorf("build mermaid graph failed. \nwanted \n%s\nand got \n%s\n", wanted, got)
		fmt.Println([]byte(normalizedGot))
		fmt.Println([]byte(normalizedWanted))
	}
}