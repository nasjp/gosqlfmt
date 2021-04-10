package main

var query = `SELECT *
FROM users1` /*hoge*/ /*SQL*/ /*hoge*/

func main() {
	query := `SELECT * FROM users2` /*SQL*/

	_ = query

	f1 := func() {
		query := `SELECT * FROM users3` /*SQL*/
		_ = query
	}

	_ = f1

	func() {
		query := `SELECT * FROM users4` /*SQL*/
		_ = query
	}()

	do1(`SELECT * FROM users5` /*SQL*/)

	h := &Hoge{
		sql: `SELECT * FROM users6`, /*SQL*/
	}
	h.do2(`SELECT * FROM users7` /*SQL*/)

	h.do3(`SELECT * FROM users8` /*SQL*/)

	do1(
		`SELECT * FROM users9`,
		/*SQL*/
	)

	do1(
		`SELECT * FROM users10`, /*SQL*/
	)

	/*SQL*/
	query2 := `SELECT * FROM users11`

	_ = query2

	query3 := `SELECT * FROM users12`
	/*SQL*/

	_ = query3

	query4 := []string{`SELECT * FROM users13`, `SELECT * FROM users14`} /*SQL*/

	_ = query4

	query5 := []string{
		`SELECT * FROM users15`,
		`SELECT * FROM users15`, /*SQL*/
	}

	_ = query5

	query6 := map[string]interface{}{
		`SELECT * FROM users17` /*SQL*/ : nil,
	}

	_ = query6

	query7 := map[interface{}]string{
		"": `SELECT * FROM users18`, /*SQL*/
	}

	_ = query7
}

type Hoge struct {
	sql string
}

func (*Hoge) run1() string {
	return `SELECT * FROM users19` /*SQL*/
}

func (*Hoge) run2() string {
	query := `SELECT * FROM users20` /*SQL*/
	return query
}

func do1(string) {}

func (*Hoge) do2(string) {}

func (Hoge) do3(string) {}
