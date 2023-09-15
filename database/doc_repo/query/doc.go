package query

const InsertDocumentSql = `
	insert into user_document (doc_type, number, user_id) values ($1, $2, $3);
`

const SelectAllDocumentsByUserId = `
	select id, doc_type, number from user_document where user_id = $1;
`

const DeleteDocumentSql = `
	delete from user_document where id = $1;
`
