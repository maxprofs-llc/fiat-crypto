/* Autogenerated: src/ExtractionOCaml/unsaturated_solinas --lang=Go --cmovznz-by-mul --widen-carry --widen-bytes 25519 10 '2^255 - 19' 32 carry_mul carry_square carry_scmul121666 carry add sub opp selectznz from_bytes */
/* curve description: 25519 */
/* requested operations: carry_mul, carry_square, carry_scmul121666, carry, add, sub, opp, selectznz, from_bytes */
/* n = 10 (from "10") */
/* s-c = 2^255 - [(1, 19)] (from "2^255 - 19") */
/* machine_wordsize = 32 (from "32") */

/* Computed values: */
/* carry_chain = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1] */

package fiat_25519



/*
 * The function fiat_25519_cmovznz_u32 is a single-word conditional move.
 * Postconditions:
 *   out1 = (if arg1 = 0 then arg2 else arg3)
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [0x0 ~> 0xffffffff]
 *   arg3: [0x0 ~> 0xffffffff]
 * Output Bounds:
 *   out1: [0x0 ~> 0xffffffff]
 */
/*inline*/
func fiat_25519_cmovznz_u32(out1 *uint32, arg1 uint32, arg2 uint32, arg3 uint32) {
  var x1 uint32 = (arg1 * 0xffffffff)
  var x2 uint32 = ((x1 & arg3) | ((^x1) & arg2))
  *out1 = x2
}

/*
 * The function fiat_25519_carry_mul multiplies two field elements and reduces the result.
 * Postconditions:
 *   eval out1 mod m = (eval arg1 * eval arg2) mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999]]
 *   arg2: [[0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333]]
 */
