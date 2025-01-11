package conf

type BootConfig struct {
	Name string
	Type string
	File string
	CR   int
	CC   int
}

type BootEntry struct {
	Color int32
	Desc  string
}

var MBR = map[[2]int]BootEntry{
	{0, 3}: {
		Color: 0x00CC00,
		Desc:  "跳转指令",
	},
	{3, 11}: {
		Color: 0x0000CC,
		Desc:  "厂商名称",
	},
	{11, 13}: {
		Color: 0xCC0000,
		Desc:  "每扇区字节数",
	},
	{13, 14}: {
		Color: 0x990000,
		Desc:  "每蔟扇区数",
	},
	{14, 16}: {
		Color: 0x009900,
		Desc:  "MBR占用扇区数",
	},
	{16, 17}: {
		Color: 0x000099,
		Desc:  "FAT个数",
	},
	{17, 19}: {
		Color: 0x99CC00,
		Desc:  "根目录区最大文件数",
	},
	{19, 21}: {
		Color: 0x0099CC,
		Desc:  "总扇区数(2bytes)",
	},
	{21, 22}: {
		Color: 0xCC9900,
		Desc:  "介质描述符",
	},
	{22, 24}: {
		Color: 0x0099CC,
		Desc:  "每个FAT所占扇区数",
	},
	{24, 26}: {
		Color: 0xCC0099,
		Desc:  "每磁道扇区数",
	},
	{26, 28}: {
		Color: 0x9900CC,
		Desc:  "磁头数",
	},
	{28, 32}: {
		Color: 0x00CC00,
		Desc:  "隐藏扇区数",
	},
	{32, 36}: {
		Color: 0x0000CC,
		Desc:  "总扇区数(4bytes)",
	},
	{36, 37}: {
		Color: 0x99CC00,
		Desc:  "中断0x13的驱动器号",
	},
	{37, 38}: {
		Color: 0xCC0000,
		Desc:  "保留位",
	},
	{38, 39}: {
		Color: 0xCC0099,
		Desc:  "拓展引导标记",
	},
	{39, 43}: {
		Color: 0x9900CC,
		Desc:  "卷序列号",
	},
	{43, 54}: {
		Color: 0x00CCCC,
		Desc:  "卷标",
	},
	{54, 62}: {
		Color: 0x0099CCC,
		Desc:  "文件系统类型",
	},
	{62, 510}: {
		Color: 0xFFFFFF,
		Desc:  "引导扇区代码",
	},
	{510, 512}: {
		Color: 0xCC0000,
		Desc:  "结束标志",
	},
}
