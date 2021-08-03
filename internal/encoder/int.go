package encoder

import "unsafe"

const intMask uint64 = 1<<(32<<(^uint(0)>>63)) - 1

var intLookup [100]uint16

// "00010203...96979899" cast to []uint16
var intLELookup = [100]uint16{
	0x3030, 0x3130, 0x3230, 0x3330, 0x3430, 0x3530, 0x3630, 0x3730, 0x3830, 0x3930,
	0x3031, 0x3131, 0x3231, 0x3331, 0x3431, 0x3531, 0x3631, 0x3731, 0x3831, 0x3931,
	0x3032, 0x3132, 0x3232, 0x3332, 0x3432, 0x3532, 0x3632, 0x3732, 0x3832, 0x3932,
	0x3033, 0x3133, 0x3233, 0x3333, 0x3433, 0x3533, 0x3633, 0x3733, 0x3833, 0x3933,
	0x3034, 0x3134, 0x3234, 0x3334, 0x3434, 0x3534, 0x3634, 0x3734, 0x3834, 0x3934,
	0x3035, 0x3135, 0x3235, 0x3335, 0x3435, 0x3535, 0x3635, 0x3735, 0x3835, 0x3935,
	0x3036, 0x3136, 0x3236, 0x3336, 0x3436, 0x3536, 0x3636, 0x3736, 0x3836, 0x3936,
	0x3037, 0x3137, 0x3237, 0x3337, 0x3437, 0x3537, 0x3637, 0x3737, 0x3837, 0x3937,
	0x3038, 0x3138, 0x3238, 0x3338, 0x3438, 0x3538, 0x3638, 0x3738, 0x3838, 0x3938,
	0x3039, 0x3139, 0x3239, 0x3339, 0x3439, 0x3539, 0x3639, 0x3739, 0x3839, 0x3939,
}

func init() {
	var b [2]byte
	*(*uint16)(unsafe.Pointer(&b)) = uint16(0xABCD)
	switch b[0] {
	case 0xCD:
		intLookup = intLELookup
	case 0xAB:
		intLookup = [100]uint16{
			0x3030, 0x3031, 0x3032, 0x3033, 0x3034, 0x3035, 0x3036, 0x3037, 0x3038, 0x3039,
			0x3130, 0x3131, 0x3132, 0x3133, 0x3134, 0x3135, 0x3136, 0x3137, 0x3138, 0x3139,
			0x3230, 0x3231, 0x3232, 0x3233, 0x3234, 0x3235, 0x3236, 0x3237, 0x3238, 0x3239,
			0x3330, 0x3331, 0x3332, 0x3333, 0x3334, 0x3335, 0x3336, 0x3337, 0x3338, 0x3339,
			0x3430, 0x3431, 0x3432, 0x3433, 0x3434, 0x3435, 0x3436, 0x3437, 0x3438, 0x3439,
			0x3530, 0x3531, 0x3532, 0x3533, 0x3534, 0x3535, 0x3536, 0x3537, 0x3538, 0x3539,
			0x3630, 0x3631, 0x3632, 0x3633, 0x3634, 0x3635, 0x3636, 0x3637, 0x3638, 0x3639,
			0x3730, 0x3731, 0x3732, 0x3733, 0x3734, 0x3735, 0x3736, 0x3737, 0x3738, 0x3739,
			0x3830, 0x3831, 0x3832, 0x3833, 0x3834, 0x3835, 0x3836, 0x3837, 0x3838, 0x3839,
			0x3930, 0x3931, 0x3932, 0x3933, 0x3934, 0x3935, 0x3936, 0x3937, 0x3938, 0x3939,
		}
	default:
		panic("could not determine endianness")
	}
}

func (e *Encoder) writeInt(data int64) {
	u64 := uint64(data)
	n := u64 & intMask
	negative := data < 0
	if !negative {
		if n < 10 {
			e.writeByte(byte(n + '0'))
			return
		} else if n < 100 {
			memmove(
				unsafe.Pointer(uintptr((*sliceHeader)(unsafe.Pointer(&e.buffer)).data)+uintptr(e.offset)),
				unsafe.Pointer(&intLELookup[n]),
				2,
			)
			e.offset += 2
			return
		}
	} else {
		n = -n & intMask
	}
	var b [22]byte
	u := (*[11]uint16)(unsafe.Pointer(&b))
	i := 11
	for n >= 100 {
		j := n % 100
		n /= 100
		i--
		u[i] = intLookup[j]
	}
	i--
	u[i] = intLookup[n]
	i *= 2 // convert to byte index
	if n < 10 {
		i++ // remove leading zero
	}
	if negative {
		i--
		b[i] = '-'
	}
	e.write(b[i:])
}

func (e *Encoder) encodeInt(data int64) {
	e.grow(24)
	e.writeByte('i')
	e.writeInt(data)
	e.writeByte('e')
}
