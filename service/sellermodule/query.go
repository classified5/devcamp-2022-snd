package sellermodule

const (
	addSellerQuery = `
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
	getSellerQuery = `
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

	getSellerAllQuery = `
	SELECT
		*
	FROM
		seller
`
)
