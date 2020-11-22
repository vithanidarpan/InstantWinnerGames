# ************************************************************
# Sequel Pro SQL dump
# Version 5438
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 8.0.0-dmr)
# Database: instant_winner_games
# Generation Time: 2020-11-21 18:31:23 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table campaigns
# ------------------------------------------------------------

CREATE TABLE `campaigns` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `start_date` datetime(3) DEFAULT NULL,
  `end_date` datetime(3) DEFAULT NULL,
  `name` text,
  `description` text,
  `picture_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;



# Dump of table cities
# ------------------------------------------------------------

CREATE TABLE `cities` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `name` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;



# Dump of table gifts
# ------------------------------------------------------------

CREATE TABLE `gifts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` text,
  `campaign_id` int(11) DEFAULT NULL,
  `picture_id` int(11) DEFAULT NULL,
  `description` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;



# Dump of table instant_winner_games
# ------------------------------------------------------------

CREATE TABLE `instant_winner_games` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` text,
  `play_time` datetime(3) DEFAULT NULL,
  `start_date` datetime(3) DEFAULT NULL,
  `end_date` datetime(3) DEFAULT NULL,
  `gift_id` int(11) DEFAULT NULL,
  `campaign_id` int(11) DEFAULT NULL,
  `place_id` int(11) DEFAULT NULL,
  `description` text,
  `won` decimal(10,0) DEFAULT NULL,
  `zip_code` longtext,
  `max_distance` float DEFAULT NULL,
  `address` longtext,
  `longitude` float DEFAULT NULL,
  `latitude` float DEFAULT NULL,
  `city` longtext,
  `code` longtext,
  `mail` longtext,
  PRIMARY KEY (`id`),
  KEY `fk_instant_winner_games_place` (`place_id`),
  CONSTRAINT `fk_instant_winner_games_place` FOREIGN KEY (`place_id`) REFERENCES `places` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;



# Dump of table instant_winner_players
# ------------------------------------------------------------

CREATE TABLE `instant_winner_players` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `instant_winner_game_id` int(11) DEFAULT NULL,
  `ip_address` text,
  `fingerprint` text,
  `email` text,
  `time` datetime(3) DEFAULT NULL,
  `result` decimal(10,0) DEFAULT NULL,
  `gift_id` int(11) DEFAULT '0',
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_instant_winner_players_instant_winner_game` (`instant_winner_game_id`),
  KEY `idx_instant_winner_players_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_instant_winner_players_instant_winner_game` FOREIGN KEY (`instant_winner_game_id`) REFERENCES `instant_winner_games` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;



# Dump of table pictures
# ------------------------------------------------------------

CREATE TABLE `pictures` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` text,
  `mime_type` text,
  `data` blob,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;



# Dump of table places
# ------------------------------------------------------------

CREATE TABLE `places` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` text,
  `address` text,
  `zip_code` text,
  `city_id` int(11) DEFAULT NULL,
  `code` text,
  `mail` text,
  `latitude` double DEFAULT NULL,
  `longitude` double DEFAULT NULL,
  `max_distance` double DEFAULT NULL,
  `city` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;



# Dump of table random_draw_games
# ------------------------------------------------------------

CREATE TABLE `random_draw_games` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` text,
  `description` text,
  `gift_id` int(11) DEFAULT NULL,
  `start_date` datetime(3) DEFAULT NULL,
  `end_date` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;



# Dump of table random_draw_players
# ------------------------------------------------------------

CREATE TABLE `random_draw_players` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `random_draw_game_id` int(11) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `ip_address` text,
  `email` text,
  `time` datetime(3) DEFAULT NULL,
  `won` decimal(10,0) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;



# Dump of table users
# ------------------------------------------------------------

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` text,
  `email` text,
  `password` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
