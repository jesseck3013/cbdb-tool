CREATE TABLE BIOG_MAIN (
    "c_personid" INTEGER NOT NULL,
    "c_name" varchar(255) DEFAULT NULL /* Hanyu Pinyin full name; auto-generated: c_surname + " " + c_mingzi */,
    "c_name_chn" varchar(255) DEFAULT NULL /* Chinese full name; auto-generated: c_surname_chn + c_mingzi_chn (no space) */,
    "c_index_year" smallint(6) DEFAULT NULL,
    "c_index_year_type_code" varchar(255) DEFAULT NULL,
    "c_index_year_source_id" INTEGER(11) DEFAULT NULL,
    "c_female" smallint(6) DEFAULT NULL,
    "c_index_addr_id" INTEGER(11) DEFAULT 0,
    "c_index_addr_type_code" smallint(6) DEFAULT NULL,
    "c_ethnicity_code" smallint(6) DEFAULT NULL,
    "c_household_status_code" smallint(6) DEFAULT NULL,
    "c_tribe" varchar(255) DEFAULT NULL,
    "c_birthyear" smallint(6) DEFAULT NULL,
    "c_by_nh_code" smallint(6) DEFAULT NULL,
    "c_by_nh_year" smallint(6) DEFAULT NULL,
    "c_by_range" smallint(6) DEFAULT NULL,
    "c_deathyear" smallint(6) DEFAULT NULL,
    "c_dy_nh_code" smallint(6) DEFAULT NULL,
    "c_dy_nh_year" smallint(6) DEFAULT NULL,
    "c_dy_range" smallint(6) DEFAULT NULL,
    "c_death_age" smallint(6) DEFAULT NULL,
    "c_death_age_range" smallint(6) DEFAULT NULL,
    "c_fl_earliest_year" smallint(6) DEFAULT NULL,
    "c_fl_ey_nh_code" smallint(6) DEFAULT NULL,
    "c_fl_ey_nh_year" smallint(6) DEFAULT NULL,
    "c_fl_ey_notes" TEXT DEFAULT NULL,
    "c_fl_latest_year" smallint(6) DEFAULT NULL,
    "c_fl_ly_nh_code" smallint(6) DEFAULT NULL,
    "c_fl_ly_nh_year" smallint(6) DEFAULT NULL,
    "c_fl_ly_notes" TEXT DEFAULT NULL,
    "c_surname" varchar(255) DEFAULT NULL /* Hanyu Pinyin romanization of the person's surname; auto-generated from c_surname_chn via pinyin lookup table */,
    "c_surname_chn" varchar(255) DEFAULT NULL /* Chinese surname; split from c_name_chn by matching longest known surname in pinyin table */,
    "c_mingzi" varchar(255) DEFAULT NULL /* Hanyu Pinyin romanization of the person's given name (excluding surname); auto-generated from c_mingzi_chn */,
    "c_mingzi_chn" varchar(255) DEFAULT NULL /* Chinese given name (excluding surname); remainder of c_name_chn after surname extraction */,
    "c_dy" smallint(6) DEFAULT NULL,
    "c_choronym_code" smallint(6) DEFAULT NULL,
    "c_notes" TEXT DEFAULT NULL,
    "c_by_intercalary" smallint(6) DEFAULT NULL,
    "c_dy_intercalary" smallint(6) DEFAULT NULL,
    "c_by_month" smallint(6) DEFAULT NULL,
    "c_dy_month" smallint(6) DEFAULT NULL,
    "c_by_day" smallint(6) DEFAULT NULL,
    "c_dy_day" smallint(6) DEFAULT NULL,
    "c_by_day_gz" smallint(6) DEFAULT NULL,
    "c_dy_day_gz" smallint(6) DEFAULT NULL,
    "c_surname_proper" varchar(255) DEFAULT NULL /* Surname in the person's native language (non-Chinese), if applicable; user-editable */,
    "c_mingzi_proper" varchar(255) DEFAULT NULL /* Given name in the person's native language (non-Chinese, excluding surname), if applicable; user-editable */,
    "c_name_proper" varchar(255) DEFAULT NULL /* Full name in the person's native language; auto-generated: c_mingzi_proper + " " + c_surname_proper (given-name-first order) */,
    "c_surname_rm" varchar(255) DEFAULT NULL /* Non-Pinyin romanization of the person's surname (e.g. Wade-Giles, McCune-Reischauer), if applicable; user-editable */,
    "c_mingzi_rm" varchar(255) DEFAULT NULL /* Non-Pinyin romanization of the person's given name (excluding surname), if applicable; user-editable */,
    "c_name_rm" varchar(255) DEFAULT NULL /* Non-Pinyin romanized full name; auto-generated: c_mingzi_rm + " " + c_surname_rm (given-name-first order) */,
    "c_created_by" varchar(255) DEFAULT NULL,
    "c_modified_by" varchar(255) DEFAULT NULL,
    "c_created_date" TEXT DEFAULT NULL,
    "c_modified_date" TEXT DEFAULT NULL,
    PRIMARY KEY ("c_personid")
);

