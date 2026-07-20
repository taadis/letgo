package unit

// 计算机存储单位
type Size uint64

const (
	Bit Size = 1 // 位, 也叫比特, 是计算机存储中的最小单位, 如: 0,1

	Byte = 8 * Bit     // 字节, 8 位 1 字节
	KB   = 1024 * Byte // KiloByte 千字节, 也就是我们常说的 KB
	MB   = 1024 * KB   // MegaByte MB 兆字节, 也就是我们常说的几兆几兆, 兆已经是百万级数量单位
	GB   = 1024 * MB   // GigaByte GB 吉字节, 也叫 "千兆", 也就是我们常说的 U 盘几 G, 电脑/手机内存几 G
	// 后面的单位, 基本上只有大数据才会用到.
	TB = 1024 * GB // Terabyte TB 太字节, 万亿字节
	PB = 1024 * TB // PetaByte PB 拍字节
	EB = 1024 * PB // ExaByte EB 艾字节, 百亿亿字节
	// more will overflows
	// ZB = 1024 * EB // ZettaByte ZB 泽字节, 十万亿亿字节
	// YB = 1024 * ZB // YottaByte YB 尧字节, 一亿亿亿字节
)

// 长度单位 (以毫米为基础单位)
type Length int64

const (
	Millimeter Length = 1

	Centimeter = 10 * Millimeter
	Meter      = 100 * Centimeter
	Kilometer  = 1000 * Meter
)

// 重量单位 (以克为基础单位)
type Weight int64

const (
	Gram Weight = 1

	Kilogram = 1000 * Gram
	Tonne    = 1000 * Kilogram // 吨
)
