-- +migrate Up
ALTER TABLE `websites`
    DROP FOREIGN KEY `fk_websites_environment_id`,
    DROP INDEX `uk_environment_id`;

ALTER TABLE `build_logs`
    DROP FOREIGN KEY `fk_build_logs_environment_id`;

ALTER TABLE `environments`
    MODIFY COLUMN `id` varchar(22) COLLATE 'utf8mb4_general_ci' NOT NULL COMMENT 'ブランチID' FIRST,
    RENAME TO `branches`;

ALTER TABLE `build_logs`
    CHANGE `environment_id` `branch_id` varchar(22) COLLATE 'utf8mb4_general_ci' NOT NULL COMMENT 'ブランチID' AFTER `finished_at`,
    ADD CONSTRAINT `fk_build_logs_branch_id` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`) ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE `websites`
    CHANGE `environment_id` `branch_id` varchar(22) COLLATE 'utf8mb4_general_ci' NOT NULL COMMENT 'ブランチID' AFTER `updated_at`,
    ADD UNIQUE `uk_branch_id` (`branch_id`),
    ADD CONSTRAINT `fk_websites_branch_id` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`) ON UPDATE RESTRICT ON DELETE RESTRICT;

-- +migrate Down
ALTER TABLE `websites`
    DROP FOREIGN KEY `fk_websites_branch_id`,
    DROP INDEX `uk_branch_id`,
    CHANGE `branch_id` `environment_id` varchar(22) COLLATE 'utf8mb4_general_ci' NOT NULL COMMENT '環境ID' AFTER `updated_at`,
    ADD UNIQUE `uk_environment_id` (`environment_id`),
    ADD CONSTRAINT `fk_websites_environment_id` FOREIGN KEY (`environment_id`) REFERENCES `branches` (`id`) ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE `build_logs`
    DROP FOREIGN KEY `fk_build_logs_branch_id`,
    CHANGE `branch_id` `environment_id` varchar(22) COLLATE 'utf8mb4_general_ci' NULL COMMENT '環境ID' AFTER `finished_at`,
    ADD CONSTRAINT `fk_build_logs_environment_id` FOREIGN KEY (`environment_id`) REFERENCES `branches` (`id`) ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE `branches`
    MODIFY COLUMN `id` varchar(22) COLLATE 'utf8mb4_general_ci' NULL COMMENT '環境ID' FIRST,
    RENAME TO `environments`;