CREATE TABLE KIN_DATA (
    "c_personid" INTEGER(11) NOT NULL,
    "c_kin_id" INTEGER(11) NOT NULL,
    "c_kin_code" smallint(6) NOT NULL,
    "c_source" INTEGER(11) DEFAULT NULL,
    "c_pages" varchar(255) DEFAULT NULL,
    "c_notes" TEXT DEFAULT NULL,
    "c_autogen_notes" TEXT DEFAULT NULL,
    "c_created_by" varchar(255) DEFAULT NULL,
    "c_modified_by" varchar(255) DEFAULT NULL,
    "c_created_date" TEXT DEFAULT NULL,
    "c_modified_date" TEXT DEFAULT NULL,
    PRIMARY KEY ("c_kin_code", "c_kin_id", "c_personid")
);

CREATE TABLE KINSHIP_CODES (
    "c_kincode" smallint(6) NOT NULL,
    "c_kin_pair1" smallint(6) NOT NULL DEFAULT 0,
    "c_kin_pair2" smallint(6) NOT NULL DEFAULT 0,
    "c_kin_pair_notes" varchar(255) DEFAULT NULL,
    "c_kinrel_chn" varchar(255) NOT NULL DEFAULT '',
    "c_kinrel" varchar(255) NOT NULL DEFAULT '',
    "c_kinrel_alt" varchar(255) DEFAULT NULL,
    "c_pick_sorting" smallint(6) DEFAULT NULL,
    "c_upstep" smallint(6) NOT NULL DEFAULT 0,
    "c_dwnstep" smallint(6) NOT NULL DEFAULT 0,
    "c_marstep" smallint(6) NOT NULL DEFAULT 0,
    "c_colstep" smallint(6) NOT NULL DEFAULT 0,
    "c_kinrel_simplified" varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY ("c_kincode")
);

CREATE TABLE DYNASTIES (
    "c_dy" smallint(6) NOT NULL,
    "c_dynasty" varchar(255) DEFAULT NULL,
    "c_dynasty_chn" varchar(255) DEFAULT NULL,
    "c_start" smallint(6) NOT NULL DEFAULT 0,
    "c_end" smallint(6) NOT NULL DEFAULT 0,
    "c_sort" smallint(6) DEFAULT NULL,
    PRIMARY KEY ("c_dy")
);

CREATE TABLE CHORONYM_CODES (
    "c_choronym_code" smallint(6) NOT NULL,
    "c_choronym_desc" varchar(255) DEFAULT NULL,
    "c_choronym_chn" varchar(255) DEFAULT NULL,
    PRIMARY KEY ("c_choronym_code")
);

CREATE TABLE ALTNAME_DATA (
    "c_personid" INTEGER(11) NOT NULL,
    "c_alt_name" varchar(255) DEFAULT NULL,
    "c_alt_name_chn" varchar(255) NOT NULL,
    "c_alt_name_type_code" smallint(6) NOT NULL,
    "c_sequence" smallint(6) DEFAULT 0,
    "c_source" INTEGER(11) DEFAULT NULL,
    "c_pages" varchar(255) DEFAULT NULL,
    "c_notes" TEXT DEFAULT NULL,
    "c_created_by" varchar(255) DEFAULT NULL,
    "c_modified_by" varchar(255) DEFAULT NULL,
    "c_created_date" TEXT DEFAULT NULL,
    "c_modified_date" TEXT DEFAULT NULL,
    PRIMARY KEY ("c_alt_name_chn", "c_alt_name_type_code", "c_personid")
);

