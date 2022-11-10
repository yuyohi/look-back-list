CREATE TABLE IF NOT EXISTS tasks (
    `id` VARCHAR(32) NOT NULL, -- 後々user_idを追加し、外部keyにする
    `title` TEXT NOT NULL,
    `detail` TEXT NOT NULL,
    `estimated_time` INT NOT NULL,
    `actual_time` INT NOT NULL,
    `is_done` boolean NOT NULL,
    `created_at` datetime(6) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
