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
left join DYNASTIES d on bm.c_dy = d.c_dy 
left join CHORONYM_CODES cc on bm.c_choronym_code = cc.c_choronym_code 
where bm.c_personid = ?;

-- name: GetAltnamesByPersonID :many
select 
ad.c_alt_name ,
ad.c_alt_name_chn ,
ac.c_name_type_desc,
ac.c_name_type_desc_chn 
from 
ALTNAME_DATA ad 
join ALTNAME_CODES ac on ad.c_alt_name_type_code = ac.c_name_type_code 
where ad.c_personid = ?;

-- name: GetPersonKinShipByPersonID :many
select 
kd.c_kin_id,
kc.c_kinrel, 
kc.c_kinrel_alt, 
kc.c_kinrel_chn, 
bm.c_name,
bm.c_name_chn 
from KIN_DATA kd
join BIOG_MAIN bm on bm.c_personid = kd.c_kin_id 
join KINSHIP_CODES kc on kd.c_kin_code = kc.c_kincode 
where kd.c_personid = ?;

-- name: GetStatusByPersonID :many
SELECT 
sc.c_status_desc ,
sc.c_status_desc_chn 
FROM
STATUS_DATA sd 
join STATUS_CODES sc on sd.c_status_code = sc.c_status_code 
where sd.c_personid = ?;

-- name: GetEntryByPersonID :many
SELECT
ed.c_age,
ed.c_year,
et.c_entry_type_desc,
ec.c_entry_desc_chn,
et.c_entry_type_level,
et.c_entry_type_desc_chn,
ed.c_exam_rank
FROM
ENTRY_DATA ed 
--join NIAN_HAO nh on nh.c_nianhao_id = ed.c_entry_nh_id 
join ENTRY_CODES ec on ec.c_entry_code  = ed.c_entry_code 
join ENTRY_CODE_TYPE_REL ectr on ectr.c_entry_code = ec.c_entry_code 
join ENTRY_TYPES et on et.c_entry_type = ectr.c_entry_type 
where ed.c_personid = ?;

-- name: GetPostingByPersonID :many
WITH ptod AS (
	SELECT * FROM POSTED_TO_OFFICE_DATA WHERE c_personid = CAST(sqlc.arg(person_id) AS INTEGER)
)
SELECT 
ac.c_appt_desc,
ac.c_appt_desc_chn,
oc2.c_office_chn,
oc2.c_office_chn_alt,
ptod.c_firstyear,
ptod.c_lastyear,
tc.c_title_chn 
FROM 
POSTING_DATA pd 
join ptod on ptod.c_posting_id = pd.c_posting_id
join APPOINTMENT_CODES ac on ac.c_appt_code = ptod.c_appt_code 
LEFT join OFFICE_CATEGORIES oc on oc.c_office_category_id = ptod.c_office_category_id 
join OFFICE_CODES oc2 on oc2.c_office_id = ptod.c_office_id 
LEFT JOIN TEXT_CODES tc on ptod.c_source = tc.c_textid 
where pd.c_personid = CAST(sqlc.arg(person_id) AS INTEGER)
ORDER BY ptod.c_firstyear;

-- name: GetTextByPersonID :many
SELECT
tc.c_title_chn,
tc.c_text_year,
tc.c_source,
tc2.c_title_chn 
FROM
BIOG_TEXT_DATA btd 
JOIN TEXT_CODES tc on tc.c_textid  = btd.c_textid 
LEFT JOIN TEXT_CODES tc2 on tc.c_source = tc2.c_textid 
WHERE btd.c_personid = ?;

-- name: GetAssociationByPersonID :many
SELECT
ad.c_assoc_id ,
bm.c_name_chn,
ad.c_assoc_first_year,
ad.c_assoc_last_year,
ac.c_assoc_desc_chn,
ac.c_assoc_pair,
ad.c_notes,
ad.c_pages,
tc.c_title_chn,
t.c_assoc_type_desc,
t.c_assoc_type_desc_chn,
t.c_assoc_type_short_desc
FROM 
ASSOC_DATA ad 
LEFT JOIN BIOG_MAIN bm on ad.c_assoc_id = bm.c_personid 
LEFT JOIN ASSOC_CODES ac on ad.c_assoc_code  = ac.c_assoc_code 
LEFT JOIN TEXT_CODES tc  on ad.c_source = tc.c_textid 
LEFT JOIN ASSOC_CODE_TYPE_REL actr on actr.c_assoc_code = ac.c_assoc_code 
LEFT JOIN ASSOC_TYPES t on t.c_assoc_type_code = actr.c_assoc_type_code 
WHERE ad.c_personid = ?;

-- name: GetInstitutionByPersonID :many
SELECT 
sinc.c_inst_name_hz,
sinc.c_inst_name_py,
bic.c_bi_role_desc,
bic.c_bi_role_chn,
bid.c_bi_begin_year,
bid.c_bi_end_year,
tc.c_title,
tc.c_title_chn,
bid.c_pages
FROM
BIOG_INST_DATA bid 
LEFT JOIN BIOG_MAIN bm on bm.c_personid = bid.c_personid 
LEFT JOIN SOCIAL_INSTITUTION_NAME_CODES sinc on bid.c_inst_name_code = sinc.c_inst_name_code 
LEFT JOIN SOCIAL_INSTITUTION_CODES sic on sic.c_inst_code = bid.c_inst_code
--LEFT JOIN SOCIAL_INSTITUTION_ADDR sia on sia.c_inst_code = bid.c_inst_code
LEFT JOIN BIOG_INST_CODES bic on bic.c_bi_role_code = bid.c_bi_role_code 
LEFT JOIN TEXT_CODES tc on tc.c_textid  = bid.c_source 
where bid.c_personid = ?
ORDER BY bid.c_bi_begin_year;;

-- name: GetPlaceByPersonID :many
SELECT
ac.c_name,
ac.c_name_chn,
bac.c_addr_desc,
bac.c_addr_desc_chn,
ac.c_admin_type,
bad.c_firstyear,
bad.c_lastyear,
bad.c_notes,
ac.x_coord,
ac.y_coord,
tc.c_title_chn,
bad.c_pages
FROM
BIOG_ADDR_DATA bad 
LEFT JOIN ADDR_CODES ac on ac.c_addr_id = bad.c_addr_id 
LEFT JOIN BIOG_ADDR_CODES bac on bac.c_addr_type = bad.c_addr_type 
LEFT JOIN TEXT_CODES tc on bad.c_source = tc.c_textid 
LEFT JOIN ADDR_BELONGS_DATA abd on abd.c_addr_id = bad.c_addr_id 
where bad.c_personid = ?
ORDER BY bad.c_firstyear;
