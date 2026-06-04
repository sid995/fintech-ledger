package version

const Name = "fintech-ledger"
const Phase0 = "0.0.0-phase0"

// Banner returns a single-line startup message for tooling smoke checks.
func Banner() string {
	return Name + " " + Phase0
}
