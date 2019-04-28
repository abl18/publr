-- Copyright 2019 Publr Authors.
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--     http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

CREATE TABLE IF NOT EXISTS `users` (
    `id` integer AUTO_INCREMENT NOT NULL,
    `email` varchar(50) NOT NULL,
    `username` varchar(50) NOT NULL,
    `fullname` varchar(50) NOT NULL,
    `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updatetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_username` (`username`),
    FULLTEXT `ft_users` (`username`, `fullname`)
);

CREATE TABLE IF NOT EXISTS `site_users` (
    `id` integer AUTO_INCREMENT NOT NULL,
    `site_domain` varchar(100) NOT NULL,
    `user_username` varchar(50) NOT NULL,
    `role` int NOT NULL,
    `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updatetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_site_users` (`user_username`, `site_domain`),
    CONSTRAINT `c_users` FOREIGN KEY `fk_users` (`user_username`) REFERENCES `users`(`username`) ON UPDATE CASCADE ON DELETE CASCADE
);