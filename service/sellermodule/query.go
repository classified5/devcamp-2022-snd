package sellermodule

const (
	addShipperQuery = `
	INSERT INTO seller (
		name,
		password,
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
	) returning id
`
	getShipperQuery = `
	SELECT
		name,
		password,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM
		seller
	WHERE
		id=$1
`

	getShipperAllQuery = `
	SELECT
		*
	FROM
		seller
`
)
