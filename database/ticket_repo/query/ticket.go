package query

const InsertTicketSql = `
	insert into tickets (start_point, end_point, start_time, end_time, company, user_id) values ($1, $2, $3, $4, $5, $6);
`

const SelectAllTickets = `
	select * from tickets;
`

const SelectUsersByTicketId = `
	select name, surname, patronymic from service_user u join tickets t on u.id = t.user_id where t.id = $1;
`

const SelectAllInfoAboutTicketByIdSql = `
	select start_point, end_point, start_time, end_time, buy_time, company from tickets where id = $1;
`

const SelectAllInfoAboutUserSql = `
	select service_user.id, name, surname, patronymic from service_user join tickets t on service_user.id = t.user_id where t.id = $1;
`

const SelectAllDocumentsByUserId = `
	select id, doc_type, number from user_document where user_id = $1;
`

const DeleteTicketSql = `
	delete from tickets where id = $1;
`
