package nouveau

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

var ErrPowerStateMismatch = errors.New("current power state mismatches newly set power state (does supplied power state exist?)")

type NouveauPowerState struct {
	State         string `json:"state"`         // A hexadecimal byte indicator.
	CoreFreqRange string `json:"coreFreqRange"` // In the format X MHz or else X-Y MHz.
	MemFreqRange  string `json:"memFreqRange"`  // In the format X MHz or else X-Y MHz.
	UsedOn        string `json:"selectedOn"`    // Empty, AC, DC or both AC DC.
	Selected      bool   `json:"selected"`      // If this power state is currently active.
	CurrentlyOn   string `json:"currentPower"`  // AC or DC, indicates current power supply, only on selected power state.
}

func NouveauSetPowerState(driDevice int, state string) error {
	file, err := os.OpenFile("/sys/kernel/debug/dri/"+strconv.Itoa(driDevice)+"/pstate", os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	n, err := file.Write([]byte(state))
	if err != nil {
		return err
	} else if n != len([]byte(state)) {
		return ErrPowerStateMismatch
	}
	currentState, _, err := NouveauGetPowerStates(driDevice)
	if err != nil {
		return err
	} else if currentState.State != state {
		return ErrPowerStateMismatch
	}
	return nil
}

func nouveauParsePowerState(state string) *NouveauPowerState {
	coreFreqIndex := strings.Index(state, "core ") + 5
	memFreqIndex := strings.Index(state, "memory")
	// AC always comes first.
	memEndIndex := len(state)
	acIndex := strings.Index(state, "AC")
	dcIndex := strings.Index(state, "DC")
	if acIndex > 0 { // -1 and 0 indicate that this is the selected state.
		// Then it's selected.
		memEndIndex = acIndex - 1
	} else if dcIndex > 0 { // If only DC is selected for this power state and not AC.
		memEndIndex = dcIndex - 1
	}
	usedOn := strings.TrimSuffix(strings.TrimSpace(state[memEndIndex:]), " *")
	return &NouveauPowerState{
		State:         state[:2],
		Selected:      strings.Contains(state, "*"),
		CoreFreqRange: state[coreFreqIndex : memFreqIndex-1],
		MemFreqRange:  state[memFreqIndex+7 : memEndIndex],
		UsedOn:        usedOn,
	}
}

func NouveauGetPowerStates(driDevice int) (*NouveauPowerState, []*NouveauPowerState, error) {
	content, err := os.ReadFile("/sys/kernel/debug/dri/" + strconv.Itoa(driDevice) + "/pstate")
	if err != nil {
		return nil, nil, err
	}
	str := string(content)
	states := strings.Split(strings.TrimSpace(str), "\n")
	currentState := states[len(states)-1]
	availableStates := states[:len(states)-1]
	parsedAvailableStates := make([]*NouveauPowerState, len(availableStates))
	for index, state := range availableStates {
		parsedAvailableStates[index] = nouveauParsePowerState(state)
	}
	// Parse the current state.
	parsedCurrentState := nouveauParsePowerState(currentState)
	parsedCurrentState.CurrentlyOn = parsedCurrentState.State
	parsedCurrentState.Selected = true
	parsedCurrentState.State = ""
	for _, state := range parsedAvailableStates {
		if state.Selected {
			parsedCurrentState.State = state.State
			parsedCurrentState.UsedOn = state.UsedOn
			state.CurrentlyOn = parsedCurrentState.CurrentlyOn
		}
	}
	return parsedCurrentState, parsedAvailableStates, nil
}
