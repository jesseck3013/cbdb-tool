-- name: GetPersonByID :one
select c_personid, c_name_chn
from BIOG_MAIN
where c_personid = ?;