CREATE TABLE ALTNAME_CODES (
    "c_name_type_code" smallint(6) NOT NULL,
    "c_name_type_desc" varchar(255) DEFAULT NULL,
    "c_name_type_desc_chn" varchar(255) DEFAULT NULL,
    PRIMARY KEY ("c_name_type_code")
);

CREATE TABLE STATUS_DATA (
    "c_personid" INTEGER(11) NOT NULL,
    "c_sequence" smallint(6) NOT NULL,
    "c_status_code" smallint(6) NOT NULL,
    "c_firstyear" smallint(6) DEFAULT NULL,
    "c_fy_nh_code" smallint(6) DEFAULT NULL,
    "c_fy_nh_year" smallint(6) DEFAULT NULL,
    "c_fy_range" smallint(6) DEFAULT NULL,
    "c_lastyear" smallint(6) DEFAULT NULL,
    "c_ly_nh_code" smallint(6) DEFAULT NULL,
    "c_ly_nh_year" smallint(6) DEFAULT NULL,
    "c_ly_range" smallint(6) DEFAULT NULL,
    "c_supplement" varchar(255) DEFAULT NULL,
    "c_source" INTEGER(11) DEFAULT NULL,
    "c_pages" varchar(255) DEFAULT NULL,
    "c_notes" TEXT DEFAULT NULL,
    "c_created_by" varchar(255) DEFAULT NULL,
    "c_modified_by" varchar(255) DEFAULT NULL,
    "c_created_date" TEXT DEFAULT NULL,
    "c_modified_date" TEXT DEFAULT NULL,
    PRIMARY KEY ("c_personid", "c_sequence", "c_status_code")
);

CREATE TABLE STATUS_CODES (
    "c_status_code" smallint(6) NOT NULL,
    "c_status_desc" varchar(255) NOT NULL DEFAULT '',
    "c_status_desc_chn" varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY ("c_status_code")
);

CREATE TABLE ENTRY_DATA (
    "c_personid" INTEGER(11) NOT NULL,
    "c_entry_code" smallint(6) NOT NULL,
    "c_sequence" smallint(6) NOT NULL,
    "c_exam_rank" varchar(255) DEFAULT NULL,
    "c_kin_code" smallint(6) NOT NULL,
    "c_kin_id" INTEGER(11) NOT NULL,
    "c_assoc_code" smallint(6) NOT NULL,
    "c_assoc_id" INTEGER(11) NOT NULL,
    "c_year" smallint(6) NOT NULL,
    "c_age" smallint(6) DEFAULT NULL,
    "c_entry_nh_id" smallint(6) DEFAULT NULL,
    "c_entry_nh_year" smallint(6) DEFAULT NULL,
    "c_entry_dy" smallint(6) DEFAULT NULL,
    "c_entry_range" smallint(6) DEFAULT NULL,
    "c_inst_code" smallint(6) NOT NULL DEFAULT 0,
    "c_inst_name_code" smallint(6) NOT NULL DEFAULT 0,
    "c_exam_field" varchar(255) DEFAULT NULL,
    "c_entry_addr_id" INTEGER(11) DEFAULT NULL,
    "c_parental_status_code" smallint(6) DEFAULT NULL,
    "c_attempt_count" smallint(6) DEFAULT NULL,
    "c_source" INTEGER(11) DEFAULT NULL,
    "c_pages" varchar(255) DEFAULT NULL,
    "c_notes" TEXT DEFAULT NULL,
    "c_posting_notes" varchar(255) DEFAULT NULL,
    "c_created_by" varchar(255) DEFAULT NULL,
    "c_modified_by" varchar(255) DEFAULT NULL,
    "c_created_date" TEXT DEFAULT NULL,
    "c_modified_date" TEXT DEFAULT NULL,
    PRIMARY KEY ("c_assoc_code", "c_assoc_id", "c_entry_code", "c_inst_code", "c_inst_name_code", "c_kin_code", "c_kin_id", "c_personid", "c_sequence", "c_year")
);

