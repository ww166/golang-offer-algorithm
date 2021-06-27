module x.localhost/sample1

go 1.16

replace x.localhost/collections => ../collections

require x.localhost/examples v0.0.0-00010101000000-000000000000

replace x.localhost/examples => ../examples
