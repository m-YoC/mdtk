package sub

import (
)

const (
	ACT_NIL = iota

	// Group A
	ACT_VERSION
	ACT_CMD_HELP
	ACT_MANUAL
	ACT_WRITE_CONFIG

	// Group B
	ACT_MAKE_LIB
	ACT_GROUPS
	ACT_TASK_HELP

	// Group C
	ACT_PATH
	ACT_DIR

	// Group D
	ACT_RUN
	ACT_SCRIPT
	ACT_RAW_SCRIPT
)

func EnumGroupA(vflag bool, chflag bool, mflag bool, write_config_flag bool) int {
	if vflag {
		return ACT_VERSION
	}
	if chflag {
		return ACT_CMD_HELP
	}
	if mflag {
		return ACT_MANUAL
	}
	if write_config_flag {
		return ACT_WRITE_CONFIG
	}
	return ACT_NIL
}

func EnumGroupB(groups_flag bool, task_help_flag bool, make_lib_flag bool) int {
	if make_lib_flag {
		return ACT_MAKE_LIB
	}
	if groups_flag {
		return ACT_GROUPS
	}
	if task_help_flag {
		return ACT_TASK_HELP
	}
	return ACT_NIL
}

func EnumGroupC_WritePath(path_flag bool, dir_flag bool) int {
	if dir_flag {
		return ACT_DIR
	}
	if path_flag {
		return ACT_PATH
	}
	return ACT_NIL
}

func EnumGroupD_RunOrWriteScript(lang_is_sub bool, script_flag bool, raw_script_flag bool) int {
	if lang_is_sub {
		return ACT_RAW_SCRIPT
	}
	
	if raw_script_flag {
		return ACT_RAW_SCRIPT
	}
	
	if script_flag {
		return ACT_SCRIPT
	}
	return ACT_RUN
}