CREATE TABLE ENTRY_CODES (
    "c_entry_code" smallint(6) NOT NULL,
    "c_entry_desc" varchar(255) NOT NULL DEFAULT '',
    "c_entry_desc_chn" varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY ("c_entry_code")
);

CREATE TABLE ENTRY_CODE_TYPE_REL (
    "c_entry_code" smallint(6) NOT NULL,
    "c_entry_type" varchar(255) NOT NULL,
    PRIMARY KEY ("c_entry_code", "c_entry_type")
);

CREATE TABLE ENTRY_TYPES (
    "c_entry_type" varchar(255) NOT NULL,
    "c_entry_type_desc" varchar(255) NOT NULL DEFAULT '',
    "c_entry_type_desc_chn" varchar(255) NOT NULL DEFAULT '',
    "c_entry_type_parent_id" varchar(255) DEFAULT NULL,
    "c_entry_type_level" smallint(6) DEFAULT NULL,
    "c_entry_type_sortorder" smallint(6) DEFAULT NULL,
    PRIMARY KEY ("c_entry_type")
);

CREATE TABLE POSTING_DATA (
    "c_personid" INTEGER(11) DEFAULT NULL,
    "c_posting_id" INTEGER(11) NOT NULL,
    "c_created_by" varchar(255) DEFAULT NULL,
    "c_created_date" TEXT DEFAULT NULL,
    "c_modified_by" varchar(255) DEFAULT NULL,
    "c_modified_date" TEXT DEFAULT NULL,
    PRIMARY KEY ("c_posting_id")
);

CREATE TABLE POSTED_TO_OFFICE_DATA (
    "c_personid" INTEGER(11) DEFAULT NULL,
    "c_office_id" INTEGER(11) NOT NULL,
    "c_posting_id" INTEGER(11) NOT NULL,
    "c_sequence" smallint(6) DEFAULT NULL,
    "c_firstyear" smallint(6) DEFAULT NULL,
    "c_fy_nh_code" smallint(6) DEFAULT NULL,
    "c_fy_nh_year" smallint(6) DEFAULT NULL,
    "c_fy_range" smallint(6) DEFAULT NULL,
    "c_lastyear" smallint(6) DEFAULT NULL,
    "c_ly_nh_code" smallint(6) DEFAULT NULL,
    "c_ly_nh_year" smallint(6) DEFAULT NULL,
    "c_ly_range" smallint(6) DEFAULT NULL,
    "c_appt_code" smallint(6) NOT NULL DEFAULT 0,
    "c_assume_office_code" smallint(6) DEFAULT NULL,
    "c_inst_code" smallint(6) DEFAULT 0,
    "c_inst_name_code" smallint(6) DEFAULT 0,
    "c_source" INTEGER(11) DEFAULT NULL,
    "c_pages" varchar(255) DEFAULT NULL,
    "c_notes" TEXT DEFAULT NULL,
    "c_office_id_backup" INTEGER(11) DEFAULT NULL,
    "c_office_category_id" smallint(6) DEFAULT NULL,
    "c_fy_intercalary" smallint(6) DEFAULT NULL,
    "c_fy_month" smallint(6) DEFAULT NULL,
    "c_ly_intercalary" smallint(6) DEFAULT NULL,
    "c_ly_month" smallint(6) DEFAULT NULL,
    "c_fy_day" smallint(6) DEFAULT NULL,
    "c_ly_day" smallint(6) DEFAULT NULL,
    "c_fy_day_gz" smallint(6) DEFAULT NULL,
    "c_ly_day_gz" smallint(6) DEFAULT NULL,
    "c_dy" smallint(6) DEFAULT NULL,
    "c_created_by" varchar(255) DEFAULT NULL,
    "c_modified_by" varchar(255) DEFAULT NULL,
    "c_created_date" TEXT DEFAULT NULL,
    "c_modified_date" TEXT DEFAULT NULL,
    PRIMARY KEY ("c_office_id", "c_posting_id")
);

