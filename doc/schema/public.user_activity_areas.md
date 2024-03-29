# public.user_activity_areas

## Description

## Columns

| Name | Type | Default | Nullable | Children | Parents | Comment |
| ---- | ---- | ------- | -------- | -------- | ------- | ------- |
| id | bigint | nextval('user_activity_areas_id_seq'::regclass) | false |  |  |  |
| user_id | bigint |  | false |  | [public.users](public.users.md) |  |
| prefecture_id | bigint |  | false |  | [public.prefectures](public.prefectures.md) |  |
| created_at | timestamp with time zone |  | false |  |  |  |
| updated_at | timestamp with time zone |  | false |  |  |  |

## Constraints

| Name | Type | Definition |
| ---- | ---- | ---------- |
| user_activity_areas_user_id_fkey | FOREIGN KEY | FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE |
| user_activity_areas_prefecture_id_fkey | FOREIGN KEY | FOREIGN KEY (prefecture_id) REFERENCES prefectures(id) ON DELETE CASCADE |
| user_activity_areas_pkey | PRIMARY KEY | PRIMARY KEY (id) |
| user_activity_areas_user_id_prefecture_id_key | UNIQUE | UNIQUE (user_id, prefecture_id) |

## Indexes

| Name | Definition |
| ---- | ---------- |
| user_activity_areas_pkey | CREATE UNIQUE INDEX user_activity_areas_pkey ON public.user_activity_areas USING btree (id) |
| user_activity_areas_user_id_prefecture_id_key | CREATE UNIQUE INDEX user_activity_areas_user_id_prefecture_id_key ON public.user_activity_areas USING btree (user_id, prefecture_id) |
| user_activity_areas_user_id_idx | CREATE INDEX user_activity_areas_user_id_idx ON public.user_activity_areas USING btree (user_id) |
| user_activity_areas_prefecture_id_idx | CREATE INDEX user_activity_areas_prefecture_id_idx ON public.user_activity_areas USING btree (prefecture_id) |

## Relations

![er](public.user_activity_areas.svg)

---

> Generated by [tbls](https://github.com/k1LoW/tbls)
