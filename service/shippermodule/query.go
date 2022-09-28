package shippermodule

const (
	addShipperQuery = `
	INSERT INTO shipper (
		name,
		image_url,
		description,
		max_weight,
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
	getShipperQuery = `
	SELECT
		name,
		image_url,
		description,
		max_weight,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM
		shipper
	WHERE
		id=$1
`

	getShipperAllQuery = `
	SELECT
		*
	FROM
		shipper
`

	updateShipperQuery = `
	UPDATE
		shipper
	SET
		%s
	WHERE
		id=%d
`
)
