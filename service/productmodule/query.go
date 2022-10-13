package productmodule

const (
	addProductQuery = `
	INSERT INTO product (
		name,
		price,
		discount,
		stock,
		created_at,
		created_by,
		updated_at,
		updated_by
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8
	) returning id
`
	getProductQuery = `
	SELECT
		name,
		price,
		discount,
		stock,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM
		product
	WHERE
		id=$1
`

	getProductAllQuery = `
	SELECT
		*
	FROM
		product
`

	updateProductQuery = `
	UPDATE
		product
	SET
		name=$1,
		price=$2,
		discount=$3,
		stock=$4,
		created_at=$5,
		created_by=$6,
		updated_at=$7,
		updated_by=$8
	WHERE
		id=$9
	returning id	
`
)
