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

INSERT INTO `users` (`email`, `username`, `fullname`) VALUES ("userdemo@mysites.site", "userdemo", "User Demo");
INSERT INTO `site_users` (user_username, site_domain, role) VALUES ("userdemo", "mysites.site", "0");

COMMIT;

START TRANSACTION;

INSERT INTO `users` (`email`, `username`, `fullname`) VALUES ("authordemo@mysites.site", "authordemo", "Author Demo");
INSERT INTO `site_users` (user_username, site_domain, role) VALUES ("authordemo", "mysites.site", "1");

COMMIT;

START TRANSACTION;

INSERT INTO `users` (`email`, `username`, `fullname`) VALUES ("ownerdemo@mysites.site", "ownerdemo", "Owner Demo");
INSERT INTO `site_users` (user_username, site_domain, role) VALUES ("ownerdemo", "mysites.site", "3");

COMMIT;