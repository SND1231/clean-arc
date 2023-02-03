CREATE TABLE `workers` (
  `worker_id` varchar(255) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `status` int(1) NOT NULL COMMENT '0:非活性, 1:活性',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`worker_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `auth` (
  `worker_id` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `change_password_count` int(1) NOT NULL,
  `fail_auth_count` int(1) NOT NULL,
  `status` int(1) NOT NULL COMMENT '0:削除, 1:利用可能, 2:一時停止',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  foreign key (`worker_id`)  references workers(`worker_id`),
  PRIMARY KEY (`worker_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
