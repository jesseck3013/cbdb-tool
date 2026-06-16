-- name: GetPersonByID :one
select 
bm.c_personid,
bm.c_name,
bm.c_name_chn,
bm.c_mingzi,
bm.c_female,
bm.c_birthyear,
bm.c_deathyear,
d.c_dynasty,
d.c_dynasty_chn,
cc.c_choronym_desc,
cc.c_choronym_chn 
from BIOG_MAIN bm 
join DYNASTIES d on bm.c_dy = d.c_dy 
join CHORONYM_CODES cc on bm.c_choronym_code = cc.c_choronym_code 
where bm.c_personid = ?;