/*inline*/
func fiat_25519_carry_mul(out1 *[10]uint32, arg1 *[10]uint32, arg2 *[10]uint32) {
  var x1 uint64 = (uint64((arg1[9])) * uint64(((arg2[9]) * (uint32(0x2) * uint32(0x13)))))
  var x2 uint64 = (uint64((arg1[9])) * uint64(((arg2[8]) * uint32(0x13))))
  var x3 uint64 = (uint64((arg1[9])) * uint64(((arg2[7]) * (uint32(0x2) * uint32(0x13)))))
  var x4 uint64 = (uint64((arg1[9])) * uint64(((arg2[6]) * uint32(0x13))))
  var x5 uint64 = (uint64((arg1[9])) * uint64(((arg2[5]) * (uint32(0x2) * uint32(0x13)))))
  var x6 uint64 = (uint64((arg1[9])) * uint64(((arg2[4]) * uint32(0x13))))
  var x7 uint64 = (uint64((arg1[9])) * uint64(((arg2[3]) * (uint32(0x2) * uint32(0x13)))))
  var x8 uint64 = (uint64((arg1[9])) * uint64(((arg2[2]) * uint32(0x13))))
  var x9 uint64 = (uint64((arg1[9])) * uint64(((arg2[1]) * (uint32(0x2) * uint32(0x13)))))
  var x10 uint64 = (uint64((arg1[8])) * uint64(((arg2[9]) * uint32(0x13))))
  var x11 uint64 = (uint64((arg1[8])) * uint64(((arg2[8]) * uint32(0x13))))
  var x12 uint64 = (uint64((arg1[8])) * uint64(((arg2[7]) * uint32(0x13))))
  var x13 uint64 = (uint64((arg1[8])) * uint64(((arg2[6]) * uint32(0x13))))
  var x14 uint64 = (uint64((arg1[8])) * uint64(((arg2[5]) * uint32(0x13))))
  var x15 uint64 = (uint64((arg1[8])) * uint64(((arg2[4]) * uint32(0x13))))
  var x16 uint64 = (uint64((arg1[8])) * uint64(((arg2[3]) * uint32(0x13))))
  var x17 uint64 = (uint64((arg1[8])) * uint64(((arg2[2]) * uint32(0x13))))
  var x18 uint64 = (uint64((arg1[7])) * uint64(((arg2[9]) * (uint32(0x2) * uint32(0x13)))))
  var x19 uint64 = (uint64((arg1[7])) * uint64(((arg2[8]) * uint32(0x13))))
  var x20 uint64 = (uint64((arg1[7])) * uint64(((arg2[7]) * (uint32(0x2) * uint32(0x13)))))
  var x21 uint64 = (uint64((arg1[7])) * uint64(((arg2[6]) * uint32(0x13))))
  var x22 uint64 = (uint64((arg1[7])) * uint64(((arg2[5]) * (uint32(0x2) * uint32(0x13)))))
  var x23 uint64 = (uint64((arg1[7])) * uint64(((arg2[4]) * uint32(0x13))))
  var x24 uint64 = (uint64((arg1[7])) * uint64(((arg2[3]) * (uint32(0x2) * uint32(0x13)))))
  var x25 uint64 = (uint64((arg1[6])) * uint64(((arg2[9]) * uint32(0x13))))
  var x26 uint64 = (uint64((arg1[6])) * uint64(((arg2[8]) * uint32(0x13))))
  var x27 uint64 = (uint64((arg1[6])) * uint64(((arg2[7]) * uint32(0x13))))
  var x28 uint64 = (uint64((arg1[6])) * uint64(((arg2[6]) * uint32(0x13))))
  var x29 uint64 = (uint64((arg1[6])) * uint64(((arg2[5]) * uint32(0x13))))
  var x30 uint64 = (uint64((arg1[6])) * uint64(((arg2[4]) * uint32(0x13))))
  var x31 uint64 = (uint64((arg1[5])) * uint64(((arg2[9]) * (uint32(0x2) * uint32(0x13)))))
  var x32 uint64 = (uint64((arg1[5])) * uint64(((arg2[8]) * uint32(0x13))))
  var x33 uint64 = (uint64((arg1[5])) * uint64(((arg2[7]) * (uint32(0x2) * uint32(0x13)))))
  var x34 uint64 = (uint64((arg1[5])) * uint64(((arg2[6]) * uint32(0x13))))
  var x35 uint64 = (uint64((arg1[5])) * uint64(((arg2[5]) * (uint32(0x2) * uint32(0x13)))))
  var x36 uint64 = (uint64((arg1[4])) * uint64(((arg2[9]) * uint32(0x13))))
  var x37 uint64 = (uint64((arg1[4])) * uint64(((arg2[8]) * uint32(0x13))))
  var x38 uint64 = (uint64((arg1[4])) * uint64(((arg2[7]) * uint32(0x13))))
  var x39 uint64 = (uint64((arg1[4])) * uint64(((arg2[6]) * uint32(0x13))))
  var x40 uint64 = (uint64((arg1[3])) * uint64(((arg2[9]) * (uint32(0x2) * uint32(0x13)))))
  var x41 uint64 = (uint64((arg1[3])) * uint64(((arg2[8]) * uint32(0x13))))
  var x42 uint64 = (uint64((arg1[3])) * uint64(((arg2[7]) * (uint32(0x2) * uint32(0x13)))))
  var x43 uint64 = (uint64((arg1[2])) * uint64(((arg2[9]) * uint32(0x13))))
  var x44 uint64 = (uint64((arg1[2])) * uint64(((arg2[8]) * uint32(0x13))))
  var x45 uint64 = (uint64((arg1[1])) * uint64(((arg2[9]) * (uint32(0x2) * uint32(0x13)))))
  var x46 uint64 = (uint64((arg1[9])) * uint64((arg2[0])))
  var x47 uint64 = (uint64((arg1[8])) * uint64((arg2[1])))
  var x48 uint64 = (uint64((arg1[8])) * uint64((arg2[0])))
  var x49 uint64 = (uint64((arg1[7])) * uint64((arg2[2])))
  var x50 uint64 = (uint64((arg1[7])) * uint64(((arg2[1]) * uint32(0x2))))
  var x51 uint64 = (uint64((arg1[7])) * uint64((arg2[0])))
  var x52 uint64 = (uint64((arg1[6])) * uint64((arg2[3])))
  var x53 uint64 = (uint64((arg1[6])) * uint64((arg2[2])))
  var x54 uint64 = (uint64((arg1[6])) * uint64((arg2[1])))
  var x55 uint64 = (uint64((arg1[6])) * uint64((arg2[0])))
  var x56 uint64 = (uint64((arg1[5])) * uint64((arg2[4])))
  var x57 uint64 = (uint64((arg1[5])) * uint64(((arg2[3]) * uint32(0x2))))
  var x58 uint64 = (uint64((arg1[5])) * uint64((arg2[2])))
  var x59 uint64 = (uint64((arg1[5])) * uint64(((arg2[1]) * uint32(0x2))))
  var x60 uint64 = (uint64((arg1[5])) * uint64((arg2[0])))
  var x61 uint64 = (uint64((arg1[4])) * uint64((arg2[5])))
  var x62 uint64 = (uint64((arg1[4])) * uint64((arg2[4])))
  var x63 uint64 = (uint64((arg1[4])) * uint64((arg2[3])))
  var x64 uint64 = (uint64((arg1[4])) * uint64((arg2[2])))
  var x65 uint64 = (uint64((arg1[4])) * uint64((arg2[1])))
  var x66 uint64 = (uint64((arg1[4])) * uint64((arg2[0])))
  var x67 uint64 = (uint64((arg1[3])) * uint64((arg2[6])))
  var x68 uint64 = (uint64((arg1[3])) * uint64(((arg2[5]) * uint32(0x2))))
  var x69 uint64 = (uint64((arg1[3])) * uint64((arg2[4])))
  var x70 uint64 = (uint64((arg1[3])) * uint64(((arg2[3]) * uint32(0x2))))
  var x71 uint64 = (uint64((arg1[3])) * uint64((arg2[2])))
  var x72 uint64 = (uint64((arg1[3])) * uint64(((arg2[1]) * uint32(0x2))))
  var x73 uint64 = (uint64((arg1[3])) * uint64((arg2[0])))
  var x74 uint64 = (uint64((arg1[2])) * uint64((arg2[7])))
  var x75 uint64 = (uint64((arg1[2])) * uint64((arg2[6])))
  var x76 uint64 = (uint64((arg1[2])) * uint64((arg2[5])))
  var x77 uint64 = (uint64((arg1[2])) * uint64((arg2[4])))
  var x78 uint64 = (uint64((arg1[2])) * uint64((arg2[3])))
  var x79 uint64 = (uint64((arg1[2])) * uint64((arg2[2])))
  var x80 uint64 = (uint64((arg1[2])) * uint64((arg2[1])))
  var x81 uint64 = (uint64((arg1[2])) * uint64((arg2[0])))
  var x82 uint64 = (uint64((arg1[1])) * uint64((arg2[8])))
  var x83 uint64 = (uint64((arg1[1])) * uint64(((arg2[7]) * uint32(0x2))))
  var x84 uint64 = (uint64((arg1[1])) * uint64((arg2[6])))
  var x85 uint64 = (uint64((arg1[1])) * uint64(((arg2[5]) * uint32(0x2))))
  var x86 uint64 = (uint64((arg1[1])) * uint64((arg2[4])))
  var x87 uint64 = (uint64((arg1[1])) * uint64(((arg2[3]) * uint32(0x2))))
  var x88 uint64 = (uint64((arg1[1])) * uint64((arg2[2])))
  var x89 uint64 = (uint64((arg1[1])) * uint64(((arg2[1]) * uint32(0x2))))
  var x90 uint64 = (uint64((arg1[1])) * uint64((arg2[0])))
  var x91 uint64 = (uint64((arg1[0])) * uint64((arg2[9])))
  var x92 uint64 = (uint64((arg1[0])) * uint64((arg2[8])))
  var x93 uint64 = (uint64((arg1[0])) * uint64((arg2[7])))
  var x94 uint64 = (uint64((arg1[0])) * uint64((arg2[6])))
  var x95 uint64 = (uint64((arg1[0])) * uint64((arg2[5])))
  var x96 uint64 = (uint64((arg1[0])) * uint64((arg2[4])))
  var x97 uint64 = (uint64((arg1[0])) * uint64((arg2[3])))
  var x98 uint64 = (uint64((arg1[0])) * uint64((arg2[2])))
  var x99 uint64 = (uint64((arg1[0])) * uint64((arg2[1])))
  var x100 uint64 = (uint64((arg1[0])) * uint64((arg2[0])))
  var x101 uint64 = (x100 + (x45 + (x44 + (x42 + (x39 + (x35 + (x30 + (x24 + (x17 + x9)))))))))
  var x102 uint64 = (x101 >> 26)
  var x103 uint32 = (uint32(x101) & 0x3ffffff)
  var x104 uint64 = (x91 + (x82 + (x74 + (x67 + (x61 + (x56 + (x52 + (x49 + (x47 + x46)))))))))
  var x105 uint64 = (x92 + (x83 + (x75 + (x68 + (x62 + (x57 + (x53 + (x50 + (x48 + x1)))))))))
  var x106 uint64 = (x93 + (x84 + (x76 + (x69 + (x63 + (x58 + (x54 + (x51 + (x10 + x2)))))))))
  var x107 uint64 = (x94 + (x85 + (x77 + (x70 + (x64 + (x59 + (x55 + (x18 + (x11 + x3)))))))))
  var x108 uint64 = (x95 + (x86 + (x78 + (x71 + (x65 + (x60 + (x25 + (x19 + (x12 + x4)))))))))
  var x109 uint64 = (x96 + (x87 + (x79 + (x72 + (x66 + (x31 + (x26 + (x20 + (x13 + x5)))))))))
  var x110 uint64 = (x97 + (x88 + (x80 + (x73 + (x36 + (x32 + (x27 + (x21 + (x14 + x6)))))))))
  var x111 uint64 = (x98 + (x89 + (x81 + (x40 + (x37 + (x33 + (x28 + (x22 + (x15 + x7)))))))))
  var x112 uint64 = (x99 + (x90 + (x43 + (x41 + (x38 + (x34 + (x29 + (x23 + (x16 + x8)))))))))
  var x113 uint64 = (x102 + x112)
  var x114 uint64 = (x113 >> 25)
  var x115 uint32 = (uint32(x113) & 0x1ffffff)
  var x116 uint64 = (x114 + x111)
  var x117 uint64 = (x116 >> 26)
  var x118 uint32 = (uint32(x116) & 0x3ffffff)
  var x119 uint64 = (x117 + x110)
  var x120 uint64 = (x119 >> 25)
  var x121 uint32 = (uint32(x119) & 0x1ffffff)
  var x122 uint64 = (x120 + x109)
  var x123 uint64 = (x122 >> 26)
  var x124 uint32 = (uint32(x122) & 0x3ffffff)
  var x125 uint64 = (x123 + x108)
  var x126 uint64 = (x125 >> 25)
  var x127 uint32 = (uint32(x125) & 0x1ffffff)
  var x128 uint64 = (x126 + x107)
  var x129 uint64 = (x128 >> 26)
  var x130 uint32 = (uint32(x128) & 0x3ffffff)
  var x131 uint64 = (x129 + x106)
  var x132 uint64 = (x131 >> 25)
  var x133 uint32 = (uint32(x131) & 0x1ffffff)
  var x134 uint64 = (x132 + x105)
  var x135 uint64 = (x134 >> 26)
  var x136 uint32 = (uint32(x134) & 0x3ffffff)
  var x137 uint64 = (x135 + x104)
  var x138 uint64 = (x137 >> 25)
  var x139 uint32 = (uint32(x137) & 0x1ffffff)
  var x140 uint64 = (x138 * uint64(0x13))
  var x141 uint64 = (uint64(x103) + x140)
  var x142 uint32 = uint32((x141 >> 26))
  var x143 uint32 = (uint32(x141) & 0x3ffffff)
  var x144 uint32 = (x142 + x115)
  var x145 uint32 = (x144 >> 25)
  var x146 uint32 = (x144 & 0x1ffffff)
  var x147 uint32 = (x145 + x118)
  out1[0] = x143
  out1[1] = x146
  out1[2] = x147
  out1[3] = x121
  out1[4] = x124
  out1[5] = x127
  out1[6] = x130
  out1[7] = x133
  out1[8] = x136
  out1[9] = x139
}

