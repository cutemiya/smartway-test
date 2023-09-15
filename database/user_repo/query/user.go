package query

const InsertUserSQL = `
	insert into service_user(name, surname, patronymic) values ($1, $2, $3);
`

const CheckUserSql = `
	select count(*) from service_user where id = $1;
`

const DeleteUserSql = `
	delete from service_user where id = $1;
`

const DeleteUserTicketsSql = `
	delete from tickets where user_id = $1;
`

const DeleteUserDocumentsSql = `
	delete from user_document where user_id = $1
`

const SelectPreviouslyFlightsSql = `
	select id, start_point, end_point, start_time, end_time, buy_time, company from tickets 
	where buy_time < $1::timestamp and end_time < $2::timestamp and user_id = $3;
`

const SelectNotFulFilledFlightsSql = `
	select id, start_point, end_point, start_time, end_time, buy_time, company from tickets 
	where buy_time > $1::timestamp and end_time > $2::timestamp and user_id = $3;
`

const SelectPreviouslyAndNotFulFilledFlightsSql = `
	select id, start_point, end_point, start_time, end_time, buy_time, company from tickets 
	where buy_time > $1::timestamp and end_time < $2::timestamp and user_id = $3;
`
