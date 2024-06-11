package fbufstat

import "encoding/json"

type Bufstat struct {
	N_itm     int `json:"n_itm"`
	Buff_size int `json:"buff_size"`
}

func New(nf int, bsize int) Bufstat {

	Ftable := Bufstat{nf, bsize}

	return Ftable

}

func (Class *Bufstat) SetStat(buf []byte) {

	var Ftable Bufstat

	json.Unmarshal(buf, &Ftable)

	Class.N_itm = Ftable.N_itm
	Class.Buff_size = Ftable.Buff_size

}

func (Class Bufstat) GetObj() Bufstat {

	return Class

}

func (Class *Bufstat) UpdateCnt() {

	Class.N_itm += 1

}

func (Class Bufstat) GetJSONObj() []byte {

	bt_arr, _ := json.Marshal(Class)

	return bt_arr

}

func (Class *Bufstat) UpdateSize(buf int) {

	Class.Buff_size += buf

}