CREATE TABLE APPOINTMENT_CODES (
    "c_appt_code" smallint(6) NOT NULL,
    "c_appt_desc_chn" varchar(255) DEFAULT NULL,
    "c_appt_desc" varchar(255) DEFAULT NULL,
    "c_appt_desc_chn_alt" varchar(255) DEFAULT NULL,
    "c_appt_desc_alt" varchar(255) DEFAULT NULL,
    "c_notes" TEXT DEFAULT NULL,
    PRIMARY KEY ("c_appt_code")
);

CREATE TABLE OFFICE_CATEGORIES (
    "c_office_category_id" smallint(6) NOT NULL,
    "c_category_desc" varchar(255) DEFAULT NULL,
    "c_category_desc_chn" varchar(255) DEFAULT NULL,
    "c_notes" varchar(255) DEFAULT NULL,
    PRIMARY KEY ("c_office_category_id")
);

CREATE TABLE OFFICE_CODES (
    "c_office_id" INTEGER(11) NOT NULL,
    "c_dy" smallint(6) NOT NULL DEFAULT 0,
    "c_office_pinyin" varchar(255) DEFAULT NULL,
    "c_office_chn" varchar(255) DEFAULT NULL,
    "c_office_pinyin_alt" varchar(255) DEFAULT NULL,
    "c_office_chn_alt" varchar(255) DEFAULT NULL,
    "c_office_trans" varchar(255) DEFAULT NULL,
    "c_office_trans_alt" varchar(255) DEFAULT NULL,
    "c_source" INTEGER(11) DEFAULT NULL,
    "c_pages" varchar(255) DEFAULT NULL,
    "c_notes" TEXT DEFAULT NULL,
    PRIMARY KEY ("c_office_id")
);

CREATE TABLE TEXT_CODES (
    "c_textid" INTEGER(11) NOT NULL,
    "c_title_chn" varchar(255) DEFAULT NULL,
    "c_title" varchar(255) DEFAULT NULL,
    "c_title_trans" varchar(255) DEFAULT NULL,
    "c_text_type_id" varchar(128) DEFAULT NULL,
    "c_text_year" smallint(6) DEFAULT NULL,
    "c_text_nh_code" smallint(6) DEFAULT NULL,
    "c_text_nh_year" smallint(6) DEFAULT NULL,
    "c_text_range_code" smallint(6) DEFAULT NULL,
    "c_bibl_cat_code" smallint(6) DEFAULT 0,
    "c_extant" smallint(6) DEFAULT NULL,
    "c_text_country" smallint(6) DEFAULT NULL,
    "c_text_dy" smallint(6) DEFAULT NULL,
    "c_source" INTEGER(11) DEFAULT NULL,
    "c_pages" varchar(255) DEFAULT NULL,
    "c_url_api" varchar(255) DEFAULT NULL,
    "c_url_api_coda" varchar(255) DEFAULT NULL,
    "c_url_homepage" varchar(255) DEFAULT NULL,
    "c_notes" TEXT DEFAULT NULL,
    "c_title_alt_chn" varchar(255) DEFAULT NULL,
    "c_created_by" varchar(255) DEFAULT NULL,
    "c_modified_by" varchar(255) DEFAULT NULL,
    "c_created_date" TEXT DEFAULT NULL,
    "c_modified_date" TEXT DEFAULT NULL,
    PRIMARY KEY ("c_textid")
);


CREATE TABLE BIOG_TEXT_DATA (
    "c_textid" INTEGER(11) NOT NULL,
    "c_personid" INTEGER(11) NOT NULL,
    "c_role_id" smallint(6) NOT NULL,
    "c_year" smallint(6) DEFAULT NULL,
    "c_nh_code" smallint(6) DEFAULT NULL,
    "c_nh_year" smallint(6) DEFAULT NULL,
    "c_range_code" smallint(6) DEFAULT NULL,
    "c_source" INTEGER(11) DEFAULT NULL,
    "c_pages" varchar(255) DEFAULT NULL,
    "c_notes" TEXT DEFAULT NULL,
    "c_created_by" varchar(255) DEFAULT NULL,
    "c_modified_by" varchar(255) DEFAULT NULL,
    "c_created_date" TEXT DEFAULT NULL,
    "c_modified_date" TEXT DEFAULT NULL,
    PRIMARY KEY ("c_personid", "c_role_id", "c_textid")
);