/*
 * The function fiat_25519_carry_square squares a field element and reduces the result.
 * Postconditions:
 *   eval out1 mod m = (eval arg1 * eval arg1) mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333]]
 */
/*inline*/
func fiat_25519_carry_square(out1 *[10]uint32, arg1 *[10]uint32) {
  var x1 uint32 = ((arg1[9]) * uint32(0x13))
  var x2 uint32 = (x1 * uint32(0x2))
  var x3 uint32 = ((arg1[9]) * uint32(0x2))
  var x4 uint32 = ((arg1[8]) * uint32(0x13))
  var x5 uint64 = (uint64(x4) * uint64(0x2))
  var x6 uint32 = ((arg1[8]) * uint32(0x2))
  var x7 uint32 = ((arg1[7]) * uint32(0x13))
  var x8 uint32 = (x7 * uint32(0x2))
  var x9 uint32 = ((arg1[7]) * uint32(0x2))
  var x10 uint32 = ((arg1[6]) * uint32(0x13))
  var x11 uint64 = (uint64(x10) * uint64(0x2))
  var x12 uint32 = ((arg1[6]) * uint32(0x2))
  var x13 uint32 = ((arg1[5]) * uint32(0x13))
  var x14 uint32 = ((arg1[5]) * uint32(0x2))
  var x15 uint32 = ((arg1[4]) * uint32(0x2))
  var x16 uint32 = ((arg1[3]) * uint32(0x2))
  var x17 uint32 = ((arg1[2]) * uint32(0x2))
  var x18 uint32 = ((arg1[1]) * uint32(0x2))
  var x19 uint64 = (uint64((arg1[9])) * uint64((x1 * uint32(0x2))))
  var x20 uint64 = (uint64((arg1[8])) * uint64(x2))
  var x21 uint64 = (uint64((arg1[8])) * uint64(x4))
  var x22 uint64 = (uint64((arg1[7])) * (uint64(x2) * uint64(0x2)))
  var x23 uint64 = (uint64((arg1[7])) * x5)
  var x24 uint64 = (uint64((arg1[7])) * uint64((x7 * uint32(0x2))))
  var x25 uint64 = (uint64((arg1[6])) * uint64(x2))
  var x26 uint64 = (uint64((arg1[6])) * x5)
  var x27 uint64 = (uint64((arg1[6])) * uint64(x8))
  var x28 uint64 = (uint64((arg1[6])) * uint64(x10))
  var x29 uint64 = (uint64((arg1[5])) * (uint64(x2) * uint64(0x2)))
  var x30 uint64 = (uint64((arg1[5])) * x5)
  var x31 uint64 = (uint64((arg1[5])) * (uint64(x8) * uint64(0x2)))
  var x32 uint64 = (uint64((arg1[5])) * x11)
  var x33 uint64 = (uint64((arg1[5])) * uint64((x13 * uint32(0x2))))
  var x34 uint64 = (uint64((arg1[4])) * uint64(x2))
  var x35 uint64 = (uint64((arg1[4])) * x5)
  var x36 uint64 = (uint64((arg1[4])) * uint64(x8))
  var x37 uint64 = (uint64((arg1[4])) * x11)
  var x38 uint64 = (uint64((arg1[4])) * uint64(x14))
  var x39 uint64 = (uint64((arg1[4])) * uint64((arg1[4])))
  var x40 uint64 = (uint64((arg1[3])) * (uint64(x2) * uint64(0x2)))
  var x41 uint64 = (uint64((arg1[3])) * x5)
  var x42 uint64 = (uint64((arg1[3])) * (uint64(x8) * uint64(0x2)))
  var x43 uint64 = (uint64((arg1[3])) * uint64(x12))
  var x44 uint64 = (uint64((arg1[3])) * uint64((x14 * uint32(0x2))))
  var x45 uint64 = (uint64((arg1[3])) * uint64(x15))
  var x46 uint64 = (uint64((arg1[3])) * uint64(((arg1[3]) * uint32(0x2))))
  var x47 uint64 = (uint64((arg1[2])) * uint64(x2))
  var x48 uint64 = (uint64((arg1[2])) * x5)
  var x49 uint64 = (uint64((arg1[2])) * uint64(x9))
  var x50 uint64 = (uint64((arg1[2])) * uint64(x12))
  var x51 uint64 = (uint64((arg1[2])) * uint64(x14))
  var x52 uint64 = (uint64((arg1[2])) * uint64(x15))
  var x53 uint64 = (uint64((arg1[2])) * uint64(x16))
  var x54 uint64 = (uint64((arg1[2])) * uint64((arg1[2])))
  var x55 uint64 = (uint64((arg1[1])) * (uint64(x2) * uint64(0x2)))
  var x56 uint64 = (uint64((arg1[1])) * uint64(x6))
  var x57 uint64 = (uint64((arg1[1])) * uint64((x9 * uint32(0x2))))
  var x58 uint64 = (uint64((arg1[1])) * uint64(x12))
  var x59 uint64 = (uint64((arg1[1])) * uint64((x14 * uint32(0x2))))
  var x60 uint64 = (uint64((arg1[1])) * uint64(x15))
  var x61 uint64 = (uint64((arg1[1])) * uint64((x16 * uint32(0x2))))
  var x62 uint64 = (uint64((arg1[1])) * uint64(x17))
  var x63 uint64 = (uint64((arg1[1])) * uint64(((arg1[1]) * uint32(0x2))))
  var x64 uint64 = (uint64((arg1[0])) * uint64(x3))
  var x65 uint64 = (uint64((arg1[0])) * uint64(x6))
  var x66 uint64 = (uint64((arg1[0])) * uint64(x9))
  var x67 uint64 = (uint64((arg1[0])) * uint64(x12))
  var x68 uint64 = (uint64((arg1[0])) * uint64(x14))
  var x69 uint64 = (uint64((arg1[0])) * uint64(x15))
  var x70 uint64 = (uint64((arg1[0])) * uint64(x16))
  var x71 uint64 = (uint64((arg1[0])) * uint64(x17))
  var x72 uint64 = (uint64((arg1[0])) * uint64(x18))
  var x73 uint64 = (uint64((arg1[0])) * uint64((arg1[0])))
  var x74 uint64 = (x73 + (x55 + (x48 + (x42 + (x37 + x33)))))
  var x75 uint64 = (x74 >> 26)
  var x76 uint32 = (uint32(x74) & 0x3ffffff)
  var x77 uint64 = (x64 + (x56 + (x49 + (x43 + x38))))
  var x78 uint64 = (x65 + (x57 + (x50 + (x44 + (x39 + x19)))))
  var x79 uint64 = (x66 + (x58 + (x51 + (x45 + x20))))
  var x80 uint64 = (x67 + (x59 + (x52 + (x46 + (x22 + x21)))))
  var x81 uint64 = (x68 + (x60 + (x53 + (x25 + x23))))
  var x82 uint64 = (x69 + (x61 + (x54 + (x29 + (x26 + x24)))))
  var x83 uint64 = (x70 + (x62 + (x34 + (x30 + x27))))
  var x84 uint64 = (x71 + (x63 + (x40 + (x35 + (x31 + x28)))))
  var x85 uint64 = (x72 + (x47 + (x41 + (x36 + x32))))
  var x86 uint64 = (x75 + x85)
  var x87 uint64 = (x86 >> 25)
  var x88 uint32 = (uint32(x86) & 0x1ffffff)
  var x89 uint64 = (x87 + x84)
  var x90 uint64 = (x89 >> 26)
  var x91 uint32 = (uint32(x89) & 0x3ffffff)
  var x92 uint64 = (x90 + x83)
  var x93 uint64 = (x92 >> 25)
  var x94 uint32 = (uint32(x92) & 0x1ffffff)
  var x95 uint64 = (x93 + x82)
  var x96 uint64 = (x95 >> 26)
  var x97 uint32 = (uint32(x95) & 0x3ffffff)
  var x98 uint64 = (x96 + x81)
  var x99 uint64 = (x98 >> 25)
  var x100 uint32 = (uint32(x98) & 0x1ffffff)
  var x101 uint64 = (x99 + x80)
  var x102 uint64 = (x101 >> 26)
  var x103 uint32 = (uint32(x101) & 0x3ffffff)
  var x104 uint64 = (x102 + x79)
  var x105 uint64 = (x104 >> 25)
  var x106 uint32 = (uint32(x104) & 0x1ffffff)
  var x107 uint64 = (x105 + x78)
  var x108 uint64 = (x107 >> 26)
  var x109 uint32 = (uint32(x107) & 0x3ffffff)
  var x110 uint64 = (x108 + x77)
  var x111 uint64 = (x110 >> 25)
  var x112 uint32 = (uint32(x110) & 0x1ffffff)
  var x113 uint64 = (x111 * uint64(0x13))
  var x114 uint64 = (uint64(x76) + x113)
  var x115 uint32 = uint32((x114 >> 26))
  var x116 uint32 = (uint32(x114) & 0x3ffffff)
  var x117 uint32 = (x115 + x88)
  var x118 uint32 = (x117 >> 25)
  var x119 uint32 = (x117 & 0x1ffffff)
  var x120 uint32 = (x118 + x91)
  out1[0] = x116
  out1[1] = x119
  out1[2] = x120
  out1[3] = x94
  out1[4] = x97
  out1[5] = x100
  out1[6] = x103
  out1[7] = x106
  out1[8] = x109
  out1[9] = x112
}

