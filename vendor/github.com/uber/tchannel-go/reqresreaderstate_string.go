// generated by stringer -type=reqResReaderState; DO NOT EDIT

package tchannel

import "fmt"

const _reqResReaderState_name = "reqResReaderPreArg1reqResReaderPreArg2reqResReaderPreArg3reqResReaderComplete"

var _reqResReaderState_index = [...]uint8{0, 19, 38, 57, 77}

func (i reqResReaderState) String() string {
	if i < 0 || i+1 >= reqResReaderState(len(_reqResReaderState_index)) {
		return fmt.Sprintf("reqResReaderState(%d)", i)
	}
	return _reqResReaderState_name[_reqResReaderState_index[i]:_reqResReaderState_index[i+1]]
}
