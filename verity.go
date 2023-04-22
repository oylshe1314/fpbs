package fpbs

const factor byte = 0xDB

func Verification(p []byte) byte {
	var code byte = factor
	for i := range p {
		code = code ^ p[i]
	}
	return code
}

func Verify(p []byte) bool {
	var code = Verification(p[:len(p)-1])
	return code == p[len(p)-1]
}
