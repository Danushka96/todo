package Models

type Table struct {
	Name string
	Fields []TableField `default:"[]"`
}

func New(name string) Table{
	return Table{name, make([]TableField,0)};
}

func (t *Table) SetTableFields(tf TableField) {
	t.Fields = append(t.Fields, tf)
}