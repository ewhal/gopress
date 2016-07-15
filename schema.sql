CREATE TABLE `post` (
  `id` int(10) unsigned NOT NULL auto_increment,
  `title` longtext,
  `summary` longtext,
  `post` longtext,
  PRIMARY KEY (`id`)
);

CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL auto_increment,
  `email` longtext,
  `password` longtext,
  `author` longtext,
  `email` longtext,
  `level` longtext,
  PRIMARY KEY (`id`)
);
