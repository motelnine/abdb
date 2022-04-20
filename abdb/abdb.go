package ABDB

import (
	"fmt"
)

type Where struct {
	Column   string
	Value    string
	Operator string
}

type Args struct {
	Type     string
	Schema   string
	Table    string
	Values   map[int]map[string]string
	Columns  []string
	Where    []Where
	Operator string
	Combiner string
}

type Params struct {
	Schema      string
	Function    string
	Args        string
}

func Call(args Params) {
	query := fmt.Sprintf(`SELECT * FROM "%s"."%s"('%s')`, args.Schema, args.Function, args.Args)
	Raw(query)
}

func constructInsert(args Args) {
	query := fmt.Sprintf(`INSERT INTO "%s"."%s"`, args.Schema, args.Table)
	Raw(query)
}

func constructSelect(args Args) {
	columns := ""

	if(args.Columns == nil) {
		columns = "*"
	} else {
		for k, v := range args.Columns {
			if (k > 0) {
				columns = columns + `,`
			}
			columns = columns + `"` + v + `"`
		}
	}

	query := fmt.Sprintf("SELECT %s FROM \"%s\".\"%s\"", columns, args.Schema, args.Table)
	where := groupWhere(args)

	if(where != "") {
		query = fmt.Sprintf("%s WHERE %s", query, where)
	}

	//query = Util.Cat(opener, where)
	Raw(query)
}

func groupIn(args Args) {
	// YOU BE HERE
}

func groupWhere(args Args) string {
	where := ""

	if(args.Combiner == "") {
		args.Combiner = "AND"
	}

	if(args.Where != nil) {
		i := 0

		if (args.Combiner == "") {
			args.Combiner = "AND"
		}

		for _, v := range args.Where {

			if (v.Operator == "") {
				v.Operator = "="
			}

			theLine := fmt.Sprintf("i:%d", i)
			fmt.Println(theLine)

			if (i > 0) {
				fmt.Println(args.Combiner)
				//where = fmt.Println(where, fmt.Sprintf(" %s ", args.Combiner))
				where = fmt.Sprintf("%s %s ", where, args.Combiner)
			}

			where = v.Column + v.Operator + v.Value
			i++
		}
	}
	return where
}

func Query(args Args) {
	switch args.Type {
		case "insert":
			constructInsert(args)
			break
		case "select":
			constructSelect(args)
			break
	}
}

func Raw(query string) {
	//fmt.Println(config.Reader)
	fmt.Println(query)
}