/*
 * The function fiat_25519_carry_scmul_121666 multiplies a field element by 121666 and reduces the result.
 * Postconditions:
 *   eval out1 mod m = (121666 * eval arg1) mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333]]
 */
/*inline*/
func fiat_25519_carry_scmul_121666(out1 *[10]uint32, arg1 *[10]uint32) {
  var x1 uint64 = (uint64(0x1db42) * uint64((arg1[9])))
  var x2 uint64 = (uint64(0x1db42) * uint64((arg1[8])))
  var x3 uint64 = (uint64(0x1db42) * uint64((arg1[7])))
  var x4 uint64 = (uint64(0x1db42) * uint64((arg1[6])))
  var x5 uint64 = (uint64(0x1db42) * uint64((arg1[5])))
  var x6 uint64 = (uint64(0x1db42) * uint64((arg1[4])))
  var x7 uint64 = (uint64(0x1db42) * uint64((arg1[3])))
  var x8 uint64 = (uint64(0x1db42) * uint64((arg1[2])))
  var x9 uint64 = (uint64(0x1db42) * uint64((arg1[1])))
  var x10 uint64 = (uint64(0x1db42) * uint64((arg1[0])))
  var x11 uint32 = uint32((x10 >> 26))
  var x12 uint32 = (uint32(x10) & 0x3ffffff)
  var x13 uint64 = (uint64(x11) + x9)
  var x14 uint32 = uint32((x13 >> 25))
  var x15 uint32 = (uint32(x13) & 0x1ffffff)
  var x16 uint64 = (uint64(x14) + x8)
  var x17 uint32 = uint32((x16 >> 26))
  var x18 uint32 = (uint32(x16) & 0x3ffffff)
  var x19 uint64 = (uint64(x17) + x7)
  var x20 uint32 = uint32((x19 >> 25))
  var x21 uint32 = (uint32(x19) & 0x1ffffff)
  var x22 uint64 = (uint64(x20) + x6)
  var x23 uint32 = uint32((x22 >> 26))
  var x24 uint32 = (uint32(x22) & 0x3ffffff)
  var x25 uint64 = (uint64(x23) + x5)
  var x26 uint32 = uint32((x25 >> 25))
  var x27 uint32 = (uint32(x25) & 0x1ffffff)
  var x28 uint64 = (uint64(x26) + x4)
  var x29 uint32 = uint32((x28 >> 26))
  var x30 uint32 = (uint32(x28) & 0x3ffffff)
  var x31 uint64 = (uint64(x29) + x3)
  var x32 uint32 = uint32((x31 >> 25))
  var x33 uint32 = (uint32(x31) & 0x1ffffff)
  var x34 uint64 = (uint64(x32) + x2)
  var x35 uint32 = uint32((x34 >> 26))
  var x36 uint32 = (uint32(x34) & 0x3ffffff)
  var x37 uint64 = (uint64(x35) + x1)
  var x38 uint32 = uint32((x37 >> 25))
  var x39 uint32 = (uint32(x37) & 0x1ffffff)
  var x40 uint32 = (x38 * uint32(0x13))
  var x41 uint32 = (x12 + x40)
  var x42 uint32 = (x41 >> 26)
  var x43 uint32 = (x41 & 0x3ffffff)
  var x44 uint32 = (x42 + x15)
  var x45 uint32 = (x44 >> 25)
  var x46 uint32 = (x44 & 0x1ffffff)
  var x47 uint32 = (x45 + x18)
  out1[0] = x43
  out1[1] = x46
  out1[2] = x47
  out1[3] = x21
  out1[4] = x24
  out1[5] = x27
  out1[6] = x30
  out1[7] = x33
  out1[8] = x36
  out1[9] = x39
}

