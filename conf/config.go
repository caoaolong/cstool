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
		Color: 0x006600,
		Desc:  "跳转指令",
	},
	{3, 11}: {
		Color: 0x000066,
		Desc:  "厂商名称",
	},
	{11, 13}: {
		Color: 0x660000,
		Desc:  "每扇区字节数",
	},
	{13, 14}: {
		Color: 0x330000,
		Desc:  "每蔟扇区数",
	},
	{14, 16}: {
		Color: 0x003300,
		Desc:  "MBR占用扇区数",
	},
	{16, 17}: {
		Color: 0x000033,
		Desc:  "FAT个数",
	},
	{17, 19}: {
		Color: 0x336600,
		Desc:  "根目录区最大文件数",
	},
	{19, 21}: {
		Color: 0x003366,
		Desc:  "总扇区数(2bytes)",
	},
	{21, 22}: {
		Color: 0x663300,
		Desc:  "介质描述符",
	},
	{22, 24}: {
		Color: 0x003366,
		Desc:  "每个FAT所占扇区数",
	},
	{24, 26}: {
		Color: 0x660033,
		Desc:  "每磁道扇区数",
	},
	{26, 28}: {
		Color: 0x330066,
		Desc:  "磁头数",
	},
	{28, 32}: {
		Color: 0x006600,
		Desc:  "隐藏扇区数",
	},
	{32, 36}: {
		Color: 0x000066,
		Desc:  "总扇区数(4bytes)",
	},
	{36, 37}: {
		Color: 0x336600,
		Desc:  "中断0x13的驱动器号",
	},
	{37, 38}: {
		Color: 0x660000,
		Desc:  "保留位",
	},
	{38, 39}: {
		Color: 0x660033,
		Desc:  "拓展引导标记",
	},
	{39, 43}: {
		Color: 0x330066,
		Desc:  "卷序列号",
	},
	{43, 54}: {
		Color: 0x006666,
		Desc:  "卷标",
	},
	{54, 62}: {
		Color: 0x003366,
		Desc:  "文件系统类型",
	},
	{62, 510}: {
		Color: 0x000000,
		Desc:  "引导扇区代码",
	},
	{510, 512}: {
		Color: 0x660000,
		Desc:  "结束标志",
	},
}
