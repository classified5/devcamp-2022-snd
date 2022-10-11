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
		name=$1,
		image_url=$2,
		description=$3,
		max_weight=$4,
		created_at=$5,
		created_by=$6,
		updated_at=$7,
		updated_by=$8
	WHERE
		id=$9
	returning id	
`

	deleteShipperQuery = `
	DELETE FROM 
		shipper 
	WHERE 
		id=$1 
	returning id
	`
)