/*
 * The function fiat_25519_carry reduces a field element.
 * Postconditions:
 *   eval out1 mod m = eval arg1 mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333]]
 */
/*inline*/
func fiat_25519_carry(out1 *[10]uint32, arg1 *[10]uint32) {
  var x1 uint32 = (arg1[0])
  var x2 uint32 = ((x1 >> 26) + (arg1[1]))
  var x3 uint32 = ((x2 >> 25) + (arg1[2]))
  var x4 uint32 = ((x3 >> 26) + (arg1[3]))
  var x5 uint32 = ((x4 >> 25) + (arg1[4]))
  var x6 uint32 = ((x5 >> 26) + (arg1[5]))
  var x7 uint32 = ((x6 >> 25) + (arg1[6]))
  var x8 uint32 = ((x7 >> 26) + (arg1[7]))
  var x9 uint32 = ((x8 >> 25) + (arg1[8]))
  var x10 uint32 = ((x9 >> 26) + (arg1[9]))
  var x11 uint32 = ((x1 & 0x3ffffff) + ((x10 >> 25) * uint32(0x13)))
  var x12 uint32 = ((x11 >> 26) + (x2 & 0x1ffffff))
  var x13 uint32 = (x11 & 0x3ffffff)
  var x14 uint32 = (x12 & 0x1ffffff)
  var x15 uint32 = ((x12 >> 25) + (x3 & 0x3ffffff))
  var x16 uint32 = (x4 & 0x1ffffff)
  var x17 uint32 = (x5 & 0x3ffffff)
  var x18 uint32 = (x6 & 0x1ffffff)
  var x19 uint32 = (x7 & 0x3ffffff)
  var x20 uint32 = (x8 & 0x1ffffff)
  var x21 uint32 = (x9 & 0x3ffffff)
  var x22 uint32 = (x10 & 0x1ffffff)
  out1[0] = x13
  out1[1] = x14
  out1[2] = x15
  out1[3] = x16
  out1[4] = x17
  out1[5] = x18
  out1[6] = x19
  out1[7] = x20
  out1[8] = x21
  out1[9] = x22
}