CREATE TABLE ASSOC_DATA (
    "c_assoc_code" smallint(6) NOT NULL,
    "c_personid" INTEGER(11) NOT NULL,
    "c_kin_code" smallint(6) NOT NULL,
    "c_kin_id" INTEGER(11) NOT NULL,
    "c_assoc_id" INTEGER(11) NOT NULL,
    "c_assoc_kin_code" smallint(6) NOT NULL,
    "c_assoc_kin_id" INTEGER(11) NOT NULL,
    "c_tertiary_personid" INTEGER(11) DEFAULT NULL,
    "c_tertiary_type_notes" TEXT DEFAULT NULL,
    "c_assoc_count" smallint(6) NOT NULL DEFAULT 1,
    "c_sequence" smallint(6) DEFAULT 0,
    "c_assoc_first_year" smallint(6) NOT NULL DEFAULT -9999,
    "c_assoc_last_year" smallint(6) DEFAULT NULL,
    "c_source" INTEGER(11) DEFAULT NULL,
    "c_pages" varchar(255) DEFAULT NULL,
    "c_notes" TEXT DEFAULT NULL,
    "c_assoc_fy_nh_code" smallint(6) DEFAULT NULL,
    "c_assoc_fy_nh_year" smallint(6) DEFAULT NULL,
    "c_assoc_fy_range" smallint(6) DEFAULT NULL,
    "c_assoc_ly_nh_code" smallint(6) DEFAULT NULL,
    "c_assoc_ly_nh_year" smallint(6) DEFAULT NULL,
    "c_assoc_ly_range" smallint(6) DEFAULT NULL,
    "c_addr_id" INTEGER(11) DEFAULT NULL,
    "c_litgenre_code" smallint(6) DEFAULT NULL,
    "c_occasion_code" smallint(6) DEFAULT NULL,
    "c_topic_code" smallint(6) DEFAULT NULL,
    "c_inst_code" smallint(6) DEFAULT 0,
    "c_inst_name_code" smallint(6) DEFAULT 0,
    "c_text_title" varchar(255) NOT NULL DEFAULT '',
    "c_assoc_claimer_id" INTEGER(11) DEFAULT NULL,
    "c_assoc_fy_intercalary" smallint(6) DEFAULT NULL,
    "c_assoc_fy_month" smallint(6) DEFAULT NULL,
    "c_assoc_fy_day" smallint(6) DEFAULT NULL,
    "c_assoc_fy_day_gz" smallint(6) DEFAULT NULL,
    "c_assoc_ly_intercalary" smallint(6) DEFAULT NULL,
    "c_assoc_ly_month" smallint(6) DEFAULT NULL,
    "c_assoc_ly_day" smallint(6) DEFAULT NULL,
    "c_assoc_ly_day_gz" smallint(6) DEFAULT NULL,
    "c_created_by" varchar(255) DEFAULT NULL,
    "c_modified_by" varchar(255) DEFAULT NULL,
    "c_created_date" TEXT DEFAULT NULL,
    "c_modified_date" TEXT DEFAULT NULL,
    PRIMARY KEY ("c_assoc_code", "c_assoc_id", "c_assoc_kin_code", "c_assoc_kin_id", "c_kin_code", "c_kin_id", "c_personid", "c_text_title", "c_assoc_first_year")
);

CREATE TABLE ASSOC_CODES (
    "c_assoc_code" smallint(6) NOT NULL,
    "c_assoc_pair" smallint(6) DEFAULT NULL,
    "c_assoc_pair2" smallint(6) DEFAULT NULL,
    "c_assoc_desc" varchar(255) DEFAULT NULL,
    "c_assoc_desc_chn" varchar(255) DEFAULT NULL,
    "c_assoc_role_type" varchar(255) DEFAULT NULL,
    "c_sortorder" smallint(6) DEFAULT NULL,
    "c_example" varchar(255) DEFAULT NULL,
    PRIMARY KEY ("c_assoc_code")
);

CREATE TABLE ASSOC_CODE_TYPE_REL (
    "c_assoc_code" smallint(6) NOT NULL,
    "c_assoc_type_code" varchar(255) NOT NULL,
    PRIMARY KEY ("c_assoc_code", "c_assoc_type_code")
);

