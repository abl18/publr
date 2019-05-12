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

START TRANSACTION;

INSERT INTO `posts` (`title`, `slug`, `html`, `image`, `published`) VALUES ("My First Post", "my-first-posts", "<p>My First Post</p>", "image.png", true);
INSERT INTO `post_sites` (`post_slug`, `site_domain`) VALUES ("my-first-posts", "mysites.site");
INSERT INTO `post_authors` (`post_slug`, `author_username`) VALUES ("my-first-posts", "authordemo");

COMMIT;

START TRANSACTION;

INSERT INTO `posts` (`title`, `slug`, `html`, `image`, `published`) VALUES ("My Second Post", "my-second-posts", "<p>My Second Post</p>", "image.png", true);
INSERT INTO `post_sites` (`post_slug`, `site_domain`) VALUES ("my-second-posts", "mysites.site");
INSERT INTO `post_authors` (`post_slug`, `author_username`) VALUES ("my-second-posts", "authordemo");

COMMIT;