/*
 * The function fiat_25519_add adds two field elements.
 * Postconditions:
 *   eval out1 mod m = (eval arg1 + eval arg2) mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333]]
 *   arg2: [[0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999]]
 */
/*inline*/
func fiat_25519_add(out1 *[10]uint32, arg1 *[10]uint32, arg2 *[10]uint32) {
  var x1 uint32 = ((arg1[0]) + (arg2[0]))
  var x2 uint32 = ((arg1[1]) + (arg2[1]))
  var x3 uint32 = ((arg1[2]) + (arg2[2]))
  var x4 uint32 = ((arg1[3]) + (arg2[3]))
  var x5 uint32 = ((arg1[4]) + (arg2[4]))
  var x6 uint32 = ((arg1[5]) + (arg2[5]))
  var x7 uint32 = ((arg1[6]) + (arg2[6]))
  var x8 uint32 = ((arg1[7]) + (arg2[7]))
  var x9 uint32 = ((arg1[8]) + (arg2[8]))
  var x10 uint32 = ((arg1[9]) + (arg2[9]))
  out1[0] = x1
  out1[1] = x2
  out1[2] = x3
  out1[3] = x4
  out1[4] = x5
  out1[5] = x6
  out1[6] = x7
  out1[7] = x8
  out1[8] = x9
  out1[9] = x10
}

