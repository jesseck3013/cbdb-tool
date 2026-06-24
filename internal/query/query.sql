-- name: GetPersonByName :many
SELECT
bm.c_personid AS person_id,
bm.c_name AS name_en,
bm.c_name_chn AS name_ch,
bm.c_birthyear AS birth_year,
bm.c_deathyear AS death_year,
d.c_dynasty AS dynasty_en,
d.c_dynasty_chn AS dynasty_ch,
d.c_start AS dynasty_start,
d.c_end AS dynasty_end
FROM 
BIOG_MAIN bm 
join DYNASTIES d on bm.c_dy = d.c_dy 
WHERE bm.c_name_chn = CAST(sqlc.arg(name) AS text)
OR 
bm.c_name like CAST(sqlc.arg(name) AS text);

-- name: GetPersonBasicInfoByID :one
select 
bm.c_personid AS person_id,
bm.c_name AS name_en,
bm.c_name_chn As name_ch,
bm.c_mingzi AS mingzi,
bm.c_female AS female,
bm.c_birthyear as birth_year,
bm.c_deathyear as death_year,
d.c_dynasty as dynasty_en,
d.c_dynasty_chn as dynasty_ch,
cc.c_choronym_desc as choronym_en,
cc.c_choronym_chn as choronym_ch
from BIOG_MAIN bm 
left join DYNASTIES d on bm.c_dy = d.c_dy 
left join CHORONYM_CODES cc on bm.c_choronym_code = cc.c_choronym_code 
where bm.c_personid = ?;

-- name: GetAltnamesByPersonID :many
select 
ad.c_alt_name as altname_en,
ad.c_alt_name_chn as altname_ch,
ac.c_name_type_desc as name_type,
ac.c_name_type_desc_chn as name_type_ch
from 
ALTNAME_DATA ad 
join ALTNAME_CODES ac on ad.c_alt_name_type_code = ac.c_name_type_code 
where ad.c_personid = ?;

-- name: GetPersonKinShipByPersonID :many
select 
kd.c_kin_id as kin_id,
kc.c_kinrel as kin_rel_en, 
kc.c_kinrel_alt as kin_alt, 
kc.c_kinrel_chn as kin_rel_ch, 
bm.c_name as name,
bm.c_name_chn as name_ch
from KIN_DATA kd
join BIOG_MAIN bm on bm.c_personid = kd.c_kin_id 
join KINSHIP_CODES kc on kd.c_kin_code = kc.c_kincode 
where kd.c_personid = ?;

-- name: GetStatusByPersonID :many
SELECT 
sc.c_status_desc as status_en,
sc.c_status_desc_chn as status_ch
FROM
STATUS_DATA sd 
join STATUS_CODES sc on sd.c_status_code = sc.c_status_code 
where sd.c_personid = ?;

-- name: GetEntryByPersonID :many
SELECT
ed.c_age as age,
ed.c_year as year,
et.c_entry_type_desc as entry_type_en,
ec.c_entry_desc_chn as entry_ch,
et.c_entry_type_level as entry_type_level,
et.c_entry_type_desc_chn as entry_type_level_ch,
ed.c_exam_rank as exam_rank
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
ac.c_appt_desc as appt_en,
ac.c_appt_desc_chn as appt_cn,
oc2.c_office_chn as office_ch,
oc2.c_office_chn_alt as office_ch_alt,
ptod.c_firstyear as first_year,
ptod.c_lastyear as last_year,
tc.c_title_chn as title_ch
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
tc.c_title_chn as title_ch,
tc.c_text_year as text_year,
tc.c_source as source,
tc2.c_title_chn as source_title
FROM
BIOG_TEXT_DATA btd 
JOIN TEXT_CODES tc on tc.c_textid  = btd.c_textid 
LEFT JOIN TEXT_CODES tc2 on tc.c_source = tc2.c_textid 
WHERE btd.c_personid = ?;

-- name: GetAssociationByPersonID :many
SELECT
ad.c_assoc_id as assoc_id,
ad.c_text_title as text_title,
bm.c_name as name_en,
bm.c_name_chn as name_ch,
ad.c_assoc_first_year as assoc_first_year,
ad.c_assoc_last_year as assoc_last_year,
ac.c_assoc_desc_chn as assoc_cn,
ac.c_assoc_pair as assoc_pair,
ad.c_notes as notes,
ad.c_pages as pages,
tc.c_title_chn as title_ch,
t.c_assoc_type_desc as assoc_type_en,
t.c_assoc_type_desc_chn as assoc_type_ch,
t.c_assoc_type_short_desc as assoc_type_short
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
sinc.c_inst_name_hz as inst_name_hz,
sinc.c_inst_name_py as inst_name_py,
bic.c_bi_role_desc as bi_role_en,
bic.c_bi_role_chn as bi_role_ch,
bid.c_bi_begin_year as bi_begin_year,
bid.c_bi_end_year as bi_end_year,
tc.c_title as title,
tc.c_title_chn as title_chn,
bid.c_pages as pages
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
ac.c_name as name_en,
ac.c_name_chn as name_ch,
bac.c_addr_desc as addr_en,
bac.c_addr_desc_chn as addr_ch,
ac.c_admin_type as admin_type,
bad.c_firstyear as first_year,
bad.c_lastyear as last_year,
bad.c_notes as notes,
ac.x_coord as x_coord,
ac.y_coord as y_coord,
tc.c_title_chn as title_ch,
bad.c_pages as pages
FROM
BIOG_ADDR_DATA bad 
LEFT JOIN ADDR_CODES ac on ac.c_addr_id = bad.c_addr_id 
LEFT JOIN BIOG_ADDR_CODES bac on bac.c_addr_type = bad.c_addr_type 
LEFT JOIN TEXT_CODES tc on bad.c_source = tc.c_textid 
LEFT JOIN ADDR_BELONGS_DATA abd on abd.c_addr_id = bad.c_addr_id 
where bad.c_personid = ?
ORDER BY bad.c_firstyear;
