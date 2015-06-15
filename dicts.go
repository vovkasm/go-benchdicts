package dicts

type Key struct {
	A, B, C string
}

func (k Key) StringKey() string {
	return k.A + "|" + k.B + "|" + k.C
}

type Dict interface {
	Set(k Key, v int)
	Get(k Key) (int, bool)
}

type Dict1 struct {
	data map[Key]int
}

func NewDict1() Dict {
	d := &Dict1{
		data: make(map[Key]int, 0),
	}
	return d
}

func (d *Dict1) Set(k Key, v int) {
	d.data[k] = v
}

func (d *Dict1) Get(k Key) (int, bool) {
	v, ok := d.data[k]
	return v, ok
}

type Dict2 struct {
	data map[string]map[string]map[string]int
}

func NewDict2() Dict {
	d := &Dict2{
		data: make(map[string]map[string]map[string]int, 0),
	}
	return d
}

func (d *Dict2) Set(k Key, v int) {
	if av, ok := d.data[k.A]; ok {
		if bv, ok := av[k.B]; ok {
			bv[k.C] = v
		} else {
			av[k.B] = map[string]int{k.C: v}
		}
	} else {
		d.data[k.A] = map[string]map[string]int{k.B: {k.C: v}}
	}
}

func (d *Dict2) Get(k Key) (int, bool) {
	v, ok := d.data[k.A][k.B][k.C]
	return v, ok
}

type Dict3 struct {
	data map[string]int
}

func NewDict3() Dict {
	d := &Dict3{data: make(map[string]int, 0)}
	return d
}

func (d *Dict3) Set(k Key, v int) {
	d.data[k.StringKey()] = v
}

func (d *Dict3) Get(k Key) (int, bool) {
	v, ok := d.data[k.StringKey()]
	return v, ok
}