/*
 * The function fiat_25519_sub subtracts two field elements.
 * Postconditions:
 *   eval out1 mod m = (eval arg1 - eval arg2) mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333]]
 *   arg2: [[0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999]]
 */
/*inline*/
func fiat_25519_sub(out1 *[10]uint32, arg1 *[10]uint32, arg2 *[10]uint32) {
  var x1 uint32 = ((0x7ffffda + (arg1[0])) - (arg2[0]))
  var x2 uint32 = ((0x3fffffe + (arg1[1])) - (arg2[1]))
  var x3 uint32 = ((0x7fffffe + (arg1[2])) - (arg2[2]))
  var x4 uint32 = ((0x3fffffe + (arg1[3])) - (arg2[3]))
  var x5 uint32 = ((0x7fffffe + (arg1[4])) - (arg2[4]))
  var x6 uint32 = ((0x3fffffe + (arg1[5])) - (arg2[5]))
  var x7 uint32 = ((0x7fffffe + (arg1[6])) - (arg2[6]))
  var x8 uint32 = ((0x3fffffe + (arg1[7])) - (arg2[7]))
  var x9 uint32 = ((0x7fffffe + (arg1[8])) - (arg2[8]))
  var x10 uint32 = ((0x3fffffe + (arg1[9])) - (arg2[9]))
  out1[0] = x1
  out1[1] = x2
  out1[2] = x3
  out1[3] = x4
  out1[4] = x5
  out1[5] = x6
  out1[6] = x7
  out1[7] = x8
  out1[8] = x9
  out1[9] = x10
}

/*
 * The function fiat_25519_opp negates a field element.
 * Postconditions:
 *   eval out1 mod m = -eval arg1 mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999], [0x0 ~> 0xd333332], [0x0 ~> 0x6999999]]
 */
/*inline*/
func fiat_25519_opp(out1 *[10]uint32, arg1 *[10]uint32) {
  var x1 uint32 = (0x7ffffda - (arg1[0]))
  var x2 uint32 = (0x3fffffe - (arg1[1]))
  var x3 uint32 = (0x7fffffe - (arg1[2]))
  var x4 uint32 = (0x3fffffe - (arg1[3]))
  var x5 uint32 = (0x7fffffe - (arg1[4]))
  var x6 uint32 = (0x3fffffe - (arg1[5]))
  var x7 uint32 = (0x7fffffe - (arg1[6]))
  var x8 uint32 = (0x3fffffe - (arg1[7]))
  var x9 uint32 = (0x7fffffe - (arg1[8]))
  var x10 uint32 = (0x3fffffe - (arg1[9]))
  out1[0] = x1
  out1[1] = x2
  out1[2] = x3
  out1[3] = x4
  out1[4] = x5
  out1[5] = x6
  out1[6] = x7
  out1[7] = x8
  out1[8] = x9
  out1[9] = x10
}

/*
 * The function fiat_25519_selectznz is a multi-limb conditional select.
 * Postconditions:
 *   eval out1 = (if arg1 = 0 then eval arg2 else eval arg3)
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [[0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff]]
 *   arg3: [[0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff]]
 */
/*inline*/
func fiat_25519_selectznz(out1 *[10]uint32, arg1 uint32, arg2 *[10]uint32, arg3 *[10]uint32) {
  var x1 uint32
  fiat_25519_cmovznz_u32(&x1, arg1, (arg2[0]), (arg3[0]))
  var x2 uint32
  fiat_25519_cmovznz_u32(&x2, arg1, (arg2[1]), (arg3[1]))
  var x3 uint32
  fiat_25519_cmovznz_u32(&x3, arg1, (arg2[2]), (arg3[2]))
  var x4 uint32
  fiat_25519_cmovznz_u32(&x4, arg1, (arg2[3]), (arg3[3]))
  var x5 uint32
  fiat_25519_cmovznz_u32(&x5, arg1, (arg2[4]), (arg3[4]))
  var x6 uint32
  fiat_25519_cmovznz_u32(&x6, arg1, (arg2[5]), (arg3[5]))
  var x7 uint32
  fiat_25519_cmovznz_u32(&x7, arg1, (arg2[6]), (arg3[6]))
  var x8 uint32
  fiat_25519_cmovznz_u32(&x8, arg1, (arg2[7]), (arg3[7]))
  var x9 uint32
  fiat_25519_cmovznz_u32(&x9, arg1, (arg2[8]), (arg3[8]))
  var x10 uint32
  fiat_25519_cmovznz_u32(&x10, arg1, (arg2[9]), (arg3[9]))
  out1[0] = x1
  out1[1] = x2
  out1[2] = x3
  out1[3] = x4
  out1[4] = x5
  out1[5] = x6
  out1[6] = x7
  out1[7] = x8
  out1[8] = x9
  out1[9] = x10
}

