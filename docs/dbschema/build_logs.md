# build_logs

## Description

ビルドログテーブル

<details>
<summary><strong>Table Definition</strong></summary>

```sql
CREATE TABLE `build_logs` (
  `id` varchar(22) NOT NULL COMMENT 'ビルドログID',
  `application_id` varchar(22) NOT NULL COMMENT 'アプリケーションID',
  `result` varchar(20) NOT NULL COMMENT 'ビルド結果',
  `started_at` datetime(6) NOT NULL COMMENT 'ビルド開始日時',
  `finished_at` datetime(6) DEFAULT NULL COMMENT 'ビルド終了日時',
  PRIMARY KEY (`id`),
  KEY `fk_build_logs_application_id` (`application_id`),
  CONSTRAINT `fk_build_logs_application_id` FOREIGN KEY (`application_id`) REFERENCES `applications` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ビルドログテーブル'
```

</details>

## Columns

| Name | Type | Default | Nullable | Children | Parents | Comment |
| ---- | ---- | ------- | -------- | -------- | ------- | ------- |
| id | varchar(22) |  | false | [artifacts](artifacts.md) |  | ビルドログID |
| application_id | varchar(22) |  | false |  | [applications](applications.md) | アプリケーションID |
| result | varchar(20) |  | false |  |  | ビルド結果 |
| started_at | datetime(6) |  | false |  |  | ビルド開始日時 |
| finished_at | datetime(6) |  | true |  |  | ビルド終了日時 |

## Constraints

| Name | Type | Definition |
| ---- | ---- | ---------- |
| fk_build_logs_application_id | FOREIGN KEY | FOREIGN KEY (application_id) REFERENCES applications (id) |
| PRIMARY | PRIMARY KEY | PRIMARY KEY (id) |

## Indexes

| Name | Definition |
| ---- | ---------- |
| fk_build_logs_application_id | KEY fk_build_logs_application_id (application_id) USING BTREE |
| PRIMARY | PRIMARY KEY (id) USING BTREE |

## Relations

![er](build_logs.svg)

---

> Generated by [tbls](https://github.com/k1LoW/tbls)