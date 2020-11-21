package policy

import (
	"errors"
	"fmt"
)

type policyType string

var KONP_POLICY = policyType("KONP")
var NP_POLICY = policyType("NP")
var GPP_POLICY = policyType("GPP")
var LC_POLICY = policyType("LC")
var RPP_POLICY = policyType("RPP")
var RPP1_POLICY = policyType("RPP1")
var RPP2_POLICY = policyType("RPP2")
var RPP3_POLICY = policyType("RPP3")
var RPP4_POLICY = policyType("RPP4")

var POLICY_NOT_FOUND_ERROR = errors.New("POLICY_NOT_FOUND_ERROR")

// New policy for buffer management
func New(name policyType) Policy {
	switch name {
	case KONP_POLICY:
		return KONP{basePolicy{Name: string(name)}}
	case NP_POLICY:
		return NP{basePolicy{Name: string(name)}}
	case GPP_POLICY:
		return GPP{basePolicy{Name: string(name)}}
	case LC_POLICY:
		return LC{basePolicy{Name: string(name)}}
    case RPP_POLICY:
        return RPP{basePolicy{Name: string(name)}}
    case RPP1_POLICY:
        return RPP1{basePolicy{Name: string(name)}}
    case RPP2_POLICY:
        return RPP2{basePolicy{Name: string(name)}}
    case RPP3_POLICY:
        return RPP3{basePolicy{Name: string(name)}}
    case RPP4_POLICY:
        return RPP4{basePolicy{Name: string(name)}}
	default:
		panic(fmt.Errorf("%w %s", POLICY_NOT_FOUND_ERROR, name))
	}
}