/*
 * The function fiat_25519_from_bytes deserializes a field element from bytes in little-endian order.
 * Postconditions:
 *   eval out1 mod m = bytes_eval arg1 mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x7f]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333], [0x0 ~> 0x4666666], [0x0 ~> 0x2333333]]
 */
/*inline*/
func fiat_25519_from_bytes(out1 *[10]uint32, arg1 *[32]uint32) {
  var x1 uint32 = ((arg1[31]) << 18)
  var x2 uint32 = ((arg1[30]) << 10)
  var x3 uint32 = ((arg1[29]) << 2)
  var x4 uint32 = ((arg1[28]) << 20)
  var x5 uint32 = ((arg1[27]) << 12)
  var x6 uint32 = ((arg1[26]) << 4)
  var x7 uint32 = ((arg1[25]) << 21)
  var x8 uint32 = ((arg1[24]) << 13)
  var x9 uint32 = ((arg1[23]) << 5)
  var x10 uint32 = ((arg1[22]) << 23)
  var x11 uint32 = ((arg1[21]) << 15)
  var x12 uint32 = ((arg1[20]) << 7)
  var x13 uint32 = ((arg1[19]) << 24)
  var x14 uint32 = ((arg1[18]) << 16)
  var x15 uint32 = ((arg1[17]) << 8)
  var x16 uint32 = (arg1[16])
  var x17 uint32 = ((arg1[15]) << 18)
  var x18 uint32 = ((arg1[14]) << 10)
  var x19 uint32 = ((arg1[13]) << 2)
  var x20 uint32 = ((arg1[12]) << 19)
  var x21 uint32 = ((arg1[11]) << 11)
  var x22 uint32 = ((arg1[10]) << 3)
  var x23 uint32 = ((arg1[9]) << 21)
  var x24 uint32 = ((arg1[8]) << 13)
  var x25 uint32 = ((arg1[7]) << 5)
  var x26 uint32 = ((arg1[6]) << 22)
  var x27 uint32 = ((arg1[5]) << 14)
  var x28 uint32 = ((arg1[4]) << 6)
  var x29 uint32 = ((arg1[3]) << 24)
  var x30 uint32 = ((arg1[2]) << 16)
  var x31 uint32 = ((arg1[1]) << 8)
  var x32 uint32 = (arg1[0])
  var x33 uint32 = (x32 + (x31 + (x30 + x29)))
  var x34 uint32 = (x33 >> 26)
  var x35 uint32 = (x33 & 0x3ffffff)
  var x36 uint32 = (x3 + (x2 + x1))
  var x37 uint32 = (x6 + (x5 + x4))
  var x38 uint32 = (x9 + (x8 + x7))
  var x39 uint32 = (x12 + (x11 + x10))
  var x40 uint32 = (x16 + (x15 + (x14 + x13)))
  var x41 uint32 = (x19 + (x18 + x17))
  var x42 uint32 = (x22 + (x21 + x20))
  var x43 uint32 = (x25 + (x24 + x23))
  var x44 uint32 = (x28 + (x27 + x26))
  var x45 uint32 = (x34 + x44)
  var x46 uint32 = (x45 >> 25)
  var x47 uint32 = (x45 & 0x1ffffff)
  var x48 uint32 = (x46 + x43)
  var x49 uint32 = (x48 >> 26)
  var x50 uint32 = (x48 & 0x3ffffff)
  var x51 uint32 = (x49 + x42)
  var x52 uint32 = (x51 >> 25)
  var x53 uint32 = (x51 & 0x1ffffff)
  var x54 uint32 = (x52 + x41)
  var x55 uint32 = (x54 & 0x3ffffff)
  var x56 uint32 = (x40 >> 25)
  var x57 uint32 = (x40 & 0x1ffffff)
  var x58 uint32 = (x56 + x39)
  var x59 uint32 = (x58 >> 26)
  var x60 uint32 = (x58 & 0x3ffffff)
  var x61 uint32 = (x59 + x38)
  var x62 uint32 = (x61 >> 25)
  var x63 uint32 = (x61 & 0x1ffffff)
  var x64 uint32 = (x62 + x37)
  var x65 uint32 = (x64 >> 26)
  var x66 uint32 = (x64 & 0x3ffffff)
  var x67 uint32 = (x65 + x36)
  out1[0] = x35
  out1[1] = x47
  out1[2] = x50
  out1[3] = x53
  out1[4] = x55
  out1[5] = x57
  out1[6] = x60
  out1[7] = x63
  out1[8] = x66
  out1[9] = x67
}