CREATE TABLE ASSOC_TYPES (
    "c_assoc_type_code" varchar(255) NOT NULL,
    "c_assoc_type_desc" varchar(255) DEFAULT NULL,
    "c_assoc_type_desc_chn" varchar(255) DEFAULT NULL,
    "c_assoc_type_parent_id" varchar(255) DEFAULT NULL,
    "c_assoc_type_level" smallint(6) DEFAULT NULL,
    "c_assoc_type_sortorder" smallint(6) DEFAULT NULL,
    "c_assoc_type_short_desc" varchar(255) DEFAULT NULL,
    PRIMARY KEY ("c_assoc_type_code")
);

CREATE TABLE BIOG_INST_DATA (
    "c_personid" INTEGER(11) NOT NULL,
    "c_inst_name_code" smallint(6) NOT NULL,
    "c_inst_code" smallint(6) NOT NULL,
    "c_bi_role_code" smallint(6) NOT NULL,
    "c_bi_begin_year" smallint(6) DEFAULT NULL,
    "c_bi_by_nh_code" smallint(6) DEFAULT NULL,
    "c_bi_by_nh_year" smallint(6) DEFAULT NULL,
    "c_bi_by_range" smallint(6) DEFAULT NULL,
    "c_bi_end_year" smallint(6) DEFAULT NULL,
    "c_bi_ey_nh_code" smallint(6) DEFAULT NULL,
    "c_bi_ey_nh_year" smallint(6) DEFAULT NULL,
    "c_bi_ey_range" smallint(6) DEFAULT NULL,
    "c_source" INTEGER(11) DEFAULT NULL,
    "c_pages" varchar(255) DEFAULT NULL,
    "c_notes" TEXT DEFAULT NULL,
    "c_created_by" varchar(255) DEFAULT NULL,
    "c_modified_by" varchar(255) DEFAULT NULL,
    "c_created_date" TEXT DEFAULT NULL,
    "c_modified_date" TEXT DEFAULT NULL,
    PRIMARY KEY ("c_bi_role_code", "c_inst_code", "c_inst_name_code", "c_personid")
);

CREATE TABLE SOCIAL_INSTITUTION_NAME_CODES (
    "c_inst_name_code" smallint(6) NOT NULL,
    "c_inst_name_hz" varchar(255) NOT NULL DEFAULT '',
    "c_inst_name_py" varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY ("c_inst_name_code")
);

CREATE TABLE SOCIAL_INSTITUTION_CODES (
    "c_inst_name_code" smallint(6) NOT NULL,
    "c_inst_code" smallint(6) NOT NULL,
    "c_inst_type_code" smallint(6) DEFAULT NULL,
    "c_inst_begin_year" smallint(6) DEFAULT NULL,
    "c_by_nianhao_code" smallint(6) DEFAULT NULL,
    "c_by_nianhao_year" smallint(6) DEFAULT NULL,
    "c_by_year_range" smallint(6) DEFAULT NULL,
    "c_inst_begin_dy" smallint(6) DEFAULT NULL,
    "c_inst_floruit_dy" smallint(6) DEFAULT NULL,
    "c_inst_first_known_year" smallint(6) DEFAULT NULL,
    "c_inst_end_year" smallint(6) DEFAULT NULL,
    "c_ey_nianhao_code" smallint(6) DEFAULT NULL,
    "c_ey_nianhao_year" smallint(6) DEFAULT NULL,
    "c_ey_year_range" smallint(6) DEFAULT NULL,
    "c_inst_end_dy" smallint(6) DEFAULT NULL,
    "c_inst_last_known_year" smallint(6) DEFAULT NULL,
    "c_source" INTEGER(11) DEFAULT NULL,
    "c_pages" varchar(255) DEFAULT NULL,
    "c_notes" TEXT DEFAULT NULL,
    PRIMARY KEY ("c_inst_code", "c_inst_name_code")
);

CREATE TABLE BIOG_INST_CODES (
    "c_bi_role_code" smallint(6) NOT NULL,
    "c_bi_role_desc" varchar(255) DEFAULT NULL,
    "c_bi_role_chn" varchar(255) DEFAULT NULL,
    "c_notes" varchar(255) DEFAULT NULL,
    PRIMARY KEY ("c_bi_role_code")
);
