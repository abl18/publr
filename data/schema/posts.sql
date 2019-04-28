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

CREATE TABLE IF NOT EXISTS `posts` (
    `id` integer AUTO_INCREMENT NOT NULL,
    `title` varchar(50) NOT NULL,
    `slug` varchar(50) NOT NULL,
    `html` text NOT NULL,
    `image` varchar(255),
    `published` boolean NOT NULL DEFAULT 0,
    `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `publishtime` timestamp NULL DEFAULT NULL,
    `updatetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_slug` (`slug`)
);

CREATE TABLE IF NOT EXISTS `post_sites` (
    `id` integer AUTO_INCREMENT NOT NULL,
    `post_slug` varchar(50) NOT NULL,
    `site_domain` varchar(100) NOT NULL,
    `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updatetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_post_sites` (`post_slug`, `site_domain`),
    CONSTRAINT `c_post_sites` FOREIGN KEY `fk_posts` (`post_slug`) REFERENCES `posts`(`slug`) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `post_authors` (
    `id` integer AUTO_INCREMENT NOT NULL,
    `post_slug` varchar(50) NOT NULL,
    `author_username` varchar(50) NOT NULL,
    `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updatetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_post_authors` (`post_slug`, `author_username`),
    CONSTRAINT `c_post_authors` FOREIGN KEY `fk_posts` (`post_slug`) REFERENCES `posts`(`slug`) ON UPDATE CASCADE ON DELETE CASCADE
);